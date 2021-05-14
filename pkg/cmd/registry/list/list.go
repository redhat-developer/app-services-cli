package list

import (
	"context"
	"encoding/json"

	srsclient "github.com/redhat-developer/app-services-cli/pkg/api/srs/client"
	flagutil "github.com/redhat-developer/app-services-cli/pkg/cmdutil/flags"
	"github.com/redhat-developer/app-services-cli/pkg/connection"
	"github.com/redhat-developer/app-services-cli/pkg/iostreams"
	"github.com/redhat-developer/app-services-cli/pkg/localize"

	"github.com/redhat-developer/app-services-cli/pkg/dump"

	"github.com/spf13/cobra"

	"github.com/redhat-developer/app-services-cli/internal/config"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/factory"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/flag"
	"github.com/redhat-developer/app-services-cli/pkg/logging"

	"gopkg.in/yaml.v2"
)

// row is the details of a Kafka instance needed to print to a table
type kafkaRow struct {
	ID     string `json:"id" header:"ID"`
	Name   string `json:"name" header:"Name"`
	URL    string `json:"registryUrl" header:"Registry URL"`
	Status string `json:"status" header:"Status"`
}

type options struct {
	outputFormat string
	// TODO ping team to discuss those options
	// page         int
	// limit        int
	// search       string

	IO         *iostreams.IOStreams
	Config     config.IConfig
	Connection factory.ConnectionFunc
	Logger     func() (logging.Logger, error)
	localizer  localize.Localizer
}

// NewListCommand creates a new command for listing kafkas.
func NewListCommand(f *factory.Factory) *cobra.Command {
	opts := &options{
		Config:     f.Config,
		Connection: f.Connection,
		Logger:     f.Logger,
		IO:         f.IOStreams,
		localizer:  f.Localizer,
	}

	cmd := &cobra.Command{
		Use:     "list",
		Short:   "List service registry",
		Long:    "",
		Example: "",
		Args:    cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			if opts.outputFormat != "" && !flagutil.IsValidInput(opts.outputFormat, flagutil.ValidOutputFormats...) {
				return flag.InvalidValueError("output", opts.outputFormat, flagutil.ValidOutputFormats...)
			}

			return runList(opts)
		},
	}

	cmd.Flags().StringVarP(&opts.outputFormat, "output", "o", "", opts.localizer.MustLocalize("kafkas.common.flag.output.description"))

	return cmd
}

func runList(opts *options) error {
	logger, err := opts.Logger()
	if err != nil {
		return err
	}

	connection, err := opts.Connection(connection.DefaultConfigSkipMasAuth)
	if err != nil {
		return err
	}

	api := connection.API()

	a := api.ServiceRegistry().GetRegistries(context.Background())

	response, _, err := a.Execute()

	if err != nil {
		return err
	}

	if len(response) == 0 && opts.outputFormat == "" {
		logger.Info(opts.localizer.MustLocalize("kafka.common.log.info.noKafkaInstances"))
		return nil
	}

	switch opts.outputFormat {
	case "json":
		data, _ := json.Marshal(response)
		_ = dump.JSON(opts.IO.Out, data)
	case "yaml", "yml":
		data, _ := yaml.Marshal(response)
		_ = dump.YAML(opts.IO.Out, data)
	default:
		rows := mapResponseItemsToRows(&response)
		dump.Table(opts.IO.Out, rows)
		logger.Info("")
	}

	return nil
}

func mapResponseItemsToRows(registries *[]srsclient.Registry) []kafkaRow {
	rows := []kafkaRow{}

	for _, k := range *registries {
		row := kafkaRow{
			ID:     string(k.GetId()),
			Name:   k.GetName(),
			URL:    k.GetRegistryUrl(),
			Status: string(k.GetStatus().Value),
		}

		rows = append(rows, row)
	}

	return rows
}
