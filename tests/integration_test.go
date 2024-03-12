package tests

import (
	"encoding/json"
	"testing"

	"github.com/FastLane-Labs/atlas-operations-relay/core"
)

func TestIntegration(t *testing.T) {
	//start solver
	solverDoneChan := make(chan struct{})
	go runSolver(solverDoneChan, true)

	//start bundler
	bundlerReceiveChan := make(chan []byte)
	bundlerSendChan := make(chan []byte)
	go runBundler(bundlerPk, bundlerReceiveChan, bundlerSendChan)

	//send user request
	userOp, err := sendUserRequest()
	if err != nil {
		t.Fatal(err)
	}

	//wait for solver to finish
	<-solverDoneChan

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
	err = sendBundleOperation(t, userOp, solverOps, NewDappOperation(userOp, solverOps))
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
	tx, err := NewAtlasTx(bundleRequest)
	if err != nil {
		t.Fatal(err)
	}

	//bundler sends bundlerResponse
	err = sendBundlerResposne(tx.Hash(), bundleRequest.Id, bundlerSendChan)
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
