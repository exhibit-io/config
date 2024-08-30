package config

// CorsConfig represents the configuration for Cross-Origin Resource Sharing.
type CorsConfig struct {
	AllowedOrigins   []string
	AllowedMethods   []string
	AllowedHeaders   []string
	AllowCredentials bool
}

// Load loads the CORS configuration from environment variables.
func (c *CorsConfig) Load() error {
	c.AllowedOrigins = GetEnvAsSlice("CORS_ALLOWED_ORIGINS", []string{"*"}, ",")
	c.AllowedMethods = GetEnvAsSlice("CORS_ALLOWED_METHODS", []string{"GET", "POST"}, ",")
	c.AllowedHeaders = GetEnvAsSlice("CORS_ALLOWED_HEADERS", []string{"*"}, ",")
	c.AllowCredentials = GetEnvAsBool("CORS_ALLOW_CREDENTIALS", false)
	return nil
}

// Validate validates the CORS configuration.
func (c *CorsConfig) Validate() error {
	// Add any necessary validation here
	return nil
}
