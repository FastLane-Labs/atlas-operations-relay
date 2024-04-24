package bundle

import (
	"context"
	"sync"
	"time"

	"github.com/FastLane-Labs/atlas-operations-relay/config"
	"github.com/FastLane-Labs/atlas-operations-relay/contract"
	"github.com/FastLane-Labs/atlas-operations-relay/contract/dAppControl"
	"github.com/FastLane-Labs/atlas-operations-relay/log"
	"github.com/FastLane-Labs/atlas-operations-relay/operation"
	"github.com/FastLane-Labs/atlas-operations-relay/relayerror"
	"github.com/FastLane-Labs/atlas-operations-relay/utils"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	ErrBundleAlreadySubmitted = relayerror.NewError(5000, "bundle for this user operation has already been submitted")
	ErrBundleNotFound         = relayerror.NewError(5001, "bundle not found")
	ErrBundleFailedSimulation = relayerror.NewError(5002, "bundle failed simulation")
)

type getDAppConfigFn func(common.Address, *operation.UserOperation) (*dAppControl.DAppConfig, *relayerror.Error)

type Manager struct {
	ethClient *ethclient.Client
	config    *config.Config

	// Indexed by userOpHash
	bundles map[common.Hash]*Bundle

	atlasDomainSeparator common.Hash

	getDAppConfig getDAppConfigFn

	mu sync.RWMutex
}

func NewManager(ethClient *ethclient.Client, config *config.Config, atlasDomainSeparator common.Hash, getDAppConfig getDAppConfigFn) *Manager {
	bm := &Manager{
		ethClient:            ethClient,
		config:               config,
		bundles:              make(map[common.Hash]*Bundle),
		atlasDomainSeparator: atlasDomainSeparator,
		getDAppConfig:        getDAppConfig,
	}

	go bm.bundlesCleaner()
	return bm
}

func (bm *Manager) bundlesCleaner() {
	for range time.Tick(10 * time.Minute) {
		bm.mu.Lock()
		for userOpHash, bundle := range bm.bundles {
			if (bundle.atlasTxHash != (common.Hash{}) || bundle.relayErr != nil) && time.Since(bundle.createdAt) > time.Hour {
				delete(bm.bundles, userOpHash)
			}
		}
		bm.mu.Unlock()
	}
}

func (bm *Manager) NewBundle(bundleOps *operation.BundleOperations) (common.Hash, *Bundle, *relayerror.Error) {
	dAppConfig, relayErr := bm.getDAppConfig(bundleOps.DAppOperation.Control, bundleOps.UserOperation)
	if relayErr != nil {
		log.Info("failed to get dapp config", "err", relayErr.Message)
		return common.Hash{}, nil, relayErr
	}

	userOpHash, relayErr := bundleOps.UserOperation.Hash(utils.FlagTrustedOpHash(dAppConfig.CallConfig))
	if relayErr != nil {
		log.Info("failed to compute user operation hash", "err", relayErr.Message)
		return common.Hash{}, nil, relayErr
	}

	relayErr = bundleOps.Validate(bm.ethClient, userOpHash, bm.config.Contracts.Atlas, bm.atlasDomainSeparator, bm.config.Relay.Gas.MaxPerUserOperation, bm.config.Relay.Gas.MaxPerDAppOperation, dAppConfig)
	if relayErr != nil {
		log.Info("invalid dapp operation", "err", relayErr.Message, "userOpHash", userOpHash.Hex())
		return common.Hash{}, nil, relayErr
	}

	solverOps := make([]operation.SolverOperation, 0, len(bundleOps.SolverOperations))
	for _, solverOp := range bundleOps.SolverOperations {
		solverOps = append(solverOps, *solverOp)
	}

	pData, err := contract.AtlasAbi.Pack("metacall", *bundleOps.UserOperation, solverOps, *bundleOps.DAppOperation)
	if err != nil {
		log.Info("failed to pack bundle", "err", err, "userOpHash", userOpHash.Hex())
		return common.Hash{}, nil, relayerror.ErrServerInternal
	}

	_, err = bm.ethClient.CallContract(
		context.Background(),
		ethereum.CallMsg{
			From: bundleOps.DAppOperation.Bundler,
			To:   &bm.config.Contracts.Atlas,
			Data: pData,
		},
		nil,
	)

	if err != nil {
		log.Info("metacall simulation failed", "err", err, "userOpHash", userOpHash.Hex())
		return common.Hash{}, nil, ErrBundleFailedSimulation.AddError(err)
	}

	bm.mu.Lock()
	defer bm.mu.Unlock()

	if _, exist := bm.bundles[userOpHash]; exist {
		return common.Hash{}, nil, ErrBundleAlreadySubmitted
	}

	bundle := NewBundle(userOpHash, bundleOps)
	bm.bundles[userOpHash] = bundle
	return userOpHash, bundle, nil
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
