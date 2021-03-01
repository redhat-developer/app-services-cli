package connection

import (
	"context"
	"crypto/x509"
	"fmt"
	"net/http"
	"net/url"

	"github.com/bf2fc6cc711aee1a0c2a/cli/internal/config"
	"github.com/bf2fc6cc711aee1a0c2a/cli/internal/localizer"
	kasclient "github.com/bf2fc6cc711aee1a0c2a/cli/pkg/api/kas/client"
	strimziadminclient "github.com/bf2fc6cc711aee1a0c2a/cli/pkg/api/strimzi-admin/client"

	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/api"

	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/logging"

	"github.com/Nerzal/gocloak/v7"

	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/auth/token"
)

// Default values:
const (
	DefaultClientID = "rhoas-cli-prod"
	DefaultURL      = "https://api.openshift.com"

	// SSO defaults
	DefaultAuthURL = "https://sso.redhat.com/auth/realms/redhat-external"
	DefaultRealm   = "redhat-external"
	// MAS SSO defaults
	DefaultMASRealm   = "mas-sso-staging"
	DefaultMasAuthURL = "https://keycloak-edge-redhat-rhoam-user-sso.apps.mas-sso-stage.1gzl.s1.devshift.org/auth/realms/mas-sso-staging"
)

var DefaultScopes = []string{
	"openid",
}

// KeycloakConnection contains the data needed to connect to the `api.openshift.com`. Don't create instances
// of this type directly, use the builder instead
type KeycloakConnection struct {
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
	logger            logging.Logger
	Config            config.IConfig
}

// RefreshTokens will fetch a refreshed copy of the access token and refresh token from the authentication server
// The new tokens will have an increased expiry time and are persisted in the config and connection
func (c *KeycloakConnection) RefreshTokens(ctx context.Context) (err error) {
	c.logger.Debug(localizer.MustLocalizeFromID("connection.refreshTokens.log.debug.refreshingTokens"))

	cfg, err := c.Config.Load()
	if err != nil {
		return err
	}

	// track if we need to update the config with new token values
	var cfgChanged bool
	if c.Token.NeedsRefresh() {
		// nolint:govet
		refreshedTk, err := c.keycloakClient.RefreshToken(ctx, c.Token.RefreshToken, c.clientID, "", DefaultRealm)
		if err != nil {
			return &AuthError{err, ""}
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

	if c.MASToken.NeedsRefresh() {
		// nolint:govet
		refreshedMasTk, err := c.masKeycloakClient.RefreshToken(ctx, c.MASToken.RefreshToken, c.clientID, "", DefaultMASRealm)
		if err != nil {
			return &AuthError{err, ""}
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
	c.logger.Debug(localizer.MustLocalizeFromID("connection.refreshTokens.log.debug.tokensRefreshed"))

	return nil
}

// Logout logs the user out from the authentication server
// Invalidating and removing the access and refresh tokens
// The user will have to log in again to access the API
func (c *KeycloakConnection) Logout(ctx context.Context) (err error) {
	err = c.keycloakClient.Logout(ctx, c.clientID, "", DefaultRealm, c.Token.RefreshToken)
	if err != nil {
		return &AuthError{err, ""}
	}

	err = c.masKeycloakClient.Logout(ctx, c.clientID, "", DefaultMASRealm, c.MASToken.RefreshToken)
	if err != nil {
		return &AuthError{err, ""}
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

	if err = c.Config.Save(cfg); err != nil {
		return err
	}

	return nil
}

// API Creates a new API type which is a single type for multiple APIs
func (c *KeycloakConnection) API() *api.API {
	var cachedKafkaServiceAPI kasclient.DefaultApi

	var cachedStrimziKafkaID string
	var cachedStrimziAdminAPI strimziadminclient.DefaultApi

	kafkaAPIFunc := func() kasclient.DefaultApi {
		if cachedKafkaServiceAPI != nil {
			return cachedKafkaServiceAPI
		}

		// create the client
		kafkaAPIClient := c.createKafkaAPIClient()

		cachedKafkaServiceAPI = kafkaAPIClient.DefaultApi

		return cachedKafkaServiceAPI
	}

	strimzAdminAPIFunc := func(kafkaID string) strimziadminclient.DefaultApi {
		// if the api client is already created, and the same Kafka ID is used
		// return the cached client
		if cachedStrimziAdminAPI != nil && kafkaID == cachedStrimziKafkaID {
			return cachedStrimziAdminAPI
		}

		// cache the Kafka ID
		cachedStrimziKafkaID = kafkaID

		// create the client
		apiClient := c.createStrimziAdminAPIClient(kafkaID)

		cachedStrimziAdminAPI = apiClient.DefaultApi

		return cachedStrimziAdminAPI
	}

	return &api.API{
		Kafka:      kafkaAPIFunc,
		TopicAdmin: strimzAdminAPIFunc,
	}
}

// Create a new Kafka API client
func (c *KeycloakConnection) createKafkaAPIClient() *kasclient.APIClient {
	cfg := kasclient.NewConfiguration()

	cfg.Scheme = c.apiURL.Scheme
	cfg.Host = c.apiURL.Host

	cfg.HTTPClient = c.defaultHTTPClient

	cfg.AddDefaultHeader("Authorization", fmt.Sprintf("Bearer %v", c.Token.AccessToken))

	apiClient := kasclient.NewAPIClient(cfg)

	return apiClient
}

// Create a new Strimzi Admin API client
func (c *KeycloakConnection) createStrimziAdminAPIClient(kafkaID string) *strimziadminclient.APIClient {
	cfg := strimziadminclient.NewConfiguration()

	cfg.Scheme = c.apiURL.Scheme
	cfg.Host = c.apiURL.Host

	cfg.HTTPClient = c.defaultHTTPClient

	cfg.AddDefaultHeader("X-Kafka-ID", kafkaID)
	cfg.AddDefaultHeader("Authorization", c.MASToken.AccessToken)
	fmt.Println(c.MASToken.AccessToken)

	apiClient := strimziadminclient.NewAPIClient(cfg)

	return apiClient
}
