package create

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/bf2fc6cc711aee1a0c2a/cli/internal/localizer"

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

	localizer.LoadMessageFiles("cmd/kafka/common", "cmd/kafka/topic/common", "cmd/kafka/topic/create")

	cmd := &cobra.Command{
		Use:     localizer.MustLocalizeFromID("kafka.topic.create.cmd.use"),
		Short:   localizer.MustLocalizeFromID("kafka.topic.create.cmd.shortDescription"),
		Long:    localizer.MustLocalizeFromID("kafka.topic.create.cmd.longDescription"),
		Example: localizer.MustLocalizeFromID("kafka.topic.create.cmd.example"),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			if len(args) == 0 {
				return fmt.Errorf(localizer.MustLocalizeFromID("kafka.topic.create.cmd.error.topicNameIsRequired"))
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
				return fmt.Errorf(localizer.MustLocalizeFromID("kafka.topic.common.error.noKafkaSelected"))
			}

			opts.kafkaID = cfg.Services.Kafka.ClusterID

			return runCmd(opts)
		},
	}

	cmd.Flags().StringVarP(&opts.outputFormat, "output", "o", "json", localizer.MustLocalize(&localizer.Config{
		MessageID: "kafka.topic.common.flag.output.description",
	}))
	cmd.Flags().Int32Var(&opts.partitions, "partitions", 1, localizer.MustLocalizeFromID("kafka.topic.common.flag.partitions"))
	cmd.Flags().Int32Var(&opts.replicas, "replicas", 1, localizer.MustLocalizeFromID("kafka.topic.common.flag.replicas"))
	cmd.Flags().IntVar(&opts.retentionMs, "retention-ms", -1, localizer.MustLocalizeFromID("kafka.topic.common.flag.retentionMs"))

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
		return errors.New(localizer.MustLocalize(&localizer.Config{
			MessageID: "kafka.common.error.notFoundByIdError",
			TemplateData: map[string]interface{}{
				"ID": opts.kafkaID,
			},
		}))
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
			return fmt.Errorf(localizer.MustLocalize(&localizer.Config{
				MessageID: "kafka.topic.common.error.unauthorized",
				TemplateData: map[string]interface{}{
					"Operation": "create",
				},
			}))
		case 409:
			return fmt.Errorf(localizer.MustLocalize(&localizer.Config{
				MessageID: "kafka.topic.create.error.conflictError",
				TemplateData: map[string]interface{}{
					"TopicName":    opts.topicName,
					"InstanceName": kafkaInstance.GetName(),
				},
			}))
		case 500:
			return fmt.Errorf("%v: %w", localizer.MustLocalizeFromID("kafka.topic.common.error.internalServerError"), topicErr)
		case 503:
			return fmt.Errorf("%v: %w", localizer.MustLocalize(&localizer.Config{
				MessageID: "kafka.topic.common.error.unableToConnectToKafka",
				TemplateData: map[string]interface{}{
					"Name": kafkaInstance.GetName(),
				},
			}), topicErr)
		default:
			return topicErr
		}
	}

	logger.Info(localizer.MustLocalize(&localizer.Config{
		MessageID: "kafka.topic.create.log.info.topicCreated",
		TemplateData: map[string]interface{}{
			"TopicName":    response.GetName(),
			"InstanceName": kafkaInstance.GetName(),
		},
	}))

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
