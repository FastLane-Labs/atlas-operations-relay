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
		Nonce:         big.NewInt(100),
		Deadline:      big.NewInt(200),
		Control:       common.HexToAddress("0x3"),
		Bundler:       common.HexToAddress("0x4"),
		UserOpHash:    common.HexToHash("0x9999999999999999999999999999999999999999999999999999999999999999"),
		CallChainHash: common.HexToHash("0x8888888888888888888888888888888888888888888888888888888888888888"),
		Signature:     []byte("signature"),
	}
}

func TestDAppOperationHash(t *testing.T) {
	t.Parallel()

	dAppOp := generateDAppOperation()
	want := common.HexToHash("0x19f8a2e775a072f45728d1d87bcfff85ca5fecdf674f261af2a5c8c70b37a453")

	result, err := dAppOp.Hash(&Eip712DomainTest)
	if err != nil {
		t.Errorf("DAppOperation.Hash() error = %v", err)
	}

	if result != want {
		t.Errorf("DAppOperation.Hash() = %v, want %v", result, want)
	}
}

func TestDAppOperationCheckSignature(t *testing.T) {
	t.Parallel()

	dAppOp := generateDAppOperation()
	dAppOp.From = common.HexToAddress("0xB764B6545d283C0E547952763F8a843394295da1")
	dAppOp.Signature = common.FromHex("0x741bd1cc70e34a39d763ae23d0d94c6a9156b10ba9a4cead3e847d4f15ad6edf4a7a60b875f1cb1795358b7a395b422659b7336f2f3a90453f8c2a16369e69d81c")

	relayErr := dAppOp.checkSignature(&Eip712DomainTest)
	if relayErr != nil {
		t.Errorf("DAppOperation.checkSignature() error = %v", relayErr)
	}
}
