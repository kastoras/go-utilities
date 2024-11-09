package utilities

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func GetEnvParam(parameter string, defaultVal string) (string, error) {
	err := godotenv.Load()
	if err != nil {
		return "", err
	}

	param := os.Getenv(parameter)
	if param != "" {
		return param, err
	}

	if defaultVal == "" {
		return "", fmt.Errorf("warning: no value set for parameter %v", param)
	}

	return defaultVal, nil
}
