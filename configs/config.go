package configs

import (
	"errors"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func GetConfig(variable string) (string, error) {
	dotenv := godotenv.Load()
	if dotenv != nil {
		log.Fatal("Error loading .env file")
		return "", errors.New("Error loading .env file")
	}
	result := os.Getenv(variable)
	if result == "" {
		log.Fatal("Error loading .env file")
		return "", errors.New("variable not found")
	}

	return result, nil
}
