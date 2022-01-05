package role

import (
	"github.com/redhat-developer/app-services-cli/pkg/cmd/registry/artifact/role/add"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/registry/artifact/role/list"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/registry/artifact/role/revoke"
	"github.com/redhat-developer/app-services-cli/pkg/core/cmdutil/factory"
	"github.com/spf13/cobra"
)

func NewRoleCommand(f *factory.Factory) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "role",
		Short:   f.Localizer.MustLocalize("registry.role.cmd.description.short"),
		Long:    f.Localizer.MustLocalize("registry.role.cmd.description.long"),
		Example: f.Localizer.MustLocalize("registry.role.cmd.example"),
		Args:    cobra.MinimumNArgs(1),
	}

	cmd.AddCommand(
		add.NewAddCommand(f),
		revoke.NewRevokeCommand(f),
		list.NewListCommand(f),
	)

	return cmd
}
