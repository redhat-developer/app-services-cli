package completion

import (
	"github.com/bf2fc6cc711aee1a0c2a/cli/internal/localizer"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/cmd/factory"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/iostreams"
	"github.com/spf13/cobra"
)

type Options struct {
	IO *iostreams.IOStreams
}

func NewCompletionCommand(f *factory.Factory) *cobra.Command {
	opts := &Options{
		IO: f.IOStreams,
	}

	cmd := &cobra.Command{
		Use:                   localizer.MustLocalizeFromID("completion.cmd.use"),
		Short:                 localizer.MustLocalizeFromID("completion.cmd.shortDescription"),
		Long:                  localizer.MustLocalizeFromID("completion.cmd.longDescription"),
		DisableFlagsInUseLine: true,
		ValidArgs:             []string{"bash", "zsh", "fish", "powershell"},
		Args:                  cobra.ExactValidArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			var err error
			switch args[0] {
			case "bash":
				err = cmd.Root().GenBashCompletion(opts.IO.Out)
			case "zsh":
				err = cmd.Root().GenZshCompletion(opts.IO.Out)
			case "fish":
				err = cmd.Root().GenFishCompletion(opts.IO.Out, true)
			case "powershell":
				err = cmd.Root().GenPowerShellCompletion(opts.IO.Out)
			}

			if err != nil {
				return err
			}

			return nil
		},
	}

	return cmd
}
