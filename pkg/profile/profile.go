// This file contains functions that add common arguments to the command line
package profile

import (
	"strings"

	"github.com/redhat-developer/app-services-cli/internal/config"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/factory"
	"github.com/redhat-developer/app-services-cli/pkg/color"
)

func DevPreviewLabel() string {
	return color.Bold(" [Preview] ")
}

func DevPreviewAnnotation() map[string]string {
	return map[string]string{"channel": "preview"}
}

func DevPreviewEnabled(f *factory.Factory) bool {
	logger, err := f.Logger()
	if err != nil {
		logger.Info("Cannot determine status of dev preview. ", err)
		return false
	}
	config, err := f.Config.Load()
	if err != nil {
		logger.Info("Cannot determine status of dev preview. ", err)
		return false
	}

	return config.DevPreviewEnabled
}

func EnableDevPreview(f *factory.Factory, enablement *string) (*config.Config, error) {
	if *enablement == "" {
		// Flag not present no action needed.
		return nil, nil
	}

	logger, err := f.Logger()
	if err != nil {
		logger.Info("Cannot enable dev preview. ", err)
		return nil, err
	}
	config, err := f.Config.Load()
	if err != nil {
		logger.Info("Cannot enable dev preview.", err)
		return nil, err
	}

	config.DevPreviewEnabled = strings.ToLower(*enablement) == "true" || *enablement == "yes" || *enablement == "y"
	err = f.Config.Save(config)
	if err != nil {
		logger.Info("Cannot enable dev preview. ", err)
		return nil, err
	}
	return config, err
}
