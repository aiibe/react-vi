package server

import (
	"net/http"
	"souksyp/react-vi/html"
	"souksyp/react-vi/store"

	"github.com/labstack/echo"
)

func Start() {
	// Init local server
	e := echo.New()

	// Hide the banner message and ASCII art logo
	e.HideBanner = true
	e.HidePort = true

	// Serves embedded index.html
	e.GET("/", func(ctx echo.Context) error {
		// Write some bytes to the response body
		w := ctx.Response().Writer
		_, err := w.Write(html.IndexPage)
		if err != nil {
			return err
		}
		return nil
	})

	// Serves data as JSON for graph
	e.GET("/data", func(c echo.Context) error {
		return c.JSON(http.StatusOK, store.NodesMap)
	})

	// Run server
	e.Logger.Fatal(e.Start(":18881"))
}
