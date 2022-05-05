package configutil

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
)

type ConfigError struct {
	err      error
	location string
	configID string
}

func (c ConfigError) Error() string {
	return c.err.Error()
}
func (c ConfigError) Unwrap() error {
	return c.err
}
func (c ConfigError) Location() string {
	return c.location
}
func (c ConfigError) ID() string {
	return c.configID
}

type ConfigFile struct {
	envName         string
	configID        string
	defaultFileName string
}

func NewConfigFile(configID string, defaultFileName string, envName string) ConfigFile {
	c := ConfigFile{
		configID:        configID,
		envName:         envName,
		defaultFileName: defaultFileName,
	}

	return c
}

// Load loads from the configuration file. If the file doesn't
// exist, no error will be returned and the given contect will
// be left untouched
func (c *ConfigFile) Load(into interface{}) error {
	path, err := c.Location()
	if err != nil {
		return err
	}

	_, err = os.Stat(path)
	if os.IsNotExist(err) {
		return nil
	}
	if err != nil {
		return ConfigError{
			err:      err,
			location: path,
			configID: c.configID,
		}
	}
	// #nosec G304
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return ConfigError{
			err:      err,
			location: path,
			configID: c.configID,
		}
	}

	err = json.Unmarshal(data, &into)
	if err != nil {
		return ConfigError{
			err:      err,
			location: path,
			configID: c.configID,
		}
	}

	return nil
}

func (c *ConfigFile) Remove() error {
	path, err := c.Location()
	if err != nil {
		return err
	}

	_, err = os.Stat(path)
	if os.IsNotExist(err) {
		return nil
	}
	err = os.Remove(path)
	if err != nil {
		return err
	}
	return nil
}

func (c *ConfigFile) Save(content interface{}) error {
	path, err := c.Location()
	if err != nil {
		return err
	}

	data, err := json.MarshalIndent(content, "", "  ")
	if err != nil {
		return ConfigError{
			err:      err,
			location: path,
			configID: c.configID,
		}
	}

	cfgDir, err := c.defaultDir()
	if err != nil {
		return err
	}
	if _, err = os.Stat(cfgDir); os.IsNotExist(err) {
		err = os.Mkdir(cfgDir, 0o700)
		if err != nil {
			return err
		}
	}
	err = ioutil.WriteFile(path, data, 0o600)
	if err != nil {
		return ConfigError{
			err:      err,
			location: path,
			configID: c.configID,
		}
	}
	return nil
}

func (c *ConfigFile) Location() (string, error) {
	path := os.Getenv(c.envName)
	if path != "" {
		return path, nil
	}

	dir, err := c.defaultDir()
	if err != nil {
		return "", err
	}

	if _, err = os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, 0o700)
		if err != nil {
			return "", err
		}
	}

	path = filepath.Join(dir, c.defaultFileName)
	if err != nil {
		return "", err
	}

	return path, nil
}

func (c *ConfigFile) defaultDir() (string, error) {
	userCfgDir, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(userCfgDir, c.configID), nil
}
