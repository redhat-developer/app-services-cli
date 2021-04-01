package api

import (
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/api/ams/amsclient"
	kasclient "github.com/bf2fc6cc711aee1a0c2a/cli/pkg/api/kas/client"
	strimziadminclient "github.com/bf2fc6cc711aee1a0c2a/cli/pkg/api/strimzi-admin/client"
)

// API is a type which defines a number of API creator functions
type API struct {
	Kafka       func() kasclient.DefaultApi
	TopicAdmin  func(kafkaID string) (strimziadminclient.DefaultApi, *kasclient.KafkaRequest, error)
	AccountMgmt func() amsclient.DefaultApi
}
