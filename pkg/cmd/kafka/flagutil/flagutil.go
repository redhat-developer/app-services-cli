package flagutil

import (
	"github.com/redhat-developer/app-services-cli/pkg/core/cmdutil/flagutil"
	"github.com/redhat-developer/app-services-cli/pkg/core/localize"
	"github.com/redhat-developer/app-services-cli/pkg/shared/connection"
	"github.com/redhat-developer/app-services-cli/pkg/shared/factory"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
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

// RegisterNameFlagCompletionFunc adds dynamic completion for the --name flag
func RegisterNameFlagCompletionFunc(cmd *cobra.Command, f *factory.Factory) error {
	return cmd.RegisterFlagCompletionFunc("name", func(cmd *cobra.Command, _ []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		var validNames []string
		directive := cobra.ShellCompDirectiveNoSpace

		conn, err := f.Connection(connection.DefaultConfigSkipMasAuth)
		if err != nil {
			return validNames, directive
		}

		req := conn.API().KafkaMgmt().GetKafkas(f.Context)
		if toComplete != "" {
			searchQ := "name like " + toComplete + "%"
			req = req.Search(searchQ)
		}
		kafkas, httpRes, err := req.Execute()
		if err != nil {
			return validNames, directive
		}
		if httpRes != nil {
			defer httpRes.Body.Close()
		}

		items := kafkas.GetItems()
		for index := range items {
			validNames = append(validNames, items[index].GetName())
		}

		return validNames, directive
	})
}
