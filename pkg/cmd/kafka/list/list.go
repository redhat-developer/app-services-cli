package list

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strconv"

	"github.com/MakeNowJust/heredoc"
	flagutil "github.com/bf2fc6cc711aee1a0c2a/cli/pkg/cmdutil/flags"

	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/dump"

	"github.com/spf13/cobra"

	"github.com/bf2fc6cc711aee1a0c2a/cli/internal/config"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/cmd/factory"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/cmdutil"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/connection"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/logging"

	"gopkg.in/yaml.v2"

	pkgKafka "github.com/bf2fc6cc711aee1a0c2a/cli/pkg/kafka"
)

// Cluster is the details of a Kafka instance
type cluster struct {
	ID                  string `json:"id" header:"ID"`
	Name                string `json:"name" header:"Name"`
	Owner               string `json:"owner" header:"Owner"`
	Kind                string `json:"kind"`
	Href                string `json:"href"`
	Status              string `json:"status" header:"Status"`
	CloudProvider       string `json:"cloud_provider" header:"Cloud Provider"`
	Region              string `json:"region" header:"Region"`
	BootstrapServerHost string `json:"bootstrapServerHost"`
	CreatedAt           string `json:"created_at"`
	UpdatedAt           string `json:"updated_at"`
}

// clusterList contains a list of Kafka instances
type clusterList struct {
	Kind  string    `json:"kind"`
	Page  int       `json:"page"`
	Size  int       `json:"size"`
	Total int       `json:"total"`
	Items []cluster `json:"items"`
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

	cmd := &cobra.Command{
		Use:   "list",
		Short: "List all Kafka instances",
		Long: heredoc.Doc(`
			List all Kafka instances.
			
			This command will provide a high level view of all Kafka instances.
			The fields displayed are: ID, Name, Owner, Status, Cloud Provider, Region.
			Use the describe command to view all fields for a specific instance.

			The instances are displayed by default in a table, but can also be displayed as JSON or YAML.
		`),
		Example: heredoc.Doc(`
			# list all Kafka instances using the default (plain) output format
			$ rhoas kafka list

			# list all Kafka instances using JSON as the output format
			$ rhoas kafka list -o json
		`),
		Args: cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			logger, err := opts.Logger()
			if err != nil {
				return err
			}

			if !flagutil.IsValidInput(opts.outputFormat, flagutil.AllowedListFormats...) {
				logger.Infof("Unknown flag value '%v' for --output. Using plain format instead", opts.outputFormat)
				opts.outputFormat = "plain"
			}

			return runList(opts)
		},
	}

	cmd.Flags().StringVarP(&opts.outputFormat, "output", "o", "plain", fmt.Sprintf("Output format of the results. Choose from %q", flagutil.AllowedListFormats))
	cmd.Flags().IntVarP(&opts.page, "page", "", 0, "Page that should be returned from server")
	cmd.Flags().IntVarP(&opts.limit, "limit", "", 100, "Limit of items that should be returned from server")
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

	a := api.Kafka.ListKafkas(context.Background())
	a = a.Page(strconv.Itoa(opts.page))
	a = a.Size(strconv.Itoa(opts.limit))
	response, _, apiErr := a.Execute()
	// modify the items to add a :443 port to the bootstrap URL
	kafkaItems := response.GetItems()
	kafkaItems = pkgKafka.TransformKafkaRequestListItems(kafkaItems)
	response.SetItems(kafkaItems)

	if apiErr.Error() != "" {
		return fmt.Errorf("Unable to list Kafka instances: %w", apiErr)
	}

	if response.Size == 0 {
		logger.Info("No Kafka instances found.")
		return nil
	}

	jsonResponse, _ := json.Marshal(response)

	outputFormat := opts.outputFormat

	var kafkaList clusterList

	if err = json.Unmarshal(jsonResponse, &kafkaList); err != nil {
		logger.Infof("Could not unmarshal Kakfa list into table, defaulting to JSON: %v", err)
		outputFormat = "json"
	}

	switch outputFormat {
	case "json":
		data, _ := json.MarshalIndent(response, "", cmdutil.DefaultJSONIndent)
		_ = dump.JSON(os.Stdout, data)
	case "yaml", "yml":
		data, _ := yaml.Marshal(response)
		_ = dump.YAML(os.Stdout, data)
	default:
		dump.Table(os.Stdout, kafkaList.Items)
		logger.Info("")
	}

	return nil
}
