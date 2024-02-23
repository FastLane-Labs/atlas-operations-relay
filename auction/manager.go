package auction

import (
	"context"
	"sync"

	"github.com/FastLane-Labs/atlas-operations-relay/log"
	"github.com/FastLane-Labs/atlas-operations-relay/operation"
	"github.com/FastLane-Labs/atlas-operations-relay/relayerror"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	ErrServerInternal        = relayerror.NewError(2201, "server internal error")
	ErrDeadlineExceeded      = relayerror.NewError(2202, "deadline exceeded")
	ErrAuctionAlreadyStarted = relayerror.NewError(2203, "auction for this user operation has already started")
	ErrAuctionNotFound       = relayerror.NewError(2204, "auction not found")
)

type Manager struct {
	ethClient *ethclient.Client

	// Indexed by userOpHash
	auctions map[common.Hash]*Auction

	mu sync.RWMutex
}

func NewManager(ethClient *ethclient.Client) *Manager {
	return &Manager{
		ethClient: ethClient,
	}
}

func (am *Manager) NewUserOperation(userOp *operation.UserOperation) (common.Hash, *relayerror.Error) {
	userOpHash, err := userOp.Hash()
	if err != nil {
		log.Info("failed to compute user operation hash", "err", err)
		return common.Hash{}, relayerror.ErrComputeUserOpHash.AddError(err)
	}

	currentBlockNumber, err := am.ethClient.BlockNumber(context.Background())
	if err != nil {
		log.Info("failed to get current block number", "err", err, "userOpHash", userOpHash.String())
		return common.Hash{}, ErrServerInternal
	}

	if userOp.Deadline.Uint64() < currentBlockNumber {
		log.Info("user operation deadline exceeded", "userOpHash", userOpHash.String())
		return common.Hash{}, ErrDeadlineExceeded
	}

	am.mu.Lock()
	defer am.mu.Unlock()

	if _, alreadyStarted := am.auctions[userOpHash]; alreadyStarted {
		log.Info("auction for this user operation has already started", "userOpHash", userOpHash.String())
		return common.Hash{}, ErrAuctionAlreadyStarted
	}

	am.auctions[userOpHash] = NewAuction(userOp)
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

func (am *Manager) NewSolverOperation(solverOp *operation.SolverOperation) *relayerror.Error {
	am.mu.RLock()
	defer am.mu.RUnlock()

	auction, ok := am.auctions[solverOp.UserOpHash]
	if !ok {
		log.Info("auction not found", "userOpHash", solverOp.UserOpHash.String())
		return ErrAuctionNotFound
	}

	return auction.addSolverOp(solverOp)
}
