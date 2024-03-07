package core

import (
	"github.com/FastLane-Labs/atlas-operations-relay/config"
	"github.com/FastLane-Labs/atlas-operations-relay/log"
	"github.com/ethereum/go-ethereum/ethclient"
)

func StartRelay(conf *config.Config, serverReadyChan chan struct{}) {
	log.InitLogger()

	if conf == nil {
		conf = config.Load()
	}

	ethClient, err := ethclient.Dial(conf.Network.RpcUrl)
	if err != nil {
		log.Error("failed to connect to the Ethereum client", "err", err)
		return
	}

	NewRelay(ethClient, conf).Run(serverReadyChan)
}
