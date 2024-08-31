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

func GetConfig[T Configurable](c *Config, key string) (T, bool) {
	if config, ok := c.configs[key]; ok {
		// Type assertion to the expected generic type T
		if typedConfig, ok := config.(T); ok {
			return typedConfig, true
		}
	}

	// Return the zero value of T and false if the key doesn't exist or the type assertion fails
	var zero T
	return zero, false
}
