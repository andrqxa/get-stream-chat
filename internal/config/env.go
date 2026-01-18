package config

import (
	"fmt"
	"os"
)

func getEnvDefault(key string, def string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return def
}

func getEnvRequired(key string) (string, error) {
	if v := os.Getenv(key); v != "" {
		return v, nil
	}
	return "", fmt.Errorf("required env var %s is not set", key)
}
