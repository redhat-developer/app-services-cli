package update

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/AlecAivazis/survey/v2"

	"github.com/redhat-developer/app-services-cli/pkg/cmdutil"
	"github.com/redhat-developer/app-services-cli/pkg/connection"
	"github.com/redhat-developer/app-services-cli/pkg/localize"

	flagutil "github.com/redhat-developer/app-services-cli/pkg/cmdutil/flags"
	topicutil "github.com/redhat-developer/app-services-cli/pkg/kafka/topic"

	"github.com/redhat-developer/app-services-cli/pkg/cmd/flag"

	"github.com/redhat-developer/app-services-cli/pkg/dump"
	"github.com/redhat-developer/app-services-cli/pkg/iostreams"
	kafkainstanceclient "github.com/redhat-developer/app-services-sdk-go/kafkainstance/apiv1internal/client"
	"gopkg.in/yaml.v2"

	"github.com/redhat-developer/app-services-cli/internal/config"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/factory"
	"github.com/redhat-developer/app-services-cli/pkg/logging"

	"github.com/spf13/cobra"
)

var (
	partitionCount     int32
	retentionPeriodMs  int
	retentionSizeBytes int
)

type Options struct {
	topicName         string
	partitionsStr     string
	retentionMsStr    string
	retentionBytesStr string
	kafkaID           string
	outputFormat      string
	interactive       bool

	IO         *iostreams.IOStreams
	Config     config.IConfig
	Connection factory.ConnectionFunc
	Logger     func() (logging.Logger, error)
	localizer  localize.Localizer
}

