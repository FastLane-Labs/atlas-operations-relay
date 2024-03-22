package operation

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
)

func generateDAppOperation() *DAppOperation {
	return &DAppOperation{
		From:          common.HexToAddress("0x1"),
		To:            common.HexToAddress("0x2"),
		Value:         big.NewInt(100),
		Gas:           big.NewInt(200),
		Nonce:         big.NewInt(300),
		Deadline:      big.NewInt(400),
		Control:       common.HexToAddress("0x3"),
		Bundler:       common.HexToAddress("0x4"),
		UserOpHash:    common.HexToHash("0x9999999999999999999999999999999999999999999999999999999999999999"),
		CallChainHash: common.HexToHash("0x8888888888888888888888888888888888888888888888888888888888888888"),
		Signature:     []byte("signature"),
	}
}

func TestDAppOperationProofHash(t *testing.T) {
	t.Parallel()

	dAppOp := generateDAppOperation()
	want := common.HexToHash("0x3b041ae0e931009f49b824007dd38b60c636ee64e1a5c33774ea8328d1c119a8")

	result, err := dAppOp.ProofHash()
	if err != nil {
		t.Errorf("DAppOperation.proofHash() error = %v", err)
	}

	if result != want {
		t.Errorf("DAppOperation.proofHash() = %v, want %v", result, want)
	}
}

func TestDAppOperationCheckSignature(t *testing.T) {
	t.Parallel()

	dAppOp := generateDAppOperation()
	dAppOp.From = common.HexToAddress("0x4af08B2fA648F2f2B0A6805f37516050353d83dc")
	dAppOp.Signature = common.FromHex("0x0ee6e7b9c532f1091efcac2b08fe7e631c22ba718bc4bf26bda5e6316461e9633c3f4fa50c8be1feec30a184b7844718ac957a10fe33fa8323e81eb7d9dfcf5e1c")
	domainSeparator := common.HexToHash("0x82b5c47bb09eca2c93143f36f8fde6567050d39f3611535aab530d4f15fa5d0f")

	relayErr := dAppOp.checkSignature(domainSeparator)
	if relayErr != nil {
		t.Errorf("DAppOperation.checkSignature() error = %v", relayErr)
	}
}
