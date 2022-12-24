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

func lineHasLib(line string, ext string) bool {
	if ext == TSX {
		return strings.HasPrefix(line, "import")
	}
	return false
}

// Read import line
// Get lib's filename and path to its the source
func getLibNameFromLine(line string) string {

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
	if filepath.Ext((libPath)) == "" {
		libName += TSX
	}

	return libName
}
