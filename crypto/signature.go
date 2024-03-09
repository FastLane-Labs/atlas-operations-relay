package crypto

import (
	"crypto/ecdsa"

	"github.com/FastLane-Labs/atlas-operations-relay/log"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

func RecoverEip712Signer(domainSeparator common.Hash, structHash common.Hash, signature []byte) (common.Address, error) {
	messageHash := crypto.Keccak256([]byte("\x19\x01"), domainSeparator.Bytes(), structHash.Bytes())
	return recoverSigner(messageHash, signature)
}

func RecoverEthereumSigner(message string, signature []byte) (common.Address, error) {
	messageHash := accounts.TextHash([]byte(message))
	return recoverSigner(messageHash, signature)
}

func recoverSigner(messageHash []byte, signature []byte) (common.Address, error) {
	changedSignature := false
	if signature[crypto.RecoveryIDOffset] == 27 || signature[crypto.RecoveryIDOffset] == 28 {
		signature[crypto.RecoveryIDOffset] -= 27 // Transform yellow paper V from 27/28 to 0/1
		changedSignature = true
	}

	pubKey, err := crypto.SigToPub(messageHash, signature)
	if err != nil {
		log.Info("failed to recover public key", "err", err)
		return common.Address{}, err
	}

	//reset the v of the signature
	if changedSignature {
		signature[crypto.RecoveryIDOffset] += 27
	}
	return crypto.PubkeyToAddress(*pubKey), nil
}

func SignEip712(domainSeparator common.Hash, proofHash common.Hash, pk *ecdsa.PrivateKey) []byte {
	payload := crypto.Keccak256Hash([]byte("\x19\x01"), domainSeparator.Bytes(), proofHash.Bytes())
	signature, err := crypto.Sign(payload.Bytes(), pk)
	if err != nil {
		panic(err)
	}
	if signature[64] < 2 {
		signature[64] += 27
	}

	return signature
}
