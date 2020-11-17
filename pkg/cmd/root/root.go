package root

import (
	"github.com/MakeNowJust/heredoc"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/cmd/authz"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/cmd/completion"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/cmd/docs"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/cmd/kafka"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/cmd/login"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/cmd/logout"
	"github.com/spf13/cobra"
)

func NewCmdRoot() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "rhmas <command> <subcommand> [flags]",
		Short: "RHMAS CLI",
		Long:  "Work with your Managed Services",
		Example: heredoc.Doc(`
			$ rhmas kafka create
			$ rhmas kafka list
			$ rhmas kafka use
		`),
	}

	cmd.AddCommand(login.NewLoginCmd())
	cmd.AddCommand(logout.NewLogoutCommand())
	cmd.AddCommand(kafka.NewKafkaCommand())
	cmd.AddCommand(authz.NewAuthGroupCommand())
	cmd.AddCommand(docs.NewDocsCommand())
	cmd.AddCommand(completion.CompletionCmd)

	return cmd
}
