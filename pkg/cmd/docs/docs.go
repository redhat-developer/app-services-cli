package docs

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/browser"
)

var flags struct {
	generate bool
	browser  bool
}

// NewDocsCommand opens docs
func NewDocsCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "docs",
		Short: "Open documentation in browser",
		Long:  "Opens the Red Hat Managed Application Services CLI documentation in your default web browser",
		Run:   runDocs,
	}

	cmd.Flags().BoolVar(&flags.browser, "browser", true, "Open documentation in your browser")

	return cmd
}

func runDocs(cmd *cobra.Command, _ []string) {
	if flags.browser {
		fmt.Fprintln(os.Stderr, "Opening documentation in your browser")
		browsercmd, err := browser.GetOpenBrowserCommand("http://localhost:3000/docs/commands/rhmas")
		if err != nil {
			fmt.Fprint(os.Stderr, err)
			os.Exit(1)
		}
		_ = browsercmd.Start()
	}
}
