package cluster

import "context"

type CustomConnection interface {
	CreateCustomResource(ctx context.Context, c *KubernetesCluster, serviceID string, opts Options) error
	CustomResourceExists(ctx context.Context, c *KubernetesCluster, serviceID string, opts Options) error
}
