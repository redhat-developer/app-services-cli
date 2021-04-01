package create

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"

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
	defaultRetentionPeriodMS = 604800000
)

type Options struct {
	topicName    string
	partitions   int32
	retentionMs  int
	kafkaID      string
	outputFormat string
	interactive  bool

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
		Args:    cobra.RangeArgs(0, 1),
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

			if err = flag.ValidateOutput(opts.outputFormat); err != nil {
				return err
			}

			if !opts.interactive {
				opts.topicName = args[0]

				if err = topic.ValidateName(opts.topicName); err != nil {
					return err
				}

				if err = topic.ValidatePartitionsN(opts.partitions); err != nil {
					return err
				}

				if err = topic.ValidateMessageRetentionPeriod(opts.retentionMs); err != nil {
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
	cmd.Flags().Int32Var(&opts.partitions, "partitions", 1, localizer.MustLocalizeFromID("kafka.topic.common.input.partitions.description"))
	cmd.Flags().IntVar(&opts.retentionMs, "retention-ms", defaultRetentionPeriodMS, localizer.MustLocalizeFromID("kafka.topic.common.input.retentionMs.description"))

	return cmd
}

// nolint:funlen
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

	topicInput := strimziadminclient.NewTopicInput{
		Name: opts.topicName,
		Settings: strimziadminclient.TopicSettings{
			NumPartitions: opts.partitions,
			Config:        topic.CreateConfig(opts.retentionMs),
		},
	}
	createTopicReq = createTopicReq.NewTopicInput(topicInput)

	response, httpRes, topicErr := createTopicReq.Execute()
	bodyBytes, err := ioutil.ReadAll(httpRes.Body)
	if err != nil {
		logger.Debug("Could not read response body")
	} else {
		logger.Debug("Response Body:", string(bodyBytes))
	}
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

func runInteractivePrompt(opts *Options) (err error) {

	conn, err := opts.Connection(connection.DefaultConfigRequireMasAuth)
	if err != nil {
		return err
	}

	api, kafkaInstance, err := conn.API().TopicAdmin(opts.kafkaID)
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

	err = survey.AskOne(
		promptName,
		&opts.topicName,
		survey.WithValidator(survey.Required),
		survey.WithValidator(topic.ValidateName),
		survey.WithValidator(topic.ValidateNameIsAvailable(api, kafkaInstance.GetName())),
	)

	if err != nil {
		return err
	}

	partitionsPrompt := &survey.Input{
		Message: localizer.MustLocalizeFromID("kafka.topic.create.input.partitions.message"),
		Help:    localizer.MustLocalizeFromID("kafka.topic.common.input.partitions.description"),
		Default: "1",
	}

	err = survey.AskOne(partitionsPrompt, &opts.partitions, survey.WithValidator(topic.ValidatePartitionsN))
	if err != nil {
		return err
	}

	retentionPrompt := &survey.Input{
		Message: localizer.MustLocalizeFromID("kafka.topic.create.input.retentionMs.message"),
		Help:    localizer.MustLocalizeFromID("kafka.topic.common.input.retentionMs.description"),
		Default: fmt.Sprintf("%v", defaultRetentionPeriodMS),
	}

	err = survey.AskOne(retentionPrompt, &opts.retentionMs, survey.WithValidator(topic.ValidateMessageRetentionPeriod))
	if err != nil {
		return err
	}

	return nil
}
