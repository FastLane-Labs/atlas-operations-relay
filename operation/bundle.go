package operation

import (
	"math/big"

	"github.com/FastLane-Labs/atlas-operations-relay/contract/dAppControl"
	"github.com/FastLane-Labs/atlas-operations-relay/log"
	"github.com/FastLane-Labs/atlas-operations-relay/relayerror"
	"github.com/FastLane-Labs/atlas-operations-relay/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/signer/core/apitypes"
)

// External representation of a bundle of operations,
// the relay receives and broadcasts bundles in this format
type BundleOperationsRaw struct {
	UserOperation    *UserOperationRaw     `json:"userOperation" validate:"required"`
	SolverOperations []*SolverOperationRaw `json:"solverOperations" validate:"required"`
	DAppOperation    *DAppOperationRaw     `json:"dAppOperation" validate:"required"`
}

func (bor *BundleOperationsRaw) IsZero() bool {
	return bor.UserOperation == nil && len(bor.SolverOperations) == 0 && bor.DAppOperation == nil
}

func (args *BundleOperationsRaw) Decode() *BundleOperations {
	var solverOps []*SolverOperation
	for _, solverOpArgs := range args.SolverOperations {
		solverOps = append(solverOps, solverOpArgs.Decode())
	}
	return &BundleOperations{
		UserOperation:    args.UserOperation.Decode(),
		SolverOperations: solverOps,
		DAppOperation:    args.DAppOperation.Decode(),
	}
}

// Internal representation of a bundle of operations
type BundleOperations struct {
	UserOperation    *UserOperation
	SolverOperations SolverOperations
	DAppOperation    *DAppOperation
}

func (b *BundleOperations) EncodeToRaw() *BundleOperationsRaw {
	var solverOpsRaw []*SolverOperationRaw
	for _, solverOp := range b.SolverOperations {
		solverOpsRaw = append(solverOpsRaw, solverOp.EncodeToRaw())
	}

	return &BundleOperationsRaw{
		UserOperation:    b.UserOperation.EncodeToRaw(),
		SolverOperations: solverOpsRaw,
		DAppOperation:    b.DAppOperation.EncodeToRaw(),
	}
}

func (b *BundleOperations) Validate(ethClient *ethclient.Client, userOpHash common.Hash, atlas common.Address, eip712Domain *apitypes.TypedDataDomain, userOpGasLimit *big.Int, dAppConfig *dAppControl.DAppConfig) *relayerror.Error {
	// Re-validate user operation
	if relayErr := b.UserOperation.Validate(ethClient, atlas, eip712Domain, userOpGasLimit); relayErr != nil {
		return relayErr
	}

	var (
		callChainHash common.Hash
		err           error
	)

	if utils.FlagVerifyCallChainHash(dAppConfig.CallConfig) {
		callChainHash, err = b.CallChainHash()
		if err != nil {
			log.Info("failed to compute call chain hash", "err", err)
			return relayerror.ErrServerInternal
		}
	}

	if relayErr := b.DAppOperation.Validate(userOpHash, b.UserOperation, callChainHash, atlas, eip712Domain); relayErr != nil {
		return relayErr
	}

	return nil
}

func (b *BundleOperations) CallChainHash() (common.Hash, error) {
	userOpAbiEncoded, err := b.UserOperation.AbiEncode()
	if err != nil {
		return common.Hash{}, err
	}

	solverOpsAbiEncoded, err := b.SolverOperations.AbiEncode()
	if err != nil {
		return common.Hash{}, err
	}

	return crypto.Keccak256Hash(
		userOpAbiEncoded,
		solverOpsAbiEncoded,
	), nil
}
