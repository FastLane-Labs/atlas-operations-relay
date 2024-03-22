package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/big"
	"math/rand"
	"net/http"
	"net/url"

	"github.com/FastLane-Labs/atlas-operations-relay/contract"
	"github.com/FastLane-Labs/atlas-operations-relay/core"
	"github.com/FastLane-Labs/atlas-operations-relay/log"
	"github.com/FastLane-Labs/atlas-operations-relay/operation"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gorilla/websocket"
)

func runSolver(sendMsgOnWs bool) {
	//solver ws connection
	conn, solverResp := getSolverWsConnection()

	if solverResp.StatusCode != 101 {
		log.Error("Expected status code 101, got", solverResp.StatusCode)
		return
	}

	//subscription id
	subscriptionId := generateRandomString(10)

	//track what's being received
	responseChan := make(chan string)
	solverInputReceiveChan := make(chan []byte)

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
				responseChan <- response.Id
			}
			if broadcast.Topic != "" {
				solverInputReceiveChan <- message
			}
		}
	}()

	//subscribe to newSolverInputs topic
	req := core.Request{
		Id:     subscriptionId,
		Method: "subscribe",
		Params: &core.RequestParams{
			Topic: "newSolverInputs",
		},
	}
	msg, err := json.Marshal(req)
	if err != nil {
		log.Error("failed to marshal request:", err)
		return
	}
	if err := conn.WriteMessage(websocket.TextMessage, msg); err != nil {
		log.Error("failed to send request:", err)
		return
	}

	//dont proceed until subscribled
	latestResponseId := <-responseChan
	if latestResponseId != subscriptionId {
		log.Error("expected subscription id", "expected", subscriptionId, "got", latestResponseId)
	}

	log.Info("solver subscribed to", "topic", req.Params.Topic)

	//wait for userOp to be received
	solverInpBroadcastBytes := <-solverInputReceiveChan

	broadcast := &core.Broadcast{}

	err = json.Unmarshal(solverInpBroadcastBytes, broadcast)
	if err != nil {
		log.Error("failed to unmarshal userOpBroadcastBytes:", err)
		return
	}

	solverInp := broadcast.Data.SolverInput
	if err := solverInp.Validate(); err != nil {
		log.Error("solverInp validation failed:", err)
		return
	}

	userOpHash := solverInp.UserOpHash
	log.Info("solver received userOp", "userOpHash", userOpHash.Hex())

	ee := executionEnvironment(solverInp.From, solverInp.Control)
	solverOp := solveUserOperation(solverInp, ee)

	if sendMsgOnWs {
		requestId, err := sendSolverOpWs(solverOp, conn)
		if err != nil {
			log.Error("failed to send solverOp on ws:", err)
		}
		responseId := <-responseChan
		if responseId != requestId {
			log.Error("expected response id", "expected", requestId, "got", responseId)
		}

	} else {
		if err := sendSolverOpHttp(solverOp); err != nil {
			log.Error("failed to send solverOp on http:", err)
		}
	}
	log.Info("solver sent solverOp", "userOpHash", solverOp.UserOpHash.Hex())
}

func solveUserOperation(solverInput *operation.SolverInput, executionEnvironment common.Address) *operation.SolverOperation {
	if solverInput.Data == nil {
		panic("need solverInput.Data for this test")
	}

	swapIntent, err := swapIntentAbiDecode(solverInput.Data)
	if err != nil {
		panic(err)
	}

	data, err := solverData(swapIntent, executionEnvironment)
	if err != nil {
		panic(err)
	}

	solverOp := &operation.SolverOperation{
		From:         solverEoa,
		To:           conf.Contracts.Atlas,
		Value:        big.NewInt(0),
		Gas:          big.NewInt(100000),
		MaxFeePerGas: big.NewInt(0).Add(solverInput.MaxFeePerGas.ToInt(), big.NewInt(1e9)),
		Deadline:     solverInput.Deadline.ToInt(),
		Solver:       simpleRfqSolver,
		Control:      solverInput.Control,
		UserOpHash:   solverInput.UserOpHash,
		BidToken:     common.Address{},
		BidAmount:    big.NewInt(1e13),
		Data:         data,
		Signature:    nil,
	}

	proofHash, err := solverOp.ProofHash()
	if err != nil {
		panic(err)
	}

	solverOp.Signature = signEip712(atlasDomainSeparator, proofHash, solverPk)

	if err := solverOp.Validate(solverInput, conf.Contracts.Atlas, atlasDomainSeparator, conf.Relay.Gas.MaxPerSolverOperation); err != nil {
		panic(err)
	}

	return solverOp
}

func solverData(swapIntent *SwapIntent, executionEnvironment common.Address) ([]byte, error) {
	method, exists := contract.SimpleRfqSolverAbi.Methods["fulfillRFQ"]
	if !exists {
		return nil, fmt.Errorf("method signature not found")
	}

	data, err := method.Inputs.Pack(swapIntent, executionEnvironment)
	if err != nil {
		return nil, fmt.Errorf("failed to pack inputs: %v", err)
	}

	return append(common.Hex2Bytes(solverFulfillFuncSelector), data...), nil
}

func getSolverWsConnection() (*websocket.Conn, *http.Response) {
	u := url.URL{Scheme: "ws", Host: "localhost:8080", Path: "/ws/solver"}

	conn, solverResp, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		panic(err)
	}
	return conn, solverResp
}

func sendSolverOpHttp(solverOp *operation.SolverOperation) error {
	solverOpJSON, err := json.Marshal(solverOp.EncodeToRaw())
	if err != nil {
		return err
	}

	resp, err := http.Post("http://localhost:8080/solverOperation", "application/json", bytes.NewReader(solverOpJSON))
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("expected status code 200, got %d", resp.StatusCode)
	}

	return nil
}

func sendSolverOpWs(solverOp *operation.SolverOperation, conn *websocket.Conn) (string, error) {
	request_id := generateRandomString(10)

	req := core.Request{
		Id:     request_id,
		Method: "submitSolverOperation",
		Params: &core.RequestParams{
			SolverOperation: solverOp.EncodeToRaw(),
		},
	}

	reqJSON, err := json.Marshal(req)
	if err != nil {
		return "", err
	}

	if err := conn.WriteMessage(websocket.TextMessage, reqJSON); err != nil {
		return "", err
	}

	return request_id, nil
}

func generateRandomString(n int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	var output []byte
	for i := 0; i < n; i++ {
		randomIndex := rand.Intn(len(charset))
		output = append(output, charset[randomIndex])
	}

	return string(output)
}
