// This file contains functions used to implement the '--debug' command line option.

package debug

import "github.com/spf13/pflag"

// AddFlag adds the verbose flag to the given set of command line flags.
func AddFlag(flags *pflag.FlagSet) {
	flags.BoolVarP(
		&enabled,
		"verbose",
		"v",
		false,
		"Enable verbose mode",
	)
}

// Enabled returns a boolean flag that indicates if the verbose mode is enabled
func Enabled() bool {
	return enabled
}

// enabled is a boolean flag that indicates that the verbose mode is enabled
var enabled bool
