package auction

import (
	"sync"
	"time"

	"github.com/FastLane-Labs/atlas-operations-relay/log"
	"github.com/FastLane-Labs/atlas-operations-relay/operation"
	"github.com/FastLane-Labs/atlas-operations-relay/relayerror"
)

const (
	AuctionDuration = 500 * time.Millisecond
)

var (
	ErrAuctionClosed  = relayerror.NewError(2100, "auction for this user operation has already ended")
	ErrAuctionOngoing = relayerror.NewError(2101, "auction for this user operation is ongoing")
)

type Auction struct {
	open bool

	userOp    *operation.UserOperation
	solverOps []*operation.SolverOperation

	completionSubs []chan []*operation.SolverOperation

	mu sync.RWMutex
}

func NewAuction(userOp *operation.UserOperation) *Auction {
	auction := &Auction{
		open:           true,
		userOp:         userOp,
		solverOps:      make([]*operation.SolverOperation, 0),
		completionSubs: make([]chan []*operation.SolverOperation, 0),
	}

	time.AfterFunc(AuctionDuration, auction.close)
	return auction
}

func (a *Auction) close() {
	a.mu.Lock()
	defer a.mu.Unlock()

	a.open = false

	for _, subChan := range a.completionSubs {
		subChan <- a.solverOps
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
		log.Info("auction for this user operation has already ended", "userOpHash", solverOp.UserOpHash.String())
		return ErrAuctionClosed
	}

	// TODO: add checks here

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
		log.Info("auction for this user operation is ongoing", "userOpHash", userOpHash.String())
		return nil, ErrAuctionOngoing
	}

	return a.solverOps, nil
}
