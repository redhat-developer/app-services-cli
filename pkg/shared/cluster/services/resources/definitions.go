package resources

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

const (
	AKCGroup   = "rhoas.redhat.com"
	AKCVersion = "v1alpha1"
)

var AKCRMeta = metav1.TypeMeta{
	Kind:       "KafkaConnection",
	APIVersion: AKCGroup + "/" + AKCVersion,
}

var AKCResource = schema.GroupVersionResource{
	Group:    AKCGroup,
	Version:  AKCVersion,
	Resource: "kafkaconnections",
}

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

// All services defined as resources
var AllResources = []schema.GroupVersionResource{
	AKCResource,
	SRCResource,
}
