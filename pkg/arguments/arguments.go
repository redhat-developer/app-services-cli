// This file contains functions that add common arguments to the command line
package arguments

import (
	"strings"

	"github.com/redhat-developer/app-services-cli/internal/config"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/debug"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/factory"
	"github.com/spf13/pflag"
)

// AddDebugFlag adds the '--debug' flag to the given set of command line flags
func AddDebugFlag(fs *pflag.FlagSet) {
	debug.AddFlag(fs)
}

// Enables dev preview in config
func EnableDevPreview(f *factory.Factory, enablement *string) (*config.Config, error) {
	if *enablement == "" {
		// Flag not present no action needed.
		return nil, nil
	}

	logger, err := f.Logger()
	if err != nil {
		logger.Info("Cannot enable dev preview")
		return nil, err
	}
	config, err := f.Config.Load()
	if err != nil {
		logger.Info("Cannot enable dev preview")
		return nil, err
	}

	config.DevPreviewEnabled = strings.ToLower(*enablement) == "true" || *enablement == "yes" || *enablement == "y"
	err = f.Config.Save(config)
	if err != nil {
		logger.Info("Cannot enable dev preview")
		return nil, err
	}
	return config, err
}
