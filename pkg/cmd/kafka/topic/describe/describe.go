package describe

import (
	"context"
	"net/http"

	kafkacmdutil "github.com/redhat-developer/app-services-cli/pkg/shared/kafkautil"

	"github.com/redhat-developer/app-services-cli/pkg/core/cmdutil/flagutil"
	"github.com/redhat-developer/app-services-cli/pkg/core/config"
	"github.com/redhat-developer/app-services-cli/pkg/core/ioutil/dump"
	"github.com/redhat-developer/app-services-cli/pkg/core/ioutil/iostreams"
	"github.com/redhat-developer/app-services-cli/pkg/core/localize"
	"github.com/redhat-developer/app-services-cli/pkg/core/logging"
	"github.com/redhat-developer/app-services-cli/pkg/shared/factory"
	"github.com/spf13/cobra"
)

type options struct {
	name         string
	kafkaID      string
	outputFormat string

	IO         *iostreams.IOStreams
	Config     config.IConfig
	Connection factory.ConnectionFunc
	Logger     logging.Logger
	localizer  localize.Localizer
	Context    context.Context
}

// NewDescribeTopicCommand gets a new command for describing a kafka topic.
func NewDescribeTopicCommand(f *factory.Factory) *cobra.Command {
	opts := &options{
		Connection: f.Connection,
		Config:     f.Config,
		Logger:     f.Logger,
		IO:         f.IOStreams,
		localizer:  f.Localizer,
		Context:    f.Context,
	}

	cmd := &cobra.Command{
		Use:     "describe",
		Short:   opts.localizer.MustLocalize("kafka.topic.describe.cmd.shortDescription"),
		Long:    opts.localizer.MustLocalize("kafka.topic.describe.cmd.longDescription"),
		Example: opts.localizer.MustLocalize("kafka.topic.describe.cmd.example"),
		Args:    cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			if opts.outputFormat != "" {
				if err = flagutil.ValidateOutput(opts.outputFormat); err != nil {
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

	flags.AddOutput(&opts.outputFormat)

	flags.StringVar(&opts.name, "name", "", opts.localizer.MustLocalize("kafka.topic.common.flag.output.description"))
	_ = cmd.RegisterFlagCompletionFunc("name", func(cmd *cobra.Command, _ []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return kafkacmdutil.FilterValidTopicNameArgs(f, toComplete)
	})
	_ = cmd.MarkFlagRequired("name")

	flagutil.EnableOutputFlagCompletion(cmd)

	return cmd
}

func runCmd(opts *options) error {
	conn, err := opts.Connection()
	if err != nil {
		return err
	}

	api, kafkaInstance, err := conn.API().KafkaAdmin(opts.kafkaID)
	if err != nil {
		return err
	}

	// fetch the topic
	topicResponse, httpRes, err := api.TopicsApi.GetTopic(opts.Context, opts.name).Execute()
	if httpRes != nil {
		defer httpRes.Body.Close()
	}

	if err != nil {
		if httpRes == nil {
			return err
		}

		topicNameTmplPair := localize.NewEntry("TopicName", opts.name)
		kafkaNameTmplPair := localize.NewEntry("InstanceName", kafkaInstance.GetName())
		operationTmplPair := localize.NewEntry("Operation", "describe")

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

	return dump.Formatted(opts.IO.Out, opts.outputFormat, topicResponse)
}
