package setting

import (
	"github.com/redhat-developer/app-services-cli/pkg/cmd/registry/setting/get"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/registry/setting/list"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/registry/setting/set"
	"github.com/redhat-developer/app-services-cli/pkg/shared/factory"
	"github.com/spf13/cobra"
)

func NewSettingCommand(f *factory.Factory) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "setting",
		Short:   f.Localizer.MustLocalize("setting.cmd.description.short"),
		Long:    f.Localizer.MustLocalize("setting.cmd.description.long"),
		Example: f.Localizer.MustLocalize("setting.cmd.example"),
		Args:    cobra.MinimumNArgs(1),
		Hidden:  true,
	}

	// add sub-commands
	cmd.AddCommand(
		list.NewListCommand(f),
		get.NewGetCommand(f),
		set.NewSetCommand(f),
	)

	return cmd
}
