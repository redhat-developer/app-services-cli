package rulecmdutil

import (
	"github.com/redhat-developer/app-services-cli/pkg/cmd/registry/registrycmdutil"
	"github.com/redhat-developer/app-services-cli/pkg/core/cmdutil/flagutil"
	"github.com/redhat-developer/app-services-cli/pkg/shared/factory"
	"github.com/spf13/cobra"
)

type flagSet struct {
	cmd     *cobra.Command
	factory *factory.Factory
	*registrycmdutil.FlagSet
}

// NewFlagSet returns a new flag set with common Service Registry rule flags
func NewFlagSet(cmd *cobra.Command, f *factory.Factory) *flagSet {
	return &flagSet{
		cmd:     cmd,
		factory: f,
		FlagSet: registrycmdutil.NewFlagSet(cmd, f),
	}
}

// AddArtifactID adds a flag for setting the artifact ID
func (fs *flagSet) AddArtifactID(artifactID *string) {
	flagName := "artifact-id"

	fs.StringVar(
		artifactID,
		flagName,
		"",
		fs.factory.Localizer.MustLocalize("artifact.common.id"),
	)

}

// AddGroup adds a flag for setting the artifact group
func (fs *flagSet) AddGroup(artifactID *string) {
	flagName := "group"

	fs.StringVarP(
		artifactID,
		flagName,
		"g",
		registrycmdutil.DefaultArtifactGroup,
		fs.factory.Localizer.MustLocalize("artifact.common.group"),
	)

}

// AddRuleType adds a flag for setting the rule type
func (fs *flagSet) AddRuleType(ruleType *string) *flagutil.FlagOptions {
	flagName := "rule-type"

	ruleTypeMap := GetRuleTypeMap()

	ruleTypes := make([]string, 0, len(ruleTypeMap))
	for i := range ruleTypeMap {
		ruleTypes = append(ruleTypes, i)
	}

	fs.StringVar(
		ruleType,
		flagName,
		"",
		fs.factory.Localizer.MustLocalize("registry.rule.common.flag.ruleType"),
	)
	_ = fs.cmd.RegisterFlagCompletionFunc(flagName, func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return ruleTypes, cobra.ShellCompDirectiveNoSpace
	})

	return flagutil.WithFlagOptions(fs.cmd, flagName)

}

// AddConfig adds a flag for setting the configuration value for a rule
func (fs *flagSet) AddConfig(config *string) *flagutil.FlagOptions {
	flagName := "config"

	validityConfigMap := GetConfigMap()
	compatibilityConfigMap := GetConfigMap()

	configs := make([]string, 0, len(validityConfigMap)+len(compatibilityConfigMap))

	var i string
	for i = range validityConfigMap {
		configs = append(configs, i)
	}
	for i = range compatibilityConfigMap {
		configs = append(configs, i)
	}

	fs.StringVar(
		config,
		flagName,
		"",
		fs.factory.Localizer.MustLocalize("registry.rule.common.flag.config"),
	)

	_ = fs.cmd.RegisterFlagCompletionFunc(flagName, func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return configs, cobra.ShellCompDirectiveNoSpace
	})

	return flagutil.WithFlagOptions(fs.cmd, flagName)

}
