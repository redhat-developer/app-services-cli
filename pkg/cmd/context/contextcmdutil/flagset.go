package contextcmdutil

import (
	"github.com/redhat-developer/app-services-cli/pkg/core/cmdutil/flagutil"
	"github.com/redhat-developer/app-services-cli/pkg/shared/factory"
	"github.com/spf13/cobra"
)

type FlagSet struct {
	cmd     *cobra.Command
	factory *factory.Factory
	*flagutil.FlagSet
}

// NewFlagSet returns a new flag set with common context flags
func NewFlagSet(cmd *cobra.Command, f *factory.Factory) *FlagSet {
	return &FlagSet{
		cmd:     cmd,
		factory: f,
		FlagSet: flagutil.NewFlagSet(cmd, f.Localizer),
	}
}

// AddContextName adds a flag for setting the name of the context
func (fs *FlagSet) AddContextName(name *string) *flagutil.FlagOptions {
	flagName := "name"

	fs.StringVar(
		name,
		flagName,
		"",
		fs.factory.Localizer.MustLocalize("context.common.flag.name"),
	)

	return flagutil.WithFlagOptions(fs.cmd, flagName)

}
