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
	ErrSolverOperationNotFound    = relayerror.NewError(4103, "solver operation not found")
)

type Auction struct {
	open       bool
	userOpHash common.Hash

	userOp                  *operation.UserOperation
	userOperationPartialRaw *operation.UserOperationPartialRaw
	solverOpsWithScore      []*operation.SolverOperationWithScore
	solverOpsStatus         map[common.Hash]*SolverStatus
	solversParticipating    map[common.Address]struct{}
	solverGasLimit          uint32

	solverOpsCompletionSubs []chan []*operation.SolverOperationWithScore

	// Indexed by solverOpHash
	solverStatusesCompletionSubs map[common.Hash][]chan *SolverStatus

	createdAt time.Time

	mu sync.RWMutex
}

func NewAuction(duration time.Duration, userOp *operation.UserOperation, userOperationPartialRaw *operation.UserOperationPartialRaw, userOpHash common.Hash, solverGasLimit uint32) *Auction {
	auction := &Auction{
		open:                         true,
		userOpHash:                   userOpHash,
		userOp:                       userOp,
		userOperationPartialRaw:      userOperationPartialRaw,
		solverOpsWithScore:           make([]*operation.SolverOperationWithScore, 0),
		solverOpsStatus:              make(map[common.Hash]*SolverStatus),
		solversParticipating:         make(map[common.Address]struct{}),
		solverGasLimit:               solverGasLimit,
		solverOpsCompletionSubs:      make([]chan []*operation.SolverOperationWithScore, 0),
		solverStatusesCompletionSubs: make(map[common.Hash][]chan *SolverStatus),
		createdAt:                    time.Now(),
	}

	time.AfterFunc(duration, auction.close)
	return auction
}

func (a *Auction) close() {
	a.mu.RLock()
	defer a.mu.RUnlock()

	log.Info("closing auction", "userOpHash", a.userOpHash.Hex())
	a.open = false

	for _, subChan := range a.solverOpsCompletionSubs {
		select {
		case subChan <- a.solverOpsWithScore:
		default:
			// Sub isn't listening, don't block
		}
	}

	for solverOpHash, subs := range a.solverStatusesCompletionSubs {
		status := a.solverOpsStatus[solverOpHash]
		for _, subChan := range subs {
			select {
			case subChan <- status:
			default:
				// Sub isn't listening, don't block
			}
		}
	}
}

func (a *Auction) addSolverOpsCompletionSub(subChan chan []*operation.SolverOperationWithScore) {
	a.solverOpsCompletionSubs = append(a.solverOpsCompletionSubs, subChan)
}

func (a *Auction) addSolverStatusesCompletionSub(solverOpHash common.Hash, subChan chan *SolverStatus) {
	if _, ok := a.solverStatusesCompletionSubs[solverOpHash]; !ok {
		a.solverStatusesCompletionSubs[solverOpHash] = make([]chan *SolverStatus, 0)
	}

	a.solverStatusesCompletionSubs[solverOpHash] = append(a.solverStatusesCompletionSubs[solverOpHash], subChan)
}

func (a *Auction) addSolverOp(solverOp *operation.SolverOperationWithScore) (common.Hash, *relayerror.Error) {
	a.mu.Lock()
	defer a.mu.Unlock()

	if !a.open {
		log.Info("auction for this user operation has already ended", "userOpHash", a.userOpHash.Hex())
		return common.Hash{}, ErrAuctionClosed
	}

	if _, alreadyParticipating := a.solversParticipating[solverOp.SolverOperation.From]; alreadyParticipating {
		log.Info("solver is already participating in this auction", "userOpHash", a.userOpHash.Hex(), "solver", solverOp.SolverOperation.From.Hex())
		return common.Hash{}, ErrSolverAlreadyParticipating
	}

	solverOpHash, relayErr := solverOp.SolverOperation.Hash()
	if relayErr != nil {
		return common.Hash{}, relayErr
	}

	a.solverOpsWithScore = append(a.solverOpsWithScore, solverOp)
	a.solverOpsStatus[solverOpHash] = SolverStatusAuctionPending
	a.solversParticipating[solverOp.SolverOperation.From] = struct{}{}

	return solverOpHash, nil
}

func (a *Auction) getSolverOps(completionChan chan []*operation.SolverOperationWithScore) ([]*operation.SolverOperationWithScore, *relayerror.Error) {
	a.mu.Lock()
	defer a.mu.Unlock()

	if completionChan != nil && a.open {
		a.addSolverOpsCompletionSub(completionChan)
		return nil, nil
	}

	if a.open {
		log.Info("auction for this user operation is ongoing", "userOpHash", a.userOpHash.Hex())
		return nil, ErrAuctionOngoing
	}

	return a.solverOpsWithScore, nil
}

func (a *Auction) getSolverOpStatus(solverOpHash common.Hash, completionChan chan *SolverStatus) (*SolverStatus, *relayerror.Error) {
	a.mu.Lock()
	defer a.mu.Unlock()

	status, ok := a.solverOpsStatus[solverOpHash]
	if !ok {
		log.Info("solver operation not found in this auction", "solverOpHash", solverOpHash.Hex(), "userOpHash", a.userOpHash.Hex())
		return nil, ErrSolverOperationNotFound
	}

	if completionChan != nil && a.open {
		a.addSolverStatusesCompletionSub(solverOpHash, completionChan)
		return nil, nil
	}

	return status, nil
}
