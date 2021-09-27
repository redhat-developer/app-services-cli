package root

import (
	"github.com/redhat-developer/app-services-cli/pkg/arguments"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/debug"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/docs"

	"github.com/redhat-developer/app-services-cli/pkg/cmd/registry"

	"github.com/redhat-developer/app-services-cli/pkg/cmd/login"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/status"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/whoami"

	"github.com/redhat-developer/app-services-cli/pkg/cmd/cluster"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/completion"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/factory"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/kafka"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/logout"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/serviceaccount"
	cliversion "github.com/redhat-developer/app-services-cli/pkg/cmd/version"
	"github.com/spf13/cobra"
)

func NewRootCommand(f *factory.Factory, version string) *cobra.Command {
	cmd := &cobra.Command{
		SilenceUsage:  true,
		SilenceErrors: true,
		Use:           "rhoas <command> <subcommand> [flags]",
		Short:         f.Localizer.MustLocalize("root.cmd.shortDescription"),
		Long:          f.Localizer.MustLocalize("root.cmd.longDescription"),
		Example:       f.Localizer.MustLocalize("root.cmd.example"),
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			f.Logger.SetDebug(debug.Enabled())
		},
	}
	fs := cmd.PersistentFlags()
	arguments.AddDebugFlag(fs)
	// this flag comes out of the box, but has its own basic usage text, so this overrides that
	var help bool

	fs.BoolVarP(&help, "help", "h", false, f.Localizer.MustLocalize("root.cmd.flag.help.description"))
	cmd.Flags().Bool("version", false, f.Localizer.MustLocalize("root.cmd.flag.version.description"))

	cmd.Version = version

	// pflag.CommandLine.AddGoFlagSet(flag.CommandLine)

	// Child commands
	cmd.AddCommand(login.NewLoginCmd(f))
	cmd.AddCommand(logout.NewLogoutCommand(f))
	cmd.AddCommand(kafka.NewKafkaCommand(f))
	cmd.AddCommand(serviceaccount.NewServiceAccountCommand(f))
	cmd.AddCommand(cluster.NewClusterCommand(f))
	cmd.AddCommand(status.NewStatusCommand(f))
	cmd.AddCommand(completion.NewCompletionCommand(f))
	cmd.AddCommand(whoami.NewWhoAmICmd(f))
	cmd.AddCommand(cliversion.NewVersionCmd(f))

	// Registry commands
	cmd.AddCommand(registry.NewServiceRegistryCommand(f))

	cmd.AddCommand(docs.NewDocsCmd(f))

	return cmd
}
