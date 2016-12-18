package main

import (
	"fmt"
	"log"
	"math/big"
	"os"
	"path"
	"strings"
)

func dx(numStr string) {
	var num big.Int
	if _, ok := num.SetString(numStr, 10); !ok {
		log.Fatalf("failed to parse %v as decimal", numStr)
	}

	fmt.Printf("0x%v\n", num.Text(16))
}

func xd(numStr string) {
	numStr = strings.TrimPrefix(strings.ToLower(numStr), "0x")

	var num big.Int
	if _, ok := num.SetString(numStr, 16); !ok {
		log.Fatalf("failed to parse %v as hex", numStr)
	}

	fmt.Println(num.Text(10))
}

func main() {
	progName := path.Base(os.Args[0])

	if len(os.Args) != 2 {
		log.Fatalf("Usage: %v num", progName)
	}
	numStr := os.Args[1]

	switch progName {
	case "dx":
		dx(numStr)
		break
	case "xd":
		xd(numStr)
		break
	default:
		log.Fatalf("Unknown mode %v", progName)
	}
}
