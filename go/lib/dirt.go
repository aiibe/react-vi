package dirt

import (
	"bufio"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/fatih/color"
	"github.com/samber/lo"
)

const (
	TSX = ".tsx"
)

func cleanString(str string, toReplace []string) string {
	lo.ForEach(toReplace, func(char string, index int) {
		str = strings.ReplaceAll(str, char, "")
	})
	return str
}

func getLib(line string, sourcePath string) (string, string) {
	str := strings.Split(line, " ")
	length := len(str)

	// Access the element at the last index of the array
	last := str[length-1]

	// Clean
	subs := []string{`'`, `"`, `;`}
	name := cleanString(last, subs)

	// Get filename
	str2 := strings.Split(name, "/")
	length2 := len(str2)
	finalName := str2[length2-1]

	// Source of file
	sourceFilePath := strings.ReplaceAll(name, "./", sourcePath+"/")

	filePath := filepath.Ext((name))
	if filePath == "" {
		sourceFilePath += TSX
		finalName += TSX

	}

	return finalName, sourceFilePath
}

func lineHasLib(line string, ext string) bool {
	if ext == TSX {
		return strings.HasPrefix(line, "import")
	}
	return false
}

func hightlight(name string, ext string) string {
	blue := color.New(color.FgBlue).SprintFunc()

	if filepath.Ext(name) == TSX {
		return blue(name)
	}

	return name
}

func sayFileSize(level int, fullPath string) string {
	branch := strings.Repeat("-", level-1)
	fileInfo, _ := os.Stat(fullPath)

	fileName := fileInfo.Name()
	fileName = hightlight(fileName, TSX)

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
