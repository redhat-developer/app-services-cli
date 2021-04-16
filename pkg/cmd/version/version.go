package version

import (
	"fmt"

	"github.com/redhat-developer/app-services-cli/internal/build"
	"github.com/redhat-developer/app-services-cli/internal/localizer"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/factory"
	"github.com/redhat-developer/app-services-cli/pkg/iostreams"
	"github.com/spf13/cobra"
)

type Options struct {
	IO *iostreams.IOStreams
}

func NewVersionCmd(f *factory.Factory) *cobra.Command {
	opts := &Options{
		IO: f.IOStreams,
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
	fmt.Fprintln(opts.IO.Out, localizer.MustLocalize(&localizer.Config{
		MessageID: "version.cmd.outputText",
		TemplateData: map[string]interface{}{
			"Version": build.Version,
		},
	}))
	return nil
}
