package main

import (
	"os"

	"github.com/FastLane-Labs/atlas-operations-relay/core"
	"github.com/FastLane-Labs/atlas-operations-relay/log"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	log.InitLogger()

	ethRpcUrl := os.Getenv("ETH_RPC_URL")

	ethClient, err := ethclient.Dial(ethRpcUrl)
	if err != nil {
		log.Error("failed to connect to the Ethereum client", "err", err)
		return
	}

	relay := core.NewRelay(ethClient)
	relay.Run()
}
