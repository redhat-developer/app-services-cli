package config

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/connection"
	"github.com/mitchellh/go-homedir"
)

// New creates a new config type
func New() *Config {
	cfg := &Config{}

	return cfg
}

// Load loads the configuration from the configuration file. If the configuration file doesn't exist
// it will return an empty configuration object.
func Load(cfg *Config) error {
	file, err := Location()
	if err != nil {
		return err
	}
	_, err = os.Stat(file)
	if os.IsNotExist(err) {
		return err
	}
	if err != nil {
		return fmt.Errorf("can't check if config file '%s' exists: %w", file, err)
	}
	// #nosec G304
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return fmt.Errorf("can't read config file '%s': %w", file, err)
	}
	err = json.Unmarshal(data, &cfg)
	if err != nil {
		return fmt.Errorf("can't parse config file '%s': %w", file, err)
	}
	return nil
}

// Save saves the given configuration to the configuration file.
func (c *Config) Save() error {
	file, err := Location()
	if err != nil {
		return err
	}
	data, err := json.MarshalIndent(c, "", "  ")
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
func (c *Config) Remove() error {
	file, err := Location()
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

func (c *Config) Connection() (conn *connection.Connection, err error) {
	builder := connection.NewBuilder()
	if c.AccessToken != "" {
		builder.WithAccessToken(c.AccessToken)
	}
	if c.RefreshToken != "" {
		builder.WithRefreshToken(c.RefreshToken)
	}
	if c.ClientID != "" {
		builder.WithClientID(c.ClientID)
	}
	if c.Scopes != nil {
		builder.WithScopes(c.Scopes...)
	}
	if c.URL != "" {
		builder.WithURL(c.URL)
	}
	if c.AuthURL == "" {
		c.AuthURL = connection.DefaultAuthURL
	}
	builder.WithAuthURL(c.AuthURL)

	builder.WithInsecure(c.Insecure)

	conn, err = builder.Build()
	if err != nil {
		return nil, err
	}

	accessTk, refreshTk, err := conn.RefreshTokens(context.TODO())
	if err != nil {
		return nil, err
	}

	accessTkChanged := accessTk != c.AccessToken
	refreshTkChanged := refreshTk != c.RefreshToken

	if accessTkChanged {
		c.AccessToken = accessTk
	}
	if refreshTkChanged {
		c.RefreshToken = refreshTk
	}

	if !accessTkChanged && refreshTkChanged {
		return conn, nil
	}

	if err != nil {
		return nil, fmt.Errorf("Unable to save config file: %w", err)
	}

	return conn, nil
}

func Location() (path string, err error) {
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
