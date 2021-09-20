package cluster

import (
	"context"
)

type ConnectArguments struct {
	OfflineAccessToken      string
	ForceCreationWithoutAsk bool
	IgnoreContext           bool
	SelectedKafka           string
	SelectedRegistry        string
	Namespace               string
	SelectedService         string
	SelectedServiceID       string
}

// Cluster defines methods used to interact with a cluster
type Cluster interface {
	Connect(ctx context.Context, connectOpts *ConnectArguments, opts Options) error
	IsRhoasOperatorAvailableOnCluster(ctx context.Context) (bool, error)
	CurrentNamespace() (string, error)
}
