package cluster

import (
	"context"
)

type ConnectArguments struct {
	OfflineAccessToken      string
	ForceCreationWithoutAsk bool
	IgnoreContext           bool
	Namespace               string
	SelectedService         string
	SelectedServiceID       string
}

// Cluster defines methods used to interact with a cluster
type Cluster interface {
	Connect(ctx context.Context, connectOpts *ConnectArguments, c CustomConnection, opts Options) error
	IsRhoasOperatorAvailableOnCluster(ctx context.Context) (bool, error)
	CurrentNamespace() (string, error)
	ExecuteServiceBinding(ctx context.Context, service CustomConnection, opts Options, options *ServiceBindingOptions) error
}
