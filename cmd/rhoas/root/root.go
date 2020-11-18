package root

import (
	"github.com/MakeNowJust/heredoc"
	"github.com/bf2fc6cc711aee1a0c2a/cli/cmd/rhoas/completion"
	"github.com/bf2fc6cc711aee1a0c2a/cli/cmd/rhoas/docs"
	"github.com/bf2fc6cc711aee1a0c2a/cli/cmd/rhoas/kafka"
	"github.com/bf2fc6cc711aee1a0c2a/cli/cmd/rhoas/login"
	"github.com/bf2fc6cc711aee1a0c2a/cli/cmd/rhoas/logout"
	"github.com/spf13/cobra"
)

func NewRootCommand() *cobra.Command {
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

	cmd.AddCommand(login.NewLoginCmd())
	cmd.AddCommand(logout.NewLogoutCommand())
	cmd.AddCommand(kafka.NewKafkaCommand())
	cmd.AddCommand(docs.NewDocsCommand())
	cmd.AddCommand(completion.CompletionCmd)

	return cmd
}
