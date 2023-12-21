package config

import (
	"github.com/idan-fishman/whichfile/pkg/validators"
	"github.com/spf13/viper"
)

// Config defines the project settings.
type Config struct {
	Log LogConfig `validate:"required"`
}

// LogConfig defines the project logging settings.
type LogConfig struct {
	Level string `validate:"required,oneof=debug info warn error panic fatal"`
}

var (
	C Config // Global config instance.
)

// LoadConfig loads the configuration from environment variables and validates it.
func LoadConfig() (Config, error) {
	// Load environment variables and set defaults.
	viper.AutomaticEnv()
	viper.SetDefault("LOG_LEVEL", "warn")

	// Populate the config struct.
	config := Config{
		Log: LogConfig{
			Level: viper.GetString("LOG_LEVEL"),
		},
	}

	// Validate the config struct.
	err := validators.V.Struct(config)
	if err != nil {
		return Config{}, err
	}

	return config, nil
}
