package auction

import (
	"sync"
	"time"

	"github.com/FastLane-Labs/atlas-operations-relay/log"
	"github.com/FastLane-Labs/atlas-operations-relay/operation"
	"github.com/FastLane-Labs/atlas-operations-relay/relayerror"
	"github.com/ethereum/go-ethereum/common"
)

var (
	ErrAuctionClosed              = relayerror.NewError(4100, "auction for this user operation has already ended")
	ErrAuctionOngoing             = relayerror.NewError(4101, "auction for this user operation is ongoing")
	ErrSolverAlreadyParticipating = relayerror.NewError(4102, "solver is already participating in this auction")
)

type Auction struct {
	open       bool
	userOpHash common.Hash

	userOp                  *operation.UserOperation
	userOperationPartialRaw *operation.UserOperationPartialRaw
	solverOpsWithScore      []*operation.SolverOperationWithScore
	solversParticipating    map[common.Address]struct{}
	solverGasLimit          uint32

	completionSubs []chan []*operation.SolverOperationWithScore

	createdAt time.Time

	mu sync.RWMutex
}

func NewAuction(duration time.Duration, userOp *operation.UserOperation, userOperationPartialRaw *operation.UserOperationPartialRaw, userOpHash common.Hash, solverGasLimit uint32) *Auction {
	auction := &Auction{
		open:                    true,
		userOpHash:              userOpHash,
		userOp:                  userOp,
		userOperationPartialRaw: userOperationPartialRaw,
		solverOpsWithScore:      make([]*operation.SolverOperationWithScore, 0),
		solversParticipating:    make(map[common.Address]struct{}),
		solverGasLimit:          solverGasLimit,
		completionSubs:          make([]chan []*operation.SolverOperationWithScore, 0),
		createdAt:               time.Now(),
	}

	time.AfterFunc(duration, auction.close)
	return auction
}

func (a *Auction) close() {
	a.mu.Lock()
	defer a.mu.Unlock()

	log.Info("closing auction", "userOpHash", a.userOpHash.Hex())
	a.open = false

	for _, subChan := range a.completionSubs {
		select {
		case subChan <- a.solverOpsWithScore:
		default:
			// Sub isn't listening, don't block
		}
	}
}

func (a *Auction) addCompletionSub(subChan chan []*operation.SolverOperationWithScore) {
	a.mu.Lock()
	defer a.mu.Unlock()

	a.completionSubs = append(a.completionSubs, subChan)
}

func (a *Auction) addSolverOp(solverOp *operation.SolverOperationWithScore) *relayerror.Error {
	a.mu.Lock()
	defer a.mu.Unlock()

	if !a.open {
		log.Info("auction for this user operation has already ended", "userOpHash", a.userOpHash.Hex())
		return ErrAuctionClosed
	}

	if _, alreadyParticipating := a.solversParticipating[solverOp.SolverOperation.From]; alreadyParticipating {
		log.Info("solver is already participating in this auction", "userOpHash", a.userOpHash.Hex(), "solver", solverOp.SolverOperation.From.Hex())
		return ErrSolverAlreadyParticipating
	}

	a.solverOpsWithScore = append(a.solverOpsWithScore, solverOp)
	a.solversParticipating[solverOp.SolverOperation.From] = struct{}{}
	return nil
}

func (a *Auction) getSolverOps(completionChan chan []*operation.SolverOperationWithScore) ([]*operation.SolverOperationWithScore, *relayerror.Error) {
	a.mu.RLock()
	open := a.open
	a.mu.RUnlock()

	if completionChan != nil && open {
		a.addCompletionSub(completionChan)
		return nil, nil
	}

	if open {
		log.Info("auction for this user operation is ongoing", "userOpHash", a.userOpHash.Hex())
		return nil, ErrAuctionOngoing
	}

	return a.solverOpsWithScore, nil
}
