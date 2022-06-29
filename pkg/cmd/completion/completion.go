package completion

import (
	"github.com/spf13/cobra"

	"github.com/redhat-developer/app-services-cli/pkg/cmd/completion/bash"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/completion/fish"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/completion/powershell"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/completion/zsh"
	"github.com/redhat-developer/app-services-cli/pkg/shared/factory"
)

func NewCompletionCommand(f *factory.Factory) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "completion",
		Short:   f.Localizer.MustLocalize("completion.cmd.shortDescription"),
		Long:    f.Localizer.MustLocalize("completion.cmd.longDescription"),
		Example: f.Localizer.MustLocalize("completion.cmd.example"),
		Args:    cobra.ExactArgs(1),
	}

	cmd.AddCommand(
		bash.NewCommand(f),
		zsh.NewCommand(f),
		fish.NewCommand(f),
		powershell.NewCommand(f),
	)

	return cmd
}
