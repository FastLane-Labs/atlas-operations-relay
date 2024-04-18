package core

import (
	"math/big"

	"github.com/FastLane-Labs/atlas-operations-relay/relayerror"
	"github.com/ethereum/go-ethereum/common"
)

func (r *Relay) balanceOfBonded(account common.Address) (*big.Int, *relayerror.Error) {
	balance, err := r.atlETHContract.BalanceOfBonded(nil, account)
	if err != nil {
		return nil, ErrCantGetBondedBalance.AddError(err)
	}
	return balance, nil
}

func (r *Relay) reputationScore(account common.Address) int {
	accessData, err := r.storageContract.AccessData(nil, account)
	if err != nil {
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
	signatories, err := r.atlasVerificationContract.GetDAppSignatories(nil, dAppControl)
	if err != nil {
		return []common.Address{}, ErrCantGetDAppSignatories.AddError(err)
	}

	return signatories, nil
}
