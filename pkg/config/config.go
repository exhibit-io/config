package config

import (
	"github.com/joho/godotenv"
	"log"
)

// Config holds all service configurations.
type Config struct {
	configs []Configurable
}

// NewConfig initializes a new Config instance with default configurations.
func New() *Config {
	return &Config{
		configs: []Configurable{
			&RedisConfig{}, // Default configurations
			&CorsConfig{},
			// Additional default configurations can be added here...
		},
	}
}

// Register allows adding external configurations.
func (c *Config) Register(config Configurable) {
	c.configs = append(c.configs, config)
}

// LoadConfig loads all registered configurations.
func (c *Config) LoadConfig() error {
	err := godotenv.Load()
	if err != nil {
		log.Default().Print("No .env file found, using default environment variables.")
	}

	for _, config := range c.configs {
		if err := config.Load(); err != nil {
			return err
		}
		if err := config.Validate(); err != nil {
			return err
		}
	}

	return nil
}
