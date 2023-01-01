package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"souksyp/react-vi/dirt"
	"souksyp/react-vi/server"
	"souksyp/react-vi/store"

	"github.com/fatih/color"
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
	server.Start()

}

func buildStaticNodes() {
	start := time.Now()

	// Accept root directory -d flag
	root := flag.String("d", "./src", "a string")
	flag.Parse()

	// Scan and store files/directories within root directory
	store.NodesMap = dirt.Scan(*root)

	// Scan dependencies for each file
	for source, value := range store.NodesMap {
		if !value.IsDir {
			readOne(source)
		}
	}

	// fmt.Println(store.NodesMap)
	fmt.Println("---------------------")
	fmt.Println("Analyze ", len(store.NodesMap), " files in ", time.Since(start))
	color.Green("Success !")
	fmt.Println()
	fmt.Println("Server is running at", color.BlueString("http://localhost:18881"))
	fmt.Println()
}

func readOne(filename string) {

	dirPath, fileName := filepath.Split(filename)
	fmt.Println("Current filename", dirPath, fileName)

	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error read file", filename)
	}

	var logMap dirt.LogMap

	log := dirt.CreateLog(logger.OutputOptions{LogLevel: logger.LevelDebug}, &logMap)

	sourceFile := logger.Source{
		Index:          0,
		KeyPath:        logger.Path{Text: filename},
		PrettyPath:     filename,
		Contents:       string(data),
		IdentifierName: filename,
	}

	ext := filepath.Ext(filename)
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

	fmt.Println("Named imports", len(ast.NamedImports))

	// Get named imports
	var imports []string

	// Link components/index.ts to components dir
	if fileName == "index.ts" || fileName == "index.tsx" {
		source3 := filepath.Clean(dirPath)
		imports = append(imports, source3)
	}

	for _, record := range ast.NamedImports {

		// Current import record
		importRecord := ast.ImportRecords[record.ImportRecordIndex]
		importRecordPath := importRecord.Path.Text

		// Alias is default
		if record.Alias == "default" {

			// Current line is possibly .css, .svg,
			if filepath.Ext(importRecordPath) == ".css" || filepath.Ext(importRecordPath) == ".svg" {
				source1 := filepath.Join(dirPath, filepath.Clean(importRecordPath))
				imports = append(imports, source1)
				// fmt.Println("ASSETS", (importRecordPath), source1)
				continue
			}

			// Handle .component and all....
			source1 := filepath.Join(dirPath, filepath.Clean(importRecordPath)+".tsx")
			source2 := filepath.Join(dirPath, filepath.Clean(importRecordPath)+".ts")
			imports = append(imports, source1)
			imports = append(imports, source2)
			// fmt.Println("ASSETS", (importRecordPatht), source1)
			continue
		}

		// Current line has alias .ts or directory with index.ts
		fmt.Println("-------", importRecordPath)
		source1 := filepath.Join(dirPath, filepath.Base(importRecordPath), "index.ts")
		source2 := filepath.Join(dirPath, filepath.Base(importRecordPath)+".ts")
		source3 := filepath.Join(dirPath, filepath.Base(importRecordPath), "index.tsx")

		source4 := filepath.Join("../src/", filepath.Base(importRecordPath), "index.ts")
		source5 := filepath.Join("../src/", filepath.Base(importRecordPath)+".ts")
		source6 := filepath.Join("../src/", filepath.Base(importRecordPath), "index.tsx")

		imports = append(imports, source1)
		imports = append(imports, source2)
		imports = append(imports, source3)

		imports = append(imports, source4)
		imports = append(imports, source5)
		imports = append(imports, source6)
		// fmt.Println("", record.Alias, source1)
	}

	// Analyze export *
	fmt.Println("Named exports star", len(ast.ExportStarImportRecords))
	for _, index := range ast.ExportStarImportRecords {
		// Current import record
		importRecord := ast.ImportRecords[index]
		importRecordPath := importRecord.Path.Text

		source1 := filepath.Join(dirPath, filepath.Base(importRecordPath)+".tsx")
		source2 := filepath.Join(dirPath, filepath.Base(importRecordPath)+".ts")
		imports = append(imports, source1)
		imports = append(imports, source2)
	}
	fmt.Println("Suggested imports", len(imports), imports)

	// Check existence of suggested imports
	currentFile := store.NodesMap[filename]
	for _, source := range imports {
		fileNode, ok := store.NodesMap[source]
		if ok {
			if !fileNode.IsDir {
				currentFile.Dependencies = append(currentFile.Dependencies, source)
				store.NodesMap[filename] = currentFile
			} else {
				dir := store.NodesMap[source]
				dir.Dependencies = append(store.NodesMap[source].Dependencies, currentFile.Path)
				store.NodesMap[source] = dir
				// currentFile.IsDir = true
			}

		}
	}

	fmt.Println("Imports found", len(currentFile.Dependencies), currentFile.Dependencies)

	fmt.Println()
}
