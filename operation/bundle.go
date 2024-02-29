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

type BundleOperations struct {
	UserOperation    *UserOperation     `json:"userOperation"`
	SolverOperations []*SolverOperation `json:"solverOperations"`
	DAppOperation    *DAppOperation     `json:"dAppOperation"`
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

	callChainHash, err := b.callChainHash(dAppConfig.CallConfig, dAppConfig.To)
	if err != nil {
		log.Info("failed to compute call chain hash", "err", err)
		return relayerror.ErrServerInternal
	}

	if relayErr := b.DAppOperation.Validate(userOpHash, b.UserOperation, callChainHash, atlas, atlasDomainSeparator, dAppOpGasLimit); relayErr != nil {
		return relayErr
	}

	return nil
}

func (b *BundleOperations) callChainHash(callConfig uint32, dAppControl common.Address) (common.Hash, error) {
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

	userOpAbiEncoded, err := b.UserOperation.abiEncode()
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
		solverOpAbiEncoded, err := solverOp.abiEncode()
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
