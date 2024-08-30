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
func (r *RedisConfig) GetURL() (*url.URL, error) {
	url := &url.URL{
		Scheme: "redis",
		Host:   fmt.Sprintf("%s:%s", r.Host, r.Port),
		Path:   strconv.Itoa(r.DB),
	}

	return url, nil
}

// GetAddr returns the Redis address in the format host:port.
func (r *RedisConfig) GetAddr() string {
	return fmt.Sprintf("%s:%s", r.Host, r.Port)
}

// Load loads the Redis configuration from environment variables.
func (r *RedisConfig) Load() error {
	r.Host = getEnv("REDIS_HOST", "localhost")
	r.Port = getEnv("REDIS_PORT", "6379")
	r.Password = getEnv("REDIS_PASSWORD", "")
	r.DB = getEnvAsInt("REDIS_DB", 0)
	return nil
}

// Validate validates the Redis configuration.
func (r *RedisConfig) Validate() error {
	if r.Host == "" || r.Port == "" {
		return fmt.Errorf("invalid Redis configuration: host and port must be specified")
	}
	return nil
}
