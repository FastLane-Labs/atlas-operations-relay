package tests

import (
	"context"
	"fmt"
	"math/big"
	"os"
	"testing"

	"github.com/FastLane-Labs/atlas-operations-relay/config"
	"github.com/FastLane-Labs/atlas-operations-relay/contract"
	"github.com/FastLane-Labs/atlas-operations-relay/contract/atlas"
	"github.com/FastLane-Labs/atlas-operations-relay/contract/atlasVerification"
	"github.com/FastLane-Labs/atlas-operations-relay/contract/dAppControl"
	"github.com/FastLane-Labs/atlas-operations-relay/core"
	relayCrypto "github.com/FastLane-Labs/atlas-operations-relay/crypto"
	"github.com/FastLane-Labs/atlas-operations-relay/operation"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

/*

Dev note:

All EOAs and contracts used here have been deployed/integrated/funded on Sepolia for the sole purpose of the integration tests.
Changing even one of those values will most likely break the tests.

For reference, here are the necessary components:

EOAs:
- user: has at least 1e12 tokenB, as defined in the demo swap intent data, and has approved Atlas to spend it (permit69)
- solver: has bonded at least 0.5 ATLETH on Atlas

Contracts:
- Atlas
- AtlasVerification
- Simulator
The above 3 contracts are dependent on each others.

- swapIntentDAppControl: is integrated and dependent on Atlas
- SimpleRfqSolver: has at least 0.01 ETH, and 200e12 tokenA, as defined in the demo swap intent

*/

var (
	ethClient            *ethclient.Client
	atlasDomainSeparator common.Hash

	tokenA                = common.HexToAddress("0x7439E9Bb6D8a84dd3A23fe621A30F95403F87fB9")
	tokenB                = common.HexToAddress("0x7b79995e5f793A07Bc00c21412e50Ecae098E7f9")
	swapIntentDAppControl = common.HexToAddress("0x2F98675731D0e659E716d890330901a8A2355813")
	simpleRfqSolver       = common.HexToAddress("0xbB4026B66fB10CF380Cb069F5ac517Ed5caAdc34")

	userPk, _ = crypto.ToECDSA(common.FromHex("0x54c371fc7e513a4574cf87ca222664580e9c3fa58ecc636a2bd811bddfac66a2"))
	userEoa   = crypto.PubkeyToAddress(userPk.PublicKey) // 0x7F14D219C66c53A638185146A20555bb6a3f234A

	solverPk, _ = crypto.ToECDSA(common.FromHex("e3818ba09a3106ecbf80cafed7510df6f81e6dfab3cfa240680ec9389ba66206"))
	solverEoa   = crypto.PubkeyToAddress(solverPk.PublicKey) // 0x915c0f7A6a3be9E2298De421C4FE6CA1E2194880

	bundlerPk, _ = crypto.ToECDSA(common.FromHex("586af44cb6f500bdcdbea0e4411916dfc4806e7df43504da5bdfe144dd78f895"))
	bundlerEoa   = crypto.PubkeyToAddress(bundlerPk.PublicKey) // 0xEdB89106f2293ed2bAAbA1e8E844306412cB39Fe

	conf = &config.Config{
		Network: config.Network{
			RpcUrl: "https://rpc.sepolia.org/",
		},
		Contracts: config.Contracts{
			Atlas:             common.HexToAddress("0x090A1134ebF3066f87dB165a910E832a87C2Ffb5"),
			AtlasVerification: common.HexToAddress("0x15261B0437b0e54d6f0Dd308148a359Bb0191A81"),
			Simulator:         common.HexToAddress("0x8F4C1158fA69f973973b6ba7A1b3F58EE9204d1A"),
		},
	}
)

func TestMain(m *testing.M) {
	var err error
	ethClient, err = ethclient.Dial(conf.Network.RpcUrl)
	if err != nil {
		panic(err)
	}

	atlasVerification, err := atlasVerification.NewAtlasVerification(conf.Contracts.AtlasVerification, ethClient)
	if err != nil {
		panic(err)
	}

	atlasDomainSeparator, err = atlasVerification.GetDomainSeparator(nil)
	if err != nil {
		panic(err)
	}

	serverReadyChan := make(chan struct{})
	// Start the relay
	go core.StartRelay(conf, serverReadyChan)
	// Wait for the server to be ready
	<-serverReadyChan
	os.Exit(m.Run())
}

func NewDemoSwapIntent() *SwapIntent {
	return &SwapIntent{
		TokenUserBuys:          tokenA,
		AmountUserBuys:         big.NewInt(200e12),
		TokenUserSells:         tokenB,
		AmountUserSells:        big.NewInt(1e12),
		AuctionBaseCurrency:    common.HexToAddress("0x0"),
		SolverMustReimburseGas: false,
		Conditions:             make([]Condition, 0),
	}
}

func NewDemoUserOperation() *operation.UserOperation {
	currentBlock, err := ethClient.BlockNumber(context.Background())
	if err != nil {
		panic(err)
	}

	data, err := NewDemoSwapIntent().abiEncodeWithSelector()
	if err != nil {
		panic(err)
	}

	userOp := &operation.UserOperation{
		From:         userEoa,
		To:           conf.Contracts.Atlas,
		Deadline:     big.NewInt(int64(currentBlock) + 1000),
		Gas:          big.NewInt(100000),
		Nonce:        big.NewInt(5),
		MaxFeePerGas: big.NewInt(20e9),
		Value:        big.NewInt(0),
		Dapp:         swapIntentDAppControl,
		Control:      swapIntentDAppControl,
		SessionKey:   common.HexToAddress("0x0"),
		Data:         data,
		Signature:    nil,
	}

	proofHash, err := userOp.ProofHash()
	if err != nil {
		panic(err)
	}

	userOp.Signature = relayCrypto.SignEip712(atlasDomainSeparator, proofHash, userPk)

	if err := userOp.Validate(ethClient, conf.Contracts.Atlas, atlasDomainSeparator, nil); err != nil {
		panic(err)
	}

	return userOp
}

func SolveUserOperation(userOp *operation.UserOperation, executionEnvironment common.Address) *operation.SolverOperation {
	userOpHash, relayErr := userOp.Hash()
	if relayErr != nil {
		panic(relayErr)
	}

	swapIntent, err := swapIntentAbiDecode(userOp.Data)
	if err != nil {
		panic(err)
	}

	data, err := solverData(swapIntent, executionEnvironment)
	if err != nil {
		panic(err)
	}

	solverOp := &operation.SolverOperation{
		From:         solverEoa,
		To:           conf.Contracts.Atlas,
		Value:        big.NewInt(0),
		Gas:          big.NewInt(100000),
		MaxFeePerGas: big.NewInt(0).Add(userOp.MaxFeePerGas, big.NewInt(1e9)),
		Deadline:     userOp.Deadline,
		Solver:       simpleRfqSolver,
		Control:      userOp.Control,
		UserOpHash:   userOpHash,
		BidToken:     common.Address{},
		BidAmount:    big.NewInt(1e13),
		Data:         data,
		Signature:    nil,
	}

	proofHash, err := solverOp.ProofHash()
	if err != nil {
		panic(err)
	}

	solverOp.Signature = relayCrypto.SignEip712(atlasDomainSeparator, proofHash, solverPk)

	if err := solverOp.Validate(userOp, conf.Contracts.Atlas, atlasDomainSeparator, nil); err != nil {
		panic(err)
	}

	return solverOp
}

func ExecutionEnvironment(user common.Address, dAppControl common.Address) common.Address {
	atlasContract, err := atlas.NewAtlas(conf.Contracts.Atlas, ethClient)
	if err != nil {
		panic(err)
	}

	executionEnvironment, err := atlasContract.GetExecutionEnvironment(nil, user, dAppControl)
	if err != nil {
		panic(err)
	}

	return executionEnvironment.ExecutionEnvironment
}

func NewDappOperation(userOp *operation.UserOperation, solverOps []*operation.SolverOperation) *operation.DAppOperation {
	userOpHash, _ := userOp.Hash()

	dAppControlContract, err := dAppControl.NewDAppControl(userOp.Control, ethClient)
	if err != nil {
		panic(err)
	}

	dAppConfig, err := dAppControlContract.GetDAppConfig(nil, dAppControl.UserOperation(*userOp))
	if err != nil {
		panic(err)
	}

	callChainHash, err := CallChainHash(dAppConfig.CallConfig, dAppConfig.To, userOp, solverOps)
	if err != nil {
		panic(err)
	}

	dAppOp := &operation.DAppOperation{
		From:          userOp.From,
		To:            conf.Contracts.Atlas,
		Value:         big.NewInt(0),
		Gas:           big.NewInt(100000),
		Nonce:         big.NewInt(0),
		Deadline:      userOp.Deadline,
		Control:       userOp.Control,
		Bundler:       bundlerEoa,
		UserOpHash:    userOpHash,
		CallChainHash: callChainHash,
		Signature:     []byte(""),
	}

	proofHash, err := dAppOp.ProofHash()
	if err != nil {
		panic(err)
	}

	//user signs the DappOp
	dAppOp.Signature = relayCrypto.SignEip712(atlasDomainSeparator, proofHash, userPk)

	return dAppOp
}

func NewAtlasTx(bundleRequest *core.BundleRequest) (*types.Transaction, error) {
	atlasContract, err := atlas.NewAtlas(conf.Contracts.Atlas, ethClient)
	if err != nil {
		panic(err)
	}

	signFn := func(addr common.Address, tx *types.Transaction) (*types.Transaction, error) {
		return types.SignTx(tx, types.NewEIP155Signer(big.NewInt(int64(11155111))), bundlerPk)
	}

	gasPrice, err := ethClient.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, fmt.Errorf("Can't get gas price suggestion: %w", err)
	}

	opts := &bind.TransactOpts{
		From:     bundlerEoa,
		GasPrice: gasPrice,
		Signer:   signFn,
		Nonce:    nil,
		Value:    nil,
		GasLimit: uint64(5_000_000),
		NoSend:   true,
	}

	atlas_userOp := atlas.UserOperation{
		From:         bundleRequest.Bundle.UserOperation.From,
		To:           bundleRequest.Bundle.UserOperation.To,
		Value:        bundleRequest.Bundle.UserOperation.Value,
		Gas:          bundleRequest.Bundle.UserOperation.Gas,
		MaxFeePerGas: bundleRequest.Bundle.UserOperation.MaxFeePerGas,
		Nonce:        bundleRequest.Bundle.UserOperation.Nonce,
		Deadline:     bundleRequest.Bundle.UserOperation.Deadline,
		Dapp:         bundleRequest.Bundle.UserOperation.Dapp,
		Control:      bundleRequest.Bundle.UserOperation.Control,
		SessionKey:   bundleRequest.Bundle.UserOperation.SessionKey,
		Data:         bundleRequest.Bundle.UserOperation.Data,
		Signature:    bundleRequest.Bundle.UserOperation.Signature,
	}

	atlas_dappOp := atlas.DAppOperation{
		From:          bundleRequest.Bundle.DAppOperation.From,
		To:            bundleRequest.Bundle.DAppOperation.To,
		Value:         bundleRequest.Bundle.DAppOperation.Value,
		Gas:           bundleRequest.Bundle.DAppOperation.Gas,
		Nonce:         bundleRequest.Bundle.DAppOperation.Nonce,
		Deadline:      bundleRequest.Bundle.DAppOperation.Deadline,
		Control:       bundleRequest.Bundle.DAppOperation.Control,
		Bundler:       bundleRequest.Bundle.DAppOperation.Bundler,
		UserOpHash:    bundleRequest.Bundle.DAppOperation.UserOpHash,
		CallChainHash: bundleRequest.Bundle.DAppOperation.CallChainHash,
		Signature:     bundleRequest.Bundle.DAppOperation.Signature,
	}

	atlas_solverOps := make([]atlas.SolverOperation, len(bundleRequest.Bundle.SolverOperations))
	for i, solverOp := range bundleRequest.Bundle.SolverOperations {
		atlas_solverOps[i] = atlas.SolverOperation{
			From:         solverOp.From,
			To:           solverOp.To,
			Value:        solverOp.Value,
			Gas:          solverOp.Gas,
			MaxFeePerGas: solverOp.MaxFeePerGas,
			Deadline:     solverOp.Deadline,
			Solver:       solverOp.Solver,
			Control:      solverOp.Control,
			UserOpHash:   solverOp.UserOpHash,
			BidToken:     solverOp.BidToken,
			BidAmount:    solverOp.BidAmount,
			Data:         solverOp.Data,
			Signature:    solverOp.Signature,
		}
	}

	return atlasContract.Metacall(opts, atlas_userOp, atlas_solverOps, atlas_dappOp)
}

