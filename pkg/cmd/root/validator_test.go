package root

import (
	"fmt"
	"os"
	"testing"

	"github.com/redhat-developer/app-services-cli/pkg/core/config"
	"github.com/redhat-developer/app-services-cli/pkg/core/localize/goi18n"
	"github.com/redhat-developer/app-services-cli/pkg/shared/factory/defaultfactory"

	"github.com/aerogear/charmil/validator"
	"github.com/aerogear/charmil/validator/rules"
	"github.com/redhat-developer/app-services-cli/internal/build"
	"github.com/redhat-developer/app-services-cli/internal/mockutil"
)

func Test_ValidateCommandsUsingCharmilValidator(t *testing.T) {
	localizer, err := goi18n.New(nil)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	buildVersion := build.Version
	cmdFactory := defaultfactory.New(build.BinaryName, localizer)
	if err != nil {
		fmt.Println(cmdFactory.IOStreams.ErrOut, err)
		os.Exit(1)
	}

	mockutil.NewConfigMock(&config.Config{})
	cmd := NewRootCommand(cmdFactory, buildVersion)

	// Testing cobra commands with default recommended config
	vali := rules.ValidatorConfig{
		ValidatorOptions: rules.ValidatorOptions{},
		ValidatorRules: rules.ValidatorRules{
			Length: rules.Length{
				Limits: map[string]rules.Limit{
					"Short":   {Min: 5},
					"Example": {Min: 10},
					"Long":    {Min: 10},
				},
			},
			Punctuation: rules.Punctuation{
				RuleOptions: validator.RuleOptions{
					Verbose: true,
				},
			},
		},
	}
	validationErr := rules.ExecuteRules(cmd, &vali)

	if len(validationErr) != 0 {
		t.Errorf("validationErr was not empty, got length: %d; want %d", len(validationErr), 0)
	}

	for _, errs := range validationErr {
		if errs.Err != nil {
			t.Logf("%s: cmd %s: %s", errs.Rule, errs.Cmd.CommandPath(), errs.Name)
		}
	}
}
