package tests

import (
	"math/big"
	"testing"

	"github.com/FastLane-Labs/atlas-operations-relay/operation"
)

func TestUserOperation(t *testing.T) {
	if err := sendUserRequest(faultyUserOpBadSignature()); err == nil {
		t.Fatal("expected error on bad signature")
	}

	if err := sendUserRequest(faultyUserOpBadDeadline()); err == nil {
		t.Fatal("expected error on bad deadline")
	}

	if err := sendUserRequest(faultyUserOpGasTooHigh()); err == nil {
		t.Fatal("expected error on gas too high")
	}
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
