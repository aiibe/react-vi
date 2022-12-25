package store

import (
	"path/filepath"
	"souksyp/react-vi/dirt"
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
func ScanNodesDependencies(rootPath string) {

	for _, file := range NodesMap {
		fileExt := filepath.Ext(file.Name)
		isTSFile := (fileExt == dirt.TSX) || (fileExt == dirt.TS)

		if isTSFile {
			depNames := dirt.GetDependencies(rootPath, file.Path, file.Name)

			for _, depName := range depNames {
				updateDependencies(file.Id, depName)
			}
		}
	}
}

// Update a file dependencies
func updateDependencies(targetFileId string, dependencyName string) {
	_, ok := NodesMap[dependencyName]
	if ok {
		currentFile := NodesMap[targetFileId]
		currentFile.Dependencies = append(currentFile.Dependencies, dependencyName)
		NodesMap[targetFileId] = currentFile
	}
}
