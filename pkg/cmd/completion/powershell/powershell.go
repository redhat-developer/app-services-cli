package powershell

import (
	"github.com/spf13/cobra"

	"github.com/redhat-developer/app-services-cli/pkg/shared/factory"
)

func NewCommand(f *factory.Factory) *cobra.Command {
	cmd := &cobra.Command{
		Use:                   "powershell",
		Short:                 f.Localizer.MustLocalize("completion.powershell.cmd.shortDescription"),
		Long:                  f.Localizer.MustLocalize("completion.powershell.cmd.longDescription"),
		Example:               f.Localizer.MustLocalize("completion.powershell.cmd.example"),
		DisableFlagsInUseLine: true,
		Args:                  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmd.Parent().Parent().GenPowerShellCompletion(f.IOStreams.Out)
		},
	}

	return cmd
}
