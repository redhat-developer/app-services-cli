package rule

import (
	"github.com/redhat-developer/app-services-cli/pkg/cmd/registry/rule/describe"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/registry/rule/disable"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/registry/rule/enable"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/registry/rule/list"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/registry/rule/update"
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
		list.NewListCommand(f),
		describe.NewDescribeCommand(f),
		update.NewUpdateCommand(f),
		disable.NewDisableCommand(f),
	)

	return cmd
}
