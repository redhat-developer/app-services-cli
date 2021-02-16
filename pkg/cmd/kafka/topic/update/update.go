package update

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"

	"github.com/bf2fc6cc711aee1a0c2a/cli/internal/localizer"

	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/cmdutil"

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

var (
	partitionCount    int32
	replicaCount      int32
	retentionPeriodMs int
)

type Options struct {
	topicName       string
	partitionsStr   string
	retentionMsStr  string
	replicaCountStr string
	kafkaID         string
	outputFormat    string

	IO         *iostreams.IOStreams
	Config     config.IConfig
	Connection func() (connection.Connection, error)
	Logger     func() (logging.Logger, error)
}

// NewUpdateTopicCommand gets a new command for updating a kafka topic.
// nolint:funlen
func NewUpdateTopicCommand(f *factory.Factory) *cobra.Command {
	opts := &Options{
		Connection: f.Connection,
		Config:     f.Config,
		Logger:     f.Logger,
		IO:         f.IOStreams,
	}

	localizer.LoadMessageFiles("cmd/kafka/topic/common", "cmd/kafka/topic/update")

	cmd := &cobra.Command{
		Use:     localizer.MustLocalizeFromID("kafka.topic.update.cmd.use"),
		Short:   localizer.MustLocalizeFromID("kafka.topic.update.cmd.shortDescription"),
		Long:    localizer.MustLocalizeFromID("kafka.topic.update.cmd.longDescription"),
		Example: localizer.MustLocalizeFromID("kafka.topic.update.cmd.example"),
		Args:    cobra.ExactArgs(1),
		// Dynamic completion of the topic name
		ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
			validNames := []string{}

			cfg, err := opts.Config.Load()
			if err != nil {
				return validNames, cobra.ShellCompDirectiveError
			}

			if !cfg.HasKafka() {
				return validNames, cobra.ShellCompDirectiveError
			}

			opts.kafkaID = cfg.Services.Kafka.ClusterID

			return cmdutil.FilterValidTopicNameArgs(f, opts.kafkaID, toComplete)
		},
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			if len(args) > 0 {
				opts.topicName = args[0]
			}

			logger, err := opts.Logger()
			if err != nil {
				return err
			}

			if opts.retentionMsStr == "" && opts.partitionsStr == "" && opts.replicaCountStr == "" {
				logger.Info(localizer.MustLocalizeFromID("kafka.topic.update.log.info.nothingToUpdate"))
			}

			if err = flag.ValidateOutput(opts.outputFormat); err != nil {
				return err
			}

			if err = topic.ValidateName(opts.topicName); err != nil {
				return err
			}

			// check if the partition flag is set
			// and if so try to convert the value from string to int32
			if opts.partitionsStr != "" {
				// convert the value from "partitions" to int32
				// nolint:govet
				p, err := strconv.ParseInt(opts.partitionsStr, 10, 32)
				if err != nil {
					return flag.InvalidArgumentError("--partitions", opts.partitionsStr, err)
				}
				partitionCount = int32(p)

				if err = topic.ValidatePartitionsN(partitionCount); err != nil {
					return err
				}
			}

			// check if the replica flag is set
			// and if so try to convert the value from string to int32
			if opts.replicaCountStr != "" {
				// convert the value from "replicas" to int32
				// nolint:govet
				p, err := strconv.ParseInt(opts.replicaCountStr, 10, 32)
				if err != nil {
					return flag.InvalidArgumentError("--replicas", opts.replicaCountStr, err)
				}
				replicaCount = int32(p)

				if err = topic.ValidateReplicationFactorN(replicaCount); err != nil {
					return err
				}
			}

			// check if the retention flag is set
			// and if so try to convert the value from string to int
			if opts.retentionMsStr != "" {
				// convert the value from "--retention-ms" to int
				// nolint:govet
				retentionPeriodMs, err = strconv.Atoi(opts.retentionMsStr)
				if err != nil {
					return flag.InvalidArgumentError("--retention-ms", opts.retentionMsStr, err)
				}

				if err = topic.ValidateMessageRetentionPeriod(retentionPeriodMs); err != nil {
					return err
				}
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
	cmd.Flags().StringVar(&opts.partitionsStr, "partitions", "", localizer.MustLocalizeFromID("kafka.topic.common.flag.partitions.description"))
	cmd.Flags().StringVar(&opts.replicaCountStr, "replicas", "", localizer.MustLocalizeFromID("kafka.topic.common.flag.replicas.description"))
	cmd.Flags().StringVar(&opts.retentionMsStr, "retention-ms", "", localizer.MustLocalizeFromID("kafka.topic.common.flag.retentionMs.description"))

	return cmd
}

// nolint:funlen
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

	// Check if the Kafka instance exists
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

	// track if any values have changed
	var needsUpdate bool

	topicToUpdate, httpRes, _ := api.TopicAdmin(opts.kafkaID).GetTopic(ctx, opts.topicName).Execute()
	if httpRes.StatusCode == 404 {
		return errors.New(localizer.MustLocalize(&localizer.Config{
			MessageID: "kafka.topic.update.error.topicNotFoundError",
			TemplateData: map[string]interface{}{
				"TopicName":    opts.topicName,
				"InstanceName": kafkaInstance.GetName(),
			},
		}))
	}

	currentPartitionCount := len(topicToUpdate.GetPartitions())

	updateTopicReq := api.TopicAdmin(opts.kafkaID).UpdateTopic(ctx, opts.topicName)

	topicSettings := &strimziadminclient.TopicSettings{}

	// Only set partitions if the flag was set
	if opts.partitionsStr != "" {
		if int(partitionCount) < currentPartitionCount {

			return errors.New(localizer.MustLocalize(&localizer.Config{
				MessageID: "kafka.topic.update.error.cannotDecreasePartitionCountError",
				TemplateData: map[string]interface{}{
					"From": currentPartitionCount,
					"To":   partitionCount,
				},
			}))
		}
		if int(partitionCount) == currentPartitionCount {
			logger.Infof(localizer.MustLocalize(&localizer.Config{
				MessageID: "kafka.topic.update.log.info.samePartitionCount",
				TemplateData: map[string]interface{}{
					"Name":  opts.topicName,
					"Count": currentPartitionCount,
				},
			}))

		} else {
			needsUpdate = true
			topicSettings.NumPartitions = &partitionCount
		}
	}

	// Update replica count
	if opts.replicaCountStr != "" {
		needsUpdate = true
		topicSettings.ReplicationFactor = &replicaCount
	}

	if opts.retentionMsStr != "" {
		needsUpdate = true
		topicConfig := topic.CreateConfig(retentionPeriodMs)
		topicSettings.SetConfig(*topicConfig)
	}

	if !needsUpdate {
		logger.Info(localizer.MustLocalizeFromID("kafka.topic.update.log.info.nothingToUpdate"))
		return nil
	}

	updateTopicReq = updateTopicReq.TopicSettings(*topicSettings)

	// update the topic
	response, httpRes, topicErr := updateTopicReq.Execute()
	// handle error
	if topicErr.Error() != "" {
		switch httpRes.StatusCode {
		case 404:
			return errors.New(localizer.MustLocalize(&localizer.Config{
				MessageID: "kafka.topic.common.error.notFoundError",
				TemplateData: map[string]interface{}{
					"TopicName":    opts.topicName,
					"InstanceName": kafkaInstance.GetName(),
				},
			}))
		case 401:
			return fmt.Errorf(localizer.MustLocalize(&localizer.Config{
				MessageID: "kafka.topic.common.error.unauthorized",
				TemplateData: map[string]interface{}{
					"Operation": "update",
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

	// the topic was updated, print it to stdout
	logger.Info(localizer.MustLocalize(&localizer.Config{
		MessageID: "kafka.topic.update.log.info.topicUpdated",
		TemplateData: map[string]interface{}{
			"TopicName":    opts.topicName,
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
