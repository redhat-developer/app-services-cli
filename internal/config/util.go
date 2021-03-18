package config

import (
	"errors"
	"os"
	"path/filepath"
	"runtime"

	"github.com/bf2fc6cc711aee1a0c2a/cli/internal/localizer"
)

// GetUserConfig gets the path to store user-specific configurations
func GetUserConfig(name string) (string, error) {
	currentOs := runtime.GOOS
	switch currentOs {
	case "windows":
		return filepath.Join(os.Getenv("APPDATA"), name), nil
	case "darwin":
		return filepath.Join(os.Getenv("HOME"), "Library", "Application Support", name), nil
	case "linux":
		baseDir := os.Getenv("XDG_CONFIG_HOME")
		if baseDir == "" {
			baseDir = filepath.Join(os.Getenv("HOME"), ".config")
		}
		return filepath.Join(baseDir, name), nil
	default:
		return "", errors.New(localizer.MustLocalize(&localizer.Config{
			MessageID: "common.error.unsupportedOperatingSystem",
			TemplateData: map[string]interface{}{
				"OS": runtime.GOOS,
			},
		}))
	}
}
