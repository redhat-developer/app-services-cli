package kas

import (
	"net/http"
	"net/url"

	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/api/kas/client"
)

// New creates a new Kafka API client
func New(apiURL *url.URL, httpClient *http.Client) *kasclient.APIClient {
	cfg := kasclient.NewConfiguration()

	cfg.Scheme = apiURL.Scheme
	cfg.Host = apiURL.Host

	cfg.HTTPClient = httpClient

	apiClient := kasclient.NewAPIClient(cfg)

	return apiClient
}
