package describe

import (
	"context"
	"encoding/json"
	"errors"

	flagutil "github.com/bf2fc6cc711aee1a0c2a/cli/pkg/cmdutil/flags"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/connection"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/iostreams"

	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/cmd/flag"

	"github.com/bf2fc6cc711aee1a0c2a/cli/internal/config"
	"github.com/bf2fc6cc711aee1a0c2a/cli/internal/localizer"
	kasclient "github.com/bf2fc6cc711aee1a0c2a/cli/pkg/api/kas/client"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/cmd/factory"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/cmdutil"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/dump"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/kafka"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

type Options struct {
	id           string
	name         string
	outputFormat string

	IO         *iostreams.IOStreams
	Config     config.IConfig
	Connection factory.ConnectionFunc
}

// NewDescribeCommand describes a Kafka instance, either by passing an `--id flag`
// or by using the kafka instance set in the config, if any
func NewDescribeCommand(f *factory.Factory) *cobra.Command {
	opts := &Options{
		Config:     f.Config,
		Connection: f.Connection,
		IO:         f.IOStreams,
	}

	cmd := &cobra.Command{
		Use:     localizer.MustLocalizeFromID("kafka.describe.cmd.use"),
		Short:   localizer.MustLocalizeFromID("kafka.describe.cmd.shortDescription"),
		Long:    localizer.MustLocalizeFromID("kafka.describe.cmd.longDescription"),
		Example: localizer.MustLocalizeFromID("kafka.describe.cmd.example"),
		Args:    cobra.RangeArgs(0, 1),
		// Dynamic completion of the Kafka name
		ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
			var searchName string
			if len(args) > 0 {
				searchName = args[0]
			}
			return cmdutil.FilterValidKafkas(f, searchName)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			validOutputFormats := flagutil.ValidOutputFormats
			if opts.outputFormat != "" && !flagutil.IsValidInput(opts.outputFormat, validOutputFormats...) {
				return flag.InvalidValueError("output", opts.outputFormat, validOutputFormats...)
			}

			if len(args) > 0 {
				opts.name = args[0]
			}

			if opts.name != "" && opts.id != "" {
				return errors.New(localizer.MustLocalizeFromID("kafka.common.error.idAndNameCannotBeUsed"))
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
				return errors.New(localizer.MustLocalizeFromID("kafka.common.error.noKafkaSelected"))
			}

			opts.id = cfg.Services.Kafka.ClusterID

			return runDescribe(opts)
		},
	}

	cmd.Flags().StringVarP(&opts.outputFormat, "output", "o", "json", localizer.MustLocalizeFromID("kafka.common.flag.output.description"))
	cmd.Flags().StringVar(&opts.id, "id", "", localizer.MustLocalizeFromID("kafka.describe.flag.id"))

	return cmd
}

func runDescribe(opts *Options) error {
	connection, err := opts.Connection(connection.DefaultConfigSkipMasAuth)
	if err != nil {
		return err
	}

	api := connection.API()

	var kafkaInstance *kasclient.KafkaRequest
	ctx := context.Background()
	if opts.name != "" {
		kafkaInstance, _, err = kafka.GetKafkaByName(ctx, api.Kafka(), opts.name)
		if err.Error() != "" {
			return err
		}
	} else {
		kafkaInstance, _, err = kafka.GetKafkaByID(ctx, api.Kafka(), opts.id)
		if err.Error() != "" {
			return err
		}
	}

	return printKafka(kafkaInstance, opts)
}

func printKafka(kafka *kasclient.KafkaRequest, opts *Options) error {
	switch opts.outputFormat {
	case "yaml", "yml":
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
