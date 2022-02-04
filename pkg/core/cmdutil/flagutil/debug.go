// This file contains functions used to implement the '--debug' command line option.

package flagutil

import "github.com/spf13/pflag"

// VerboseFlag adds the verbose flag to the given set of command line flags.
func VerboseFlag(flags *pflag.FlagSet) {
	flags.BoolVarP(
		&enabled,
		"verbose",
		"v",
		false,
		"Enable verbose mode",
	)
}

// DebugEnabled returns a boolean flag that indicates if the verbose mode is enabled
func DebugEnabled() bool {
	return enabled
}

// enabled is a boolean flag that indicates that the verbose mode is enabled
var enabled bool
