package core

import (
	"math/big"

	"github.com/FastLane-Labs/atlas-operations-relay/auction"
	"github.com/FastLane-Labs/atlas-operations-relay/bundle"
	"github.com/FastLane-Labs/atlas-operations-relay/config"
	"github.com/FastLane-Labs/atlas-operations-relay/contract/atlETH"
	"github.com/FastLane-Labs/atlas-operations-relay/contract/atlasVerification"
	"github.com/FastLane-Labs/atlas-operations-relay/operation"
	"github.com/FastLane-Labs/atlas-operations-relay/relayerror"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/ethclient"
)

const (
	SolverOpSuccessfullySubmitted = "solver operation successfully submitted"
	BundleSuccessfullySubmitted   = "bundle successfully submitted"
)

var (
	ErrForwardBundle          = relayerror.NewError(3100, "failed to forward bundle")
	ErrCantGetDAppSignatories = relayerror.NewError(3101, "failed to get dapp signatories")
	ErrCantGetBondedBalance   = relayerror.NewError(3102, "failed to get atlEth bonded balance")
)

type Relay struct {
	config *config.Config
	server *Server

	auctionManager *auction.Manager
	bundleManager  *bundle.Manager
}

func NewRelay(ethClient *ethclient.Client, config *config.Config) *Relay {
	atlETHContract, err := atlETH.NewAtlETH(config.Contracts.Atlas, ethClient)
	if err != nil {
		panic(err)
	}

	balanceOfBonded := func(account common.Address) (*big.Int, *relayerror.Error) {
		balance, err := atlETHContract.BalanceOfBonded(nil, account)
		if err != nil {
			return nil, ErrCantGetBondedBalance.AddError(err)
		}

		return balance, nil
	}

	atlasVerificationContract, err := atlasVerification.NewAtlasVerification(config.Contracts.AtlasVerification, ethClient)
	if err != nil {
		panic(err)
	}

	getDAppSignatories := func(dAppControl common.Address) ([]common.Address, *relayerror.Error) {
		signatories, err := atlasVerificationContract.GetDAppSignatories(nil, dAppControl)
		if err != nil {
			return []common.Address{}, ErrCantGetDAppSignatories.AddError(err)
		}

		return signatories, nil
	}

	atlasDomainSeparator, err := atlasVerificationContract.GetDomainSeparator(nil)
	if err != nil {
		panic(err)
	}

	auctionManager := auction.NewManager(ethClient, config, atlasDomainSeparator, balanceOfBonded)

	relay := &Relay{
		config:         config,
		auctionManager: auctionManager,
		bundleManager:  bundle.NewManager(ethClient, config, atlasDomainSeparator),
	}

	relay.server = NewServer(NewRouter(NewApi(relay)), auctionManager.NewSolverOperation, getDAppSignatories)
	return relay
}

func (r *Relay) Run(serverReadyChan chan struct{}) {
	r.server.ListenAndServe(serverReadyChan)
}

type UserOperationArgs struct {
	From         common.Address `json:"from"`
	To           common.Address `json:"to"`
	Value        *hexutil.Big   `json:"value"`
	Gas          hexutil.Uint64 `json:"gas"`
	MaxFeePerGas *hexutil.Big   `json:"maxFeePerGas"`
	Nonce        hexutil.Uint64 `json:"nonce"`
	Deadline     hexutil.Uint64 `json:"deadline"`
	Dapp         common.Address `json:"dapp"`
	Control      common.Address `json:"control"`
	SessionKey   common.Address `json:"sessionKey"`
	Data         hexutil.Bytes  `json:"data"`
	Signature    hexutil.Bytes  `json:"signature"`
}

func (r *Relay) submitUserOperation(args *UserOperationArgs) (common.Hash, *relayerror.Error) {
	userOp := &operation.UserOperation{
		From:         args.From,
		To:           args.To,
		Value:        args.Value.ToInt(),
		Gas:          new(big.Int).SetUint64(uint64(args.Gas)),
		MaxFeePerGas: args.MaxFeePerGas.ToInt(),
		Nonce:        new(big.Int).SetUint64(uint64(args.Nonce)),
		Deadline:     new(big.Int).SetUint64(uint64(args.Deadline)),
		Dapp:         args.Dapp,
		Control:      args.Control,
		SessionKey:   args.SessionKey,
		Data:         args.Data,
		Signature:    args.Signature,
	}

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

	if err := r.server.ForwardBundle(bundleOps, bundle.SetAtlasTxHash, bundle.SetRelayError); err != nil {
		userOphash, _ := bundleOps.UserOperation.Hash()
		r.bundleManager.UnregisterBundle(userOphash)
		return "", ErrForwardBundle.AddError(err)
	}

	return BundleSuccessfullySubmitted, nil
}

func (r *Relay) getBundleHash(userOpHash common.Hash, completionChan chan *bundle.Bundle) (common.Hash, *relayerror.Error) {
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
