package config

import (
	"fmt"
)

type ServerConfig struct {
	Host string
	Port string
}

func (s ServerConfig) GetAddr() string {
	return fmt.Sprintf("%s:%s", s.Host, s.Port)
}

// Load loads the Redis configuration from environment variables.
func (s *ServerConfig) Load() error {
	s.Host = GetEnv("HOST", "0.0.0.0")
	s.Port = GetEnv("PORT", "8080")
	return nil
}

// Validate validates the Redis configuration.
func (s *ServerConfig) Validate() error {
	if s.Host == "" || s.Port == "" {
		return fmt.Errorf("invalid Redis configuration: host and port must be specified")
	}
	return nil
}
