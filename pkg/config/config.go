package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/mitchellh/go-homedir"

	sdk "github.com/openshift-online/ocm-sdk-go"
)

// Config is the type used to track the config of the client
type Config struct {
	AccessToken  string           `json:"access_token,omitempty" doc:"Bearer access token."`
	RefreshToken string           `json:"refresh_token,omitempty" doc:"Offline or refresh token."`
	TokenURL     string           `json:"token_url,omitempty" doc:"OpenID token URL."`
	Services     ServiceConfigMap `json:"services,omitempty"`
	URL          string           `json:"url,omitempty" doc:"URL of the API gateway. The value can be the complete URL or an alias. The valid aliases are 'production', 'staging' and 'integration'."`
}

// ServiceConfigMap is a map of configs for the managed application services
type ServiceConfigMap struct {
	Kafka KafkaConfig `json:"kafka,omitempty"`
}

// KafkaConfig is the config for the managed Kafka service
type KafkaConfig struct {
	ClusterID string `json:"clusterId"`
}

// SetKafka sets the current Kafka cluster
func (s *ServiceConfigMap) SetKafka(k *KafkaConfig) {
	s.Kafka = *k
}

// Save saves the given configuration to the configuration file.
func Save(cfg *Config) error {
	file, err := Location()
	if err != nil {
		return err
	}
	data, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		return fmt.Errorf("can't marshal config: %v", err)
	}
	err = ioutil.WriteFile(file, data, 0600)
	if err != nil {
		return fmt.Errorf("can't write file '%s': %v", file, err)
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
		err = fmt.Errorf("can't check if config file '%s' exists: %v", file, err)
		return
	}
	// #nosec G304
	data, err := ioutil.ReadFile(file)
	if err != nil {
		err = fmt.Errorf("can't read config file '%s': %v", file, err)
		return
	}
	cfg = new(Config)
	err = json.Unmarshal(data, cfg)
	if err != nil {
		err = fmt.Errorf("can't parse config file '%s': %v", file, err)
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
	if rhmasConfig := os.Getenv("RHMASCLI_CONFIG"); rhmasConfig != "" {
		path = rhmasConfig
	} else {
		home, err := homedir.Dir()
		if err != nil {
			return "", err
		}
		path = filepath.Join(home, ".rhmascli.json")
	}
	return path, nil
}

// Connection creates a connection using this configuration.
func (c *Config) Connection() (connection *sdk.Connection, err error) {
	if err != nil {
		return
	}

	builder := sdk.NewConnectionBuilder()
	if c.TokenURL != "" {
		builder.TokenURL(c.TokenURL)
	}

	// TODO read these from CLI
	builder.Client(sdk.DefaultClientID, sdk.DefaultClientSecret)
	builder.Scopes(sdk.DefaultScopes...)
	builder.URL(c.URL)

	tokens := make([]string, 0, 2)
	if c.AccessToken != "" {
		tokens = append(tokens, c.AccessToken)
	}
	if c.RefreshToken != "" {
		tokens = append(tokens, c.RefreshToken)
	}
	if len(tokens) > 0 {
		builder.Tokens(tokens...)
	}
	// disable TLS certification verification for now.
	builder.Insecure(true)

	// Create the connection:
	connection, err = builder.Build()
	if err != nil {
		return
	}

	return
}

// CheckTokenValidity checks if the configuration contains either credentials or tokens that haven't expired, so
// that it can be used to perform authenticated requests.
func (c *Config) CheckTokenValidity() (tokenIsValid bool, err error) {
	now := time.Now()
	if c.AccessToken != "" {
		var expires bool
		var left time.Duration
		var accessToken *jwt.Token
		accessToken, err = parseToken(c.AccessToken)
		if err != nil {
			return
		}
		expires, left, err = sdk.GetTokenExpiry(accessToken, now)
		if err != nil {
			return
		}
		if !expires || left > 5*time.Second {
			tokenIsValid = true
			return
		}
	}
	if c.RefreshToken != "" {
		var expires bool
		var left time.Duration
		var refreshToken *jwt.Token
		refreshToken, err = parseToken(c.RefreshToken)
		if err != nil {
			return
		}
		expires, left, err = sdk.GetTokenExpiry(refreshToken, now)
		if err != nil {
			return
		}
		if !expires || left > 10*time.Second {
			tokenIsValid = true
			return
		}
	}
	return
}

func parseToken(textToken string) (token *jwt.Token, err error) {
	parser := new(jwt.Parser)
	token, _, err = parser.ParseUnverified(textToken, jwt.MapClaims{})
	if err != nil {
		err = fmt.Errorf("can't parse token: %v", err)
		return
	}
	return token, nil
}
