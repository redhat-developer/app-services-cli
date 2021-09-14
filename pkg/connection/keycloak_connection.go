package connection

import (
	"context"
	"crypto/x509"
	"fmt"
	"github.com/redhat-developer/app-services-cli/internal/build"
	"net"
	"net/http"
	"net/url"

	kafkainstance "github.com/redhat-developer/app-services-sdk-go/kafkainstance/apiv1internal"
	kafkainstanceclient "github.com/redhat-developer/app-services-sdk-go/kafkainstance/apiv1internal/client"

	kafkamgmt "github.com/redhat-developer/app-services-sdk-go/kafkamgmt/apiv1"
	kafkamgmtclient "github.com/redhat-developer/app-services-sdk-go/kafkamgmt/apiv1/client"

	registryinstance "github.com/redhat-developer/app-services-sdk-go/registryinstance/apiv1internal"
	registryinstanceclient "github.com/redhat-developer/app-services-sdk-go/registryinstance/apiv1internal/client"

	registrymgmt "github.com/redhat-developer/app-services-sdk-go/registrymgmt/apiv1"
	registrymgmtclient "github.com/redhat-developer/app-services-sdk-go/registrymgmt/apiv1/client"

	"golang.org/x/oauth2"

	"github.com/redhat-developer/app-services-cli/pkg/api/ams/amsclient"
	"github.com/redhat-developer/app-services-cli/pkg/api/kas"
	"github.com/redhat-developer/app-services-cli/pkg/api/rbac"

	"github.com/redhat-developer/app-services-cli/pkg/kafka/kafkaerr"

	"github.com/redhat-developer/app-services-cli/internal/config"

	"github.com/redhat-developer/app-services-cli/pkg/api"

	"github.com/redhat-developer/app-services-cli/pkg/logging"

	"github.com/Nerzal/gocloak/v7"

	"github.com/redhat-developer/app-services-cli/pkg/auth/token"
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
	consoleURL        *url.URL
	defaultRealm      string
	masRealm          string
	logger            logging.Logger
	Config            config.IConfig
	connectionConfig  *Config
}

