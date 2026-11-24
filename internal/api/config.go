package api

import (
	"github.com/google/uuid"
	"github.com/kelseyhightower/envconfig"
)

// Config holds HTTP server settings and the identity returned by /health.
type Config struct {
	AppPort     int    `default:"8080" envconfig:"APP_PORT"`
	ServiceName string `default:"urlshortner" envconfig:"SERVICE_NAME"`
	InstanceID  string `envconfig:"ID"`
}

// NewConfig reads APP_PORT, SERVICE_NAME, and ID from the environment.
// APP_PORT defaults to 8080 and SERVICE_NAME defaults to "urlshortner".
// It generates a UUID when ID is empty.
// It returns an error if an environment value cannot be parsed.
func NewConfig() (*Config, error) {
	cfg := &Config{}

	err := envconfig.Process("", cfg)
	if err != nil {
		return nil, err
	}

	if cfg.InstanceID == "" {
		cfg.InstanceID = uuid.NewString()
	}

	return cfg, nil
}
