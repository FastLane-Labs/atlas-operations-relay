package config

import (
	"encoding/json"
	"flag"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/common"
)

type Config struct {
	Network *struct {
		RpcUrl string `json:"rpc_url"`
	} `json:"network"`

	Contracts *struct {
		Atlas             common.Address `json:"atlas,omitempty"`
		AtlasVerification common.Address `json:"atlasVerification,omitempty"`
		Simulator         common.Address `json:"simulator,omitempty"`
	} `json:"contracts"`

	Relay *struct {
		Gas *struct {
			MaxPerUserOperation   *big.Int `json:"max_per_user_operation,omitempty"`
			MaxPerSolverOperation *big.Int `json:"max_per_solver_operation,omitempty"`
			MaxPerDAppOperation   *big.Int `json:"max_per_dApp_operation,omitempty"`
		} `json:"gas"`
	} `json:"relay"`
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

	config.parseFlags()

	if envRpcUrl, set := os.LookupEnv("ETH_RPC_URL"); set && len(envRpcUrl) > 0 {
		config.Network.RpcUrl = envRpcUrl
	}

	return &config
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
		c.Contracts.Atlas = common.HexToAddress(*contractsAtlasPtr)
	}

	if len(*contractsAtlasVerificationPtr) > 0 {
		c.Contracts.AtlasVerification = common.HexToAddress(*contractsAtlasVerificationPtr)
	}

	if len(*contractsSimulatorPtr) > 0 {
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
