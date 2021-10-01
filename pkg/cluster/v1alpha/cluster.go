package v1alpha

import (
	"context"
)

// TODO rename options
type ConnectArguments struct {
	OfflineAccessToken      string
	ForceCreationWithoutAsk bool
	IgnoreContext           bool
	Namespace               string
	SelectedServiceType     string
	SelectedServiceID       string
}

type ServiceBindingOptions struct {
	ServiceName             string
	Namespace               string
	AppName                 string
	ForceCreationWithoutAsk bool
	BindingName             string
	BindAsFiles             bool
	DeploymentConfigEnabled bool
}

// Cluster defines methods used to interact with a cluster
// TODO rename to Cluster API
type Cluster interface {
	Connect(ctx context.Context, connectOpts *ConnectArguments, opts InputOptions) error
	ExecuteServiceBinding(ctx context.Context, service CustomConnection, opts InputOptions, options *ServiceBindingOptions) error
	IsRhoasOperatorAvailableOnCluster(ctx context.Context) (bool, error)
	CurrentNamespace() (string, error)
}
