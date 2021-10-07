// flags package is a helper package for processing and interactive command line flags
package flags

import (
	"fmt"

	"github.com/redhat-developer/app-services-cli/pkg/dump"
	"github.com/redhat-developer/app-services-cli/pkg/localize"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

var (
	ValidOutputFormats       = []string{dump.JSONFormat, dump.YAMLFormat, dump.YMLFormat}
	CredentialsOutputFormats = []string{"env", "json", "properties"}
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

// EnableStaticFlagCompletion enables autocompletion for flags with predefined valid values
func EnableStaticFlagCompletion(cmd *cobra.Command, flagName string, validValues []string) {
	_ = cmd.RegisterFlagCompletionFunc(flagName, func(cmd *cobra.Command, _ []string, _ string) ([]string, cobra.ShellCompDirective) {
		return validValues, cobra.ShellCompDirectiveNoSpace
	})
}

// // EnableOutputFlagCompletion enables autocompletion for output flag
func EnableOutputFlagCompletion(cmd *cobra.Command) {
	_ = cmd.RegisterFlagCompletionFunc("output", func(cmd *cobra.Command, _ []string, _ string) ([]string, cobra.ShellCompDirective) {
		return ValidOutputFormats, cobra.ShellCompDirectiveNoSpace
	})
}

func FlagDescription(localizer localize.Localizer, messageID string, validOptions ...string) string {
	description := localizer.MustLocalize(messageID)
	if description[len(description)-1:] != "." {
		description += "."
	}

	chooseFrom := localizer.MustLocalize("flag.common.chooseFrom")

	var options string
	for i, val := range validOptions {
		options += fmt.Sprintf("\"%v\"", val)
		if i < len(validOptions)-1 {
			options += ", "
		}
	}

	return fmt.Sprintf("%v %v%v", description, chooseFrom, options)
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
