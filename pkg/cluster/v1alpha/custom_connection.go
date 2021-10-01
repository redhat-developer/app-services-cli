package v1alpha

import (
	"context"
)

// CustomConnection interface defines type for custom resource structs
type CustomConnection interface {
	CreateCustomResource(ctx context.Context, serviceID string) error
	BindCustomConnection(ctx context.Context, serviceName string, options BindOperationOptions, clients *KubernetesClients) error
	CustomResourceExists(ctx context.Context, serviceID string) (int, error)
}
