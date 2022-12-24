package main

import (
	"encoding/json"
	"fmt"
	"log"
	"souksyp/react-vi/dirt"
	"souksyp/react-vi/store"

	"time"
)

// Start
func main() {
	start := time.Now()

	// Store all files/directories within root directory
	const root = "../src"
	store.NodesMap = dirt.Scan(root)
	store.ScanNodesDependencies()

	// Marshal Indent
	jason, err := json.MarshalIndent(store.NodesMap, "", "  ")
	if err != nil {
		log.Fatalf(err.Error())
	}
	fmt.Print(string(jason))

	fmt.Println("-------------")
	elapsed := time.Since(start)
	fmt.Println(elapsed)
}
