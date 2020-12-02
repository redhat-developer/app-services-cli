package factory

import (
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/cmdutil"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/config"
)

func New(appVersion string) *cmdutil.Factory {
	var cachedConfig *config.Config
	var configError error
	configFunc := func() (config.Config, error) {
		if cachedConfig != nil || configError != nil {
			return *cachedConfig, configError
		}
		cachedConfig, configError := config.Load()
		return *cachedConfig, configError
	}

	return &cmdutil.Factory{
		Config: configFunc,
	}
}
