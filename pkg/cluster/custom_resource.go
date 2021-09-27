package cluster

import "context"

type CustomConnection interface {
	CreateCustomResource(ctx context.Context, c *KubernetesCluster, serviceID string, namespace string, opts Options) error
	CustomResourceExists(ctx context.Context, c *KubernetesCluster, namespace string, serviceName string, opts Options) error
}
