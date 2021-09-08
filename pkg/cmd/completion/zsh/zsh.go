package zsh

import (
	"github.com/redhat-developer/app-services-cli/pkg/cmd/factory"
	"github.com/spf13/cobra"
)

func NewCommand(f *factory.Factory) *cobra.Command {
	cmd := &cobra.Command{
		Use:                   "zsh",
		Short:                 f.Localizer.MustLocalize("completion.zsh.cmd.shortDescription"),
		Long:                  f.Localizer.MustLocalize("completion.zsh.cmd.longDescription"),
		Example:               f.Localizer.MustLocalize("completion.zsh.cmd.example"),
		DisableFlagsInUseLine: true,
		Args:                  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmd.Parent().Parent().GenZshCompletion(f.IOStreams.Out)
		},
	}

	return cmd
}
