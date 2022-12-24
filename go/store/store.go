package store

import (
	"path/filepath"
	dirt "souksyp/react-vi/dirt"
	"strings"
)

// Declare files map
var NodesMap map[string]dirt.FileNode

// Get only .tsx files
func FilterTSX() []dirt.FileNode {
	var files []dirt.FileNode

	for _, file := range NodesMap {
		if !file.IsDir && filepath.Ext(file.Path) == ".tsx" && !strings.Contains(file.Name, "test") {
			files = append(files, file)
		}
	}

	return files
}

// Scan dependencies
func ScanNodesDependencies() {
	tsxFiles := FilterTSX()

	for _, file := range tsxFiles {
		depNames := dirt.GetDependencies(file.Path, file.Name)
		for _, depName := range depNames {
			isTSX := filepath.Ext(depName) == dirt.TSX
			if isTSX {
				updateDependencies(file.Name, depName)
			}
		}
	}
}

// Update a file dependencies
func updateDependencies(targetFileName string, dependencyName string) {
	_, ok := NodesMap[dependencyName]
	if ok {
		currentFile := NodesMap[targetFileName]
		currentFile.Dependencies = append(currentFile.Dependencies, dependencyName)
		NodesMap[targetFileName] = currentFile
	}
}
