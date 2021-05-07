package create

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"

	"github.com/AlecAivazis/survey/v2"

	"github.com/redhat-developer/app-services-cli/pkg/connection"
	topicutil "github.com/redhat-developer/app-services-cli/pkg/kafka/topic"
	"github.com/redhat-developer/app-services-cli/pkg/localize"

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

const (
	defaultRetentionPeriodMS = 604800000
	defaultRetentionSize     = -1
)

type Options struct {
	topicName      string
	partitions     int32
	retentionMs    int
	retentionBytes int
	kafkaID        string
	outputFormat   string
	interactive    bool

	IO         *iostreams.IOStreams
	Config     config.IConfig
	Connection factory.ConnectionFunc
	Logger     func() (logging.Logger, error)
	localizer  localize.Localizer
}

// NewCreateTopicCommand gets a new command for creating kafka topic.
func NewCreateTopicCommand(f *factory.Factory) *cobra.Command {
	opts := &Options{
		Connection: f.Connection,
		Config:     f.Config,
		Logger:     f.Logger,
		IO:         f.IOStreams,
		localizer:  f.Localizer,
	}

	cmd := &cobra.Command{
		Use:     opts.localizer.MustLocalize("kafka.topic.create.cmd.use"),
		Short:   opts.localizer.MustLocalize("kafka.topic.create.cmd.shortDescription"),
		Long:    opts.localizer.MustLocalize("kafka.topic.create.cmd.longDescription"),
		Example: opts.localizer.MustLocalize("kafka.topic.create.cmd.example"),
		Args:    cobra.RangeArgs(0, 1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			if !opts.IO.CanPrompt() && len(args) == 0 {
				return errors.New(opts.localizer.MustLocalize("argument.error.requiredWhenNonInteractive", localize.NewEntry("Argument", "name")))
			} else if len(args) == 0 {
				opts.interactive = true
			}

			if err = flag.ValidateOutput(opts.outputFormat); err != nil {
				return err
			}

			if !opts.interactive {
				opts.topicName = args[0]

				if err = topicutil.ValidateName(opts.topicName); err != nil {
					return err
				}

				if err = topicutil.ValidatePartitionsN(opts.partitions); err != nil {
					return err
				}

				if err = topicutil.ValidateMessageRetentionPeriod(opts.retentionMs); err != nil {
					return err
				}

				if err = topicutil.ValidateMessageRetentionSize(opts.retentionBytes); err != nil {
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
				return fmt.Errorf(opts.localizer.MustLocalize("kafka.topic.common.error.noKafkaSelected"))
			}

			opts.kafkaID = cfg.Services.Kafka.ClusterID

			return runCmd(opts)
		},
	}

	cmd.Flags().StringVarP(&opts.outputFormat, "output", "o", "json", opts.localizer.MustLocalize("kafka.topic.common.flag.output.description"))
	cmd.Flags().Int32Var(&opts.partitions, "partitions", 1, opts.localizer.MustLocalize("kafka.topic.common.input.partitions.description"))
	cmd.Flags().IntVar(&opts.retentionMs, "retention-ms", defaultRetentionPeriodMS, opts.localizer.MustLocalize("kafka.topic.common.input.retentionMs.description"))
	cmd.Flags().IntVar(&opts.retentionBytes, "retention-bytes", defaultRetentionSize, opts.localizer.MustLocalize("kafka.topic.common.input.retentionBytes.description"))

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
			Config:        createConfigEntries(opts),
		},
	}
	createTopicReq = createTopicReq.NewTopicInput(topicInput)

	response, httpRes, err := createTopicReq.Execute()
	if err != nil {
		if httpRes == nil {
			return err
		}

		operationTmplPair := localize.NewEntry("Operation", "create")
		switch httpRes.StatusCode {
		case 401:
			return errors.New(opts.localizer.MustLocalize("kafka.topic.common.error.unauthorized", operationTmplPair))
		case 403:
			return errors.New(opts.localizer.MustLocalize("kafka.topic.common.error.forbidden", operationTmplPair))
		case 409:
			return errors.New(opts.localizer.MustLocalize("kafka.topic.create.error.conflictError", localize.NewEntry("TopicName", opts.topicName), localize.NewEntry("InstanceName", kafkaInstance.GetName())))
		case 500:
			return errors.New(opts.localizer.MustLocalize("kafka.topic.common.error.internalServerError"))
		case 503:
			return errors.New(opts.localizer.MustLocalize("kafka.topic.common.error.unableToConnectToKafka", localize.NewEntry("Name", kafkaInstance.GetName())))
		default:
			return err
		}
	}

	logger.Info(opts.localizer.MustLocalize("kafka.topic.create.log.info.topicCreated", localize.NewEntry("TopicName", response.GetName()), localize.NewEntry("InstanceName", kafkaInstance.GetName())))

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

	logger.Debug(opts.localizer.MustLocalize("common.log.debug.startingInteractivePrompt"))

	promptName := &survey.Input{
		Message: opts.localizer.MustLocalize("kafka.topic.common.input.name.message"),
		Help:    opts.localizer.MustLocalize("kafka.topic.common.input.name.help"),
	}

	err = survey.AskOne(
		promptName,
		&opts.topicName,
		survey.WithValidator(survey.Required),
		survey.WithValidator(topicutil.ValidateName),
		survey.WithValidator(topicutil.ValidateNameIsAvailable(api, kafkaInstance.GetName())),
	)

	if err != nil {
		return err
	}

	partitionsPrompt := &survey.Input{
		Message: opts.localizer.MustLocalize("kafka.topic.create.input.partitions.message"),
		Help:    opts.localizer.MustLocalize("kafka.topic.common.input.partitions.description"),
		Default: "1",
	}

	err = survey.AskOne(partitionsPrompt, &opts.partitions, survey.WithValidator(topicutil.ValidatePartitionsN))
	if err != nil {
		return err
	}

	retentionMsPrompt := &survey.Input{
		Message: opts.localizer.MustLocalize("kafka.topic.create.input.retentionMs.message"),
		Help:    opts.localizer.MustLocalize("kafka.topic.common.input.retentionMs.description"),
		Default: strconv.Itoa(defaultRetentionPeriodMS),
	}

	err = survey.AskOne(retentionMsPrompt, &opts.retentionMs, survey.WithValidator(topicutil.ValidateMessageRetentionPeriod))
	if err != nil {
		return err
	}

	retentionBytesPrompt := &survey.Input{
		Message: opts.localizer.MustLocalize("kafka.topic.create.input.retentionBytes.message"),
		Help:    opts.localizer.MustLocalize("kafka.topic.common.input.retentionBytes.description"),
		Default: strconv.Itoa(defaultRetentionSize),
	}

	err = survey.AskOne(retentionBytesPrompt, &opts.retentionBytes, survey.WithValidator(topicutil.ValidateMessageRetentionSize))
	if err != nil {
		return err
	}

	return nil
}

func createConfigEntries(opts *Options) *[]strimziadminclient.ConfigEntry {
	retentionMsStr := strconv.Itoa(opts.retentionMs)
	retentionBytesStr := strconv.Itoa(opts.retentionBytes)
	configEntryMap := map[string]*string{
		topicutil.RetentionMsKey:   &retentionMsStr,
		topicutil.RetentionSizeKey: &retentionBytesStr,
	}
	return topicutil.CreateConfigEntries(configEntryMap)
}
