package core

import (
	"github.com/FastLane-Labs/atlas-operations-relay/auction"
	"github.com/FastLane-Labs/atlas-operations-relay/bundle"
	"github.com/FastLane-Labs/atlas-operations-relay/operation"
	"github.com/FastLane-Labs/atlas-operations-relay/relayerror"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

const (
	SolverOpSuccessfullySubmitted = "solver operation successfully submitted"
	BundleSuccessfullySubmitted   = "bundle successfully submitted"
)

var (
	ErrForwardBundle = relayerror.NewError(1200, "failed to forward bundle")
)

type Relay struct {
	server *Server

	auctionManager *auction.Manager
	bundleManager  *bundle.Manager
}

func NewRelay(ethClient *ethclient.Client) *Relay {
	auctionManager := auction.NewManager(ethClient)

	relay := &Relay{
		auctionManager: auctionManager,
		bundleManager:  bundle.NewManager(ethClient),
	}
	relay.server = NewServer(NewRouter(NewApi(relay)), auctionManager.NewSolverOperation)
	return relay
}

func (r *Relay) Run() {
	r.server.ListenAndServe()
}

func (r *Relay) submitUserOperation(userOp *operation.UserOperation) (common.Hash, *relayerror.Error) {
	userOpHash, relayErr := r.auctionManager.NewUserOperation(userOp)
	if relayErr != nil {
		return common.Hash{}, relayErr
	}

	go r.server.BroadcastUserOperation(userOp)
	return userOpHash, nil
}

func (r *Relay) getSolverOperations(userOpHash common.Hash, completionChan chan []*operation.SolverOperation) ([]*operation.SolverOperation, *relayerror.Error) {
	return r.auctionManager.GetSolverOperations(userOpHash, completionChan)
}

func (r *Relay) submitBundleOperations(bundleOps *operation.BundleOperations) (string, *relayerror.Error) {
	bundle, relayErr := r.bundleManager.NewBundle(bundleOps)
	if relayErr != nil {
		return "", relayErr
	}

	if err := r.server.ForwardBundle(bundleOps, bundle.SetAtlasTxHash); err != nil {
		userOphash, _ := bundleOps.UserOperation.Hash()
		r.bundleManager.UnregisterBundle(userOphash)
		return "", ErrForwardBundle.AddError(err)
	}

	return BundleSuccessfullySubmitted, nil
}

func (r *Relay) getBundleHash(userOpHash common.Hash, completionChan chan common.Hash) (common.Hash, *relayerror.Error) {
	return r.bundleManager.GetBundleHash(userOpHash, completionChan)
}

func (r *Relay) submitSolverOperation(solverOp *operation.SolverOperation) (string, *relayerror.Error) {
	var result string

	relayErr := r.auctionManager.NewSolverOperation(solverOp)
	if relayErr == nil {
		result = SolverOpSuccessfullySubmitted
	}

	return result, relayErr
}
