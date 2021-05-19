package api

import (
	"github.com/redhat-developer/app-services-cli/pkg/api/ams/amsclient"
	strimziadminclient "github.com/redhat-developer/app-services-cli/pkg/api/strimzi-admin/client"
	kafkamgmtv1 "github.com/redhat-developer/app-services-sdk-go/kafka/mgmt/apiv1"
)

// API is a type which defines a number of API creator functions
type API struct {
	Kafka       func() kafkamgmtv1.DefaultApi
	TopicAdmin  func(kafkaID string) (strimziadminclient.DefaultApi, *kafkamgmtv1.KafkaRequest, error)
	AccountMgmt func() amsclient.DefaultApi
}
