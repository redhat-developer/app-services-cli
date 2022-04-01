package request

import (
	"errors"
	"fmt"

	"github.com/redhat-developer/app-services-cli/pkg/api/generic"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/registry/artifact/util"
	"github.com/redhat-developer/app-services-cli/pkg/shared/connection"
	"github.com/redhat-developer/app-services-cli/pkg/shared/factory"
	"github.com/spf13/cobra"
)

type options struct {
	factory *factory.Factory

	urlPath string
	method  string
}

type GenericAPI interface {
	GenericAPI() generic.GenericAPI
}

func NewRequestCmd(f *factory.Factory) *cobra.Command {
	opts := &options{
		factory: f,
	}

	cmd := &cobra.Command{
		Use:   "request",
		Short: "Allows you to perform API requests against the API server",
		// TODO paths in examples need to get command name. Command name needs to be dynamic
		Example: `
		  # Perform a GET request to the specified path
		  rhoas request --path /api/kafkas_mgmt/v1/kafkas
		  
		  # Perform a POST request to the specified path
		  cat request.json | rhoas request --path "/api/kafkas_mgmt/v1/kafkas?async=true" --method post `,
		Args: cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return runCmd(opts)
		},
	}
	cmd.Flags().StringVar(&opts.urlPath, "path", "", "Path to send request. For example /api/kafkas_mgmt/v1/kafkas?async=true")
	cmd.Flags().StringVar(&opts.method, "method", "GET", "HTTP method to use. (get, post)")
	return cmd
}

func runCmd(opts *options) (err error) {
	f := opts.factory
	if opts.urlPath == "" {
		return errors.New("--path is required")
	}
	f.Logger.Info("Performing request to", opts.urlPath)
	conn, err := f.Connection(connection.DefaultConfigSkipMasAuth)

	if err != nil {
		return err
	}

	var data interface{}
	var response interface{}
	if opts.method == "post" {
		f.Logger.Info("POST request. Reading file from standard input")
		specifiedFile, err1 := util.CreateFileFromStdin()
		if err1 != nil {
			return err
		}
		data, response, err = conn.API().GenericAPI().POST(f.Context, opts.urlPath, specifiedFile)
	} else {
		data, response, err = conn.API().GenericAPI().GET(f.Context, opts.urlPath)
	}

	if err != nil || data == nil {
		f.Logger.Info("Fetching data failed", err, response)
		return err
	}

	fmt.Fprint(f.IOStreams.Out, data)
	return nil
}
