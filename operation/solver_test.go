package operation

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
)

func generateSolverOperation() *SolverOperation {
	return &SolverOperation{
		From:         common.HexToAddress("0x1"),
		To:           common.HexToAddress("0x2"),
		Value:        big.NewInt(100),
		Gas:          big.NewInt(200),
		MaxFeePerGas: big.NewInt(300),
		Deadline:     big.NewInt(400),
		Solver:       common.HexToAddress("0x3"),
		Control:      common.HexToAddress("0x4"),
		UserOpHash:   common.HexToHash("0x9999999999999999999999999999999999999999999999999999999999999999"),
		BidToken:     common.HexToAddress("0x5"),
		BidAmount:    big.NewInt(500),
		Data:         []byte("data"),
		Signature:    []byte("signature"),
	}
}

func TestSolverOperationAbiEncode(t *testing.T) {
	t.Parallel()

	solverOp := generateSolverOperation()
	want := common.FromHex("0x000000000000000000000000000000000000000000000000000000000000002000000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000002000000000000000000000000000000000000000000000000000000000000006400000000000000000000000000000000000000000000000000000000000000c8000000000000000000000000000000000000000000000000000000000000012c0000000000000000000000000000000000000000000000000000000000000190000000000000000000000000000000000000000000000000000000000000000300000000000000000000000000000000000000000000000000000000000000049999999999999999999999999999999999999999999999999999999999999999000000000000000000000000000000000000000000000000000000000000000500000000000000000000000000000000000000000000000000000000000001f400000000000000000000000000000000000000000000000000000000000001a000000000000000000000000000000000000000000000000000000000000001e00000000000000000000000000000000000000000000000000000000000000004646174610000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000097369676e61747572650000000000000000000000000000000000000000000000")

	result, err := solverOp.AbiEncode()
	if err != nil {
		t.Errorf("SolverOperation.abiEncode() error = %v", err)
	}

	if string(result) != string(want) {
		t.Errorf("SolverOperation.abiEncode() = %v, want %v", result, want)
	}
}

func TestSolverOperationProofHash(t *testing.T) {
	t.Parallel()

	solverOp := generateSolverOperation()
	want := common.HexToHash("0xf50a1c5bfaa8fbe671ccd340a43661b90c70cdbde9ea42254f7a0606ef39ab79")

	result, err := solverOp.ProofHash()
	if err != nil {
		t.Errorf("SolverOperation.proofHash() error = %v", err)
	}

	if result != want {
		t.Errorf("SolverOperation.proofHash() = %v, want %v", result, want)
	}
}

func TestSolverOperationCheckSignature(t *testing.T) {
	t.Parallel()

	solverOp := generateSolverOperation()
	solverOp.From = common.HexToAddress("0xbab20eE47D62C288EBAFd1A938B193561E48C1Da")
	solverOp.Signature = common.FromHex("0xb6c553426561570bd818ffb611f5182a9b0ce672aa2c73568a8fc5af7bbe71dc596141b7ac63286bba9145cc052dedef1874baea53ee02447ba7750ee97f23831c")
	domainSeparator := common.HexToHash("0x82b5c47bb09eca2c93143f36f8fde6567050d39f3611535aab530d4f15fa5d0f")

	relayErr := solverOp.checkSignature(domainSeparator)
	if relayErr != nil {
		t.Errorf("SolverOperation.checkSignature() error = %v", relayErr)
	}
}
