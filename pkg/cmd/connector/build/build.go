package build

import (
	"encoding/json"
	"os"

	"github.com/redhat-developer/app-services-cli/pkg/core/cmdutil/flagutil"
	"github.com/redhat-developer/app-services-cli/pkg/core/ioutil/dump"
	"github.com/redhat-developer/app-services-cli/pkg/core/localize"
	connectormgmtclient "github.com/redhat-developer/app-services-sdk-go/connectormgmt/apiv1/client"
	connectorerror "github.com/redhat-developer/app-services-sdk-go/connectormgmt/apiv1/error"
	"github.com/wtrocki/survey-json-schema/pkg/surveyjson"

	"github.com/redhat-developer/app-services-cli/pkg/shared/connection"
	"github.com/redhat-developer/app-services-cli/pkg/shared/factory"

	"github.com/spf13/cobra"
)

type options struct {
	outputFile    string
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
	flags.StringVar(&opts.outputFile, "output-file", "", f.Localizer.MustLocalize("connector.build.file.flag.description"))
	flags.StringVar(&opts.connectorType, "type", "", f.Localizer.MustLocalize("connector.build.type.flag.description"))
	flags.StringVar(&opts.name, "name", "", f.Localizer.MustLocalize("connector.build.name.flag.description"))
	flags.AddOutput(&opts.outputFormat)

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

	if opts.connectorType == "" {
		return opts.f.Localizer.MustLocalizeError("connector.type.error.notFound", localize.NewEntry("Id", opts.connectorType))
	}

	api := conn.API()

	request := api.ConnectorsMgmt().ConnectorTypesApi.GetConnectorTypeByID(f.Context, opts.connectorType)
	response, _, err := request.Execute()

	if apiErr := connectorerror.GetAPIError(err); apiErr != nil {
		switch apiErr.GetCode() {
		case connectorerror.ERROR_7:
			return opts.f.Localizer.MustLocalizeError("connector.type.error.notFound", localize.NewEntry("Id", opts.connectorType))
		}
	}
	if err != nil {
		return err
	}

	opts.f.Logger.Info(opts.f.Localizer.MustLocalize("connector.build.info.msg"))

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

	var connectorSpecification map[string]interface{}
	err = json.Unmarshal(result, &connectorSpecification)
	if err != nil {
		return err
	}

	connector := createConnector(opts, connectorSpecification)

	if opts.outputFile != "" {
		file, err1 := os.OpenFile(opts.outputFile, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0600)
		if err1 != nil {
			return err1
		}
		defer file.Close()
		if err1 = dump.Formatted(file, opts.outputFormat, connector); err1 != nil {
			return err1
		}
	} else {
		if err = dump.Formatted(f.IOStreams.Out, opts.outputFormat, connector); err != nil {
			return err
		}
	}

	f.Logger.Info(f.Localizer.MustLocalize("connector.build.info.success"))

	return nil
}

func createConnector(opts *options, connectorSpecification map[string]interface{}) connectormgmtclient.ConnectorRequest {
	connectorChannel := connectormgmtclient.Channel(connectormgmtclient.CHANNEL_STABLE)
	connector := connectormgmtclient.ConnectorRequest{
		Name:            opts.name,
		Channel:         &connectorChannel,
		ConnectorTypeId: opts.connectorType,
		DesiredState:    connectormgmtclient.CONNECTORDESIREDSTATE_READY,
		NamespaceId:     "",
		ServiceAccount:  *connectormgmtclient.NewServiceAccount("", ""),
		Kafka: connectormgmtclient.KafkaConnectionSettings{
			Id:  "",
			Url: "",
		},
		SchemaRegistry: &connectormgmtclient.SchemaRegistryConnectionSettings{
			Id:  "",
			Url: "",
		},
		Connector: connectorSpecification,
	}
	return connector
}
