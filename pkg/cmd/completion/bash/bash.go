package bash

import (
	"github.com/redhat-developer/app-services-cli/internal/localizer"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/factory"
	"github.com/spf13/cobra"
)

func NewCommand(f *factory.Factory) *cobra.Command {
	cmd := &cobra.Command{
		Use:                   localizer.MustLocalizeFromID("completion.bash.cmd.use"),
		Short:                 localizer.MustLocalizeFromID("completion.bash.cmd.shortDescription"),
		Long:                  localizer.MustLocalizeFromID("completion.bash.cmd.longDescription"),
		DisableFlagsInUseLine: true,
		Args:                  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmd.Parent().Parent().GenBashCompletion(f.IOStreams.Out)
		},
	}

	return cmd
}
