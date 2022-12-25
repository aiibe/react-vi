package main

import (
	"flag"
	"fmt"
	"souksyp/react-vi/dirt"
	"souksyp/react-vi/server"
	"souksyp/react-vi/store"

	"time"
)

// Start
func main() {
	//
	buildStaticNodes()

	// Open browser
	// dirt.OpenBrowser("http://localhost:18881/")

	// Start server
	server.Start()

}

func buildStaticNodes() {
	start := time.Now()

	// Accept root directory -d flag
	root := flag.String("d", "./src", "a string")
	flag.Parse()

	// Scan and store files/directories within root directory
	store.NodesMap = dirt.Scan(*root)

	// Scan for deps inside files
	store.ScanNodesDependencies()

	fmt.Println("Executed time", time.Since(start))
}
