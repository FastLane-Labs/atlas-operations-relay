package auction

import (
	"context"
	"encoding/hex"
	"fmt"
	"math/big"
	"sync"
	"time"

	"github.com/FastLane-Labs/atlas-operations-relay/config"
	"github.com/FastLane-Labs/atlas-operations-relay/contract"
	"github.com/FastLane-Labs/atlas-operations-relay/contract/dAppControl"
	"github.com/FastLane-Labs/atlas-operations-relay/log"
	"github.com/FastLane-Labs/atlas-operations-relay/operation"
	"github.com/FastLane-Labs/atlas-operations-relay/relayerror"
	"github.com/FastLane-Labs/atlas-operations-relay/utils"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	ErrAuctionAlreadyStarted    = relayerror.NewError(4000, "auction for this user operation has already started")
	ErrAuctionNotFound          = relayerror.NewError(4001, "auction not found")
	ErrUserOpFailedSimulation   = relayerror.NewError(4002, "user operation failed simulation")
	ErrSolverOpFailedSimulation = relayerror.NewError(4003, "solver operation failed simulation")
	ErrNotEnoughBondedBalance   = relayerror.NewError(4004, "not enough atlEth bonded balance")
)

var (
	ATLETH_BONDED_BALANCE_MULTI_REQ = big.NewInt(1)
)

type solverGasLimitFn func(common.Address) (uint32, *relayerror.Error)
type balanceOfBondedFn func(common.Address) (*big.Int, *relayerror.Error)
type reputationScoreFn func(common.Address) int
type getDAppConfigFn func(common.Address, *operation.UserOperation) (*dAppControl.DAppConfig, *relayerror.Error)
type auctionCompleteCallbackFn func(*operation.BundleOperations)

type Manager struct {
	ethClient *ethclient.Client
	config    *config.Config

	// Indexed by userOpHash
	auctions map[common.Hash]*Auction

	solverOpHashToAuction map[common.Hash]*Auction

	solverGasLimit          solverGasLimitFn
	balanceOfBonded         balanceOfBondedFn
	reputationScore         reputationScoreFn
	getDAppConfig           getDAppConfigFn
	auctionCompleteCallback auctionCompleteCallbackFn

	mu sync.RWMutex
}

func NewManager(ethClient *ethclient.Client, config *config.Config, solverGasLimit solverGasLimitFn, balanceOfBonded balanceOfBondedFn, reputationScore reputationScoreFn, getDAppConfig getDAppConfigFn, auctionCompleteCallback auctionCompleteCallbackFn) *Manager {
	am := &Manager{
		ethClient:               ethClient,
		config:                  config,
		auctions:                make(map[common.Hash]*Auction),
		solverOpHashToAuction:   make(map[common.Hash]*Auction),
		solverGasLimit:          solverGasLimit,
		balanceOfBonded:         balanceOfBonded,
		reputationScore:         reputationScore,
		getDAppConfig:           getDAppConfig,
		auctionCompleteCallback: auctionCompleteCallback,
	}

	go am.auctionsCleaner()
	return am
}

func (am *Manager) auctionsCleaner() {
	for range time.Tick(10 * time.Minute) {
		am.mu.Lock()
		for userOpHash, auction := range am.auctions {
			if !auction.isOpen() && time.Since(auction.createdAt) > time.Hour {
				delete(am.auctions, userOpHash)
			}
		}
		am.mu.Unlock()
	}
}

