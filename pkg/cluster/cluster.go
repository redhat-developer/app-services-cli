package cluster

import (
	"context"
)

// Cluster defines methods used to interact with a cluster
type Cluster interface {
	Connect(ctx context.Context, interactiveSelect bool, token string) error
	IsKafkaConnectionCRDInstalled(ctx context.Context) (bool, error)
	CurrentNamespace() (string, error)
}
