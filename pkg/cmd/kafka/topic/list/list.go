package list

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/MakeNowJust/heredoc"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/api/kas"
	strimziadminclient "github.com/bf2fc6cc711aee1a0c2a/cli/pkg/api/strimzi-admin/client"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/color"

	"gopkg.in/yaml.v2"

	"github.com/bf2fc6cc711aee1a0c2a/cli/internal/config"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/cmd/factory"
	flagutil "github.com/bf2fc6cc711aee1a0c2a/cli/pkg/cmdutil/flags"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/connection"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/dump"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/iostreams"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/logging"
	"github.com/spf13/cobra"
)

type Options struct {
	Config     config.IConfig
	IO         *iostreams.IOStreams
	Connection func() (connection.Connection, error)
	Logger     func() (logging.Logger, error)

	kafkaID string
	output  string
}

type topicRow struct {
	Name            string `json:"name,omitempty" header:"Name"`
	PartitionsCount int    `json:"partitions_count,omitempty" header:"Partitions"`
}

// NewListTopicCommand gets a new command for getting kafkas.
func NewListTopicCommand(f *factory.Factory) *cobra.Command {
	opts := &Options{
		Config:     f.Config,
		Connection: f.Connection,
		Logger:     f.Logger,
		IO:         f.IOStreams,
	}

	cmd := &cobra.Command{
		Use:   "list",
		Short: "List Kafka topics",
		Long: heredoc.Docf(`
			List all topics in a Kafka instance
		`),
		Example: heredoc.Doc(`
			# list all topics
			$ rhoas kafka topic list

			# list all topics as JSON
			$ rhoas kafka topic list -o json
		`),
		RunE: func(cmd *cobra.Command, _ []string) error {
			logger, err := opts.Logger()
			if err != nil {
				return err
			}

			if opts.output != "" && !flagutil.IsValidInput(opts.output, flagutil.ValidOutputFormats...) {
				logger.Infof("Unknown flag value '%v' for --output. Using table format instead", opts.output)
				opts.output = ""
			}

			return runCmd(opts)
		},
	}

	cmd.Flags().StringVarP(&opts.output, "output", "o", "", fmt.Sprintf("Output format of the results. Choose from %q.", flagutil.ValidOutputFormats))

	return cmd
}

func runCmd(opts *Options) error {
	conn, err := opts.Connection()
	if err != nil {
		return err
	}

	logger, err := opts.Logger()
	if err != nil {
		return err
	}

	api := conn.API()

	ctx := context.Background()
	kafkaInstance, _, apiErr := api.Kafka().GetKafkaById(ctx, opts.kafkaID).Execute()
	if kas.IsErr(apiErr, kas.ErrorNotFound) {
		return fmt.Errorf("Kafka instance with ID '%v' not found", opts.kafkaID)
	} else if apiErr.Error() != "" {
		return apiErr
	}

	a := api.TopicAdmin(opts.kafkaID).GetTopicsList(context.Background())
	topicData, _, topicErr := a.Execute()

	if topicErr.Error() != "" {
		return topicErr
	}

	if topicData.GetCount() == 0 {
		logger.Infof("Kafka instance %v has no topics. Run %v to create a topic.", color.Info(kafkaInstance.GetName()), color.CodeSnippet("rhoas kafka topic create"))

		return nil
	}

	stdout := opts.IO.Out
	switch opts.output {
	case "json":
		data, _ := json.Marshal(topicData)
		_ = dump.JSON(stdout, data)
	case "yaml", "yml":
		data, _ := yaml.Marshal(topicData)
		_ = dump.YAML(stdout, data)
	default:
		topics := topicData.GetTopics()
		rows := mapTopicResultsToTableFormat(topics)
		dump.Table(stdout, rows)
	}

	return err
}

func mapTopicResultsToTableFormat(topics []strimziadminclient.Topic) []topicRow {
	var rows []topicRow = []topicRow{}

	for _, t := range topics {
		row := topicRow{
			Name:            t.GetName(),
			PartitionsCount: len(t.GetPartitions()),
		}
		rows = append(rows, row)
	}

	return rows
}
