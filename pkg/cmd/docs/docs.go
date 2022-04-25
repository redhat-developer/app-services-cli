package docs

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"time"

	rhoasdoc "github.com/redhat-developer/app-services-cli/internal/doc"
	"github.com/redhat-developer/app-services-cli/pkg/core/logging"
	"github.com/redhat-developer/app-services-cli/pkg/shared/factory"

	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"
)

const (
	markdown = "md"
	asciidoc = "adoc"
	man      = "man"
)

type options struct {
	dir        string
	format     string
	filePrefix string
	logger     logging.Logger
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
	cmd.Flags().StringVar(&opts.filePrefix, "file-prefix", "ref-cli-", "Prefix for each documentation file")
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
		docsDir := path.Clean(opts.dir + "/modules")
		assembliesDir := path.Clean(opts.dir + "/assemblies")
		err = os.MkdirAll(docsDir, 0755)
		if err != nil {
			return err
		}
		err = os.MkdirAll(assembliesDir, 0755)
		if err != nil {
			return err
		}

		options := rhoasdoc.GeneratorOptions{
			Dir:           docsDir,
			GenerateIndex: true,
			IndexFile:     assembliesDir + "/assembly-cli-command-reference.adoc"}

		options.FileNameGenerator = func(c *cobra.Command) string {
			basename := opts.filePrefix + rhoasdoc.GetNormalizedCommandPath(c) + ".adoc"
			return filepath.Join(options.Dir, basename)
		}

		err = rhoasdoc.GenAsciidocTree(cmd.Root(), &options)
	}

	if err != nil {
		return err
	}

	opts.logger.Infof("Documentation successfully generated into %v", opts.dir)

	return nil
}
