package api

import (
	"github.com/redhat-developer/app-services-cli/pkg/api/ams/amsclient"
	kasclient "github.com/redhat-developer/app-services-cli/pkg/api/kas/client"
	strimziadminclient "github.com/redhat-developer/app-services-cli/pkg/api/strimzi-admin/client"
)

// API is a type which defines a number of API creator functions
type API struct {
	Kafka       func() kasclient.DefaultApi
	TopicAdmin  func(kafkaID string) (strimziadminclient.DefaultApi, *kasclient.KafkaRequest, error)
	AccountMgmt func() amsclient.DefaultApi
}
