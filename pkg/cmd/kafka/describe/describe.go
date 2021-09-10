package describe

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	flagutil "github.com/redhat-developer/app-services-cli/pkg/cmdutil/flags"
	"github.com/redhat-developer/app-services-cli/pkg/connection"
	"github.com/redhat-developer/app-services-cli/pkg/iostreams"
	kafkacmdutil "github.com/redhat-developer/app-services-cli/pkg/kafka/cmdutil"
	"github.com/redhat-developer/app-services-cli/pkg/localize"

	"github.com/redhat-developer/app-services-cli/pkg/cmd/flag"

	"github.com/redhat-developer/app-services-cli/internal/config"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/factory"
	"github.com/redhat-developer/app-services-cli/pkg/dump"
	"github.com/redhat-developer/app-services-cli/pkg/kafka"
	"github.com/redhat-developer/app-services-cli/pkg/logging"
	kafkamgmtclient "github.com/redhat-developer/app-services-sdk-go/kafkamgmt/apiv1/client"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

type args struct {
	id           string
	name         string
	outputFormat string

	IO         *iostreams.IOStreams
	Config     config.IConfig
	Connection factory.ConnectionFunc
	Logger     logging.Logger
	localizer  localize.Localizer
	Context    context.Context
}

// NewDescribeCommand describes a Kafka instance, either by passing an `--id flag`
// or by using the kafka instance set in the config, if any
func NewDescribeCommand(f *factory.Factory) *cobra.Command {
	opts := &args{
		Config:     f.Config,
		Connection: f.Connection,
		IO:         f.IOStreams,
		Logger:     f.Logger,
		localizer:  f.Localizer,
		Context:    f.Context,
	}

	cmd := &cobra.Command{
		Use:     "describe",
		Short:   opts.localizer.MustLocalize("kafka.describe.cmd.shortDescription"),
		Long:    opts.localizer.MustLocalize("kafka.describe.cmd.longDescription"),
		Example: opts.localizer.MustLocalize("kafka.describe.cmd.example"),
		Args:    cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			validOutputFormats := flagutil.ValidOutputFormats
			if opts.outputFormat != "" && !flagutil.IsValidInput(opts.outputFormat, validOutputFormats...) {
				return flag.InvalidValueError("output", opts.outputFormat, validOutputFormats...)
			}

			if opts.name != "" && opts.id != "" {
				return errors.New(opts.localizer.MustLocalize("service.error.idAndNameCannotBeUsed"))
			}

			if opts.id != "" || opts.name != "" {
				return runDescribe(opts)
			}

			cfg, err := opts.Config.Load()
			if err != nil {
				return err
			}

			var kafkaConfig *config.KafkaConfig
			if cfg.Services.Kafka == kafkaConfig || cfg.Services.Kafka.ClusterID == "" {
				return errors.New(opts.localizer.MustLocalize("kafka.common.error.noKafkaSelected"))
			}

			opts.id = cfg.Services.Kafka.ClusterID

			return runDescribe(opts)
		},
	}

	cmd.Flags().StringVarP(&opts.outputFormat, "output", "o", "json", opts.localizer.MustLocalize("kafka.common.flag.output.description"))
	cmd.Flags().StringVar(&opts.id, "id", "", opts.localizer.MustLocalize("kafka.describe.flag.id"))
	cmd.Flags().StringVar(&opts.name, "name", "", opts.localizer.MustLocalize("kafka.describe.flag.name"))

	if err := kafkacmdutil.RegisterNameFlagCompletionFunc(cmd, f); err != nil {
		opts.Logger.Debug(opts.localizer.MustLocalize("kafka.common.error.load.completions.name.flag"), err)
	}
	flagutil.EnableOutputFlagCompletion(cmd)

	return cmd
}

func runDescribe(opts *args) error {
	conn, err := opts.Connection(connection.DefaultConfigSkipMasAuth)
	if err != nil {
		return err
	}

	api := conn.API()

	var kafkaInstance *kafkamgmtclient.KafkaRequest
	var httpRes *http.Response
	if opts.name != "" {
		kafkaInstance, httpRes, err = kafka.GetKafkaByName(opts.Context, api.Kafka(), opts.name)
		if httpRes != nil {
			defer httpRes.Body.Close()
		}
		if err != nil {
			return err
		}
	} else {
		kafkaInstance, httpRes, err = kafka.GetKafkaByID(opts.Context, api.Kafka(), opts.id)
		if httpRes != nil {
			defer httpRes.Body.Close()
		}
		if err != nil {
			return err
		}
	}

	return printKafka(kafkaInstance, opts)
}

func printKafka(kafka *kafkamgmtclient.KafkaRequest, opts *args) error {
	switch opts.outputFormat {
	case dump.YAMLFormat, dump.YMLFormat:
		data, err := yaml.Marshal(kafka)
		if err != nil {
			return err
		}
		return dump.YAML(opts.IO.Out, data)
	default:
		data, err := json.Marshal(kafka)
		if err != nil {
			return err
		}
		return dump.JSON(opts.IO.Out, data)
	}
}
