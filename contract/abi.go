package contract

import (
	"github.com/FastLane-Labs/atlas-operations-relay/contract/dAppControl"
	"github.com/FastLane-Labs/atlas-operations-relay/contract/simulator"
)

var (
	DappControlAbi, _ = dAppControl.DAppControlMetaData.GetAbi()
	SimulatorAbi, _   = simulator.SimulatorMetaData.GetAbi()
)
