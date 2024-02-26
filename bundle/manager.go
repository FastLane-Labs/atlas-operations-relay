package bundle

import (
	"context"
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
	ErrBundleAlreadySubmitted = relayerror.NewError(3200, "bundle for this user operation has already been submitted")
	ErrBundleNotFound         = relayerror.NewError(3201, "bundle not found")
	ErrBundleFailedSimulation = relayerror.NewError(3202, "bundle failed simulation")
)

type Manager struct {
	ethClient *ethclient.Client
	config    *config.Config

	// Indexed by userOpHash
	bundles map[common.Hash]*Bundle

	mu sync.RWMutex
}

func NewManager(ethClient *ethclient.Client, config *config.Config) *Manager {
	return &Manager{
		ethClient: ethClient,
		config:    config,
	}
}

func (bm *Manager) NewBundle(bundleOps *operation.BundleOperations) (*Bundle, *relayerror.Error) {
	userOpHash, err := bundleOps.UserOperation.Hash()
	if err != nil {
		log.Info("failed to compute user operation hash", "err", err)
		return nil, relayerror.ErrComputeUserOpHash.AddError(err)
	}

	solverOps := make([]operation.SolverOperation, 0, len(bundleOps.SolverOperations))
	for _, solverOp := range bundleOps.SolverOperations {
		solverOps = append(solverOps, *solverOp)
	}

	pData, err := contract.SimulatorAbi.Pack("simSolverCalls", *bundleOps.UserOperation, solverOps, *bundleOps.DAppOperation)
	if err != nil {
		log.Info("failed to pack bundle", "err", err, "userOpHash", userOpHash.Hex())
		return nil, relayerror.ErrServerInternal
	}

	bData, err := bm.ethClient.CallContract(context.Background(), ethereum.CallMsg{To: &bm.config.Contracts.Simulator, Data: pData}, nil)
	if err != nil {
		log.Info("failed to call simulator contract", "err", err, "userOpHash", userOpHash.Hex())
		return nil, relayerror.ErrServerInternal
	}

	validOp, err := contract.SimulatorAbi.Unpack("simSolverCalls", bData)
	if err != nil {
		log.Info("failed to unpack simSolverCalls return data", "err", err, "userOpHash", userOpHash.Hex())
		return nil, relayerror.ErrServerInternal
	}

	if !validOp[0].(bool) {
		log.Info("bundle failed simulation", "userOpHash", userOpHash.Hex())
		return nil, ErrBundleFailedSimulation
	}

	bm.mu.Lock()
	defer bm.mu.Unlock()

	if _, exist := bm.bundles[userOpHash]; exist {
		return nil, ErrBundleAlreadySubmitted
	}

	bundle := NewBundle(userOpHash, bundleOps)
	bm.bundles[userOpHash] = bundle
	return bundle, nil
}

func (bm *Manager) UnregisterBundle(userOpHash common.Hash) {
	bm.mu.Lock()
	defer bm.mu.Unlock()

	delete(bm.bundles, userOpHash)
}

func (bm *Manager) GetBundleHash(userOpHash common.Hash, completionChan chan *Bundle) (common.Hash, *relayerror.Error) {
	bm.mu.RLock()
	defer bm.mu.RUnlock()

	bundle, ok := bm.bundles[userOpHash]
	if !ok {
		log.Info("bundle not found", "userOpHash", userOpHash.Hex())
		return common.Hash{}, ErrBundleNotFound
	}

	return bundle.getAtlasTxHash(completionChan)
}
