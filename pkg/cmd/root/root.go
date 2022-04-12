package root

import (
	"github.com/redhat-developer/app-services-cli/pkg/cmd/cluster"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/completion"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/context"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/docs"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/generate"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/kafka"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/login"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/logout"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/registry"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/request"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/serviceaccount"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/status"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/token"
	cliversion "github.com/redhat-developer/app-services-cli/pkg/cmd/version"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/whoami"
	"github.com/redhat-developer/app-services-cli/pkg/core/cmdutil/flagutil"
	"github.com/redhat-developer/app-services-cli/pkg/shared/factory"
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
	}

	fs := cmd.PersistentFlags()
	flagutil.VerboseFlag(fs)
	// this flag comes out of the box, but has its own basic usage text, so this overrides that
	var help bool

	fs.BoolVarP(&help, "help", "h", false, f.Localizer.MustLocalize("root.cmd.flag.help.description"))
	cmd.Flags().Bool("version", false, f.Localizer.MustLocalize("root.cmd.flag.version.description"))

	cmd.Version = version

	// Child commands
	cmd.AddCommand(login.NewLoginCmd(f))
	cmd.AddCommand(logout.NewLogoutCommand(f))
	cmd.AddCommand(kafka.NewKafkaCommand(f))
	cmd.AddCommand(serviceaccount.NewServiceAccountCommand(f))
	cmd.AddCommand(cluster.NewClusterCommand(f))
	cmd.AddCommand(status.NewStatusCommand(f))
	cmd.AddCommand(generate.NewGenerateCommand(f))
	cmd.AddCommand(completion.NewCompletionCommand(f))
	cmd.AddCommand(whoami.NewWhoAmICmd(f))
	cmd.AddCommand(cliversion.NewVersionCmd(f))
	cmd.AddCommand(token.NewAuthTokenCmd(f))
	// Registry commands
	cmd.AddCommand(registry.NewServiceRegistryCommand(f))

	cmd.AddCommand(docs.NewDocsCmd(f))
	cmd.AddCommand(request.NewCallCmd(f))
	cmd.AddCommand(context.NewContextCmd(f))

	return cmd
}
