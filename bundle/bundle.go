package bundle

import (
	"sync"
	"time"

	"github.com/FastLane-Labs/atlas-operations-relay/log"
	"github.com/FastLane-Labs/atlas-operations-relay/operation"
	"github.com/FastLane-Labs/atlas-operations-relay/relayerror"
	"github.com/ethereum/go-ethereum/common"
)

var (
	ErrNotBundledYet  = relayerror.NewError(5101, "bundle has not been bundled yet")
	ErrAlreadyBundled = relayerror.NewError(5102, "bundle has already been bundled")
	ErrAlreadyErrored = relayerror.NewError(5103, "bundle has already errored")
)

type Bundle struct {
	userOpHash common.Hash
	ops        *operation.BundleOperations

	atlasTxHashes []common.Hash
	relayErr      *relayerror.Error

	complete       bool
	completionSubs []chan *Bundle

	createdAt time.Time

	mu sync.RWMutex
}

func NewBundle(userOpHash common.Hash, bundleOps *operation.BundleOperations) *Bundle {
	return &Bundle{
		userOpHash:     userOpHash,
		ops:            bundleOps,
		atlasTxHashes:  make([]common.Hash, 0),
		completionSubs: make([]chan *Bundle, 0),
		createdAt:      time.Now(),
	}
}

func (b *Bundle) GetResult() (interface{}, *relayerror.Error) {
	b.mu.RLock()
	defer b.mu.RUnlock()
	if b.isMultiBundler() {
		return b.atlasTxHashes, nil
	}

	return b.atlasTxHashes[0], nil
}

func (b *Bundle) SetAtlasTxHash(txHash common.Hash, multiBundler bool) *relayerror.Error {
	b.mu.Lock()
	defer b.mu.Unlock()

	if !multiBundler && len(b.atlasTxHashes) == 1 {
		log.Info("bundle has already been bundled", "bundledTxHash", b.atlasTxHashes[0].Hex(), "newTxHash", txHash.Hex())
		return ErrAlreadyBundled
	}

	b.atlasTxHashes = append(b.atlasTxHashes, txHash)

	if !multiBundler {
		b.notifyCompletionSubs()
	}

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

func (b *Bundle) NotifyCompletionSubs() {
	b.notifyCompletionSubs()
}

func (b *Bundle) notifyCompletionSubs() {
	b.complete = true
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

// Return an array of hashes if the bundle is multi-bundler, otherwise return a single hash
func (b *Bundle) getAtlasTxHashes(completionChan chan *Bundle) (interface{}, *relayerror.Error) {
	atlasTxHashes, relayErr := b.GetResult()
	if relayErr != nil {
		return nil, relayErr
	}

	if completionChan != nil && !b.complete {
		b.addCompletionSub(completionChan)
		return nil, nil
	}

	if !b.complete {
		log.Info("bundle has not been completed yet", "userOpHash", b.userOpHash.Hex())
		return nil, ErrNotBundledYet
	}

	return atlasTxHashes, nil
}

func (b *Bundle) isMultiBundler() bool {
	return b.ops.DAppOperation.Bundler == common.Address{}
}
