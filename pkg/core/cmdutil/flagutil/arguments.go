// This file contains functions that add common arguments to the command line
package flagutil

import (
	"github.com/spf13/pflag"
)

// AddDebugFlag adds the '--debug' flag to the given set of command line flags
func AddDebugFlag(fs *pflag.FlagSet) {
	VerboseFlag(fs)
}
