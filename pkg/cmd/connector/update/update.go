package update

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/redhat-developer/app-services-cli/pkg/cmd/connector/connectorcmdutil"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/registry/artifact/util"
	"github.com/redhat-developer/app-services-cli/pkg/core/cmdutil/flagutil"
	"github.com/redhat-developer/app-services-cli/pkg/core/ioutil/dump"
	"github.com/redhat-developer/app-services-cli/pkg/shared/connection"
	"github.com/redhat-developer/app-services-cli/pkg/shared/factory"
	"github.com/spf13/cobra"
)

type options struct {
	file string

	outputFormat string
	f            *factory.Factory
	id           string
}

// NewUpdateCommand creates a new command to update a connector
func NewUpdateCommand(f *factory.Factory) *cobra.Command {

	opts := &options{
		f: f,
	}

	cmd := &cobra.Command{
		Use:     "update",
		Short:   f.Localizer.MustLocalize("connector.update.cmd.shortDescription"),
		Long:    f.Localizer.MustLocalize("connector.update.cmd.longDescription"),
		Example: f.Localizer.MustLocalize("connector.update.cmd.example"),
		Hidden:  true,
		Args:    cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {

			validOutputFormats := flagutil.ValidOutputFormats
			if opts.outputFormat != "" && !flagutil.IsValidInput(opts.outputFormat, validOutputFormats...) {
				return flagutil.InvalidValueError("output", opts.outputFormat, validOutputFormats...)
			}

			return runUpdateCommand(opts)
		},
	}

	flags := connectorcmdutil.NewFlagSet(cmd, f)
	flags.StringVar(&opts.file, "file", "", f.Localizer.MustLocalize("connector.file.flag.description"))
	flags.AddConnectorID(&opts.id)
	flags.AddOutput(&opts.outputFormat)

	return cmd

}

func runUpdateCommand(opts *options) error {
	f := opts.f

	var conn connection.Connection
	conn, err := f.Connection(connection.DefaultConfigSkipMasAuth)
	if err != nil {
		return err
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
	var connector map[string]interface{}
	err = json.Unmarshal(byteValue, &connector)
	if err != nil {
		return err
	}

	api := conn.API()

	a := api.ConnectorsMgmt().ConnectorsApi.PatchConnector(f.Context, opts.id)
	a = a.Body(connector)

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

	f.Logger.Info(f.Localizer.MustLocalize("connector.update.info.success"))

	return nil
}
