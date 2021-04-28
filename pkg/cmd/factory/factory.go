package factory

import (
	"github.com/redhat-developer/app-services-cli/internal/config"
	"github.com/redhat-developer/app-services-cli/pkg/connection"
	"github.com/redhat-developer/app-services-cli/pkg/iostreams"
	"github.com/redhat-developer/app-services-cli/pkg/locales"
	"github.com/redhat-developer/app-services-cli/pkg/logging"
)

// Factory is an abstract type which provides access to
// the root configuration and connections for the CLI
type Factory struct {
	// Type which defines the streams for the CLI
	IOStreams *iostreams.IOStreams
	// Interface to read/write to the config
	Config config.IConfig
	// Creates a connection to the API
	Connection ConnectionFunc
	// Returns a logger to create leveled logs in the application
	Logger func() (logging.Logger, error)
	// Localizer provides text to the commands
	Localizer locales.Localizer
}

type ConnectionFunc func(cfg *connection.Config) (connection.Connection, error)
