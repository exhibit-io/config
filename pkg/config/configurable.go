package config

// Configurable is an interface that all configuration structs must implement.
type Configurable interface {
	Load() error
	Validate() error
}
