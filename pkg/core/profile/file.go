package profile

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

// NewFile creates a new context type
func NewFile() IContext {
	cfg := &File{}

	return cfg
}

// File is a type which describes a context file
type File struct{}

const errorFormat = "%v: %w"

const envName = "RHOASCONTEXT"

const customContextLocation = "RHOAS_CUSTOM_CONTEXT"

// Load loads the profiles from the context file. If the context file doesn't exist
// it will return an empty context object.
func (c *File) Load() (*Context, error) {
	file, err := c.Location()
	if err != nil {
		return nil, err
	}

	_, err = os.Stat(file)
	if os.IsNotExist(err) {
		return nil, err
	}
	if err != nil {
		return nil, fmt.Errorf(errorFormat, "unable to check if context file exists", err)
	}
	// #nosec G304
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, fmt.Errorf(errorFormat, "unable to read context file", err)
	}
	var cfg Context
	err = json.Unmarshal(data, &cfg)
	if err != nil {
		return nil, fmt.Errorf(errorFormat, "unable to parse contexts", err)
	}
	return &cfg, nil
}

// Save saves the given profiles to the context file.
func (c *File) Save(cfg *Context) error {
	file, err := c.Location()
	if err != nil {
		return err
	}
	data, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		return fmt.Errorf("%v: %w", "unable to marshal context", err)
	}
	rhoasCfgDir, err := DefaultDir()
	if err != nil {
		return err
	}
	if _, err = os.Stat(rhoasCfgDir); os.IsNotExist(err) {
		err = os.Mkdir(rhoasCfgDir, 0o700)
		if err != nil {
			return err
		}
	}
	err = ioutil.WriteFile(file, data, 0o600)
	if err != nil {
		return fmt.Errorf(errorFormat, "unable to save context", err)
	}
	return nil
}

// Remove removes the context file.
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

// Location gets the path to the context file
func (c *File) Location() (path string, err error) {

	if rhoasCustomContext := os.Getenv(customContextLocation); rhoasCustomContext != "" {
		return rhoasCustomContext, err
	}

	if rhoasContext := os.Getenv(envName); rhoasContext != "" {
		path = rhoasContext
	} else {
		rhoasCtxDir, err := DefaultDir()
		if err != nil {
			return "", err
		}
		path = filepath.Join(rhoasCtxDir, "contexts.json")
		if err != nil {
			return "", err
		}
	}
	return path, nil
}

// Checks if config has custom location
func HasCustomLocation() bool {
	rhoasContext := os.Getenv(customContextLocation)
	return rhoasContext != ""
}

// DefaultDir returns the default parent directory of the context file
func DefaultDir() (string, error) {
	userCfgDir, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(userCfgDir, "rhoas"), nil
}
