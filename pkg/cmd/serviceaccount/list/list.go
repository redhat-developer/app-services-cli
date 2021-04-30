package list

import (
	"context"
	"encoding/json"

	"github.com/redhat-developer/app-services-cli/internal/config"
	kasclient "github.com/redhat-developer/app-services-cli/pkg/api/kas/client"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/factory"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/flag"
	"github.com/redhat-developer/app-services-cli/pkg/cmdutil"
	flagutil "github.com/redhat-developer/app-services-cli/pkg/cmdutil/flags"
	"github.com/redhat-developer/app-services-cli/pkg/connection"
	"github.com/redhat-developer/app-services-cli/pkg/dump"
	"github.com/redhat-developer/app-services-cli/pkg/iostreams"
	"github.com/redhat-developer/app-services-cli/pkg/localize"
	"github.com/redhat-developer/app-services-cli/pkg/logging"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

type Options struct {
	Config     config.IConfig
	Connection factory.ConnectionFunc
	Logger     func() (logging.Logger, error)
	IO         *iostreams.IOStreams
	localizer  localize.Localizer

	output string
}

// svcAcctRow contains the properties used to
// populate the list of service accounts into a table row
type svcAcctRow struct {
	ID        string `json:"id" header:"ID"`
	Name      string `json:"name" header:"Name"`
	ClientID  string `json:"clientID" header:"Client ID"`
	Owner     string `json:"owner" header:"Owner"`
	CreatedAt string `json:"createdAt" header:"Created At"`
}

// NewListCommand creates a new command to list service accounts
func NewListCommand(f *factory.Factory) *cobra.Command {
	opts := &Options{
		Config:     f.Config,
		Connection: f.Connection,
		Logger:     f.Logger,
		IO:         f.IOStreams,
		localizer:  f.Localizer,
	}

	cmd := &cobra.Command{
		Use:     opts.localizer.MustLocalize("serviceAccount.list.cmd.use"),
		Short:   opts.localizer.MustLocalize("serviceAccount.list.cmd.shortDescription"),
		Long:    opts.localizer.MustLocalize("serviceAccount.list.cmd.longDescription"),
		Example: opts.localizer.MustLocalize("serviceAccount.list.cmd.example"),
		Args:    cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			if opts.output != "" && !flagutil.IsValidInput(opts.output, flagutil.ValidOutputFormats...) {
				return flag.InvalidValueError("output", opts.output, flagutil.ValidOutputFormats...)
			}

			return runList(opts)
		},
	}

	cmd.Flags().StringVarP(&opts.output, "output", "o", "", opts.localizer.MustLocalize("serviceAccount.list.flag.output.description"))

	return cmd
}

func runList(opts *Options) (err error) {
	logger, err := opts.Logger()
	if err != nil {
		return err
	}

	connection, err := opts.Connection(connection.DefaultConfigSkipMasAuth)
	if err != nil {
		return err
	}

	api := connection.API()

	a := api.Kafka().ListServiceAccounts(context.Background())
	res, _, err := a.Execute()

	if err != nil {
		return err
	}

	serviceaccounts := res.GetItems()
	if len(serviceaccounts) == 0 && opts.output == "" {
		logger.Info(opts.localizer.MustLocalize("serviceAccount.list.log.info.noneFound"))
		return nil
	}

	outStream := opts.IO.Out
	switch opts.output {
	case "json":
		data, _ := json.MarshalIndent(res, "", cmdutil.DefaultJSONIndent)
		_ = dump.JSON(outStream, data)
	case "yaml", "yml":
		data, _ := yaml.Marshal(res)
		_ = dump.YAML(outStream, data)
	default:
		rows := mapResponseItemsToRows(serviceaccounts)
		dump.Table(outStream, rows)
	}

	return nil
}

func mapResponseItemsToRows(svcAccts []kasclient.ServiceAccountListItem) []svcAcctRow {
	rows := []svcAcctRow{}

	for _, sa := range svcAccts {
		row := svcAcctRow{
			ID:        sa.GetId(),
			Name:      sa.GetName(),
			ClientID:  sa.GetClientID(),
			Owner:     sa.GetOwner(),
			CreatedAt: sa.GetCreatedAt().String(),
		}

		rows = append(rows, row)
	}

	return rows
}
