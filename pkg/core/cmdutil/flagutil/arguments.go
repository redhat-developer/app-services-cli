// This file contains functions that add common arguments to the command line
package flagutil

import (
	"github.com/redhat-developer/app-services-cli/pkg/cmd/debug"
	"github.com/spf13/pflag"
)

// AddDebugFlag adds the '--debug' flag to the given set of command line flags
func AddDebugFlag(fs *pflag.FlagSet) {
	debug.AddFlag(fs)
}
