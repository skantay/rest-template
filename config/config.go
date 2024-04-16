package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct{}

func New(path string) (Config, error) {
	viper.SetConfigFile(path)

	var config Config

	if err := viper.ReadInConfig(); err != nil {
		return config, fmt.Errorf("failed to read config: %w", err)
	}

	if err := viper.Unmarshal(&config); err != nil {
		return config, fmt.Errorf("failed to unmarshal config: %w", err)
	}

	return config, nil
}
