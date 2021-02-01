package flag

import (
	"fmt"

	"github.com/spf13/pflag"
)

func AddLimit(flags *pflag.FlagSet, p *int, defaultVal int) {
	flags.IntVar(
		p,
		"limit",
		defaultVal,
		"Set a limit on the number of results to show",
	)
}

func AddPage(flags *pflag.FlagSet, p *int, defaultVal int) {
	flags.IntVar(
		p,
		"page",
		defaultVal,
		"Set the current page of the results",
	)
}

func AddOutput(flags *pflag.FlagSet, p *string, defaultVal string, options []string) {
	flags.StringVarP(
		p,
		"output",
		"o",
		defaultVal,
		fmt.Sprintf("Output format of the results. Choose from %q", options),
	)
}
