package bundle

import (
	"sync"

	"github.com/FastLane-Labs/atlas-operations-relay/log"
	"github.com/FastLane-Labs/atlas-operations-relay/operation"
	"github.com/FastLane-Labs/atlas-operations-relay/relayerror"
	"github.com/ethereum/go-ethereum/common"
)

var (
	ErrNotBundledYet  = relayerror.NewError(3100, "bundle has not been bundled yet")
	ErrAlreadyBundled = relayerror.NewError(3101, "bundle has already been bundled")
	ErrAlreadyErrored = relayerror.NewError(3102, "bundle has already errored")
)

type Bundle struct {
	userOpHash common.Hash
	ops        *operation.BundleOperations

	atlasTxHash common.Hash
	relayErr    *relayerror.Error

	completionSubs []chan *Bundle

	mu sync.RWMutex
}

func NewBundle(userOpHash common.Hash, bundleOps *operation.BundleOperations) *Bundle {
	return &Bundle{
		userOpHash:     userOpHash,
		ops:            bundleOps,
		completionSubs: make([]chan *Bundle, 0),
	}
}

func (b *Bundle) GetResult() (common.Hash, *relayerror.Error) {
	b.mu.RLock()
	defer b.mu.RUnlock()
	return b.atlasTxHash, b.relayErr
}

func (b *Bundle) SetAtlasTxHash(txHash common.Hash) *relayerror.Error {
	b.mu.Lock()
	defer b.mu.Unlock()

	if b.atlasTxHash != (common.Hash{}) {
		log.Info("bundle has already been bundled", "bundledTxHash", b.atlasTxHash.Hex(), "newTxHash", txHash.Hex())
		return ErrAlreadyBundled
	}

	b.atlasTxHash = txHash
	b.notifyCompletionSubs()

	return nil
}

func (b *Bundle) SetRelayError(relayErr *relayerror.Error) *relayerror.Error {
	b.mu.Lock()
	defer b.mu.Unlock()

	if b.relayErr != nil {
		log.Info("bundle has already errored", "relayErr", b.relayErr, "newRelayErr", relayErr)
		return ErrAlreadyErrored
	}

	b.relayErr = relayErr
	b.notifyCompletionSubs()

	return nil
}

func (b *Bundle) notifyCompletionSubs() {
	for _, subChan := range b.completionSubs {
		select {
		case subChan <- b:
		default:
			// Sub isn't listening, don't block
		}
	}
}

func (b *Bundle) addCompletionSub(subChan chan *Bundle) {
	b.mu.Lock()
	defer b.mu.Unlock()

	b.completionSubs = append(b.completionSubs, subChan)
}

func (b *Bundle) getAtlasTxHash(completionChan chan *Bundle) (common.Hash, *relayerror.Error) {
	atlasTxHash, relayErr := b.GetResult()
	if relayErr != nil {
		return common.Hash{}, relayErr
	}

	bundled := atlasTxHash != (common.Hash{})
	if completionChan != nil && !bundled {
		b.addCompletionSub(completionChan)
		return common.Hash{}, nil
	}

	if !bundled {
		log.Info("bundle has not been bundled yet", "userOpHash", b.userOpHash.Hex())
		return common.Hash{}, ErrNotBundledYet
	}

	return atlasTxHash, nil
}
