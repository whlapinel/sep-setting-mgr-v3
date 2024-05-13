package main

import (
	"log"
	"os"
	"sep_setting_mgr/internal/assets"
	"sep_setting_mgr/internal/database"

	"github.com/joho/godotenv"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
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
	e := echo.New()
	db, err := database.InitializeDB(false)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	err = MountHandlers(e, db)
	if err != nil {
		log.Fatal(err)
	}
	e.Use(middleware.Logger())
	// scripts, styles and images are embedded in the binary
	assets.RegisterStatic(e)
	e.Logger.Fatal(e.Start(":1323"))
}
