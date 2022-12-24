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

// func printFileSize(level int, fullPath string) string {
// 	branch := strings.Repeat("-", level-1)
// 	fileInfo, _ := os.Stat(fullPath)

// 	fileName := fileInfo.Name()
// 	fileName = hightlightTSX(fileName, TSX)

// 	if len(branch) == 0 {
// 		return fmt.Sprintf("%s (%d B)", fileName, fileInfo.Size())
// 	} else {
// 		return fmt.Sprintf("%s %s (%d B)", branch, fileName, fileInfo.Size())
// 	}
// }

func ReadFileLines(fullPath string, fileName string, level int) []string {
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
			libFileName := getLibFromLine(line)
			dependencies = append(dependencies, libFileName)
		}
	}

	return dependencies
}

type FileInfo struct {
	Name         string
	IsDir        bool
	Path         string
	Dependencies []string
}

func MapFiles(sourceDir string) map[string]FileInfo {
	Files := map[string]FileInfo{}

	err := filepath.WalkDir(sourceDir, func(path string, file fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		entryFound := FileInfo{Name: file.Name(), IsDir: file.IsDir(), Path: path, Dependencies: []string{}}
		Files[entryFound.Name] = entryFound

		return nil
	})

	if err != nil {
		fmt.Println(err)
	}

	return Files
}
