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

	// The implementation of `rhoas kafka use` command has been aliased here as `rhoas context kafka-use`
	kafkaUseCmd := kafkaUse.NewUseCommand(f)
	kafkaUseCmd.Use = "kafka-use"

	// The implementation of `rhoas service-registry use` command has been aliased here as `rhoas context service-registry-use`
	registryUseCmd := registryUse.NewUseCommand(f)
	registryUseCmd.Use = "service-registry-use"

	cmd.AddCommand(
		use.NewUseCommand(f),
		status.NewStatusCommand(f),
		list.NewListCommand(f),
		create.NewCreateCommand(f),
		delete.NewDeleteCommand(f),
		kafkaUseCmd,
		registryUseCmd,

		// `rhoas status` cmd has been re-used as `rhoas context status`
		status.NewStatusCommand(f),
	)
	return cmd
}
