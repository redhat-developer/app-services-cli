package mock

import (
	"context"
	"fmt"

	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/api/managedservices"
)

type Connection struct {
	AccessToken  string
	RefreshToken string
}

func (c *Connection) RefreshTokens(ctx context.Context) (string, string, error) {
	if c.RefreshToken == "expired" {
		return "", "", fmt.Errorf("Refresh token already expired")
	}

	return "valid", "valid", nil
}

func (c *Connection) Logout(ctx context.Context) error {
	if c.AccessToken == "expired" && c.RefreshToken == "expired" {
		return fmt.Errorf("Could not log out as tokens are expired")
	}

	return nil
}

func (c *Connection) NewMASClient() *managedservices.APIClient {
	return nil
}
