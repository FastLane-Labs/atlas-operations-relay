package tests

import (
	"encoding/json"
	"os"
	"testing"
	"time"

	"github.com/FastLane-Labs/atlas-operations-relay/config"
	"github.com/FastLane-Labs/atlas-operations-relay/contract/atlasVerification"
	"github.com/FastLane-Labs/atlas-operations-relay/core"
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

	chainId int64 = 11155111 //sepolia

	sendAtlasTx bool = false //true -> send tx to the network, false -> do everything apart from sending

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
		Relay: config.Relay{
			Auction: config.Auction{
				Duration: 3 * time.Second,
			},
		},
	}
)

func TestMain(m *testing.M) {
	conf.Validate()

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

func TestIntegration(t *testing.T) {
	//start solver
	go runSolver(true)

	//start bundler
	bundlerReceiveChan := make(chan []byte)
	bundlerSendChan := make(chan []byte)
	go runBundler(bundlerPk, bundlerReceiveChan, bundlerSendChan)

	//send user request
	userOp, err := sendUserRequest()
	if err != nil {
		t.Fatal(err)
	}

	//user requests solver solutions
	userOpHash, _ := userOp.Hash()
	solverOps, err := retreiveSolverOps(userOpHash, true)
	if err != nil {
		t.Fatal(err)
	}

	if len(solverOps) == 0 {
		t.Fatal("expected at least one solver operation")
	}

	//send bundleOps to the relay
	err = sendBundleOperation(userOp, solverOps, newDappOperation(userOp, solverOps))
	if err != nil {
		t.Fatal(err)
	}

	//wait till bundler receives bundleOps
	bundleOpBytes := <-bundlerReceiveChan
	bundleRequest := &core.BundleRequest{}
	err = json.Unmarshal(bundleOpBytes, bundleRequest)
	if err != nil {
		t.Fatal(err)
	}

	//send atlas tx
	tx, err := newAtlasTx(bundleRequest)
	if err != nil {
		t.Fatal(err)
	}

	//bundler sends bundlerResponse
	err = sendBundlerResponse(tx.Hash(), bundleRequest.Id, bundlerSendChan)
	if err != nil {
		t.Fatal(err)
	}

	//user asks relay for txHash
	txHash, err := retrieveAtlasTxHash(userOpHash, true)
	if err != nil {
		t.Fatal(err)
	}

	if txHash != tx.Hash() {
		t.Fatalf("expected txHash %s, got %s", tx.Hash(), txHash)
	}
}
