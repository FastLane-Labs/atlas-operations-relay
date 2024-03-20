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
	ErrAuctionClosed  = relayerror.NewError(4100, "auction for this user operation has already ended")
	ErrAuctionOngoing = relayerror.NewError(4101, "auction for this user operation is ongoing")
)

type Auction struct {
	open       bool
	userOpHash common.Hash

	userOp      *operation.UserOperation
	solverInput *operation.SolverInput
	solverOps   []*operation.SolverOperation

	completionSubs []chan []*operation.SolverOperation

	createdAt time.Time

	mu sync.RWMutex
}

func NewAuction(duration time.Duration, userOp *operation.UserOperation, solverInput *operation.SolverInput, userOpHash common.Hash) *Auction {
	auction := &Auction{
		open:           true,
		userOpHash:     userOpHash,
		userOp:         userOp,
		solverInput:    solverInput,
		solverOps:      make([]*operation.SolverOperation, 0),
		completionSubs: make([]chan []*operation.SolverOperation, 0),
		createdAt:      time.Now(),
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
		case subChan <- a.solverOps:
		default:
			// Sub isn't listening, don't block
		}
	}
}

func (a *Auction) addCompletionSub(subChan chan []*operation.SolverOperation) {
	a.mu.Lock()
	defer a.mu.Unlock()

	a.completionSubs = append(a.completionSubs, subChan)
}

func (a *Auction) addSolverOp(solverOp *operation.SolverOperation) *relayerror.Error {
	a.mu.Lock()
	defer a.mu.Unlock()

	if !a.open {
		log.Info("auction for this user operation has already ended", "userOpHash", solverOp.UserOpHash.Hex())
		return ErrAuctionClosed
	}

	a.solverOps = append(a.solverOps, solverOp)
	return nil
}

func (a *Auction) getSolverOps(completionChan chan []*operation.SolverOperation) ([]*operation.SolverOperation, *relayerror.Error) {
	a.mu.RLock()
	open := a.open
	a.mu.RUnlock()

	if completionChan != nil && open {
		a.addCompletionSub(completionChan)
		return nil, nil
	}

	if open {
		userOpHash, _ := a.userOp.Hash()
		log.Info("auction for this user operation is ongoing", "userOpHash", userOpHash.Hex())
		return nil, ErrAuctionOngoing
	}

	return a.solverOps, nil
}
