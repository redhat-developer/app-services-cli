package request

import (
	"context"
	"errors"

	"github.com/redhat-developer/app-services-cli/pkg/core/ioutil/iostreams"
	"github.com/redhat-developer/app-services-cli/pkg/core/localize"
	"github.com/redhat-developer/app-services-cli/pkg/core/logging"
	"github.com/redhat-developer/app-services-cli/pkg/shared/connection"
	"github.com/redhat-developer/app-services-cli/pkg/shared/factory"
	"github.com/spf13/cobra"
)

type options struct {
	IO         *iostreams.IOStreams
	Logger     logging.Logger
	localizer  localize.Localizer
	Context    context.Context
	Connection factory.ConnectionFunc

	urlPath string
	body    string
}

func NewCallCmd(f *factory.Factory) *cobra.Command {
	opts := &options{
		IO:         f.IOStreams,
		Logger:     f.Logger,
		localizer:  f.Localizer,
		Context:    f.Context,
		Connection: f.Connection,
	}

	cmd := &cobra.Command{
		Use:    "request",
		Short:  "Allows you to perform API requests against the API server",
		Hidden: true,
		Args:   cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return runCmd(opts)
		},
	}
	cmd.Flags().StringVar(&opts.urlPath, "path", "", "Path to send request. For example /api/kafkas_mgmt/v1/kafkas?async=true")
	cmd.Flags().StringVar(&opts.body, "body", "", "If body present then it will be used as request body for post request")
	return cmd
}

func runCmd(opts *options) (err error) {
	if opts.urlPath == "" {
		return errors.New("--path is required")
	}
	opts.Logger.Info("Performing request to", opts.urlPath)
	conn, err := opts.Connection(connection.DefaultConfigSkipMasAuth)

	if err != nil {
		return err
	}

	var data interface{}
	if opts.body == "" {
		data, _, err = conn.API().GenericAPI().GET(opts.Context, opts.urlPath)
	} else {
		data, _, err = conn.API().GenericAPI().POST(opts.Context, opts.urlPath, opts.body)
	}

	if err != nil || data == nil {
		opts.Logger.Info("Fetching data failed", err)
		return nil
	}

	opts.Logger.Info("Response:", data)
	return nil
}
