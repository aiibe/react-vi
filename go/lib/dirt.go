package dirt

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/samber/lo"
)

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
	tsxFiles := lo.Filter(files, func(x fs.FileInfo, index int) bool {
		return filepath.Ext(x.Name()) == ".tsx"
	})

	// Print tsx files
	for _, file := range tsxFiles {
		status := fmt.Sprintf("%s - %d Bytes", file.Name(), file.Size())
		fmt.Println(status)
	}
}
