package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/redhat-developer/app-services-cli/pkg/api/kas"
	"github.com/redhat-developer/app-services-cli/pkg/doc"
	"github.com/redhat-developer/app-services-cli/pkg/dump"
	"github.com/redhat-developer/app-services-cli/pkg/localize"
	"github.com/redhat-developer/app-services-cli/pkg/localize/goi18n"

	"github.com/redhat-developer/app-services-cli/pkg/cmdutil"

	"github.com/redhat-developer/app-services-cli/internal/build"

	"github.com/redhat-developer/app-services-cli/internal/config"

	"github.com/redhat-developer/app-services-cli/pkg/cmd/debug"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/factory"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/root"
	"github.com/spf13/cobra"
)

var generateDocs = os.Getenv("GENERATE_DOCS") == "true"

func main() {
	localizer, err := goi18n.New(nil)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	fmt.Println(localizer.MustLocalize("root.cmd.use"))
	os.Exit(0)

	buildVersion := build.Version
	cmdFactory := factory.New(build.Version, localizer)
	logger, err := cmdFactory.Logger()
	if err != nil {
		fmt.Println(cmdFactory.IOStreams.ErrOut, err)
		os.Exit(1)
	}

	err = initConfig(cmdFactory)
	if err != nil {
		logger.Errorf(localizer.MustLocalize("main.config.error", localize.NewEntry("Error", err)))
		os.Exit(1)
	}

	rootCmd := root.NewRootCommand(cmdFactory, buildVersion)

	rootCmd.InitDefaultHelpCmd()

	if generateDocs {
		generateDocumentation(rootCmd)
		os.Exit(0)
	}

	err = rootCmd.Execute()
	if debug.Enabled() {
		build.CheckForUpdate(context.Background(), logger, localizer)
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

func initConfig(f *factory.Factory) error {
	if !config.HasCustomLocation() {
		rhoasCfgDir, err := config.DefaultDir()
		if err != nil {
			return err
		}

		// create rhoas config directory
		if _, err = os.Stat(rhoasCfgDir); os.IsNotExist(err) {
			err = os.MkdirAll(rhoasCfgDir, 0o700)
			if err != nil {
				return err
			}
		}
	}

	cfgFile, err := f.Config.Load()

	if cfgFile != nil {
		return err
	}

	if !os.IsNotExist(err) {
		return err
	}

	cfgFile = &config.Config{}
	if err := f.Config.Save(cfgFile); err != nil {
		return err
	}
	return nil
}
