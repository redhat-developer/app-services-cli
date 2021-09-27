package main

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/redhat-developer/app-services-cli/pkg/icon"

	"github.com/redhat-developer/app-services-cli/pkg/localize"
	"github.com/redhat-developer/app-services-cli/pkg/localize/goi18n"

	"github.com/redhat-developer/app-services-cli/internal/build"

	"github.com/redhat-developer/app-services-cli/internal/config"

	"github.com/redhat-developer/app-services-cli/pkg/cmd/debug"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/factory"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/root"
)

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

	err = rootCmd.Execute()

	if err == nil {
		if debug.Enabled() {
			build.CheckForUpdate(cmdFactory.Context, cmdFactory.Logger, localizer)
		}
		return
	}
	cmdFactory.Logger.Errorf("%v\n", rootError(err, localizer))
	build.CheckForUpdate(context.Background(), cmdFactory.Logger, localizer)
	os.Exit(1)
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
	prefix := icon.ErrorPrefix()
	errMessage := err.Error()
	if prefix == icon.CrossMark {
		errMessage = firstCharToUpper(errMessage)
	}
	return fmt.Errorf("%v %v. %v", icon.ErrorPrefix(), errMessage, localizer.MustLocalize("common.log.error.verboseModeHint"))
}

// Ensure that the first character in the string is capitalized
func firstCharToUpper(message string) string {
	return strings.ToUpper(message[:1]) + message[1:]
}
