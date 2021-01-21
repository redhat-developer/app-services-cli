package connection

import (
	"context"
	"crypto/x509"
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/api"

	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/api/serviceapi"

	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/logging"

	"github.com/Nerzal/gocloak/v7"

	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/auth/token"
)

// Default values:
const (
	// #nosec G101
	DefaultAuthURL  = "https://sso.redhat.com/auth/realms/redhat-external"
	DefaultClientID = "rhoas-cli-prod"
	DefaultURL      = "https://api.openshift.com"
)

var DefaultScopes = []string{
	"openid",
}

// KeycloakConnection contains the data needed to connect to the `api.openshift.com`. Don't create instances
// of this type directly, use the builder instead
type KeycloakConnection struct {
	trustedCAs *x509.CertPool
	insecure   bool
	client     *http.Client
	clientID   string
	Token      *token.Token
	scopes     []string
	keycloak   gocloak.GoCloak
	apiURL     *url.URL
	logger     logging.Logger
}

// RefreshTokens will fetch a refreshed copy of the access token and refresh token from the authentication server
// The new tokens will have an increased expiry time and are persisted in the config and connection
func (c *KeycloakConnection) RefreshTokens(ctx context.Context) (accessToken string, refreshToken string, err error) {
	c.logger.Debug("Refreshing access tokens")
	refreshedTk, err := c.keycloak.RefreshToken(ctx, c.Token.RefreshToken, c.clientID, "", "redhat-external")
	if err != nil {
		return "", "", &AuthError{err, ""}
	}

	if refreshedTk.AccessToken != c.Token.AccessToken {
		c.Token.AccessToken = refreshedTk.AccessToken
	}
	if refreshedTk.RefreshToken != c.Token.RefreshToken {
		c.Token.RefreshToken = refreshedTk.RefreshToken
	}

	c.logger.Debug("Access tokens successfully refreshed.")
	if c.logger.DebugEnabled() {
		b, _ := json.Marshal(c.Token)
		c.logger.Debug(string(b))
	}

	return refreshedTk.AccessToken, refreshedTk.RefreshToken, nil
}

// Logout logs the user out from the authentication server
// Invalidating and removing the access and refresh tokens
// The user will have to log in again to access the API
func (c *KeycloakConnection) Logout(ctx context.Context) error {
	err := c.keycloak.Logout(ctx, c.clientID, "", "redhat-external", c.Token.RefreshToken)
	if err != nil {
		return &AuthError{err, ""}
	}

	c.Token.AccessToken = ""
	c.Token.RefreshToken = ""

	return nil
}

// API Creates a new API type which is a single type for multiple APIs
func (c *KeycloakConnection) API() *api.API {
	kafkaAPIClient := serviceapi.New(c.apiURL, c.client, c.Token.AccessToken)

	a := &api.API{
		Kafka: kafkaAPIClient.DefaultApi,
	}

	return a
}
