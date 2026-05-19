package config

import (
	"github.com/pkg/errors"

	"github.com/kelseyhightower/envconfig"
)

// Config represents the application configuration.
type Config struct {
	API struct {
		Port         int `envconfig:"API_PORT" default:"3000"`
		ReadTimeout  int `envconfig:"API_READ_TIMEOUT" default:"7"`
		WriteTimeout int `envconfig:"API_WRITE_TIMEOUT" default:"5"`
		IdleTimeout  int `envconfig:"API_IDLE_TIMEOUT" default:"5"`
		Timeout      int `envconfig:"API_TIMEOUT" default:"5"`
	}
}

// NewConfig creates a new initialised application Config.
func NewConfig() (*Config, error) {
	var config Config

	if err := envconfig.Process("", &config); err != nil {
		return nil, errors.Wrap(err, "error parsing config")
	}

	return &config, nil
}
