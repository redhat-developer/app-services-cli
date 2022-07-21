package build

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/redhat-developer/app-services-cli/pkg/core/cmdutil/flagutil"
	"github.com/wtrocki/survey-json-schema/pkg/surveyjson"

	"github.com/redhat-developer/app-services-cli/pkg/shared/connection"
	"github.com/redhat-developer/app-services-cli/pkg/shared/factory"

	"github.com/spf13/cobra"
)

type options struct {
	name          string
	connectorType string
	outputFormat  string
	f             *factory.Factory
}

// NewBuildCommand builds a new command to build a Connector
func NewBuildCommand(f *factory.Factory) *cobra.Command {
	opts := &options{
		f: f,
	}

	cmd := &cobra.Command{
		Use:     "build",
		Short:   f.Localizer.MustLocalize("connector.build.cmd.shortDescription"),
		Long:    f.Localizer.MustLocalize("connector.build.cmd.longDescription"),
		Example: f.Localizer.MustLocalize("connector.build.cmd.example"),
		Hidden:  true,
		Args:    cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			validOutputFormats := flagutil.ValidOutputFormats
			if opts.outputFormat != "" && !flagutil.IsValidInput(opts.outputFormat, validOutputFormats...) {
				return flagutil.InvalidValueError("output", opts.outputFormat, validOutputFormats...)
			}

			return runBuild(opts)
		},
	}
	flags := flagutil.NewFlagSet(cmd, f.Localizer)
	flags.StringVar(&opts.name, "name", "", f.Localizer.MustLocalize("connector.build.name.flag.description"))
	flags.StringVar(&opts.connectorType, "type", "", f.Localizer.MustLocalize("connector.build.type.flag.description"))
	// TODO connector type suggestions
	return cmd
}

func runBuild(opts *options) error {
	f := opts.f

	var conn connection.Connection
	conn, err := f.Connection()
	if err != nil {
		return err
	}

	if opts.name == "" {
		opts.name = "connector.json"
	}

	api := conn.API()
	if opts.connectorType == "" {
		return fmt.Errorf("Interactive mode not supported yet")
	}

	request := api.ConnectorsMgmt().ConnectorTypesApi.GetConnectorTypeByID(f.Context, opts.connectorType)
	response, _, err := request.Execute()
	if err != nil {
		return err
	}
	// TODO handle errors for connectors

	// Creates JSONSchema based of
	schemaOptions := surveyjson.JSONSchemaOptions{
		Out:                 os.Stdout,
		In:                  os.Stdin,
		OutErr:              os.Stderr,
		AskExisting:         true,
		AutoAcceptDefaults:  false,
		NoAsk:               false,
		IgnoreMissingValues: false,
	}

	initialValues := make(map[string]interface{})

	schemaBytes, err := json.Marshal(response.Schema)
	if err != nil {
		return err
	}

	result, err := schemaOptions.GenerateValues(schemaBytes, initialValues)
	if err != nil {
		return err
	}
	// TODO file output
	fmt.Fprint(os.Stdin, string(result))

	return nil
}
