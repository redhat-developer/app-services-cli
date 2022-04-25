package servicecontext

// Context is a type which describes the properties of context file
type Context struct {
	Contexts       map[string]ServiceConfig `json:"contexts,omitempty"`
	CurrentContext string                   `json:"current_context"`
}

// ServiceConfig is a map of identifiers for the application services
type ServiceConfig struct {
	KafkaID           string `json:"kafkaID"`
	ServiceRegistryID string `json:"serviceregistryID"`
}

// IContext is an interface which describes functions for context file
type IContext interface {
	Load() (*Context, error)
	Save(*Context) error
	Remove() error
	Location() (string, error)
}
