package list

import (
	"context"
	"fmt"
	"time"

	"github.com/redhat-developer/app-services-cli/internal/build"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/serviceaccount/svcaccountcmdutil/validation"
	"github.com/redhat-developer/app-services-cli/pkg/core/cmdutil"
	"github.com/redhat-developer/app-services-cli/pkg/core/cmdutil/flagutil"
	"github.com/redhat-developer/app-services-cli/pkg/core/config"
	"github.com/redhat-developer/app-services-cli/pkg/core/ioutil/dump"
	"github.com/redhat-developer/app-services-cli/pkg/core/ioutil/iostreams"
	"github.com/redhat-developer/app-services-cli/pkg/core/localize"
	"github.com/redhat-developer/app-services-cli/pkg/core/logging"
	"github.com/redhat-developer/app-services-cli/pkg/shared/factory"

	"github.com/spf13/cobra"

	svcacctmgmtclient "github.com/redhat-developer/app-services-sdk-go/serviceaccountmgmt/apiv1/client"
)

type options struct {
	Config     config.IConfig
	Connection factory.ConnectionFunc
	Logger     logging.Logger
	IO         *iostreams.IOStreams
	localizer  localize.Localizer
	Context    context.Context

	output       string
	enableAuthV2 bool
	page         int32
	size         int32
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

			validator := &validation.Validator{
				Localizer: opts.localizer,
			}

			if err := validator.ValidatePage(opts.page); err != nil {
				return err
			}

			if err := validator.ValidateSize(opts.size); err != nil {
				return err
			}

			return runList(opts)
		},
	}

	cmd.Flags().StringVarP(&opts.output, "output", "o", "", opts.localizer.MustLocalize("serviceAccount.list.flag.output.description"))
	cmd.Flags().BoolVar(&opts.enableAuthV2, "enable-auth-v2", false, opts.localizer.MustLocalize("serviceAccount.common.flag.enableAuthV2"))

	_ = cmd.Flags().MarkDeprecated("enable-auth-v2", opts.localizer.MustLocalize("serviceAccount.common.flag.deprecated.enableAuthV2"))

	cmd.Flags().Int32VarP(&opts.page, "page", "", int32(cmdutil.ConvertPageValueToInt32(build.DefaultPageNumber)), opts.localizer.MustLocalize("serviceAccount.list.flag.page.description"))
	// Default has been set to 100 to preserve how list worked before
	cmd.Flags().Int32VarP(&opts.size, "size", "", 100, opts.localizer.MustLocalize("serviceAccount.list.flag.size.description"))

	flagutil.EnableOutputFlagCompletion(cmd)

	return cmd
}

func runList(opts *options) (err error) {
	conn, err := opts.Connection()
	if err != nil {
		return err
	}

	a := conn.API().ServiceAccountMgmt().GetServiceAccounts(opts.Context)

	// Calculate offset based on page and size provided
	calculatedFirst := (opts.page - 1) * opts.size
	a = a.First(calculatedFirst)
	a = a.Max(opts.size)

	serviceaccounts, _, err := a.Execute()
	if err != nil {
		return err
	}

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
		return dump.Formatted(opts.IO.Out, opts.output, serviceaccounts)
	}

	return nil
}

func mapResponseItemsToRows(svcAccts []svcacctmgmtclient.ServiceAccountData) []svcAcctRow {
	rows := make([]svcAcctRow, len(svcAccts))

	for i, sa := range svcAccts {

		row := svcAcctRow{
			ID:        sa.GetId(),
			Name:      sa.GetName(),
			ClientID:  sa.GetClientId(),
			Owner:     sa.GetCreatedBy(),
			CreatedAt: unixTimestampToUTC(sa.GetCreatedAt()),
		}

		rows[i] = row
	}

	return rows
}

// unixTimestampToUTC converts a unix timestamp to the corresponding local Time
func unixTimestampToUTC(timestamp int64) string {
	localTime := time.Unix(timestamp, 0)
	return fmt.Sprint(localTime)
}
