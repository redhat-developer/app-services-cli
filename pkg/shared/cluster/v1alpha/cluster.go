package v1alpha

import (
	"context"

	"github.com/redhat-developer/app-services-cli/pkg/core/config"
	"github.com/redhat-developer/app-services-cli/pkg/core/ioutil/iostreams"
	"github.com/redhat-developer/app-services-cli/pkg/core/localize"
	"github.com/redhat-developer/app-services-cli/pkg/core/logging"
	"github.com/redhat-developer/app-services-cli/pkg/shared/connection"
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
	Namespace               string
	ServiceType             string
	ServiceName             string
	IgnoreContext           bool
}

// BindOperationOptions contains input flags for bind method
type BindOperationOptions struct {
	Namespace               string
	ServiceName             string
	ServiceType             string
	AppName                 string
	ForceCreationWithoutAsk bool
	BindingName             string
	BindAsFiles             bool
	DeploymentConfigEnabled bool
	IgnoreContext           bool
}

type CleanOperationOptions struct {
	ForceDeleteWithoutAsk bool
	Namespace             string
}

// status of the Operator
type OperatorStatus struct {
	ServiceBindingOperatorAvailable bool
	RHOASOperatorAvailable          bool
	LatestRHOASVersionAvailable     bool
}

// ClusterUserAPI -  interact with kuberentes clusters in order to connect and bind resources
type ClusterUserAPI interface {
	ExecuteConnect(connectOpts *ConnectOperationOptions) error
	ExecuteServiceBinding(options *BindOperationOptions) error
	ExecuteStatus() (OperatorStatus, error)
	ExecuteClean(cleanOptions *CleanOperationOptions) error
}
