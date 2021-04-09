package config

import (
	"errors"
	"os"
	"path/filepath"
	"runtime"

	"github.com/redhat-developer/app-services-cli/internal/localizer"
)

// gets the path to store user-specific configurations
func getUserConfig(name string) (string, error) {
	currentOs := runtime.GOOS
	switch currentOs {
	case "windows":
		return filepath.Join(os.Getenv("APPDATA"), name), nil
	case "darwin":
		return filepath.Join(os.Getenv("HOME"), "Library", "Application Support", name), nil
	case "linux":
		baseDir := getLinuxConfigBaseDir()
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

// Gets the config directory for Linux
func getLinuxConfigBaseDir() string {
	baseDir := os.Getenv("XDG_CONFIG_HOME")
	if baseDir != "" {
		return baseDir
	}
	baseDir = filepath.Join(os.Getenv("HOME"), ".config")
	// The config directory does not exist on the current environment
	// use the HOME directory
	if _, err := os.Stat(baseDir); os.IsNotExist(err) {
		baseDir = os.Getenv("HOME")
	}
	return baseDir
}