// RefreshTokens will fetch a refreshed copy of the access token and refresh token from the authentication server
// The new tokens will have an increased expiry time and are persisted in the config and connection
func (c *KeycloakConnection) RefreshTokens(ctx context.Context) (err error) {
	c.logger.Debug("Refreshing tokens")

	cfg, err := c.Config.Load()
	if err != nil {
		return err
	}

	// track if we need to update the config with new token values
	var cfgChanged bool
	if c.connectionConfig.RequireAuth {
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
func (c *KeycloakConnection) Logout(ctx context.Context) (err error) {
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
// nolint:funlen
func (c *KeycloakConnection) API() *api.API {
	amsAPIFunc := func() amsclient.DefaultApi {
		amsAPIClient := c.createAmsAPIClient()

		return amsAPIClient.DefaultApi
	}

	kafkaAPIFunc := func() kafkamgmtclient.DefaultApi {
		// create the client
		kafkaAPIClient := c.createKafkaAPIClient()

		return kafkaAPIClient.DefaultApi
	}

	serviceAccountAPIFunc := func() kafkamgmtclient.SecurityApi {
		apiClient := c.createKafkaAPIClient()

		return apiClient.SecurityApi
	}

	registryAPIFunc := func() registrymgmtclient.RegistriesApi {
		srsAPIClient := c.createServiceRegistryAPIClient()

		return srsAPIClient.RegistriesApi
	}

	rbacAPI := rbac.RbacAPI{
		PrincipalAPI: func() rbac.PrincipalAPI {
			cl := c.createOAuthTransport(c.Token.AccessToken)
			cfg := rbac.Config{
				HTTPClient: cl,
				Debug:      c.logger.DebugEnabled(),
				BaseURL:    c.consoleURL,
			}
			return rbac.NewPrincipalAPIClient(&cfg)
		},
	}

	kafkaAdminAPIFunc := func(kafkaID string) (*kafkainstanceclient.APIClient, *kafkamgmtclient.KafkaRequest, error) {
		kafkaAPI := kafkaAPIFunc()

		kafkaInstance, resp, err := kafkaAPI.GetKafkaById(context.Background(), kafkaID).Execute()
		defer resp.Body.Close()
		if kas.IsErr(err, kas.ErrorNotFound) {
			return nil, nil, kafkaerr.NotFoundByIDError(kafkaID)
		} else if err != nil {
			return nil, nil, fmt.Errorf("%w", err)
		}

		kafkaStatus := kafkaInstance.GetStatus()

		switch kafkaStatus {
		case "provisioning", "accepted":
			err = fmt.Errorf(`Kafka instance "%v" is not ready yet`, kafkaInstance.GetName())
			return nil, nil, err
		case "failed":
			err = fmt.Errorf(`Kafka instance "%v" has failed`, kafkaInstance.GetName())
			return nil, nil, err
		case "deprovision":
			err = fmt.Errorf(`Kafka instance "%v" is being deprovisioned`, kafkaInstance.GetName())
			return nil, nil, err
		case "deleting":
			err = fmt.Errorf(`Kafka instance "%v" is being deleted`, kafkaInstance.GetName())
			return nil, nil, err
		}

		bootstrapURL := kafkaInstance.GetBootstrapServerHost()
		if bootstrapURL == "" {
			err = fmt.Errorf(`bootstrap URL is missing for Kafka instance "%v"`, kafkaInstance.GetName())

			return nil, nil, err
		}

		// create the client
		client := c.createKafkaAdminAPI(bootstrapURL)

		return client, &kafkaInstance, nil
	}

	registryInstanceAPIFunc := func(registryID string) (*registryinstanceclient.APIClient, *registrymgmtclient.Registry, error) {
		registryAPI := registryAPIFunc()

		instance, resp, err := registryAPI.GetRegistry(context.Background(), registryID).Execute()
		defer resp.Body.Close()
		if err != nil {
			return nil, nil, fmt.Errorf("%w", err)
		}

		status := instance.GetStatus()
		// nolint
		switch status {
		case "provisioning", "accepted":
			err = fmt.Errorf(`service registry instance "%v" is not ready yet`, instance.GetName())
			return nil, nil, err
		case "failed":
			err = fmt.Errorf(`service registry instance "%v" has failed`, instance.GetName())
			return nil, nil, err
		case "deprovision":
			err = fmt.Errorf(`service registry instance "%v" is being deprovisioned`, instance.GetName())
			return nil, nil, err
		case "deleting":
			err = fmt.Errorf(`service registry instance "%v" is being deleted`, instance.GetName())
			return nil, nil, err
		}

		registryUrl := instance.GetRegistryUrl()
		if registryUrl == "" {
			err = fmt.Errorf(`URL is missing for Service Registry instance "%v"`, instance.GetName())

			return nil, nil, err
		}

		// create the client
		client := c.createServiceRegistryInstanceAPI(registryUrl)

		return client, &instance, nil
	}

	return &api.API{
		Kafka:                   kafkaAPIFunc,
		ServiceAccount:          serviceAccountAPIFunc,
		KafkaAdmin:              kafkaAdminAPIFunc,
		ServiceRegistryInstance: registryInstanceAPIFunc,
		AccountMgmt:             amsAPIFunc,
		ServiceRegistryMgmt:     registryAPIFunc,
		RBAC:                    rbacAPI,
	}
}

// Create a new Kafka API client
func (c *KeycloakConnection) createKafkaAPIClient() *kafkamgmtclient.APIClient {
	tc := c.createOAuthTransport(c.Token.AccessToken)
	client := kafkamgmt.NewAPIClient(&kafkamgmt.Config{
		BaseURL:    c.apiURL.String(),
		Debug:      c.logger.DebugEnabled(),
		HTTPClient: tc,
		UserAgent:  build.DefaultUserAgentPrefix + build.Version,
	})

	return client
}

// Create a new Registry API client
func (c *KeycloakConnection) createServiceRegistryAPIClient() *registrymgmtclient.APIClient {
	tc := c.createOAuthTransport(c.Token.AccessToken)
	client := registrymgmt.NewAPIClient(&registrymgmt.Config{
		BaseURL:    c.apiURL.String(),
		Debug:      c.logger.DebugEnabled(),
		HTTPClient: tc,
		UserAgent:  build.DefaultUserAgentPrefix + build.Version,
	})

	return client
}

// Create a new KafkaAdmin API client
func (c *KeycloakConnection) createKafkaAdminAPI(bootstrapURL string) *kafkainstanceclient.APIClient {
	host, port, _ := net.SplitHostPort(bootstrapURL)

	var apiURL *url.URL

	if host == "localhost" {
		apiURL = &url.URL{
			Scheme: "http",
			Host:   fmt.Sprintf("localhost:%v", port),
		}
		apiURL.Scheme = "http"
		apiURL.Path = "/data/kafka"
	} else {
		apiHost := fmt.Sprintf("admin-server-%v", host)
		apiURL, _ = url.Parse(apiHost)
		apiURL.Scheme = "https"
		apiURL.Path = "/rest"
		apiURL.Host = fmt.Sprintf("admin-server-%v", host)
	}

	c.logger.Debugf("Making request to %v", apiURL.String())

	client := kafkainstance.NewAPIClient(&kafkainstance.Config{
		BaseURL:    apiURL.String(),
		Debug:      c.logger.DebugEnabled(),
		HTTPClient: c.createOAuthTransport(c.MASToken.AccessToken),
		UserAgent:  build.DefaultUserAgentPrefix + build.Version,
	})

	return client
}

// Create a new RegistryInstance API client
func (c *KeycloakConnection) createServiceRegistryInstanceAPI(registryUrl string) *registryinstanceclient.APIClient {
	host, port, _ := net.SplitHostPort(registryUrl)

	var baseURL string
	if host == "localhost" {
		var apiURL = &url.URL{
			Scheme: "http",
			Host:   fmt.Sprintf("localhost:%v", port),
		}
		apiURL.Scheme = "http"
		apiURL.Path = "/data/registry"
		baseURL = apiURL.String()
	} else {
		baseURL = registryUrl + "/apis/registry/v2"
	}

	c.logger.Debugf("Making request to %v", baseURL)

	client := registryinstance.NewAPIClient(&registryinstance.Config{
		BaseURL:    baseURL,
		Debug:      c.logger.DebugEnabled(),
		HTTPClient: c.createOAuthTransport(c.MASToken.AccessToken),
		UserAgent:  build.DefaultUserAgentPrefix + build.Version,
	})

	return client
}

func (c *KeycloakConnection) createAmsAPIClient() *amsclient.APIClient {
	cfg := amsclient.NewConfiguration()

	cfg.Scheme = c.apiURL.Scheme
	cfg.Host = c.apiURL.Host
	cfg.UserAgent = build.DefaultUserAgentPrefix + build.Version

	cfg.HTTPClient = c.createOAuthTransport(c.Token.AccessToken)

	apiClient := amsclient.NewAPIClient(cfg)

	return apiClient
}

// wraps the HTTP client with an OAuth2 Transport layer to provide automatic token refreshing
func (c *KeycloakConnection) createOAuthTransport(accessToken string) *http.Client {
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{
			AccessToken: accessToken,
		},
	)

	return &http.Client{
		Transport: &oauth2.Transport{
			Base:   c.defaultHTTPClient.Transport,
			Source: oauth2.ReuseTokenSource(nil, ts),
		},
	}
}
