package main

import (
	"fmt"
	"io/ioutil"
	dirt "souksyp/react-vi/lib"

	"github.com/ije/esbuild-internal/config"
	"github.com/ije/esbuild-internal/js_parser"
	"github.com/ije/esbuild-internal/logger"
)

func main() {
	filename := "../src/index.tsx"

	// Read file
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error")
	}

	sourceFile := logger.Source{
		Index:          0,
		KeyPath:        logger.Path{Text: filename},
		PrettyPath:     filename,
		Contents:       string(data),
		IdentifierName: filename,
	}

	var logMap dirt.LogMap
	log := dirt.NewLogMap(logger.OutputOptions{LogLevel: logger.LevelDebug}, &logMap)

	// AST
	options := config.Options{Mode: config.ModeBundle}
	options.JSX = config.JSXOptions{
		Parse: true,
	}
	ast, _ := js_parser.Parse(log, sourceFile, js_parser.OptionsFromConfig(&options))

	fmt.Println(ast.ImportRecords)
}
