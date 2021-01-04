package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/mitchellh/go-homedir"
)

// NewFile creates a new config type
func NewFile() IConfig {
	cfg := &File{}

	return cfg
}

// File is a type which describes a config file
type File struct{}

// Load loads the configuration from the configuration file. If the configuration file doesn't exist
// it will return an empty configuration object.
func (c *File) Load() (*Config, error) {
	file, err := c.Location()
	if err != nil {
		return nil, err
	}
	_, err = os.Stat(file)
	if os.IsNotExist(err) {
		return nil, err
	}
	if err != nil {
		return nil, fmt.Errorf("can't check if config file '%s' exists: %w", file, err)
	}
	// #nosec G304
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, fmt.Errorf("can't read config file '%s': %w", file, err)
	}
	var cfg Config
	err = json.Unmarshal(data, &cfg)
	if err != nil {
		return nil, fmt.Errorf("can't parse config file '%s': %w", file, err)
	}
	return &cfg, nil
}

// Save saves the given configuration to the configuration file.
func (c *File) Save(cfg *Config) error {
	file, err := c.Location()
	if err != nil {
		return err
	}
	data, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		return fmt.Errorf("can't marshal config: %w", err)
	}
	err = ioutil.WriteFile(file, data, 0600)
	if err != nil {
		return fmt.Errorf("can't write file '%s': %w", file, err)
	}
	return nil
}

// Remove removes the configuration file.
func (c *File) Remove() error {
	file, err := c.Location()
	if err != nil {
		return err
	}
	_, err = os.Stat(file)
	if os.IsNotExist(err) {
		return nil
	}
	err = os.Remove(file)
	if err != nil {
		return err
	}
	return nil
}

// Location gets the path to the config file
func (c *File) Location() (path string, err error) {
	if rhoasConfig := os.Getenv("RHOASCLI_CONFIG"); rhoasConfig != "" {
		path = rhoasConfig
	} else {
		home, err := homedir.Dir()
		if err != nil {
			return "", err
		}
		path = filepath.Join(home, ".rhoascli.json")
	}
	return path, nil
}
