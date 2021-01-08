// This file contains functions that add common arguments to the command line

package arguments

import (
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/cmd/debug"
	"github.com/spf13/pflag"
)

// AddDebugFlag adds the '--debug' flag to the given set of command line flags
func AddDebugFlag(fs *pflag.FlagSet) {
	debug.AddFlag(fs)
}
