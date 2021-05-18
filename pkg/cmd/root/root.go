package root

import (
	"flag"

	"github.com/redhat-developer/app-services-cli/pkg/cmd/registry"
	"github.com/redhat-developer/app-services-cli/pkg/profile"

	"github.com/redhat-developer/app-services-cli/internal/build"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/login"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/status"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/whoami"
	"github.com/redhat-developer/app-services-cli/pkg/localize"

	"github.com/redhat-developer/app-services-cli/pkg/arguments"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/cluster"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/completion"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/factory"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/kafka"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/logout"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/serviceaccount"
	cliversion "github.com/redhat-developer/app-services-cli/pkg/cmd/version"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

func NewRootCommand(f *factory.Factory, version string) *cobra.Command {
	var devpreview string

	cmd := &cobra.Command{
		SilenceUsage:  true,
		SilenceErrors: true,
		Use:           f.Localizer.MustLocalize("root.cmd.use"),
		Short:         f.Localizer.MustLocalize("root.cmd.shortDescription"),
		Long:          f.Localizer.MustLocalize("root.cmd.longDescription"),
		Example:       f.Localizer.MustLocalize("root.cmd.example"),
		RunE: func(cmd *cobra.Command, _ []string) error {
			if devpreview == "" {
				return cmd.Help()
			}

			_, err := profile.EnableDevPreview(f, &devpreview)
			return err
		},
	}
	fs := cmd.PersistentFlags()
	arguments.AddDebugFlag(fs)
	// this flag comes out of the box, but has its own basic usage text, so this overrides that
	var help bool
	fs.BoolVarP(&help, "help", "h", false, f.Localizer.MustLocalize("root.cmd.flag.help.description"))
	cmd.Flags().StringVarP(&devpreview, "devpreview", "", "", f.Localizer.MustLocalize("root.cmd.flag.devpreview.description"))

	cmd.Version = version

	cmd.SetVersionTemplate(f.Localizer.MustLocalize("version.cmd.outputText", localize.NewEntry("Version", build.Version)))
	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)

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

	// Dev preview commands
	if profile.DevPreviewEnabled(f) {
		cmd.AddCommand(registry.NewServiceRegistryCommand(f))
	}

	return cmd
}
