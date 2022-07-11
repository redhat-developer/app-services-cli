package connection

import (
	"context"

	"github.com/redhat-developer/app-services-cli/pkg/shared/connection/api"
)

// FIXLATER This entire class can be removed however it should be done
// after connectors commands are finished thus we do not have many conflicts.
type Config struct {
	RequireAuth bool
}

// DefaultConfigSkipMasAuth is used when running all commads
var DefaultConfigSkipMasAuth = &Config{
	RequireAuth: true,
}

var DefaultConfigRequireMasAuth = &Config{
	RequireAuth: true,
}

// Connection is an interface which defines methods for interacting
// with the control plane API and the authentication server
//go:generate moq -out connection_mock.go . Connection
type Connection interface {
	// Method to refresh the OAuth tokens
	RefreshTokens(ctx context.Context) error
	// Method to perform a logout request to the authentication server
	Logout(ctx context.Context) error
	// Method to create the API clients
	API() api.API
}
