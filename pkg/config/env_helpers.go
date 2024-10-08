package config

import (
	"os"
	"strconv"
	"strings"
)

// getEnv gets an environment variable or returns a fallback value.
func GetEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}

// GetEnvAsInt gets an environment variable as an integer or returns a fallback value.
func GetEnvAsInt(key string, fallback int) int {
	valueStr := GetEnv(key, "")
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}
	return fallback
}

// GetEnvAsBool gets an environment variable as a boolean or returns a fallback value.
func GetEnvAsBool(key string, fallback bool) bool {
	valueStr := GetEnv(key, "")
	if value, err := strconv.ParseBool(valueStr); err == nil {
		return value
	}
	return fallback
}

// GetEnvAsSlice gets an environment variable as a slice or returns a fallback value.
func GetEnvAsSlice(key string, fallback []string, separator string) []string {
	valueStr := GetEnv(key, "")
	if valueStr == "" {
		return fallback
	}
	return strings.Split(valueStr, separator)
}
