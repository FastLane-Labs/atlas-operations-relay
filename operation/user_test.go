package operation

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
)

func generateUserOperation() *UserOperation {
	return &UserOperation{
		From:         common.HexToAddress("0x1"),
		To:           common.HexToAddress("0x2"),
		Deadline:     big.NewInt(100),
		Gas:          big.NewInt(200),
		Nonce:        big.NewInt(300),
		MaxFeePerGas: big.NewInt(400),
		Value:        big.NewInt(500),
		Dapp:         common.HexToAddress("0x3"),
		Control:      common.HexToAddress("0x4"),
		SessionKey:   common.HexToAddress("0x5"),
		Data:         []byte("data"),
		Signature:    []byte("signature"),
	}
}

func TestUserOperationHash(t *testing.T) {
	t.Parallel()

	userOp := generateUserOperation()
	want := common.HexToHash("0x0e658352ea19af267ebd61688d3056cb0e432713bf4adde58dac46419eb25485")

	result, err := userOp.Hash(false)
	if err != nil {
		t.Errorf("UserOperation.Hash() error = %v", err)
	}

	if result != want {
		t.Errorf("UserOperation.Hash() = %v, want %v", result, want)
	}
}

func TestUserOperationAltHash(t *testing.T) {
	t.Parallel()

	userOp := generateUserOperation()
	want := common.HexToHash("0xb715f6dca6a71f65387c6837476bf0ed9f9169a2b93fefed67d2fa82a8e1e1dc")

	result, err := userOp.Hash(true)
	if err != nil {
		t.Errorf("UserOperation.Hash() error = %v", err)
	}

	if result != want {
		t.Errorf("UserOperation.Hash() = %v, want %v", result, want)
	}
}

func TestUserOperationAbiEncode(t *testing.T) {
	t.Parallel()

	userOp := generateUserOperation()
	want := common.FromHex("0x00000000000000000000000000000000000000000000000000000000000000200000000000000000000000000000000000000000000000000000000000000001000000000000000000000000000000000000000000000000000000000000000200000000000000000000000000000000000000000000000000000000000001f400000000000000000000000000000000000000000000000000000000000000c80000000000000000000000000000000000000000000000000000000000000190000000000000000000000000000000000000000000000000000000000000012c0000000000000000000000000000000000000000000000000000000000000064000000000000000000000000000000000000000000000000000000000000000300000000000000000000000000000000000000000000000000000000000000040000000000000000000000000000000000000000000000000000000000000005000000000000000000000000000000000000000000000000000000000000018000000000000000000000000000000000000000000000000000000000000001c00000000000000000000000000000000000000000000000000000000000000004646174610000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000097369676e61747572650000000000000000000000000000000000000000000000")

	result, err := userOp.AbiEncode()
	if err != nil {
		t.Errorf("UserOperation.abiEncode() error = %v", err)
	}

	if string(result) != string(want) {
		t.Errorf("UserOperation.abiEncode() = %v, want %v", result, want)
	}
}

func TestUserOperationProofHash(t *testing.T) {
	t.Parallel()

	userOp := generateUserOperation()
	want := common.HexToHash("0x53db56ff45b29fe0f2f24afb71f70a5f0cc09089d6ef3cbc6dc4c3479492a92f")

	result, err := userOp.ProofHash()
	if err != nil {
		t.Errorf("UserOperation.proofHash() error = %v", err)
	}

	if result != want {
		t.Errorf("UserOperation.proofHash() = %v, want %v", result, want)
	}
}

func TestUserOperationCheckSignature(t *testing.T) {
	t.Parallel()

	userOp := generateUserOperation()
	userOp.From = common.HexToAddress("0xeBfc39Ea4C2540a49d22700C7f08Ed517C54049B")
	userOp.Signature = common.FromHex("0xc0a0b163768a2369d6891e0dc7ad20c51a9618b4744474de36f1ca9cda6963e777e25c89a3ad053758ad95c859cd945fb8e5897cf9653b2f4e222f5eeafd314a1b")
	domainSeparator := common.HexToHash("0x82b5c47bb09eca2c93143f36f8fde6567050d39f3611535aab530d4f15fa5d0f")

	relayErr := userOp.checkSignature(domainSeparator)
	if relayErr != nil {
		t.Errorf("UserOperation.checkSignature() error = %v", relayErr)
	}
}
