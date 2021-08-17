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
		Use:   "artifact",
		Short: "Manage Service Registry Artifacts",
		Long: `
Apicurio Registry Artifacts enables developers to manage and share the structure of their data. 
For example, client applications can dynamically push or pull the latest updates to or from the registry without needing to redeploy.
Apicurio Registry also enables developers to create rules that govern how registry content can evolve over time. 
For example, this includes rules for content validation and version compatibility.

Registry commands enable client applications to manage the artifacts in the registry. 
This set of commands provide create, read, update, and delete operations for schema and API artifacts, rules, versions, and metadata.
`,
		Example: `
## Create artifact in my-group from schema.json file
rhoas service-registry artifact create --artifact-id=my-artifact --group=my-group artifact.json

## Get artifact content
rhoas service-registry artifact get --artifact-id=my-artifact --group=my-group file.json 

## Delete artifact
rhoas service-registry artifact delete --artifact-id=my-artifact

## Get artifact metadata
rhoas service-registry artifact metadata --artifact-id=my-artifact --group=my-group

## Update artifact
rhoas service-registry artifact update --artifact-id=my-artifact artifact-new.json

## List Artifacts
rhoas service-registry artifact list --group=my-group --limit=10 page=1

## View artifact versions
rhoas service-registry artifact versions --artifact-id=my-artifact --group=my-group
		`,
		Args: cobra.MinimumNArgs(1),
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
		metadata.NewMetadataCommand(f),
		versions.NewVersionsCommand(f),
		download.NewDownloadCommand(f),
	)

	return cmd
}
