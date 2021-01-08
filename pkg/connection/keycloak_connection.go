package connection

import (
	"context"
	"crypto/x509"
	"fmt"
	"net/http"
	"net/url"
	"sync"

	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/api/managedservices"
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
	tokenMutex *sync.Mutex
}

// RefreshTokens will fetch a refreshed copy of the access token and refresh token from the authentication server
// The new tokens will have an increased expiry time and are persisted in the config and connection
func (c *KeycloakConnection) RefreshTokens(ctx context.Context) (accessToken string, refreshToken string, err error) {
	// ensure this method is not executed concurrently,
	// as multiple attributes of the connection are updated
	c.tokenMutex.Lock()
	defer c.tokenMutex.Unlock()

	c.logger.Debug("Refreshing access tokens")
	refreshedTk, err := c.keycloak.RefreshToken(ctx, c.Token.RefreshToken, c.clientID, "", "redhat-external")
	if err != nil {
		return "", "", &AuthError{err, ""}
	}
	c.logger.Debug("Access tokens successfully refreshed")

	if refreshedTk.AccessToken != c.Token.AccessToken {
		c.Token.AccessToken = refreshedTk.AccessToken
	}
	if refreshedTk.RefreshToken != c.Token.RefreshToken {
		c.Token.RefreshToken = refreshedTk.RefreshToken
	}

	return refreshedTk.AccessToken, refreshedTk.RefreshToken, nil
}

// Logout logs the user out from the authentication server
// Invalidating and removing the access and refresh tokens
// The user will have to log in again to access the API
func (c *KeycloakConnection) Logout(ctx context.Context) error {
	// ensure this method is not executed concurrently,
	// as multiple attributes of the connection are updated
	c.tokenMutex.Lock()
	defer c.tokenMutex.Unlock()

	err := c.keycloak.Logout(ctx, c.clientID, "", "redhat-external", c.Token.RefreshToken)
	if err != nil {
		return &AuthError{err, ""}
	}

	c.Token.AccessToken = ""
	c.Token.RefreshToken = ""

	return nil
}

// NewAPIClient creates a new Managed Services API Client
// The current access token is passed as a Bearer token in the request
// to authorize the request
func (c *KeycloakConnection) NewAPIClient() *managedservices.APIClient {
	masCfg := managedservices.NewConfiguration()

	masCfg.Scheme = c.apiURL.Scheme
	masCfg.Host = c.apiURL.Host

	masCfg.HTTPClient = c.client

	masCfg.AddDefaultHeader("Authorization", fmt.Sprintf("Bearer %s", c.Token.AccessToken))

	masClient := managedservices.NewAPIClient(masCfg)

	return masClient
}
