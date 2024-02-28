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
		MaxFeePerGas:  big.NewInt(300),
		Nonce:         big.NewInt(400),
		Deadline:      big.NewInt(500),
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
	want := common.HexToHash("0xf9b9b34f358309d76ec72f2d9be0af79c824c36999896c7a10c41949f5ff070a")

	result, err := dAppOp.proofHash()
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
	dAppOp.Signature = common.FromHex("0xb05222fcd6215c9f7b3880f74a813146bb4e1ef7476b8c4e8000d5ece49d6055098b79f9900703d0bf1bed633ac79f9a4197bddf24a152fd4a40b43d7a3b29461c")
	domainSeparator := common.HexToHash("0x82b5c47bb09eca2c93143f36f8fde6567050d39f3611535aab530d4f15fa5d0f")

	relayErr := dAppOp.checkSignature(domainSeparator)
	if relayErr != nil {
		t.Errorf("DAppOperation.checkSignature() error = %v", relayErr)
	}
}
