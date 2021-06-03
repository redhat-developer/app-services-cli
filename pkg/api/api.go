package api

import (
	"github.com/redhat-developer/app-services-cli/pkg/api/ams/amsclient"
	strimziadminclient "github.com/redhat-developer/app-services-cli/pkg/api/strimzi-admin/client"
	kafkamgmtclient "github.com/redhat-developer/app-services-sdk-go/kafkamgmt/apiv1/client"
)

// API is a type which defines a number of API creator functions
type API struct {
	Kafka          func() kafkamgmtclient.DefaultApi
	ServiceAccount func() kafkamgmtclient.SecurityApi
	TopicAdmin     func(kafkaID string) (strimziadminclient.DefaultApi, *kafkamgmtclient.KafkaRequest, error)
	AccountMgmt    func() amsclient.DefaultApi
}
