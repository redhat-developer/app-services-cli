package cluster

import (
	"context"
)

// CustomConnection interface defines type for custom resource structs
type CustomConnection interface {
	CreateCustomResource(ctx context.Context, c *KubernetesCluster, serviceID string) error
	CustomResourceExists(ctx context.Context, c *KubernetesCluster, serviceID string) (int, error)
	BindCustomConnection(ctx context.Context, serviceName string, options ServiceBindingOptions, clients *KubernetesClients) error
}
