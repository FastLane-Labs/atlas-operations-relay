package crypto

import (
	"github.com/FastLane-Labs/atlas-operations-relay/log"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

func GetSigner(domainSeparator common.Hash, structHash common.Hash, signature []byte) (common.Address, error) {
	hash := crypto.Keccak256([]byte("\x19\x01"), domainSeparator.Bytes(), structHash.Bytes())

	if signature[crypto.RecoveryIDOffset] == 27 || signature[crypto.RecoveryIDOffset] == 28 {
		signature[crypto.RecoveryIDOffset] -= 27 // Transform yellow paper V from 27/28 to 0/1
	}

	recovered, err := crypto.SigToPub(hash, signature)
	if err != nil {
		log.Info("failed to recover public key", "err", err)
		return common.Address{}, err
	}

	return crypto.PubkeyToAddress(*recovered), nil
}
