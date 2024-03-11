package tests

import (
	"context"
	"math/big"
	"os"
	"testing"

	"github.com/FastLane-Labs/atlas-operations-relay/config"
	"github.com/FastLane-Labs/atlas-operations-relay/contract/atlas"
	"github.com/FastLane-Labs/atlas-operations-relay/contract/atlasVerification"
	"github.com/FastLane-Labs/atlas-operations-relay/core"
	relayCrypto "github.com/FastLane-Labs/atlas-operations-relay/crypto"
	"github.com/FastLane-Labs/atlas-operations-relay/operation"
	"github.com/ethereum/go-ethereum/common"
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
		Nonce:        big.NewInt(4),
		MaxFeePerGas: big.NewInt(1000000000),
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
		MaxFeePerGas: big.NewInt(5000000000),
		Deadline:     userOp.Deadline,
		Solver:       simpleRfqSolver,
		Control:      userOp.Control,
		UserOpHash:   userOpHash,
		BidToken:     common.HexToAddress("0x0"),
		BidAmount:    big.NewInt(1e14),
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
