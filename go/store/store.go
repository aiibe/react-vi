package store

import (
	"souksyp/react-vi/dirt"
)

// Declare files map
var NodesMap map[string]dirt.FileNode

// Update a file dependencies
func UpdateDependencies(targetFileId string, dependencyName string) {
	_, ok := NodesMap[dependencyName]
	if ok {
		currentFile := NodesMap[targetFileId]
		currentFile.Dependencies = append(currentFile.Dependencies, dependencyName)
		NodesMap[targetFileId] = currentFile
	}
}
