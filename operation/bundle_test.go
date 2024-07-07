package operation

import (
	"testing"

	"github.com/FastLane-Labs/atlas-operations-relay/utils"
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

func TestBundleOperationsCallChainHashWithRequiredPreOpsAndWithSolverOps(t *testing.T) {
	t.Parallel()

	bundleOps := generateBundleOperations(3)
	want := common.HexToHash("0xce835ea8087710762b1d392a3225f5cb50adb278093945e1835b3cc5f3033a82")

	result, err := bundleOps.CallChainHash(uint32(1)<<uint32(utils.IndexRequirePreOps), common.HexToAddress("0x4"))
	if err != nil {
		t.Errorf("BundleOperations.callChainHash() error = %v", err)
	}

	if result != want {
		t.Errorf("BundleOperations.callChainHash() = %v, want %v", result, want)
	}
}

func TestBundleOperationsCallChainHashWithRequiredPreOpsAndWithoutSolverOps(t *testing.T) {
	t.Parallel()

	bundleOps := generateBundleOperations(0)
	want := common.HexToHash("0x38402f35236801c6382b2a79ff78e24d7208744b7253bfb99a8ab19bcab8f824")

	result, err := bundleOps.CallChainHash(uint32(1)<<uint32(utils.IndexRequirePreOps), common.HexToAddress("0x4"))
	if err != nil {
		t.Errorf("BundleOperations.callChainHash() error = %v", err)
	}

	if result != want {
		t.Errorf("BundleOperations.callChainHash() = %v, want %v", result, want)
	}
}

func TestBundleOperationsCallChainHashWithoutRequiredPreOpsAndWithSolverOps(t *testing.T) {
	t.Parallel()

	bundleOps := generateBundleOperations(3)
	want := common.HexToHash("0x8a71f907fe61688772ede6e7bb91efa992fe86c27917862adf533984dd56a2b8")

	result, err := bundleOps.CallChainHash(0, common.HexToAddress("0x4"))
	if err != nil {
		t.Errorf("BundleOperations.callChainHash() error = %v", err)
	}

	if result != want {
		t.Errorf("BundleOperations.callChainHash() = %v, want %v", result, want)
	}
}

func TestBundleOperationsCallChainHashWithoutRequiredPreOpsAndWithoutSolverOps(t *testing.T) {
	t.Parallel()

	bundleOps := generateBundleOperations(0)
	want := common.HexToHash("0x1feca496343f60c6fd5bfa97ec935fed62285b814ef720ac633dabb1c6e25777")

	result, err := bundleOps.CallChainHash(0, common.HexToAddress("0x4"))
	if err != nil {
		t.Errorf("BundleOperations.callChainHash() error = %v", err)
	}

	if result != want {
		t.Errorf("BundleOperations.callChainHash() = %v, want %v", result, want)
	}
}
