package connection

import (
	"context"
	"crypto/x509"
	"errors"
	"fmt"
	"net"
	"net/http"
	"net/url"
	"strings"

	"github.com/bf2fc6cc711aee1a0c2a/cli/internal/config"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/api/kas"

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
	// MAS SSO defaults
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
	defaultRealm      string
	masRealm          string
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
	// nolint:govet
	refreshedTk, err := c.keycloakClient.RefreshToken(ctx, c.Token.RefreshToken, c.clientID, "", c.defaultRealm)
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

	// nolint:govet
	refreshedMasTk, err := c.masKeycloakClient.RefreshToken(ctx, c.MASToken.RefreshToken, c.clientID, "", c.masRealm)
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
	err = c.keycloakClient.Logout(ctx, c.clientID, "", c.defaultRealm, c.Token.RefreshToken)
	if err != nil {
		return &AuthError{err, ""}
	}

	err = c.masKeycloakClient.Logout(ctx, c.clientID, "", c.masRealm, c.MASToken.RefreshToken)
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
// nolint:funlen
func (c *KeycloakConnection) API() *api.API {
	var cachedKafkaServiceAPI kasclient.DefaultApi
	var cachedKafkaID string
	var cachedKafkaAdminAPI strimziadminclient.DefaultApi
	var cachedKafkaRequest *kasclient.KafkaRequest
	var cachedKafkaAdminErr error

	kafkaAPIFunc := func() kasclient.DefaultApi {
		if cachedKafkaServiceAPI != nil {
			return cachedKafkaServiceAPI
		}

		// create the client
		kafkaAPIClient := c.createKafkaAPIClient()

		cachedKafkaServiceAPI = kafkaAPIClient.DefaultApi

		return cachedKafkaServiceAPI
	}

	kafkaAdminAPIFunc := func(kafkaID string) (strimziadminclient.DefaultApi, *kasclient.KafkaRequest, error) {
		// if the api client is already created, and the same Kafka ID is used
		// return the cached client
		if cachedKafkaAdminAPI != nil && kafkaID == cachedKafkaID {
			return cachedKafkaAdminAPI, cachedKafkaRequest, cachedKafkaAdminErr
		}

		api := kafkaAPIFunc()

		kafkaInstance, resp, apiErr := api.GetKafkaById(context.Background(), kafkaID).Execute()
		defer resp.Body.Close()
		if kas.IsErr(apiErr, kas.ErrorNotFound) {
			cachedKafkaAdminAPI = nil
			cachedKafkaRequest = nil
			cachedKafkaAdminErr = errors.New(localizer.MustLocalize(&localizer.Config{
				MessageID: "kafka.common.error.notFoundByIdError",
				TemplateData: map[string]interface{}{
					"ID": kafkaID,
				},
			}))

			return cachedKafkaAdminAPI, cachedKafkaRequest, cachedKafkaAdminErr
		} else if apiErr.Error() != "" {
			cachedKafkaAdminAPI = nil
			cachedKafkaRequest = nil
			cachedKafkaAdminErr = fmt.Errorf("%v", apiErr)

			return cachedKafkaAdminAPI, cachedKafkaRequest, cachedKafkaAdminErr
		}

		cachedKafkaRequest = &kafkaInstance
		// cache the Kafka ID
		cachedKafkaID = kafkaID

		kafkaStatus := kafkaInstance.GetStatus()
		if kafkaStatus != "ready" {
			cachedKafkaAdminAPI = nil
			cachedKafkaRequest = nil
			cachedKafkaAdminErr = errors.New(localizer.MustLocalize(&localizer.Config{
				MessageID: "kafka.common.error.notReadyError",
				TemplateData: map[string]interface{}{
					"Name": kafkaInstance.GetName(),
				},
			}))

			return cachedKafkaAdminAPI, cachedKafkaRequest, cachedKafkaAdminErr
		}

		bootstrapURL := kafkaInstance.GetBootstrapServerHost()
		if bootstrapURL == "" {
			cachedKafkaAdminAPI = nil
			cachedKafkaRequest = nil
			cachedKafkaAdminErr = fmt.Errorf(localizer.MustLocalize(&localizer.Config{
				MessageID: "connection.error.missingBootstrapURL",
				TemplateData: map[string]interface{}{
					"Name": kafkaInstance.GetName(),
				},
			}))

			return cachedKafkaAdminAPI, cachedKafkaRequest, cachedKafkaAdminErr
		}

		// create the client
		apiClient := c.createKafkaAdminAPI(bootstrapURL)

		cachedKafkaAdminAPI = apiClient.DefaultApi
		cachedKafkaAdminErr = nil

		return cachedKafkaAdminAPI, cachedKafkaRequest, cachedKafkaAdminErr
	}

	return &api.API{
		Kafka:      kafkaAPIFunc,
		TopicAdmin: kafkaAdminAPIFunc,
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

// Create a new KafkaAdmin API client
func (c *KeycloakConnection) createKafkaAdminAPI(bootstrapURL string) *strimziadminclient.APIClient {
	cfg := strimziadminclient.NewConfiguration()

	host, _, _ := net.SplitHostPort(bootstrapURL)

	cfg.Scheme = "https"
	cfg.Host = fmt.Sprintf("admin-server-%v", host)
	c.logger.Debugf("Making request to %v://%v", cfg.Scheme, cfg.Host)

	cfg.HTTPClient = c.defaultHTTPClient

	cfg.AddDefaultHeader("Authorization", c.MASToken.AccessToken)

	apiClient := strimziadminclient.NewAPIClient(cfg)

	return apiClient
}

// get the realm from the Keycloak URL
func getKeycloakRealm(url *url.URL) (realm string, ok bool) {
	parts := strings.Split(url.Path, "/")
	for i, part := range parts {
		if part == "realms" {
			realm = parts[i+1]
			ok = true
		}
	}

	return realm, ok
}
