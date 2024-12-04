package utilities

import (
	"fmt"
	"log"	
	"os"

	"github.com/joho/godotenv"
)

func GetEnvParam(parameter string, defaultVal string) (string, error) {
	param := os.Getenv(parameter)
	if param != "" {
		return param, nil
	}

	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: No env parameters set or error loading .env file")
		return "", err
	}

	param = os.Getenv(parameter)
	if param != "" {
		return param, nil
	}
	
	if defaultVal == "" {
		return "", fmt.Errorf("warning: no value set for parameter %v", parameter)
	}

	return defaultVal, nil
}
