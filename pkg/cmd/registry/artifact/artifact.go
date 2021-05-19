package artifact

import (
	"github.com/spf13/cobra"

	"github.com/redhat-developer/app-services-cli/pkg/cmd/factory"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/registry/artifact/create"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/registry/artifact/delete"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/registry/artifact/fetch"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/registry/artifact/list"
)

const (
	Name      = "name"
	Operation = "operation"
)

// NewArtifactsCommand uses currently configured Registry to create artifacts
func NewArtifactCommand(f *factory.Factory) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "artifact",
		Short: "Managed Service Registry groups and artifacts",
		Long:  "",
	}

	cmd.AddCommand(
		create.NewCreateCommand(f),
		delete.NewDeleteCommand(f),
		list.NewListCommand(f),
		fetch.NewFetchCommand(f),
	)

	return cmd
}
