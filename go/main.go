package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"souksyp/react-vi/dirt"
	"souksyp/react-vi/store"

	"time"
)

// Start
func main() {
	start := time.Now()

	root := flag.String("d", "./src", "a string")
	flag.Parse()

	// Store all files/directories within root directory
	store.NodesMap = dirt.Scan(*root)
	store.ScanNodesDependencies()

	// Marshal Indent
	jason, err := json.MarshalIndent(store.NodesMap, "", "  ")
	if err != nil {
		log.Fatalf(err.Error())
	}
	fmt.Print(string(jason))

	fmt.Println("-------------")
	elapsed := time.Since(start)
	fmt.Println(elapsed, *root)
}
