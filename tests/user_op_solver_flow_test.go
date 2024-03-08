package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"testing"
	"time"

	"github.com/FastLane-Labs/atlas-operations-relay/core"
	"github.com/FastLane-Labs/atlas-operations-relay/log"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gorilla/websocket"
)

func TestUserOpToSolverFLow(t *testing.T) {

	//solver ws connection
	u := url.URL{Scheme: "ws", Host: "localhost:8080", Path: "/ws/solver"}
	fmt.Printf("Connecting to %s\n", u.String())

	conn, solverResp, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	if solverResp.StatusCode != 101 {
		t.Errorf("Expected status code %d, got %d", 101, solverResp.StatusCode)
	}

	//track what's being received
	subscribedChan := make(chan struct{})
	userOpReceivedChan := make(chan struct{})

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

			errResp := json.Unmarshal(message, response)
			errBroadcast := json.Unmarshal(message, broadcast)

			if errResp != nil && errBroadcast != nil {
				panic("cannot handle msg " + string(message))
			}
			if errResp == nil {
				subscribedChan <- struct{}{}
			}
			if errBroadcast == nil {
				userOpReceivedChan <- struct{}{}
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
		t.Error("not subscribed to newUserOperations topic")
	}

	//eth client connection
	ethClient, err := ethclient.Dial(conf.Network.RpcUrl)
	if err != nil {
		log.Error("failed to connect to the Ethereum client", "err", err)
		return
	}

	userOp := generateDemoValidUserOp(ethClient, conf)
	userOpJSON, err := json.Marshal(userOp)
	if err != nil {
		panic(err)
	}

	_, err = http.Post("http://localhost:8080/userOperation", "application/json", bytes.NewReader(userOpJSON))
	if err != nil {
		fmt.Println("Error sending HTTP request:", err)
		return
	}

	if !waitOnChanFor(userOpReceivedChan, 1*time.Second) {
		t.Error("user operation not received by solver")
	}
}

func waitOnChanFor(ch chan struct{}, t time.Duration) bool {
	select {
	case <-ch:
		return true
	case <-time.After(t):
		fmt.Println("timout")
		return false
	}
}
