package config

import (
	"crypto/ecdsa"
	"encoding/json"
	"flag"
	"math/big"
	"os"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

var (
	defaultAuctionDuration   = 500 * time.Millisecond
	defaultMaxSolutions      = uint64(10)
	defaultOperationGasLimit = big.NewInt(1000000)
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
		Auction struct {
			Duration     uint64 `json:"duration,omitempty"`
			MaxSolutions uint64 `json:"max_solutions,omitempty"`
		} `json:"auction,omitempty"`
		Gas struct {
			MaxPerUserOperation uint64 `json:"max_per_user_operation,omitempty"`
			MaxPerDAppOperation uint64 `json:"max_per_dApp_operation,omitempty"`
		} `json:"gas,omitempty"`
	} `json:"relay,omitempty"`
}

type Network struct {
	RpcUrl string
}

type Contracts struct {
	Atlas             common.Address
	AtlasVerification common.Address
	Simulator         common.Address
}

type Auction struct {
	Duration     time.Duration
	MaxSolutions uint64
}

type Gas struct {
	MaxPerUserOperation *big.Int
	MaxPerDAppOperation *big.Int
}

type Relay struct {
	Auction     Auction
	Gas         Gas
	Signatories map[common.Address]*ecdsa.PrivateKey
}

type Config struct {
	Network   Network
	Contracts Contracts
	Relay     Relay
}

func Load() *Config {
	config := &Config{
		Relay: Relay{
			Signatories: make(map[common.Address]*ecdsa.PrivateKey),
		},
	}

	config.parseConfigFile()
	config.parseFlags()
	config.parseEnv()
	config.Validate()

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

	c.Relay.Auction.Duration = time.Duration(configJson.Relay.Auction.Duration * uint64(time.Millisecond))
	c.Relay.Gas.MaxPerUserOperation = new(big.Int).SetUint64(configJson.Relay.Gas.MaxPerUserOperation)
	c.Relay.Gas.MaxPerDAppOperation = new(big.Int).SetUint64(configJson.Relay.Gas.MaxPerDAppOperation)
}

func (c *Config) parseFlags() {
	networkRpcUrlPtr := flag.String("network.rpc_url", "", "Ethereum RPC URL")
	contractsAtlasPtr := flag.String("contracts.atlas", "", "Atlas contract address")
	contractsAtlasVerificationPtr := flag.String("contracts.atlasVerification", "", "AtlasVerification contract address")
	contractsSimulatorPtr := flag.String("contracts.simulator", "", "Simulator contract address")
	relayAuctionDurationPtr := flag.Uint64("relay.auction.duration", 0, "Auction duration in milliseconds")
	relayAuctionMaxSolutionsPtr := flag.Uint64("relay.auction.max_solutions", 0, "Max solutions per auction")
	relayGasMaxPerUserOperationPtr := flag.Uint64("relay.gas.max_per_user_operation", 0, "Max gas per user operation")
	relayGasMaxPerDAppOperationPtr := flag.Uint64("relay.gas.max_per_dApp_operation", 0, "Max gas per dApp operation")
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

	if *relayAuctionDurationPtr > 0 {
		c.Relay.Auction.Duration = time.Duration(*relayAuctionDurationPtr * uint64(time.Millisecond))
	}

	if *relayAuctionMaxSolutionsPtr > 0 {
		c.Relay.Auction.MaxSolutions = *relayAuctionMaxSolutionsPtr
	}

	if *relayGasMaxPerUserOperationPtr > 0 {
		c.Relay.Gas.MaxPerUserOperation = new(big.Int).SetUint64(*relayGasMaxPerUserOperationPtr)
	}

	if *relayGasMaxPerDAppOperationPtr > 0 {
		c.Relay.Gas.MaxPerDAppOperation = new(big.Int).SetUint64(*relayGasMaxPerDAppOperationPtr)
	}
}

func (c *Config) parseEnv() {
	if envRpcUrl, set := os.LookupEnv("ETH_RPC_URL"); set && len(envRpcUrl) > 0 {
		c.Network.RpcUrl = envRpcUrl
	}

	if envSignatoriesPks, set := os.LookupEnv("SIGNATORIES_PKS"); set && len(envSignatoriesPks) > 0 {
		for _, pk := range strings.Split(envSignatoriesPks, ",") {
			privateKey, err := crypto.HexToECDSA(pk)
			if err != nil {
				panic(err)
			}
			c.Relay.Signatories[crypto.PubkeyToAddress(privateKey.PublicKey)] = privateKey

		}
	}
}

func (c *Config) Validate() {
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

	if c.Relay.Auction.Duration == 0 {
		c.Relay.Auction.Duration = defaultAuctionDuration
	}

	if c.Relay.Auction.MaxSolutions == 0 {
		c.Relay.Auction.MaxSolutions = defaultMaxSolutions
	}

	if c.Relay.Gas.MaxPerUserOperation == nil || c.Relay.Gas.MaxPerUserOperation.Cmp(common.Big0) == 0 {
		c.Relay.Gas.MaxPerUserOperation = new(big.Int).Set(defaultOperationGasLimit)
	}

	if c.Relay.Gas.MaxPerDAppOperation == nil || c.Relay.Gas.MaxPerDAppOperation.Cmp(common.Big0) == 0 {
		c.Relay.Gas.MaxPerDAppOperation = new(big.Int).Set(defaultOperationGasLimit)
	}
}
