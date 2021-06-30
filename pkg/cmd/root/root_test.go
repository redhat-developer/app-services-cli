package root

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/aerogear/charmil/validator/rules"
	"github.com/redhat-developer/app-services-cli/internal/build"
	"github.com/redhat-developer/app-services-cli/internal/config"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/factory"
	"github.com/redhat-developer/app-services-cli/pkg/localize/goi18n"
)

func Test_ExecuteCommand(t *testing.T) {

	localizer, err := goi18n.New(nil)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	buildVersion := build.Version
	cmdFactory := factory.New(build.Version, localizer)
	if err != nil {
		fmt.Println(cmdFactory.IOStreams.ErrOut, err)
		os.Exit(1)
	}

	initConfig(cmdFactory)

	cmd := NewRootCommand(cmdFactory, buildVersion)

	// Testing cobra commands with default recommended config
	var vali rules.RuleConfig
	validationErr := vali.ExecuteRules(cmd)

	if len(validationErr) != 0 {
		t.Errorf("validationErr was not empty, got length: %d; want %d", len(validationErr), 0)
	}

	for _, errs := range validationErr {
		if errs.Err != nil {
			t.Fatalf("%s: cmd %s: %s", errs.Rule, errs.Cmd.CommandPath(), errs.Name)
		}
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
		err = os.MkdirAll(rhoasCfgDir, 0700)
		if err != nil {
			return err
		}
	}
	if _, err = os.Stat(oldFilePath); err == nil {
		return os.Rename(oldFilePath, cfgPath)
	}
	return nil
}
