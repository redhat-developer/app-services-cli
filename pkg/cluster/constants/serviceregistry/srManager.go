package serviceregistry

import (
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

var (
	SRCGroup   = "rhoas.redhat.com"
	SRCVersion = "v1alpha1"
)

var RegistryResourceMeta = metav1.TypeMeta{
	Kind:       "ServiceRegistryConnection",
	APIVersion: SRCGroup + "/" + SRCVersion,
}

var tokenSecretName = "rh-cloud-services-accesstoken-cli"

/*  #nosec */
var serviceAccountSecretName = "rh-cloud-services-service-account"

var SRCResource = schema.GroupVersionResource{
	Group:    SRCGroup,
	Version:  SRCVersion,
	Resource: "serviceregistryconnections",
}

func GetServiceRegistryAPIURL(namespace string) string {
	return fmt.Sprintf("/apis/rhoas.redhat.com/v1alpha1/namespaces/%v/serviceregistryconnections", namespace)
}

func CreateSRObject(crName string, namespace string, registryID string) *ServiceRegsitryConnection {
	serviceRegistryCR := &ServiceRegsitryConnection{
		ObjectMeta: metav1.ObjectMeta{
			Name:      crName,
			Namespace: namespace,
		},
		TypeMeta: RegistryResourceMeta,
		Spec: ServiceRegsitryConnectionSpec{
			ServiceRegistryId:     registryID,
			AccessTokenSecretName: tokenSecretName,
			Credentials: CredentialsSpec{
				SecretName: serviceAccountSecretName,
			},
		},
	}

	return serviceRegistryCR
}
