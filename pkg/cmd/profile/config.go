package profile

import (
	"github.com/MakeNowJust/heredoc"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/factory"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/profile/generate"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/profile/manage"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/profile/status"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/profile/view"

	"github.com/spf13/cobra"
)

func NewConfigCommand(f *factory.Factory) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "profile",
		Short: "Generates and manages services profiles",
		Long: heredoc.Doc(`
		Group of command that provide abstraction for managed services configuration.

		Config commands give users ability to:

		- Manage profiles of configurations containing multiple services
		Profiles 

		- Generate various service confiurations for local development and kubernetes
	
		Generates and manages services configurations that can be used 
		to configure your application deployments to be able to
		connect with RHOAS managed services.
		`),
		Example: ``,
		Args:    cobra.ExactArgs(1),
	}

	cmd.AddCommand(
		view.NewViewCommand(f),
		status.NewStatusCommand(f),
		manage.NewManageCommand(f),
		generate.NewGenerateCommand(f),
	)

	return cmd
}
