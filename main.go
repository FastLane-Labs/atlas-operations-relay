package main

import (
	"github.com/FastLane-Labs/atlas-operations-relay/config"
	"github.com/FastLane-Labs/atlas-operations-relay/core"
	"github.com/FastLane-Labs/atlas-operations-relay/log"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	log.InitLogger()
	config := config.Load()

	ethClient, err := ethclient.Dial(config.Network.RpcUrl)
	if err != nil {
		log.Error("failed to connect to the Ethereum client", "err", err)
		return
	}

	relay := core.NewRelay(ethClient, config)
	relay.Run()
}
