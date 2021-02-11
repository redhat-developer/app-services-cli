package update

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/color"

	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/cmdutil"

	"github.com/MakeNowJust/heredoc"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/kafka/topic"

	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/cmd/flag"
	flagutil "github.com/bf2fc6cc711aee1a0c2a/cli/pkg/cmdutil/flags"

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

	cmd := &cobra.Command{
		Use:   "update",
		Short: "Update a Kafka topic",
		Long:  "Update a topic in the current Kafka instance",
		Example: heredoc.Doc(`
			# update the number of replicas for a topic
			$ rhoas kafka topic update topic-1 --replication-factor 3
		`),
		Args: cobra.ExactArgs(1),
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

			if opts.retentionMsStr == "" && opts.partitionsStr == "" && opts.replicaCountStr == "" {
				return fmt.Errorf(`nothing to update`)
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
				return fmt.Errorf("No Kafka instance selected. Use the '--id' flag or set one in context with the 'use' command")
			}

			opts.kafkaID = cfg.Services.Kafka.ClusterID

			return runCmd(opts)
		},
	}

	fs := cmd.Flags()
	flag.AddOutput(fs, &opts.outputFormat, "json", flagutil.ValidOutputFormats)
	cmd.Flags().StringVar(&opts.partitionsStr, "partitions", "", "The number of partitions in the topic")
	cmd.Flags().StringVar(&opts.retentionMsStr, "retention-ms", "", "The period of time in milliseconds the broker will retain a partition log before deleting it")
	cmd.Flags().StringVar(&opts.replicaCountStr, "replicas", "", "The replication factor for the topic")

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
		return fmt.Errorf("Kafka instance with ID '%v' not found", opts.kafkaID)
	} else if apiErr.Error() != "" {
		return apiErr
	}

	// track if any values have changed
	var needsUpdate bool

	topicToUpdate, httpRes, _ := api.TopicAdmin(opts.kafkaID).GetTopic(ctx, opts.topicName).Execute()
	if httpRes.StatusCode == 404 {
		return fmt.Errorf("topic '%v' not found in Kafka instance '%v'", opts.topicName, kafkaInstance.GetName())
	}

	currentPartitionCount := len(topicToUpdate.GetPartitions())

	updateTopicReq := api.TopicAdmin(opts.kafkaID).UpdateTopic(ctx, opts.topicName)

	topicSettings := &strimziadminclient.TopicSettings{}

	// Only set partitions if the flag was set
	if opts.partitionsStr != "" {
		if int(partitionCount) < currentPartitionCount {
			return fmt.Errorf("number of topic partitions cannot be decreased from %v to %v", currentPartitionCount, partitionCount)
		}
		if int(partitionCount) == currentPartitionCount {
			logger.Infof("The number of partitions set (%v) is the same as the current number of partitions", partitionCount)
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
		logger.Info("No topic values have been changed, nothing to update")
		return nil
	}

	updateTopicReq = updateTopicReq.TopicSettings(*topicSettings)

	// update the topic
	response, httpRes, topicErr := updateTopicReq.Execute()
	// handle error
	if topicErr.Error() != "" {
		switch httpRes.StatusCode {
		case 401:
			return fmt.Errorf("you are unauthorized to update this topic")
		case 404:
			return fmt.Errorf("topic '%v' not found in Kafka instance '%v'", opts.topicName, kafkaInstance.GetName())
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

	// the topic was updated, print it to stdout
	logger.Infof("Topic %v updated in Kafka instance %v:", color.Info(response.GetName()), color.Info((kafkaInstance.GetName())))
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
