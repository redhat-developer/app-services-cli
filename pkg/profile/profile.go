// This file contains functions that help to manage visibility of early stage commands
package profile

import (
	"github.com/redhat-developer/app-services-cli/internal/config"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/factory"
)

// Visual element displayed in help
func ApplyDevPreviewLabel(message string) string {
	return "[Preview] " + message
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
func EnableDevPreview(f *factory.Factory, enablement bool) (*config.Config, error) {
	logger, err := f.Logger()
	if err != nil {
		logger.Info(f.Localizer.MustLocalize("profile.error.enablement"), err)
		return nil, err
	}

	config, err := f.Config.Load()
	if err != nil {
		logger.Info(f.Localizer.MustLocalize("profile.error.enablement"), err)
		return nil, err
	}

	config.DevPreviewEnabled = enablement
	err = f.Config.Save(config)
	if err != nil {
		logger.Info(f.Localizer.MustLocalize("profile.error.enablement"), err)
		return nil, err
	}
	if config.DevPreviewEnabled {
		logger.Info(f.Localizer.MustLocalize("profile.status.devpreview.enabled"))
	} else {
		logger.Info(f.Localizer.MustLocalize("profile.status.devpreview.disabled"))
	}
	return config, err
}
