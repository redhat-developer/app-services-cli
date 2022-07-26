package build

import (
	"encoding/json"
	"os"

	"github.com/redhat-developer/app-services-cli/pkg/cmd/connector/connectorcmdutil"
	"github.com/redhat-developer/app-services-cli/pkg/core/cmdutil/flagutil"
	"github.com/redhat-developer/app-services-cli/pkg/core/ioutil/color"
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
	overwrite     bool

	f *factory.Factory
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
			if !opts.f.IOStreams.CanPrompt() {
				return flagutil.InvalidValueError("output", opts.outputFormat, validOutputFormats...)
			}

			// If the  file already exists, and the --overwrite flag is not set then return an error
			// indicating that the user should explicitly request overwriting of the file
			if _, err := os.Stat(opts.outputFile); err == nil && !opts.overwrite {
				return opts.f.Localizer.MustLocalizeError("connector.common.error.FileAlreadyExists", localize.NewEntry("Name", color.CodeSnippet(opts.outputFile)))
			}

			return runBuild(opts)
		},
	}
	flags := flagutil.NewFlagSet(cmd, f.Localizer)
	flags.StringVar(&opts.outputFile, "output-file", "", f.Localizer.MustLocalize("connector.build.file.flag.description"))
	flags.StringVar(&opts.connectorType, "type", "", f.Localizer.MustLocalize("connector.build.type.flag.description"))
	flags.StringVar(&opts.name, "name", "", f.Localizer.MustLocalize("connector.build.name.flag.description"))
	cmd.Flags().BoolVar(&opts.overwrite, "overwrite", false, opts.f.Localizer.MustLocalize("connector.build.overwrite.flag.description"))
	flags.AddOutput(&opts.outputFormat)

	_ = cmd.MarkFlagRequired("type")

	_ = cmd.RegisterFlagCompletionFunc("type", func(cmd *cobra.Command, _ []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return connectorcmdutil.FilterValidTypesArgs(f, toComplete)
	})

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
		if apiErr.GetCode() == connectorerror.ERROR_7 {
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

	connector := createConnectorObject(opts, connectorSpecification)

	if opts.outputFile == "" {
		if opts.outputFormat == "" {
			opts.outputFile = "connector.json"
		} else {
			opts.outputFile = "connector." + opts.outputFormat
		}
	}

	file, err := os.Create(opts.outputFile)
	if err != nil {
		return err
	}
	defer file.Close()
	if err = dump.Formatted(file, opts.outputFormat, connector); err != nil {
		return err
	}

	f.Logger.Info(f.Localizer.MustLocalize("connector.build.info.success"))

	return nil
}

func createConnectorObject(opts *options, connectorSpecification map[string]interface{}) connectormgmtclient.ConnectorRequest {
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
