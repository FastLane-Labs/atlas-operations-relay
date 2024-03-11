package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/FastLane-Labs/atlas-operations-relay/auction"
	"github.com/FastLane-Labs/atlas-operations-relay/operation"
	"github.com/ethereum/go-ethereum/crypto"
)

func TestBundlerFlow(t *testing.T) {
	//start bundler
	sampleBundlerPk, err := crypto.GenerateKey()
	if err != nil {
		panic(err)
	}

	startBundler(t, sampleBundlerPk)

	//start solver
	solverDoneChan := make(chan struct{})

	go runSolver(solverDoneChan)
	userOp, err := sendUserRequest()
	if err != nil {
		t.Fatal(err)
	}

	<-solverDoneChan

	//wait for auction to end
	time.Sleep(auction.AuctionDuration)

	//user requests for solver solutions
	userOpHash, _ := userOp.Hash()
	solverOps, err := retreiveSolverOps(userOpHash)
	if err != nil {
		t.Fatal(err)
	}

	if len(solverOps) == 0 {
		t.Fatal("expected at least one solver operation")
	}

	//send bundleOps to the relay
	sendBundleOperation(t, userOp, solverOps, NewDappOperation(userOp, solverOps))

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

	resp, err := http.Post("http://localhost:8080/bundleOperations", "application/json", bytes.NewReader(bundleOpsJSON))
	if err != nil {
		return err
	}

	fmt.Println("Response Status:", resp.Status)
	fmt.Println("Response Headers:", resp.Header)
	fmt.Println("Response Body:", resp.Body)

	return nil
}
