package contract

import (
	"github.com/FastLane-Labs/atlas-operations-relay/contract/atlas"
	"github.com/FastLane-Labs/atlas-operations-relay/contract/dAppControl"
	"github.com/FastLane-Labs/atlas-operations-relay/contract/simpleRfqSolver"
	"github.com/FastLane-Labs/atlas-operations-relay/contract/simulator"
)

var (
	AtlasAbi, _           = atlas.AtlasMetaData.GetAbi()
	DappControlAbi, _     = dAppControl.DAppControlMetaData.GetAbi()
	SimulatorAbi, _       = simulator.SimulatorMetaData.GetAbi()
	SimpleRfqSolverAbi, _ = simpleRfqSolver.SimpleRfqSolverMetaData.GetAbi()
)
