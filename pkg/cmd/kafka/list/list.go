package list

import (
	"context"
	"encoding/json"
	"os"
	"strconv"

	kafkamsg "github.com/bf2fc6cc711aee1a0c2a/cli/internal/localizer/msg/kafka"

	kasclient "github.com/bf2fc6cc711aee1a0c2a/cli/pkg/api/kas/client"
	flagutil "github.com/bf2fc6cc711aee1a0c2a/cli/pkg/cmdutil/flags"

	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/dump"

	"github.com/spf13/cobra"

	"github.com/bf2fc6cc711aee1a0c2a/cli/internal/config"
	"github.com/bf2fc6cc711aee1a0c2a/cli/internal/localizer"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/cmd/factory"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/cmd/flag"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/connection"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/logging"

	"gopkg.in/yaml.v2"
)

// row is the details of a Kafka instance needed to print to a table
type kafkaRow struct {
	ID            string `json:"id" header:"ID"`
	Name          string `json:"name" header:"Name"`
	Owner         string `json:"owner" header:"Owner"`
	Status        string `json:"status" header:"Status"`
	CloudProvider string `json:"cloud_provider" header:"Cloud Provider"`
	Region        string `json:"region" header:"Region"`
}

type options struct {
	outputFormat string
	page         int
	limit        int

	Config     config.IConfig
	Connection func() (connection.Connection, error)
	Logger     func() (logging.Logger, error)
}

// NewListCommand creates a new command for listing kafkas.
func NewListCommand(f *factory.Factory) *cobra.Command {
	opts := &options{
		page:       0,
		limit:      100,
		Config:     f.Config,
		Connection: f.Connection,
		Logger:     f.Logger,
	}

	localizer.LoadMessageFiles("cmd/kafka", "cmd/kafka/list")

	cmd := &cobra.Command{
		Use:   localizer.MustLocalizeFromID("kafka.list.cmd.use"),
		Short: localizer.MustLocalizeFromID("kafka.list.cmd.shortDescription"),
		Long:  localizer.MustLocalizeFromID("kafka.list.cmd.longDescription"),
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			if opts.outputFormat != "" && !flagutil.IsValidInput(opts.outputFormat, flagutil.ValidOutputFormats...) {
				return flag.InvalidValueError("output", opts.outputFormat, flagutil.ValidOutputFormats...)
			}

			return runList(opts)
		},
	}

	cmd.Flags().StringVarP(&opts.outputFormat, "output", "o", "", localizer.MustLocalize(&localizer.Config{
		MessageID:   kafkamsg.OutputFlagDescription,
		PluralCount: 2,
	}))
	cmd.Flags().IntVarP(&opts.page, "page", "", 0, localizer.MustLocalizeFromID("kafka.list.flag.page"))
	cmd.Flags().IntVarP(&opts.limit, "limit", "", 100, localizer.MustLocalizeFromID("kafka.list.flag.limit"))

	return cmd
}

func runList(opts *options) error {
	logger, err := opts.Logger()
	if err != nil {
		return err
	}

	connection, err := opts.Connection()
	if err != nil {
		return err
	}

	api := connection.API()

	a := api.Kafka().ListKafkas(context.Background())
	a = a.Page(strconv.Itoa(opts.page))
	a = a.Size(strconv.Itoa(opts.limit))
	response, _, apiErr := a.Execute()

	if apiErr.Error() != "" {
		return apiErr
	}

	if response.Size == 0 {
		logger.Info(localizer.MustLocalizeFromID("kafka.list.log.info.noKafkaInstances"))
		return nil
	}

	switch opts.outputFormat {
	case "json":
		data, _ := json.Marshal(response)
		_ = dump.JSON(os.Stdout, data)
	case "yaml", "yml":
		data, _ := yaml.Marshal(response)
		_ = dump.YAML(os.Stdout, data)
	default:
		rows := mapResponseItemsToRows(response.GetItems())
		dump.Table(os.Stdout, rows)
		logger.Info("")
	}

	return nil
}

func mapResponseItemsToRows(kafkas []kasclient.KafkaRequest) []kafkaRow {
	rows := []kafkaRow{}

	for _, k := range kafkas {
		row := kafkaRow{
			ID:            k.GetId(),
			Name:          k.GetName(),
			Owner:         k.GetOwner(),
			Status:        k.GetStatus(),
			CloudProvider: k.GetCloudProvider(),
			Region:        k.GetRegion(),
		}

		rows = append(rows, row)
	}

	return rows
}
