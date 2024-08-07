package tests

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"math/big"
	"net/http"
	"net/url"
	"strings"

	"github.com/FastLane-Labs/atlas-operations-relay/contract/atlasVerification"
	"github.com/FastLane-Labs/atlas-operations-relay/contract/dAppControl"
	"github.com/FastLane-Labs/atlas-operations-relay/log"
	"github.com/FastLane-Labs/atlas-operations-relay/operation"
	"github.com/FastLane-Labs/atlas-operations-relay/utils"
	"github.com/ethereum/go-ethereum/common"
)

func newDemoSwapIntent() *SwapIntent {
	return &SwapIntent{
		TokenUserBuys:       tokenA,
		AmountUserBuys:      big.NewInt(200e12),
		TokenUserSells:      tokenB,
		AmountUserSells:     big.NewInt(1e12),
		AuctionBaseCurrency: common.HexToAddress("0x0"),
		Conditions:          make([]Condition, 0),
	}
}

func newDemoUserOperation(sessionKey common.Address) *operation.UserOperation {
	currentBlock, err := ethClient.BlockNumber(context.Background())
	if err != nil {
		panic(err)
	}

	data, err := newDemoSwapIntent().abiEncodeWithSelector()
	if err != nil {
		panic(err)
	}

	atlasVerification, err := atlasVerification.NewAtlasVerification(conf.Contracts.AtlasVerification, ethClient)
	if err != nil {
		panic(err)
	}

	nonce, err := atlasVerification.GetUserNextNonce(nil, userEoa, false)
	if err != nil {
		panic(err)
	}

	dAppControlContract, err := dAppControl.NewDAppControl(swapIntentDAppControl, ethClient)
	if err != nil {
		panic(err)
	}

	callConfig, err := dAppControlContract.CALLCONFIG(nil)
	if err != nil {
		panic(err)
	}

	userOp := &operation.UserOperation{
		From:         userEoa,
		To:           conf.Contracts.Atlas,
		Deadline:     big.NewInt(int64(currentBlock) + 1000),
		Gas:          big.NewInt(100000),
		Nonce:        nonce,
		MaxFeePerGas: big.NewInt(150e9),
		Value:        big.NewInt(0),
		Dapp:         swapIntentDAppControl,
		Control:      swapIntentDAppControl,
		CallConfig:   callConfig,
		SessionKey:   sessionKey,
		Data:         data,
		Signature:    nil,
	}

	// User always signs the full hash, hence `false` is passed
	userOpHashForSigning, relayErr := userOp.Hash(false, &conf.Relay.Eip712.Domain)
	if relayErr != nil {
		panic(relayErr)
	}

	userOp.Signature, _ = utils.SignMessage(userOpHashForSigning.Bytes(), userPk)

	if err := userOp.Validate(ethClient, conf.Contracts.Atlas, &conf.Relay.Eip712.Domain, conf.Relay.Gas.MaxPerUserOperation); err != nil {
		panic(err)
	}

	return userOp
}

func sendUserRequest(userOp *operation.UserOperation) error {
	userOpHash, relayErr := userOp.Hash(utils.FlagTrustedOpHash(userOp.CallConfig), &conf.Relay.Eip712.Domain)
	if relayErr != nil {
		return relayErr
	}

	userOpWithHints := &operation.UserOperationWithHintsRaw{
		UserOperation: userOp.EncodeToRaw(),
		Hints:         []common.Address{},
	}

	reqJSON, err := json.Marshal(userOpWithHints)
	if err != nil {
		return fmt.Errorf("failed to marshal userOp: %w", err)
	}

	resp, err := http.Post("http://localhost:8080/userOperation", "application/json", bytes.NewReader(reqJSON))
	if err != nil {
		return fmt.Errorf("failed to send userOp: %w", err)
	}

	log.Info("user sent userOp", "userOpHash", userOpHash.Hex())

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("expected status code 200, got %d", resp.StatusCode)
	}

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %w", err)
	}

	userOpHashInResp := common.HexToHash(strings.Trim(string(bodyBytes), "\""))

	if userOpHashInResp != userOpHash {
		return fmt.Errorf("expected userOpHash %s, got %s", userOpHash, userOpHashInResp)
	}

	return nil
}

func retreiveSolverOps(userOpHash common.Hash, wait bool) ([]*operation.SolverOperation, error) {
	u := url.URL{
		Scheme: "http",
		Host:   "localhost:8080",
		Path:   "/solverOperations",
		RawQuery: url.Values{
			"operationHash": []string{userOpHash.Hex()},
			"wait":          []string{fmt.Sprintf("%t", wait)},
		}.Encode(),
	}

	resp, err := http.Get(u.String())
	if err != nil {
		return make([]*operation.SolverOperation, 0), err
	}

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, fmt.Errorf("failed to read response body: %w", err)
		}
		return nil, fmt.Errorf("expected status code 200, got %d: %s", resp.StatusCode, string(bodyBytes))
	}

	var solverOpsWithScoreRaw []*operation.SolverOperationWithScoreRaw

	if err = json.NewDecoder(resp.Body).Decode(&solverOpsWithScoreRaw); err != nil {
		return make([]*operation.SolverOperation, 0), err
	}

	solverOps := make([]*operation.SolverOperation, 0, len(solverOpsWithScoreRaw))
	for _, solverOpWithScoreRaw := range solverOpsWithScoreRaw {
		solverOps = append(solverOps, solverOpWithScoreRaw.SolverOperation.Decode())
	}

	return solverOps, nil
}

func sendBundleOperation(userOp *operation.UserOperation, solverOps []*operation.SolverOperation, dAppOp *operation.DAppOperation) error {
	bundleOps := &operation.BundleOperations{
		UserOperation:    userOp,
		SolverOperations: solverOps,
		DAppOperation:    dAppOp,
	}

	bundleOpsJSON, err := json.Marshal(bundleOps.EncodeToRaw())
	if err != nil {
		return err
	}

	resp, err := http.Post("http://localhost:8080/bundleOperations", "application/json", bytes.NewReader(bundleOpsJSON))
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			return fmt.Errorf("failed to read response body: %w", err)
		}
		return fmt.Errorf("expected status code 200, got %d: %s", resp.StatusCode, string(bodyBytes))
	}

	userOpHash, _ := userOp.Hash(false, &conf.Relay.Eip712.Domain)
	log.Info("relay sent bundleOps", "userOpHash", userOpHash.Hex(), "nSolverOps", len(bundleOps.SolverOperations))

	return nil
}

func retrieveAtlasTxHash(userOpHash common.Hash, wait bool) (common.Hash, error) {
	u := url.URL{
		Scheme: "http",
		Host:   "localhost:8080",
		Path:   "/bundleHash",
		RawQuery: url.Values{
			"operationHash": []string{userOpHash.Hex()},
			"wait":          []string{fmt.Sprintf("%t", wait)},
		}.Encode(),
	}

	resp, err := http.Get(u.String())
	if err != nil {
		return common.Hash{}, err
	}

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			return common.Hash{}, fmt.Errorf("failed to read response body: %w", err)
		}
		return common.Hash{}, fmt.Errorf("expected status code 200, got %d: %s", resp.StatusCode, string(bodyBytes))
	}

	var txHash common.Hash
	err = json.NewDecoder(resp.Body).Decode(&txHash)
	if err != nil {
		return common.Hash{}, err
	}

	return txHash, nil
}
