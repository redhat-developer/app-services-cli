package connector_type

import (
	"github.com/redhat-developer/app-services-cli/internal/doc"

	"github.com/redhat-developer/app-services-cli/pkg/shared/factory"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/connector/connector_type/list"
	"github.com/spf13/cobra"
)

// NewTypeCommand creates a new command to use different connector types
func NewTypeCommand(f *factory.Factory) *cobra.Command {
	cmd := &cobra.Command{
		Use:         "type",
		Annotations: map[string]string{doc.AnnotationName: "Connectors commands"},
		Short:       f.Localizer.MustLocalize("connector.cmd.shortDescription"),
		Long:        f.Localizer.MustLocalize("connector.cmd.longDescription"),
		Example:     f.Localizer.MustLocalize("connector.cmd.example"),
		Args:        cobra.MinimumNArgs(1),
	}

	// add sub-commands
	cmd.AddCommand(
		list.NewListCommand(f),
	)

	return cmd
}
