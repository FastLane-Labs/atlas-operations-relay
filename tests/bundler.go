package tests

import (
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"math/big"
	"net/url"
	"time"

	"github.com/FastLane-Labs/atlas-operations-relay/contract/atlas"
	"github.com/FastLane-Labs/atlas-operations-relay/core"
	"github.com/FastLane-Labs/atlas-operations-relay/log"
	"github.com/FastLane-Labs/atlas-operations-relay/operation"
	"github.com/FastLane-Labs/atlas-operations-relay/utils"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/gorilla/websocket"
)

func runBundler(bundlerPk *ecdsa.PrivateKey, bundlerReceiveChan chan []byte, bundlerSendChan chan []byte) {
	bundlerAddr := crypto.PubkeyToAddress(bundlerPk.PublicKey)
	timestamp := time.Now().Unix()
	signatureContent := fmt.Sprintf("%s:%d", bundlerAddr, timestamp)
	signature, err := utils.SignEthereumMessage([]byte(signatureContent), bundlerPk)
	if err != nil {
		log.Error("failed to sign message:", err)
		return
	}

	u := url.URL{
		Scheme: "ws",
		Host:   "localhost:8080",
		Path:   "/ws/bundler",
		RawQuery: url.Values{
			"address":   []string{bundlerAddr.Hex()},
			"timestamp": []string{fmt.Sprintf("%d", timestamp)},
			"signature": []string{hexutil.Encode(signature)},
		}.Encode(),
	}

	conn, resp, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Error("dial:", err)
		return
	}

	if resp.StatusCode != 101 {
		log.Error("Expected status code 101, got", resp.StatusCode)
		return
	}

	log.Info("bundler connected", "bundlerEoa", bundlerAddr.Hex())

	// start listening on connection
	go func() {
		for {
			_, message, err := conn.ReadMessage()
			if err != nil {
				log.Error("bundler ws error:", err)
				if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
					log.Error("error:", err)
				}
				break
			}

			bundlerReceiveChan <- message
		}
	}()

	// start listening on bundlerSendChan
	go func() {
		for {
			message := <-bundlerSendChan
			err := conn.WriteMessage(websocket.TextMessage, message)
			if err != nil {
				log.Error("bundler ws error:", err)
				break
			}
		}
	}()
}

func sendBundlerResponse(txHash common.Hash, id string, bundlerSendChan chan []byte) error {
	bundleResponse := &core.BundleResponse{
		Id:     id,
		Result: txHash,
	}

	bundleResponseJSON, err := json.Marshal(bundleResponse)
	if err != nil {
		return err
	}

	bundlerSendChan <- bundleResponseJSON

	log.Info("bundler sent bundleResponse", "txHash", txHash.Hex())
	return nil
}

func newAtlasTx(bundleRequest *operation.BundleOperations) (*types.Transaction, error) {
	atlasContract, err := atlas.NewAtlas(conf.Contracts.Atlas, ethClient)
	if err != nil {
		panic(err)
	}

	signFn := func(addr common.Address, tx *types.Transaction) (*types.Transaction, error) {
		return types.SignTx(tx, types.NewEIP155Signer(big.NewInt(chainId)), bundlerPk)
	}

	gasPrice, err := ethClient.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, fmt.Errorf("can't get gas price suggestion: %w", err)
	}

	opts := &bind.TransactOpts{
		From:     bundlerEoa,
		GasPrice: gasPrice,
		Signer:   signFn,
		Nonce:    nil,
		Value:    nil,
		GasLimit: validationGasLimit + solverGasLimit,
		NoSend:   !sendAtlasTx,
	}

	atlas_userOp := atlas.UserOperation(*bundleRequest.UserOperation)
	atlas_dappOp := atlas.DAppOperation{
		From:          bundleRequest.DAppOperation.From,
		To:            bundleRequest.DAppOperation.To,
		Nonce:         bundleRequest.DAppOperation.Nonce,
		Deadline:      bundleRequest.DAppOperation.Deadline,
		Control:       bundleRequest.DAppOperation.Control,
		Bundler:       bundleRequest.DAppOperation.Bundler,
		UserOpHash:    bundleRequest.DAppOperation.UserOpHash,
		CallChainHash: bundleRequest.DAppOperation.CallChainHash,
		Signature:     bundleRequest.DAppOperation.Signature,
	}

	atlas_solverOps := make([]atlas.SolverOperation, len(bundleRequest.SolverOperations))
	for i, solverOp := range bundleRequest.SolverOperations {
		atlas_solverOps[i] = atlas.SolverOperation{
			From:         solverOp.From,
			To:           solverOp.To,
			Value:        solverOp.Value,
			Gas:          solverOp.Gas,
			MaxFeePerGas: solverOp.MaxFeePerGas,
			Deadline:     solverOp.Deadline,
			Solver:       solverOp.Solver,
			Control:      solverOp.Control,
			UserOpHash:   solverOp.UserOpHash,
			BidToken:     solverOp.BidToken,
			BidAmount:    solverOp.BidAmount,
			Data:         solverOp.Data,
			Signature:    solverOp.Signature,
		}
	}

	return atlasContract.Metacall(opts, atlas_userOp, atlas_solverOps, atlas_dappOp)
}
