package tests

import (
	"bytes"
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"math/big"
	"net/http"
	"strings"
	"testing"

	"github.com/FastLane-Labs/atlas-operations-relay/config"
	"github.com/FastLane-Labs/atlas-operations-relay/contract/atlasVerification"
	"github.com/FastLane-Labs/atlas-operations-relay/log"
	"github.com/FastLane-Labs/atlas-operations-relay/operation"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func TestUserOpFlow(t *testing.T) {
	ethClient, err := ethclient.Dial(conf.Network.RpcUrl)
	if err != nil {
		log.Error("failed to connect to the Ethereum client", "err", err)
		return
	}

	userOp := generateDemoValidUserOp(ethClient, conf)
	userOpHash, err := userOp.Hash()
	userOpJSON, err := json.Marshal(userOp)
	if err != nil {
		panic(err)
	}

	resp, err := http.Post("http://localhost:8080/userOperation", "application/json", bytes.NewReader(userOpJSON))
	if err != nil {
		fmt.Println("Error sending HTTP request:", err)
		return
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, resp.StatusCode)
	}

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	userOpHashReceived := common.HexToHash(strings.Trim(string(bodyBytes), "\""))

	if userOpHashReceived != userOpHash {
		t.Errorf("Expected user operation hash %s, got %s", userOpHash.Hex(), userOpHashReceived.Hex())
	}
}

func generateDemoValidUserOp(ethClient *ethclient.Client, config *config.Config) *operation.UserOperation {
	currentBlock, err := ethClient.BlockNumber(context.Background())
	if err != nil {
		panic(err)
	}

	atlasVerificationContract, err := atlasVerification.NewAtlasVerification(config.Contracts.AtlasVerification, ethClient)
	if err != nil {
		panic(err)
	}

	atlasDomainSeparator, err := atlasVerificationContract.GetDomainSeparator(nil)
	if err != nil {
		panic(err)
	}

	privateKeyBytes, err := hex.DecodeString("54c371fc7e513a4574cf87ca222664580e9c3fa58ecc636a2bd811bddfac66a2")
	if err != nil {
		panic(err)
	}

	privateKey, err := crypto.ToECDSA(privateKeyBytes)
	if err != nil {
		panic(err)
	}

	publicKey := crypto.PubkeyToAddress(privateKey.PublicKey)

	dataBytes, err := hex.DecodeString("83a6992a00000000000000000000000000000000000000000000000000000000000000200000000000000000000000007439E9Bb6D8a84dd3A23fe621A30F95403F87fB90000000000000000000000000000000000000000000000000000b5e620f480000000000000000000000000007b79995e5f793a07bc00c21412e50ecae098e7f9000000000000000000000000000000000000000000000000000000e8d4a510000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000e00000000000000000000000000000000000000000000000000000000000000000")
	if err != nil {
		panic(err)
	}

	userOp := &operation.UserOperation{
		From:         publicKey,
		To:           config.Contracts.Atlas,
		Deadline:     big.NewInt(int64(currentBlock) + 1),
		Gas:          big.NewInt(100000),
		Nonce:        big.NewInt(4),
		MaxFeePerGas: big.NewInt(1000000000),
		Value:        big.NewInt(0),
		Dapp:         common.HexToAddress("0x2F98675731D0e659E716d890330901a8A2355813"),
		Control:      common.HexToAddress("0x2F98675731D0e659E716d890330901a8A2355813"),
		SessionKey:   common.HexToAddress("0x0"),
		Data:         dataBytes,
		Signature:    nil,
	}

	proofHash, err := userOp.ProofHash()
	if err != nil {
		panic(err)
	}

	payload := crypto.Keccak256Hash([]byte("\x19\x01"), atlasDomainSeparator[:], proofHash[:])
	signature, err := crypto.Sign(payload[:], privateKey)
	if err != nil {
		panic(err)
	}

	userOp.Signature = signature

	if err := userOp.Validate(ethClient, config.Contracts.Atlas, atlasDomainSeparator, nil); err != nil {
		panic("user operation is invalid " + err.Message)
	}

	return userOp
}
