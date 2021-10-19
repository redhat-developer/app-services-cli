// Package flags is a helper package for processing and interactive command line flags
package flags

import (
	"fmt"
	"sort"

	"github.com/redhat-developer/app-services-cli/pkg/dump"
	"github.com/redhat-developer/app-services-cli/pkg/localize"
	"github.com/redhat-developer/app-services-cli/pkg/serviceaccount/credentials"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

var (
	ValidOutputFormats       = []string{dump.JSONFormat, dump.YAMLFormat, dump.YMLFormat}
	CredentialsOutputFormats = []string{credentials.EnvFormat, credentials.JSONFormat, credentials.PropertiesFormat}
)

// IsValidInput checks if the input value is in the range of valid values
func IsValidInput(input string, validValues ...string) bool {
	for _, b := range validValues {
		if input == b {
			return true
		}
	}

	return false
}

// FlagDescription creates a flag description and adds a list of valid options (if any)
func FlagDescription(localizer localize.Localizer, messageID string, validOptions ...string) string {
	// ensure consistent order
	sort.Strings(validOptions)

	description := localizer.MustLocalize(messageID)
	if description[len(description)-1:] != "." {
		description += "."
	}

	var chooseFrom string
	if len(validOptions) > 0 {
		chooseFrom = localizer.MustLocalize("flag.common.chooseFrom")

		for i, val := range validOptions {
			chooseFrom += fmt.Sprintf("\"%v\"", val)
			if i < len(validOptions)-1 {
				chooseFrom += ", "
			}
		}
	}

	return fmt.Sprintf("%v %v", description, chooseFrom)
}

type FlagSet struct {
	flags     *pflag.FlagSet
	cmd       *cobra.Command
	localizer localize.Localizer
}

func NewFlagSet(cmd *cobra.Command, localizer localize.Localizer) *FlagSet {
	return &FlagSet{
		flags:     cmd.Flags(),
		cmd:       cmd,
		localizer: localizer,
	}
}

// AddOutput adds an output flag to the command
func (fs *FlagSet) AddOutput(output *string) {
	flagName := "output"

	fs.flags.StringVarP(
		output,
		flagName,
		"o",
		dump.EmptyFormat,
		FlagDescription(fs.localizer, "flag.common.output.description", ValidOutputFormats...),
	)

	_ = fs.cmd.RegisterFlagCompletionFunc(flagName, func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return ValidOutputFormats, cobra.ShellCompDirectiveNoSpace
	})
}

// AddYes adds a "yes" flag to the command
func (fs *FlagSet) AddYes(yes *bool) {
	flagName := "yes"

	fs.flags.BoolVarP(
		yes,
		flagName,
		"y",
		false,
		FlagDescription(fs.localizer, "flag.common.yes.description"),
	)

	_ = fs.cmd.RegisterFlagCompletionFunc(flagName, func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return ValidOutputFormats, cobra.ShellCompDirectiveNoSpace
	})
}
