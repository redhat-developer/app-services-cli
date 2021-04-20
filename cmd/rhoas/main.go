package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/markbates/pkger"
	"github.com/redhat-developer/app-services-cli/pkg/api/kas"
	"github.com/redhat-developer/app-services-cli/pkg/doc"
	"github.com/redhat-developer/app-services-cli/pkg/dump"

	"github.com/redhat-developer/app-services-cli/pkg/cmdutil"

	"github.com/redhat-developer/app-services-cli/internal/build"
	"github.com/redhat-developer/app-services-cli/internal/localizer"

	"github.com/redhat-developer/app-services-cli/internal/config"

	"github.com/redhat-developer/app-services-cli/pkg/cmd/debug"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/factory"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/root"
	"github.com/spf13/cobra"
)

var (
	generateDocs = os.Getenv("GENERATE_DOCS") == "true"
)

// load all locale files
func loadStaticFiles() error {
	err := localizer.IncludeAssetsAndLoadMessageFiles()
	if err != nil {
		return err
	}

	return pkger.Walk("/static", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		return nil
	})
}

func main() {
	err := loadStaticFiles()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	buildVersion := build.Version
	cmdFactory := factory.New(build.Version)
	logger, err := cmdFactory.Logger()
	if err != nil {
		fmt.Println(cmdFactory.IOStreams.ErrOut, err)
		os.Exit(1)
	}

	initConfig(cmdFactory)

	rootCmd := root.NewRootCommand(cmdFactory, buildVersion)

	rootCmd.InitDefaultHelpCmd()

	if generateDocs {
		generateDocumentation(rootCmd)
		os.Exit(0)
	}

	err = rootCmd.Execute()
	if debug.Enabled() {
		build.CheckForUpdate(context.Background(), logger)
	}
	if err == nil {
		return
	}

	if e, ok := kas.GetAPIError(err); ok {
		logger.Error("Error:", e.GetReason())
		if debug.Enabled() {
			errJSON, _ := json.Marshal(e)
			_ = dump.JSON(cmdFactory.IOStreams.ErrOut, errJSON)
		}
		os.Exit(1)
	}

	if err = cmdutil.CheckSurveyError(err); err != nil {
		logger.Error("Error:", err)
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

	err := doc.GenAsciidocTreeCustom(rootCommand, "./docs/commands", filePrepender, linkHandler)
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
		fmt.Fprintln(f.IOStreams.ErrOut, err)
		os.Exit(1)
	}

	cfgFile = &config.Config{}
	if err := f.Config.Save(cfgFile); err != nil {
		fmt.Fprintln(f.IOStreams.ErrOut, err)
		os.Exit(1)
	}
}
