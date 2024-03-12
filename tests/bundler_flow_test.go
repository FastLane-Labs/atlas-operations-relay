package tests

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"net/http"
	"testing"
	"time"

	"github.com/FastLane-Labs/atlas-operations-relay/auction"
	"github.com/FastLane-Labs/atlas-operations-relay/contract/atlas"
	"github.com/FastLane-Labs/atlas-operations-relay/core"
	"github.com/FastLane-Labs/atlas-operations-relay/operation"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

func TestBundlerFlow(t *testing.T) {
	//start solver
	solverDoneChan := make(chan struct{})
	go runSolver(solverDoneChan)

	//start bundler
	bundlerReceiveChan := make(chan []byte)
	go runBundler(bundlerPk, bundlerReceiveChan)

	//send user request
	userOp, err := sendUserRequest()
	if err != nil {
		t.Fatal(err)
	}

	//wait for solver to finish
	<-solverDoneChan

	//wait for auction to end
	time.Sleep(auction.AuctionDuration)

	//user requests solver solutions
	userOpHash, _ := userOp.Hash()
	solverOps, err := retreiveSolverOps(userOpHash)
	if err != nil {
		t.Fatal(err)
	}

	if len(solverOps) == 0 {
		t.Fatal("expected at least one solver operation")
	}

	//send bundleOps to the relay
	err = sendBundleOperation(t, userOp, solverOps, NewDappOperation(userOp, solverOps))
	if err != nil {
		t.Fatal(err)
	}

	//wait till bundler receives bundleOps
	bundleOpBytes := <-bundlerReceiveChan
	bundleRequest := &core.BundleRequest{}
	err = json.Unmarshal(bundleOpBytes, bundleRequest)
	if err != nil {
		t.Fatal(err)
	}

	//send atlas tx
	atlasContract, err := atlas.NewAtlas(conf.Contracts.Atlas, ethClient)
	if err != nil {
		panic(err)
	}

	signFn := func(addr common.Address, tx *types.Transaction) (*types.Transaction, error) {
		return types.SignTx(tx, types.NewEIP155Signer(big.NewInt(int64(11155111))), bundlerPk)
	}

	gasPrice, err := ethClient.SuggestGasPrice(context.Background())
	if err != nil {
		fmt.Println("Can't get gas price suggestion", err)
		return
	}

	fmt.Println("Gas price", gasPrice.Text(10))

	opts := &bind.TransactOpts{
		From:     bundlerEoa,
		GasPrice: gasPrice,
		Signer:   signFn,
		Nonce:    nil,
		Value:    nil,
		GasLimit: uint64(5_000_000),
		NoSend:   true,
	}

	atlas_userOp := atlas.UserOperation{
		From:         bundleRequest.Bundle.UserOperation.From,
		To:           bundleRequest.Bundle.UserOperation.To,
		Value:        bundleRequest.Bundle.UserOperation.Value,
		Gas:          bundleRequest.Bundle.UserOperation.Gas,
		MaxFeePerGas: bundleRequest.Bundle.UserOperation.MaxFeePerGas,
		Nonce:        bundleRequest.Bundle.UserOperation.Nonce,
		Deadline:     bundleRequest.Bundle.UserOperation.Deadline,
		Dapp:         bundleRequest.Bundle.UserOperation.Dapp,
		Control:      bundleRequest.Bundle.UserOperation.Control,
		SessionKey:   bundleRequest.Bundle.UserOperation.SessionKey,
		Data:         bundleRequest.Bundle.UserOperation.Data,
		Signature:    bundleRequest.Bundle.UserOperation.Signature,
	}

	atlas_dappOp := atlas.DAppOperation{
		From:          bundleRequest.Bundle.DAppOperation.From,
		To:            bundleRequest.Bundle.DAppOperation.To,
		Value:         bundleRequest.Bundle.DAppOperation.Value,
		Gas:           bundleRequest.Bundle.DAppOperation.Gas,
		Nonce:         bundleRequest.Bundle.DAppOperation.Nonce,
		Deadline:      bundleRequest.Bundle.DAppOperation.Deadline,
		Control:       bundleRequest.Bundle.DAppOperation.Control,
		Bundler:       bundleRequest.Bundle.DAppOperation.Bundler,
		UserOpHash:    bundleRequest.Bundle.DAppOperation.UserOpHash,
		CallChainHash: bundleRequest.Bundle.DAppOperation.CallChainHash,
		Signature:     bundleRequest.Bundle.DAppOperation.Signature,
	}

	atlas_solverOps := make([]atlas.SolverOperation, len(bundleRequest.Bundle.SolverOperations))
	for i, solverOp := range bundleRequest.Bundle.SolverOperations {
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

	tx, err := atlasContract.Metacall(opts, atlas_userOp, atlas_solverOps, atlas_dappOp)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Atlas tx sent", tx.Hash().Hex())
}

func sendBundleOperation(t *testing.T, userOp *operation.UserOperation, solverOps []*operation.SolverOperation, dAppOp *operation.DAppOperation) error {
	bundleOps := &operation.BundleOperations{
		UserOperation:    userOp,
		SolverOperations: solverOps,
		DAppOperation:    dAppOp,
	}

	bundleOpsJSON, err := json.Marshal(bundleOps)
	if err != nil {
		t.Errorf("failed to marshal bundle operations: %v", err)
	}

	_, err = http.Post("http://localhost:8080/bundleOperations", "application/json", bytes.NewReader(bundleOpsJSON))
	if err != nil {
		return err
	}

	return nil
}
