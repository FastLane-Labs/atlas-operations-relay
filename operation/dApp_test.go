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
	want := common.HexToHash("0xd71b1989258522f5316e5fabd7b6e171897f96c5807f8261da600d1aac7da71d")

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

	dAppOp.From = common.HexToAddress("0xb8BE63EC949a4aBD643404Ad1c78d3cca29D671A")
	domainSeparator := common.HexToHash("0x82b5c47bb09eca2c93143f36f8fde6567050d39f3611535aab530d4f15fa5d0f")

	dAppOp.Signature = common.FromHex("0f4c1dc7452c3e85b138bad755fa531dc034b795c236594db2f10ff49be1ee1a28db2f9064117b352c67272afbf5a513feadef9d9b7c428f55e96067920386861b")

	relayErr := dAppOp.checkSignature(domainSeparator)
	if relayErr != nil {
		t.Errorf("DAppOperation.checkSignature() error = %v", relayErr)
	}
}