func (am *Manager) NewUserOperation(userOp *operation.UserOperation, hints []common.Address) (common.Hash, *operation.UserOperationPartialRaw, *relayerror.Error) {
	dAppConfig, relayErr := am.getDAppConfig(userOp.Control, userOp)
	if relayErr != nil {
		log.Info("failed to get dapp config", "err", relayErr.Message)
		return common.Hash{}, nil, relayErr
	}

	userOpHash, relayErr := userOp.Hash(utils.FlagTrustedOpHash(dAppConfig.CallConfig), &am.config.Relay.Eip712.Domain)
	if relayErr != nil {
		log.Info("failed to compute user operation hash", "err", relayErr.Message)
		return common.Hash{}, nil, relayErr
	}

	if relayErr := userOp.Validate(am.ethClient, am.config.Contracts.Atlas, &am.config.Relay.Eip712.Domain, am.config.Relay.Gas.MaxPerUserOperation); relayErr != nil {
		log.Info("invalid user operation", "err", relayErr.Message, "userOpHash", userOpHash.Hex())
		return common.Hash{}, nil, relayErr
	}

	if relayErr := am.simulateUserOperation(userOp, userOpHash); relayErr != nil {
		log.Info("user operation failed simulation", "err", relayErr.Message, "userOpHash", userOpHash.Hex())
		return common.Hash{}, nil, relayErr
	}

	solverGasLimit, relayErr := am.solverGasLimit(userOp.Control)
	if relayErr != nil {
		log.Info("failed to get solver gas limit", "err", relayErr.Message, "userOpHash", userOpHash.Hex())
		return common.Hash{}, nil, relayErr
	}

	userOperationPartialRaw := operation.NewUserOperationPartialRaw(userOpHash, userOp, hints)

	am.mu.Lock()
	defer am.mu.Unlock()

	if _, alreadyStarted := am.auctions[userOpHash]; alreadyStarted {
		log.Info("auction for this user operation has already started", "userOpHash", userOpHash.Hex())
		return common.Hash{}, nil, ErrAuctionAlreadyStarted
	}

	am.auctions[userOpHash] = NewAuction(am.config.Relay.Auction.Duration, am.config.Relay.Auction.MaxSolutions, userOp, userOperationPartialRaw, userOpHash, solverGasLimit, am.simulateSolverOperation, am.auctionCompleteCallback)
	return userOpHash, userOperationPartialRaw, nil
}

func (am *Manager) GetSolverOperations(userOpHash common.Hash, completionChan chan []*operation.SolverOperationWithScore) ([]*operation.SolverOperationWithScore, *relayerror.Error) {
	am.mu.RLock()
	defer am.mu.RUnlock()

	auction, ok := am.auctions[userOpHash]
	if !ok {
		log.Info("auction not found", "userOpHash", userOpHash.Hex())
		return nil, ErrAuctionNotFound
	}

	return auction.getSolverOps(completionChan)
}

