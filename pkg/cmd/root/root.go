package root

import (
	"flag"
	"fmt"

	"github.com/redhat-developer/app-services-cli/internal/localizer"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/status"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/whoami"

	"github.com/redhat-developer/app-services-cli/pkg/arguments"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/cluster"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/completion"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/factory"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/kafka"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/login"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/logout"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/serviceaccount"
	cliversion "github.com/redhat-developer/app-services-cli/pkg/cmd/version"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

func NewRootCommand(cmdFactory *factory.Factory, version string) *cobra.Command {
	cmd := &cobra.Command{
		SilenceUsage:  true,
		SilenceErrors: true,
		Use:           localizer.MustLocalizeFromID("root.cmd.use"),
		Short:         localizer.MustLocalizeFromID("root.cmd.shortDescription"),
		Long:          localizer.MustLocalizeFromID("root.cmd.longDescription"),
		Example:       localizer.MustLocalizeFromID("root.cmd.example"),
	}

	cmd.Version = version

	cmd.SetVersionTemplate(fmt.Sprintln(localizer.MustLocalize(&localizer.Config{
		MessageID: "version.cmd.outputText",
		TemplateData: map[string]interface{}{
			"Version": version,
		},
	})))

	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)

	fs := cmd.PersistentFlags()
	arguments.AddDebugFlag(fs)

	// this flag comes out of the box, but has its own basic usage text, so this overrides that
	var help bool
	fs.BoolVarP(&help, "help", "h", false, localizer.MustLocalizeFromID("root.cmd.flag.help.description"))

	// Child commands
	cmd.AddCommand(login.NewLoginCmd(cmdFactory))
	cmd.AddCommand(logout.NewLogoutCommand(cmdFactory))
	cmd.AddCommand(kafka.NewKafkaCommand(cmdFactory))
	cmd.AddCommand(serviceaccount.NewServiceAccountCommand(cmdFactory))
	cmd.AddCommand(cluster.NewClusterCommand(cmdFactory))
	cmd.AddCommand(status.NewStatusCommand(cmdFactory))
	cmd.AddCommand(completion.NewCompletionCommand(cmdFactory))
	cmd.AddCommand(whoami.NewWhoAmICmd(cmdFactory))
	cmd.AddCommand(cliversion.NewVersionCmd(cmdFactory))

	return cmd
}
