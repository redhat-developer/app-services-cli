package v1alpha

import (
	"context"

	"github.com/redhat-developer/app-services-cli/internal/config"
	"github.com/redhat-developer/app-services-cli/pkg/connection"
	"github.com/redhat-developer/app-services-cli/pkg/iostreams"
	"github.com/redhat-developer/app-services-cli/pkg/localize"
	"github.com/redhat-developer/app-services-cli/pkg/logging"
)

// CommandEnvironment provides number of abstractions provided by CLI
type CommandEnvironment struct {
	Connection connection.Connection
	Config     config.IConfig
	Logger     logging.Logger
	IO         *iostreams.IOStreams
	Localizer  localize.Localizer
	Context    context.Context
}

// ConnectOperationOptions contains input flags for connect method
type ConnectOperationOptions struct {
	OfflineAccessToken      string
	ForceCreationWithoutAsk bool
	IgnoreContext           bool
	Namespace               string
	SelectedServiceType     string
	SelectedServiceID       string
}

// BindOperationOptions contains input flags for bind method
type BindOperationOptions struct {
	ServiceName             string
	Namespace               string
	AppName                 string
	ForceCreationWithoutAsk bool
	BindingName             string
	BindAsFiles             bool
	DeploymentConfigEnabled bool
}

// ClusterUserAPI -  interact with kuberentes clusters in order to connect and bind resources
type ClusterUserAPI interface {
	ExecuteConnect(connectOpts *ConnectOperationOptions) error
	ExecuteServiceBinding(options *BindOperationOptions) error
	IsRhoasOperatorAvailableOnCluster() (bool, error)
	IsSBOOperatorAvailableOnCluster() (bool, error)
}
