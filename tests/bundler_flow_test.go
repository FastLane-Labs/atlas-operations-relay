package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"
	"time"

	"github.com/FastLane-Labs/atlas-operations-relay/auction"
	"github.com/FastLane-Labs/atlas-operations-relay/core"
	"github.com/FastLane-Labs/atlas-operations-relay/operation"
	"github.com/ethereum/go-ethereum/common"
)

func TestBundlerFlow(t *testing.T) {
	//start solver
	solverDoneChan := make(chan struct{})
	go runSolver(solverDoneChan)

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

	//wait for auction to end
	time.Sleep(auction.AuctionDuration)

	//user requests solver solutions
	userOpHash, _ := userOp.Hash()
	solverOps, err := retreiveSolverOps(userOpHash)
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

	//bundler sends bundleResponse
	err = sendBundleResposne(tx.Hash(), bundleRequest.Id, bundlerSendChan)
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

func sendBundleOperation(t *testing.T, userOp *operation.UserOperation, solverOps []*operation.SolverOperation, dAppOp *operation.DAppOperation) error {
	bundleOps := &operation.BundleOperations{
		UserOperation:    userOp,
		SolverOperations: solverOps,
		DAppOperation:    dAppOp,
	}

	bundleOpsJSON, err := json.Marshal(bundleOps)
	if err != nil {
		t.Errorf("failed to marshal bundle operations: %v", err)
	}

	_, err = http.Post("http://localhost:8080/bundleOperations", "application/json", bytes.NewReader(bundleOpsJSON))
	if err != nil {
		return err
	}

	return nil
}

func sendBundleResposne(txHash common.Hash, id string,bundlerSendChan chan []byte) error {
	bundleResponse := &core.BundleResponse{
		Id:     id,
		Result: txHash,
	}

	bundleResponseJSON, err := json.Marshal(bundleResponse)
	if err != nil {
		return err
	}

	bundlerSendChan <- bundleResponseJSON
	return nil
}
