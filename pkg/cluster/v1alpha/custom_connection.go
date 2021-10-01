package v1alpha

import (
	"context"
)

// RHOASKubernetesService interface defines type for custom resource structs
type RHOASKubernetesService interface {
	CreateCustomResource(ctx context.Context, serviceID string) error
	BindCustomConnection(ctx context.Context, serviceName string, options BindOperationOptions) error
	CustomResourceExists(ctx context.Context, serviceID string) (int, error)
}
