package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/FastLane-Labs/atlas-operations-relay/core"
	"github.com/FastLane-Labs/atlas-operations-relay/log"
	"github.com/FastLane-Labs/atlas-operations-relay/operation"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gorilla/websocket"
)

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

func runSolver(doneChan chan struct{}, sendMsgOnWs bool) {
	//solver ws connection
	conn, solverResp := getSolverWsConnection()

	if solverResp.StatusCode != 101 {
		log.Error("Expected status code 101, got", solverResp.StatusCode)
		return
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
		log.Error("failed to marshal request:", err)
		return
	}
	if err := conn.WriteMessage(websocket.TextMessage, msg); err != nil {
		log.Error("failed to send request:", err)
		return
	}

	//dont proceed until subscribled
	<-subscribedChan

	log.Info("solver subscribed to", "topic", req.Params.Topic)

	//wait for userOp to be received
	userOpBroadcastBytes := <-userOpReceivedChan

	broadcast := &core.Broadcast{}

	err = json.Unmarshal(userOpBroadcastBytes, broadcast)
	if err != nil {
		log.Error("failed to unmarshal userOpBroadcastBytes:", err)
		return
	}

	userOp := broadcast.Data.UserOperation
	userOpHash, _ := userOp.Hash()
	log.Info("solver received userOp", "userOpHash", userOpHash.Hex())

	ee := ExecutionEnvironment(userOp.From, userOp.Control)
	solverOp := SolveUserOperation(userOp, ee)

	if sendMsgOnWs {
		if err := sendSolverOpWs(solverOp, conn); err != nil {
			log.Error("failed to send solverOp on ws:", err)
		}
	} else {
		if err := sendSolverOpHttp(solverOp); err != nil {
			log.Error("failed to send solverOp on http:", err)
		}
	}
	log.Info("solver sent solverOp", "userOpHash", solverOp.UserOpHash.Hex())
	
	doneChan <- struct{}{}
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
	solverOpJSON, err := json.Marshal(solverOp)
	if err != nil {
		return err
	}

	resp, err := http.Post("http://localhost:8080/solverOperation", "application/json", bytes.NewReader(solverOpJSON))
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Expected status code 200, got %d", resp.StatusCode)
	}

	return nil
}

func sendSolverOpWs(solverOp *operation.SolverOperation, conn *websocket.Conn) error {
	req := core.Request{
		Id:     solverOp.UserOpHash.Hex(),
		Method: "submitSolverOperation",
		Params: &core.RequestParams{
			SolverOperation: solverOp,
		},
	}

	reqJSON, err := json.Marshal(req)
	if err != nil {
		return err
	}

	if err := conn.WriteMessage(websocket.TextMessage, reqJSON); err != nil {
		return err
	}

	return nil
}

/**
func TestSolverOp(t *testing.T) {
	solverDoneChan := make(chan struct{})

	go runSolver(solverDoneChan, false)

	userOp, err := sendUserRequest()
	if err != nil {
		t.Fatal(err)
	}

	<-solverDoneChan

	//user requests for solver solutions
	userOpHash, _ := userOp.Hash()
	solverOps, err := retreiveSolverOps(userOpHash, true)
	if err != nil {
		t.Fatal(err)
	}

	if len(solverOps) == 0 {
		t.Fatal("expected at least one solver operation")
	}
}
**/
