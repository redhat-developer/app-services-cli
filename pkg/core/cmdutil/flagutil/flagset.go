// Package flagutil is a helper package for processing and interactive command line flags
package flagutil

import (
	"github.com/redhat-developer/app-services-cli/internal/build"
	"github.com/redhat-developer/app-services-cli/pkg/core/cmdutil"
	"github.com/redhat-developer/app-services-cli/pkg/core/ioutil/dump"
	"github.com/redhat-developer/app-services-cli/pkg/core/localize"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

var (
	ValidOutputFormats     = []string{dump.JSONFormat, dump.YAMLFormat, dump.YMLFormat, dump.EmptyFormat, dump.TableFormat}
	ValidListOutputFormats = append(ValidOutputFormats, dump.TableFormat)
)

type FlagSet struct {
	*pflag.FlagSet
	cmd       *cobra.Command
	localizer localize.Localizer
}

// NewFlagSet returns a new flag set with some common flags
func NewFlagSet(cmd *cobra.Command, localizer localize.Localizer) *FlagSet {
	return &FlagSet{
		FlagSet:   cmd.Flags(),
		cmd:       cmd,
		localizer: localizer,
	}
}

// AddOutput executes AddOutputFormatted with required inputs
func (fs *FlagSet) AddOutput(output *string) {
	fs.AddOutputFormatted(output, false, nil)
}

// AddOutputFormatted adds an output flag to the command
func (fs *FlagSet) AddOutputFormatted(output *string, isTable bool, format *string) {
	flagName := "output"

	selectedFormat := dump.JSONFormat

	if format != nil && *format != "" {
		selectedFormat = *format

	}

	suggestions := ValidOutputFormats

	if isTable {
		suggestions = ValidListOutputFormats
	}

	fs.StringVarP(
		output,
		flagName,
		"o",
		selectedFormat,
		FlagDescription(fs.localizer, "flag.common.output.description", suggestions...),
	)

	_ = fs.cmd.RegisterFlagCompletionFunc(flagName, func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return suggestions, cobra.ShellCompDirectiveNoSpace
	})
}

// AddYes adds a "yes" flag to the command
func (fs *FlagSet) AddYes(yes *bool) {
	flagName := "yes"

	fs.BoolVarP(
		yes,
		flagName,
		"y",
		false,
		FlagDescription(fs.localizer, "flag.common.yes.description"),
	)

	_ = fs.cmd.RegisterFlagCompletionFunc(flagName, func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return ValidOutputFormats, cobra.ShellCompDirectiveNoSpace
	})
}

// AddPage adds a "page" flag to the command
func (fs *FlagSet) AddPage(page *int32) {
	flagName := "page"

	fs.Int32Var(
		page,
		flagName,
		cmdutil.ConvertPageValueToInt32(build.DefaultPageNumber),
		FlagDescription(fs.localizer, "kafka.common.flag.page.description"),
	)

}

// AddSize adds a "size" flag to the command
func (fs *FlagSet) AddSize(page *int32) {
	flagName := "size"

	fs.Int32Var(
		page,
		flagName,
		cmdutil.ConvertPageValueToInt32(build.DefaultPageSize),
		FlagDescription(fs.localizer, "kafka.common.flag.size.description"),
	)
}

// AddBypassTermsCheck adds a flag to allow bypassing
// of the terms check before creating an instance
func (fs *FlagSet) AddBypassTermsCheck(bypass *bool) {
	flagName := "bypass-terms-check"

	fs.BoolVar(
		bypass,
		flagName,
		false,
		"Bypass the terms and conditions pre-check and other AMS checks",
	)

	_ = fs.MarkHidden(flagName)
}

// WithFlagOptions returns additional functions to custom the default flag settings
func WithFlagOptions(cmd *cobra.Command, flagName string) *FlagOptions {
	return &FlagOptions{
		Required: func() error {
			return cmd.MarkFlagRequired(flagName)
		},
	}
}

// FlagOptions defines additional flag options
type FlagOptions struct {
	Required func() error
}

// ValidateOutput checks if value v is a valid value for --output
func ValidateOutput(v string) error {
	isValid := IsValidInput(v, ValidOutputFormats...)

	if isValid {
		return nil
	}

	return InvalidValueError("output", v, ValidOutputFormats...)
}

// ValidateOutputForTable checks if value v is a valid value for --output adding table
func ValidateOutputForTable(v string) error {
	isValid := IsValidInput(v, ValidListOutputFormats...)

	if isValid {
		return nil
	}

	return InvalidValueError("output", v, ValidListOutputFormats...)
}
