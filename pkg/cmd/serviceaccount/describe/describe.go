package describe

import (
	"context"
	"net/http"

	"github.com/redhat-developer/app-services-cli/internal/config"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/factory"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/flag"
	flagutil "github.com/redhat-developer/app-services-cli/pkg/cmdutil/flagutil"
	"github.com/redhat-developer/app-services-cli/pkg/connection"
	"github.com/redhat-developer/app-services-cli/pkg/dump"
	"github.com/redhat-developer/app-services-cli/pkg/iostreams"
	"github.com/redhat-developer/app-services-cli/pkg/localize"
	"github.com/spf13/cobra"
)

type options struct {
	id           string
	outputFormat string

	IO         *iostreams.IOStreams
	Config     config.IConfig
	Connection factory.ConnectionFunc
	localizer  localize.Localizer
	Context    context.Context
}

func NewDescribeCommand(f *factory.Factory) *cobra.Command {
	opts := &options{
		Config:     f.Config,
		Connection: f.Connection,
		IO:         f.IOStreams,
		localizer:  f.Localizer,
		Context:    f.Context,
	}

	cmd := &cobra.Command{
		Use:     "describe",
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

func runDescribe(opts *options) error {
	conn, err := opts.Connection(connection.DefaultConfigSkipMasAuth)
	if err != nil {
		return err
	}

	api := conn.API()

	res, httpRes, err := api.ServiceAccount().GetServiceAccountById(opts.Context, opts.id).Execute()
	if httpRes != nil {
		defer httpRes.Body.Close()
	}
	if err != nil {
		if httpRes == nil {
			return err
		}

		switch httpRes.StatusCode {
		case http.StatusNotFound:
			return opts.localizer.MustLocalizeError("serviceAccount.common.error.notFoundError", localize.NewEntry("ID", opts.id))
		default:
			return err
		}
	}

	return dump.Formatted(opts.IO.Out, opts.outputFormat, res)
}
