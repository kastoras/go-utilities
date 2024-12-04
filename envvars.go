package utilities

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func GetEnvParam(parameter string, defaultVal string) (string, error) {
	// First, try to get the environment variable from the system (e.g., Docker Compose environment)
	param := os.Getenv(parameter)
	if param != "" {
		return param, nil
	}

	// If not found, try to load from .env file
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: No env parameters set or error loading .env file")
		return "", err
	}

	param = os.Getenv(parameter)
	if param != "" {
		return param, nil
	}

	// If still not found, return the default value or an error if no default is provided
	if defaultVal == "" {
		return "", fmt.Errorf("warning: no value set for parameter %v", parameter)
	}

	return defaultVal, nil
}
