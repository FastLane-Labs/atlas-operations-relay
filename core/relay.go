package core

import (
	"crypto/ecdsa"
	"math/big"

	"github.com/FastLane-Labs/atlas-operations-relay/auction"
	"github.com/FastLane-Labs/atlas-operations-relay/bundle"
	"github.com/FastLane-Labs/atlas-operations-relay/config"
	"github.com/FastLane-Labs/atlas-operations-relay/contract/atlETH"
	"github.com/FastLane-Labs/atlas-operations-relay/contract/atlasVerification"
	"github.com/FastLane-Labs/atlas-operations-relay/contract/storage"
	"github.com/FastLane-Labs/atlas-operations-relay/log"
	"github.com/FastLane-Labs/atlas-operations-relay/operation"
	"github.com/FastLane-Labs/atlas-operations-relay/relayerror"
	"github.com/FastLane-Labs/atlas-operations-relay/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

const (
	BundleSuccessfullySubmitted = "bundle successfully submitted"
)

var (
	ErrForwardBundle          = relayerror.NewError(3100, "failed to forward bundle")
	ErrCantGetDAppSignatories = relayerror.NewError(3101, "failed to get dapp signatories")
	ErrCantGetBondedBalance   = relayerror.NewError(3102, "failed to get atlEth bonded balance")
	ErrRelayIsOffline         = relayerror.NewError(3103, "relay is offline")
)

type Relay struct {
	config    *config.Config
	ethClient *ethclient.Client
	server    *Server

	auctionManager *auction.Manager
	bundleManager  *bundle.Manager

	atlETHContract            *atlETH.AtlETH
	storageContract           *storage.Storage
	atlasVerificationContract *atlasVerification.AtlasVerification

	atlasDomainSeparator common.Hash
}

func NewRelay(ethClient *ethclient.Client, config *config.Config) *Relay {
	atlETHContract, err := atlETH.NewAtlETH(config.Contracts.Atlas, ethClient)
	if err != nil {
		panic(err)
	}

	storageContract, err := storage.NewStorage(config.Contracts.Atlas, ethClient)
	if err != nil {
		panic(err)
	}

	atlasVerificationContract, err := atlasVerification.NewAtlasVerification(config.Contracts.AtlasVerification, ethClient)
	if err != nil {
		panic(err)
	}

	atlasDomainSeparator, err := atlasVerificationContract.GetDomainSeparator(nil)
	if err != nil {
		panic(err)
	}

	r := &Relay{
		config:                    config,
		ethClient:                 ethClient,
		atlETHContract:            atlETHContract,
		storageContract:           storageContract,
		atlasVerificationContract: atlasVerificationContract,
		atlasDomainSeparator:      atlasDomainSeparator,
	}

	r.auctionManager = auction.NewManager(ethClient, config, atlasDomainSeparator, r.solverGasLimit, r.balanceOfBonded, r.reputationScore, r.getDAppConfig, r.auctionCompleteCallback)
	r.bundleManager = bundle.NewManager(ethClient, config, atlasDomainSeparator, r.getDAppConfig, r.auctionManager.IsAuctionOngoing)
	r.server = NewServer(NewRouter(NewApi(r)), r.auctionManager.NewSolverOperation, r.auctionManager.GetSolverOperationStatus, r.getDAppSignatories)

	return r
}

func (r *Relay) Run(serverReadyChan chan struct{}) {
	r.server.ListenAndServe(serverReadyChan)
}

// Called when an auction is complete, bundleOps is missing the DAppOperation, we'll generate it if one of the relay's
// signatories is allowed by the dApp and submit the complete bundle.
func (r *Relay) auctionCompleteCallback(bundleOps *operation.BundleOperations) {
	dAppConfig, relayErr := r.getDAppConfig(bundleOps.UserOperation.Control, bundleOps.UserOperation)
	if relayErr != nil {
		log.Info("failed to get dapp config", "err", relayErr.Message)
		return
	}

	if utils.FlagUserAuctioneer(dAppConfig.CallConfig) || utils.FlagSolverAuctioneer(dAppConfig.CallConfig) || utils.FlagUnknownAuctioneer(dAppConfig.CallConfig) {
		// Another allowed auctioneer will generate the DAppOperation and submit the bundle
		return
	}

	userOpHash, relayErr := bundleOps.UserOperation.Hash(utils.FlagTrustedOpHash(dAppConfig.CallConfig))
	if relayErr != nil {
		log.Info("failed to compute user operation hash", "err", relayErr.Message)
		return
	}

	allowedSignatories, relayErr := r.getDAppSignatories(bundleOps.UserOperation.Control)
	if relayErr != nil {
		log.Info("failed to get dApp signatories", "control", bundleOps.UserOperation.Control.Hex(), "err", relayErr.Message)
		return
	}

	var (
		selectedSignatory   common.Address
		selectedSignatoryPk *ecdsa.PrivateKey
	)

	for _, signatory := range allowedSignatories {
		if pk, ok := r.config.Relay.Signatories[signatory]; ok {
			log.Info("signatory available", "signatory", signatory.Hex())
			selectedSignatory = signatory
			selectedSignatoryPk = pk
			break
		}
	}

	if selectedSignatoryPk == nil {
		// No local signatory available, the server will attempt to contact the appropiate signatory
		r.server.NewSignatoryRequest(userOpHash, bundleOps.UserOperation, bundleOps.SolverOperations, allowedSignatories, r.submitBundleOperations, r.bundleManager.RegisterBundleError)
		return
	}

	callChainHash, err := bundleOps.CallChainHash(dAppConfig.CallConfig, dAppConfig.To)
	if err != nil {
		log.Info("failed to compute call chain hash", "err", err)
		return
	}

	nonce, relayErr := r.getDappNextNonce(selectedSignatory)
	if relayErr != nil {
		log.Info("failed to get next nonce", "err", relayErr.Message)
		return
	}

	bundleOps.DAppOperation = &operation.DAppOperation{
		From:          selectedSignatory,
		To:            r.config.Contracts.Atlas,
		Nonce:         nonce,
		Deadline:      new(big.Int).Set(bundleOps.UserOperation.Deadline),
		Control:       bundleOps.UserOperation.Control,
		Bundler:       selectedSignatory, // Signatory is a whitelisted bundler and will be tried first
		UserOpHash:    userOpHash,
		CallChainHash: callChainHash,
	}

	proofHash, err := bundleOps.DAppOperation.ProofHash()
	if err != nil {
		log.Info("failed to compute dApp proof hash", "err", err)
		return
	}

	bundleOps.DAppOperation.Signature, err = utils.SignEip712Message(r.atlasDomainSeparator, proofHash, selectedSignatoryPk)
	if err != nil {
		log.Info("failed to sign dApp proof hash", "err", err)
		return
	}

	// All set to submit the bundle
	if _, relayErr := r.submitBundleOperations(bundleOps); relayErr != nil {
		log.Info("failed to submit bundle", "err", relayErr.Message)
	}
}

func (r *Relay) submitUserOperation(userOp *operation.UserOperation, hints []common.Address) (common.Hash, *relayerror.Error) {
	userOpHash, userOpPartialRaw, relayErr := r.auctionManager.NewUserOperation(userOp, hints)
	if relayErr != nil {
		return common.Hash{}, relayErr
	}

	go r.server.BroadcastUserOperationPartial(userOpPartialRaw)
	return userOpHash, nil
}

func (r *Relay) getSolverOperations(userOpHash common.Hash, completionChan chan []*operation.SolverOperationWithScore) (operation.SolverOperationsWithScore, *relayerror.Error) {
	return r.auctionManager.GetSolverOperations(userOpHash, completionChan)
}

func (r *Relay) submitBundleOperations(bundleOps *operation.BundleOperations) (string, *relayerror.Error) {
	userOpHash, bundle, relayErr := r.bundleManager.NewBundle(bundleOps)
	if relayErr != nil {
		return "", relayErr
	}

	if err := r.server.ForwardBundle(bundleOps, bundle.SetAtlasTxHash, bundle.SetRelayError); err != nil {
		r.bundleManager.UnregisterBundle(userOpHash)
		return "", ErrForwardBundle.AddError(err)
	}

	return BundleSuccessfullySubmitted, nil
}

func (r *Relay) getBundleHash(userOpHash common.Hash, completionChan chan *bundle.Bundle) (common.Hash, *relayerror.Error) {
	return r.bundleManager.GetBundleHash(userOpHash, completionChan)
}

func (r *Relay) submitSolverOperation(solverOp *operation.SolverOperation) (common.Hash, *relayerror.Error) {
	return r.auctionManager.NewSolverOperation(solverOp)
}

func (r *Relay) getSolverOperationStatus(solverOpHash common.Hash, completionChan chan *auction.SolverStatus) (*auction.SolverStatus, *relayerror.Error) {
	return r.auctionManager.GetSolverOperationStatus(solverOpHash, completionChan)
}
