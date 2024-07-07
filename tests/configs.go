package tests

import (
	"time"

	"github.com/FastLane-Labs/atlas-operations-relay/config"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/signer/core/apitypes"
)

var (
	ethClient *ethclient.Client

	chainId     int64 = 11155111 //sepolia
	sendAtlasTx bool  = false    //true -> send tx to the network, false -> do everything apart from sending

	validationGasLimit = uint64(500_000)
	solverGasLimit     = uint64(1_000_000)

	tokenA                = common.HexToAddress("0x7439E9Bb6D8a84dd3A23fe621A30F95403F87fB9")
	tokenB                = common.HexToAddress("0x7b79995e5f793A07Bc00c21412e50Ecae098E7f9")
	swapIntentDAppControl = common.HexToAddress("0x60d7B59c6743C25b29a7aEe6F5a37c07B1A6Cff3")
	simpleRfqSolver       = common.HexToAddress("0x12f1D679c09A050c2b8259e4B1dE14e315F07822")

	userPk, _ = crypto.ToECDSA(common.FromHex("0d3414024a8d727a824933d47460fd9ea5d65f88feec92761a476405cf2d5922"))
	userEoa   = crypto.PubkeyToAddress(userPk.PublicKey) // 0xeA402251DA4365c12BF9A3C9d88029A04988A712

	solverPk, _ = crypto.ToECDSA(common.FromHex("e3818ba09a3106ecbf80cafed7510df6f81e6dfab3cfa240680ec9389ba66206"))
	solverEoa   = crypto.PubkeyToAddress(solverPk.PublicKey) // 0x915c0f7A6a3be9E2298De421C4FE6CA1E2194880

	bundlerPk, _ = crypto.ToECDSA(common.FromHex("586af44cb6f500bdcdbea0e4411916dfc4806e7df43504da5bdfe144dd78f895"))
	bundlerEoa   = crypto.PubkeyToAddress(bundlerPk.PublicKey) // 0xEdB89106f2293ed2bAAbA1e8E844306412cB39Fe

	// The governance account that deployed the dApp control contract is the same as the bundler.
	// Separating variables for clarity.
	governancePk  = bundlerPk
	governanceEoa = bundlerEoa

	conf = &config.Config{
		Network: config.Network{
			ChainId: uint64(chainId),
			RpcUrl:  "https://rpc.sepolia.org/",
		},
		Contracts: config.Contracts{
			Atlas:             common.HexToAddress("0x1Be854EeA3D753db001aC7A1aaE7Eb30f9B1166a"),
			AtlasVerification: common.HexToAddress("0x26Bb4e798116Bb01f26A47EDA2597814BDC18467"),
			Simulator:         common.HexToAddress("0xCF9Db077FFC7Ae39210e00468bf94021adFb51a0"),
		},
		Relay: config.Relay{
			Simulations: true,
			Auction: config.Auction{
				Duration: 5 * time.Second,
			},
			Eip712: config.Eip712{
				Domain: apitypes.TypedDataDomain{
					Name:              "AtlasVerification",
					Version:           "1.0",
					ChainId:           math.NewHexOrDecimal256(chainId),
					VerifyingContract: "0x26Bb4e798116Bb01f26A47EDA2597814BDC18467",
				},
			},
		},
	}
)
