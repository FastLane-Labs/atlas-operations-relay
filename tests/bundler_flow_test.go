package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"
	"time"

	"github.com/FastLane-Labs/atlas-operations-relay/auction"
	"github.com/FastLane-Labs/atlas-operations-relay/operation"
)

func TestBundlerFlow(t *testing.T) {
	//start solver
	solverDoneChan := make(chan struct{})
	go runSolver(solverDoneChan)

	//start bundler
	go runBundler(bundlerPk)

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
