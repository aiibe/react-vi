package main

import (
	"encoding/json"
	"fmt"
	"log"
	"path/filepath"
	dirt "souksyp/react-vi/lib"
	"strings"
	"time"

	"github.com/fatih/color"
	"github.com/samber/lo"
)

func main() {
	start := time.Now()

	root := "../src"

	// Get files map
	mapFiles := dirt.MapFiles(root)

	files := lo.Values(mapFiles)

	// Get only .tsx files only
	tsx := lo.Filter(files, func(file dirt.FileInfo, index int) bool {
		return !file.IsDir && filepath.Ext(file.Path) == ".tsx" && !strings.Contains(file.Name, "test")
	})

	//
	green := color.New(color.FgGreen).SprintFunc()

	// Scan each .tsx files
	lo.ForEach(tsx, func(file dirt.FileInfo, index int) {
		libNames := dirt.ReadFileLines(file.Path, file.Name, 1)
		lo.ForEach(libNames, func(name string, i int) {
			isTSX := filepath.Ext(name) == dirt.TSX
			if isTSX {
				entry, ok := mapFiles[name]
				if ok {
					currentFile := mapFiles[file.Name]
					currentFile.Dependencies = append(currentFile.Dependencies, entry.Name)
					mapFiles[file.Name] = currentFile
				}
			}
		})

		fmt.Println(mapFiles[file.Name].Name, green(len(mapFiles[file.Name].Dependencies)))

		lo.ForEach(mapFiles[file.Name].Dependencies, func(dp string, _ int) {
			fmt.Println("-", green(dp))
		})
	})

	elapsed := time.Since(start)
	fmt.Println("-------------")
	fmt.Println(elapsed)
	fmt.Println("-------------")

	//MarshalIndent
	jason, err := json.MarshalIndent(mapFiles, "", "  ")
	if err != nil {
		log.Fatalf(err.Error())
	}

	fmt.Print(string(jason))
	// fmt.Printf("%#v\n", mapFiles)
}
