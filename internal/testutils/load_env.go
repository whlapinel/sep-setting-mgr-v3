package testutils

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

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
