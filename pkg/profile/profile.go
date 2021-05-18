// This file contains functions that help to manage visibility of early stage commands
package profile

import (
	"strings"

	"github.com/redhat-developer/app-services-cli/internal/config"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/factory"
	"github.com/redhat-developer/app-services-cli/pkg/color"
)

// Visual element displayed in help
func DevPreviewLabel() string {
	return color.Info("[Preview] ")
}

// Annotation used in templates and tools like documentation generation
func DevPreviewAnnotation() map[string]string {
	return map[string]string{"channel": "preview"}
}

// Check if preview is enabled
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

// Enable dev preview
func EnableDevPreview(f *factory.Factory, enablement *string) (*config.Config, error) {
	logger, err := f.Logger()
	if err != nil {
		logger.Info("Cannot enable dev preview. ", err)
		return nil, err
	}

	if *enablement == "" {
		logger.Info("Skip")
		// Flag not present no action needed.
		return nil, nil
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
	if config.DevPreviewEnabled {
		logger.Info("Developer Preview commands activated. Use help command to view them.")
	} else {
		logger.Info("Developer Preview commands deactivated.")
	}
	return config, err
}
