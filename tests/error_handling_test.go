package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/big"
	"math/rand"
	"net/http"
	"testing"

	"github.com/FastLane-Labs/atlas-operations-relay/operation"
	"github.com/FastLane-Labs/atlas-operations-relay/utils"
	"github.com/ethereum/go-ethereum/common"
)

func TestErrorHandling(t *testing.T) {
	if err := sendUserRequest(faultyUserOpBadSignature()); err == nil {
		t.Fatal("expected error on bad signature")
	}

	if err := sendUserRequest(faultyUserOpBadDeadline()); err == nil {
		t.Fatal("expected error on bad deadline")
	}

	if err := sendUserRequest(faultyUserOpGasTooHigh()); err == nil {
		t.Fatal("expected error on gas too high")
	}

	if err := sendUserRequest(faultyUserOpBadData()); err == nil {
		t.Fatal("expected error on bad data")
	}

	badSolverOpTest(t, solveUserOperationBadSignature)
	badSolverOpTest(t, solveUserOperationBadGas)
	badSolverOpTest(t, solveUserOperationBadData)
}

func TestBadPostReqJson(t *testing.T) {
	badJson := map[string]string{
		"ab": "cd",
	}

	reqJSON, err := json.Marshal(badJson)
	if err != nil {
		t.Fatal(err)
	}

	resp1, err := http.Post("http://localhost:8080/userOperation", "application/json", bytes.NewReader(reqJSON))
	if err != nil {
		t.Fatal(err)
	}

	if resp1.StatusCode != http.StatusBadRequest {
		t.Fatalf("expected status code %d, got %d", http.StatusBadRequest, resp1.StatusCode)
	}

	resp2, err := http.Post("http://localhost:8080/solverOperation", "application/json", bytes.NewReader(reqJSON))
	if err != nil {
		t.Fatal(err)
	}

	if resp2.StatusCode != http.StatusBadRequest {
		t.Fatalf("expected status code %d, got %d", http.StatusBadRequest, resp2.StatusCode)
	}

	resp3, err := http.Post("http://localhost:8080/bundleOperations", "application/json", bytes.NewReader(reqJSON))
	if err != nil {
		t.Fatal(err)
	}

	if resp3.StatusCode != http.StatusBadRequest {
		t.Fatalf("expected status code %d, got %d", http.StatusBadRequest, resp3.StatusCode)
	}
}

func badSolverOpTest(t *testing.T, solveUserOpFunc solveUserOpFunc) {
	solverDone := make(chan struct{})
	go runSolver(true, solveUserOpFunc, solverDone)

	//successful user operation
	userOp := newDemoUserOperation(common.HexToAddress(fmt.Sprintf("0x%x", rand.Int63())))
	if err := sendUserRequest(userOp); err != nil {
		t.Fatal(err)
	}

	userOpHash, _ := userOp.Hash(utils.FlagTrustedOpHash(userOp.CallConfig), &conf.Relay.Eip712.Domain)
	solverOps, err := retreiveSolverOps(userOpHash, true)
	if err != nil {
		t.Fatal(err)
	}

	if len(solverOps) != 0 {
		t.Fatal("expected zero solver operations")
	}

	close(solverDone)
}

func faultyUserOpBadSignature() *operation.UserOperation {
	userOp := newDemoUserOperation(common.HexToAddress(fmt.Sprintf("0x%x", rand.Int63())))
	userOp.Signature = []byte("bad signature")
	return userOp
}

func faultyUserOpBadDeadline() *operation.UserOperation {
	userOp := newDemoUserOperation(common.HexToAddress(fmt.Sprintf("0x%x", rand.Int63())))
	userOp.Deadline = big.NewInt(0)
	return userOp
}

func faultyUserOpGasTooHigh() *operation.UserOperation {
	userOp := newDemoUserOperation(common.HexToAddress(fmt.Sprintf("0x%x", rand.Int63())))
	userOp.MaxFeePerGas = big.NewInt(1000_000)
	return userOp
}

func faultyUserOpBadData() *operation.UserOperation {
	userOp := newDemoUserOperation(common.HexToAddress(fmt.Sprintf("0x%x", rand.Int63())))
	userOp.Data = []byte("bad data")
	return userOp
}

func solveUserOperationBadSignature(userOperationPartialRaw *operation.UserOperationPartialRaw, executionEnvironment common.Address) *operation.SolverOperation {
	solverOp := solveUserOperation(userOperationPartialRaw, executionEnvironment)
	solverOp.Signature = []byte("bad signature")
	return solverOp
}

func solveUserOperationBadGas(userOperationPartialRaw *operation.UserOperationPartialRaw, executionEnvironment common.Address) *operation.SolverOperation {
	solverOp := solveUserOperation(userOperationPartialRaw, executionEnvironment)
	solverOp.MaxFeePerGas = big.NewInt(0)
	return solverOp
}

func solveUserOperationBadData(userOperationPartialRaw *operation.UserOperationPartialRaw, executionEnvironment common.Address) *operation.SolverOperation {
	solverOp := solveUserOperation(userOperationPartialRaw, executionEnvironment)
	solverOp.Data = []byte("bad data")
	return solverOp
}
