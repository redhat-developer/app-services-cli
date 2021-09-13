package main

import (
	"errors"
	"fmt"
	"github.com/redhat-developer/app-services-cli/pkg/api/kafka"
	kafkamgmtclient "github.com/redhat-developer/app-services-sdk-go/kafkamgmt/apiv1/client"
	"os"
	"strings"

	"github.com/redhat-developer/app-services-cli/pkg/icon"

	"github.com/redhat-developer/app-services-cli/pkg/doc"
	"github.com/redhat-developer/app-services-cli/pkg/localize"
	"github.com/redhat-developer/app-services-cli/pkg/localize/goi18n"

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

	buildVersion := build.Version
	cmdFactory := factory.New(localizer)

	err = initConfig(cmdFactory)
	if err != nil {
		cmdFactory.Logger.Errorf(localizer.MustLocalize("main.config.error", localize.NewEntry("Error", err)))
		os.Exit(1)
	}

	rootCmd := root.NewRootCommand(cmdFactory, buildVersion)

	rootCmd.InitDefaultHelpCmd()

	if generateDocs {
		generateDocumentation(rootCmd)
		os.Exit(0)
	}

	err = rootCmd.Execute()

	if err == nil {
		if debug.Enabled() {
			build.CheckForUpdate(cmdFactory.Context, cmdFactory.Logger, localizer)
		}
		return
	}
	cmdFactory.Logger.Error(rootError(err, localizer))
	build.CheckForUpdate(cmdFactory.Context, cmdFactory.Logger, localizer)
	os.Exit(1)
}

/**
* Generates documentation files
 */
func generateDocumentation(rootCommand *cobra.Command) {
	fmt.Fprintln(os.Stderr, "\n🛠️  Generating rhoas command-line reference documentation...")
	filePrepender := func(filename string) string {
		return ""
	}

	rootCommand.DisableAutoGenTag = true

	linkHandler := func(s string) string { return s }

	if err := doc.GenAsciidocTreeCustom(rootCommand, "./docs/commands", filePrepender, linkHandler); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Fprintln(os.Stderr, "\n✅ Command-line reference documentation has been generated successfully")
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

// rootError creates the root error which is printed to the console
// it wraps the error which has been returned from subcommands with a prefix
func rootError(err error, localizer localize.Localizer) error {
	var kafkaAPIError kafkamgmtclient.GenericOpenAPIError

	if errors.As(err, &kafkaAPIError) {
		err = kafka.TransformError(err)
	}

	errMessage := err.Error()

	prefix := icon.ErrorPrefix()
	if prefix == icon.CrossMark {
		errMessage = firstCharToUpper(errMessage)
	}
	return fmt.Errorf("%v %v. %v", icon.ErrorPrefix(), errMessage, localizer.MustLocalize("common.log.error.verboseModeHint"))
}

// Ensure that the first character in the string is capitalized
func firstCharToUpper(message string) string {
	return strings.ToUpper(message[:1]) + message[1:]
}
