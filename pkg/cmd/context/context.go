package context

import (
	namespaceUse "github.com/redhat-developer/app-services-cli/pkg/cmd/connector/namespace/use"
	connectorUse "github.com/redhat-developer/app-services-cli/pkg/cmd/connector/use"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/context/create"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/context/delete"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/context/list"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/context/unset"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/context/use"
	kafkaUse "github.com/redhat-developer/app-services-cli/pkg/cmd/kafka/use"
	registryUse "github.com/redhat-developer/app-services-cli/pkg/cmd/registry/use"

	"github.com/redhat-developer/app-services-cli/pkg/cmd/status"
	"github.com/redhat-developer/app-services-cli/pkg/shared/factory"
	"github.com/spf13/cobra"
)

// NewContextCmd creates a new command to manage service contexts
func NewContextCmd(f *factory.Factory) *cobra.Command {

	cmd := &cobra.Command{
		Use:     "context",
		Short:   f.Localizer.MustLocalize("context.cmd.shortDescription"),
		Long:    f.Localizer.MustLocalize("context.cmd.longDescription"),
		Example: f.Localizer.MustLocalize("context.cmd.example"),
		Args:    cobra.NoArgs,
	}

	// The implementation of `rhoas kafka use` command has been aliased here as `rhoas context set-kafka`
	kafkaUseCmd := kafkaUse.NewUseCommand(f)
	kafkaUseCmd.Use = "set-kafka"
	kafkaUseCmd.Example = f.Localizer.MustLocalize("context.setKafka.cmd.example")

	// The implementation of `rhoas connector namespace use` command has been aliased here as `rhoas context set-namespace`
	connectorUseCmd := connectorUse.NewUseCommand(f)
	connectorUseCmd.Use = "set-connector"
	connectorUseCmd.Example = f.Localizer.MustLocalize("context.setConnector.cmd.example")

	// The implementation of `rhoas connector use` command has been aliased here as `rhoas context set-connector`
	namespaceUseCmd := namespaceUse.NewUseCommand(f)
	namespaceUseCmd.Use = "set-namespace"
	namespaceUseCmd.Example = f.Localizer.MustLocalize("context.setNamespace.cmd.example")

	// `rhoas status` cmd has been re-used as `rhoas context status`
	statusCmd := status.NewStatusCommand(f)
	statusCmd.Long = f.Localizer.MustLocalize("context.status.cmd.longDescription")
	statusCmd.Example = f.Localizer.MustLocalize("context.status.cmd.example")

	// The implementation of `rhoas service-registry use` command has been aliased here as `rhoas context set-service-registry`
	registryUseCmd := registryUse.NewUseCommand(f)
	registryUseCmd.Use = "set-service-registry"
	registryUseCmd.Example = f.Localizer.MustLocalize("context.setRegistry.cmd.example")

	cmd.AddCommand(
		use.NewUseCommand(f),
		list.NewListCommand(f),
		create.NewCreateCommand(f),
		delete.NewDeleteCommand(f),
		unset.NewUnsetCommand(f),

		// reused sub-commands
		kafkaUseCmd,
		connectorUseCmd,
		namespaceUseCmd,
		registryUseCmd,
		statusCmd,
	)
	return cmd
}
