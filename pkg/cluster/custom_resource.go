package cluster

import (
	"context"

	"k8s.io/client-go/dynamic"
)

type CustomConnection interface {
	CreateCustomResource(ctx context.Context, c *KubernetesCluster, serviceID string) error
	CustomResourceExists(ctx context.Context, c *KubernetesCluster, serviceID string) error
	BindCustomConnection(ctx context.Context, serviceName string, options ServiceBindingOptions, clients *KubernetesClients) error
	CustomConnectionExists(ctx context.Context, dynamicClient dynamic.Interface, serviceName string, ns string) error
}
