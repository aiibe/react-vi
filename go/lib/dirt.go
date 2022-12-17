package dirt

import (
	"bufio"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/samber/lo"
)

func ReadLines(path string) {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
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
		status := fmt.Sprintf("%s - %d Bytes", file.Name(), file.Size())
		fmt.Println(status)

		filePath := sourceDir + "/" + file.Name()
		ReadLines(filePath)
	})
}
