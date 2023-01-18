package dedicated

import (
	"github.com/redhat-developer/app-services-cli/pkg/cmd/dedicated/register"
	"github.com/redhat-developer/app-services-cli/pkg/shared/factory"
	"github.com/spf13/cobra"
)

// TODO add localizer and descriptions
func NewDedicatedCmd(f *factory.Factory) *cobra.Command {
	cmd := &cobra.Command{
		Use: "dedicated",
		// Short:   f.Localizer.MustLocalize("kafka.topic.cmd.shortDescription"),
		Short: "shortDescription",
		// Long:    f.Localizer.MustLocalize("kafka.topic.cmd.longDescription"),
		Long: "longDescription",
		// Example: f.Localizer.MustLocalize("kafka.topic.cmd.example"),
		Example: "example",
	}

	cmd.AddCommand(
		register.NewRegisterClusterCommand(f),
	)

	return cmd
}
