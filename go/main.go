package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"souksyp/react-vi/dirt"
	"souksyp/react-vi/store"
	"strings"

	"github.com/ije/esbuild-internal/config"
	"github.com/ije/esbuild-internal/js_parser"
	"github.com/ije/esbuild-internal/logger"

	"time"
)

// Start
func main() {
	//
	buildStaticNodes()

	// Open browser
	// dirt.OpenBrowser("http://localhost:18881/")

	// Start server
	// server.Start()

}

func buildStaticNodes() {
	start := time.Now()

	// Accept root directory -d flag
	root := flag.String("d", "./src", "a string")
	flag.Parse()

	// Scan and store files/directories within root directory
	store.NodesMap = dirt.Scan(*root)

	// Scan dependencies for each file
	firstFilename := "../src/index.tsx"

	fmt.Println(firstFilename)

	data, err := os.ReadFile(firstFilename)
	if err != nil {
		fmt.Println("Error read file", firstFilename)
	}

	var logMap dirt.LogMap

	log := dirt.CreateLog(logger.OutputOptions{LogLevel: logger.LevelDebug}, &logMap)

	sourceFile := logger.Source{
		Index:          0,
		KeyPath:        logger.Path{Text: firstFilename},
		PrettyPath:     firstFilename,
		Contents:       string(data),
		IdentifierName: firstFilename,
	}

	ext := filepath.Ext(firstFilename)
	options := config.Options{Mode: config.ModeBundle}
	if ext == ".ts" || ext == ".tsx" {
		options.TS = config.TSOptions{
			Parse: true,
		}
	}

	if ext == ".jsx" || ext == ".tsx" {
		options.JSX = config.JSXOptions{
			Parse: true,
		}
	}

	ast, _ := js_parser.Parse(log, sourceFile, js_parser.OptionsFromConfig(&options))

	fmt.Println(len(ast.ImportRecords))

	localDependencies := []string{}
	for _, record := range ast.ImportRecords {
		if strings.HasPrefix(record.Path.Text, "./") {
			localDependencies = append(localDependencies, record.Path.Text)
			fmt.Println(filepath.Join(*root, filepath.Base(record.Path.Text)))
		}
	}

	fmt.Println(localDependencies)

	fmt.Println("Executed time", time.Since(start))
}
