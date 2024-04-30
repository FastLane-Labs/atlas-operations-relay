package tests

import (
	"github.com/FastLane-Labs/atlas-operations-relay/contract/atlas"
	"github.com/ethereum/go-ethereum/common"
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
