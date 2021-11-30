package create

import (
	"context"
	"net/http"
	"strconv"

	"github.com/redhat-developer/app-services-cli/pkg/kafkautil/topicutil"

	"github.com/redhat-developer/app-services-cli/pkg/core/cmdutil/factory"
	"github.com/redhat-developer/app-services-cli/pkg/core/cmdutil/flagutil"
	"github.com/redhat-developer/app-services-cli/pkg/core/config"
	"github.com/redhat-developer/app-services-cli/pkg/core/connection"
	"github.com/redhat-developer/app-services-cli/pkg/core/ioutil/dump"
	"github.com/redhat-developer/app-services-cli/pkg/core/ioutil/iostreams"
	"github.com/redhat-developer/app-services-cli/pkg/core/localize"
	"github.com/redhat-developer/app-services-cli/pkg/core/logging"
	kafkainstanceclient "github.com/redhat-developer/app-services-sdk-go/kafkainstance/apiv1internal/client"

	"github.com/AlecAivazis/survey/v2"

	"github.com/spf13/cobra"
)

const (
	defaultRetentionPeriodMS = 604800000
	defaultRetentionSize     = -1
	defaultCleanupPolicy     = "delete"
)

type options struct {
	topicName      string
	partitions     int32
	retentionMs    int
	retentionBytes int
	kafkaID        string
	outputFormat   string
	cleanupPolicy  string
	interactive    bool

	IO         *iostreams.IOStreams
	Config     config.IConfig
	Connection factory.ConnectionFunc
	Logger     logging.Logger
	localizer  localize.Localizer
	Context    context.Context
}

