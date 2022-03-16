package context

import (
	"path/filepath"

	"github.com/redhat-developer/app-services-cli/pkg/cmd/context/create"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/context/delete"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/context/list"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/context/use"
	kafkaUse "github.com/redhat-developer/app-services-cli/pkg/cmd/kafka/use"
	registryUse "github.com/redhat-developer/app-services-cli/pkg/cmd/registry/use"
	"github.com/redhat-developer/app-services-cli/pkg/core/localize"
	"github.com/redhat-developer/app-services-cli/pkg/core/servicecontext"

	"github.com/redhat-developer/app-services-cli/pkg/cmd/status"
	"github.com/redhat-developer/app-services-cli/pkg/shared/factory"
	"github.com/spf13/cobra"
)

// NewContextCmd creates a new command to manage service contexts
func NewContextCmd(f *factory.Factory) *cobra.Command {

	ctxDirLocation, _ := servicecontext.DefaultDir()

	ctxLocation := filepath.Join(ctxDirLocation, "contexts.json")

	cmd := &cobra.Command{
		Use:     "context",
		Short:   f.Localizer.MustLocalize("context.cmd.shortDescription"),
		Long:    f.Localizer.MustLocalize("context.cmd.longDescription", localize.NewEntry("ContextPath", ctxLocation)),
		Example: f.Localizer.MustLocalize("context.cmd.example"),
		Args:    cobra.NoArgs,
	}

	// The implementation of `rhoas kafka use` command has been aliased here as `rhoas context kafka-use`
	kafkaUseCmd := kafkaUse.NewUseCommand(f)
	kafkaUseCmd.Use = "kafka-use"
	kafkaUseCmd.Example = f.Localizer.MustLocalize("context.kafkaUse.cmd.example")

	// `rhoas status` cmd has been re-used as `rhoas context status`
	statusCmd := status.NewStatusCommand(f)
	statusCmd.Example = f.Localizer.MustLocalize("context.status.cmd.example")

	// The implementation of `rhoas service-registry use` command has been aliased here as `rhoas context service-registry-use`
	registryUseCmd := registryUse.NewUseCommand(f)
	registryUseCmd.Use = "service-registry-use"
	registryUseCmd.Example = f.Localizer.MustLocalize("context.registryUse.cmd.example")

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
