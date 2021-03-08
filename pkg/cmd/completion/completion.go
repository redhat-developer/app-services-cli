package completion

import (
	"errors"

	"github.com/bf2fc6cc711aee1a0c2a/cli/internal/localizer"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/cmd/completion/bash"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/cmd/completion/fish"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/cmd/completion/zsh"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/cmd/factory"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/iostreams"
	"github.com/spf13/cobra"
)

type Options struct {
	IO *iostreams.IOStreams
}

func NewCompletionCommand(f *factory.Factory) *cobra.Command {
	cmd := &cobra.Command{
		Use:   localizer.MustLocalizeFromID("completion.cmd.use"),
		Short: localizer.MustLocalizeFromID("completion.cmd.shortDescription"),
		Long:  localizer.MustLocalizeFromID("completion.cmd.longDescription"),
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				return errors.New(localizer.MustLocalizeFromID("completion.cmd.error.subcommandRequired"))
			}
			return nil
		},
	}

	cmd.AddCommand(
		bash.NewCommand(f),
		zsh.NewCommand(f),
		fish.NewCommand(f),
	)

	return cmd
}
