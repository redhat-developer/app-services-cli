package root

import (
	"flag"

	"github.com/MakeNowJust/heredoc"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/arguments"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/cmd/cluster"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/cmd/completion"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/cmd/factory"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/cmd/kafka"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/cmd/login"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/cmd/logout"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/cmd/serviceaccount"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

func NewRootCommand(cmdFactory *factory.Factory, version string) *cobra.Command {
	cmd := &cobra.Command{
		SilenceUsage:  true,
		SilenceErrors: true,
		Use:           "rhoas <command> <subcommand> [flags]",
		Short:         "RHOAS CLI",
		Long:          "Manage your application services directly from the command-line.",
		Example: heredoc.Doc(`
			# authenticate securely through your web browser
			$ rhoas login

			# create a Kafka instance
			$ rhoas kafka create --name "my-kafka-cluster"

			# create a service account with SASL/PLAIN credentials
			$ rhoas serviceaccount create -o json

			# connect your Kubernetes/OpenShift cluster to a service
			$ rhoas cluster connect
		`),
	}

	cmd.Version = version

	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)

	fs := cmd.PersistentFlags()
	arguments.AddDebugFlag(fs)

	// Child commands
	cmd.AddCommand(login.NewLoginCmd(cmdFactory))
	cmd.AddCommand(logout.NewLogoutCommand(cmdFactory))
	cmd.AddCommand(kafka.NewKafkaCommand(cmdFactory))
	cmd.AddCommand(serviceaccount.NewServiceAccountCommand(cmdFactory))
	cmd.AddCommand(cluster.NewClusterCommand(cmdFactory))
	cmd.AddCommand(completion.CompletionCmd)

	return cmd
}
