package connectorcmdutil

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

// NewFlagSet returns a new flag set with common Service Registry flags
func NewFlagSet(cmd *cobra.Command, f *factory.Factory) *FlagSet {
	return &FlagSet{
		cmd:     cmd,
		factory: f,
		FlagSet: flagutil.NewFlagSet(cmd, f.Localizer),
	}
}

// AddConnectorID adds a flag for specifying the connector ID
func (fs *FlagSet) AddConnectorID(ruleType *string) *flagutil.FlagOptions {
	flagName := "id"

	fs.StringVar(
		ruleType,
		flagName,
		"",
		fs.factory.Localizer.MustLocalize("connector.common.flag.id.description"),
	)

	return flagutil.WithFlagOptions(fs.cmd, flagName)

}
