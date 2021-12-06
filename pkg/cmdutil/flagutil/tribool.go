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

var validTribools = []string{
	TRIBOOL_TRUE,
	TRIBOOL_FALSE,
	TRIBOOL_DEFAULT,
}

func isTriboolValid(val string) error {
	switch Tribool(val) {
	case TRIBOOL_FALSE, TRIBOOL_TRUE:
		return nil
	}

	return errors.New("invalid tribool")
}

type Tribool string

func (s *Tribool) Set(val string) error {
	err := isTriboolValid(val)
	if err != nil {
		return err
	}
	*s = Tribool(val)
	return nil
}

func (s *Tribool) Type() string {
	return "Tribool"
}

func (s *Tribool) String() string { return string(*s) }

// TriBoolVar defines a tribool flag with specified name, default value, and usage string.
// The argument p points to a tribool variable in which to store the value of the flag.
func (fs *FlagSet) TriBoolVar(p *Tribool, name string, value Tribool, usage string) {
	fs.VarP(newTriBoolValue(value, p), name, "", usage)

	_ = fs.cmd.RegisterFlagCompletionFunc(name, func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return validTribools, cobra.ShellCompDirectiveNoSpace
	})
}

func newTriBoolValue(val Tribool, p *Tribool) *Tribool {
	*p = Tribool(val)
	return (*Tribool)(p)
}
