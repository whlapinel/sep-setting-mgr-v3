package main

import (
	"sep_setting_mgr/internal/assets"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	// Echo instance
	e := echo.New()
	err := MountHandlers(e)
	if err != nil {
		e.Logger.Fatal(err)
	}
	e.Use(middleware.Logger())
	// scripts, styles and images are embedded in the binary
	assets.RegisterStatic(e)
	e.Logger.Fatal(e.Start(":1323"))
}
