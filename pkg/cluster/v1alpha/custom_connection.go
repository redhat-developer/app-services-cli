package v1alpha

// RHOASKubernetesService interface defines type for custom resource structs
type RHOASKubernetesService interface {
	CreateCustomResource(serviceID string) error
	BindCustomConnection(serviceName string, options BindOperationOptions) error
	CustomResourceExists(serviceID string) (int, error)
}
