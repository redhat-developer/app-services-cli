package connection

import (
	"context"

	"github.com/redhat-developer/app-services-cli/pkg/api"
)

type Config struct {
	RequireAuth    bool
	RequireMASAuth bool
}

// DefaultConfigSkipMasAuth is used when running commands which do  not require authenticatation with MAS-SSO
var DefaultConfigSkipMasAuth = &Config{
	RequireAuth:    true,
	RequireMASAuth: false,
}

// DefaultConfigRequireMasAuth is used when running commands which must authenticate with MAS-SSO
var DefaultConfigRequireMasAuth = &Config{
	RequireAuth:    true,
	RequireMASAuth: true,
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
	API() *api.API
}
