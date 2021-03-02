package serviceaccount

import (
	"github.com/bf2fc6cc711aee1a0c2a/cli/internal/localizer"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/cmd/factory"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/cmd/serviceaccount/create"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/cmd/serviceaccount/delete"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/cmd/serviceaccount/describe"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/cmd/serviceaccount/list"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/cmd/serviceaccount/resetcredentials"
	"github.com/spf13/cobra"
)

// NewServiceAccountCommand creates a new command sub-group to manage service accounts
func NewServiceAccountCommand(f *factory.Factory) *cobra.Command {
	cmd := &cobra.Command{
		Use:   localizer.MustLocalizeFromID("serviceAccount.cmd.use"),
		Short: localizer.MustLocalizeFromID("serviceAccount.cmd.shortDescription"),
		Long:  localizer.MustLocalizeFromID("serviceAccount.cmd.longDescription"),
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
