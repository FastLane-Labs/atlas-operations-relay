package utils

import (
	"crypto/ecdsa"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

func SignEthereumMessage(data []byte, pk *ecdsa.PrivateKey) ([]byte, error) {
	sig, err := crypto.Sign(accounts.TextHash(data), pk)
	if err != nil {
		return nil, err
	}
	if sig[crypto.RecoveryIDOffset] == 0 || sig[crypto.RecoveryIDOffset] == 1 {
		sig[crypto.RecoveryIDOffset] += 27
	}
	return sig, nil
}

func SignEip712Message(domainSeparator common.Hash, proofHash common.Hash, pk *ecdsa.PrivateKey) ([]byte, error) {
	payload := crypto.Keccak256Hash([]byte("\x19\x01"), domainSeparator.Bytes(), proofHash.Bytes())
	sig, err := crypto.Sign(payload.Bytes(), pk)
	if err != nil {
		return nil, err
	}
	if sig[crypto.RecoveryIDOffset] == 0 || sig[crypto.RecoveryIDOffset] == 1 {
		sig[crypto.RecoveryIDOffset] += 27
	}
	return sig, nil
}

func RecoverEthereumSigner(message string, signature []byte) (common.Address, error) {
	messageHash := accounts.TextHash([]byte(message))
	return recoverSigner(messageHash, signature)
}

func RecoverEip712Signer(domainSeparator common.Hash, structHash common.Hash, signature []byte) (common.Address, error) {
	messageHash := crypto.Keccak256([]byte("\x19\x01"), domainSeparator.Bytes(), structHash.Bytes())
	return recoverSigner(messageHash, signature)
}

func recoverSigner(messageHash []byte, signature []byte) (common.Address, error) {
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
