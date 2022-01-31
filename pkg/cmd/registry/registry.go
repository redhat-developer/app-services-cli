// Package registry REST API exposed via the serve command.
package registry

import (
	"github.com/redhat-developer/app-services-cli/internal/doc"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/registry/artifact"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/registry/artifact/role"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/registry/create"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/registry/delete"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/registry/describe"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/registry/list"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/registry/use"
	"github.com/redhat-developer/app-services-cli/pkg/shared/factory"
	"github.com/spf13/cobra"
)

func NewServiceRegistryCommand(f *factory.Factory) *cobra.Command {
	cmd := &cobra.Command{
		Use:         "service-registry",
		Annotations: map[string]string{doc.AnnotationName: "Service Registry commands"},
		Short:       f.Localizer.MustLocalize("registry.cmd.shortDescription"),
		Long:        f.Localizer.MustLocalize("registry.cmd.longDescription"),
		Example:     f.Localizer.MustLocalize("registry.cmd.example"),
		Args:        cobra.MinimumNArgs(1),
	}

	// add sub-commands
	cmd.AddCommand(
		create.NewCreateCommand(f),
		describe.NewDescribeCommand(f),
		delete.NewDeleteCommand(f),
		list.NewListCommand(f),
		use.NewUseCommand(f),
		artifact.NewArtifactsCommand(f),
		role.NewRoleCommand(f),
	)

	return cmd
}
