package operation

import (
	"testing"

	"github.com/ethereum/go-ethereum/common"
)

func generateBundleOperations(solverOpsCount int) *BundleOperations {
	solverOps := make([]*SolverOperation, solverOpsCount)
	for i := 0; i < solverOpsCount; i++ {
		solverOps[i] = generateSolverOperation()
	}

	return &BundleOperations{
		generateUserOperation(),
		solverOps,
		generateDAppOperation(),
	}
}

func TestCallChainHashWithSolverOperations(t *testing.T) {
	t.Parallel()

	bundleOps := generateBundleOperations(3)
	want := common.HexToHash("0x8a71f907fe61688772ede6e7bb91efa992fe86c27917862adf533984dd56a2b8")

	result, err := bundleOps.CallChainHash()
	if err != nil {
		t.Errorf("CallChainHash() error = %v", err)
	}

	if result != want {
		t.Errorf("CallChainHash() = %v, want %v", result, want)
	}
}

func TestCallChainHashWithoutSolverOperations(t *testing.T) {
	t.Parallel()

	bundleOps := generateBundleOperations(0)
	want := common.HexToHash("0x1feca496343f60c6fd5bfa97ec935fed62285b814ef720ac633dabb1c6e25777")

	result, err := bundleOps.CallChainHash()
	if err != nil {
		t.Errorf("CallChainHash() error = %v", err)
	}

	if result != want {
		t.Errorf("CallChainHash() = %v, want %v", result, want)
	}
}
