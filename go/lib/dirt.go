package dirt

import (
	"bufio"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/samber/lo"
)

func lineHasImport(ext string, line string) bool {
	if ext == ".tsx" {
		return strings.HasPrefix(line, "import")
	}
	return false
}

func ReadLines(filePath string) {
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// Filter only lines start with import
	for scanner.Scan() {
		line := scanner.Text()
		isImportLine := lineHasImport(".tsx", line)

		if isImportLine {
			fmt.Println(line)
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

	// Filter .tsx files
	tsxFiles := lo.Filter(files, func(file fs.FileInfo, index int) bool {
		return filepath.Ext(file.Name()) == ".tsx"
	})

	// Read lines
	lo.ForEach(tsxFiles, func(file fs.FileInfo, index int) {
		fileInfo := fmt.Sprintf("%s - %d Bytes", file.Name(), file.Size())
		fmt.Println(fileInfo)
		fmt.Println("-------------------")

		filePath := sourceDir + "/" + file.Name()
		ReadLines(filePath)

		fmt.Println("")
	})
}
