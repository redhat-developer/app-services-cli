package kcconnection

import (
	"context"
	"crypto/x509"
	"net/http"
	"net/url"

	"github.com/redhat-developer/app-services-cli/pkg/core/auth/token"

	"github.com/redhat-developer/app-services-cli/pkg/core/config"
	"github.com/redhat-developer/app-services-cli/pkg/core/logging"
	"github.com/redhat-developer/app-services-cli/pkg/shared/connection"
	"github.com/redhat-developer/app-services-cli/pkg/shared/connection/api"
	"github.com/redhat-developer/app-services-cli/pkg/shared/connection/api/defaultapi"

	"github.com/redhat-developer/app-services-cli/internal/build"

	"github.com/Nerzal/gocloak/v7"
)

var DefaultScopes = []string{
	"openid",
}

// Connection contains the data needed to connect to the `api.openshift.com`. Don't create instances
// of this type directly, use the builder instead
type Connection struct {
	trustedCAs        *x509.CertPool
	insecure          bool
	defaultHTTPClient *http.Client
	clientID          string
	Token             *token.Token
	MASToken          *token.Token
	scopes            []string
	keycloakClient    gocloak.GoCloak
	masKeycloakClient gocloak.GoCloak
	apiURL            *url.URL
	consoleURL        *url.URL
	defaultRealm      string
	masRealm          string
	logger            logging.Logger
	Config            config.IConfig
	connectionConfig  *connection.Config
}

// RefreshTokens will fetch a refreshed copy of the access token and refresh token from the authentication server
// The new tokens will have an increased expiry time and are persisted in the config and connection
func (c *Connection) RefreshTokens(ctx context.Context) (err error) {

	cfg, err := c.Config.Load()
	if err != nil {
		return err
	}

	// track if we need to update the config with new token values
	var cfgChanged bool
	if c.connectionConfig.RequireAuth {
		c.logger.Debug("Refreshing tokens")
		// nolint:govet
		refreshedTk, err := c.keycloakClient.RefreshToken(ctx, c.Token.RefreshToken, c.clientID, "", c.defaultRealm)
		if err != nil {
			return &AuthError{err}
		}

		if refreshedTk.AccessToken != c.Token.AccessToken {
			c.Token.AccessToken = refreshedTk.AccessToken
			cfg.AccessToken = refreshedTk.AccessToken
			cfgChanged = true
		}
		if refreshedTk.RefreshToken != c.Token.RefreshToken {
			c.Token.RefreshToken = refreshedTk.RefreshToken
			cfg.RefreshToken = refreshedTk.RefreshToken
			cfgChanged = true
		}
	}

	if c.connectionConfig.RequireMASAuth {
		c.logger.Debug("Refreshing MAS SSO tokens")
		// nolint:govet
		refreshedMasTk, err := c.masKeycloakClient.RefreshToken(ctx, c.MASToken.RefreshToken, c.clientID, "", c.masRealm)
		if err != nil {
			return &MasAuthError{err}
		}
		if refreshedMasTk.AccessToken != c.MASToken.AccessToken {
			c.MASToken.AccessToken = refreshedMasTk.AccessToken
			cfg.MasAccessToken = refreshedMasTk.AccessToken
			cfgChanged = true
		}
		if refreshedMasTk.RefreshToken != c.MASToken.RefreshToken {
			c.MASToken.RefreshToken = refreshedMasTk.RefreshToken
			cfg.MasRefreshToken = refreshedMasTk.RefreshToken
			cfgChanged = true
		}
	}

	if !cfgChanged {
		return nil
	}

	if err = c.Config.Save(cfg); err != nil {
		return err
	}
	c.logger.Debug("Tokens refreshed")

	return nil
}

// Logout logs the user out from the authentication server
// Invalidating and removing the access and refresh tokens
// The user will have to log in again to access the API
func (c *Connection) Logout(ctx context.Context) (err error) {
	err = c.keycloakClient.Logout(ctx, c.clientID, "", c.defaultRealm, c.Token.RefreshToken)
	if err != nil {
		return &AuthError{err}
	}

	if c.MASToken.RefreshToken != "" {
		err = c.masKeycloakClient.Logout(ctx, c.clientID, "", c.masRealm, c.MASToken.RefreshToken)
		if err != nil {
			return &AuthError{err}
		}
	}

	c.Token.AccessToken = ""
	c.Token.RefreshToken = ""
	c.MASToken.AccessToken = ""
	c.MASToken.RefreshToken = ""

	cfg, err := c.Config.Load()
	if err != nil {
		return err
	}

	cfg.AccessToken = ""
	cfg.RefreshToken = ""
	cfg.MasAccessToken = ""
	cfg.MasRefreshToken = ""

	return c.Config.Save(cfg)
}

// API Creates a new API type which is a single type for multiple APIs
func (c *Connection) API() api.API {
	apiClient := defaultapi.New(&defaultapi.Config{
		HTTPClient:     c.defaultHTTPClient,
		UserAgent:      build.DefaultUserAgentPrefix + build.Version,
		MasAccessToken: c.MASToken.AccessToken,
		AccessToken:    c.Token.AccessToken,
		ApiURL:         c.apiURL,
		ConsoleURL:     c.consoleURL,
		Logger:         c.logger,
	})

	return apiClient
}
