package connection

import (
	"context"

	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/api"
)

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
