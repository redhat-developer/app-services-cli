package version

import (
	"context"
	"fmt"

	"github.com/redhat-developer/app-services-cli/internal/build"
	"github.com/redhat-developer/app-services-cli/pkg/core/cmdutil/factory"
	"github.com/redhat-developer/app-services-cli/pkg/core/cmdutil/flagutil"
	"github.com/redhat-developer/app-services-cli/pkg/core/ioutil/iostreams"
	"github.com/redhat-developer/app-services-cli/pkg/core/localize"
	"github.com/redhat-developer/app-services-cli/pkg/core/logging"
	"github.com/spf13/cobra"
)

type options struct {
	IO        *iostreams.IOStreams
	Logger    logging.Logger
	localizer localize.Localizer
	Context   context.Context
}

func NewVersionCmd(f *factory.Factory) *cobra.Command {
	opts := &options{
		IO:        f.IOStreams,
		Logger:    f.Logger,
		localizer: f.Localizer,
		Context:   f.Context,
	}

	cmd := &cobra.Command{
		Use:    "version",
		Short:  opts.localizer.MustLocalize("version.cmd.shortDescription"),
		Hidden: true,
		Args:   cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			return runCmd(opts)
		},
	}

	return cmd
}

func runCmd(opts *options) (err error) {
	fmt.Fprintln(opts.IO.Out, opts.localizer.MustLocalize("version.cmd.outputText", localize.NewEntry("Version", build.Version)))

	// debug mode checks this for a version update also.
	// so we check if is enabled first so as not to print it twice
	if !flagutil.DebugEnabled() {
		build.CheckForUpdate(opts.Context, build.Version, opts.Logger, opts.localizer)
	}
	return nil
}
