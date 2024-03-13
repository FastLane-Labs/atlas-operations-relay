package tests

import (
	"crypto/ecdsa"
	"fmt"

	"github.com/FastLane-Labs/atlas-operations-relay/contract/atlas"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

func executionEnvironment(user common.Address, dAppControl common.Address) common.Address {
	atlasContract, err := atlas.NewAtlas(conf.Contracts.Atlas, ethClient)
	if err != nil {
		panic(err)
	}

	executionEnvironment, err := atlasContract.GetExecutionEnvironment(nil, user, dAppControl)
	if err != nil {
		panic(err)
	}

	return executionEnvironment.ExecutionEnvironment
}

func signMessage(data []byte, privKey *ecdsa.PrivateKey) ([]byte, error) {
	prefix := fmt.Sprintf("\x19Ethereum Signed Message:\n%d", len(data))
	prefixedData := []byte(prefix + string(data))

	hash := crypto.Keccak256Hash(prefixedData)

	signature, err := crypto.Sign(hash.Bytes(), privKey)
	if err != nil {
		return nil, err
	}

	// According to the yellow paper, the V value in signature (27 or 28) is expected
	// by most libraries, including web3. However, Go Ethereum `Sign` method produces
	// 0 or 1 for V. So we adjust V back to 27 or 28 to ensure compatibility.
	if signature[64] < 27 {
		signature[64] += 27
	}

	return signature, nil
}

func signEip712(domainSeparator common.Hash, proofHash common.Hash, pk *ecdsa.PrivateKey) []byte {
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
