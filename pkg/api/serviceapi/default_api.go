package serviceapi

import (
	"net/http"
	"net/url"

	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/api/serviceapi/client"
)

// New creates a new Kafka API client
func New(apiURL *url.URL, httpClient *http.Client, accessToken string) *client.APIClient {
	masCfg := client.NewConfiguration()

	masCfg.Scheme = apiURL.Scheme
	masCfg.Host = apiURL.Host

	masCfg.HTTPClient = httpClient

	apiClient := client.NewAPIClient(masCfg)

	return apiClient
}
