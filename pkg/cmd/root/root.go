package root

import (
	"github.com/MakeNowJust/heredoc"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/cmd/cluster"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/cmd/completion"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/cmd/docs"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/cmd/factory"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/cmd/kafka"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/cmd/login"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/cmd/logout"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/cmd/serviceaccount"
	"github.com/spf13/cobra"
)

func NewRootCommand(cmdFactory *factory.Factory, version string) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "rhoas <command> <subcommand> [flags]",
		Short: "rhoas cli",
		Long:  "Work with your Managed Services",
		Example: heredoc.Doc(`
			$ rhoas kafka create
			$ rhoas kafka list
			$ rhoas kafka use
		`),
	}

	cmd.Version = version

	// Child commands
	cmd.AddCommand(login.NewLoginCmd(cmdFactory))
	cmd.AddCommand(logout.NewLogoutCommand(cmdFactory))
	cmd.AddCommand(kafka.NewKafkaCommand(cmdFactory))
	cmd.AddCommand(serviceaccount.NewServiceAccountCommand(cmdFactory))
	cmd.AddCommand(docs.NewDocsCommand())
	cmd.AddCommand(cluster.NewClusterCommand(cmdFactory))
	cmd.AddCommand(completion.CompletionCmd)

	return cmd
}
