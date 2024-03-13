package tests

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/FastLane-Labs/atlas-operations-relay/contract/atlasVerification"
	"github.com/FastLane-Labs/atlas-operations-relay/core"
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
