package tests

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/big"
	"net/http"
	"testing"
	"time"

	"github.com/FastLane-Labs/atlas-operations-relay/contract/atlasVerification"
	"github.com/FastLane-Labs/atlas-operations-relay/core"
	"github.com/FastLane-Labs/atlas-operations-relay/log"
	"github.com/FastLane-Labs/atlas-operations-relay/operation"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gorilla/websocket"
)

func TestSolverSolves(t *testing.T) {
	doneChan := make(chan struct{})

	go runSolver(ethClient, doneChan)

	go runUser(ethClient)

	<-doneChan
}

func runUser(ethClient *ethclient.Client) {
	userOp := generateDemoValidUserOp(ethClient, conf)
	userOpHash, _ := userOp.Hash()
	userOpJSON, err := json.Marshal(userOp)
	if err != nil {
		panic(err)
	}

	_, err = http.Post("http://localhost:8080/userOperation", "application/json", bytes.NewReader(userOpJSON))
	if err != nil {
		fmt.Println("Error sending HTTP request:", err)
		return
	}

	fmt.Println("Sent user operation", userOpHash.Hex())
}

func runSolver(ethClient *ethclient.Client, doneChan chan struct{}) {
	//solver ws connection
	conn, solverResp := getWsConnection()
	defer conn.Close()

	if solverResp.StatusCode != 101 {
		panic("Expected status code 101, got " + fmt.Sprint(solverResp.StatusCode))
	}

	//track what's being received
	subscribedChan := make(chan struct{})
	userOpReceivedChan := make(chan []byte)

	//start listening on connection
	go func() {
		for {
			_, message, err := conn.ReadMessage()
			if err != nil {
				if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
					log.Error("error:", err)
				}
				break
			}

			response := &core.Response{}
			broadcast := &core.Broadcast{}

			json.Unmarshal(message, response)
			json.Unmarshal(message, broadcast)

			//if response and broadcast are both empty, panic
			if response.Id == "" && broadcast.Topic == "" {
				panic("cannot handle msg " + string(message))
			}
			if response.Id != "" {
				subscribedChan <- struct{}{}
			}
			if broadcast.Topic != "" {
				userOpReceivedChan <- message
			}
		}
	}()

	//subscribe to newUserOperations topic
	req := core.Request{
		Id:     "1234",
		Method: "subscribe",
		Params: &core.RequestParams{
			Topic: "newUserOperations",
		},
	}
	msg, err := json.Marshal(req)
	if err != nil {
		panic(err)
	}
	if err := conn.WriteMessage(websocket.TextMessage, msg); err != nil {
		panic(err)
	}

	//dont proceed until subscribled
	if !waitOnChanFor(subscribedChan, 1*time.Second) {
		panic("not subscribed to newUserOperations topic")
	}

	//wait for userOp to be received
	userOpBroadcastBytes := <-userOpReceivedChan
	broadcast := &core.Broadcast{}

	err = json.Unmarshal(userOpBroadcastBytes, broadcast)
	if err != nil {
		panic(err)
	}

	userOp := broadcast.Data.UserOperation
	userOpHash, _ := userOp.Hash()
	fmt.Println("Received user operation", userOpHash.Hex())

	//generate solver operation
	solverOp := generateSolverOperation(userOp, ethClient)

	//send solver operation
	solverOpJSON, err := json.Marshal(solverOp)
	if err != nil {
		panic(err)
	}

	resp, err := http.Post("http://localhost:8080/solverOperation", "application/json", bytes.NewReader(solverOpJSON))
	if err != nil {
		fmt.Println("Error sending HTTP request:", err)
		return
	}

	defer resp.Body.Close()
	fmt.Println("Sent solver operation", userOpHash.Hex())
	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	fmt.Println("response Body:", resp.Body)

	doneChan <- struct{}{}
}

func generateSolverOperation(userOperation *operation.UserOperation, ethClient *ethclient.Client) *operation.SolverOperation {
	//generates a solver operation specific to the user operation generated by `generateDemoValidUserOp`

	atlasVerificationContract, err := atlasVerification.NewAtlasVerification(conf.Contracts.AtlasVerification, ethClient)
	if err != nil {
		panic(err)
	}

	atlasDomainSeparator, err := atlasVerificationContract.GetDomainSeparator(nil)
	if err != nil {
		panic(err)
	}

	privateKeyBytes, err := hex.DecodeString("e3818ba09a3106ecbf80cafed7510df6f81e6dfab3cfa240680ec9389ba66206")
	if err != nil {
		panic(err)
	}

	privateKey, err := crypto.ToECDSA(privateKeyBytes)
	if err != nil {
		panic(err)
	}

	publicKey := crypto.PubkeyToAddress(privateKey.PublicKey)

	userOpHash, _ := userOperation.Hash()

	solverOp := &operation.SolverOperation{
		From:         publicKey,
		To:           conf.Contracts.Atlas,
		Value:        big.NewInt(0),
		Gas:          big.NewInt(100000),
		MaxFeePerGas: big.NewInt(1000000000),
		Deadline:     userOperation.Deadline,
		Solver:       common.HexToAddress("0x91eb17bBeD5557E70A7c0D9c69C413aecA820765"),
		Control:      common.HexToAddress("0x2F98675731D0e659E716d890330901a8A2355813"),
		UserOpHash:   userOpHash,
		BidToken:     common.HexToAddress("0x00"),
		BidAmount:    big.NewInt(1000_000_000),
		Data:         []byte(""),
		Signature:    nil,
	}

	proofHash, err := solverOp.ProofHash()
	if err != nil {
		panic(err)
	}

	payload := crypto.Keccak256Hash([]byte("\x19\x01"), atlasDomainSeparator[:], proofHash[:])
	signature, err := crypto.Sign(payload[:], privateKey)
	if err != nil {
		panic(err)
	}

	solverOp.Signature = signature

	if err := solverOp.Validate(userOperation, conf.Contracts.Atlas, atlasDomainSeparator, nil); err != nil {
		panic("user operation is invalid " + err.Message)
	}

	return solverOp
}
