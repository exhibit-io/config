package config

import (
	"fmt"
	"net/url"
	"strconv"
)

// RedisConfig represents the configuration for Redis.
type RedisConfig struct {
	Host     string
	Port     string
	Password string
	DB       int
}

// GetURL constructs and returns the Redis URL as a *url.URL type.
func (s *RedisConfig) GetURL() (*url.URL, error) {
	url := &url.URL{
		Scheme: "redis",
		Host:   fmt.Sprintf("%s:%s", s.Host, s.Port),
		Path:   strconv.Itoa(s.DB),
	}

	return url, nil
}

// GetAddr returns the Redis address in the format host:port.
func (s *RedisConfig) GetAddr() string {
	return fmt.Sprintf("%s:%s", s.Host, s.Port)
}

// Load loads the Redis configuration from environment variables.
func (s *RedisConfig) Load() error {
	s.Host = GetEnv("REDIS_HOST", "localhost")
	s.Port = GetEnv("REDIS_PORT", "6379")
	s.Password = GetEnv("REDIS_PASSWORD", "")
	s.DB = GetEnvAsInt("REDIS_DB", 0)
	return nil
}

// Validate validates the Redis configuration.
func (s *RedisConfig) Validate() error {
	if s.Host == "" || s.Port == "" {
		return fmt.Errorf("invalid Redis configuration: host and port must be specified")
	}
	return nil
}
