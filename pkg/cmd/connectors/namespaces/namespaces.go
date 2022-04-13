package namespaces

import (
	"github.com/redhat-developer/app-services-cli/pkg/cmd/connectors/namespaces/create"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/connectors/namespaces/list"
	"github.com/redhat-developer/app-services-cli/pkg/shared/factory"
	"github.com/spf13/cobra"
)

func NewNameSpaceCommand(f *factory.Factory) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "namespace",
		Short:   f.Localizer.MustLocalize("connector.namespace.cmd.shortDescription"),
		Long:    f.Localizer.MustLocalize("connector.namespace.cmd.longDescription"),
		Example: f.Localizer.MustLocalize("connector.namespace.cmd.example"),
		Args:    cobra.MinimumNArgs(1),
	}

	// add sub-commands
	cmd.AddCommand(
		list.NewListCommand(f),
		create.NewCreateCommand(f),
	)

	return cmd
}