func (am *Manager) IsAuctionOngoing(userOpHash common.Hash) bool {
	am.mu.RLock()
	defer am.mu.RUnlock()
	auction, ok := am.auctions[userOpHash]
	return ok && auction.isOpen()
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

func (am *Manager) NewSolverOperation(solverOp *operation.SolverOperation) (common.Hash, *relayerror.Error) {
	auction := am.getAuction(solverOp.UserOpHash)
	if auction == nil {
		return common.Hash{}, ErrAuctionNotFound
	}

	relayErr := solverOp.Validate(auction.userOp, am.config.Contracts.Atlas, &am.config.Relay.Eip712.Domain, auction.solverGasLimit)
	if relayErr != nil {
		log.Info("invalid solver operation", "err", relayErr.Message, "userOpHash", auction.userOpHash.Hex())
		return common.Hash{}, relayErr
	}

	bondedBalance, relayErr := am.balanceOfBonded(solverOp.From)
	if relayErr != nil {
		return common.Hash{}, relayErr
	}

	// Bonded balance must be >= ATLETH_BONDED_BALANCE_MULTI_REQ * (gas * maxFeePerGas)
	atlEthRequired := new(big.Int).Mul(ATLETH_BONDED_BALANCE_MULTI_REQ, new(big.Int).Mul(solverOp.Gas, solverOp.MaxFeePerGas))
	if bondedBalance.Cmp(atlEthRequired) < 0 {
		log.Info("not enough bonded balance", "userOpHash", solverOp.UserOpHash.Hex(), "from", solverOp.From.Hex(), "bondedBalance", bondedBalance.String(), "atlEthRequired", atlEthRequired.String())
		return common.Hash{}, ErrNotEnoughBondedBalance
	}

	solverOpWithScore := &operation.SolverOperationWithScore{
		SolverOperation: solverOp,
		Score:           am.reputationScore(solverOp.From),
	}

	solverOpHash, relayErr := auction.addSolverOp(solverOpWithScore, &am.config.Relay.Eip712.Domain)
	if relayErr != nil {
		return common.Hash{}, relayErr
	}

	am.solverOpHashToAuction[solverOpHash] = auction

	log.Info("valid solver operation", "userOpHash", auction.userOpHash.Hex(), "solverOpHash", solverOpHash.Hex(), "from", solverOp.From.Hex(), "score", solverOpWithScore.Score)

	return solverOpHash, relayErr
}

func (am *Manager) GetSolverOperationStatus(solverOpHash common.Hash, completionChan chan *SolverStatus) (*SolverStatus, *relayerror.Error) {
	am.mu.RLock()
	defer am.mu.RUnlock()

	auction, ok := am.solverOpHashToAuction[solverOpHash]
	if !ok {
		log.Info("auction not found", "solverOpHash", solverOpHash.Hex())
		return nil, ErrAuctionNotFound
	}

	return auction.getSolverOpStatus(solverOpHash, completionChan)
}

func (am *Manager) simulateUserOperation(userOp *operation.UserOperation, userOpHash common.Hash) *relayerror.Error {
	if !am.config.Relay.Simulations {
		// Simulations disabled
		return nil
	}

	pData, err := contract.SimulatorAbi.Pack("simUserOperation", *userOp)
	if err != nil {
		log.Info("failed to pack user operation", "err", err, "userOpHash", userOpHash.Hex())
		return relayerror.ErrServerInternal
	}

	bData, err := am.ethClient.CallContract(
		context.Background(),
		ethereum.CallMsg{
			To:        &am.config.Contracts.Simulator,
			Gas:       userOp.Gas.Uint64() + 1500000, // Add gas for validateCalls and others
			GasFeeCap: new(big.Int).Set(userOp.MaxFeePerGas),
			Value:     new(big.Int).Set(userOp.Value),
			Data:      pData,
		},
		nil)
	if err != nil {
		log.Info("failed to call simulator contract", "err", err, "userOpHash", userOpHash.Hex())
		return relayerror.ErrServerInternal
	}

	validOp, err := contract.SimulatorAbi.Unpack("simUserOperation", bData)
	if err != nil {
		log.Info("failed to unpack simUserOperation return data", "err", err, "userOpHash", userOpHash.Hex())
		return relayerror.ErrServerInternal
	}

	if !validOp[0].(bool) {
		result := validOp[1].(uint8)
		validCallResult := validOp[2].(*big.Int)
		log.Info("user operation failed simulation", "userOpHash", userOpHash.Hex(), "result", result, "validCallResult", validCallResult.String())
		log.Debug("user operation pData", "pData", hex.EncodeToString(pData))
		return ErrUserOpFailedSimulation.AddMessage(fmt.Sprintf("result: %d, validCallResult: %s", result, validCallResult.String()))
	}

	return nil
}

func (am *Manager) simulateSolverOperation(userOp *operation.UserOperation, userOpHash common.Hash, solverOp *operation.SolverOperation) *relayerror.Error {
	if !am.config.Relay.Simulations {
		// Simulations disabled
		return nil
	}

	dAppOp, err := operation.GenerateSimulationDAppOperation(userOpHash, userOp, []*operation.SolverOperation{solverOp})
	if err != nil {
		log.Info("failed to generate simulation dapp operation", "err", err, "userOpHash", userOpHash.Hex())
		return relayerror.ErrServerInternal
	}

	pData, err := contract.SimulatorAbi.Pack("simSolverCall", *userOp, *solverOp, *dAppOp)
	if err != nil {
		log.Info("failed to pack solver operation", "err", err, "userOpHash", userOpHash.Hex())
		return relayerror.ErrServerInternal
	}

	gasPrice := new(big.Int).Set(userOp.MaxFeePerGas)
	if solverOp.MaxFeePerGas.Cmp(userOp.MaxFeePerGas) > 0 {
		gasPrice.Set(solverOp.MaxFeePerGas)
	}

	bData, err := am.ethClient.CallContract(
		context.Background(),
		ethereum.CallMsg{
			To:        &am.config.Contracts.Simulator,
			Gas:       userOp.Gas.Uint64() + solverOp.Gas.Uint64() + 1500000, // Add gas for validateCalls and others
			GasFeeCap: gasPrice,
			Value:     new(big.Int).Set(userOp.Value),
			Data:      pData,
		},
		nil,
	)
	if err != nil {
		log.Info("failed to call simulator contract", "err", err, "userOpHash", userOpHash.Hex())
		return relayerror.ErrServerInternal
	}

	validOp, err := contract.SimulatorAbi.Unpack("simSolverCall", bData)
	if err != nil {
		log.Info("failed to unpack simSolverCall return data", "err", err, "userOpHash", userOpHash.Hex())
		return relayerror.ErrServerInternal
	}

	if !validOp[0].(bool) {
		result := validOp[1].(uint8)
		solverOutcomeResult := validOp[2].(*big.Int)
		log.Info("solver operation failed simulation", "userOpHash", userOpHash.Hex(), "result", result, "solverOutcomeResult", solverOutcomeResult.String())
		log.Debug("solver operation pData", "pData", hex.EncodeToString(pData))
		return ErrSolverOpFailedSimulation.AddMessage(fmt.Sprintf("result: %d, solverOutcomeResult: %s", result, solverOutcomeResult.String()))
	}

	return nil
}
