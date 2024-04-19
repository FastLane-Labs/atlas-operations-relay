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

	chainId     int64 = 11155111 //sepolia
	sendAtlasTx bool  = false    //true -> send tx to the network, false -> do everything apart from sending

	validationGasLimit = uint64(500_000)
	solverGasLimit     = uint64(1_000_000)

	tokenA                = common.HexToAddress("0x7439E9Bb6D8a84dd3A23fe621A30F95403F87fB9")
	tokenB                = common.HexToAddress("0x7b79995e5f793A07Bc00c21412e50Ecae098E7f9")
	swapIntentDAppControl = common.HexToAddress("0x632767B79E67E30914c221CdE4383d1FeaE4768E")
	simpleRfqSolver       = common.HexToAddress("0x8729a958873d1B44B5A089d654bb9112A33cE624")

	userPk, _ = crypto.ToECDSA(common.FromHex("0d3414024a8d727a824933d47460fd9ea5d65f88feec92761a476405cf2d5922"))
	userEoa   = crypto.PubkeyToAddress(userPk.PublicKey) // 0xeA402251DA4365c12BF9A3C9d88029A04988A712

	solverPk, _ = crypto.ToECDSA(common.FromHex("e3818ba09a3106ecbf80cafed7510df6f81e6dfab3cfa240680ec9389ba66206"))
	solverEoa   = crypto.PubkeyToAddress(solverPk.PublicKey) // 0x915c0f7A6a3be9E2298De421C4FE6CA1E2194880

	bundlerPk, _ = crypto.ToECDSA(common.FromHex("586af44cb6f500bdcdbea0e4411916dfc4806e7df43504da5bdfe144dd78f895"))
	bundlerEoa   = crypto.PubkeyToAddress(bundlerPk.PublicKey) // 0xEdB89106f2293ed2bAAbA1e8E844306412cB39Fe

	conf = &config.Config{
		Network: config.Network{
			RpcUrl: "https://rpc.sepolia.org/",
		},
		Contracts: config.Contracts{
			Atlas:             common.HexToAddress("0xa892eb9F79E0D1b6277B3456b0a8FE770386f6DB"),
			AtlasVerification: common.HexToAddress("0xeeB91b2d317e3A747E88c1CA542ae31E32B87FDF"),
			Simulator:         common.HexToAddress("0xAAdF6272cCE4121Db92da224C28d1B59C9feF4d5"),
		},
		Relay: config.Relay{
			Auction: config.Auction{
				Duration: 5 * time.Second,
			},
		},
	}
)
