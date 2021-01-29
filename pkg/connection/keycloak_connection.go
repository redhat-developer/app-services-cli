package connection

import (
	"context"
	"crypto/x509"
	"fmt"
	"net/http"
	"net/url"
	"time"

	kasclient "github.com/bf2fc6cc711aee1a0c2a/cli/pkg/api/kas/client"
	strimziadminclient "github.com/bf2fc6cc711aee1a0c2a/cli/pkg/api/strimzi-admin/client"

	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/api"

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
	trustedCAs        *x509.CertPool
	insecure          bool
	defaultHTTPClient *http.Client
	clientID          string
	Token             *token.Token
	scopes            []string
	keycloak          gocloak.GoCloak
	apiURL            *url.URL
	logger            logging.Logger
}

// RefreshTokens will fetch a refreshed copy of the access token and refresh token from the authentication server
// The new tokens will have an increased expiry time and are persisted in the config and connection
func (c *KeycloakConnection) RefreshTokens(ctx context.Context) (accessToken string, refreshToken string, err error) {
	c.logger.Debug("Refreshing access tokens")

	if !c.tokenNeedsRefresh() {
		return c.Token.AccessToken, c.Token.RefreshToken, nil
	}

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

func (c *KeycloakConnection) tokenNeedsRefresh() bool {
	t := time.Now()
	expires, left, err := token.GetExpiry(c.Token.AccessToken, t)
	if err != nil {
		c.logger.Debug("Error while checking token expiry:", err)
	}

	if !expires || left > 5*time.Minute {
		c.logger.Debug("Token is still valid.", "Expires in", left)
		return false
	}

	return true
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
	cfg.AddDefaultHeader("X-API-OpenShift-Com-Token", c.Token.AccessToken)

	apiClient := strimziadminclient.NewAPIClient(cfg)

	return apiClient
}
