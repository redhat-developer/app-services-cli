package list

import (
	"context"

	"github.com/redhat-developer/app-services-cli/pkg/core/cmdutil/flagutil"
	"github.com/redhat-developer/app-services-cli/pkg/core/config"
	"github.com/redhat-developer/app-services-cli/pkg/core/ioutil/dump"
	"github.com/redhat-developer/app-services-cli/pkg/core/ioutil/iostreams"
	"github.com/redhat-developer/app-services-cli/pkg/core/localize"
	"github.com/redhat-developer/app-services-cli/pkg/core/logging"
	"github.com/redhat-developer/app-services-cli/pkg/shared/connection"
	"github.com/redhat-developer/app-services-cli/pkg/shared/factory"
	kafkamgmtclient "github.com/redhat-developer/app-services-sdk-go/kafkamgmt/apiv1/client"
	"github.com/spf13/cobra"
)

type options struct {
	Config     config.IConfig
	Connection factory.ConnectionFunc
	Logger     logging.Logger
	IO         *iostreams.IOStreams
	localizer  localize.Localizer
	Context    context.Context

	output string
}

// svcAcctRow contains the properties used to
// populate the list of service accounts into a table row
type svcAcctRow struct {
	ID        string `json:"id" header:"ID"`
	ClientID  string `json:"clientID" header:"Client ID"`
	Name      string `json:"name" header:"Short Description"`
	Owner     string `json:"owner" header:"Owner"`
	CreatedAt string `json:"createdAt" header:"Created At"`
}

// NewListCommand creates a new command to list service accounts
func NewListCommand(f *factory.Factory) *cobra.Command {
	opts := &options{
		Config:     f.Config,
		Connection: f.Connection,
		Logger:     f.Logger,
		IO:         f.IOStreams,
		localizer:  f.Localizer,
		Context:    f.Context,
	}

	cmd := &cobra.Command{
		Use:     "list",
		Short:   opts.localizer.MustLocalize("serviceAccount.list.cmd.shortDescription"),
		Long:    opts.localizer.MustLocalize("serviceAccount.list.cmd.longDescription"),
		Example: opts.localizer.MustLocalize("serviceAccount.list.cmd.example"),
		Args:    cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			if opts.output != "" && !flagutil.IsValidInput(opts.output, flagutil.ValidOutputFormats...) {
				return flagutil.InvalidValueError("output", opts.output, flagutil.ValidOutputFormats...)
			}

			return runList(opts)
		},
	}

	cmd.Flags().StringVarP(&opts.output, "output", "o", "", opts.localizer.MustLocalize("serviceAccount.list.flag.output.description"))

	flagutil.EnableOutputFlagCompletion(cmd)

	return cmd
}

func runList(opts *options) (err error) {
	conn, err := opts.Connection(connection.DefaultConfigSkipMasAuth)
	if err != nil {
		return err
	}

	res, _, err := conn.API().ServiceAccountMgmt().GetServiceAccounts(opts.Context).Execute()
	if err != nil {
		return err
	}

	serviceaccounts := res.GetItems()
	if len(serviceaccounts) == 0 && opts.output == "" {
		opts.Logger.Info(opts.localizer.MustLocalize("serviceAccount.list.log.info.noneFound"))
		return nil
	}

	outStream := opts.IO.Out
	switch opts.output {
	case dump.EmptyFormat:
		rows := mapResponseItemsToRows(serviceaccounts)
		dump.Table(outStream, rows)
	default:
		return dump.Formatted(opts.IO.Out, opts.output, res)
	}

	return nil
}

func mapResponseItemsToRows(svcAccts []kafkamgmtclient.ServiceAccountListItem) []svcAcctRow {
	rows := make([]svcAcctRow, len(svcAccts))

	for i, sa := range svcAccts {
		row := svcAcctRow{
			ID:        sa.GetId(),
			Name:      sa.GetName(),
			ClientID:  sa.GetClientId(),
			Owner:     sa.GetOwner(),
			CreatedAt: sa.GetCreatedAt().String(),
		}

		rows[i] = row
	}

	return rows
}
