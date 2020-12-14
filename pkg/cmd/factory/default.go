package factory

import (
	"os"

	"github.com/bf2fc6cc711aee1a0c2a/cli/internal/config"
)

func New(cliVersion string) *Factory {
	var cfg *config.Config
	var configError error
	configFunc := func() (config.Config, error) {
		if cfg != nil || configError != nil {
			return *cfg, configError
		}

		configError = config.Load(cfg)
		if os.IsNotExist(configError) {
			cfg = config.New()
			configError = nil
		}
		return *cfg, configError
	}

	return &Factory{
		Config: configFunc,
	}
}
