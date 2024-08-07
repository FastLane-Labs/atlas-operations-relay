package core

import (
	"math/big"

	"github.com/FastLane-Labs/atlas-operations-relay/contract/dAppControl"
	"github.com/FastLane-Labs/atlas-operations-relay/log"
	"github.com/FastLane-Labs/atlas-operations-relay/operation"
	"github.com/FastLane-Labs/atlas-operations-relay/relayerror"
	"github.com/ethereum/go-ethereum/common"
)

func (r *Relay) solverGasLimit(dAppControlAddress common.Address) (uint32, *relayerror.Error) {
	dAppControlContract, err := dAppControl.NewDAppControl(dAppControlAddress, r.ethClient)
	if err != nil {
		log.Error("solverGasLimit: failed to create dAppControl contract", "err", err)
		return 0, relayerror.ErrServerInternal
	}

	solverGasLimit, err := dAppControlContract.GetSolverGasLimit(nil)
	if err != nil {
		log.Error("solverGasLimit: failed to get solver gas limit", "err", err)
		return 0, relayerror.ErrServerInternal
	}

	return solverGasLimit, nil
}

func (r *Relay) balanceOfBonded(account common.Address) (*big.Int, *relayerror.Error) {
	balance, err := r.atlETHContract.BalanceOfBonded(nil, account)
	if err != nil {
		log.Error("balanceOfBonded: failed to get bonded balance", "err", err)
		return nil, ErrCantGetBondedBalance.AddError(err)
	}
	return balance, nil
}

func (r *Relay) reputationScore(account common.Address) int {
	accessData, err := r.storageContract.AccessData(nil, account)
	if err != nil {
		log.Error("reputationScore: failed to get access data", "err", err)
		return 0
	}
	wins := accessData.AuctionWins.Int64()
	fails := accessData.AuctionFails.Int64()
	total := wins + fails
	if total == 0 {
		return 0
	}
	return int((wins * 100) / total)
}

func (r *Relay) getDAppSignatories(dAppControl common.Address) ([]common.Address, *relayerror.Error) {
	signatories, err := r.atlasVerificationContract.DAppSignatories(nil, dAppControl)
	if err != nil {
		log.Error("getDAppSignatories: failed to get dApp signatories", "err", err)
		return []common.Address{}, ErrCantGetDAppSignatories.AddError(err)
	}

	return signatories, nil
}

func (r *Relay) getDAppConfig(dAppControlAddress common.Address, userOp *operation.UserOperation) (*dAppControl.DAppConfig, *relayerror.Error) {
	dAppControlContract, err := dAppControl.NewDAppControl(dAppControlAddress, r.ethClient)
	if err != nil {
		log.Info("failed to instantiate dapp control contract", "err", err)
		return nil, relayerror.ErrServerInternal
	}

	dAppConfig, err := dAppControlContract.GetDAppConfig(nil, dAppControl.UserOperation(*userOp))
	if err != nil {
		log.Info("failed to get dapp config", "err", err)
		return nil, relayerror.ErrServerInternal
	}

	return &dAppConfig, nil
}

func (r *Relay) getUserNextNonce(account common.Address, requireSequencedNonces bool) (*big.Int, *relayerror.Error) {
	nonce, err := r.atlasVerificationContract.GetUserNextNonce(nil, account, requireSequencedNonces)
	if err != nil {
		log.Error("failed to get next nonce", "err", err)
		return nil, relayerror.ErrServerInternal
	}

	return nonce, nil
}

func (r *Relay) getDappNextNonce(account common.Address) (*big.Int, *relayerror.Error) {
	nonce, err := r.atlasVerificationContract.GetDAppNextNonce(nil, account)
	if err != nil {
		log.Error("failed to get next nonce", "err", err)
		return nil, relayerror.ErrServerInternal
	}

	return nonce, nil
}
