package serviceaccount

import (
	"github.com/MakeNowJust/heredoc"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/cmd/factory"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/cmd/serviceaccount/create"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/cmd/serviceaccount/delete"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/cmd/serviceaccount/list"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/cmd/serviceaccount/resetcredentials"
	"github.com/spf13/cobra"
)

// NewServiceAccountCommand creates a new command sub-group to manage service accounts
func NewServiceAccountCommand(f *factory.Factory) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "serviceaccount",
		Short: "Create, list, delete and update service accounts",
		Long: heredoc.Doc(`
			Use these commands to create, delete, and list service accounts. You can also reset the credentials for service account.

			Service accounts include SASL/PLAIN credentials which can then be used to authenticate applications and tools with your services.
		`),
		Example: heredoc.Doc(`
			# create a service account
			$ rhoas serviceaccount creare

			# list service accounts
			$ rhoas serviceaccount list

			# delete a service account
			$ rhoas serviceaccount delete --id "173c1ad9-932d-4007-ae0f-4da74f4d2ccd"

			# reset credentials on a service account
			$ rhoas serviceaccount reset-credentials
		`),
		Args: cobra.ExactArgs(1),
	}

	cmd.AddCommand(
		create.NewCreateCommand(f),
		list.NewListCommand(f),
		delete.NewDeleteCommand(f),
		resetcredentials.NewResetCredentialsCommand(f),
	)

	return cmd
}
