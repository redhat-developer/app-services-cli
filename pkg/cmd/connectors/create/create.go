package create

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/redhat-developer/app-services-cli/pkg/cmd/registry/artifact/util"
	"github.com/redhat-developer/app-services-cli/pkg/core/cmdutil/flagutil"
	"github.com/redhat-developer/app-services-cli/pkg/core/ioutil/dump"
	"github.com/redhat-developer/app-services-cli/pkg/shared/connection"
	"github.com/redhat-developer/app-services-cli/pkg/shared/factory"
	connectormgmtclient "github.com/redhat-developer/app-services-sdk-go/connectormgmt/apiv1/client"

	"github.com/spf13/cobra"
)

type options struct {
	file string

	outputFormat string
	f            *factory.Factory
	interactive  bool
}

func NewCreateCommand(f *factory.Factory) *cobra.Command {
	opts := &options{
		f: f,
	}

	cmd := &cobra.Command{
		Use:     "create",
		Short:   f.Localizer.MustLocalize("connector.create.cmd.shortDescription"),
		Long:    f.Localizer.MustLocalize("connector.create.cmd.longDescription"),
		Example: f.Localizer.MustLocalize("connector.create.cmd.example"),
		Args:    cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {

			validOutputFormats := flagutil.ValidOutputFormats
			if opts.outputFormat != "" && !flagutil.IsValidInput(opts.outputFormat, validOutputFormats...) {
				return flagutil.InvalidValueError("output", opts.outputFormat, validOutputFormats...)
			}

			return runCreate(opts)
		},
	}
	flags := flagutil.NewFlagSet(cmd, f.Localizer)
	flags.StringVar(&opts.file, "file", "", f.Localizer.MustLocalize("connector.file.flag.description"))
	flags.AddOutput(&opts.outputFormat)

	return cmd
}

func runCreate(opts *options) error {
	f := opts.f

	var conn connection.Connection
	conn, err := f.Connection(connection.DefaultConfigSkipMasAuth)
	if err != nil {
		return err
	}

	if opts.interactive {
		// TODO
	}
	var specifiedFile *os.File
	if opts.file == "" {
		file, err1 := util.CreateFileFromStdin()
		if err1 != nil {
			return err
		}
		specifiedFile = file
	} else {
		if util.IsURL(opts.file) {
			specifiedFile, err = util.GetContentFromFileURL(f.Context, opts.file)
		} else {
			specifiedFile, err = os.Open(opts.file)
		}
		if err != nil {
			return err
		}
	}
	defer specifiedFile.Close()

	byteValue, err := ioutil.ReadAll(specifiedFile)
	if err != nil {
		return err
	}
	var connector connectormgmtclient.ConnectorRequest
	err = json.Unmarshal(byteValue, &connector)
	if err != nil {
		return err
	}

	api := conn.API()

	a := api.ConnectorsMgmt().ConnectorsApi.CreateConnector(f.Context)
	a = a.ConnectorRequest(connector)
	a = a.Async(true)

	response, httpRes, err := a.Execute()
	if httpRes != nil {
		defer httpRes.Body.Close()
	}

	if err != nil {
		return err
	}

	if err = dump.Formatted(f.IOStreams.Out, opts.outputFormat, response); err != nil {
		return err
	}

	f.Logger.Info(f.Localizer.MustLocalize("connectors.create.info.success"))

	return nil
}
