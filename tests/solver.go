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
	"github.com/FastLane-Labs/atlas-operations-relay/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gorilla/websocket"
)

type solveUserOpFunc func(*operation.UserOperationPartialRaw, common.Address) *operation.SolverOperation

func runSolver(sendMsgOnWs bool,
	solveUserOpFunc solveUserOpFunc,
	doneChan chan struct{}) {

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
	userOperationPartialReceiveChan := make(chan []byte)

	//start listening on connection
	go func() {
		for {
			select {
			case <-doneChan:
				return
			default:
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
					userOperationPartialReceiveChan <- message
				}
			}
		}
	}()

	//subscribe to newUserOperations topic
	req := core.Request{
		Id:     subscriptionId,
		Method: "subscribe",
		Params: &core.RequestParams{
			Topic: "newUserOperations",
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
	solverInpBroadcastBytes := <-userOperationPartialReceiveChan

	broadcast := &core.Broadcast{}

	err = json.Unmarshal(solverInpBroadcastBytes, broadcast)
	if err != nil {
		log.Error("failed to unmarshal userOpBroadcastBytes:", err)
		return
	}

	userOpPartial := broadcast.Data.UserOperationPartial
	isHinted := len(userOpPartial.Hints) > 0
	isDirect := userOpPartial.Data != nil || userOpPartial.From != common.Address{} || userOpPartial.Value != nil

	if isHinted && isDirect {
		log.Error("userOpPartial is both hinted and direct")
		return
	} else if !isHinted && !isDirect {
		log.Error("userOpPartial is neither hinted nor direct")
	}

	userOpHash := userOpPartial.UserOpHash
	log.Info("solver received userOperationPartial", "userOpHash", userOpHash.Hex())

	ee := executionEnvironment(userOpPartial.From, userOpPartial.Control)
	solverOp := solveUserOpFunc(userOpPartial, ee)

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
			log.Error("failed to send solverOp on http", "err", err)
		}
	}
	log.Info("solver sent solverOp", "userOpHash", solverOp.UserOpHash.Hex())
}

func solveUserOperation(userOperationPartialRaw *operation.UserOperationPartialRaw, executionEnvironment common.Address) *operation.SolverOperation {
	if userOperationPartialRaw.Data == nil {
		panic("need userOperationPartialRaw.Data for this test")
	}

	swapIntent, err := swapIntentAbiDecode(userOperationPartialRaw.Data)
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
		Gas:          big.NewInt(int64(solverGasLimit)),
		MaxFeePerGas: big.NewInt(0).Add(userOperationPartialRaw.MaxFeePerGas.ToInt(), big.NewInt(1e9)),
		Deadline:     userOperationPartialRaw.Deadline.ToInt(),
		Solver:       simpleRfqSolver,
		Control:      userOperationPartialRaw.Control,
		UserOpHash:   userOperationPartialRaw.UserOpHash,
		BidToken:     common.Address{},
		BidAmount:    big.NewInt(1e10),
		Data:         data,
		Signature:    nil,
	}

	solverOpHash, relayErr := solverOp.Hash(&conf.Relay.Eip712.Domain)
	if relayErr != nil {
		panic(relayErr)
	}

	solverOp.Signature, _ = utils.SignMessage(solverOpHash.Bytes(), solverPk)

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

	return append(method.ID, data...), nil
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
