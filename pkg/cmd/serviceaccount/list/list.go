package list

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/redhat-developer/app-services-cli/internal/config"
	"github.com/redhat-developer/app-services-cli/internal/localizer"
	kasclient "github.com/redhat-developer/app-services-cli/pkg/api/kas/client"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/factory"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/flag"
	"github.com/redhat-developer/app-services-cli/pkg/cmdutil"
	flagutil "github.com/redhat-developer/app-services-cli/pkg/cmdutil/flags"
	"github.com/redhat-developer/app-services-cli/pkg/connection"
	"github.com/redhat-developer/app-services-cli/pkg/dump"
	"github.com/redhat-developer/app-services-cli/pkg/iostreams"
	"github.com/redhat-developer/app-services-cli/pkg/logging"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

type Options struct {
	Config     config.IConfig
	Connection factory.ConnectionFunc
	Logger     func() (logging.Logger, error)
	IO         *iostreams.IOStreams

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
	}

	cmd := &cobra.Command{
		Use:     localizer.MustLocalizeFromID("serviceAccount.list.cmd.use"),
		Short:   localizer.MustLocalizeFromID("serviceAccount.list.cmd.shortDescription"),
		Long:    localizer.MustLocalizeFromID("serviceAccount.list.cmd.longDescription"),
		Example: localizer.MustLocalizeFromID("serviceAccount.list.cmd.example"),
		RunE: func(cmd *cobra.Command, _ []string) error {
			if opts.output != "" && !flagutil.IsValidInput(opts.output, flagutil.ValidOutputFormats...) {
				return flag.InvalidValueError("output", opts.output, flagutil.ValidOutputFormats...)
			}

			return runList(opts)
		},
	}

	cmd.Flags().StringVarP(&opts.output, "output", "o", "", localizer.MustLocalizeFromID("serviceAccount.common.flag.output.description"))

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
	res, httpRes, err := a.Execute()

	if err != nil {
		if httpRes == nil {
			return err
		}

		switch httpRes.StatusCode {
		case 403:
			return fmt.Errorf("%v: %w", localizer.MustLocalize(&localizer.Config{
				MessageID:   "serviceAccount.common.error.forbidden",
				PluralCount: 2,
				TemplateData: map[string]interface{}{
					"Operation": "list",
				},
			}), err)
		case 500:
			return fmt.Errorf("%v: %w", localizer.MustLocalizeFromID("serviceAccount.common.error.internalServerError"), err)
		default:
			return err
		}
	}

	serviceaccounts := res.GetItems()
	if len(serviceaccounts) == 0 && opts.output == "" {
		logger.Info(localizer.MustLocalizeFromID("serviceAccount.list.log.info.noneFound"))
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
