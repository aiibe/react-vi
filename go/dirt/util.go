package dirt

import (
	"os"
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
	toRemove := []string{`'`, `"`, `;`, `./`}
	libPath := stripAll(lastBlock, toRemove)

	// Possibly an .tsx  import
	if filepath.Ext((libPath)) != ".svg" && filepath.Ext((libPath)) != ".css" {
		libPath3 := libPath + TSX
		libPath2 := strings.Replace(path, name, libPath3, 1)

		_, err := os.Stat(libPath2)
		if err != nil {
			if os.IsNotExist(err) {
				libPath4 := libPath + "/index.ts"
				libPath5 := strings.Replace(path, name, libPath4, 1)

				_, err := os.Stat(libPath5)
				if err == nil {
					return libPath5
				}
			}
		}
		return libPath2
	}

	return strings.Replace(path, name, libPath, 1)
}
