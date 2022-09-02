package config

// IConfig is an interface which describes the functions
// needed to read/write from a config
//
//go:generate moq -out ./config_mock.go . IConfig
type IConfig interface {
	Load() (*Config, error)
	Save(config *Config) error
	Remove() error
	Location() (string, error)
}

// Config is a type which describes the properties which can be in the config
type Config struct {
	AccessToken  string           `json:"access_token,omitempty" doc:"Bearer access token."`
	RefreshToken string           `json:"refresh_token,omitempty" doc:"Offline or refresh token."`
	Services     ServiceConfigMap `json:"services,omitempty"`
	APIUrl       string           `json:"api_url,omitempty" doc:"URL of the API gateway. The value can be the complete URL or an alias. The valid aliases are 'production', 'staging' and 'integration'."`
	AuthURL      string           `json:"auth_url,omitempty" doc:"URL of the authentication server"`
	ClientID     string           `json:"client_id,omitempty" doc:"OpenID client identifier."`
	Insecure     bool             `json:"insecure,omitempty" doc:"Enables insecure communication with the server. This disables verification of TLS certificates and host names."`
	Scopes       []string         `json:"scopes,omitempty" doc:"OpenID scope. If this option is used it will replace completely the default scopes. Can be repeated multiple times to specify multiple scopes."`
	Telemetry    string           `json:"telemetry,omitempty" doc:"Flag used to enable telemetry for user."`
	LastUpdated  int64            `json:"last_updated,omitempty" doc:"Timestamp of the last update cli"`
	EnableAuthV2 bool             `json:"enable_auth_v2,omitempty" doc:"Enables use of new Service Account SDK"`
}

// ServiceConfigMap is a map of configs for the application services
type ServiceConfigMap struct {
	Kafka           *KafkaConfig           `json:"kafka,omitempty"`
	ServiceRegistry *ServiceRegistryConfig `json:"serviceregistry,omitempty"`
}

// KafkaConfig is the config for the Kafka service
type KafkaConfig struct {
	ClusterID string `json:"clusterId"`
}

type ServiceRegistryConfig struct {
	InstanceID string `json:"instanceId"`
	Name       string `json:"name"`
}

// GetKafkaIdOk returns the current Kafka instance ID and whether it exists
func (c *Config) GetKafkaIdOk() (string, bool) {
	if c.Services.Kafka != nil && c.Services.Kafka.ClusterID != "" {
		return c.Services.Kafka.ClusterID, true
	}

	return "", false
}

// GetServiceRegistryIdOk returns the service registry instance ID and whether it exists or not
func (c *Config) GetServiceRegistryIdOk() (string, bool) {
	if c.Services.ServiceRegistry != nil && c.Services.ServiceRegistry.InstanceID != "" {
		return c.Services.ServiceRegistry.InstanceID, true
	}

	return "", false
}
