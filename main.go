package main

import (
	"log"
	"os"

	"github.com/FastLane-Labs/atlas-operations-relay/core"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	ethRpcUrl := os.Getenv("ETH_RPC_URL")

	ethClient, err := ethclient.Dial(ethRpcUrl)
	if err != nil {
		log.Fatalf("could not initialize eth client: %s", err)
	}

	relay := core.NewRelay(ethClient)
	relay.Run()
}
