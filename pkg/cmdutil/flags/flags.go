// flags package is a helper package for processing and interactive command line flags
package flags

import (
	"fmt"
	"os"

	"github.com/spf13/pflag"
)

var (
	AllowedListFormats       = []string{"plain", "json", "yml", "yaml"}
	CredentialsOutputFormats = []string{"env", "json", "properties", "kube"}
)

// MustGetDefinedString attempts to get a non-empty string flag from the provided flag set or panic
func MustGetDefinedString(flagName string, flags *pflag.FlagSet) string {
	flagVal := MustGetString(flagName, flags)
	if flagVal == "" {
		fmt.Fprintln(os.Stderr, undefinedValueMessage(flagName))
		os.Exit(1)
	}
	return flagVal
}

// MustGetString attempts to get a string flag from the provided flag set or panic
func MustGetString(flagName string, flags *pflag.FlagSet) string {
	flagVal, err := flags.GetString(flagName)
	if err != nil {
		fmt.Fprintln(os.Stderr, notFoundMessage(flagName, err))
		os.Exit(1)
	}
	return flagVal
}

func GetString(flagName string, flags *pflag.FlagSet) string {
	flagVal, err := flags.GetString(flagName)
	if err != nil {
		return ""
	}
	return flagVal
}

func GetBool(flagName string, flags *pflag.FlagSet) string {
	flagVal, err := flags.GetString(flagName)
	if err != nil {
		return ""
	}
	return flagVal
}

// MustGetBool attempts to get a boolean flag from the provided flag set or panic
func MustGetBool(flagName string, flags *pflag.FlagSet) bool {
	flagVal, err := flags.GetBool(flagName)
	if err != nil {
		fmt.Fprintln(os.Stderr, notFoundMessage(flagName, err))
		os.Exit(1)
	}
	return flagVal
}

// IsValidInput checks if the input value is in the range of valid values
func IsValidInput(input string, validValues ...string) bool {
	for _, b := range validValues {
		if input == b {
			return true
		}
	}

	return false
}

func undefinedValueMessage(flagName string) string {
	return fmt.Sprintf("flag %s has undefined value", flagName)
}

func notFoundMessage(flagName string, err error) string {
	return fmt.Sprintf("could not get flag %s from flag set: %s", flagName, err.Error())
}
