package services

import "k8s.io/apimachinery/pkg/runtime/schema"

// ServiceDetails contains metadata for service including structure that should be used
// to create service CR
type ServiceDetails struct {
	Type               string
	ID                 string
	Name               string
	KubernetesResource interface{}
	GroupMetadata      schema.GroupVersionResource
}

// RHOASKubernetesService interface defines type for custom resource structs
type RHOASKubernetesService interface {

	// Build Custom Resource representing desired service that should be created
	BuildServiceDetails(serviceName string, namespace string, ignoreConfigContext bool) (*ServiceDetails, error)
}
