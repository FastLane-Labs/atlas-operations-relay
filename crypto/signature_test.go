package crypto

import (
	"testing"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

func TestUserOperationHash(t *testing.T) {
	t.Parallel()

	domainSeparator := common.HexToHash("0x82b5c47bb09eca2c93143f36f8fde6567050d39f3611535aab530d4f15fa5d0f")
	randomData := common.HexToHash("0x1234567890123456789012345678901234567890123456789012345678901234")
	payload := crypto.Keccak256Hash([]byte("\x19\x01"), domainSeparator.Bytes(), randomData.Bytes())

	privateKey, err := crypto.GenerateKey()
	if err != nil {
		t.Errorf("crypto.GenerateKey() error = %v", err)
	}

	signature, err := crypto.Sign(payload.Bytes(), privateKey)
	if err != nil {
		t.Errorf("crypto.Sign() error = %v", err)
	}

	want := crypto.PubkeyToAddress(privateKey.PublicKey)

	signer, err := GetSigner(domainSeparator, randomData, signature)
	if err != nil {
		t.Errorf("GetSigner() error = %v", err)
	}

	if signer != want {
		t.Errorf("GetSigner() = %v, want %v", signer, want)
	}
}

func TestBundlerAuthentication(t *testing.T) {
	t.Parallel()

	privateKey, err := crypto.GenerateKey()
	if err != nil {
		t.Errorf("crypto.GenerateKey() error = %v", err)
	}

	content := "0x0000000000000000000000000000000000000000:0"

	signature, err := crypto.Sign(accounts.TextHash([]byte(content)), privateKey)
	if err != nil {
		t.Errorf("crypto.Sign() error = %v", err)
	}

	want := crypto.PubkeyToAddress(privateKey.PublicKey)
	signer, err := RecoverEthereumSigner(content, signature)
	if err != nil {
		t.Errorf("RecoverEthereumSigner() error = %v", err)
	}

	if signer != want {
		t.Errorf("RecoverEthereumSigner() = %v, want %v", signer, want)
	}
}
