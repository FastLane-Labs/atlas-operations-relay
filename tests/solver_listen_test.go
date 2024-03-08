package tests

import (
	"fmt"
	"log"
	"net/url"
	"testing"

	"github.com/gorilla/websocket"
)

func TestSolverListen(t *testing.T) {
	u := url.URL{Scheme: "ws", Host: "localhost:8080", Path: "/ws/solver"}
	fmt.Printf("Connecting to %s\n", u.String())

	c, resp, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer c.Close()

	if resp.StatusCode != 101 {
		t.Errorf("Expected status code %d, got %d", 101, resp.StatusCode)
	}
}
