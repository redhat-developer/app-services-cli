package create

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/MakeNowJust/heredoc"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/color"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/kafka/topic"

	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/cmd/flag"

	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/api/kas"
	strimziadminclient "github.com/bf2fc6cc711aee1a0c2a/cli/pkg/api/strimzi-admin/client"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/dump"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/iostreams"
	"gopkg.in/yaml.v2"

	"github.com/bf2fc6cc711aee1a0c2a/cli/internal/config"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/cmd/factory"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/connection"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/logging"

	"github.com/spf13/cobra"
)

const (
	Partitions = "partitions"
	Replicas   = "replicas"
)

type Options struct {
	topicName    string
	partitions   int32
	replicas     int32
	retentionMs  int
	kafkaID      string
	outputFormat string

	IO         *iostreams.IOStreams
	Config     config.IConfig
	Connection func() (connection.Connection, error)
	Logger     func() (logging.Logger, error)
}

// NewCreateTopicCommand gets a new command for creating kafka topic.
func NewCreateTopicCommand(f *factory.Factory) *cobra.Command {
	opts := &Options{
		Connection: f.Connection,
		Config:     f.Config,
		Logger:     f.Logger,
		IO:         f.IOStreams,
	}

	cmd := &cobra.Command{
		Use:   "create",
		Short: "Create a Kafka topic",
		Long: heredoc.Doc(`
			Create topic in the current Kafka instance.

			This command lets you create a topic, set a desired number of 
			partitions, replicas and retention period or else use the default values.
		`),
		Example: heredoc.Doc(`
			# create a topic
			$ rhoas kafka topic create topic-1
		`),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			if len(args) == 0 {
				return fmt.Errorf(`Topic name is required. Run "rhoas kafka topic create <topic-name>"`)
			}
			opts.topicName = args[0]

			if err = flag.ValidateOutput(opts.outputFormat); err != nil {
				return err
			}

			if err = topic.ValidateName(opts.topicName); err != nil {
				return err
			}

			if err = topic.ValidatePartitionsN(opts.partitions); err != nil {
				return err
			}

			if err = topic.ValidateReplicationFactorN(opts.replicas); err != nil {
				return err
			}

			if err = topic.ValidateMessageRetentionPeriod(opts.retentionMs); err != nil {
				return err
			}

			if opts.kafkaID != "" {
				return runCmd(opts)
			}

			cfg, err := opts.Config.Load()
			if err != nil {
				return err
			}

			if !cfg.HasKafka() {
				return fmt.Errorf("No Kafka instance selected. Use the '--id' flag or set one in context with the 'use' command")
			}

			opts.kafkaID = cfg.Services.Kafka.ClusterID

			return runCmd(opts)
		},
	}

	fs := cmd.Flags()
	flag.AddOutput(fs, &opts.outputFormat, "json", []string{"ks"})
	cmd.Flags().Int32Var(&opts.partitions, "partitions", 1, "The number of partitions in the topic")
	cmd.Flags().Int32Var(&opts.replicas, "replicas", 1, "The replication factor for the topic")
	cmd.Flags().IntVar(&opts.retentionMs, "retention-ms", -1, "The period of time in milliseconds the broker will retain a partition log before deleting it")

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

	createTopicReq := api.TopicAdmin(opts.kafkaID).CreateTopic(ctx)

	topicInput := strimziadminclient.NewTopicInput{
		Name: opts.topicName,
		Settings: &strimziadminclient.TopicSettings{
			ReplicationFactor: &opts.replicas,
			NumPartitions:     &opts.partitions,
			Config:            topic.CreateConfig(opts.retentionMs),
		},
	}
	createTopicReq = createTopicReq.NewTopicInput(topicInput)

	response, httpRes, topicErr := createTopicReq.Execute()
	if topicErr.Error() != "" {
		switch httpRes.StatusCode {
		case 401:
			return fmt.Errorf("you are unauthorized to create this topic")
		case 409:
			return fmt.Errorf("topic '%v' already exists in Kafka instance '%v'", opts.topicName, kafkaInstance.GetName())
		case 500:
			return fmt.Errorf("internal server error: %w", topicErr)
		case 503:
			return fmt.Errorf("unable to connect to Kafka instance '%v': %w", kafkaInstance.GetName(), topicErr)
		default:
			return topicErr
		}
	}

	logger.Infof("Topic %v created in Kafka instance %v:", color.Info(response.GetName()), color.Info((kafkaInstance.GetName())))
	switch opts.outputFormat {
	case "json":
		data, _ := json.Marshal(response)
		_ = dump.JSON(opts.IO.Out, data)
	case "yaml", "yml":
		data, _ := yaml.Marshal(response)
		_ = dump.YAML(opts.IO.Out, data)
	}

	return nil
}
