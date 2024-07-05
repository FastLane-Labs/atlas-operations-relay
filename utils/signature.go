package utils

import (
	"crypto/ecdsa"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

func SignMessage(data []byte, pk *ecdsa.PrivateKey) ([]byte, error) {
	sig, err := crypto.Sign(data, pk)
	if err != nil {
		return nil, err
	}
	if sig[crypto.RecoveryIDOffset] == 0 || sig[crypto.RecoveryIDOffset] == 1 {
		sig[crypto.RecoveryIDOffset] += 27
	}
	return sig, nil
}

func SignEthereumMessage(data []byte, pk *ecdsa.PrivateKey) ([]byte, error) {
	return SignMessage(accounts.TextHash(data), pk)
}

func RecoverSigner(messageHash []byte, signature []byte) (common.Address, error) {
	sig := make([]byte, len(signature))
	copy(sig, signature)
	if sig[crypto.RecoveryIDOffset] == 27 || sig[crypto.RecoveryIDOffset] == 28 {
		sig[crypto.RecoveryIDOffset] -= 27
	}
	pubKey, err := crypto.SigToPub(messageHash, sig)
	if err != nil {
		return common.Address{}, err
	}
	return crypto.PubkeyToAddress(*pubKey), nil
}

func RecoverEthereumSigner(message string, signature []byte) (common.Address, error) {
	messageHash := accounts.TextHash([]byte(message))
	return RecoverSigner(messageHash, signature)
}
