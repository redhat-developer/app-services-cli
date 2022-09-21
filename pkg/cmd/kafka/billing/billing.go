package billing

import (
	"fmt"

	"github.com/redhat-developer/app-services-cli/pkg/cmd/kafka/create"
	kafkaFlagutil "github.com/redhat-developer/app-services-cli/pkg/cmd/kafka/flagutil"

	"github.com/redhat-developer/app-services-cli/pkg/core/cmdutil/flagutil"
	"github.com/redhat-developer/app-services-cli/pkg/core/ioutil/dump"
	"github.com/redhat-developer/app-services-cli/pkg/core/servicecontext"
	"github.com/redhat-developer/app-services-cli/pkg/shared/accountmgmtutil"
	"github.com/redhat-developer/app-services-cli/pkg/shared/factory"
	"github.com/redhat-developer/app-services-cli/pkg/shared/remote"

	"github.com/spf13/cobra"
)

// row is the details of a Kafka instance needed to print to a table
type billingRow struct {
	CloudProvider string `json:"cloud_provider" header:"Cloud Provider"`
	Region        string `json:"region" header:"Region"`
	BillingType   string `json:"billing" header:"Billing"`
	Plan          string `json:"plan" header:"Plan"`
}

type options struct {
	outputFormat string

	f              *factory.Factory
	ServiceContext servicecontext.IContext
}

// NewListCommand creates a new command for listing kafkas.
func NewBillingCommand(f *factory.Factory) *cobra.Command {
	opts := &options{
		f: f,
	}

	cmd := &cobra.Command{
		Use:     "billing",
		Short:   opts.f.Localizer.MustLocalize("kafka.billing.cmd.shortDescription"),
		Long:    opts.f.Localizer.MustLocalize("kafka.billing.cmd.longDescription"),
		Example: opts.f.Localizer.MustLocalize("kafka.billing.cmd.example"),
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

	err, constants := remote.GetRemoteServiceConstants(opts.f.Context, opts.f.Logger)
	if err != nil {
		return err
	}

	orgQuotas, err := accountmgmtutil.GetOrgQuotas(opts.f, &constants.Kafka.Ams)
	if err != nil {
		return err
	}

	billingTypes := create.FetchSupportedBillingModels(orgQuotas)

	billingArray := make([]billingRow, 0, 10)
	if len(billingTypes) == 0 {
		billingArray = append(billingArray, billingRow{
			BillingType: create.DeveloperType,
		})
	} else {

		quotas := append(orgQuotas.MarketplaceQuotas, orgQuotas.StandardQuotas...)

		for _, quota := range quotas {

			providers, err := create.GetEnabledCloudProviderNames(opts.f)
			if err != nil {
				return err
			}
			for _, providerId := range providers {
				regions, err := create.GetEnabledCloudRegionIDs(opts.f, providerId, nil)
				if err != nil {
					return err
				}
				for _, region := range regions {
					kafkaSizes, _ := create.FetchValidKafkaSizes(opts.f, "aws", "us-east-1", quota)
					sizeLabels := create.GetValidKafkaSizesLabels(kafkaSizes)
					for _, sizeLabel := range sizeLabels {
						billingArray = append(billingArray, billingRow{
							CloudProvider: providerId,
							Region:        region,
							BillingType:   quota.BillingModel,
							Plan:          fmt.Sprintf("%s.%s", quota.BillingModel, sizeLabel),
						})
					}
				}
			}

		}
		// for _, billing := range billingTypes {

		// 	kafkaSizes := create.FetchValidKafkaSizes(opts.f, "aws", "us-east-1", orgQuotas)

		// 	billingArray = append(billingArray, billingRow{
		// 		BillingType: billing,
		// 	})
		// }
	}

	switch opts.outputFormat {
	case dump.EmptyFormat:
		dump.Table(opts.f.IOStreams.Out, billingArray)
		opts.f.Logger.Info("")
	default:
		return dump.Formatted(opts.f.IOStreams.Out, opts.outputFormat, billingArray)
	}
	return nil
}
