package configs

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func GetConfig(variable string) (string, error) {
	dotenv := godotenv.Load(
		"../.env",
	)
	if dotenv != nil {
		log.Fatal("Error loading .env file", dotenv)
		return "", errors.New("Error loading .env file")
	}
	fmt.Println("variable: ", variable)
	result := os.Getenv(variable)
	if result == "" {
		log.Fatal("variable not found")
		return "", errors.New("variable not found")
	}

	return result, nil
}
