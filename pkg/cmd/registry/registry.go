// Package kafka instance contains commands for interacting with cluster logic of the service directly instead of through the
// REST API exposed via the serve command.
package registry

import (
	"github.com/redhat-developer/app-services-cli/pkg/cmd/factory"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/registry/create"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/registry/delete"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/registry/describe"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/registry/list"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/registry/use"
	"github.com/redhat-developer/app-services-cli/pkg/profile"
	"github.com/spf13/cobra"
)

var example string = `

`

func NewServiceRegistryCommand(f *factory.Factory) *cobra.Command {
	cmd := &cobra.Command{
		Use:         "registry",
		Hidden:      !profile.DevPreviewEnabled(f),
		Annotations: profile.DevPreviewAnnotation(),

		Short:   profile.ApplyDevPreviewLabel(f.Localizer.MustLocalize("registry.cmd.shortDescription")),
		Long:    f.Localizer.MustLocalize("registry.cmd.longDescription"),
		Example: f.Localizer.MustLocalize("registry.cmd.example"),
		Args:    cobra.MinimumNArgs(1),
	}

	// add sub-commands
	cmd.AddCommand(
		create.NewCreateCommand(f),
		describe.NewDescribeCommand(f),
		delete.NewDeleteCommand(f),
		list.NewListCommand(f),
		use.NewUseCommand(f),
	)

	return cmd
}
