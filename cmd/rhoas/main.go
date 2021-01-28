package main

import (
	"fmt"
	"os"

	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/cmdutil"

	"github.com/bf2fc6cc711aee1a0c2a/cli/internal/build"

	"github.com/bf2fc6cc711aee1a0c2a/cli/internal/config"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/api/serviceapi"

	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/cmd/factory"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/cmd/root"
	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"
)

var (
	generateDocs = os.Getenv("GENERATE_DOCS") == "true"
)

func main() {
	buildVersion := build.Version
	cmdFactory := factory.New(build.Version)
	stderr := cmdFactory.IOStreams.ErrOut

	initConfig(cmdFactory)

	rootCmd := root.NewRootCommand(cmdFactory, buildVersion)

	rootCmd.InitDefaultHelpCmd()

	if generateDocs {
		generateDocumentation(rootCmd)
	}

	err := rootCmd.Execute()
	if err == nil {
		return
	}

	if e, ok := serviceapi.GetAPIError(err); ok {
		fmt.Fprintf(stderr, "Error: %v\n", e.GetReason())
		os.Exit(1)
	}

	if err = cmdutil.CheckSurveyError(err); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

/**
* Generates documentation files
 */
func generateDocumentation(rootCommand *cobra.Command) {
	fmt.Fprint(os.Stderr, "Generating docs.\n\n")
	filePrepender := func(filename string) string {
		return ""
	}

	rootCommand.DisableAutoGenTag = true

	linkHandler := func(s string) string { return s }

	err := doc.GenMarkdownTreeCustom(rootCommand, "./docs/commands", filePrepender, linkHandler)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func initConfig(f *factory.Factory) {
	cfgFile, err := f.Config.Load()
	if cfgFile != nil {
		return
	}
	if !os.IsNotExist(err) {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	cfgFile = &config.Config{}
	if err := f.Config.Save(cfgFile); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
