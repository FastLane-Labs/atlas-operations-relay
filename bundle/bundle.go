package bundle

import (
	"sync"

	"github.com/FastLane-Labs/atlas-operations-relay/operation"
	"github.com/FastLane-Labs/atlas-operations-relay/relayerror"
	"github.com/ethereum/go-ethereum/common"
)

var (
	ErrNotBundledYet  = relayerror.NewError(3100, "bundle has not been bundled yet")
	ErrAlreadyBundled = relayerror.NewError(3101, "bundle has already been bundled")
)

type Bundle struct {
	ops         *operation.BundleOperations
	atlasTxHash common.Hash

	completionSubs []chan common.Hash

	mu sync.RWMutex
}

func NewBundle(bundleOps *operation.BundleOperations) *Bundle {
	return &Bundle{
		ops:            bundleOps,
		completionSubs: make([]chan common.Hash, 0),
	}
}

func (b *Bundle) hasBeenBundled() bool {
	b.mu.RLock()
	defer b.mu.RUnlock()
	return b.atlasTxHash != (common.Hash{})
}

func (b *Bundle) SetAtlasTxHash(txHash common.Hash) *relayerror.Error {
	b.mu.Lock()
	defer b.mu.Unlock()

	if b.atlasTxHash != (common.Hash{}) {
		return ErrAlreadyBundled
	}

	b.atlasTxHash = txHash

	for _, subChan := range b.completionSubs {
		subChan <- b.atlasTxHash
	}

	return nil
}

func (b *Bundle) addCompletionSub(subChan chan common.Hash) {
	b.mu.Lock()
	defer b.mu.Unlock()

	b.completionSubs = append(b.completionSubs, subChan)
}

func (b *Bundle) getAtlasTxHash(completionChan chan common.Hash) (common.Hash, *relayerror.Error) {
	bundled := b.hasBeenBundled()

	if completionChan != nil && !bundled {
		b.addCompletionSub(completionChan)
		return common.Hash{}, nil
	}

	if !bundled {
		return common.Hash{}, ErrNotBundledYet
	}

	return b.atlasTxHash, nil
}
