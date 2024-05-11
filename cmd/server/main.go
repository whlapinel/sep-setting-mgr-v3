package main

import (
	"sep_setting_mgr/internal/assets"

	"github.com/labstack/echo"
)

func main() {
	// Echo instance
	e := echo.New()
	err := MountHandlers(e)
	if err != nil {
		e.Logger.Fatal(err)
	}
	assets.RegisterStatic(e)
	e.Logger.Fatal(e.Start(":1323"))
}
