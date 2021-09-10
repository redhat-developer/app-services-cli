package version

import (
	"context"
	"fmt"

	"github.com/redhat-developer/app-services-cli/internal/build"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/debug"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/factory"
	"github.com/redhat-developer/app-services-cli/pkg/iostreams"
	"github.com/redhat-developer/app-services-cli/pkg/localize"
	"github.com/redhat-developer/app-services-cli/pkg/logging"
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
	if !debug.Enabled() {
		build.CheckForUpdate(opts.Context, opts.Logger, opts.localizer)
	}
	return nil
}
