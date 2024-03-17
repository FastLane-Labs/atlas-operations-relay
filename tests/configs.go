package tests

import (
	"time"

	"github.com/FastLane-Labs/atlas-operations-relay/config"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	ethClient            *ethclient.Client
	atlasDomainSeparator common.Hash

	chainId int64 = 11155111 //sepolia

	sendAtlasTx bool = false //true -> send tx to the network, false -> do everything apart from sending

	tokenA                = common.HexToAddress("0x7439E9Bb6D8a84dd3A23fe621A30F95403F87fB9")
	tokenB                = common.HexToAddress("0x7b79995e5f793A07Bc00c21412e50Ecae098E7f9")
	swapIntentDAppControl = common.HexToAddress("0x2F98675731D0e659E716d890330901a8A2355813")
	simpleRfqSolver       = common.HexToAddress("0xbB4026B66fB10CF380Cb069F5ac517Ed5caAdc34")

	userPk, _ = crypto.ToECDSA(common.FromHex("0x54c371fc7e513a4574cf87ca222664580e9c3fa58ecc636a2bd811bddfac66a2"))
	userEoa   = crypto.PubkeyToAddress(userPk.PublicKey) // 0x7F14D219C66c53A638185146A20555bb6a3f234A

	solverPk, _ = crypto.ToECDSA(common.FromHex("e3818ba09a3106ecbf80cafed7510df6f81e6dfab3cfa240680ec9389ba66206"))
	solverEoa   = crypto.PubkeyToAddress(solverPk.PublicKey) // 0x915c0f7A6a3be9E2298De421C4FE6CA1E2194880

	bundlerPk, _ = crypto.ToECDSA(common.FromHex("586af44cb6f500bdcdbea0e4411916dfc4806e7df43504da5bdfe144dd78f895"))
	bundlerEoa   = crypto.PubkeyToAddress(bundlerPk.PublicKey) // 0xEdB89106f2293ed2bAAbA1e8E844306412cB39Fe

	conf = &config.Config{
		Network: config.Network{
			RpcUrl: "https://rpc.sepolia.org/",
		},
		Contracts: config.Contracts{
			Atlas:             common.HexToAddress("0x79d1379195f1Ed373eF8c58aC36F9C1045f8684d"),
			AtlasVerification: common.HexToAddress("0xf8e760554aB5c7E1Da51fFeD9C48de78bddD4c53"),
			Simulator:         common.HexToAddress("0x165877D0E2646bf7B42621D1551a23b94B14EfF9"),
		},
		Relay: config.Relay{
			Auction: config.Auction{
				Duration: 3 * time.Second,
			},
		},
	}
)
