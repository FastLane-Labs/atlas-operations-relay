package tests

import (
	"math/big"
	"testing"

	"github.com/FastLane-Labs/atlas-operations-relay/operation"
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

func badSolverOpTest(t *testing.T, solveUserOpFunc solveUserOpFunc) {
	solverDone := make(chan struct{})
	go runSolver(true, solveUserOpFunc, solverDone)

	//successful user operation
	userOp := newDemoUserOperation()
	if err := sendUserRequest(userOp); err != nil {
		t.Fatal(err)
	}

	userOpHash, _ := userOp.Hash()
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
	userOp := newDemoUserOperation()
	userOp.Signature = []byte("bad signature")
	return userOp
}

func faultyUserOpBadDeadline() *operation.UserOperation {
	userOp := newDemoUserOperation()
	userOp.Deadline = big.NewInt(0)
	return userOp
}

func faultyUserOpGasTooHigh() *operation.UserOperation {
	userOp := newDemoUserOperation()
	userOp.MaxFeePerGas = big.NewInt(1000_000)
	return userOp
}

func faultyUserOpBadData() *operation.UserOperation {
	userOp := newDemoUserOperation()
	userOp.Data = []byte("bad data")
	return userOp
}

func solveUserOperationBadSignature(solverInput *operation.SolverInput, executionEnvironment common.Address) *operation.SolverOperation {
	solverOp := solveUserOperation(solverInput, executionEnvironment)
	solverOp.Signature = []byte("bad signature")
	return solverOp
}

func solveUserOperationBadGas(solverInput *operation.SolverInput, executionEnvironment common.Address) *operation.SolverOperation {
	solverOp := solveUserOperation(solverInput, executionEnvironment)
	solverOp.MaxFeePerGas = big.NewInt(0)
	return solverOp
}

func solveUserOperationBadData(solverInput *operation.SolverInput, executionEnvironment common.Address) *operation.SolverOperation {
	solverOp := solveUserOperation(solverInput, executionEnvironment)
	solverOp.Data = []byte("bad data")
	return solverOp
}
