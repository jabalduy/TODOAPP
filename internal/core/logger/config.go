package core_logger

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

type LoggerConfig struct {
	Level  string `envconfig:"LOGGER_LEVEL" required="true"`
	Folder string `envconfig:"LOGGER_CONFIG" required="true"`
}

func NewConfig() (LoggerConfig, error) {
	var config LoggerConfig

	if err := envconfig.Process("", &config); err != nil {
		return LoggerConfig{}, fmt.Errorf("process envconfig: %w", err)
	}

	return config, nil
}
