package main

import (
	//import echo
	"sep_setting_mgr/cmd/server/routes"

	"github.com/labstack/echo"
)

func main() {
	// Echo instance
	e := echo.New()
	routes.RegisterRoutes(e)
	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
