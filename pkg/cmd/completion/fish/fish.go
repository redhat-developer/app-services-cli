package fish

import (
	"github.com/bf2fc6cc711aee1a0c2a/cli/internal/localizer"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/cmd/factory"
	"github.com/spf13/cobra"
)

func NewCommand(f *factory.Factory) *cobra.Command {
	cmd := &cobra.Command{
		Use:                   localizer.MustLocalizeFromID("completion.fish.cmd.use"),
		Short:                 localizer.MustLocalizeFromID("completion.fish.cmd.shortDescription"),
		Long:                  localizer.MustLocalizeFromID("completion.fish.cmd.longDescription"),
		DisableFlagsInUseLine: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmd.Parent().Parent().GenFishCompletion(f.IOStreams.Out, false)
		},
	}

	return cmd
}
