package zsh

import (
	"github.com/redhat-developer/app-services-cli/pkg/cmd/factory"
	"github.com/spf13/cobra"
)

func NewCommand(f *factory.Factory) *cobra.Command {
	cmd := &cobra.Command{
		Use:                   f.Localizer.LoadMessage("completion.zsh.cmd.use"),
		Short:                 f.Localizer.LoadMessage("completion.zsh.cmd.shortDescription"),
		Long:                  f.Localizer.LoadMessage("completion.zsh.cmd.longDescription"),
		DisableFlagsInUseLine: true,
		Args:                  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmd.Parent().Parent().GenZshCompletion(f.IOStreams.Out)
		},
	}

	return cmd
}
