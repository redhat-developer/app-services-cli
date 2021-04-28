package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/redhat-developer/app-services-cli/pkg/api/kas"
	"github.com/redhat-developer/app-services-cli/pkg/doc"
	"github.com/redhat-developer/app-services-cli/pkg/dump"
	"github.com/redhat-developer/app-services-cli/pkg/locales"

	"github.com/redhat-developer/app-services-cli/pkg/cmdutil"

	"github.com/redhat-developer/app-services-cli/internal/build"

	"github.com/redhat-developer/app-services-cli/internal/config"

	"github.com/redhat-developer/app-services-cli/pkg/cmd/debug"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/factory"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/root"
	"github.com/spf13/cobra"
)

var (
	generateDocs = os.Getenv("GENERATE_DOCS") == "true"
)

func main() {
	localizer, err := locales.New(nil)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	buildVersion := build.Version
	cmdFactory := factory.New(build.Version, localizer)
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
	// check if the config file is located in the old default location
	// if so, move it to the new location
	err := moveConfigFile(f.Config)
	if err != nil {
		fmt.Fprintf(f.IOStreams.ErrOut, "Error migrating config file to new location: %v", err)
	}

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

// check if the config file is located in the old default location
// if so, move it to the new location
func moveConfigFile(cfg config.IConfig) error {
	cfgPath, err := cfg.Location()
	if err != nil {
		return err
	}
	rhoasCfgDir, err := config.DefaultDir()
	if err != nil {
		return err
	}
	userCfgDir, err := os.UserConfigDir()
	if err != nil {
		return err
	}
	oldFilePath := filepath.Join(userCfgDir, ".rhoascli.json")
	if os.Getenv("RHOASCONFIG") == oldFilePath {
		return nil
	}
	// create rhoas config directory
	if _, err = os.Stat(rhoasCfgDir); os.IsNotExist(err) {
		err = os.Mkdir(rhoasCfgDir, 0700)
		if err != nil {
			return err
		}
	}
	if _, err = os.Stat(oldFilePath); err == nil {
		return os.Rename(oldFilePath, cfgPath)
	}
	return nil
}
