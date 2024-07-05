package operation

import (
	"testing"

	"github.com/ethereum/go-ethereum/common"
)

func generateBundleOperations() *BundleOperations {
	return &BundleOperations{
		generateUserOperation(),
		[]*SolverOperation{
			generateSolverOperation(),
			generateSolverOperation(),
			generateSolverOperation(),
		},
		generateDAppOperation(),
	}
}

func TestBundleOperationsCallChainHashWithRequiredPreOps(t *testing.T) {
	t.Parallel()

	bundleOps := generateBundleOperations()
	want := common.HexToHash("0x65d9f83fb5c1b5fcce7fd60775bb6b40eccdd87d220ac4000f3e119d6c3d2a47")

	result, err := bundleOps.CallChainHash(4, common.HexToAddress("0x4"))
	if err != nil {
		t.Errorf("BundleOperations.callChainHash() error = %v", err)
	}

	if result != want {
		t.Errorf("BundleOperations.callChainHash() = %v, want %v", result, want)
	}
}

func TestBundleOperationsCallChainHashWithoutRequiredPreOps(t *testing.T) {
	t.Parallel()

	bundleOps := generateBundleOperations()
	want := common.HexToHash("0x7a856dd01991620aa2adcb177ef5cdcc9c1f524920d6791817f9272e131f869c")

	result, err := bundleOps.CallChainHash(0, common.HexToAddress("0x4"))
	if err != nil {
		t.Errorf("BundleOperations.callChainHash() error = %v", err)
	}

	if result != want {
		t.Errorf("BundleOperations.callChainHash() = %v, want %v", result, want)
	}
}
