package config

import (
	"github.com/joho/godotenv"
	"log"
)

// Config holds all service configurations.
type Config struct {
	configs map[string]Configurable
}

// Default initializes a new Config instance with default configurations.
func Default() *Config {
	return &Config{
		configs: map[string]Configurable{
			"redis":  &RedisConfig{}, // Default configurations
			"cors":   &CorsConfig{},
			"server": &ServerConfig{},
			// Additional default configurations can be added here...
		},
	}
}

// New initializes an empty Config which needs to be populated with Configurable
func New() *Config {
	return &Config{
		configs: map[string]Configurable{},
	}
}

// Register allows adding external configurations.
func (c *Config) Register(key string, config Configurable) {
	c.configs[key] = config
}

// LoadConfig loads all registered configurations.
func (c *Config) LoadConfig() error {
	err := godotenv.Load()
	if err != nil {
		log.Default().Print("No .env file found, using default environment variables.")
	}

	for key, config := range c.configs {
		if err := config.Load(); err != nil {
			return err
		}
		if err := config.Validate(); err != nil {
			return err

		}
		log.Printf("Loaded [%s] configuration.", key)
	}

	return nil
}
