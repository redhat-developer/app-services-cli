package build

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/redhat-developer/app-services-cli/pkg/core/cmdutil/flagutil"
	"github.com/redhat-developer/app-services-cli/pkg/core/localize"
	connectorerror "github.com/redhat-developer/app-services-sdk-go/connectormgmt/apiv1/error"
	"github.com/wtrocki/survey-json-schema/pkg/surveyjson"

	"github.com/redhat-developer/app-services-cli/pkg/shared/connection"
	"github.com/redhat-developer/app-services-cli/pkg/shared/factory"

	"github.com/spf13/cobra"
)

type options struct {
	outputFile    string
	connectorType string
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
			return runBuild(opts)
		},
	}
	flags := flagutil.NewFlagSet(cmd, f.Localizer)
	flags.StringVar(&opts.outputFile, "output-file", "", f.Localizer.MustLocalize("connector.build.name.flag.description"))
	flags.StringVar(&opts.connectorType, "type", "", f.Localizer.MustLocalize("connector.build.type.flag.description"))

	_ = cmd.MarkFlagRequired("type")

	return cmd
}

func runBuild(opts *options) error {
	f := opts.f

	var conn connection.Connection
	conn, err := f.Connection()
	if err != nil {
		return err
	}

	if opts.outputFile == "" {
		opts.outputFile = "connector.json"
	}

	api := conn.API()

	request := api.ConnectorsMgmt().ConnectorTypesApi.GetConnectorTypeByID(f.Context, opts.connectorType)
	response, _, err := request.Execute()

	if apiErr := connectorerror.GetAPIError(err); apiErr != nil {
		switch apiErr.GetCode() {
		case connectorerror.ERROR_7:
			return opts.f.Localizer.MustLocalizeError("connector.type.error.notFound", localize.NewEntry("Id", opts.connectorType))
		default:
			return err
		}
	}
	if err != nil {
		return err
	}

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

	if opts.outputFile != "" {
		err := os.WriteFile(opts.outputFile, result, 0600)
		if err != nil {
			return err
		}
	} else {
		// Print to stdout
		fmt.Fprintf(os.Stdout, "%v\n", string(result))
	}

	return nil
}
