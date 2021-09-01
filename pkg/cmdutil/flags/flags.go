// flags package is a helper package for processing and interactive command line flags
package flags

import (
	"github.com/redhat-developer/app-services-cli/pkg/dump"
	"github.com/redhat-developer/app-services-cli/pkg/kafka/consumergroup"
	"github.com/spf13/cobra"
)

var (
	ValidOutputFormats       = []string{dump.JSONFormat, dump.YAMLFormat, dump.YMLFormat}
	CredentialsOutputFormats = []string{"env", "json", "properties"}
	ValidOffsets             = []string{consumergroup.AbsoluteOffset, consumergroup.EarliestOffset, consumergroup.TimestampOffset, consumergroup.LatestOffset}
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
