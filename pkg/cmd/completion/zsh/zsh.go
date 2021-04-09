package zsh

import (
	"github.com/redhat-developer/app-services-cli/internal/localizer"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/factory"
	"github.com/spf13/cobra"
)

func NewCommand(f *factory.Factory) *cobra.Command {
	cmd := &cobra.Command{
		Use:                   localizer.MustLocalizeFromID("completion.zsh.cmd.use"),
		Short:                 localizer.MustLocalizeFromID("completion.zsh.cmd.shortDescription"),
		Long:                  localizer.MustLocalizeFromID("completion.zsh.cmd.longDescription"),
		DisableFlagsInUseLine: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmd.Parent().Parent().GenZshCompletion(f.IOStreams.Out)
		},
	}

	return cmd
}
