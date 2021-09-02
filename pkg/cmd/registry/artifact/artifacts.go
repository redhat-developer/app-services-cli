package artifact

import (
	"github.com/redhat-developer/app-services-cli/pkg/cmd/factory"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/registry/artifact/crud/create"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/registry/artifact/crud/delete"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/registry/artifact/crud/get"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/registry/artifact/crud/list"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/registry/artifact/crud/update"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/registry/artifact/download"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/registry/artifact/metadata"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/registry/artifact/versions"
	"github.com/spf13/cobra"
)

func NewArtifactsCommand(f *factory.Factory) *cobra.Command {
	cmd := &cobra.Command{
		Use:     f.Localizer.MustLocalize("artifact.cmd.use"),
		Short:   f.Localizer.MustLocalize("artifact.cmd.description.short"),
		Long:    f.Localizer.MustLocalize("artifact.cmd.description.long"),
		Example: f.Localizer.MustLocalize("artifact.cmd.example"),
		Args:    cobra.MinimumNArgs(1),
	}

	// add sub-commands
	cmd.AddCommand(
		// CRUD
		create.NewCreateCommand(f),
		get.NewGetCommand(f),
		delete.NewDeleteCommand(f),
		list.NewListCommand(f),
		update.NewUpdateCommand(f),

		// Misc
		metadata.NewGetMetadataCommand(f),
		metadata.NewSetMetadataCommand(f),
		versions.NewVersionsCommand(f),
		download.NewDownloadCommand(f),
	)

	return cmd
}