func CallChainHash(callConfig uint32, dAppControl common.Address, userOp *operation.UserOperation, solverOps []*operation.SolverOperation) (common.Hash, error) {
	counter := big.NewInt(0)
	var callSequenceHash common.Hash

	if callConfig&4 != 0 {
		// Require preOps
		preOpsEncoded, err := contract.DappControlAbi.Pack("preOpsCall", userOp)
		if err != nil {
			return common.Hash{}, err
		}

		callSequenceHash = crypto.Keccak256Hash(
			callSequenceHash.Bytes(),
			dAppControl.Bytes(),
			preOpsEncoded,
			math.U256Bytes(counter),
		)

		counter.Add(counter, common.Big1)
	}

	userOpAbiEncoded, err := userOp.AbiEncode()
	if err != nil {
		return common.Hash{}, err
	}

	callSequenceHash = crypto.Keccak256Hash(
		callSequenceHash.Bytes(),
		userOpAbiEncoded,
		math.U256Bytes(counter),
	)

	counter.Add(counter, common.Big1)

	for _, solverOp := range solverOps {
		solverOpAbiEncoded, err := solverOp.AbiEncode()
		if err != nil {
			return common.Hash{}, err
		}

		callSequenceHash = crypto.Keccak256Hash(
			callSequenceHash.Bytes(),
			solverOpAbiEncoded,
			math.U256Bytes(counter),
		)

		counter.Add(counter, common.Big1)
	}

	return callSequenceHash, nil
}
