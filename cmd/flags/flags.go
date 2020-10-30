// flags package is a helper package for processing and interactive command line flags
package flags

import (
	"fmt"

	"github.com/golang/glog"
	"github.com/spf13/pflag"
)

// MustGetDefinedString attempts to get a non-empty string flag from the provided flag set or panic
func MustGetDefinedString(flagName string, flags *pflag.FlagSet) string {
	flagVal := MustGetString(flagName, flags)
	if flagVal == "" {
		glog.Fatal(undefinedValueMessage(flagName))
	}
	return flagVal
}

// MustGetString attempts to get a string flag from the provided flag set or panic
func MustGetString(flagName string, flags *pflag.FlagSet) string {
	flagVal, err := flags.GetString(flagName)
	if err != nil {
		glog.Fatalf(notFoundMessage(flagName, err))
	}
	return flagVal
}

// MustGetBool attempts to get a boolean flag from the provided flag set or panic
func MustGetBool(flagName string, flags *pflag.FlagSet) bool {
	flagVal, err := flags.GetBool(flagName)
	if err != nil {
		glog.Fatalf(notFoundMessage(flagName, err))
	}
	return flagVal
}

func undefinedValueMessage(flagName string) string {
	return fmt.Sprintf("flag %s has undefined value", flagName)
}

func notFoundMessage(flagName string, err error) string {
	return fmt.Sprintf("could not get flag %s from flag set: %s", flagName, err.Error())
}
