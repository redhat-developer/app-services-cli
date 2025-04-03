package group

import (
	"github.com/redhat-developer/app-services-cli/pkg/cmd/registry/group/create"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/registry/group/get"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/registry/group/list"
	"github.com/redhat-developer/app-services-cli/pkg/shared/factory"
	"github.com/spf13/cobra"
)

func NewGroupCommand(f *factory.Factory) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "group",
		Short:   f.Localizer.MustLocalize("group.cmd.description.short"),
		Long:    f.Localizer.MustLocalize("group.cmd.description.long"),
		Example: f.Localizer.MustLocalize("group.cmd.example"),
		Args:    cobra.MinimumNArgs(1),
		Hidden:  true,
	}

	// add sub-commands
	cmd.AddCommand(
		list.NewListCommand(f),
		create.NewCreateCommand(f),
		get.NewGetCommand(f),
	)

	return cmd
}
