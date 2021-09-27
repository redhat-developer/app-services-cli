package docs

import (
	"fmt"
	"time"

	"github.com/redhat-developer/app-services-cli/pkg/cmd/factory"
	rhoasdoc "github.com/redhat-developer/app-services-cli/pkg/doc"
	"github.com/redhat-developer/app-services-cli/pkg/logging"
	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"
)

const (
	markdown = "md"
	asciidoc = "adoc"
	man      = "man"
)

type options struct {
	dir    string
	format string

	logger logging.Logger
}

// NewDocsCmd creates a new docs command
func NewDocsCmd(f *factory.Factory) *cobra.Command {
	opts := options{
		logger: f.Logger,
	}

	cmd := &cobra.Command{
		Use:    "docs",
		Short:  "Generate documentation files in format of your choice",
		Hidden: true,
		Args:   cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			return runCmd(cmd, &opts)
		},
	}

	cmd.Flags().StringVar(&opts.format, "file-format", "md", "Output format of the generated documentation. Valid options are: 'md' (markdown), 'adoc' (Asciidoc) and 'man'")
	cmd.Flags().StringVar(&opts.dir, "dir", "./docs", "The directory to output the generated documentation files")

	return cmd
}

func runCmd(cmd *cobra.Command, opts *options) (err error) {
	cmd.Root().DisableAutoGenTag = true

	switch opts.format {
	case markdown:
		err = doc.GenMarkdownTree(cmd.Root(), opts.dir)
	case man:
		year := time.Now().Year()
		header := &doc.GenManHeader{
			Title:   "rhoas",
			Section: "1",
			Source:  fmt.Sprintf("Copyright (c) %v Red Hat, Inc.", year),
		}
		err = doc.GenManTree(cmd.Root(), header, opts.dir)
	case asciidoc:
		filePrepender := func(filename string) string { return "" }
		linkHandler := func(s string) string { return s }

		err = rhoasdoc.GenAsciidocTreeCustom(cmd.Root(), opts.dir, filePrepender, linkHandler)
	}

	if err != nil {
		return err
	}

	opts.logger.Infof("Documentation successfully generated into %v", opts.dir)

	return nil
}
