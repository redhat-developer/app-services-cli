package defaultapi

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"net/url"

	connectormgmt "github.com/redhat-developer/app-services-sdk-go/connectormgmt/apiv1"
	connectormgmtclient "github.com/redhat-developer/app-services-sdk-go/connectormgmt/apiv1/client"
	kafkamgmt "github.com/redhat-developer/app-services-sdk-go/kafkamgmt/apiv1"

	"github.com/redhat-developer/app-services-cli/pkg/shared/kafkautil"

	"github.com/redhat-developer/app-services-cli/internal/build"
	"github.com/redhat-developer/app-services-cli/pkg/api/generic"
	"github.com/redhat-developer/app-services-cli/pkg/api/rbac"
	"github.com/redhat-developer/app-services-cli/pkg/shared/connection/api"
	"github.com/redhat-developer/app-services-cli/pkg/shared/svcstatus"
	amsclient "github.com/redhat-developer/app-services-sdk-go/accountmgmt/apiv1/client"
	kafkainstance "github.com/redhat-developer/app-services-sdk-go/kafkainstance/apiv1"
	kafkainstanceclient "github.com/redhat-developer/app-services-sdk-go/kafkainstance/apiv1/client"
	kafkamgmtclient "github.com/redhat-developer/app-services-sdk-go/kafkamgmt/apiv1/client"
	kafkamgmtv1errors "github.com/redhat-developer/app-services-sdk-go/kafkamgmt/apiv1/error"
	registryinstance "github.com/redhat-developer/app-services-sdk-go/registryinstance/apiv1internal"
	registryinstanceclient "github.com/redhat-developer/app-services-sdk-go/registryinstance/apiv1internal/client"
	registrymgmt "github.com/redhat-developer/app-services-sdk-go/registrymgmt/apiv1"
	registrymgmtclient "github.com/redhat-developer/app-services-sdk-go/registrymgmt/apiv1/client"
	"golang.org/x/oauth2"

	svcacctmgmt "github.com/redhat-developer/app-services-sdk-go/serviceaccountmgmt/apiv1"
	svcacctmgmtclient "github.com/redhat-developer/app-services-sdk-go/serviceaccountmgmt/apiv1/client"
)

// defaultAPI is a type which defines a number of API creator functions
type defaultAPI struct {
	api.Config
}

// New creates a new default API client wrapper
func New(cfg *api.Config) api.API {
	return &defaultAPI{
		Config: *cfg,
	}
}

func (a *defaultAPI) GetConfig() api.Config {
	return a.Config
}

// KafkaMgmt returns a new Kafka Management API client instance
func (a *defaultAPI) KafkaMgmt() kafkamgmtclient.DefaultApi {
	tc := a.CreateOAuthTransport(a.AccessToken)
	client := kafkamgmt.NewAPIClient(&kafkamgmt.Config{
		BaseURL:    a.ApiURL.String(),
		Debug:      a.Logger.DebugEnabled(),
		HTTPClient: tc,
		UserAgent:  a.UserAgent,
	})

	return client.DefaultApi
}

// ServiceRegistryMgmt return a new Service Registry Management API client instance
func (a *defaultAPI) ServiceRegistryMgmt() registrymgmtclient.RegistriesApi {
	tc := a.CreateOAuthTransport(a.AccessToken)
	client := registrymgmt.NewAPIClient(&registrymgmt.Config{
		BaseURL:    a.ApiURL.String(),
		Debug:      a.Logger.DebugEnabled(),
		HTTPClient: tc,
		UserAgent:  build.DefaultUserAgentPrefix + build.Version,
	})

	return client.RegistriesApi
}

// ServiceAccountMgmt return a new Service Account Management API client instance
func (a *defaultAPI) ServiceAccountMgmt() svcacctmgmtclient.ServiceAccountsApi {
	tc := a.CreateOAuthTransport(a.AccessToken)
	client := svcacctmgmt.NewAPIClient(&svcacctmgmt.Config{
		BaseURL:    a.AuthURL.String(),
		Debug:      a.Logger.DebugEnabled(),
		HTTPClient: tc,
		UserAgent:  a.UserAgent,
	})

	return client.ServiceAccountsApi
}

// KafkaAdmin returns a new Kafka Admin API client instance, with the Kafka configuration object
func (a *defaultAPI) KafkaAdmin(instanceID string) (*kafkainstanceclient.APIClient, *kafkamgmtclient.KafkaRequest, error) {
	kafkaAPI := a.KafkaMgmt()

	kafkaInstance, resp, err := kafkaAPI.GetKafkaById(context.Background(), instanceID).Execute()
	if resp != nil {
		defer resp.Body.Close()
	}
	if kafkamgmtv1errors.IsAPIError(err, kafkamgmtv1errors.ERROR_7) {
		return nil, nil, kafkautil.NotFoundByIDError(instanceID)
	} else if err != nil {
		return nil, nil, fmt.Errorf("%w", err)
	}

	kafkaStatus := kafkaInstance.GetStatus()

	switch kafkaStatus {
	case svcstatus.StatusProvisioning, svcstatus.StatusAccepted:
		err = fmt.Errorf(`kafka instance "%v" is not ready yet`, kafkaInstance.GetName())
		return nil, nil, err
	case svcstatus.StatusFailed:
		err = fmt.Errorf(`kafka instance "%v" has failed`, kafkaInstance.GetName())
		return nil, nil, err
	case svcstatus.StatusDeprovision:
		err = fmt.Errorf(`kafka instance "%v" is being deprovisioned`, kafkaInstance.GetName())
		return nil, nil, err
	case svcstatus.StatusDeleting:
		err = fmt.Errorf(`kafka instance "%v" is being deleted`, kafkaInstance.GetName())
		return nil, nil, err
	}

	bootstrapURL := kafkaInstance.GetBootstrapServerHost()
	if bootstrapURL == "" {
		err = fmt.Errorf(`bootstrap URL is missing for Kafka instance "%v"`, kafkaInstance.GetName())

		return nil, nil, err
	}

	apiURL := kafkaInstance.GetAdminApiServerUrl()
	a.Logger.Debugf("Making request to %v", apiURL)

	token := a.AccessToken

	client := kafkainstance.NewAPIClient(&kafkainstance.Config{
		BaseURL:    apiURL,
		Debug:      a.Logger.DebugEnabled(),
		HTTPClient: a.CreateOAuthTransport(token),
		UserAgent:  a.UserAgent,
	})

	return client, &kafkaInstance, nil
}

