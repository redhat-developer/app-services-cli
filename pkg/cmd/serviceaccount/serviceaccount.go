package serviceaccount

import (
	"github.com/redhat-developer/app-services-cli/pkg/cmd/factory"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/serviceaccount/create"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/serviceaccount/delete"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/serviceaccount/describe"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/serviceaccount/list"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/serviceaccount/resetcredentials"
	"github.com/spf13/cobra"
)

// NewServiceAccountCommand creates a new command sub-group to manage service accounts
func NewServiceAccountCommand(f *factory.Factory) *cobra.Command {
	cmd := &cobra.Command{
		Use:   f.Localizer.MustLocalize("serviceAccount.cmd.use"),
		Short: f.Localizer.MustLocalize("serviceAccount.cmd.shortDescription"),
		Long:  f.Localizer.MustLocalize("serviceAccount.cmd.longDescription"),
		Args:  cobra.ExactArgs(1),
	}

	cmd.AddCommand(
		create.NewCreateCommand(f),
		list.NewListCommand(f),
		delete.NewDeleteCommand(f),
		resetcredentials.NewResetCredentialsCommand(f),
		describe.NewDescribeCommand(f),
	)

	return cmd
}
