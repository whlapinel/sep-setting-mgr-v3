package main

import (
	"log"
	"os"
	"sep_setting_mgr/internal/assets"
	"sep_setting_mgr/internal/repositories"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	clearData := false
	LoadEnvironment()
	e := echo.New()
	e.Use(logger)
	db, err := repositories.InitializeDB(false)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	if clearData {
		err = repositories.ClearDatabase(db)
		if err != nil {
			log.Fatal(err)
		}
	}
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
