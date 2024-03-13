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

	"github.com/FastLane-Labs/atlas-operations-relay/log"
	"github.com/FastLane-Labs/atlas-operations-relay/operation"
	"github.com/ethereum/go-ethereum/common"
)

func newDemoSwapIntent() *SwapIntent {
	return &SwapIntent{
		TokenUserBuys:          tokenA,
		AmountUserBuys:         big.NewInt(200e12),
		TokenUserSells:         tokenB,
		AmountUserSells:        big.NewInt(1e12),
		AuctionBaseCurrency:    common.HexToAddress("0x0"),
		SolverMustReimburseGas: false,
		Conditions:             make([]Condition, 0),
	}
}

func newDemoUserOperation() *operation.UserOperation {
	currentBlock, err := ethClient.BlockNumber(context.Background())
	if err != nil {
		panic(err)
	}

	data, err := newDemoSwapIntent().abiEncodeWithSelector()
	if err != nil {
		panic(err)
	}

	userOp := &operation.UserOperation{
		From:         userEoa,
		To:           conf.Contracts.Atlas,
		Deadline:     big.NewInt(int64(currentBlock) + 1000),
		Gas:          big.NewInt(100000),
		Nonce:        big.NewInt(7),
		MaxFeePerGas: big.NewInt(20e9),
		Value:        big.NewInt(0),
		Dapp:         swapIntentDAppControl,
		Control:      swapIntentDAppControl,
		SessionKey:   common.HexToAddress("0x0"),
		Data:         data,
		Signature:    nil,
	}

	proofHash, err := userOp.ProofHash()
	if err != nil {
		panic(err)
	}

	userOp.Signature = signEip712(atlasDomainSeparator, proofHash, userPk)

	if err := userOp.Validate(ethClient, conf.Contracts.Atlas, atlasDomainSeparator, conf.Relay.Gas.MaxPerUserOperation); err != nil {
		panic(err)
	}

	return userOp
}

func sendUserRequest() (*operation.UserOperation, error) {
	userOp := newDemoUserOperation()
	userOpHash, relayErr := userOp.Hash()
	if relayErr != nil {
		return nil, relayErr
	}

	userOpJSON, err := json.Marshal(userOp)
	if err != nil {
		return nil, fmt.Errorf("Failed to marshal userOp: %w", err)
	}

	resp, err := http.Post("http://localhost:8080/userOperation", "application/json", bytes.NewReader(userOpJSON))
	if err != nil {
		return nil, fmt.Errorf("Failed to send userOp: %w", err)
	}

	log.Info("user sent userOp", "userOpHash", userOpHash.Hex())

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Expected status code 200, got %d", resp.StatusCode)
	}

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("Failed to read response body: %w", err)
	}

	userOpHashInResp := common.HexToHash(strings.Trim(string(bodyBytes), "\""))

	if userOpHashInResp != userOpHash {
		return nil, fmt.Errorf("Expected userOpHash %s, got %s", userOpHash, userOpHashInResp)
	}

	return userOp, nil
}

func retreiveSolverOps(userOpHash common.Hash, wait bool) ([]*operation.SolverOperation, error) {
	u := url.URL{
		Scheme: "http",
		Host:   "localhost:8080",
		Path:   "/solverOperations",
		RawQuery: url.Values{
			"userOpHash": []string{userOpHash.Hex()},
			"wait":       []string{fmt.Sprintf("%t", wait)},
		}.Encode(),
	}

	resp, err := http.Get(u.String())
	if err != nil {
		return make([]*operation.SolverOperation, 0), err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return make([]*operation.SolverOperation, 0), fmt.Errorf("Expected status code 200, got %d", resp.StatusCode)
	}

	var solverOps []*operation.SolverOperation
	err = json.NewDecoder(resp.Body).Decode(&solverOps)

	if err != nil {
		return make([]*operation.SolverOperation, 0), err
	}

	return solverOps, nil
}

func sendBundleOperation(userOp *operation.UserOperation, solverOps []*operation.SolverOperation, dAppOp *operation.DAppOperation) error {
	bundleOps := &operation.BundleOperations{
		UserOperation:    userOp,
		SolverOperations: solverOps,
		DAppOperation:    dAppOp,
	}

	bundleOpsJSON, err := json.Marshal(bundleOps)
	if err != nil {
		return err
	}

	_, err = http.Post("http://localhost:8080/bundleOperations", "application/json", bytes.NewReader(bundleOpsJSON))
	if err != nil {
		return err
	}

	userOpHash, _ := userOp.Hash()
	log.Info("relay sent bundleOps", "userOpHash", userOpHash.Hex(), "nSolverOps", len(bundleOps.SolverOperations))

	return nil
}

func retrieveAtlasTxHash(userOpHash common.Hash, wait bool) (common.Hash, error) {
	u := url.URL{
		Scheme: "http",
		Host:   "localhost:8080",
		Path:   "/bundleHash",
		RawQuery: url.Values{
			"userOpHash": []string{userOpHash.Hex()},
			"wait":       []string{fmt.Sprintf("%t", wait)},
		}.Encode(),
	}

	resp, err := http.Get(u.String())
	if err != nil {
		return common.Hash{}, err
	}

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return common.Hash{}, fmt.Errorf("Expected status code 200, got %d", resp.StatusCode)
	}

	var txHash common.Hash
	err = json.NewDecoder(resp.Body).Decode(&txHash)
	if err != nil {
		return common.Hash{}, err
	}

	return txHash, nil
}