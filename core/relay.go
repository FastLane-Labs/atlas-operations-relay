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
	Gas          *hexutil.Big   `json:"gas"`
	MaxFeePerGas *hexutil.Big   `json:"maxFeePerGas"`
	Nonce        *hexutil.Big   `json:"nonce"`
	Deadline     *hexutil.Big   `json:"deadline"`
	Dapp         common.Address `json:"dapp"`
	Control      common.Address `json:"control"`
	SessionKey   common.Address `json:"sessionKey"`
	Data         hexutil.Bytes  `json:"data"`
	Signature    hexutil.Bytes  `json:"signature"`
}

func (args *UserOperationArgs) toUserOperation() *operation.UserOperation {
	return &operation.UserOperation{
		From:         args.From,
		To:           args.To,
		Value:        args.Value.ToInt(),
		Gas:          args.Gas.ToInt(),
		MaxFeePerGas: args.MaxFeePerGas.ToInt(),
		Nonce:        args.Nonce.ToInt(),
		Deadline:     args.Deadline.ToInt(),
		Dapp:         args.Dapp,
		Control:      args.Control,
		SessionKey:   args.SessionKey,
		Data:         args.Data,
		Signature:    args.Signature,
	}

}

type SolverOperationArgs struct {
	From         common.Address `json:"from"`
	To           common.Address `json:"to"`
	Value        *hexutil.Big   `json:"value"`
	Gas          *hexutil.Big   `json:"gas"`
	MaxFeePerGas *hexutil.Big   `json:"maxFeePerGas"`
	Deadline     *hexutil.Big   `json:"deadline"`
	Solver       common.Address `json:"solver"`
	Control      common.Address `json:"control"`
	UserOpHash   common.Hash    `json:"userOpHash"`
	BidToken     common.Address `json:"bidToken"`
	BidAmount    *hexutil.Big   `json:"bidAmount"`
	Data         hexutil.Bytes  `json:"data"`
	Signature    hexutil.Bytes  `json:"signature"`
}

func (args *SolverOperationArgs) toSolverOperation() *operation.SolverOperation {
	return &operation.SolverOperation{
		From:         args.From,
		To:           args.To,
		Value:        args.Value.ToInt(),
		Gas:          args.Gas.ToInt(),
		MaxFeePerGas: args.MaxFeePerGas.ToInt(),
		Deadline:     args.Deadline.ToInt(),
		Solver:       args.Solver,
		Control:      args.Control,
		UserOpHash:   args.UserOpHash,
		BidToken:     args.BidToken,
		BidAmount:    args.BidAmount.ToInt(),
		Data:         args.Data,
		Signature:    args.Signature,
	}
}

type DAppOperationArgs struct {
	From          common.Address `json:"from"`
	To            common.Address `json:"to"`
	Value         *hexutil.Big   `json:"value"`
	Gas           *hexutil.Big   `json:"gas"`
	Nonce         *hexutil.Big   `json:"nonce"`
	Deadline      *hexutil.Big   `json:"deadline"`
	Control       common.Address `json:"control"`
	Bundler       common.Address `json:"bundler"`
	UserOpHash    common.Hash    `json:"userOpHash"`
	CallChainHash common.Hash    `json:"callChainHash"`
	Signature     hexutil.Bytes  `json:"signature"`
}

func (args *DAppOperationArgs) toDAppOperation() *operation.DAppOperation {
	return &operation.DAppOperation{
		From:          args.From,
		To:            args.To,
		Value:         args.Value.ToInt(),
		Gas:           args.Gas.ToInt(),
		Nonce:         args.Nonce.ToInt(),
		Deadline:      args.Deadline.ToInt(),
		Control:       args.Control,
		Bundler:       args.Bundler,
		UserOpHash:    args.UserOpHash,
		CallChainHash: args.CallChainHash,
		Signature:     args.Signature,
	}
}

type BundleOperationArgs struct {
	UserOperation    *UserOperationArgs     `json:"userOperation"`
	SolverOperations []*SolverOperationArgs `json:"solverOperations"`
	DAppOperation    *DAppOperationArgs     `json:"dAppOperation"`
}

func (args *BundleOperationArgs) toBundleOperations() *operation.BundleOperations {
	var solverOps []*operation.SolverOperation
	for _, solverOpArgs := range args.SolverOperations {
		solverOps = append(solverOps, solverOpArgs.toSolverOperation())
	}
	return &operation.BundleOperations{
		UserOperation:    args.UserOperation.toUserOperation(),
		SolverOperations: solverOps,
		DAppOperation:    args.DAppOperation.toDAppOperation(),
	}
}

func (r *Relay) submitUserOperation(args *UserOperationArgs) (common.Hash, *relayerror.Error) {
	userOp := args.toUserOperation()

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

func (r *Relay) submitBundleOperations(args BundleOperationArgs) (string, *relayerror.Error) {
	bundleOps := args.toBundleOperations()
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

func (r *Relay) submitSolverOperation(args *SolverOperationArgs) (string, *relayerror.Error) {
	var result string

	relayErr := r.auctionManager.NewSolverOperation(args.toSolverOperation())
	if relayErr == nil {
		result = SolverOpSuccessfullySubmitted
	}

	return result, relayErr
}
