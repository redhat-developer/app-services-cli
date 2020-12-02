package root

import (
	"github.com/MakeNowJust/heredoc"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/cmd/completion"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/cmd/docs"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/cmd/kafka"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/cmd/login"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/cmd/logout"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/cmdutil"
	"github.com/spf13/cobra"
)

func NewRootCommand(f *cmdutil.Factory, version string) *cobra.Command {
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
	cmd.AddCommand(login.NewLoginCmd())
	cmd.AddCommand(logout.NewLogoutCommand())
	cmd.AddCommand(kafka.NewKafkaCommand(f))
	cmd.AddCommand(docs.NewDocsCommand())
	cmd.AddCommand(completion.CompletionCmd)

	return cmd
}
