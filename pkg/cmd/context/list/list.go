package list

import (
	"context"
	"fmt"

	"github.com/redhat-developer/app-services-cli/pkg/core/ioutil/icon"
	"github.com/redhat-developer/app-services-cli/pkg/core/ioutil/iostreams"
	"github.com/redhat-developer/app-services-cli/pkg/core/localize"
	"github.com/redhat-developer/app-services-cli/pkg/core/logging"
	"github.com/redhat-developer/app-services-cli/pkg/core/servicecontext"
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
}

// NewListCommand creates a new command to list available contexts
func NewListCommand(f *factory.Factory) *cobra.Command {

	opts := &options{
		Connection:     f.Connection,
		IO:             f.IOStreams,
		Logger:         f.Logger,
		localizer:      f.Localizer,
		ServiceContext: f.ServiceContext,
	}

	cmd := &cobra.Command{
		Use:     "list",
		Short:   f.Localizer.MustLocalize("context.list.cmd.shortDescription"),
		Long:    f.Localizer.MustLocalize("context.list.cmd.longDescription"),
		Example: f.Localizer.MustLocalize("context.list.cmd.example"),
		Args:    cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return runList(opts)
		},
	}

	return cmd
}

func runList(opts *options) error {

	svcContext, err := opts.ServiceContext.Load()
	if err != nil {
		return err
	}

	profiles := svcContext.Contexts

	if profiles == nil {
		profiles = make(map[string]servicecontext.ServiceConfig)
	}

	currentCtx := svcContext.CurrentContext

	var profileList string

	for name := range profiles {
		if currentCtx != "" && name == currentCtx {
			profileList += fmt.Sprintln(name, icon.SuccessPrefix())
		} else {
			profileList += fmt.Sprintln(name)
		}
	}

	opts.Logger.Info(profileList)
	opts.Logger.Info(opts.localizer.MustLocalize("context.list.log.info.describeHint"))

	return nil
}
