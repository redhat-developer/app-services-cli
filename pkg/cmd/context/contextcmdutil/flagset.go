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

	svcContext, _ := fs.factory.ServiceContext.Load()

	svcContextsMap := svcContext.Contexts
	flagName := "name"

	fs.StringVar(
		name,
		flagName,
		"",
		fs.factory.Localizer.MustLocalize("context.common.flag.name"),
	)

	contextNames := make([]string, 0, len(svcContextsMap))
	for k := range svcContextsMap {
		contextNames = append(contextNames, k)
	}

	_ = fs.cmd.RegisterFlagCompletionFunc(flagName, func(cmd *cobra.Command, _ []string, _ string) ([]string, cobra.ShellCompDirective) {
		return contextNames, cobra.ShellCompDirectiveNoSpace
	})

	return flagutil.WithFlagOptions(fs.cmd, flagName)

}
