package update

import (
	"context"
	"net/http"
	"strings"

	"github.com/redhat-developer/app-services-cli/pkg/cmd/kafka/kafkacmdutil"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/kafka/topic/topiccmdutil"

	"github.com/AlecAivazis/survey/v2"
	"github.com/redhat-developer/app-services-cli/pkg/core/cmdutil/flagutil"
	"github.com/redhat-developer/app-services-cli/pkg/core/config"
	"github.com/redhat-developer/app-services-cli/pkg/core/ioutil/iostreams"
	"github.com/redhat-developer/app-services-cli/pkg/core/localize"
	"github.com/redhat-developer/app-services-cli/pkg/core/logging"
	"github.com/redhat-developer/app-services-cli/pkg/shared/connection"
	"github.com/redhat-developer/app-services-cli/pkg/shared/factory"
	kafkainstanceclient "github.com/redhat-developer/app-services-sdk-go/kafkainstance/apiv1internal/client"

	"github.com/spf13/cobra"
)

var (
	partitionCount     int32
	retentionPeriodMs  int
	retentionSizeBytes int
)

type options struct {
	name              string
	partitionsStr     string
	retentionMsStr    string
	retentionBytesStr string
	kafkaID           string
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
	opts := &options{
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
			validator := topiccmdutil.Validator{
				Localizer: opts.localizer,
			}

			if !opts.IO.CanPrompt() && opts.retentionMsStr == "" && opts.partitionsStr == "" && opts.retentionBytesStr == "" {
				return opts.localizer.MustLocalizeError("argument.error.requiredWhenNonInteractive", localize.NewEntry("Argument", "name"))
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
					validPolicy := flagutil.IsValidInput(opts.cleanupPolicy, topiccmdutil.ValidCleanupPolicies...)
					if !validPolicy {
						return flagutil.InvalidValueError("cleanup-policy", opts.cleanupPolicy, topiccmdutil.ValidCleanupPolicies...)
					}
				}

			}

			// check if the partition flag is set
			if opts.partitionsStr != "" {
				// nolint:govet
				partitionCount, err = topiccmdutil.ConvertPartitionsToInt(opts.partitionsStr)
				if err != nil {
					return err
				}

				if err = validator.ValidatePartitionsN(partitionCount); err != nil {
					return err
				}
			}

			if opts.retentionMsStr != "" {
				retentionPeriodMs, err = topiccmdutil.ConvertRetentionMsToInt(opts.retentionMsStr)
				if err != nil {
					return err
				}

				if err = validator.ValidateMessageRetentionPeriod(retentionPeriodMs); err != nil {
					return err
				}
			}

			if opts.retentionBytesStr != "" {
				retentionSizeBytes, err = topiccmdutil.ConvertRetentionBytesToInt(opts.retentionBytesStr)
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

			instanceID, ok := cfg.GetKafkaIdOk()
			if !ok {
				return opts.localizer.MustLocalizeError("kafka.topic.common.error.noKafkaSelected")
			}

			opts.kafkaID = instanceID

			return runCmd(opts)
		},
	}

	flags := flagutil.NewFlagSet(cmd, opts.localizer)

	flags.StringVar(&opts.retentionMsStr, "retention-ms", "", opts.localizer.MustLocalize("kafka.topic.common.input.retentionMs.description"))
	flags.StringVar(&opts.retentionBytesStr, "retention-bytes", "", opts.localizer.MustLocalize("kafka.topic.common.input.retentionBytes.description"))
	flags.StringVar(&opts.cleanupPolicy, "cleanup-policy", "", opts.localizer.MustLocalize("kafka.topic.common.input.cleanupPolicy.description"))
	flags.StringVar(&opts.partitionsStr, "partitions", "", opts.localizer.MustLocalize("kafka.topic.common.input.partitions.description"))

	flags.StringVar(&opts.name, "name", "", opts.localizer.MustLocalize("kafka.topic.common.flag.name.description"))
	_ = cmd.RegisterFlagCompletionFunc("name", func(cmd *cobra.Command, _ []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return kafkacmdutil.FilterValidTopicNameArgs(f, toComplete)
	})
	_ = cmd.MarkFlagRequired("name")

	flagutil.EnableOutputFlagCompletion(cmd)

	flagutil.EnableStaticFlagCompletion(cmd, "cleanup-policy", topiccmdutil.ValidCleanupPolicies)

	return cmd
}

