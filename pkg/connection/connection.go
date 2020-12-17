package connection

import (
	"context"
	"crypto/x509"
	"fmt"
	"net/http"
	"net/url"

	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/api/managedservices"

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

// Connection contains the data needed to connect to the `api.openshift.com`. Don't create instances
// of this type directly, use the builder instead
type Connection struct {
	trustedCAs *x509.CertPool
	insecure   bool
	client     *http.Client
	clientID   string
	Token      *token.Token
	scopes     []string
	authClient gocloak.GoCloak
	apiURL     *url.URL
}

type IConnection interface {
	RefreshTokens(ctx context.Context) (string, string, error)
	Logout(ctx context.Context) error
	HTTPClient() *http.Client
	NewMASClient() *managedservices.APIClient
}

func (c *Connection) RefreshTokens(ctx context.Context) (accessToken string, refreshToken string, err error) {
	refreshedTk, err := c.authClient.RefreshToken(ctx, c.Token.RefreshToken, c.clientID, "", "redhat-external")
	if err != nil {
		return "", "", err
	}

	if refreshedTk.AccessToken != c.Token.AccessToken {
		c.Token.AccessToken = refreshedTk.AccessToken
	}
	if refreshedTk.RefreshToken != c.Token.RefreshToken {
		c.Token.RefreshToken = refreshedTk.RefreshToken
	}

	return refreshedTk.AccessToken, refreshedTk.RefreshToken, nil
}

func (c *Connection) Logout(ctx context.Context) error {
	err := c.authClient.Logout(ctx, c.clientID, "", "redhat-external", c.Token.RefreshToken)
	if err != nil {
		return err
	}

	c.Token.AccessToken = ""
	c.Token.RefreshToken = ""

	return nil
}

func (c *Connection) HTTPClient() *http.Client {
	return c.client
}

func (c *Connection) NewMASClient() *managedservices.APIClient {
	masCfg := managedservices.NewConfiguration()

	masCfg.Scheme = c.apiURL.Scheme
	masCfg.Host = c.apiURL.Host

	masCfg.HTTPClient = c.HTTPClient()

	masCfg.AddDefaultHeader("Authorization", fmt.Sprintf("Bearer %s", c.Token.AccessToken))

	masClient := managedservices.NewAPIClient(masCfg)

	return masClient
}
