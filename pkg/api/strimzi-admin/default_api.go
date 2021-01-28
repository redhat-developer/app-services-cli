package strimziadmin

import (
	"net/http"
	"net/url"

	strimziadminclient "github.com/bf2fc6cc711aee1a0c2a/cli/pkg/api/strimzi-admin/client"
)

// New creates a new Kafka API client
func New(apiURL *url.URL, httpClient *http.Client) *strimziadminclient.APIClient {
	cfg := strimziadminclient.NewConfiguration()

	cfg.Scheme = apiURL.Scheme
	cfg.Host = apiURL.Host

	cfg.HTTPClient = httpClient

	apiClient := strimziadminclient.NewAPIClient(cfg)

	return apiClient
}
