package config

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"time"

	"github.com/Nerzal/gocloak/v7"
	"github.com/dgrijalva/jwt-go"

	"github.com/mitchellh/go-homedir"
)

const (
	AuthURL = "https://sso.redhat.com/auth/realms/redhat-external"
)

// Config is the type used to track the config of the client
type Config struct {
	AccessToken  string           `json:"access_token,omitempty" doc:"Bearer access token."`
	RefreshToken string           `json:"refresh_token,omitempty" doc:"Offline or refresh token."`
	Services     ServiceConfigMap `json:"services,omitempty"`
	URL          string           `json:"url,omitempty" doc:"URL of the API gateway. The value can be the complete URL or an alias. The valid aliases are 'production', 'staging' and 'integration'."`
	ClientID     string           `json:"client_id,omitempty" doc:"OpenID client identifier."`
	Insecure     bool             `json:"insecure,omitempty" doc:"Enables insecure communication with the server. This disables verification of TLS certificates and host names."`
	Scopes       []string         `json:"scopes,omitempty" doc:"OpenID scope. If this option is used it will replace completely the default scopes. Can be repeated multiple times to specify multiple scopes."`
}

// ServiceConfigMap is a map of configs for the managed application services
type ServiceConfigMap struct {
	Kafka *KafkaConfig `json:"kafka,omitempty"`
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

func (c *Config) CreateHTTPClient() *http.Client {
	// #nosec 402
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: c.Insecure},
	}
	return &http.Client{Transport: tr}
}

func (c *Config) Armed() (tokenIsValid bool, err error) {
	now := time.Now()
	if c.AccessToken != "" {
		var expires bool
		var left time.Duration
		var accessToken *jwt.Token
		accessToken, err = parseToken(c.AccessToken)
		if err != nil {
			return
		}
		expires, left, err = getTokenExpiry(accessToken, now)
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
		expires, left, err = getTokenExpiry(refreshToken, now)
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

func (c *Config) Logout() error {
	client := c.NewClient()
	err := client.Logout(context.TODO(), c.ClientID, "", "redhat-external", c.RefreshToken)
	if err != nil {
		return err
	}

	return nil
}

func (c *Config) TokenRefresh() error {
	client := c.NewClient()
	tk, err := client.RefreshToken(context.TODO(), c.RefreshToken, c.ClientID, "", "redhat-external")
	if err != nil {
		return err
	}

	c.SetAccessToken(tk.AccessToken)
	c.SetRefreshToken(tk.RefreshToken)

	return nil
}

// Create a new Keycloak client
func (c *Config) NewClient() gocloak.GoCloak {
	authURL, _ := url.Parse(AuthURL)
	authURLBase, _ := url.Parse(authURL.Scheme + "://" + authURL.Host)
	client := gocloak.NewClient(authURLBase.String())
	restyClient := *client.RestyClient()
	// #nosec 402
	restyClient.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: c.Insecure})
	client.SetRestyClient(&restyClient)
	return client
}

func parseToken(textToken string) (token *jwt.Token, err error) {
	parser := new(jwt.Parser)
	token, _, err = parser.ParseUnverified(textToken, jwt.MapClaims{})
	if err != nil {
		err = fmt.Errorf("can't parse token: %w", err)
		return
	}
	return token, nil
}

func getTokenExpiry(token *jwt.Token, now time.Time) (expires bool,
	left time.Duration, err error) {
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		err = fmt.Errorf("expected map claims bug got %T", claims)
		return
	}
	var exp float64
	claim, ok := claims["exp"]
	if ok {
		exp, ok = claim.(float64)
		if !ok {
			err = fmt.Errorf("expected floating point 'exp' but got %T", claim)
			return
		}
	}
	if exp == 0 {
		expires = false
		left = 0
	} else {
		expires = true
		left = time.Unix(int64(exp), 0).Sub(now)
	}

	return
}
