package describe

import (
	"context"
	"encoding/json"
	"errors"

	flagutil "github.com/redhat-developer/app-services-cli/pkg/cmdutil/flags"
	"github.com/redhat-developer/app-services-cli/pkg/connection"
	"github.com/redhat-developer/app-services-cli/pkg/iostreams"
	"github.com/redhat-developer/app-services-cli/pkg/localize"
	"github.com/redhat-developer/app-services-cli/pkg/serviceregistry"

	"github.com/redhat-developer/app-services-cli/pkg/cmd/flag"

	"github.com/redhat-developer/app-services-cli/internal/config"
	srsclient "github.com/redhat-developer/app-services-cli/pkg/api/srs/client"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/factory"
	"github.com/redhat-developer/app-services-cli/pkg/cmdutil"
	"github.com/redhat-developer/app-services-cli/pkg/dump"
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
	localizer  localize.Localizer
}

// NewDescribeCommand describes a Kafka instance, either by passing an `--id flag`
// or by using the kafka instance set in the config, if any
func NewDescribeCommand(f *factory.Factory) *cobra.Command {
	opts := &Options{
		Config:     f.Config,
		Connection: f.Connection,
		IO:         f.IOStreams,
		localizer:  f.Localizer,
	}

	cmd := &cobra.Command{
		Use:     "describe",
		Short:   "Describe service registry",
		Long:    "",
		Example: "",
		Args:    cobra.RangeArgs(0, 1),
		// Dynamic completion of the Kafka name
		ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
			return cmdutil.FilterValidKafkas(f, toComplete)
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
				return errors.New(opts.localizer.MustLocalize("kafka.common.error.idAndNameCannotBeUsed"))
			}

			if opts.id != "" || opts.name != "" {
				return runDescribe(opts)
			}

			// TODO implement config part
			return errors.New("Not implemented. Use id or name as argument")

			// cfg, err := opts.Config.Load()
			// if err != nil {
			// 	return err
			// }

			// var kafkaConfig *config.KafkaConfig
			// if cfg.Services.Kafka == kafkaConfig || cfg.Services.Kafka.ClusterID == "" {
			// 	return errors.New(opts.localizer.MustLocalize("kafka.common.error.noKafkaSelected"))
			// }

			// opts.id = cfg.Services.Kafka.ClusterID

			// return runDescribe(opts)
		},
	}

	cmd.Flags().StringVarP(&opts.outputFormat, "output", "o", "json", opts.localizer.MustLocalize("kafka.common.flag.output.description"))
	cmd.Flags().StringVar(&opts.id, "id", "", opts.localizer.MustLocalize("kafka.describe.flag.id"))

	return cmd
}

func runDescribe(opts *Options) error {
	connection, err := opts.Connection(connection.DefaultConfigSkipMasAuth)
	if err != nil {
		return err
	}

	api := connection.API()

	var registry *srsclient.Registry
	ctx := context.Background()
	if opts.name != "" {
		registry, _, err = serviceregistry.GetServiceRegistryByName(ctx, api.ServiceRegistry(), opts.name)
		if err != nil {
			return err
		}
	} else {
		registry, _, err = serviceregistry.GetServiceRegistryByID(ctx, api.ServiceRegistry(), opts.id)
		if err != nil {
			return err
		}
	}

	return printService(registry, opts)
}

// TODO
func printService(registry interface{}, opts *Options) error {
	switch opts.outputFormat {
	case "yaml", "yml":
		data, err := yaml.Marshal(registry)
		if err != nil {
			return err
		}
		return dump.YAML(opts.IO.Out, data)
	default:
		data, err := json.Marshal(registry)
		if err != nil {
			return err
		}
		return dump.JSON(opts.IO.Out, data)
	}
}
