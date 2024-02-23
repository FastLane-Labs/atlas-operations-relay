package bundle

import (
	"sync"

	"github.com/FastLane-Labs/atlas-operations-relay/log"
	"github.com/FastLane-Labs/atlas-operations-relay/operation"
	"github.com/FastLane-Labs/atlas-operations-relay/relayerror"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	ErrBundleAlreadySubmitted = relayerror.NewError(3200, "bundle for this user operation has already been submitted")
	ErrBundleNotFound         = relayerror.NewError(3201, "bundle not found")
)

type Manager struct {
	ethClient *ethclient.Client

	// Indexed by userOpHash
	bundles map[common.Hash]*Bundle

	mu sync.RWMutex
}

func NewManager(ethClient *ethclient.Client) *Manager {
	return &Manager{
		ethClient: ethClient,
	}
}

func (bm *Manager) NewBundle(bundleOps *operation.BundleOperations) (*Bundle, *relayerror.Error) {
	userOpHash, err := bundleOps.UserOperation.Hash()
	if err != nil {
		log.Info("failed to compute user operation hash", "err", err)
		return nil, relayerror.ErrComputeUserOpHash.AddError(err)
	}

	// TODO: lots of checks here

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

func (bm *Manager) GetBundleHash(userOpHash common.Hash, completionChan chan common.Hash) (common.Hash, *relayerror.Error) {
	bm.mu.RLock()
	defer bm.mu.RUnlock()

	bundle, ok := bm.bundles[userOpHash]
	if !ok {
		log.Info("bundle not found", "userOpHash", userOpHash.String())
		return common.Hash{}, ErrBundleNotFound
	}

	return bundle.getAtlasTxHash(completionChan)
}