// NewCreateTopicCommand gets a new command for creating kafka topic.
func NewCreateTopicCommand(f *factory.Factory) *cobra.Command {
	opts := &options{
		Connection: f.Connection,
		Config:     f.Config,
		Logger:     f.Logger,
		IO:         f.IOStreams,
		localizer:  f.Localizer,
		Context:    f.Context,
	}

	cmd := &cobra.Command{
		Use:     "create",
		Short:   opts.localizer.MustLocalize("kafka.topic.create.cmd.shortDescription"),
		Long:    opts.localizer.MustLocalize("kafka.topic.create.cmd.longDescription"),
		Example: opts.localizer.MustLocalize("kafka.topic.create.cmd.example"),
		Args:    cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			if !opts.IO.CanPrompt() && opts.topicName == "" {
				return opts.localizer.MustLocalizeError("argument.error.requiredWhenNonInteractive", localize.NewEntry("Argument", "name"))
			} else if opts.topicName == "" {
				opts.interactive = true
			}

			if opts.outputFormat != "" {
				if err = flagutil.ValidateOutput(opts.outputFormat); err != nil {
					return err
				}
			}

			// check that a valid --cleanup-policy flag value is used
			validPolicy := flagutil.IsValidInput(opts.cleanupPolicy, topicutil.ValidCleanupPolicies...)
			if !validPolicy {
				return flagutil.InvalidValueError("cleanup-policy", opts.cleanupPolicy, topicutil.ValidCleanupPolicies...)
			}

			if !opts.interactive {
				validator := topicutil.Validator{
					Localizer: opts.localizer,
				}

				if err = validator.ValidateName(opts.topicName); err != nil {
					return err
				}

				if err = validator.ValidatePartitionsN(opts.partitions); err != nil {
					return err
				}

				if err = validator.ValidateMessageRetentionPeriod(opts.retentionMs); err != nil {
					return err
				}

				if err = validator.ValidateMessageRetentionSize(opts.retentionBytes); err != nil {
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

			instanceID, ok := cfg.GetKafkaIdOk()
			if !ok {
				return opts.localizer.MustLocalizeError("kafka.topic.common.error.noKafkaSelected")
			}

			opts.kafkaID = instanceID

			return runCmd(opts)
		},
	}

	flags := flagutil.NewFlagSet(cmd, opts.localizer)

	flags.StringVar(&opts.topicName, "name", "", opts.localizer.MustLocalize("kafka.topic.common.flag.name.description"))
	flags.Int32Var(&opts.partitions, "partitions", 1, opts.localizer.MustLocalize("kafka.topic.common.input.partitions.description"))
	flags.IntVar(&opts.retentionMs, "retention-ms", defaultRetentionPeriodMS, opts.localizer.MustLocalize("kafka.topic.common.input.retentionMs.description"))
	flags.IntVar(&opts.retentionBytes, "retention-bytes", defaultRetentionSize, opts.localizer.MustLocalize("kafka.topic.common.input.retentionBytes.description"))
	flags.StringVar(&opts.cleanupPolicy, "cleanup-policy", defaultCleanupPolicy, opts.localizer.MustLocalize("kafka.topic.common.input.cleanupPolicy.description"))
	flags.AddOutput(&opts.outputFormat)

	flagutil.EnableOutputFlagCompletion(cmd)
	flagutil.EnableStaticFlagCompletion(cmd, "cleanup-policy", topicutil.ValidCleanupPolicies)

	return cmd
}

// nolint:funlen
func runCmd(opts *options) error {
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

	api, kafkaInstance, err := conn.API().KafkaAdmin(opts.kafkaID)
	if err != nil {
		return err
	}

	createTopicReq := api.TopicsApi.CreateTopic(opts.Context)

	topicInput := kafkainstanceclient.NewTopicInput{
		Name: opts.topicName,
		Settings: kafkainstanceclient.TopicSettings{
			NumPartitions: opts.partitions,
			Config:        createConfigEntries(opts),
		},
	}
	createTopicReq = createTopicReq.NewTopicInput(topicInput)

	response, httpRes, err := createTopicReq.Execute()
	if httpRes != nil {
		defer httpRes.Body.Close()
	}

	if err != nil {
		if httpRes == nil {
			return err
		}

		operationTmplPair := localize.NewEntry("Operation", "create")
		switch httpRes.StatusCode {
		case http.StatusUnauthorized:
			return opts.localizer.MustLocalizeError("kafka.topic.common.error.unauthorized", operationTmplPair)
		case http.StatusForbidden:
			return opts.localizer.MustLocalizeError("kafka.topic.common.error.forbidden", operationTmplPair)
		case http.StatusConflict:
			return opts.localizer.MustLocalizeError("kafka.topic.create.error.conflictError", localize.NewEntry("TopicName", opts.topicName), localize.NewEntry("InstanceName", kafkaInstance.GetName()))
		case http.StatusInternalServerError:
			return opts.localizer.MustLocalizeError("kafka.topic.common.error.internalServerError")
		case http.StatusServiceUnavailable:
			return opts.localizer.MustLocalizeError("kafka.topic.common.error.unableToConnectToKafka", localize.NewEntry("Name", kafkaInstance.GetName()))
		default:
			return err
		}
	}

	opts.Logger.Info(opts.localizer.MustLocalize("kafka.topic.create.log.info.topicCreated", localize.NewEntry("TopicName", response.GetName()), localize.NewEntry("InstanceName", kafkaInstance.GetName())))

	return dump.Formatted(opts.IO.Out, opts.outputFormat, response)
}

func runInteractivePrompt(opts *options) (err error) {
	validator := topicutil.Validator{
		Localizer:  opts.localizer,
		InstanceID: opts.kafkaID,
		Connection: opts.Connection,
	}

	opts.Logger.Debug(opts.localizer.MustLocalize("common.log.debug.startingInteractivePrompt"))

	promptName := &survey.Input{
		Message: opts.localizer.MustLocalize("kafka.topic.common.input.name.message"),
		Help:    opts.localizer.MustLocalize("kafka.topic.common.input.name.help"),
	}

	err = survey.AskOne(
		promptName,
		&opts.topicName,
		survey.WithValidator(survey.Required),
		survey.WithValidator(validator.ValidateName),
		survey.WithValidator(validator.ValidateNameIsAvailable),
	)

	if err != nil {
		return err
	}

	partitionsPrompt := &survey.Input{
		Message: opts.localizer.MustLocalize("kafka.topic.create.input.partitions.message"),
		Help:    opts.localizer.MustLocalize("kafka.topic.common.input.partitions.description"),
		Default: "1",
	}

	err = survey.AskOne(partitionsPrompt, &opts.partitions, survey.WithValidator(validator.ValidatePartitionsN))
	if err != nil {
		return err
	}

	retentionMsPrompt := &survey.Input{
		Message: opts.localizer.MustLocalize("kafka.topic.create.input.retentionMs.message"),
		Help:    opts.localizer.MustLocalize("kafka.topic.common.input.retentionMs.description"),
		Default: strconv.Itoa(defaultRetentionPeriodMS),
	}

	err = survey.AskOne(retentionMsPrompt, &opts.retentionMs, survey.WithValidator(validator.ValidateMessageRetentionPeriod))
	if err != nil {
		return err
	}

	retentionBytesPrompt := &survey.Input{
		Message: opts.localizer.MustLocalize("kafka.topic.create.input.retentionBytes.message"),
		Help:    opts.localizer.MustLocalize("kafka.topic.common.input.retentionBytes.description"),
		Default: strconv.Itoa(defaultRetentionSize),
	}

	err = survey.AskOne(retentionBytesPrompt, &opts.retentionBytes, survey.WithValidator(validator.ValidateMessageRetentionSize))
	if err != nil {
		return err
	}

	cleanupPolicyPrompt := &survey.Select{
		Message: opts.localizer.MustLocalize("kafka.topic.create.input.cleanupPolicy.message"),
		Help:    opts.localizer.MustLocalize("kafka.topic.common.input.cleanupPolicy.description"),
		Options: topicutil.ValidCleanupPolicies,
		Default: defaultCleanupPolicy,
	}

	err = survey.AskOne(cleanupPolicyPrompt, &opts.cleanupPolicy)
	if err != nil {
		return err
	}

	return nil
}

func createConfigEntries(opts *options) *[]kafkainstanceclient.ConfigEntry {
	retentionMsStr := strconv.Itoa(opts.retentionMs)
	retentionBytesStr := strconv.Itoa(opts.retentionBytes)
	cleanupPolicyStr := opts.cleanupPolicy
	configEntryMap := map[string]*string{
		topicutil.RetentionMsKey:   &retentionMsStr,
		topicutil.RetentionSizeKey: &retentionBytesStr,
		topicutil.CleanupPolicy:    &cleanupPolicyStr,
	}
	return topicutil.CreateConfigEntries(configEntryMap)
}
