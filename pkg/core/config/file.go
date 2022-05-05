package config

import (
	"fmt"
	"strings"

	"github.com/redhat-developer/app-services-cli/pkg/core/configutil"
)

const errorFormat = "%v: %w"
const defaultFileName = "config.json"

// NewFile creates a new config type
func NewFile(id string) IConfig {
	return &File{
		ConfigFile: configutil.NewConfigFile(
			id,
			defaultFileName,
			strings.ToUpper(id)+"CONFIG"),
	}
}

// File is a type which describes a config file
type File struct {
	configutil.ConfigFile
}

// Load loads the configuration from the configuration file. If the configuration file doesn't exist
// it will return an empty configuration object.
func (c *File) Load() (*Config, error) {
	cfg := &Config{}

	if err := c.ConfigFile.Load(cfg); err != nil {
		return cfg, fmt.Errorf(errorFormat, "unable to load config file", err)
	}

	return cfg, nil
}

// Save saves the given configuration to the configuration file.
func (c *File) Save(cfg *Config) error {
	if err := c.ConfigFile.Save(cfg); err != nil {
		return fmt.Errorf(errorFormat, "unable to save config file", err)
	}

	return nil
}

// Remove removes the configuration file.
func (c *File) Remove() error {
	if err := c.ConfigFile.Remove(); err != nil {
		return fmt.Errorf(errorFormat, "unable to remove config file", err)
	}

	return nil
}

// Returns the configuration file location.
func (c *File) Location() (string, error) {
	path, err := c.ConfigFile.Location()
	if err != nil {
		return "", fmt.Errorf(errorFormat, "unable to determine config file location", err)
	}

	return path, nil
}
