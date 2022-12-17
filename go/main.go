package main

import (
	dirt "souksyp/react-vi/lib"
)

// Run main app
func main() {

	// Source dir
	sourcePath := "../src"

	// Read src files
	dirt.Scan(sourcePath)
}
