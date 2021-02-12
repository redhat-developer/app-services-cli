package config

// IConfig is an interface which describes the functions
// needed to read/write from a config
//go:generate moq -out ./config_mock.go . IConfig
type IConfig interface {
	Load() (*Config, error)
	Save(config *Config) error
	Remove() error
	Location() (string, error)
}

// Config is a type which describes the properties which can be in the config
type Config struct {
	AccessToken  string           `json:"access_token" doc:"Bearer access token."`
	RefreshToken string           `json:"refresh_token" doc:"Offline or refresh token."`
	Services     ServiceConfigMap `json:"services"`
	APIGateway   string           `json:"api_gateway" doc:"URL of the API gateway. The value can be the complete URL or an alias. The valid aliases are 'production', 'staging' and 'integration'."`
	AuthURL      string           `json:"auth_url" doc:"URL of the authentication server"`
	ClientID     string           `json:"client_id" doc:"OpenID client identifier."`
	Insecure     bool             `json:"insecure" doc:"Enables insecure communication with the server. This disables verification of TLS certificates and host names."`
	Scopes       []string         `json:"scopes" doc:"OpenID scope. If this option is used it will replace completely the default scopes. Can be repeated multiple times to specify multiple scopes."`
	ServiceAuth  ServiceAuth      `json:"serviceAuth"`
}

// ServiceAuth for cli authentication within enabled services
type ServiceAuth struct {
	ClientID     string `json:"clientID"`
	ClientSecret string `json:"clientSecret"`
}

// ServiceConfigMap is a map of configs for the managed application services
type ServiceConfigMap struct {
	Kafka *KafkaConfig `json:"kafka"`
}

// KafkaConfig is the config for the managed Kafka service
type KafkaConfig struct {
	ClusterID string `json:"clusterId"`
}

func (c *Config) HasKafka() bool {
	return c.Services.Kafka != nil &&
		c.Services.Kafka.ClusterID != ""
}
