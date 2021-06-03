package api

import (
	"github.com/redhat-developer/app-services-cli/pkg/api/ams/amsclient"
	kafkainstanceclient "github.com/redhat-developer/app-services-sdk-go/kafkainstance/apiv1internal/client"
	kafkamgmtclient "github.com/redhat-developer/app-services-sdk-go/kafkamgmt/apiv1/client"
)

// API is a type which defines a number of API creator functions
type API struct {
	Kafka          func() kafkamgmtclient.DefaultApi
	ServiceAccount func() kafkamgmtclient.SecurityApi
	KafkaAdmin     func(kafkaID string) (kafkainstanceclient.DefaultApi, *kafkamgmtclient.KafkaRequest, error)
	AccountMgmt    func() amsclient.DefaultApi
}
