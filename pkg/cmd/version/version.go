package version

import (
	"fmt"

	"github.com/bf2fc6cc711aee1a0c2a/cli/internal/build"
	"github.com/bf2fc6cc711aee1a0c2a/cli/internal/localizer"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/cmd/factory"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/iostreams"
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