// ServiceRegistryInstance returns a new Service Registry API client instance, with the Registry configuration object
func (a *defaultAPI) ServiceRegistryInstance(instanceID string) (*registryinstanceclient.APIClient, *registrymgmtclient.Registry, error) {
	registryAPI := a.ServiceRegistryMgmt()

	instance, resp, err := registryAPI.GetRegistry(context.Background(), instanceID).Execute()
	defer resp.Body.Close()
	if err != nil {
		return nil, nil, fmt.Errorf("%w", err)
	}

	status := svcstatus.ServiceStatus(instance.GetStatus())
	// nolint
	switch status {
	case svcstatus.StatusProvisioning, svcstatus.StatusAccepted:
		err = fmt.Errorf(`service registry instance "%v" is not ready yet`, instance.GetName())
		return nil, nil, err
	case svcstatus.StatusFailed:
		err = fmt.Errorf(`service registry instance "%v" has failed`, instance.GetName())
		return nil, nil, err
	case svcstatus.StatusDeprovision:
		err = fmt.Errorf(`service registry instance "%v" is being deprovisioned`, instance.GetName())
		return nil, nil, err
	case svcstatus.StatusDeleting:
		err = fmt.Errorf(`service registry instance "%v" is being deleted`, instance.GetName())
		return nil, nil, err
	}

	registryUrl := instance.GetRegistryUrl()
	if registryUrl == "" {
		err = fmt.Errorf(`URL is missing for Service Registry instance "%v"`, instance.GetName())

		return nil, nil, err
	}

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

	a.Logger.Debugf("Making request to %v", baseURL)

	token := a.AccessToken

	client := registryinstance.NewAPIClient(&registryinstance.Config{
		BaseURL:    baseURL,
		Debug:      a.Logger.DebugEnabled(),
		HTTPClient: a.CreateOAuthTransport(token),
		UserAgent:  build.DefaultUserAgentPrefix + build.Version,
	})

	return client, &instance, nil
}

func (a *defaultAPI) GenericAPI() generic.GenericAPI {
	tc := a.CreateOAuthTransport(a.AccessToken)
	client := generic.NewGenericAPIClient(&generic.Config{
		BaseURL:    a.ApiURL.String(),
		Debug:      a.Logger.DebugEnabled(),
		HTTPClient: tc,
	})

	return client
}

func (a *defaultAPI) ConnectorsMgmt() connectormgmtclient.APIClient {
	tc := a.CreateOAuthTransport(a.AccessToken)
	client := connectormgmt.NewAPIClient(&connectormgmt.Config{
		BaseURL:    a.ApiURL.String(),
		Debug:      a.Logger.DebugEnabled(),
		HTTPClient: tc,
	})

	return *client
}

// AccountMgmt returns a new Account Management API client instance
func (a *defaultAPI) AccountMgmt() amsclient.AppServicesApi {
	cfg := amsclient.NewConfiguration()

	cfg.Scheme = a.ApiURL.Scheme
	cfg.Host = a.ApiURL.Host
	cfg.UserAgent = a.UserAgent

	cfg.HTTPClient = a.CreateOAuthTransport(a.AccessToken)

	apiClient := amsclient.NewAPIClient(cfg)

	return apiClient.AppServicesApi
}

// RBAC returns a new RBAC API client instance
func (a *defaultAPI) RBAC() rbac.RbacAPI {
	rbacAPI := rbac.RbacAPI{
		PrincipalAPI: func() rbac.PrincipalAPI {
			cl := a.CreateOAuthTransport(a.AccessToken)
			cfg := rbac.Config{
				HTTPClient: cl,
				Debug:      a.Logger.DebugEnabled(),
				BaseURL:    a.ConsoleURL,
			}
			return rbac.NewPrincipalAPIClient(&cfg)
		},
	}
	return rbacAPI
}

// wraps the HTTP client with an OAuth2 Transport layer to provide automatic token refreshing
func (a *defaultAPI) CreateOAuthTransport(accessToken string) *http.Client {
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{
			AccessToken: accessToken,
		},
	)

	return &http.Client{
		Transport: &oauth2.Transport{
			Base:   a.HTTPClient.Transport,
			Source: oauth2.ReuseTokenSource(nil, ts),
		},
	}
}
