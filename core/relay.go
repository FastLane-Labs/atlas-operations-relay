package core

import (
	"github.com/FastLane-Labs/atlas-operations-relay/auction"
	"github.com/FastLane-Labs/atlas-operations-relay/bundle"
	"github.com/FastLane-Labs/atlas-operations-relay/config"
	"github.com/FastLane-Labs/atlas-operations-relay/contract/atlETH"
	"github.com/FastLane-Labs/atlas-operations-relay/contract/atlasVerification"
	"github.com/FastLane-Labs/atlas-operations-relay/contract/storage"
	"github.com/FastLane-Labs/atlas-operations-relay/operation"
	"github.com/FastLane-Labs/atlas-operations-relay/relayerror"
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
	}

	r.auctionManager = auction.NewManager(ethClient, config, atlasDomainSeparator, r.solverGasLimit, r.balanceOfBonded, r.reputationScore, r.getDAppConfig)
	r.bundleManager = bundle.NewManager(ethClient, config, atlasDomainSeparator, r.getDAppConfig)
	r.server = NewServer(NewRouter(NewApi(r)), r.auctionManager.NewSolverOperation, r.auctionManager.GetSolverOperationStatus, r.getDAppSignatories)

	return r
}

func (r *Relay) Run(serverReadyChan chan struct{}) {
	r.server.ListenAndServe(serverReadyChan)
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
