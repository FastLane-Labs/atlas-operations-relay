package auction

import (
	"context"
	"math/big"
	"sync"
	"time"

	"github.com/FastLane-Labs/atlas-operations-relay/config"
	"github.com/FastLane-Labs/atlas-operations-relay/contract"
	"github.com/FastLane-Labs/atlas-operations-relay/log"
	"github.com/FastLane-Labs/atlas-operations-relay/operation"
	"github.com/FastLane-Labs/atlas-operations-relay/relayerror"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	ErrAuctionAlreadyStarted    = relayerror.NewError(2202, "auction for this user operation has already started")
	ErrAuctionNotFound          = relayerror.NewError(2203, "auction not found")
	ErrUserOpFailedSimulation   = relayerror.NewError(2204, "user operation failed simulation")
	ErrSolverOpFailedSimulation = relayerror.NewError(2205, "solver operation failed simulation")
	ErrNotEnoughBondedBalance   = relayerror.NewError(2206, "not enough atlEth bonded balance")
)

var (
	ATLETH_BONDED_BALANCE_MULTI_REQ = big.NewInt(3)
)

type balanceOfBondedFn func(common.Address) (*big.Int, *relayerror.Error)

type Manager struct {
	ethClient *ethclient.Client
	config    *config.Config

	// Indexed by userOpHash
	auctions map[common.Hash]*Auction

	atlasDomainSeparator common.Hash

	balanceOfBonded balanceOfBondedFn

	mu sync.RWMutex
}

func NewManager(ethClient *ethclient.Client, config *config.Config, atlasDomainSeparator common.Hash, balanceOfBonded balanceOfBondedFn) *Manager {
	am := &Manager{
		ethClient:            ethClient,
		config:               config,
		auctions:             make(map[common.Hash]*Auction),
		atlasDomainSeparator: atlasDomainSeparator,
		balanceOfBonded:      balanceOfBonded,
	}

	go am.auctionsCleaner()
	return am
}

func (am *Manager) auctionsCleaner() {
	for range time.Tick(10 * time.Minute) {
		am.mu.Lock()
		for userOpHash, auction := range am.auctions {
			if !auction.open && time.Since(auction.createdAt) > time.Hour {
				delete(am.auctions, userOpHash)
			}
		}
		am.mu.Unlock()
	}
}

func (am *Manager) NewUserOperation(userOp *operation.UserOperation) (common.Hash, *relayerror.Error) {
	userOpHash, relayErr := userOp.Hash()
	if relayErr != nil {
		log.Info("failed to compute user operation hash", "err", relayErr.Message)
		return common.Hash{}, relayErr
	}

	if relayErr := userOp.Validate(am.ethClient, am.config.Contracts.Atlas, am.atlasDomainSeparator, am.config.Relay.Gas.MaxPerUserOperation); relayErr != nil {
		log.Info("invalid user operation", "err", relayErr.Message, "userOpHash", userOpHash.Hex())
		return common.Hash{}, relayErr
	}

	pData, err := contract.SimulatorAbi.Pack("simUserOperation", *userOp)
	if err != nil {
		log.Info("failed to pack user operation", "err", err, "userOpHash", userOpHash.Hex())
		return common.Hash{}, relayerror.ErrServerInternal
	}

	bData, err := am.ethClient.CallContract(context.Background(), ethereum.CallMsg{To: &am.config.Contracts.Simulator, Data: pData}, nil)
	if err != nil {
		log.Info("failed to call simulator contract", "err", err, "userOpHash", userOpHash.Hex())
		return common.Hash{}, relayerror.ErrServerInternal
	}

	validOp, err := contract.SimulatorAbi.Unpack("simUserOperation", bData)
	if err != nil {
		log.Info("failed to unpack simUserOperation return data", "err", err, "userOpHash", userOpHash.Hex())
		return common.Hash{}, relayerror.ErrServerInternal
	}

	if !validOp[0].(bool) {
		log.Info("user operation failed simulation", "userOpHash", userOpHash.Hex())
		return common.Hash{}, ErrUserOpFailedSimulation
	}

	am.mu.Lock()
	defer am.mu.Unlock()

	if _, alreadyStarted := am.auctions[userOpHash]; alreadyStarted {
		log.Info("auction for this user operation has already started", "userOpHash", userOpHash.Hex())
		return common.Hash{}, ErrAuctionAlreadyStarted
	}

	am.auctions[userOpHash] = NewAuction(am.config.Relay.Auction.Duration, userOp, userOpHash)
	return userOpHash, nil
}

