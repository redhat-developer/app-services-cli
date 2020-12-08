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

// Config is the type used to track the config of the client
type Config struct {
	AccessToken  string           `json:"access_token,omitempty" doc:"Bearer access token."`
	RefreshToken string           `json:"refresh_token,omitempty" doc:"Offline or refresh token."`
	Services     ServiceConfigMap `json:"services"`
	URL          string           `json:"url,omitempty" doc:"URL of the API gateway. The value can be the complete URL or an alias. The valid aliases are 'production', 'staging' and 'integration'."`
	AuthURL      string           `json:"auth_url" doc:"URL of the authentication server"`
	ClientID     string           `json:"client_id,omitempty" doc:"OpenID client identifier."`
	Insecure     bool             `json:"insecure,omitempty" doc:"Enables insecure communication with the server. This disables verification of TLS certificates and host names."`
	Scopes       []string         `json:"scopes,omitempty" doc:"OpenID scope. If this option is used it will replace completely the default scopes. Can be repeated multiple times to specify multiple scopes."`
}

// ServiceConfigMap is a map of configs for the managed application services
type ServiceConfigMap struct {
	Kafka *KafkaConfig `json:"kafka"`
}

// KafkaConfig is the config for the managed Kafka service
type KafkaConfig struct {
	ClusterHost string `json:"clusterHost"`
	ClusterID   string `json:"clusterId"`
	ClusterName string `json:"clusterName"`
}

func (c *Config) SetAccessToken(accessToken string) {
	c.AccessToken = accessToken
}

func (c *Config) SetRefreshToken(refreshToken string) {
	c.RefreshToken = refreshToken
}

func (c *Config) SetClientID(clientID string) {
	c.ClientID = clientID
}

func (c *Config) SetScopes(scopes []string) {
	c.Scopes = scopes
}

func (c *Config) SetURL(url string) {
	c.URL = url
}

func (c *Config) SetAuthURL(authURL string) {
	c.AuthURL = authURL
}

func (c *Config) SetInsecure(insecure bool) {
	c.Insecure = insecure
}

func (c *Config) HasKafka() bool {
	return c.Services.Kafka != nil
}

// SetKafka sets the current Kafka cluster
func (s *ServiceConfigMap) SetKafka(k *KafkaConfig) {
	s.Kafka = k
}

// Remove the current Kafka cluster from the config
func (s *ServiceConfigMap) RemoveKafka() {
	s.Kafka = &KafkaConfig{
		ClusterID:   "",
		ClusterHost: "",
		ClusterName: "",
	}
}

// Save saves the given configuration to the configuration file.
func Save(cfg *Config) error {
	file, err := Location()
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

// Load loads the configuration from the configuration file. If the configuration file doesn't exist
// it will return an empty configuration object.
func Load() (cfg *Config, err error) {
	file, err := Location()
	if err != nil {
		return
	}
	_, err = os.Stat(file)
	if os.IsNotExist(err) {
		cfg = nil
		err = nil
		return
	}
	if err != nil {
		err = fmt.Errorf("can't check if config file '%s' exists: %w", file, err)
		return
	}
	// #nosec G304
	data, err := ioutil.ReadFile(file)
	if err != nil {
		err = fmt.Errorf("can't read config file '%s': %w", file, err)
		return
	}
	cfg = new(Config)
	err = json.Unmarshal(data, cfg)
	if err != nil {
		err = fmt.Errorf("can't parse config file '%s': %w", file, err)
		return
	}
	return
}

// Remove removes the configuration file.
func Remove() error {
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

// Location returns the location of the configuration file.
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
		c.SetAccessToken(accessTk)
	}
	if refreshTkChanged {
		c.SetRefreshToken(refreshTk)
	}

	if !accessTkChanged && refreshTkChanged {
		return conn, nil
	}

	err = Save(c)
	if err != nil {
		return nil, fmt.Errorf("Unable to save config file: %w", err)
	}

	return conn, nil
}
