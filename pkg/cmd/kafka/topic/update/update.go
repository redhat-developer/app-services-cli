package update

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strings"

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
	name              string
	partitionsStr     string
	retentionMsStr    string
	retentionBytesStr string
	kafkaID           string
	outputFormat      string
	interactive       bool
	cleanupPolicy     string

	IO         *iostreams.IOStreams
	Config     config.IConfig
	Connection factory.ConnectionFunc
	Logger     logging.Logger
	localizer  localize.Localizer
	Context    context.Context
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
		Context:    f.Context,
	}

	cmd := &cobra.Command{
		Use:     "update",
		Short:   opts.localizer.MustLocalize("kafka.topic.update.cmd.shortDescription"),
		Long:    opts.localizer.MustLocalize("kafka.topic.update.cmd.longDescription"),
		Example: opts.localizer.MustLocalize("kafka.topic.update.cmd.example"),
		Args:    cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			validator := topicutil.Validator{
				Localizer: opts.localizer,
			}

			if !opts.IO.CanPrompt() && opts.retentionMsStr == "" && opts.partitionsStr == "" && opts.retentionBytesStr == "" {
				return errors.New(opts.localizer.MustLocalize("argument.error.requiredWhenNonInteractive", localize.NewEntry("Argument", "name")))
			} else if opts.retentionMsStr == "" && opts.partitionsStr == "" && opts.retentionBytesStr == "" && opts.cleanupPolicy == "" {
				opts.interactive = true
			}

			if !opts.interactive {
				if opts.retentionMsStr == "" && opts.partitionsStr == "" && opts.retentionBytesStr == "" && opts.cleanupPolicy == "" {
					opts.Logger.Info(opts.localizer.MustLocalize("kafka.topic.update.log.info.nothingToUpdate"))
					return nil
				}

				if err = validator.ValidateName(opts.name); err != nil {
					return err
				}

				// check that a valid --cleanup-policy flag value is used
				if opts.cleanupPolicy != "" {
					validPolicy := flagutil.IsValidInput(opts.cleanupPolicy, topicutil.ValidCleanupPolicies...)
					if !validPolicy {
						return flag.InvalidValueError("cleanup-policy", opts.cleanupPolicy, topicutil.ValidCleanupPolicies...)
					}
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
				return opts.localizer.MustLocalizeError("kafka.topic.common.error.noKafkaSelected")
			}

			opts.kafkaID = cfg.Services.Kafka.ClusterID

			return runCmd(opts)
		},
	}

	cmd.Flags().StringVarP(&opts.outputFormat, "output", "o", "json", opts.localizer.MustLocalize("kafka.topic.common.flag.output.description"))
	cmd.Flags().StringVar(&opts.retentionMsStr, "retention-ms", "", opts.localizer.MustLocalize("kafka.topic.common.input.retentionMs.description"))
	cmd.Flags().StringVar(&opts.retentionBytesStr, "retention-bytes", "", opts.localizer.MustLocalize("kafka.topic.common.input.retentionBytes.description"))
	cmd.Flags().StringVar(&opts.cleanupPolicy, "cleanup-policy", "", opts.localizer.MustLocalize("kafka.topic.common.input.cleanupPolicy.description"))
	cmd.Flags().StringVar(&opts.partitionsStr, "partitions", "", opts.localizer.MustLocalize("kafka.topic.common.input.partitions.description"))

	cmd.Flags().StringVar(&opts.name, "name", "", opts.localizer.MustLocalize("kafka.topic.common.flag.name.description"))
	_ = cmd.RegisterFlagCompletionFunc("name", func(cmd *cobra.Command, _ []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return cmdutil.FilterValidTopicNameArgs(f, toComplete)
	})
	_ = cmd.MarkFlagRequired("name")

	flagutil.EnableOutputFlagCompletion(cmd)

	flagutil.EnableStaticFlagCompletion(cmd, "cleanup-policy", topicutil.ValidCleanupPolicies)

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

		if opts.partitionsStr != "" {
			partitionCount, err = topicutil.ConvertPartitionsToInt(opts.partitionsStr)
			if err != nil {
				return err
			}
		}

	}

	conn, err := opts.Connection(connection.DefaultConfigRequireMasAuth)
	if err != nil {
		return err
	}

	api, kafkaInstance, err := conn.API().KafkaAdmin(opts.kafkaID)
	if err != nil {
		return err
	}

	// track if any values have changed
	var needsUpdate bool

	topic, httpRes, err := api.TopicsApi.GetTopic(opts.Context, opts.name).Execute()
	if httpRes != nil {
		defer httpRes.Body.Close()
	}

	topicNameTmplPair := localize.NewEntry("TopicName", opts.name)
	kafkaNameTmplPair := localize.NewEntry("InstanceName", kafkaInstance.GetName())
	if err != nil {
		if httpRes == nil {
			return err
		}
		if httpRes.StatusCode == http.StatusNotFound {
			return errors.New(opts.localizer.MustLocalize("kafka.topic.common.error.topicNotFoundError", topicNameTmplPair, kafkaNameTmplPair))
		}
	}

	// map to store the config entries which will be updated
	configEntryMap := map[string]*string{}

	updateTopicReq := api.TopicsApi.UpdateTopic(opts.Context, opts.name)

	topicSettings := &kafkainstanceclient.UpdateTopicInput{}

	if opts.retentionMsStr != "" {
		needsUpdate = true
		configEntryMap[topicutil.RetentionMsKey] = &opts.retentionMsStr
	}

	if opts.retentionBytesStr != "" {
		needsUpdate = true
		configEntryMap[topicutil.RetentionSizeKey] = &opts.retentionBytesStr
	}

	if opts.cleanupPolicy != "" && strings.Compare(opts.cleanupPolicy, topicutil.GetConfigValue(topic.GetConfig(), topicutil.CleanupPolicy)) != 0 {
		needsUpdate = true
		configEntryMap[topicutil.CleanupPolicy] = &opts.cleanupPolicy
	}

	if opts.partitionsStr != "" {
		needsUpdate = true
		topicSettings.SetNumPartitions(partitionCount)
	}

	if !needsUpdate {
		opts.Logger.Info(opts.localizer.MustLocalize("kafka.topic.update.log.info.nothingToUpdate"))
		return nil
	}

	if len(configEntryMap) > 0 {
		configEntries := topicutil.CreateConfigEntries(configEntryMap)
		topicSettings.SetConfig(*configEntries)
	}

	updateTopicReq = updateTopicReq.UpdateTopicInput(*topicSettings)

	// update the topic
	response, httpRes, err := updateTopicReq.Execute()
	if httpRes != nil {
		defer httpRes.Body.Close()
	}

	// handle error
	if err != nil {
		if httpRes == nil {
			return err
		}

		operationTmplPair := localize.NewEntry("Operation", "update")
		switch httpRes.StatusCode {
		case http.StatusNotFound:
			return errors.New(opts.localizer.MustLocalize("kafka.topic.common.error.notFoundError", topicNameTmplPair, kafkaNameTmplPair))
		case http.StatusUnauthorized:
			return errors.New(opts.localizer.MustLocalize("kafka.topic.common.error.unauthorized", operationTmplPair))
		case http.StatusForbidden:
			return errors.New(opts.localizer.MustLocalize("kafka.topic.common.error.forbidden", operationTmplPair))
		case http.StatusInternalServerError:
			return errors.New(opts.localizer.MustLocalize("kafka.topic.common.error.internalServerError"))
		case http.StatusServiceUnavailable:
			return errors.New(opts.localizer.MustLocalize("kafka.topic.common.error.unableToConnectToKafka", localize.NewEntry("Name", kafkaInstance.GetName())))
		default:
			return err
		}
	}

	opts.Logger.Info(opts.localizer.MustLocalize("kafka.topic.update.log.info.topicUpdated", topicNameTmplPair, kafkaNameTmplPair))

	switch opts.outputFormat {
	case dump.JSONFormat:
		data, _ := json.Marshal(response)
		_ = dump.JSON(opts.IO.Out, data)
	case dump.YAMLFormat, dump.YMLFormat:
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
	topic, httpRes, err := api.TopicsApi.GetTopic(opts.Context, opts.name).Execute()
	if httpRes != nil {
		defer httpRes.Body.Close()
	}

	topicNameTmplPair := localize.NewEntry("TopicName", opts.name)
	kafkaNameTmplPair := localize.NewEntry("InstanceName", kafkaInstance.GetName())
	if err != nil {
		if httpRes == nil {
			return err
		}
		if httpRes.StatusCode == http.StatusNotFound {
			return errors.New(opts.localizer.MustLocalize("kafka.topic.common.error.topicNotFoundError", topicNameTmplPair, kafkaNameTmplPair))
		}
	}

	validator := topicutil.Validator{
		Localizer: opts.localizer,
	}

	opts.Logger.Debug(opts.localizer.MustLocalize("common.log.debug.startingInteractivePrompt"))

	partitionsPrompt := &survey.Input{
		Message: opts.localizer.MustLocalize("kafka.topic.update.input.partitions.message"),
		Help:    opts.localizer.MustLocalize("kafka.topic.update.input.partitions.help"),
	}

	validator.CurPartitions = len(*topic.Partitions)

	err = survey.AskOne(partitionsPrompt, &opts.partitionsStr, survey.WithValidator(validator.ValidatePartitionsN))
	if err != nil {
		return err
	}

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

	cleanupPolicyPrompt := &survey.Select{
		Message: opts.localizer.MustLocalize("kafka.topic.update.input.cleanupPolicy.message"),
		Help:    opts.localizer.MustLocalize("kafka.topic.update.input.cleanupPolicy.help"),
		Options: topicutil.ValidCleanupPolicies,
		Default: topicutil.GetConfigValue(topic.GetConfig(), topicutil.CleanupPolicy),
	}

	err = survey.AskOne(cleanupPolicyPrompt, &opts.cleanupPolicy)
	if err != nil {
		return err
	}

	return nil
}
