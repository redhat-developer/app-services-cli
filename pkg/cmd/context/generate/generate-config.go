package generate

import (
	"context"

	"github.com/redhat-developer/app-services-cli/pkg/cmd/context/contextcmdutil"
	"github.com/redhat-developer/app-services-cli/pkg/core/cmdutil/flagutil"
	"github.com/redhat-developer/app-services-cli/pkg/core/ioutil/iostreams"
	"github.com/redhat-developer/app-services-cli/pkg/core/localize"
	"github.com/redhat-developer/app-services-cli/pkg/core/logging"
	"github.com/redhat-developer/app-services-cli/pkg/core/servicecontext"
	"github.com/redhat-developer/app-services-cli/pkg/shared/contextutil"
	"github.com/redhat-developer/app-services-cli/pkg/shared/factory"
	"github.com/spf13/cobra"
)

type options struct {
	IO             *iostreams.IOStreams
	Logger         logging.Logger
	Connection     factory.ConnectionFunc
	localizer      localize.Localizer
	Context        context.Context
	ServiceContext servicecontext.IContext

	name       string
	configType string
}

// NewGenerateCommand creates configuration files for service context
func NewGenerateCommand(f *factory.Factory) *cobra.Command {

	opts := &options{
		Connection:     f.Connection,
		IO:             f.IOStreams,
		Logger:         f.Logger,
		localizer:      f.Localizer,
		ServiceContext: f.ServiceContext,
	}

	cmd := &cobra.Command{
		Use:     "generate-config",
		Short:   f.Localizer.MustLocalize("context.generate.cmd.shortDescription"),
		Long:    f.Localizer.MustLocalize("context.generate.cmd.longDescription"),
		Example: f.Localizer.MustLocalize("context.generate.cmd.example"),
		Args:    cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {

			if !opts.IO.CanPrompt() && opts.name == "" {
				return flagutil.RequiredWhenNonInteractiveError("name")
			}

			return runGenerate(opts)
		},
	}

	flags := contextcmdutil.NewFlagSet(cmd, f)
	flags.AddContextName(&opts.name)
	flags.StringVar(&opts.configType, "type", "", opts.localizer.MustLocalize("context.generate.flag.type"))

	return cmd
}

func runGenerate(opts *options) error {

	svcContext, err := opts.ServiceContext.Load()
	if err != nil {
		return err
	}

	var svcConfig *servicecontext.ServiceConfig

	if opts.name == "" {
		svcConfig, err = contextutil.GetCurrentContext(svcContext, opts.localizer)
		if err != nil {
			return err
		}
	} else {
		svcConfig, err = contextutil.GetContext(svcContext, opts.localizer, opts.name)
		if err != nil {
			return err
		}
	}

	if err = BuildConfiguration(svcConfig, opts); err != nil {
		return err
	}

	return nil
}
