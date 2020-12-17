package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/bf2fc6cc711aee1a0c2a/cli/internal/config"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/api/managedservices"

	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/cmd/factory"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/cmd/root"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/version"
	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"
)

var (
	generateDocs = os.Getenv("GENERATE_DOCS") == "true"
)

func main() {
	cmdFactory := factory.New(version.CLI_VERSION)
	initConfig(cmdFactory.Config)

	rootCmd := root.NewRootCommand(cmdFactory, version.CLI_VERSION)

	rootCmd.SilenceErrors = true
	rootCmd.InitDefaultHelpCmd()

	if generateDocs {
		generateDocumentation(rootCmd)
	}

	err := rootCmd.Execute()
	if err == nil {
		return
	}

	// Attempt to unwrap the descriptive API error message
	var apiError managedservices.GenericOpenAPIError
	if ok := errors.As(err, &apiError); ok {
		errModel := apiError.Model()

		e, ok := errModel.(managedservices.Error)
		if ok {
			fmt.Fprintf(os.Stderr, "Error: %v\n", e.Reason)
			os.Exit(1)
		}
	}

	fmt.Fprintf(os.Stderr, "Error: %v\n", err)
	os.Exit(1)
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

func initConfig(cfg config.IConfig) {
	cfgFile, err := cfg.Load()
	if cfgFile != nil {
		return
	}
	if !os.IsNotExist(err) {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	cfgFile = &config.Config{}
	if err := cfg.Save(cfgFile); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
