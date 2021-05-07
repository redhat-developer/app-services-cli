package completion

import (
	"github.com/redhat-developer/app-services-cli/pkg/cmd/completion/bash"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/completion/fish"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/completion/zsh"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/factory"
	"github.com/spf13/cobra"
)

func NewCompletionCommand(f *factory.Factory) *cobra.Command {
	cmd := &cobra.Command{
		Use:   f.Localizer.MustLocalize("completion.cmd.use"),
		Short: f.Localizer.MustLocalize("completion.cmd.shortDescription"),
		Long:  f.Localizer.MustLocalize("completion.cmd.longDescription"),
		Args:  cobra.ExactArgs(1),
	}

	cmd.AddCommand(
		bash.NewCommand(f),
		zsh.NewCommand(f),
		fish.NewCommand(f),
	)

	return cmd
}
