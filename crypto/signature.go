package crypto

import (
	"github.com/FastLane-Labs/atlas-operations-relay/log"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

func GetSigner(domainSeparator common.Hash, structHash common.Hash, signature []byte) (common.Address, error) {
	// add prefix \x19\x01 ?
	hash := crypto.Keccak256(domainSeparator.Bytes(), structHash.Bytes())
	compressedPublicKey, err := crypto.Ecrecover(hash, signature)
	if err != nil {
		log.Info("failed to recover public key", "err", err)
		return common.Address{}, err
	}

	publicKey, err := crypto.DecompressPubkey(compressedPublicKey)
	if err != nil {
		log.Info("failed to decompress public key", "err", err)
		return common.Address{}, err
	}

	return crypto.PubkeyToAddress(*publicKey), nil
}
