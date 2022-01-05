package flagutil

import (
	"github.com/redhat-developer/app-services-cli/pkg/core/cmdutil/flagutil"
	"github.com/redhat-developer/app-services-cli/pkg/core/localize"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

const (
	// FlagProvider is a flag representing an OCM provider ID
	FlagProvider = "provider"
	// FlagRegion is a flag representing an OCM region ID
	FlagRegion = "region"
)

type flagSet struct {
	flags     *pflag.FlagSet
	cmd       *cobra.Command
	localizer localize.Localizer
	*flagutil.FlagSet
}

// NewFlagSet returns a new flag set for creating common Kafka-command flags
func NewFlagSet(cmd *cobra.Command, localizer localize.Localizer) *flagSet {
	return &flagSet{
		cmd:       cmd,
		flags:     cmd.Flags(),
		localizer: localizer,
		FlagSet:   flagutil.NewFlagSet(cmd, localizer),
	}
}

// AddInstanceID adds a flag for setting the Kafka instance ID
func (fs *flagSet) AddInstanceID(instanceID *string) {
	flagName := "instance-id"

	fs.flags.StringVar(
		instanceID,
		flagName,
		"",
		flagutil.FlagDescription(fs.localizer, "kafka.common.flag.instanceID.description"),
	)
}
