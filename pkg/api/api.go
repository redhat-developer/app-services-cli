package api

import (
	"github.com/redhat-developer/app-services-cli/pkg/api/ams/amsclient"
	kafkainstanceclient "github.com/redhat-developer/app-services-sdk-go/kafkainstance/apiv1internal/client"
	kafkamgmtclient "github.com/redhat-developer/app-services-sdk-go/kafkamgmt/apiv1/client"
	registryinstanceclient "github.com/redhat-developer/app-services-sdk-go/registryinstance/apiv1internal/client"
	registrymgmtclient "github.com/redhat-developer/app-services-sdk-go/registrymgmt/apiv1/client"
)

// API is a type which defines a number of API creator functions
type API struct {
	Kafka                   func() kafkamgmtclient.DefaultApi
	ServiceAccount          func() kafkamgmtclient.SecurityApi
	KafkaAdmin              func(kafkaID string) (*kafkainstanceclient.APIClient, *kafkamgmtclient.KafkaRequest, error)
	ServiceRegistryInstance func(registryID string) (*registryinstanceclient.APIClient, *registrymgmtclient.RegistryRest, error)
	AccountMgmt             func() amsclient.DefaultApi

	ServiceRegistryMgmt func() registrymgmtclient.RegistriesApi
}
