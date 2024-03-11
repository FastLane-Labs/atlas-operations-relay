package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"testing"
	"time"

	"github.com/FastLane-Labs/atlas-operations-relay/auction"
	"github.com/FastLane-Labs/atlas-operations-relay/core"
	"github.com/FastLane-Labs/atlas-operations-relay/log"
	"github.com/FastLane-Labs/atlas-operations-relay/operation"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gorilla/websocket"
)

func TestSolverSolves(t *testing.T) {
	solverDoneChan := make(chan struct{})

	go runSolver(solverDoneChan)
	userOp, err := sendUserRequest()
	if err != nil {
		t.Fatal(err)
	}

	<-solverDoneChan

	//wait for auction to end
	time.Sleep(auction.AuctionDuration)

	//user requests for solver solutions
	userOpHash, _ := userOp.Hash()
	solverOps, err := retreiveSolverOps(userOpHash)
	if err != nil {
		t.Fatal(err)
	}

	if len(solverOps) == 0 {
		t.Fatal("expected at least one solver operation")
	}
}

func sendUserRequest() (*operation.UserOperation, error) {
	userOp := NewDemoUserOperation()
	userOpJSON, err := json.Marshal(userOp)
	if err != nil {
		panic(err)
	}

	_, err = http.Post("http://localhost:8080/userOperation", "application/json", bytes.NewReader(userOpJSON))
	if err != nil {
		return nil, err
	}

	return userOp, nil
}

func retreiveSolverOps(userOpHash common.Hash) ([]*operation.SolverOperation, error) {
	u := url.URL{
		Scheme: "http",
		Host:   "localhost:8080",
		Path:   "/solverOperations",
		RawQuery: url.Values{
			"userOpHash": []string{userOpHash.Hex()},
		}.Encode(),
	}

	resp, err := http.Get(u.String())
	if err != nil {
		fmt.Println("Error sending HTTP request:", err)
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

func runSolver(doneChan chan struct{}) {
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

	ee := ExecutionEnvironment(userOp.From, userOp.Control)

	solverOp := SolveUserOperation(userOp, ee)

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

	if resp.StatusCode != http.StatusOK {
		panic("Expected status code 200, got " + fmt.Sprint(resp.StatusCode))
	}

	doneChan <- struct{}{}
}
