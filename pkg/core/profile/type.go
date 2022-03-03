package profile

type Context struct {
	Contexts       map[string]ServiceConfig `json:"contexts,omitempty"`
	CurrentContext string                   `json:"current_context,omitempty"`
}

// ServiceConfig is a map of identifiers for the application services
type ServiceConfig struct {
	KafkaID           string `json:"kafkaID"`
	ServiceRegistryID string `json:"serviceregistryID"`
}

type IContext interface {
	Load() (*Context, error)
	Save(config *Context) error
	Remove() error
	Location() (string, error)
}
