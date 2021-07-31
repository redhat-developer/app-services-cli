package root

import (
	"fmt"
	"os"
	"testing"

	"github.com/aerogear/charmil/validator"
	"github.com/aerogear/charmil/validator/rules"
	"github.com/redhat-developer/app-services-cli/internal/build"
	"github.com/redhat-developer/app-services-cli/internal/config"
	"github.com/redhat-developer/app-services-cli/internal/mockutil"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/factory"
	"github.com/redhat-developer/app-services-cli/pkg/localize/goi18n"
)

func Test_ValidateCommandsUsingCharmilValidator(t *testing.T) {
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

	mockutil.NewConfigMock(&config.Config{})
	cmd := NewRootCommand(cmdFactory, buildVersion)

	// Testing cobra commands with default recommended config
	vali := rules.ValidatorConfig{
		ValidatorOptions: rules.ValidatorOptions{
			SkipCommands: map[string]bool{
				"rhoas kafka*":           true,
				"rhoas completion*":      true,
				"rhoas cluster":          true,
				"rhoas logout":           true,
				"rhoas service-account*": true,
			},
		},
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
