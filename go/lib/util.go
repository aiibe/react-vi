package dirt

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/fatih/color"
	"github.com/samber/lo"
)

func stripAll(str string, toReplace []string) string {
	lo.ForEach(toReplace, func(char string, index int) {
		str = strings.ReplaceAll(str, char, "")
	})
	return str
}

func lineHasLib(line string, ext string) bool {
	if ext == TSX {
		return strings.HasPrefix(line, "import")
	}
	return false
}

func hightlightTSX(name string, ext string) string {
	blue := color.New(color.FgBlue).SprintFunc()

	if filepath.Ext(name) == TSX {
		return blue(name)
	}

	return name
}

// Read import line
// Get lib's filename and path to its the source
func getLib(line string, sourcePath string) (string, string) {

	// Split line into workable blocks
	lineBlocks := strings.Split(line, " ")
	lineBlocksLength := len(lineBlocks)

	// Get last block containing relative path to the lib
	lastBlock := lineBlocks[lineBlocksLength-1]

	// Get the lib's path clean
	toRemove := []string{`'`, `"`, `;`}
	libPath := stripAll(lastBlock, toRemove)

	// Get the filename from the path
	pathBlocks := strings.Split(libPath, "/")
	pathBlocksLength := len(pathBlocks)
	libName := pathBlocks[pathBlocksLength-1]

	// Build lib's full path
	sourceFilePath := strings.ReplaceAll(libPath, "./", sourcePath)

	red := color.New(color.FgGreen).SprintFunc()
	fmt.Println(red(libName, "|", sourceFilePath))

	// Possibly an .tsx or .ts import
	fileExt := filepath.Ext((libPath))
	if fileExt == "" {
		// !!! mutation
		sourceFilePath += TSX
		libName += TSX
	}

	return libName, sourceFilePath
}
