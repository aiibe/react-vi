package dirt

import (
	"bufio"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

const (
	TSX = ".tsx"
	TS  = ".ts"
)

type FileNode struct {
	Name         string
	IsDir        bool
	Path         string
	Dependencies []string
}

func GetDependencies(fullPath string, fileName string) []string {
	var dependencies []string

	// Open file
	file, err := os.Open(fullPath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Scan line by line
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		hasImportLine := lineHasLib(line, TSX)

		// Filter import lines
		if hasImportLine {
			libFileName := getLibNameFromLine(line)
			dependencies = append(dependencies, libFileName)
		}
	}

	return dependencies
}

func Scan(sourceDir string) map[string]FileNode {
	Files := map[string]FileNode{}

	err := filepath.WalkDir(sourceDir, func(path string, file fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		entryFound := FileNode{Name: file.Name(), IsDir: file.IsDir(), Path: path, Dependencies: []string{}}
		Files[entryFound.Name] = entryFound

		return nil
	})

	if err != nil {
		fmt.Println(err)
	}

	return Files
}
