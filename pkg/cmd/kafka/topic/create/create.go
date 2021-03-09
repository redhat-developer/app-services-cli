package create

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/AlecAivazis/survey/v2"
	"github.com/bf2fc6cc711aee1a0c2a/cli/internal/localizer"

	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/connection"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/kafka/topic"

	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/cmd/flag"

	strimziadminclient "github.com/bf2fc6cc711aee1a0c2a/cli/pkg/api/strimzi-admin/client"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/dump"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/iostreams"
	"gopkg.in/yaml.v2"

	"github.com/bf2fc6cc711aee1a0c2a/cli/internal/config"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/cmd/factory"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/logging"

	"github.com/spf13/cobra"
)

const (
	Partitions = "partitions"
	Replicas   = "replicas"
)

var (
	partitionCount    int32
	retentionPeriodMs int
)

type Options struct {
	topicName      string
	partitionsStr  string
	retentionMsStr string
	kafkaID        string
	outputFormat   string
	interactive    bool

	IO         *iostreams.IOStreams
	Config     config.IConfig
	Connection factory.ConnectionFunc
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
		Use:     localizer.MustLocalizeFromID("kafka.topic.create.cmd.use"),
		Short:   localizer.MustLocalizeFromID("kafka.topic.create.cmd.shortDescription"),
		Long:    localizer.MustLocalizeFromID("kafka.topic.create.cmd.longDescription"),
		Example: localizer.MustLocalizeFromID("kafka.topic.create.cmd.example"),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			if !opts.IO.CanPrompt() && len(args) == 0 {
				return fmt.Errorf(localizer.MustLocalize(&localizer.Config{
					MessageID: "argument.error.requiredWhenNonInteractive",
					TemplateData: map[string]interface{}{
						"Argument": "Name",
					},
				}))
			} else if len(args) == 0 {
				opts.interactive = true
			}

			if !opts.interactive {
				opts.topicName = args[0]

				if err = topic.ValidateName(opts.topicName); err != nil {
					return err
				}

				if opts.partitionsStr != "" {
					partitionCount, err := topic.ConvertPartitionsToInt(opts.partitionsStr)
					if err != nil {
						return err
					}

					if err = topic.ValidatePartitionsN(partitionCount); err != nil {
						return err
					}
				} else {
					partitionCount = 1
				}

				if opts.retentionMsStr != "" {
					retentionPeriodMs, err = topic.ConvertRetentionMsToInt(opts.retentionMsStr)
					if err != nil {
						return err
					}

					if err = topic.ValidateMessageRetentionPeriod(retentionPeriodMs); err != nil {
						return err
					}
				} else {
					retentionPeriodMs = -1
				}
			}

			if err = flag.ValidateOutput(opts.outputFormat); err != nil {
				return err
			}

			if opts.retentionMsStr != "" {
				// convert the value from "--retention-ms" to int
				// nolint:govet
				retentionPeriodMs, err = topic.ConvertRetentionMsToInt(opts.retentionMsStr)
				if err != nil {
					return err
				}

				if err = topic.ValidateMessageRetentionPeriod(retentionPeriodMs); err != nil {
					return err
				}
			}

			if opts.partitionsStr != "" {
				partitionCount, err := topic.ConvertPartitionsToInt(opts.partitionsStr)
				if err != nil {
					return err
				}

				if err = topic.ValidatePartitionsN(partitionCount); err != nil {
					return err
				}
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
	cmd.Flags().StringVar(&opts.partitionsStr, "partitions", "", localizer.MustLocalizeFromID("kafka.topic.common.flag.partitions.description"))
	cmd.Flags().StringVar(&opts.retentionMsStr, "retention-ms", "", localizer.MustLocalizeFromID("kafka.topic.common.flag.retentionMs.description"))

	return cmd
}

func runCmd(opts *Options) error {

	if opts.interactive {
		// run the create command interactively
		err := runInteractivePrompt(opts)
		if err != nil {
			return err
		}
	}

	conn, err := opts.Connection(connection.DefaultConfigRequireMasAuth)
	if err != nil {
		return err
	}

	logger, err := opts.Logger()
	if err != nil {
		return err
	}

	ctx := context.Background()
	api, kafkaInstance, err := conn.API().TopicAdmin(opts.kafkaID)
	if err != nil {
		return err
	}

	createTopicReq := api.CreateTopic(ctx)

	var replicas int32 = 3
	topicInput := strimziadminclient.NewTopicInput{
		Name: opts.topicName,
		Settings: &strimziadminclient.TopicSettings{
			ReplicationFactor: &replicas,
			NumPartitions:     &partitionCount,
			Config:            topic.CreateConfig(retentionPeriodMs),
		},
	}
	createTopicReq = createTopicReq.NewTopicInput(topicInput)

	response, httpRes, topicErr := createTopicReq.Execute()
	if topicErr.Error() != "" {
		if httpRes == nil {
			return topicErr
		}

		switch httpRes.StatusCode {
		case 401:
			return fmt.Errorf(localizer.MustLocalize(&localizer.Config{
				MessageID: "kafka.topic.common.error.unauthorized",
				TemplateData: map[string]interface{}{
					"Operation": "create",
				},
			}))
		case 403:
			return errors.New(localizer.MustLocalize(&localizer.Config{
				MessageID: "kafka.topic.common.error.forbidden",
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
			return errors.New(localizer.MustLocalizeFromID("kafka.topic.common.error.internalServerError"))
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

func runInteractivePrompt(opts *Options) (err error) {

	_, err = opts.Connection()
	if err != nil {
		return err
	}

	logger, err := opts.Logger()
	if err != nil {
		return err
	}

	logger.Debug(localizer.MustLocalizeFromID("common.log.debug.startingInteractivePrompt"))

	promptName := &survey.Input{
		Message: localizer.MustLocalizeFromID("kafka.topic.common.input.name.message"),
		Help:    localizer.MustLocalizeFromID("kafka.topic.common.input.name.help"),
	}

	err = survey.AskOne(promptName, &opts.topicName, survey.WithValidator(survey.Required))
	if err != nil {
		return err
	}

	if err = topic.ValidateName(opts.topicName); err != nil {
		return err
	}

	if opts.partitionsStr == "" {
		logger.Debug(localizer.MustLocalizeFromID("kafka.topic.common.log.debug.interactive.partitionsNotSet"))

		partitionsPrompt := &survey.Input{
			Message: localizer.MustLocalizeFromID("kafka.topic.common.input.partitions.message"),
			Help:    localizer.MustLocalizeFromID("kafka.topic.common.input.partitions.help"),
		}

		err = survey.AskOne(partitionsPrompt, &opts.partitionsStr)
		if err != nil {
			return err
		}

		if opts.partitionsStr != "" {
			partitionCount, err := topic.ConvertPartitionsToInt(opts.partitionsStr)
			if err != nil {
				return err
			}

			if err = topic.ValidatePartitionsN(partitionCount); err != nil {
				return err
			}
		}
	}

	if opts.retentionMsStr == "" {
		logger.Debug(localizer.MustLocalizeFromID("kafka.topic.common.log.debug.interactive.retentionMsNotSet"))

		retentionPrompt := &survey.Input{
			Message: localizer.MustLocalizeFromID("kafka.topic.common.input.retentionMs.message"),
			Help:    localizer.MustLocalizeFromID("kafka.topic.common.input.retentionMs.help"),
		}

		err = survey.AskOne(retentionPrompt, &opts.retentionMsStr)
		if err != nil {
			return err
		}

		if opts.retentionMsStr != "" {
			// convert the value from "--retention-ms" to int
			// nolint:govet
			retentionPeriodMs, err = topic.ConvertRetentionMsToInt(opts.retentionMsStr)
			if err != nil {
				return err
			}

			if err = topic.ValidateMessageRetentionPeriod(retentionPeriodMs); err != nil {
				return err
			}
		}
	}

	return nil
}
