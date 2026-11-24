package api

import (
	"github.com/google/uuid"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	AppPort     int    `default:"8080" envconfig:"APP_PORT"`
	ServiceName string `default:"urlshortner" envconfig:"SERVICE_NAME"`
	InstanceID  string `envconfig:"ID"`
}

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
