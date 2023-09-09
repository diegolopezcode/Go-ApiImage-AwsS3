package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func GetConfig(variable string) string {
	dotenv := godotenv.Load()
	if dotenv != nil {
		log.Fatal("Error loading .env file")
	}
	return os.Getenv(variable)
}
