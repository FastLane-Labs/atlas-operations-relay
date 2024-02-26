package auction

import (
	"context"
	"math/big"
	"sync"

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
	ErrToFieldMustBeAtlas       = relayerror.NewError(2200, "'to' field must be atlas contract address")
	ErrDeadlineExceeded         = relayerror.NewError(2201, "deadline exceeded")
	ErrAuctionAlreadyStarted    = relayerror.NewError(2202, "auction for this user operation has already started")
	ErrAuctionNotFound          = relayerror.NewError(2203, "auction not found")
	ErrUserOpFailedSimulation   = relayerror.NewError(2204, "user operation failed simulation")
	ErrSolverOpFailedSimulation = relayerror.NewError(2205, "solver operation failed simulation")
	ErrInvalidSignature         = relayerror.NewError(2206, "invalid signature")
)

var (
	solverGasLimit = big.NewInt(1000000)
)

type Manager struct {
	ethClient *ethclient.Client
	config    *config.Config

	// Indexed by userOpHash
	auctions map[common.Hash]*Auction

	mu sync.RWMutex
}

func NewManager(ethClient *ethclient.Client, config *config.Config) *Manager {
	return &Manager{
		ethClient: ethClient,
		config:    config,
	}
}

func (am *Manager) checkUserOperation(userOp *operation.UserOperation) *relayerror.Error {
	if userOp.To != am.config.Contracts.Atlas {
		return ErrToFieldMustBeAtlas
	}

	// gas limit check?

	currentBlock, err := am.ethClient.BlockNumber(context.Background())
	if err != nil {
		log.Info("failed to get current block number", "err", err)
		return relayerror.ErrServerInternal
	}

	if userOp.Deadline.Uint64() <= currentBlock {
		return ErrDeadlineExceeded
	}

	// signature check
	

	return nil
}

func (am *Manager) NewUserOperation(userOp *operation.UserOperation) (common.Hash, *relayerror.Error) {
	userOpHash, err := userOp.Hash()
	if err != nil {
		log.Info("failed to compute user operation hash", "err", err)
		return common.Hash{}, relayerror.ErrComputeUserOpHash.AddError(err)
	}

	if relayErr := am.checkUserOperation(userOp); relayErr != nil {
		log.Info("invalid user operation", "err", relayErr.Message, "userOpHash", userOpHash.Hex())
		return common.Hash{}, relayErr
	}

	pData, err := contract.SimulatorAbi.Pack("simUserOperation", *userOp)
	if err != nil {
		log.Info("failed to pack user operation", "err", err, "userOpHash", userOpHash.String())
		return common.Hash{}, relayerror.ErrServerInternal
	}

	bData, err := am.ethClient.CallContract(context.Background(), ethereum.CallMsg{To: &am.config.Contracts.Simulator, Data: pData}, nil)
	if err != nil {
		log.Info("failed to call simulator contract", "err", err, "userOpHash", userOpHash.String())
		return common.Hash{}, relayerror.ErrServerInternal
	}

	validOp, err := contract.SimulatorAbi.Unpack("simUserOperation", bData)
	if err != nil {
		log.Info("failed to unpack simUserOperation return data", "err", err, "userOpHash", userOpHash.String())
		return common.Hash{}, relayerror.ErrServerInternal
	}

	if !validOp[0].(bool) {
		log.Info("user operation failed simulation", "userOpHash", userOpHash.String())
		return common.Hash{}, ErrUserOpFailedSimulation
	}

	am.mu.Lock()
	defer am.mu.Unlock()

	if _, alreadyStarted := am.auctions[userOpHash]; alreadyStarted {
		log.Info("auction for this user operation has already started", "userOpHash", userOpHash.String())
		return common.Hash{}, ErrAuctionAlreadyStarted
	}

	am.auctions[userOpHash] = NewAuction(userOp, userOpHash)
	return userOpHash, nil
}

func (am *Manager) GetSolverOperations(userOpHash common.Hash, completionChan chan []*operation.SolverOperation) ([]*operation.SolverOperation, *relayerror.Error) {
	am.mu.RLock()
	defer am.mu.RUnlock()

	auction, ok := am.auctions[userOpHash]
	if !ok {
		log.Info("auction not found", "userOpHash", userOpHash.String())
		return nil, ErrAuctionNotFound
	}

	return auction.getSolverOps(completionChan)
}

func (am *Manager) getAuction(userOpHash common.Hash) *Auction {
	am.mu.RLock()
	defer am.mu.RUnlock()

	auction, ok := am.auctions[userOpHash]
	if !ok {
		log.Info("auction not found", "userOpHash", userOpHash.String())
		return nil
	}

	return auction
}

func (am *Manager) NewSolverOperation(solverOp *operation.SolverOperation) *relayerror.Error {
	auction := am.getAuction(solverOp.UserOpHash)
	if auction == nil {
		return ErrAuctionNotFound
	}

	dAppOp := operation.GenerateSimulationDAppOperation(auction.userOp)

	pData, err := contract.SimulatorAbi.Pack("simSolverCall", *auction.userOp, []operation.SolverOperation{*solverOp}, *dAppOp)
	if err != nil {
		log.Info("failed to pack solver operation", "err", err, "userOpHash", auction.userOpHash.String())
		return relayerror.ErrServerInternal
	}

	// TODO set proper "to" address to simulator contract
	bData, err := am.ethClient.CallContract(context.Background(), ethereum.CallMsg{To: &common.MaxAddress, Data: pData}, nil)
	if err != nil {
		log.Info("failed to call simulator contract", "err", err, "userOpHash", auction.userOpHash.String())
		return relayerror.ErrServerInternal
	}

	validOp, err := contract.SimulatorAbi.Unpack("simSolverCall", bData)
	if err != nil {
		log.Info("failed to unpack simSolverCall return data", "err", err, "userOpHash", auction.userOpHash.String())
		return relayerror.ErrServerInternal
	}

	if !validOp[0].(bool) {
		log.Info("solver operation failed simulation", "userOpHash", auction.userOpHash.String())
		return ErrSolverOpFailedSimulation
	}

	return auction.addSolverOp(solverOp)
}
