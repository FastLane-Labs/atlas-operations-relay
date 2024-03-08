package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/common"
)

func TestUserOpFlow(t *testing.T) {
	userOp := NewDemoUserOperation()

	userOpHash, relayErr := userOp.Hash()
	if relayErr != nil {
		panic(relayErr)
	}

	userOpJSON, err := json.Marshal(userOp)
	if err != nil {
		panic(err)
	}

	resp, err := http.Post("http://localhost:8080/userOperation", "application/json", bytes.NewReader(userOpJSON))
	if err != nil {
		fmt.Println("Error sending HTTP request:", err)
		return
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, resp.StatusCode)
	}

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	userOpHashReceived := common.HexToHash(strings.Trim(string(bodyBytes), "\""))

	if userOpHashReceived != userOpHash {
		t.Errorf("Expected user operation hash %s, got %s", userOpHash.Hex(), userOpHashReceived.Hex())
	}
}
