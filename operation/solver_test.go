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

func TestSolverOperationHash(t *testing.T) {
	t.Parallel()

	solverOp := generateSolverOperation()
	want := common.HexToHash("0x36ff5301f530e9175cca32acbac7bc6798d5be0c04372ba3f05149b2455c5405")

	result, err := solverOp.Hash(&Eip712DomainTest)
	if err != nil {
		t.Errorf("SolverOperation.Hash() error = %v", err)
	}

	if result != want {
		t.Errorf("SolverOperation.Hash() = %v, want %v", result, want)
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

func TestSolverOperationCheckSignature(t *testing.T) {
	t.Parallel()

	solverOp := generateSolverOperation()
	solverOp.From = common.HexToAddress("0xB764B6545d283C0E547952763F8a843394295da1")
	solverOp.Signature = common.FromHex("0x8e8d89974eb665dea669d922ad26d055b835d426fb60885a40d439ea213deb8204ce569ad60caa3fe8892656fa189a28ab2c71b02135b9b9d746d6ac12b87b6c1c")

	relayErr := solverOp.checkSignature(&Eip712DomainTest)
	if relayErr != nil {
		t.Errorf("SolverOperation.checkSignature() error = %v", relayErr)
	}
}
