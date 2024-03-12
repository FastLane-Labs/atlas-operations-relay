package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"testing"

	"github.com/FastLane-Labs/atlas-operations-relay/operation"
	"github.com/ethereum/go-ethereum/common"
)

func TestUserOpFlow(t *testing.T) {
	_, err := sendUserRequest()
	if err != nil {
		t.Fatal(err)
	}
}

func sendUserRequest() (*operation.UserOperation, error) {
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
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Expected status code 200, got %d", resp.StatusCode)
	}

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	userOpHashReceived := common.HexToHash(strings.Trim(string(bodyBytes), "\""))

	if userOpHashReceived != userOpHash {
		return nil, fmt.Errorf("Expected userOpHash %s, got %s", userOpHash, userOpHashReceived)
	}

	return userOp, nil
}

func retrieveAtlasTxHash(userOpHash common.Hash, wait bool) (common.Hash, error) {
	u := url.URL{
		Scheme: "http",
		Host:   "localhost:8080",
		Path:   "/bundleHash",
		RawQuery: url.Values{
			"userOpHash": []string{userOpHash.Hex()},
			"wait":       []string{fmt.Sprintf("%t", wait)},
		}.Encode(),
	}

	resp, err := http.Get(u.String())
	if err != nil {
		return common.Hash{}, err
	}

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return common.Hash{}, fmt.Errorf("Expected status code 200, got %d", resp.StatusCode)
	}

	var txHash common.Hash
	err = json.NewDecoder(resp.Body).Decode(&txHash)
	if err != nil {
		return common.Hash{}, err
	}

	return txHash, nil
}
