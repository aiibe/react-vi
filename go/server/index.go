package server

import (
	"net/http"
	"souksyp/react-vi/public"
	"souksyp/react-vi/store"

	"github.com/labstack/echo/v4"
)

func Start() {

	// Init local server
	e := echo.New()

	// Hide the banner message and ASCII art logo
	e.HideBanner = true
	e.HidePort = true

	// Serve React app
	e.StaticFS("/", public.FrontApp)

	// Serves data as JSON for graph
	e.GET("/data", func(c echo.Context) error {
		return c.JSON(http.StatusOK, store.NodesMap)
	})

	// Run server
	e.Logger.Fatal(e.Start(":18881"))
}
