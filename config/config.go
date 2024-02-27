package config

import (
	"encoding/json"
	"os"

	"github.com/ethereum/go-ethereum/common"
)

type Config struct {
	Network *struct {
		RpcUrl string `json:"rpc_url"`
	} `json:"network"`

	Contracts *struct {
		Atlas             common.Address `json:"atlas"`
		AtlasVerification common.Address `json:"atlasVerification"`
		Simulator         common.Address `json:"simulator"`
	} `json:"contracts"`
}

func Load() *Config {
	var config Config

	file, err := os.Open("config.json")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&config); err != nil {
		panic(err)
	}

	if envRpcUrl, set := os.LookupEnv("ETH_RPC_URL"); set && len(envRpcUrl) > 0 {
		config.Network.RpcUrl = envRpcUrl
	}

	return &config
}
