package dirt

import (
	"path/filepath"
	"strings"

	"github.com/samber/lo"
)

func stripAll(str string, toReplace []string) string {
	lo.ForEach(toReplace, func(char string, index int) {
		str = strings.ReplaceAll(str, char, "")
	})
	return str
}

func isImportLine(line string) bool {
	return strings.HasPrefix(line, "import")
}

func isExportLine(line string) bool {
	return strings.HasPrefix(line, "export")
}

// Read import line
// Get lib's filename and path to its the source
func getLibNameFromLine(path string, name string, line string) string {

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

	// Possibly an .tsx  import
	if filepath.Ext((libPath)) != ".svg" || filepath.Ext((libPath)) != ".css" {
		libName += TSX
	}

	libPath = strings.Replace(path, name, libName, 1)

	return libPath
}
