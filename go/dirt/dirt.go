package dirt

import (
	"fmt"
	"io/fs"
	"os/exec"
	"path/filepath"
	"runtime"

	"github.com/ije/esbuild-internal/logger"
)

const (
	TSX = ".tsx"
	TS  = ".ts"
)

type FileNode struct {
	Id           string
	Name         string
	IsDir        bool
	Path         string
	Dependencies []string
}

func Scan(sourceDir string) map[string]FileNode {
	Files := map[string]FileNode{}

	err := filepath.WalkDir(sourceDir, func(path string, file fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		// Populate files map
		entryFound := FileNode{Id: path, Name: file.Name(), IsDir: file.IsDir(), Path: path, Dependencies: []string{}}
		Files[entryFound.Id] = entryFound
		return nil
	})

	if err != nil {
		fmt.Println(err)
	}

	return Files
}

func OpenBrowser(url string) bool {
	var args []string
	switch runtime.GOOS {
	case "darwin":
		args = []string{"open"}
	case "windows":
		args = []string{"cmd", "/c", "start"}
	default:
		args = []string{"xdg-open"}
	}
	cmd := exec.Command(args[0], append(args[1:], url)...)
	return cmd.Start() == nil
}

type LogMap struct {
	Verbose []string
	Debug   []string
	Info    []string
	Err     []string
	Warning []string
}

func CreateLog(options logger.OutputOptions, logMap *LogMap) logger.Log {
	hasErrors := false
	var msgs []logger.Msg

	return logger.Log{
		Level: options.LogLevel,
		AddMsg: func(msg logger.Msg) {
			msgs = append(msgs, msg)

			switch msg.Kind {
			case logger.Verbose:
				if options.LogLevel <= logger.LevelVerbose {
					logMap.Verbose = append(logMap.Verbose, msgString(&msg))
				}

			case logger.Debug:
				if options.LogLevel <= logger.LevelDebug {
					logMap.Debug = append(logMap.Debug, msgString(&msg))
				}

			case logger.Info:
				if options.LogLevel <= logger.LevelInfo {
					logMap.Info = append(logMap.Info, msgString(&msg))
				}

			case logger.Error:
				if options.LogLevel <= logger.LevelError {
					hasErrors = true
					logMap.Err = append(logMap.Err, msgString(&msg))
				}

			case logger.Warning:
				if options.LogLevel <= logger.LevelWarning {
					logMap.Warning = append(logMap.Warning, msgString(&msg))
				}
			}
		},

		HasErrors: func() bool {
			return hasErrors
		},

		AlmostDone: func() {
			// noop
		},

		Done: func() []logger.Msg {
			return msgs
		},
	}
}

func msgString(msg *logger.Msg) string {
	if loc := msg.Data.Location; loc != nil {
		return fmt.Sprintf("%s: %s\n", loc.File, msg.Data.Text)
	}
	return fmt.Sprintf("%s: \n", msg.Data.Text)
}
