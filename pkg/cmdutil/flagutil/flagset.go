// Package flagutil is a helper package for processing and interactive command line flags
package flagutil

import (
	"github.com/redhat-developer/app-services-cli/pkg/dump"
	"github.com/redhat-developer/app-services-cli/pkg/localize"
	"github.com/redhat-developer/app-services-cli/pkg/serviceaccount/credentials"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

var (
	ValidOutputFormats       = []string{dump.JSONFormat, dump.YAMLFormat, dump.YMLFormat}
	CredentialsOutputFormats = []string{credentials.EnvFormat, credentials.JSONFormat, credentials.PropertiesFormat}
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

// AddOutput adds an output flag to the command
func (fs *FlagSet) AddOutput(output *string) {
	flagName := "output"

	fs.StringVarP(
		output,
		flagName,
		"o",
		dump.EmptyFormat,
		FlagDescription(fs.localizer, "flag.common.output.description", ValidOutputFormats...),
	)

	_ = fs.cmd.RegisterFlagCompletionFunc(flagName, func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return ValidOutputFormats, cobra.ShellCompDirectiveNoSpace
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
