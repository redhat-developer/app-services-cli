package cluster

import (
	"context"
)

// Cluster defines methods used to interact with a cluster
type Cluster interface {
	Connect(ctx context.Context, secretName string, forceSelect bool) error
	IsKafkaConnectionCRDInstalled(ctx context.Context) (bool, error)
	CurrentNamespace() (string, error)
}
