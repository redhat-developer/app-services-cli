package fish

import (
	"github.com/redhat-developer/app-services-cli/pkg/cmd/factory"
	"github.com/spf13/cobra"
)

func NewCommand(f *factory.Factory) *cobra.Command {
	cmd := &cobra.Command{
		Use:                   "fish",
		Short:                 f.Localizer.MustLocalize("completion.fish.cmd.shortDescription"),
		Long:                  f.Localizer.MustLocalize("completion.fish.cmd.longDescription"),
		Example:               f.Localizer.MustLocalize("completion.fish.cmd.example"),
		DisableFlagsInUseLine: true,
		Args:                  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmd.Parent().Parent().GenFishCompletion(f.IOStreams.Out, false)
		},
	}

	return cmd
}
