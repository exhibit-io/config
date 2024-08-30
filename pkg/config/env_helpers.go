package config

import (
	"os"
	"strconv"
	"strings"
)

// getEnv gets an environment variable or returns a fallback value.
func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}

// getEnvAsInt gets an environment variable as an integer or returns a fallback value.
func getEnvAsInt(key string, fallback int) int {
	valueStr := getEnv(key, "")
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}
	return fallback
}

// getEnvAsBool gets an environment variable as a boolean or returns a fallback value.
func getEnvAsBool(key string, fallback bool) bool {
	valueStr := getEnv(key, "")
	if value, err := strconv.ParseBool(valueStr); err == nil {
		return value
	}
	return fallback
}

// getEnvAsSlice gets an environment variable as a slice or returns a fallback value.
func getEnvAsSlice(key string, fallback []string, separator string) []string {
	valueStr := getEnv(key, "")
	if valueStr == "" {
		return fallback
	}
	return strings.Split(valueStr, separator)
}
