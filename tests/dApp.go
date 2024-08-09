package tests

import (
	"math/big"

	"github.com/FastLane-Labs/atlas-operations-relay/operation"
	"github.com/FastLane-Labs/atlas-operations-relay/utils"
)

func newDappOperation(userOp *operation.UserOperation, solverOps []*operation.SolverOperation) *operation.DAppOperation {
	userOpHash, _ := userOp.Hash(utils.FlagTrustedOpHash(userOp.CallConfig), &conf.Relay.Eip712.Domain)

	callChainHash, err := (&operation.BundleOperations{UserOperation: userOp, SolverOperations: solverOps}).CallChainHash()
	if err != nil {
		panic(err)
	}

	dAppOp := &operation.DAppOperation{
		From:          governanceEoa,
		To:            conf.Contracts.Atlas,
		Nonce:         big.NewInt(0),
		Deadline:      userOp.Deadline,
		Control:       userOp.Control,
		Bundler:       bundlerEoa,
		UserOpHash:    userOpHash,
		CallChainHash: callChainHash,
		Signature:     []byte(""),
	}

	dAppOpHash, relayErr := dAppOp.Hash(&conf.Relay.Eip712.Domain)
	if relayErr != nil {
		panic(relayErr)
	}

	dAppOp.Signature, _ = utils.SignMessage(dAppOpHash.Bytes(), governancePk)

	return dAppOp
}
