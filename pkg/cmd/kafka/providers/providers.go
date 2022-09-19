package providers

import (
	"github.com/redhat-developer/app-services-cli/pkg/cmd/kafka/create"
	kafkaFlagutil "github.com/redhat-developer/app-services-cli/pkg/cmd/kafka/flagutil"

	"github.com/redhat-developer/app-services-cli/pkg/core/cmdutil/flagutil"
	"github.com/redhat-developer/app-services-cli/pkg/core/ioutil/dump"
	"github.com/redhat-developer/app-services-cli/pkg/core/servicecontext"
	"github.com/redhat-developer/app-services-cli/pkg/shared/factory"

	"github.com/spf13/cobra"
)

// row is the details of a Kafka instance needed to print to a table
type providerRow struct {
	CloudProvider string `json:"cloud_provider" header:"Cloud Provider"`
	Region        string `json:"region" header:"Region"`
}

type options struct {
	outputFormat string

	f              *factory.Factory
	ServiceContext servicecontext.IContext
}

// NewListCommand creates a new command for listing kafkas.
func NewProviderCommand(f *factory.Factory) *cobra.Command {
	opts := &options{
		ServiceContext: f.ServiceContext,
	}

	cmd := &cobra.Command{
		Use:     "providers",
		Short:   opts.f.Localizer.MustLocalize("kafka.provider.cmd.shortDescription"),
		Long:    opts.f.Localizer.MustLocalize("kafka.provider.cmd.longDescription"),
		Example: opts.f.Localizer.MustLocalize("kafka.provider.cmd.example"),
		Args:    cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			if opts.outputFormat != "" && !flagutil.IsValidInput(opts.outputFormat, flagutil.ValidOutputFormats...) {
				return flagutil.InvalidValueError("output", opts.outputFormat, flagutil.ValidOutputFormats...)
			}

			return runList(opts)
		},
	}

	flags := kafkaFlagutil.NewFlagSet(cmd, opts.f.Localizer)

	flags.AddOutput(&opts.outputFormat)

	return cmd
}

func runList(opts *options) error {
	providers, err := create.GetEnabledCloudProviderNames(opts.f)
	if err != nil {
		return err
	}
	providersArray := make([]providerRow, 0, 10)
	for _, providerId := range providers {
		regions, err := create.GetEnabledCloudRegionIDs(opts.f, providerId, nil)
		if err != nil {
			return err
		}
		for _, region := range regions {
			providersArray = append(providersArray, providerRow{
				CloudProvider: providerId,
				Region:        region,
			})
		}
	}

	switch opts.outputFormat {
	case dump.EmptyFormat:
		dump.Table(opts.f.IOStreams.Out, providersArray)
		opts.f.Logger.Info("")
	default:
		return dump.Formatted(opts.f.IOStreams.Out, opts.outputFormat, providersArray)
	}
	return nil
}
