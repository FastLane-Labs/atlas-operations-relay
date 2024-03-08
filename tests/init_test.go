package tests

import (
	"os"
	"testing"

	"github.com/FastLane-Labs/atlas-operations-relay/config"
	"github.com/FastLane-Labs/atlas-operations-relay/core"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	ethClient *ethclient.Client

	conf = &config.Config{
		Network: config.Network{
			RpcUrl: "https://rpc.sepolia.org/",
		},
		Contracts: config.Contracts{
			Atlas:             common.HexToAddress("0x090A1134ebF3066f87dB165a910E832a87C2Ffb5"),
			AtlasVerification: common.HexToAddress("0x15261B0437b0e54d6f0Dd308148a359Bb0191A81"),
			Simulator:         common.HexToAddress("0x8F4C1158fA69f973973b6ba7A1b3F58EE9204d1A"),
		},
	}
)

func TestMain(m *testing.M) {
	var err error

	ethClient, err = ethclient.Dial(conf.Network.RpcUrl)
	if err != nil {
		panic(err)
	}

	serverReadyChan := make(chan struct{})
	// Start the relay
	go core.StartRelay(conf, serverReadyChan)
	// Wait for the server to be ready
	<-serverReadyChan
	os.Exit(m.Run())
}
