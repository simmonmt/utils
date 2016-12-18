package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

var (
	secondsLayouts = []string{time.RFC3339}
)

func main() {
	progName := os.Args[0]
	if len(os.Args) != 2 || os.Args[1] == "-h" {
		fmt.Println("Allowed: 2016-12-18T18:28:31.762710887Z")
		fmt.Println("         2016-12-18T18:28:31Z")
		fmt.Println("")
		fmt.Printf("Usage %v time\n", progName)
		os.Exit(2)
	}

	timeStr := os.Args[1]

	if strings.Contains(timeStr, ".") {
		if tm, err := time.Parse(time.RFC3339Nano, timeStr); err == nil {
			fmt.Println(tm.UnixNano())
			os.Exit(0)
		}
	}

	for _, layout := range secondsLayouts {
		if tm, err := time.Parse(layout, timeStr); err == nil {
			fmt.Println(tm.Unix())
			os.Exit(0)
		}
	}

	log.Fatalf("Failed to parse %v", timeStr)
}
