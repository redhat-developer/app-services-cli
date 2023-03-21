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
		Use:     "openshift-cluster",
		Short:   f.Localizer.MustLocalize("kafka.openshiftCluster.cmd.shortDescription"),
		Long:    f.Localizer.MustLocalize("kafka.openshiftCluster.cmd.longDescription"),
		Example: f.Localizer.MustLocalize("kafka.openshiftCluster.cmd.example"),
	}

	cmd.Aliases = append(cmd.Aliases, "oc")

	cmd.AddCommand(
		register.NewRegisterClusterCommand(f),
		list.NewListClusterCommand(f),
		deregister.NewDeRegisterClusterCommand(f),
	)

	cmd.Hidden = true

	return cmd
}
