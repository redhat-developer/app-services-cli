package context

import (
	"github.com/redhat-developer/app-services-cli/pkg/cmd/context/create"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/context/delete"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/context/list"
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

	// The implementation of `rhoas kafka use` command has been aliased here as `rhoas context use-kafka`
	kafkaUseCmd := kafkaUse.NewUseCommand(f)
	kafkaUseCmd.Use = "use-kafka"
	kafkaUseCmd.Example = f.Localizer.MustLocalize("context.useKafka.cmd.example")

	// `rhoas status` cmd has been re-used as `rhoas context status`
	statusCmd := status.NewStatusCommand(f)
	statusCmd.Example = f.Localizer.MustLocalize("context.status.cmd.example")

	// The implementation of `rhoas service-registry use` command has been aliased here as `rhoas context use-service-registry`
	registryUseCmd := registryUse.NewUseCommand(f)
	registryUseCmd.Use = "use-service-registry"
	registryUseCmd.Example = f.Localizer.MustLocalize("context.useRegistry.cmd.example")

	cmd.AddCommand(
		use.NewUseCommand(f),
		list.NewListCommand(f),
		create.NewCreateCommand(f),
		delete.NewDeleteCommand(f),

		// reused sub-commands
		kafkaUseCmd,
		registryUseCmd,
		statusCmd,
	)
	return cmd
}
