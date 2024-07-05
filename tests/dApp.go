package tests

import (
	"math/big"

	"github.com/FastLane-Labs/atlas-operations-relay/contract/dAppControl"
	"github.com/FastLane-Labs/atlas-operations-relay/operation"
	"github.com/FastLane-Labs/atlas-operations-relay/utils"
)

func newDappOperation(userOp *operation.UserOperation, solverOps []*operation.SolverOperation) *operation.DAppOperation {
	userOpHash, _ := userOp.Hash(false, &conf.Relay.Eip712.Domain)

	dAppControlContract, err := dAppControl.NewDAppControl(userOp.Control, ethClient)
	if err != nil {
		panic(err)
	}

	dAppConfig, err := dAppControlContract.GetDAppConfig(nil, dAppControl.UserOperation(*userOp))
	if err != nil {
		panic(err)
	}

	callChainHash, err := (&operation.BundleOperations{UserOperation: userOp, SolverOperations: solverOps}).CallChainHash(dAppConfig.CallConfig, dAppConfig.To)
	if err != nil {
		panic(err)
	}

	dAppOp := &operation.DAppOperation{
		From:          userOp.From,
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

	dAppOp.Signature, _ = utils.SignMessage(dAppOpHash.Bytes(), userPk)

	return dAppOp
}
