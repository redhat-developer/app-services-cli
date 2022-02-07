package registrycmdutil

import (
	"github.com/redhat-developer/app-services-cli/pkg/core/cmdutil/flagutil"
	"github.com/redhat-developer/app-services-cli/pkg/shared/factory"
	"github.com/spf13/cobra"
)

type flagSet struct {
	cmd     *cobra.Command
	factory *factory.Factory
	*flagutil.FlagSet
}

// NewFlagSet returns a new flag set with common Service Registry flags
func NewFlagSet(cmd *cobra.Command, f *factory.Factory) *flagSet {
	return &flagSet{
		cmd:     cmd,
		factory: f,
		FlagSet: flagutil.NewFlagSet(cmd, f.Localizer),
	}
}

// AddRegistryInstance adds a flag for setting the Service Registry instance ID
func (fs *flagSet) AddRegistryInstance(registryID *string) {
	flagName := "instance-id"

	fs.StringVar(
		registryID,
		flagName,
		"",
		fs.factory.Localizer.MustLocalize("artifact.common.instance.id"),
	)

}
