package auction

import (
	"sort"
	"sync"
	"time"

	"github.com/FastLane-Labs/atlas-operations-relay/log"
	"github.com/FastLane-Labs/atlas-operations-relay/operation"
	"github.com/FastLane-Labs/atlas-operations-relay/relayerror"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/signer/core/apitypes"
)

var (
	ErrAuctionClosed              = relayerror.NewError(4100, "auction for this user operation has already ended")
	ErrAuctionOngoing             = relayerror.NewError(4101, "auction for this user operation is ongoing")
	ErrSolverAlreadyParticipating = relayerror.NewError(4102, "solver is already participating in this auction")
	ErrSolverOperationNotFound    = relayerror.NewError(4103, "solver operation not found")
)

type simulateSolverOperationFn func(*operation.UserOperation, common.Hash, *operation.SolverOperation) *relayerror.Error

type Auction struct {
	open         bool
	userOpHash   common.Hash
	maxSolutions uint64

	userOp                  *operation.UserOperation
	userOperationPartialRaw *operation.UserOperationPartialRaw
	solverOpsWithScore      []*operation.SolverOperationWithScore
	solverOpsStatus         map[common.Hash]*SolverStatus
	solversParticipating    map[common.Address]struct{}
	solverGasLimit          uint32

	solverOpsCompletionSubs []chan []*operation.SolverOperationWithScore

	// Indexed by solverOpHash
	solverStatusesCompletionSubs map[common.Hash][]chan *SolverStatus

	simulateSolverOperation simulateSolverOperationFn
	auctionCompleteCallback auctionCompleteCallbackFn

	createdAt time.Time

	mu sync.Mutex
}

func NewAuction(duration time.Duration, maxSolutions uint64, userOp *operation.UserOperation, userOperationPartialRaw *operation.UserOperationPartialRaw, userOpHash common.Hash, solverGasLimit uint32, simulateSolverOperation simulateSolverOperationFn, auctionCompleteCallback auctionCompleteCallbackFn) *Auction {
	auction := &Auction{
		open:                         true,
		userOpHash:                   userOpHash,
		maxSolutions:                 maxSolutions,
		userOp:                       userOp,
		userOperationPartialRaw:      userOperationPartialRaw,
		solverOpsWithScore:           make([]*operation.SolverOperationWithScore, 0),
		solverOpsStatus:              make(map[common.Hash]*SolverStatus),
		solversParticipating:         make(map[common.Address]struct{}),
		solverGasLimit:               solverGasLimit,
		solverOpsCompletionSubs:      make([]chan []*operation.SolverOperationWithScore, 0),
		solverStatusesCompletionSubs: make(map[common.Hash][]chan *SolverStatus),
		simulateSolverOperation:      simulateSolverOperation,
		auctionCompleteCallback:      auctionCompleteCallback,
		createdAt:                    time.Now(),
	}

	log.Info("starting auction", "userOpHash", userOpHash.Hex(), "duration", common.PrettyDuration(duration), "maxSolutions", maxSolutions)
	time.AfterFunc(duration, auction.close)

	return auction
}

func (a *Auction) isOpen() bool {
	a.mu.Lock()
	defer a.mu.Unlock()
	return a.open
}

func (a *Auction) close() {
	a.mu.Lock()
	defer a.mu.Unlock()

	log.Info("closing auction", "userOpHash", a.userOpHash.Hex())
	a.open = false

	// Sort solverOps by score
	sort.Slice(a.solverOpsWithScore, func(i, j int) bool {
		return a.solverOpsWithScore[i].Score > a.solverOpsWithScore[j].Score
	})

	if len(a.solverOpsWithScore) > int(a.maxSolutions) {
		// Update status of solvers that didn't make the cut
		for _, solverOp := range a.solverOpsWithScore[a.maxSolutions:] {
			a.solverOpsStatus[solverOp.SolverOpHash] = SolverStatusNotIncluded
		}

		// Truncate solverOps to maxSolutions
		a.solverOpsWithScore = a.solverOpsWithScore[:a.maxSolutions]
	}

	var (
		solverOpsWithScoreTmp = make([]*operation.SolverOperationWithScore, 0)
		mu                    sync.Mutex
		wg                    sync.WaitGroup
	)

	// Simulate competing solver operations
	for _, solverOpWithScore := range a.solverOpsWithScore {
		wg.Add(1)
		go func(solverOpWithScore *operation.SolverOperationWithScore) {
			defer wg.Done()

			relayErr := a.simulateSolverOperation(a.userOp, a.userOpHash, solverOpWithScore.SolverOperation)
			if relayErr != nil {
				// Update status
				mu.Lock()
				a.solverOpsStatus[solverOpWithScore.SolverOpHash] = SolverStatusFailedSimulation
				mu.Unlock()
				log.Info("solver operation failed simulation", "userOpHash", a.userOpHash.Hex(), "solverOpHash", solverOpWithScore.SolverOpHash.Hex(), "err", relayErr.Message)
				return
			}

			mu.Lock()
			solverOpsWithScoreTmp = append(solverOpsWithScoreTmp, solverOpWithScore)

			// Update status
			a.solverOpsStatus[solverOpWithScore.SolverOpHash] = SolverStatusIncluded
			mu.Unlock()

			log.Info("solver operation added to auction's final list", "userOpHash", a.userOpHash.Hex(), "solverOpHash", solverOpWithScore.SolverOpHash.Hex())
		}(solverOpWithScore)
	}

	wg.Wait()

	// Re-sort since simulations were parallelized
	sort.Slice(solverOpsWithScoreTmp, func(i, j int) bool {
		return solverOpsWithScoreTmp[i].Score > solverOpsWithScoreTmp[j].Score
	})

	a.solverOpsWithScore = solverOpsWithScoreTmp

	// Update waiting subscribers
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

	// Notify auction complete callback
	solverOps := make([]*operation.SolverOperation, 0, len(a.solverOpsWithScore))
	for _, solverOpWithScore := range a.solverOpsWithScore {
		solverOps = append(solverOps, solverOpWithScore.SolverOperation)
	}

	bundle := &operation.BundleOperations{
		UserOperation:    a.userOp,
		SolverOperations: solverOps,
		// DAppOperation will be generated by the relay if it can
	}

	a.auctionCompleteCallback(bundle)
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

func (a *Auction) addSolverOp(solverOp *operation.SolverOperationWithScore, eip712Domain *apitypes.TypedDataDomain) (common.Hash, *relayerror.Error) {
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

	solverOpHash, relayErr := solverOp.SolverOperation.Hash(eip712Domain)
	if relayErr != nil {
		return common.Hash{}, relayErr
	}

	solverOp.SolverOpHash = solverOpHash

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
