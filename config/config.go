package config

import (
	"encoding/json"
	"flag"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/common"
)

type configJson struct {
	Network struct {
		RpcUrl string `json:"rpc_url"`
	} `json:"network"`

	Contracts struct {
		Atlas             string `json:"atlas"`
		AtlasVerification string `json:"atlasVerification"`
		Simulator         string `json:"simulator"`
	} `json:"contracts"`

	Relay struct {
		Gas struct {
			MaxPerUserOperation   uint64 `json:"max_per_user_operation"`
			MaxPerSolverOperation uint64 `json:"max_per_solver_operation"`
			MaxPerDAppOperation   uint64 `json:"max_per_dApp_operation"`
		} `json:"gas"`
	} `json:"relay"`
}

type Network struct {
	RpcUrl string `json:"rpc_url"`
}

type Contracts struct {
	Atlas             common.Address `json:"atlas"`
	AtlasVerification common.Address `json:"atlasVerification"`
	Simulator         common.Address `json:"simulator"`
}

type Gas struct {
	MaxPerUserOperation   *big.Int `json:"max_per_user_operation"`
	MaxPerSolverOperation *big.Int `json:"max_per_solver_operation"`
	MaxPerDAppOperation   *big.Int `json:"max_per_dApp_operation"`
}

type Relay struct {
	Gas Gas `json:"gas"`
}

type Config struct {
	Network   Network
	Contracts Contracts
	Relay     Relay
}

func Load() *Config {
	config := &Config{}

	config.parseConfigFile()
	config.parseFlags()
	config.parseEnv()
	config.validate()

	return config
}

func (c *Config) parseConfigFile() {
	var configJson configJson

	file, err := os.Open("config.json")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&configJson); err != nil {
		panic(err)
	}

	if len(configJson.Network.RpcUrl) > 0 {
		c.Network.RpcUrl = configJson.Network.RpcUrl
	}

	if len(configJson.Contracts.Atlas) > 0 {
		if !common.IsHexAddress(configJson.Contracts.Atlas) {
			panic("contracts.atlas is not a valid address")
		}
		c.Contracts.Atlas = common.HexToAddress(configJson.Contracts.Atlas)
	}

	if len(configJson.Contracts.AtlasVerification) > 0 {
		if !common.IsHexAddress(configJson.Contracts.AtlasVerification) {
			panic("contracts.atlasVerification is not a valid address")
		}
		c.Contracts.AtlasVerification = common.HexToAddress(configJson.Contracts.AtlasVerification)
	}

	if len(configJson.Contracts.Simulator) > 0 {
		if !common.IsHexAddress(configJson.Contracts.Simulator) {
			panic("contracts.simulator is not a valid address")
		}
		c.Contracts.Simulator = common.HexToAddress(configJson.Contracts.Simulator)
	}

	c.Relay.Gas.MaxPerUserOperation = new(big.Int).SetUint64(configJson.Relay.Gas.MaxPerUserOperation)
	c.Relay.Gas.MaxPerSolverOperation = new(big.Int).SetUint64(configJson.Relay.Gas.MaxPerSolverOperation)
	c.Relay.Gas.MaxPerDAppOperation = new(big.Int).SetUint64(configJson.Relay.Gas.MaxPerDAppOperation)
}

func (c *Config) parseFlags() {
	networkRpcUrlPtr := flag.String("network.rpc_url", "", "Ethereum RPC URL")
	contractsAtlasPtr := flag.String("contracts.atlas", "", "Atlas contract address")
	contractsAtlasVerificationPtr := flag.String("contracts.atlasVerification", "", "AtlasVerification contract address")
	contractsSimulatorPtr := flag.String("contracts.simulator", "", "Simulator contract address")
	relayGasMaxPerUserOperationPtr := flag.String("relay.gas.max_per_user_operation", "", "Max gas per user operation")
	relayGasMaxPerSolverOperationPtr := flag.String("relay.gas.max_per_solver_operation", "", "Max gas per solver operation")
	relayGasMaxPerDAppOperationPtr := flag.String("relay.gas.max_per_dApp_operation", "", "Max gas per dApp operation")
	flag.Parse()

	if len(*networkRpcUrlPtr) > 0 {
		c.Network.RpcUrl = *networkRpcUrlPtr
	}

	if len(*contractsAtlasPtr) > 0 {
		if !common.IsHexAddress(*contractsAtlasPtr) {
			panic("contracts.atlas is not a valid address")
		}
		c.Contracts.Atlas = common.HexToAddress(*contractsAtlasPtr)
	}

	if len(*contractsAtlasVerificationPtr) > 0 {
		if !common.IsHexAddress(*contractsAtlasVerificationPtr) {
			panic("contracts.atlasVerification is not a valid address")
		}
		c.Contracts.AtlasVerification = common.HexToAddress(*contractsAtlasVerificationPtr)
	}

	if len(*contractsSimulatorPtr) > 0 {
		if !common.IsHexAddress(*contractsSimulatorPtr) {
			panic("contracts.simulator is not a valid address")
		}
		c.Contracts.Simulator = common.HexToAddress(*contractsSimulatorPtr)
	}

	if len(*relayGasMaxPerUserOperationPtr) > 0 {
		c.Relay.Gas.MaxPerUserOperation = new(big.Int)
		c.Relay.Gas.MaxPerUserOperation.SetString(*relayGasMaxPerUserOperationPtr, 10)
	}

	if len(*relayGasMaxPerSolverOperationPtr) > 0 {
		c.Relay.Gas.MaxPerSolverOperation = new(big.Int)
		c.Relay.Gas.MaxPerSolverOperation.SetString(*relayGasMaxPerSolverOperationPtr, 10)
	}

	if len(*relayGasMaxPerDAppOperationPtr) > 0 {
		c.Relay.Gas.MaxPerDAppOperation = new(big.Int)
		c.Relay.Gas.MaxPerDAppOperation.SetString(*relayGasMaxPerDAppOperationPtr, 10)
	}
}

func (c *Config) parseEnv() {
	if envRpcUrl, set := os.LookupEnv("ETH_RPC_URL"); set && len(envRpcUrl) > 0 {
		c.Network.RpcUrl = envRpcUrl
	}
}

func (c *Config) validate() {
	if len(c.Network.RpcUrl) == 0 {
		panic("network.rpc_url is required")
	}

	if c.Contracts.Atlas == (common.Address{}) {
		panic("contracts.atlas is required")
	}

	if c.Contracts.AtlasVerification == (common.Address{}) {
		panic("contracts.atlasVerification is required")
	}

	if c.Contracts.Simulator == (common.Address{}) {
		panic("contracts.simulator is required")
	}
}
