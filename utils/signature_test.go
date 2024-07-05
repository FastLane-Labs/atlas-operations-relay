package utils

import (
	"testing"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/crypto"
)

func TestRecoverEthereumSigner(t *testing.T) {
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
