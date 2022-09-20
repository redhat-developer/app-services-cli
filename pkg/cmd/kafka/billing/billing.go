package billing

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
type billingRow struct {
	BillingType string `json:"billing" header:"Billing"`
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
	billingTypes, _ := create.GetBillingModelCompletionValues(opts.f)

	providersArray := make([]billingRow, 0, 10)
	if len(billingTypes) == 0 {
		providersArray = append(providersArray, billingRow{
			BillingType: create.DeveloperType,
		})
	} else {
		for _, billing := range billingTypes {
			providersArray = append(providersArray, billingRow{
				BillingType: billing,
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
