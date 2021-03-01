package cluster

import (
	"context"
)

type ConnectArguments struct {
	OfflineAccessToken      string
	ForceCreationWithoutAsk bool
	IgnoreContext           bool
	SelectedKafka           string
	Namespace               string
}

// Cluster defines methods used to interact with a cluster
type Cluster interface {
	Connect(ctx context.Context, opts *ConnectArguments) error
	IsRhoasOperatorAvailableOnCluster(ctx context.Context) (bool, error)
	CurrentNamespace() (string, error)
}
