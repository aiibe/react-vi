package public

import (
	"embed"

	"github.com/labstack/echo/v4"
)

//go:embed all:build
var Build embed.FS

var FrontApp = echo.MustSubFS(Build, "build")
