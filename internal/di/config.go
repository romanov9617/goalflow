package di

import (
	"fmt"
	"os"

	"pglease/internal/adapter/config"
)

const cfgPathEnvVar = "CONFIG_PATH"

func NewConfig() (*config.Config, error) {
	cfgPath := os.Getenv(cfgPathEnvVar)
	if cfgPath == "" {
		return nil, fmt.Errorf("di.NewConfig: ConfigPath is empty")
	}
	file, err := os.Open(cfgPath)
	if err != nil {
		return nil, fmt.Errorf("di.NewConfig: open cfgPath: %w", err)
	}

	cfg, err := config.LoadConfig(file)
	if err != nil {
		return nil, fmt.Errorf("di.NewConfig: LoadConfig: %w", err)
	}
	return cfg, nil
}
