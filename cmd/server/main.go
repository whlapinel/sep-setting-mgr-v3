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
	demo := false
	clearDB := false
	LoadEnvironment()
	if os.Getenv("DEMO") == "true" {
		demo = true
	}
	if os.Getenv("CLEAR_DB") == "true" {
		clearDB = true
	}
	log.Println("Environment: " + os.Getenv("ENV"))
	log.Println("Host: " + os.Getenv("HOST"))

	log.Println("Starting server on port " + os.Getenv("PORT") + " in " + os.Getenv("ENV") + " mode")
	e := echo.New()
	e.Use(logger)
	db, err := repositories.InitializeDB(false)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if clearDB {
		err = repositories.ClearDatabase(db)
		if err != nil {
			log.Fatal(err)
		}
	}
	err = MountHandlers(e, db, demo)
	if err != nil {
		log.Fatal(err)
	}
	// scripts, styles and images are embedded in the binary
	assets.RegisterStatic(e)
	if os.Getenv("ENV") == "development" {
		e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))
	} else {
		e.Logger.Fatal(e.StartAutoTLS(os.Getenv("HOST") + ":" + os.Getenv("PORT")))
	}
}

func LoadEnvironment() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println(err)
		log.Fatal("Error loading .env file")
	}
	if os.Getenv("ENV") == "development" {
		godotenv.Load(".env.development")
	}
	if os.Getenv("ENV") == "production" {
		godotenv.Load(".env.production")
	}
}
