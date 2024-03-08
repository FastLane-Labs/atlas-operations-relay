package tests

import (
	"testing"

	"github.com/ethereum/go-ethereum/common"
)

func TestSwapIntentAbiEncode(t *testing.T) {
	t.Parallel()

	swapIntent := NewDemoSwapIntent()
	want := common.FromHex("0x00000000000000000000000000000000000000000000000000000000000000200000000000000000000000007439e9bb6d8a84dd3a23fe621a30f95403f87fb90000000000000000000000000000000000000000000000000000b5e620f480000000000000000000000000007b79995e5f793a07bc00c21412e50ecae098e7f9000000000000000000000000000000000000000000000000000000e8d4a510000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000e00000000000000000000000000000000000000000000000000000000000000000")

	result, err := swapIntent.abiEncode()
	if err != nil {
		t.Errorf("SwapIntent.abiEncode() error = %v", err)
	}

	if string(result) != string(want) {
		t.Errorf("SwapIntent.abiEncode() = %v, want %v", result, want)
	}
}

func TestSwapIntentAbiEncodeWithSelector(t *testing.T) {
	t.Parallel()

	swapIntent := NewDemoSwapIntent()
	want := common.FromHex("0x83a6992a00000000000000000000000000000000000000000000000000000000000000200000000000000000000000007439e9bb6d8a84dd3a23fe621a30f95403f87fb90000000000000000000000000000000000000000000000000000b5e620f480000000000000000000000000007b79995e5f793a07bc00c21412e50ecae098e7f9000000000000000000000000000000000000000000000000000000e8d4a510000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000e00000000000000000000000000000000000000000000000000000000000000000")

	result, err := swapIntent.abiEncodeWithSelector()
	if err != nil {
		t.Errorf("SwapIntent.abiEncodeWithSelector() error = %v", err)
	}

	if string(result) != string(want) {
		t.Errorf("SwapIntent.abiEncodeWithSelector() = %v, want %v", result, want)
	}
}
