package describe

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/redhat-developer/app-services-cli/internal/localizer"
	"github.com/redhat-developer/app-services-cli/pkg/cmdutil"
	"github.com/redhat-developer/app-services-cli/pkg/connection"

	"github.com/redhat-developer/app-services-cli/pkg/cmd/flag"

	"github.com/redhat-developer/app-services-cli/pkg/dump"
	"github.com/redhat-developer/app-services-cli/pkg/iostreams"
	"gopkg.in/yaml.v2"

	"github.com/redhat-developer/app-services-cli/internal/config"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/factory"
	"github.com/redhat-developer/app-services-cli/pkg/logging"

	"github.com/spf13/cobra"
)

type Options struct {
	topicName    string
	kafkaID      string
	outputFormat string

	IO         *iostreams.IOStreams
	Config     config.IConfig
	Connection factory.ConnectionFunc
	Logger     func() (logging.Logger, error)
}

// NewDescribeTopicCommand gets a new command for describing a kafka topic.
func NewDescribeTopicCommand(f *factory.Factory) *cobra.Command {
	opts := &Options{
		Connection: f.Connection,
		Config:     f.Config,
		Logger:     f.Logger,
		IO:         f.IOStreams,
	}

	cmd := &cobra.Command{
		Use:     localizer.MustLocalizeFromID("kafka.topic.describe.cmd.use"),
		Short:   localizer.MustLocalizeFromID("kafka.topic.describe.cmd.shortDescription"),
		Long:    localizer.MustLocalizeFromID("kafka.topic.describe.cmd.longDescription"),
		Example: localizer.MustLocalizeFromID("kafka.topic.describe.cmd.example"),
		Args:    cobra.ExactValidArgs(1),
		// dynamic completion of topic names
		ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
			validNames := []string{}

			var searchName string
			if len(args) > 0 {
				searchName = args[0]
			}

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
			if len(args) > 0 {
				opts.topicName = args[0]
			}

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

			if !cfg.HasKafka() {
				return errors.New(localizer.MustLocalizeFromID("kafka.topic.common.error.noKafkaSelected"))
			}

			opts.kafkaID = cfg.Services.Kafka.ClusterID

			return runCmd(opts)
		},
	}

	cmd.Flags().StringVarP(&opts.outputFormat, "output", "o", "json", localizer.MustLocalize(&localizer.Config{
		MessageID: "kafka.topic.common.flag.output.description",
	}))

	return cmd
}

func runCmd(opts *Options) error {
	conn, err := opts.Connection(connection.DefaultConfigRequireMasAuth)
	if err != nil {
		return err
	}

	api, kafkaInstance, err := conn.API().TopicAdmin(opts.kafkaID)
	if err != nil {
		return err
	}

	// fetch the topic
	topicResponse, httpRes, err := api.
		GetTopic(context.Background(), opts.topicName).
		Execute()

	if err != nil {
		if httpRes == nil {
			return err
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
					"Operation": "view",
				},
			}))
		case 403:
			return errors.New(localizer.MustLocalize(&localizer.Config{
				MessageID: "kafka.topic.common.error.forbidden",
				TemplateData: map[string]interface{}{
					"Operation": "view",
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
			}), err)
		default:
			return err
		}
	}

	switch opts.outputFormat {
	case "json":
		data, _ := json.Marshal(topicResponse)
		_ = dump.JSON(opts.IO.Out, data)
	case "yaml", "yml":
		data, _ := yaml.Marshal(topicResponse)
		_ = dump.YAML(opts.IO.Out, data)
	}

	return nil
}
