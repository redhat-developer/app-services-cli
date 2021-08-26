package describe

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/redhat-developer/app-services-cli/internal/config"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/factory"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/flag"
	flagutil "github.com/redhat-developer/app-services-cli/pkg/cmdutil/flags"
	"github.com/redhat-developer/app-services-cli/pkg/connection"
	"github.com/redhat-developer/app-services-cli/pkg/dump"
	"github.com/redhat-developer/app-services-cli/pkg/iostreams"
	"github.com/redhat-developer/app-services-cli/pkg/localize"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

type Options struct {
	id           string
	outputFormat string

	IO         *iostreams.IOStreams
	Config     config.IConfig
	Connection factory.ConnectionFunc
	localizer  localize.Localizer
}

func NewDescribeCommand(f *factory.Factory) *cobra.Command {
	opts := &Options{
		Config:     f.Config,
		Connection: f.Connection,
		IO:         f.IOStreams,
		localizer:  f.Localizer,
	}

	cmd := &cobra.Command{
		Use:     opts.localizer.MustLocalize("serviceAccount.describe.cmd.use"),
		Short:   opts.localizer.MustLocalize("serviceAccount.describe.cmd.shortDescription"),
		Long:    opts.localizer.MustLocalize("serviceAccount.describe.cmd.longDescription"),
		Example: opts.localizer.MustLocalize("serviceAccount.describe.cmd.example"),
		Args:    cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			validOutputFormats := flagutil.ValidOutputFormats
			if opts.outputFormat != "" && !flagutil.IsValidInput(opts.outputFormat, validOutputFormats...) {
				return flag.InvalidValueError("output", opts.outputFormat, validOutputFormats...)
			}

			return runDescribe(opts)
		},
	}

	cmd.Flags().StringVar(&opts.id, "id", "", opts.localizer.MustLocalize("serviceAccount.describe.flag.id.description"))
	cmd.Flags().StringVarP(&opts.outputFormat, "output", "o", "json", opts.localizer.MustLocalize("serviceAccount.common.flag.output.description"))

	_ = cmd.MarkFlagRequired("id")

	flagutil.EnableOutputFlagCompletion(cmd)

	return cmd
}

func runDescribe(opts *Options) error {
	connection, err := opts.Connection(connection.DefaultConfigSkipMasAuth)
	if err != nil {
		return err
	}

	api := connection.API()

	res, httpRes, err := api.ServiceAccount().GetServiceAccountById(context.Background(), opts.id).Execute()
	defer httpRes.Body.Close()
	if err != nil {
		if httpRes == nil {
			return err
		}

		switch httpRes.StatusCode {
		case http.StatusNotFound:
			return errors.New(opts.localizer.MustLocalize("serviceAccount.common.error.notFoundError", localize.NewEntry("ID", opts.id)))
		default:
			return err
		}
	}

	switch opts.outputFormat {
	case dump.YAMLFormat, dump.YMLFormat:
		data, _ := yaml.Marshal(res)
		_ = dump.YAML(opts.IO.Out, data)
	default:
		data, _ := json.Marshal(res)
		_ = dump.JSON(opts.IO.Out, data)
	}

	return nil
}