// NewUpdateTopicCommand gets a new command for updating a kafka topic.
// nolint:funlen
func NewUpdateTopicCommand(f *factory.Factory) *cobra.Command {
	opts := &Options{
		Connection: f.Connection,
		Config:     f.Config,
		Logger:     f.Logger,
		IO:         f.IOStreams,
		localizer:  f.Localizer,
	}

	cmd := &cobra.Command{
		Use:     opts.localizer.MustLocalize("kafka.topic.update.cmd.use"),
		Short:   opts.localizer.MustLocalize("kafka.topic.update.cmd.shortDescription"),
		Long:    opts.localizer.MustLocalize("kafka.topic.update.cmd.longDescription"),
		Example: opts.localizer.MustLocalize("kafka.topic.update.cmd.example"),
		Args:    cobra.ExactValidArgs(1),
		// Dynamic completion of the topic name
		ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
			return cmdutil.FilterValidTopicNameArgs(f, toComplete)
		},
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			validator := &topicutil.Validator{
				Localizer: opts.localizer,
			}

			if !opts.IO.CanPrompt() && opts.retentionMsStr == "" && opts.partitionsStr == "" && opts.retentionBytesStr == "" {
				return errors.New(opts.localizer.MustLocalize("argument.error.requiredWhenNonInteractive", localize.NewEntry("Argument", "name")))
			} else if opts.retentionMsStr == "" && opts.partitionsStr == "" && opts.retentionBytesStr == "" {
				opts.interactive = true
			}

			opts.topicName = args[0]

			if !opts.interactive {

				// nolint:govet
				logger, err := opts.Logger()
				if err != nil {
					return err
				}

				if opts.retentionMsStr == "" && opts.partitionsStr == "" && opts.retentionBytesStr == "" {
					logger.Info(opts.localizer.MustLocalize("kafka.topic.update.log.info.nothingToUpdate"))
					return nil
				}

				if err = validator.ValidateName(opts.topicName); err != nil {
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

				if err = validator.ValidatePartitionsN(partitionCount); err != nil {
					return err
				}
			}

			if opts.retentionMsStr != "" {
				retentionPeriodMs, err = topicutil.ConvertRetentionMsToInt(opts.retentionMsStr)
				if err != nil {
					return err
				}

				if err = validator.ValidateMessageRetentionPeriod(retentionPeriodMs); err != nil {
					return err
				}
			}

			if opts.retentionBytesStr != "" {
				retentionSizeBytes, err = topicutil.ConvertRetentionBytesToInt(opts.retentionBytesStr)
				if err != nil {
					return err
				}

				if err = validator.ValidateMessageRetentionSize(retentionSizeBytes); err != nil {
					return err
				}
			}

			cfg, err := opts.Config.Load()
			if err != nil {
				return err
			}

			if !cfg.HasKafka() {
				return fmt.Errorf(opts.localizer.MustLocalize("kafka.topic.common.error.noKafkaSelected"))
			}

			opts.kafkaID = cfg.Services.Kafka.ClusterID

			return runCmd(opts)
		},
	}

	cmd.Flags().StringVarP(&opts.outputFormat, "output", "o", "json", opts.localizer.MustLocalize("kafka.topic.common.flag.output.description"))
	cmd.Flags().StringVar(&opts.retentionMsStr, "retention-ms", "", opts.localizer.MustLocalize("kafka.topic.common.input.retentionMs.description"))
	cmd.Flags().StringVar(&opts.retentionBytesStr, "retention-bytes", "", opts.localizer.MustLocalize("kafka.topic.common.input.retentionBytes.description"))

	flagutil.EnableOutputFlagCompletion(cmd)

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

		if opts.retentionBytesStr != "" {
			retentionSizeBytes, err = topicutil.ConvertRetentionBytesToInt(opts.retentionBytesStr)
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
	api, kafkaInstance, err := conn.API().KafkaAdmin(opts.kafkaID)
	if err != nil {
		return err
	}

	// track if any values have changed
	var needsUpdate bool

	_, httpRes, err := api.GetTopic(context.Background(), opts.topicName).Execute()

	topicNameTmplPair := localize.NewEntry("TopicName", opts.topicName)
	kafkaNameTmplPair := localize.NewEntry("InstanceName", kafkaInstance.GetName())
	if err != nil {
		if httpRes == nil {
			return err
		}
		if httpRes.StatusCode == 404 {
			return errors.New(opts.localizer.MustLocalize("kafka.topic.common.error.topicNotFoundError", topicNameTmplPair, kafkaNameTmplPair))
		}
	}

	// map to store the config entries which will be updated
	configEntryMap := map[string]*string{}

	updateTopicReq := api.UpdateTopic(context.Background(), opts.topicName)

	topicSettings := &kafkainstanceclient.UpdateTopicInput{}

	if opts.retentionMsStr != "" {
		needsUpdate = true
		configEntryMap[topicutil.RetentionMsKey] = &opts.retentionMsStr
	}

	if opts.retentionBytesStr != "" {
		needsUpdate = true
		configEntryMap[topicutil.RetentionSizeKey] = &opts.retentionBytesStr
	}

	if !needsUpdate {
		logger.Info(opts.localizer.MustLocalize("kafka.topic.update.log.info.nothingToUpdate"))
		return nil
	}

	if len(configEntryMap) > 0 {
		configEntries := topicutil.CreateConfigEntries(configEntryMap)
		topicSettings.SetConfig(*configEntries)
	}

	updateTopicReq = updateTopicReq.UpdateTopicInput(*topicSettings)

	// update the topic
	response, httpRes, err := updateTopicReq.Execute()
	// handle error
	if err != nil {
		if httpRes == nil {
			return err
		}

		operationTmplPair := localize.NewEntry("Operation", "update")
		switch httpRes.StatusCode {
		case 404:
			return errors.New(opts.localizer.MustLocalize("kafka.topic.common.error.notFoundError", topicNameTmplPair, kafkaNameTmplPair))
		case 401:
			return errors.New(opts.localizer.MustLocalize("kafka.topic.common.error.unauthorized", operationTmplPair))
		case 403:
			return errors.New(opts.localizer.MustLocalize("kafka.topic.common.error.forbidden", operationTmplPair))
		case 500:
			return errors.New(opts.localizer.MustLocalize("kafka.topic.common.error.internalServerError"))
		case 503:
			return errors.New(opts.localizer.MustLocalize("kafka.topic.common.error.unableToConnectToKafka", localize.NewEntry("Name", kafkaInstance.GetName())))
		default:
			return err
		}
	}

	logger.Info(opts.localizer.MustLocalize("kafka.topic.update.log.info.topicUpdated", topicNameTmplPair, kafkaNameTmplPair))

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

	api, kafkaInstance, err := conn.API().KafkaAdmin(opts.kafkaID)
	if err != nil {
		return err
	}

	// check if topic exists
	_, httpRes, err := api.GetTopic(context.Background(), opts.topicName).
		Execute()

	topicNameTmplPair := localize.NewEntry("TopicName", opts.topicName)
	kafkaNameTmplPair := localize.NewEntry("InstanceName", kafkaInstance.GetName())
	if err != nil {
		if httpRes == nil {
			return err
		}
		if httpRes.StatusCode == 404 {
			return errors.New(opts.localizer.MustLocalize("kafka.topic.common.error.topicNotFoundError", topicNameTmplPair, kafkaNameTmplPair))
		}
	}

	logger, err := opts.Logger()
	if err != nil {
		return err
	}

	validator := &topicutil.Validator{
		Localizer: opts.localizer,
	}

	logger.Debug(opts.localizer.MustLocalize("common.log.debug.startingInteractivePrompt"))

	retentionMsPrompt := &survey.Input{
		Message: opts.localizer.MustLocalize("kafka.topic.update.input.retentionMs.message"),
		Help:    opts.localizer.MustLocalize("kafka.topic.update.input.retentionMs.help"),
	}

	err = survey.AskOne(retentionMsPrompt, &opts.retentionMsStr, survey.WithValidator(validator.ValidateMessageRetentionPeriod))
	if err != nil {
		return err
	}

	retentionBytesPrompt := &survey.Input{
		Message: opts.localizer.MustLocalize("kafka.topic.update.input.retentionBytes.message"),
		Help:    opts.localizer.MustLocalize("kafka.topic.update.input.retentionBytes.help"),
	}

	err = survey.AskOne(retentionBytesPrompt, &opts.retentionBytesStr, survey.WithValidator(validator.ValidateMessageRetentionSize))
	if err != nil {
		return err
	}

	return nil
}
