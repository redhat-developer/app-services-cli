package dedicated

import (
	"github.com/redhat-developer/app-services-cli/pkg/cmd/dedicated/deregister"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/dedicated/list"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/dedicated/register"
	"github.com/redhat-developer/app-services-cli/pkg/shared/factory"
	"github.com/spf13/cobra"
)

func NewDedicatedCmd(f *factory.Factory) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "dedicated",
		Short:   f.Localizer.MustLocalize("dedicated.cmd.shortDescription"),
		Long:    f.Localizer.MustLocalize("dedicated.cmd.longDescription"),
		Example: f.Localizer.MustLocalize("dedicated.cmd.example"),
	}

	cmd.AddCommand(
		register.NewRegisterClusterCommand(f),
		list.NewListClusterCommand(f),
		deregister.NewDeRegisterClusterCommand(f),
	)

	cmd.Hidden = true

	return cmd
}
