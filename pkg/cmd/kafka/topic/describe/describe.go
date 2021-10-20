package describe

import (
	"context"
	"net/http"

	"github.com/redhat-developer/app-services-cli/pkg/cmdutil"
	"github.com/redhat-developer/app-services-cli/pkg/connection"
	"github.com/redhat-developer/app-services-cli/pkg/localize"

	"github.com/redhat-developer/app-services-cli/pkg/cmd/flag"

	"github.com/redhat-developer/app-services-cli/internal/config"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/factory"
	"github.com/redhat-developer/app-services-cli/pkg/dump"
	"github.com/redhat-developer/app-services-cli/pkg/iostreams"
	"github.com/redhat-developer/app-services-cli/pkg/logging"

	flagutil "github.com/redhat-developer/app-services-cli/pkg/cmdutil/flagutil"

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
				if err = flag.ValidateOutput(opts.outputFormat); err != nil {
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

	cmd.Flags().StringVarP(&opts.outputFormat, "output", "o", "json", opts.localizer.MustLocalize("kafka.topic.common.flag.output.description"))

	cmd.Flags().StringVar(&opts.name, "name", "", opts.localizer.MustLocalize("kafka.topic.common.flag.output.description"))
	_ = cmd.RegisterFlagCompletionFunc("name", func(cmd *cobra.Command, _ []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return cmdutil.FilterValidTopicNameArgs(f, toComplete)
	})
	_ = cmd.MarkFlagRequired("name")

	flagutil.EnableOutputFlagCompletion(cmd)

	return cmd
}

func runCmd(opts *options) error {
	conn, err := opts.Connection(connection.DefaultConfigRequireMasAuth)
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
