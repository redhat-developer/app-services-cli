package api

import (
	"github.com/redhat-developer/app-services-cli/pkg/api/generic"
	"github.com/redhat-developer/app-services-cli/pkg/api/rbac"
	connectormgmtclient "github.com/redhat-developer/app-services-sdk-go/connectormgmt/apiv1/client"

	amsclient "github.com/redhat-developer/app-services-sdk-go/accountmgmt/apiv1/client"
	kafkainstanceclient "github.com/redhat-developer/app-services-sdk-go/kafkainstance/apiv1internal/client"
	kafkamgmtclient "github.com/redhat-developer/app-services-sdk-go/kafkamgmt/apiv1/client"
	registryinstanceclient "github.com/redhat-developer/app-services-sdk-go/registryinstance/apiv1internal/client"
	registrymgmtclient "github.com/redhat-developer/app-services-sdk-go/registrymgmt/apiv1/client"
)

type API interface {
	KafkaMgmt() kafkamgmtclient.DefaultApi
	ServiceRegistryMgmt() registrymgmtclient.RegistriesApi
	ConnectorsMgmt() connectormgmtclient.APIClient
	ServiceAccountMgmt() kafkamgmtclient.SecurityApi
	KafkaAdmin(instanceID string) (*kafkainstanceclient.APIClient, *kafkamgmtclient.KafkaRequest, error)
	ServiceRegistryInstance(instanceID string) (*registryinstanceclient.APIClient, *registrymgmtclient.Registry, error)
	AccountMgmt() amsclient.AppServicesApi
	RBAC() rbac.RbacAPI
	GenericAPI() generic.GenericAPI
}
