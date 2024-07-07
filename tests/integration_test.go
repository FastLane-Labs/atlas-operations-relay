package tests

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"testing"

	"github.com/FastLane-Labs/atlas-operations-relay/core"
	"github.com/FastLane-Labs/atlas-operations-relay/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func TestMain(m *testing.M) {
	conf.Validate()

	var err error
	ethClient, err = ethclient.Dial(conf.Network.RpcUrl)
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
	go runSolver(true, solveUserOperation, make(chan struct{}))

	//start bundler
	bundlerReceiveChan := make(chan []byte)
	bundlerSendChan := make(chan []byte)
	go runBundler(bundlerPk, bundlerReceiveChan, bundlerSendChan)

	//send user request
	userOp := newDemoUserOperation(governanceEoa)
	err := sendUserRequest(userOp)
	if err != nil {
		t.Fatal(err)
	}

	//user requests solver solutions
	userOpHash, _ := userOp.Hash(utils.FlagTrustedOpHash(userOp.CallConfig), &conf.Relay.Eip712.Domain)
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
	tx, err := newAtlasTx(bundleRequest.Bundle.Decode())
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

func TestSolverHttp(t *testing.T) {
	//start solver
	go runSolver(false, solveUserOperation, make(chan struct{}))

	//send user request
	userOp := newDemoUserOperation(common.HexToAddress(fmt.Sprintf("0x%x", rand.Int63())))
	err := sendUserRequest(userOp)
	if err != nil {
		t.Fatal(err)
	}

	//user requests solver solutions
	userOpHash, _ := userOp.Hash(utils.FlagTrustedOpHash(userOp.CallConfig), &conf.Relay.Eip712.Domain)
	solverOps, err := retreiveSolverOps(userOpHash, true)
	if err != nil {
		t.Fatal(err)
	}

	if len(solverOps) == 0 {
		t.Fatal("expected at least one solver operation")
	}
}