func (am *Manager) GetSolverOperations(userOpHash common.Hash, completionChan chan []*operation.SolverOperation) ([]*operation.SolverOperation, *relayerror.Error) {
	am.mu.RLock()
	defer am.mu.RUnlock()

	auction, ok := am.auctions[userOpHash]
	if !ok {
		log.Info("auction not found", "userOpHash", userOpHash.Hex())
		return nil, ErrAuctionNotFound
	}

	return auction.getSolverOps(completionChan)
}

func (am *Manager) getAuction(userOpHash common.Hash) *Auction {
	am.mu.RLock()
	defer am.mu.RUnlock()

	auction, ok := am.auctions[userOpHash]
	if !ok {
		log.Info("auction not found", "userOpHash", userOpHash.Hex())
		return nil
	}

	return auction
}

func (am *Manager) NewSolverOperation(solverOp *operation.SolverOperation) *relayerror.Error {
	auction := am.getAuction(solverOp.UserOpHash)
	if auction == nil {
		return ErrAuctionNotFound
	}

	relayErr := solverOp.Validate(auction.userOp, am.config.Contracts.Atlas, am.atlasDomainSeparator, am.config.Relay.Gas.MaxPerSolverOperation)
	if relayErr != nil {
		log.Info("invalid solver operation", "err", relayErr.Message, "userOpHash", auction.userOpHash.Hex())
		return relayErr
	}

	bondedBalance, relayErr := am.balanceOfBonded(solverOp.From)
	if relayErr != nil {
		return relayErr
	}

	// Bonded balance must be >= ATLETH_BONDED_BALANCE_MULTI_REQ * (gas * maxFeePerGas)
	atlEthRequired := new(big.Int).Mul(ATLETH_BONDED_BALANCE_MULTI_REQ, new(big.Int).Mul(solverOp.Gas, solverOp.MaxFeePerGas))
	if bondedBalance.Cmp(atlEthRequired) < 0 {
		log.Info("not enough bonded balance", "userOpHash", solverOp.UserOpHash.Hex(), "from", solverOp.From.Hex(), "bondedBalance", bondedBalance.String(), "atlEthRequired", atlEthRequired.String())
		return ErrNotEnoughBondedBalance
	}

	dAppOp := operation.GenerateSimulationDAppOperation(auction.userOp)

	pData, err := contract.SimulatorAbi.Pack("simSolverCall", *auction.userOp, *solverOp, *dAppOp)
	if err != nil {
		log.Info("failed to pack solver operation", "err", err, "userOpHash", auction.userOpHash.Hex())
		return relayerror.ErrServerInternal
	}

	bData, err := am.ethClient.CallContract(context.Background(), ethereum.CallMsg{To: &am.config.Contracts.Simulator, Data: pData}, nil)
	if err != nil {
		log.Info("failed to call simulator contract", "err", err, "userOpHash", auction.userOpHash.Hex())
		return relayerror.ErrServerInternal
	}

	validOp, err := contract.SimulatorAbi.Unpack("simSolverCall", bData)
	if err != nil {
		log.Info("failed to unpack simSolverCall return data", "err", err, "userOpHash", auction.userOpHash.Hex())
		return relayerror.ErrServerInternal
	}

	if !validOp[0].(bool) {
		log.Info("solver operation failed simulation", "userOpHash", auction.userOpHash.Hex())
		return ErrSolverOpFailedSimulation
	}

	return auction.addSolverOp(solverOp)
}
