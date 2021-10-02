package serviceregistry

import (
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

const (
	SRCGroup   = "rhoas.redhat.com"
	SRCVersion = "v1alpha1"
)

var RegistryResourceMeta = metav1.TypeMeta{
	Kind:       "ServiceRegistryConnection",
	APIVersion: SRCGroup + "/" + SRCVersion,
}

var SRCResource = schema.GroupVersionResource{
	Group:    SRCGroup,
	Version:  SRCVersion,
	Resource: "serviceregistryconnections",
}

func GetServiceRegistryAPIURL(namespace string) string {
	return fmt.Sprintf("/apis/rhoas.redhat.com/v1alpha1/namespaces/%v/serviceregistryconnections", namespace)
}
