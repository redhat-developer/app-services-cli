package kafkainstance

import (
	"github.com/jackdelahunt/app-services-sdk-core/app-services-sdk-go/core"

	apiv1 "github.com/jackdelahunt/app-services-sdk-core/app-services-sdk-go/kafkainstance/apiv1/client"
)

// APIConfig defines the available configuration options
// to customise the API client settings
type Config = core.APIConfig

// NewAPIClient returns a new KafkaManagement v1 API client
// using a custom config
func NewAPIClient(cfg *Config) *apiv1.APIClient {
	apiCfg := apiv1.NewConfiguration()
	if cfg == nil {
		return apiv1.NewAPIClient(apiCfg)
	}

	if cfg.HTTPClient != nil {
		apiCfg.HTTPClient = cfg.HTTPClient
	}
	if cfg.BaseURL != "" {
		apiCfg.Servers = []apiv1.ServerConfiguration{
			{
				URL: cfg.BaseURL,
			},
		}
	}

	apiCfg.Debug = cfg.Debug
	apiCfg.UserAgent = cfg.UserAgent

	client := apiv1.NewAPIClient(apiCfg)

	return client
}
