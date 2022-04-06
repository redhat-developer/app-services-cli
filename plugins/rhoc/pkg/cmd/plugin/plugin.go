package plugin

import (
	"github.com/redhat-developer/app-services-cli/pkg/cmd/completion"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/login"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/logout"
	"github.com/redhat-developer/app-services-cli/pkg/core/cmdutil/flagutil"
	"github.com/redhat-developer/app-services-cli/pkg/shared/factory"
	"github.com/redhat-developer/app-services-cli/plugins/request/pkg/cmd/request"
	"github.com/spf13/cobra"
)

func NewPluginCommand(f *factory.Factory, version string) *cobra.Command {
	cmd := &cobra.Command{
		SilenceUsage:  true,
		SilenceErrors: true,
		Use:           "rhrequest <command> <subcommand> [flags]",
		Short:         "Request plugin",
		Example:       "rhrequest request -h",
	}

	fs := cmd.PersistentFlags()
	flagutil.VerboseFlag(fs)
	// this flag comes out of the box, but has its own basic usage text, so this overrides that
	var help bool

	// TODO - how to handle internationalization from multiple sources?
	fs.BoolVarP(&help, "help", "h", false, f.Localizer.MustLocalize("root.cmd.flag.help.description"))

	// TODO - wrap into SDK (Auth commands)
	cmd.AddCommand(login.NewLoginCmd(f))
	cmd.AddCommand(logout.NewLogoutCommand(f))
	cmd.AddCommand(completion.NewCompletionCommand(f))

	// Plugin command
	cmd.AddCommand(request.NewRequestCmd(f))

	return cmd
}
