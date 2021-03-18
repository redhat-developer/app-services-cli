package config

import (
	"os"
	"path/filepath"
	"runtime"
)

// GetUserConfig gets the path to store user-specific configurations
func GetUserConfig(name string) string {
	currentOs := runtime.GOOS
	switch currentOs {
	case "windows":
		return filepath.Join(os.Getenv("APPDATA"), name)
	case "darwin":
		return filepath.Join(os.Getenv("HOME"), "Library", "Application Support", name)
	default:
		baseDir := os.Getenv("XDG_CONFIG_HOME")
		if baseDir == "" {
			baseDir = filepath.Join(os.Getenv("HOME"), ".config")
		}
		return filepath.Join(baseDir, name)
	}
}
