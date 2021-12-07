package flagutil

import (
	"errors"

	"github.com/spf13/cobra"
)

const (
	TRIBOOL_TRUE    = "true"
	TRIBOOL_FALSE   = "false"
	TRIBOOL_DEFAULT = ""
)

// ValidTribools is an array of valid tribool string values
var ValidTribools = []string{
	TRIBOOL_TRUE,
	TRIBOOL_FALSE,
	TRIBOOL_DEFAULT,
}

// IsTriboolValid validates if a string corresponds to a valid tribool value
func IsTriboolValid(val string) error {
	switch Tribool(val) {
	case TRIBOOL_FALSE, TRIBOOL_TRUE, TRIBOOL_DEFAULT:
		return nil
	}

	return errors.New("invalid tribool")
}

// Tribool is a tri-state boolean where extra state corresponds to ""
type Tribool string

// Set accepts the CLI input as string and assigns it to the tribool variable
func (s *Tribool) Set(val string) error {
	err := IsTriboolValid(val)
	if err != nil {
		return err
	}
	*s = Tribool(val)
	return nil
}

// Type returns the type of flag
func (s *Tribool) Type() string {
	return "Tribool"
}

// String returns the tribool as a string
func (s *Tribool) String() string { return string(*s) }

// TriBoolVar defines a tribool flag with specified name, default value, and usage string.
// The argument p points to a tribool variable in which to store the value of the flag.
func (fs *FlagSet) TriBoolVar(p *Tribool, name string, value Tribool, usage string) {
	fs.VarP(newTriBoolValue(value, p), name, "", usage)

	_ = fs.cmd.RegisterFlagCompletionFunc(name, func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return ValidTribools, cobra.ShellCompDirectiveNoSpace
	})
}

func newTriBoolValue(val Tribool, p *Tribool) *Tribool {
	*p = Tribool(val)
	return (*Tribool)(p)
}
