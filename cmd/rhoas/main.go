package main

import (
	"fmt"
	"os"

	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/cmd/root"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/config"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/version"
	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"
)

var (
	generateDocs = os.Getenv("GENERATE_DOCS") == "true"
)

func main() {
	cobra.OnInitialize(initConfig)

	rootCmd := root.NewRootCommand(version.CLI_VERSION)
	rootCmd.InitDefaultHelpCmd()

	if generateDocs {
		generateDocumentation(rootCmd)
	}

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "error running command: %v\n", err)
	}
}

func initConfig() {
	cfgFile, err := config.Load()
	if cfgFile != nil {
		return
	}
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	cfgFile = &config.Config{}
	if err := config.Save(cfgFile); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	cfgPath, _ := config.Location()
	fmt.Fprintf(os.Stderr, "Saved config to %v\n", cfgPath)
}

/**
* Generates documentation files
 */
func generateDocumentation(rootCommand *cobra.Command) {
	fmt.Fprint(os.Stderr, "Generating docs.\n\n")
	filePrepender := func(filename string) string {
		return ""
	}

	linkHandler := func(s string) string { return s }

	err := doc.GenMarkdownTreeCustom(rootCommand, "./docs/commands", filePrepender, linkHandler)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
