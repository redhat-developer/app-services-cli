package version

import (
	"context"
	"fmt"

	"github.com/redhat-developer/app-services-cli/internal/build"
	"github.com/redhat-developer/app-services-cli/internal/localizer"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/debug"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/factory"
	"github.com/redhat-developer/app-services-cli/pkg/iostreams"
	"github.com/redhat-developer/app-services-cli/pkg/locales"
	"github.com/redhat-developer/app-services-cli/pkg/logging"
	"github.com/spf13/cobra"
)

type Options struct {
	IO        *iostreams.IOStreams
	Logger    func() (logging.Logger, error)
	localizer locales.Localizer
}

func NewVersionCmd(f *factory.Factory) *cobra.Command {
	opts := &Options{
		IO:     f.IOStreams,
		Logger: f.Logger,
	}

	cmd := &cobra.Command{
		Use:     localizer.MustLocalizeFromID("version.cmd.use"),
		Hidden:  true,
		Short:   localizer.MustLocalizeFromID("version.cmd.shortDescription"),
		Example: localizer.MustLocalizeFromID("whoami.cmd.example"),
		Args:    cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			return runCmd(opts)
		},
	}

	return cmd
}

func runCmd(opts *Options) (err error) {
	fmt.Fprintln(opts.IO.Out, opts.localizer.LoadMessageWithConfig("version.cmd.outputText", &locales.LocalizerConfig{
		TemplateData: map[string]interface{}{
			"Version": build.Version,
		},
	}))

	logger, err := opts.Logger()
	if err != nil {
		return nil
	}

	// debug mode checks this for a version update also.
	// so we check if is enabled first so as not to print it twice
	if !debug.Enabled() {
		build.CheckForUpdate(context.Background(), logger)
	}
	return nil
}
