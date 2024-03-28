package main

import (
	"fmt"
	"os"

	"github.com/FastLane-Labs/atlas-operations-relay/core"
)

var (
	Version = ""
)

func main() {
	if len(os.Args) > 1 && os.Args[1] == "version" {
		if Version == ""{
			os.Exit(0)
		}
		fmt.Printf("Version: %s\n", Version)
		os.Exit(0)
	}

	core.StartRelay(nil, nil)
}
