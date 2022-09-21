package billing

import (
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
	BillingType string `json:"billing" header:"Billing"`
	AccountID   string `json:"marketplace_account_id" header:"Marketplace Account ID"`
	Provider    string `json:"marketplace_provider" header:"Marketplace Provider"`
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

	paidQuotas := orgQuotas.MarketplaceQuotas
	paidQuotas = append(paidQuotas, orgQuotas.StandardQuotas...)

	billingArray := make([]billingRow, 0, 10)

	if len(paidQuotas) == 0 {
		opts.f.Logger.Info(opts.f.Localizer.MustLocalize("kafka.billing.log.info.noStandardInstancesAvailable"))
		return nil
	}

	for _, quota := range paidQuotas {

		if quota.BillingModel == create.StandardType {

			billingArray = append(billingArray, billingRow{
				BillingType: quota.BillingModel,
			})

		} else {

			for _, cloudAccount := range *quota.CloudAccounts {
				billingArray = append(billingArray, billingRow{
					BillingType: quota.BillingModel,
					AccountID:   *cloudAccount.CloudAccountId,
					Provider:    *cloudAccount.CloudProviderId,
				})
			}

		}

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