// nolint:funlen
func runCmd(opts *options) error {
	if opts.interactive {
		// run the update command interactively
		err := runInteractivePrompt(opts)
		if err != nil {
			return err
		}

		if opts.retentionMsStr != "" {
			retentionPeriodMs, err = topiccmdutil.ConvertRetentionMsToInt(opts.retentionMsStr)
			if err != nil {
				return err
			}
		}

		if opts.retentionBytesStr != "" {
			retentionSizeBytes, err = topiccmdutil.ConvertRetentionBytesToInt(opts.retentionBytesStr)
			if err != nil {
				return err
			}
		}

		if opts.partitionsStr != "" {
			partitionCount, err = topiccmdutil.ConvertPartitionsToInt(opts.partitionsStr)
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
			return opts.localizer.MustLocalizeError("kafka.topic.common.error.topicNotFoundError", topicNameTmplPair, kafkaNameTmplPair)
		}
	}

	// map to store the config entries which will be updated
	configEntryMap := map[string]*string{}

	updateTopicReq := api.TopicsApi.UpdateTopic(opts.Context, opts.name)

	topicSettings := &kafkainstanceclient.TopicSettings{}

	if opts.retentionMsStr != "" {
		needsUpdate = true
		configEntryMap[topiccmdutil.RetentionMsKey] = &opts.retentionMsStr
	}

	if opts.retentionBytesStr != "" {
		needsUpdate = true
		configEntryMap[topiccmdutil.RetentionSizeKey] = &opts.retentionBytesStr
	}

	if opts.cleanupPolicy != "" && strings.Compare(opts.cleanupPolicy, topiccmdutil.GetConfigValue(topic.GetConfig(), topiccmdutil.CleanupPolicy)) != 0 {
		needsUpdate = true
		configEntryMap[topiccmdutil.CleanupPolicy] = &opts.cleanupPolicy
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
		configEntries := topiccmdutil.CreateConfigEntries(configEntryMap)
		topicSettings.SetConfig(*configEntries)
	}

	updateTopicReq = updateTopicReq.TopicSettings(*topicSettings)

	// update the topic
	_, httpRes, err = updateTopicReq.Execute()
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
			return opts.localizer.MustLocalizeError("kafka.topic.common.error.notFoundError", topicNameTmplPair, kafkaNameTmplPair)
		case http.StatusUnauthorized:
			return opts.localizer.MustLocalizeError("kafka.topic.common.error.unauthorized", operationTmplPair)
		case http.StatusForbidden:
			return opts.localizer.MustLocalizeError("kafka.topic.common.error.forbidden", operationTmplPair)
		case http.StatusInternalServerError:
			return opts.localizer.MustLocalizeError("kafka.topic.common.error.internalServerError")
		case http.StatusServiceUnavailable:
			return opts.localizer.MustLocalizeError("kafka.topic.common.error.unableToConnectToKafka", localize.NewEntry("Name", kafkaInstance.GetName()))
		default:
			return err
		}
	}

	opts.Logger.Info(opts.localizer.MustLocalize("kafka.topic.update.log.info.topicUpdated", topicNameTmplPair, kafkaNameTmplPair))
	return nil
}

func runInteractivePrompt(opts *options) (err error) {
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
			return opts.localizer.MustLocalizeError("kafka.topic.common.error.topicNotFoundError", topicNameTmplPair, kafkaNameTmplPair)
		}
	}

	validator := topiccmdutil.Validator{
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
		Options: topiccmdutil.ValidCleanupPolicies,
		Default: topiccmdutil.GetConfigValue(topic.GetConfig(), topiccmdutil.CleanupPolicy),
	}

	err = survey.AskOne(cleanupPolicyPrompt, &opts.cleanupPolicy)
	if err != nil {
		return err
	}

	return nil
}
