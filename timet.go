package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

var (
	secondsLayouts = []string{
		time.RFC3339,
		"2006-01-02T15:04:05",        // RFC3339 without zone
	}
)

func main() {
	progName := os.Args[0]
	if len(os.Args) > 2 || (len(os.Args) == 2 && os.Args[1] == "-h") {
		fmt.Println("Allowed: 2016-12-18T18:28:31.762710887Z")
		fmt.Println("         2016-12-18T18:28:31Z")
		fmt.Println("")
		fmt.Printf("Usage %v time\n", progName)
		os.Exit(2)
	}

	if len(os.Args) == 1 {
		fmt.Println(time.Now().Unix())
		os.Exit(0)
	}

	timeStr := os.Args[1]

	if val, err := strconv.ParseUint(timeStr, 0, 64); err == nil {
		for val > 9999999999 {
			val /= 1000
		}

		tm := time.Unix(int64(val), 0)
		gtm := tm.In(time.UTC)

		fmt.Printf("Local: %-40v %v\n", tm.Format(time.UnixDate), tm.Format(time.RFC3339))
		fmt.Printf("UTC  : %-40v %v\n", gtm.Format(time.UnixDate), gtm.Format(time.RFC3339))
		os.Exit(0)
	}

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
