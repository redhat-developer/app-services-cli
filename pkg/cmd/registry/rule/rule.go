package rule

import (
	"github.com/redhat-developer/app-services-cli/pkg/cmd/registry/rule/enable"
	"github.com/redhat-developer/app-services-cli/pkg/shared/factory"
	"github.com/spf13/cobra"
)

func NewRuleCommand(f *factory.Factory) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "rule",
		Short:   f.Localizer.MustLocalize("registry.rule.cmd.description.short"),
		Long:    f.Localizer.MustLocalize("registry.rule.cmd.description.long"),
		Example: f.Localizer.MustLocalize("registry.rule.cmd.example"),
		Args:    cobra.MinimumNArgs(1),
	}

	// add sub-commands
	cmd.AddCommand(
		enable.NewEnableCommand(f),
	)

	return cmd
}
