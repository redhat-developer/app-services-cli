package update

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/AlecAivazis/survey/v2"
	"github.com/redhat-developer/app-services-cli/internal/localizer"

	"github.com/redhat-developer/app-services-cli/pkg/cmdutil"
	"github.com/redhat-developer/app-services-cli/pkg/connection"

	topicutil "github.com/redhat-developer/app-services-cli/pkg/kafka/topic"

	"github.com/redhat-developer/app-services-cli/pkg/cmd/flag"

	strimziadminclient "github.com/redhat-developer/app-services-cli/pkg/api/strimzi-admin/client"
	"github.com/redhat-developer/app-services-cli/pkg/dump"
	"github.com/redhat-developer/app-services-cli/pkg/iostreams"
	"gopkg.in/yaml.v2"

	"github.com/redhat-developer/app-services-cli/internal/config"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/factory"
	"github.com/redhat-developer/app-services-cli/pkg/logging"

	"github.com/spf13/cobra"
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
		Use:     localizer.MustLocalizeFromID("kafka.topic.update.cmd.use"),
		Short:   localizer.MustLocalizeFromID("kafka.topic.update.cmd.shortDescription"),
		Long:    localizer.MustLocalizeFromID("kafka.topic.update.cmd.longDescription"),
		Example: localizer.MustLocalizeFromID("kafka.topic.update.cmd.example"),
		Args:    cobra.ExactArgs(1),
		// Dynamic completion of the topic name
		ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
			validNames := []string{}

			var searchName string = args[0]

			cfg, err := opts.Config.Load()
			if err != nil {
				return validNames, cobra.ShellCompDirectiveError
			}

			if !cfg.HasKafka() {
				return validNames, cobra.ShellCompDirectiveError
			}

			return cmdutil.FilterValidTopicNameArgs(f, cfg.Services.Kafka.ClusterID, searchName)
		},
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			if !opts.IO.CanPrompt() && opts.retentionMsStr == "" && opts.partitionsStr == "" {
				return fmt.Errorf(localizer.MustLocalize(&localizer.Config{
					MessageID: "argument.error.requiredWhenNonInteractive",
					TemplateData: map[string]interface{}{
						"Argument": "Name",
					},
				}))
			} else if opts.retentionMsStr == "" && opts.partitionsStr == "" {
				opts.interactive = true
			}

			opts.topicName = args[0]

			if !opts.interactive {

				// nolint:govet
				logger, err := opts.Logger()
				if err != nil {
					return err
				}

				if opts.retentionMsStr == "" && opts.partitionsStr == "" {
					logger.Info(localizer.MustLocalizeFromID("kafka.topic.update.log.info.nothingToUpdate"))
					return nil
				}

				if err = topicutil.ValidateName(opts.topicName); err != nil {
					return err
				}
			}

			if err = flag.ValidateOutput(opts.outputFormat); err != nil {
				return err
			}

			// check if the partition flag is set
			if opts.partitionsStr != "" {
				// nolint:govet
				partitionCount, err = topicutil.ConvertPartitionsToInt(opts.partitionsStr)
				if err != nil {
					return err
				}

				if err = topicutil.ValidatePartitionsN(partitionCount); err != nil {
					return err
				}
			}

			if opts.retentionMsStr != "" {
				retentionPeriodMs, err = topicutil.ConvertRetentionMsToInt(opts.retentionMsStr)
				if err != nil {
					return err
				}

				if err = topicutil.ValidateMessageRetentionPeriod(retentionPeriodMs); err != nil {
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
	cmd.Flags().StringVar(&opts.retentionMsStr, "retention-ms", "", localizer.MustLocalizeFromID("kafka.topic.common.input.retentionMs.description"))

	return cmd
}

// nolint:funlen
func runCmd(opts *Options) error {

	if opts.interactive {
		// run the update command interactively
		err := runInteractivePrompt(opts)
		if err != nil {
			return err
		}

		if opts.retentionMsStr != "" {
			retentionPeriodMs, err = topicutil.ConvertRetentionMsToInt(opts.retentionMsStr)
			if err != nil {
				return err
			}
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
	api, kafkaInstance, err := conn.API().TopicAdmin(opts.kafkaID)
	if err != nil {
		return err
	}

	// track if any values have changed
	var needsUpdate bool

	_, httpRes, topicErr := api.GetTopic(context.Background(), opts.topicName).Execute()

	if topicErr.Error() != "" {
		if httpRes == nil {
			return topicErr
		}
		if httpRes.StatusCode == 404 {
			return errors.New(localizer.MustLocalize(&localizer.Config{
				MessageID: "kafka.topic.common.error.topicNotFoundError",
				TemplateData: map[string]interface{}{
					"TopicName":    opts.topicName,
					"InstanceName": kafkaInstance.GetName(),
				},
			}))
		}
	}

	// map to store the config entries which will be updated
	var configEntryMap map[string]*string = map[string]*string{}

	updateTopicReq := api.UpdateTopic(context.Background(), opts.topicName)

	topicSettings := &strimziadminclient.UpdateTopicInput{}

	if opts.retentionMsStr != "" {
		needsUpdate = true
		configEntryMap[topicutil.RetentionMsKey] = &opts.retentionMsStr
	}

	if !needsUpdate {
		logger.Info(localizer.MustLocalizeFromID("kafka.topic.update.log.info.nothingToUpdate"))
		return nil
	}

	if len(configEntryMap) > 0 {
		configEntries := topicutil.CreateConfigEntries(configEntryMap)
		topicSettings.SetConfig(*configEntries)
	}

	updateTopicReq = updateTopicReq.UpdateTopicInput(*topicSettings)

	// update the topic
	response, httpRes, topicErr := updateTopicReq.Execute()
	// handle error
	if topicErr.Error() != "" {
		if httpRes == nil {
			return topicErr
		}

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
		case 403:
			return errors.New(localizer.MustLocalize(&localizer.Config{
				MessageID: "kafka.topic.common.error.forbidden",
				TemplateData: map[string]interface{}{
					"Operation": "update",
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

func runInteractivePrompt(opts *Options) (err error) {

	_, err = opts.Connection(connection.DefaultConfigRequireMasAuth)
	if err != nil {
		return err
	}

	logger, err := opts.Logger()
	if err != nil {
		return err
	}

	logger.Debug(localizer.MustLocalizeFromID("common.log.debug.startingInteractivePrompt"))

	retentionPrompt := &survey.Input{
		Message: localizer.MustLocalizeFromID("kafka.topic.update.input.retentionMs.message"),
		Help:    localizer.MustLocalizeFromID("kafka.topic.update.input.retentionMs.help"),
	}

	err = survey.AskOne(retentionPrompt, &opts.retentionMsStr, survey.WithValidator(topicutil.ValidateMessageRetentionPeriod))
	if err != nil {
		return err
	}

	return nil
}
