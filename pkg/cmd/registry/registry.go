// Package kafka instance contains commands for interacting with cluster logic of the service directly instead of through the
// REST API exposed via the serve command.
package registry

import (
	"github.com/spf13/cobra"

	"github.com/redhat-developer/app-services-cli/pkg/cmd/factory"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/kafka/create"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/kafka/delete"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/kafka/describe"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/kafka/list"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/kafka/use"
)

func NewServiceRegistryCommand(f *factory.Factory) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "registry",
		Short: "Service Registry commands",
		Args:  cobra.MinimumNArgs(1),
	}

	// add sub-commands
	cmd.AddCommand(
		create.NewCreateCommand(f),
		describe.NewDescribeCommand(f),
		delete.NewDeleteCommand(f),
		list.NewListCommand(f),
		use.NewUseCommand(f),
	)

	return cmd
}
