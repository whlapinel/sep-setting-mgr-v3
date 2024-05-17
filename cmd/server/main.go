package main

import (
	"fmt"
	"log"
	"os"
	"sep_setting_mgr/internal/assets"
	"sep_setting_mgr/internal/database"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	LoadEnvironment()
	e := echo.New()
	e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogURI:     true,
		LogStatus:  true,
		LogMethod:  true,
		LogLatency: true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			const (
				reset  = "\033[0m"
				red    = "\033[31m"
				green  = "\033[32m"
				yellow = "\033[33m"
			)
			statusColor := reset
			if v.Status >= 400 {
				statusColor = red
			} else if v.Status >= 300 {
				statusColor = yellow
			} else {
				statusColor = green
			}
			methodWidth := 5
			uriWidth := 25
			statusWidth := 6
			customWidth := 12
			latencyWidth := 15
			value, _ := c.Get("id").(int)
			logLine := fmt.Sprintf("%-*s %-*s %s%-*d%s %-*d %-*s",
				methodWidth, v.Method,
				uriWidth, v.URI,
				statusColor,
				statusWidth, v.Status,
				reset,
				customWidth, value,
				latencyWidth, v.Latency,
			)
			fmt.Println(logLine)
			return nil
		},
	}))
	db, err := database.InitializeDB(false)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	err = MountHandlers(e, db)
	if err != nil {
		log.Fatal(err)
	}
	// scripts, styles and images are embedded in the binary
	assets.RegisterStatic(e)
	e.Logger.Fatal(e.Start(":1323"))
}

func LoadEnvironment() {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	if os.Getenv("ENV") == "development" {
		godotenv.Load("../../.env.development")
	}
	if os.Getenv("ENV") == "production" {
		godotenv.Load("../../.env.production")
	}
}
