// This file contains functions used to implement the '--insecure' command line option.

package insecure

import (
	"github.com/bf2fc6cc711aee1a0c2a/cli/internal/localizer"
	"github.com/spf13/pflag"
)

// AddFlag adds the debug flag to the given set of command line flags.
func AddFlag(flags *pflag.FlagSet) {
	flags.BoolVar(
		&insecure,
		"insecure",
		false,
		localizer.MustLocalizeFromID("login.flag.insecure"),
	)
}

// Insecure returns a boolean flag that indicates if the insecure mode is enabled
func Insecure() *bool {
	return &insecure
}

// insecure is a boolean flag that indicates that the debug mode is insecure
var insecure bool
