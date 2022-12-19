package dirt

import (
	"bufio"
	"fmt"
	"io/fs"
	"os"
	"strings"

	"github.com/samber/lo"
)

const (
	TSX = ".tsx"
	TS  = ".ts"
)

func sayFileSize(level int, fullPath string) string {
	branch := strings.Repeat("-", level-1)
	fileInfo, _ := os.Stat(fullPath)

	fileName := fileInfo.Name()
	fileName = hightlightTSX(fileName, TSX)

	if len(branch) == 0 {
		return fmt.Sprintf("%s (%d B)", fileName, fileInfo.Size())
	} else {
		return fmt.Sprintf("%s %s (%d B)", branch, fileName, fileInfo.Size())
	}
}

func readFileLines(sourcePath string, fileName string, level int) {

	// Get full path
	fullPath := sourcePath + "/" + fileName

	// Open file
	file, err := os.Open(fullPath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Announcer
	say := sayFileSize(level, fullPath)
	fmt.Println(say)

	// Scan line by line
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		hasImportLine := lineHasLib(line, TSX)

		// Filter import lines
		if hasImportLine {
			libName, libPath := getLib(line, sourcePath)

			// Show only local lib
			if strings.HasPrefix(libPath, sourcePath) {
				if _, err := os.Stat(libPath); os.IsNotExist(err) {
					// fmt.Println("Not found", libName)

				} else {
					// fmt.Println("Found", libName, "in", fileName)
					readFileLines(sourcePath, libName, level+1)
				}
			}
		}
	}

}

func Scan(sourceDir string) {

	// Read dir
	dir, err := os.Open(sourceDir)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer dir.Close()

	// Read files in dir
	files, err := dir.Readdir(-1)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Get index.tsx
	entryFiles := lo.Filter(files, func(file fs.FileInfo, index int) bool {
		return file.Name() == "index.tsx"
	})

	// Read lines for each entry file
	lo.ForEach(entryFiles, func(file fs.FileInfo, index int) {
		fileInfo := fmt.Sprintf("%s - %d Bytes", file.Name(), file.Size())
		fmt.Println(fileInfo)
		fmt.Println("-------------------")

		readFileLines(sourceDir, file.Name(), 1)

		fmt.Println("")
	})
}
