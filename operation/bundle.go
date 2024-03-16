package operation

import (
	"math/big"

	"github.com/FastLane-Labs/atlas-operations-relay/contract"
	"github.com/FastLane-Labs/atlas-operations-relay/contract/dAppControl"
	"github.com/FastLane-Labs/atlas-operations-relay/log"
	"github.com/FastLane-Labs/atlas-operations-relay/relayerror"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

// External representation of a bundle of operations,
// the relay receives and broadcasts bundles in this format
type BundleOperationsRaw struct {
	UserOperation    *UserOperationRaw     `json:"userOperation"`
	SolverOperations []*SolverOperationRaw `json:"solverOperations"`
	DAppOperation    *DAppOperationRaw     `json:"dAppOperation"`
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
	SolverOperations []*SolverOperation
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

func (b *BundleOperations) Validate(ethClient *ethclient.Client, userOpHash common.Hash, atlas common.Address, atlasDomainSeparator common.Hash, userOpGasLimit *big.Int, dAppOpGasLimit *big.Int) *relayerror.Error {
	// Re-validate user operation
	if relayErr := b.UserOperation.Validate(ethClient, atlas, atlasDomainSeparator, userOpGasLimit); relayErr != nil {
		return relayErr
	}

	dAppControlContract, err := dAppControl.NewDAppControl(b.DAppOperation.Control, ethClient)
	if err != nil {
		log.Info("failed to instantiate dapp control contract", "err", err)
		return relayerror.ErrServerInternal
	}

	dAppConfig, err := dAppControlContract.GetDAppConfig(nil, dAppControl.UserOperation(*b.UserOperation))
	if err != nil {
		log.Info("failed to get dapp config", "err", err)
		return relayerror.ErrServerInternal
	}

	callChainHash, err := b.CallChainHash(dAppConfig.CallConfig, dAppConfig.To)
	if err != nil {
		log.Info("failed to compute call chain hash", "err", err)
		return relayerror.ErrServerInternal
	}

	if relayErr := b.DAppOperation.Validate(userOpHash, b.UserOperation, callChainHash, atlas, atlasDomainSeparator, dAppOpGasLimit); relayErr != nil {
		return relayErr
	}

	return nil
}

func (b *BundleOperations) CallChainHash(callConfig uint32, dAppControl common.Address) (common.Hash, error) {
	counter := big.NewInt(0)
	var callSequenceHash common.Hash

	if callConfig&4 != 0 {
		// Require preOps
		preOpsEncoded, err := contract.DappControlAbi.Pack("preOpsCall", b.UserOperation)
		if err != nil {
			return common.Hash{}, err
		}

		callSequenceHash = crypto.Keccak256Hash(
			callSequenceHash.Bytes(),
			dAppControl.Bytes(),
			preOpsEncoded,
			math.U256Bytes(counter),
		)

		counter.Add(counter, common.Big1)
	}

	userOpAbiEncoded, err := b.UserOperation.AbiEncode()
	if err != nil {
		return common.Hash{}, err
	}

	callSequenceHash = crypto.Keccak256Hash(
		callSequenceHash.Bytes(),
		userOpAbiEncoded,
		math.U256Bytes(counter),
	)

	counter.Add(counter, common.Big1)

	for _, solverOp := range b.SolverOperations {
		solverOpAbiEncoded, err := solverOp.AbiEncode()
		if err != nil {
			return common.Hash{}, err
		}

		callSequenceHash = crypto.Keccak256Hash(
			callSequenceHash.Bytes(),
			solverOpAbiEncoded,
			math.U256Bytes(counter),
		)

		counter.Add(counter, common.Big1)
	}

	return callSequenceHash, nil
}
