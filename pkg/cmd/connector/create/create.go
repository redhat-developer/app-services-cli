package create

import (
	"encoding/json"
	"io/ioutil"
	"os"

	// embed static HTML file
	_ "embed"

	"github.com/pkg/errors"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/registry/artifact/util"
	"github.com/redhat-developer/app-services-cli/pkg/core/cmdutil/flagutil"
	"github.com/redhat-developer/app-services-cli/pkg/core/ioutil/dump"
	"github.com/redhat-developer/app-services-cli/pkg/core/localize"
	connectorerror "github.com/redhat-developer/app-services-sdk-go/connectormgmt/apiv1/error"
	"github.com/wtrocki/survey-json-schema/pkg/surveyjson"

	"github.com/redhat-developer/app-services-cli/pkg/shared/connection"
	"github.com/redhat-developer/app-services-cli/pkg/shared/factory"
	connectormgmtclient "github.com/redhat-developer/app-services-sdk-go/connectormgmt/apiv1/client"

	kafkamgmtclient "github.com/redhat-developer/app-services-sdk-go/kafkamgmt/apiv1/client"
	"github.com/spf13/cobra"
)

//go:embed schema.json
var requestSchema []byte

type options struct {
	file           string
	kafkaId        string
	namespace      string
	name           string
	outputFormat   string
	serviceAccount bool
	f              *factory.Factory
}

// NewCreateCommand creates a new command to create a Connector
func NewCreateCommand(f *factory.Factory) *cobra.Command {
	opts := &options{
		f: f,
	}

	cmd := &cobra.Command{
		Use:     "create",
		Short:   f.Localizer.MustLocalize("connector.create.cmd.shortDescription"),
		Long:    f.Localizer.MustLocalize("connector.create.cmd.longDescription"),
		Example: f.Localizer.MustLocalize("connector.create.cmd.example"),
		Hidden:  true,
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
	flags.StringVarP(&opts.file, "file", "", "f", f.Localizer.MustLocalize("connector.file.flag.description"))
	flags.StringVar(&opts.kafkaId, "kafka", "", f.Localizer.MustLocalize("connector.flag.kafka.description"))
	flags.StringVar(&opts.namespace, "namespace", "", f.Localizer.MustLocalize("connector.flag.namespace.description"))
	flags.StringVar(&opts.name, "name", "", f.Localizer.MustLocalize("connector.flag.name.description"))
	flags.BoolVar(&opts.serviceAccount, "create-service-account", false, f.Localizer.MustLocalize("connector.flag.sa.description"))
	flags.AddOutput(&opts.outputFormat)

	return cmd
}

func runCreate(opts *options) error {
	f := opts.f
	// Load the connector from the file
	fileContent, err := readFileFromInput(opts)
	if err != nil {
		return err
	}

	var userConnector connectormgmtclient.ConnectorRequest
	err = json.Unmarshal(fileContent, &userConnector)
	if err != nil {
		return errors.Wrap(err, opts.f.Localizer.MustLocalize("connector.message.reading.file.error"))
	}

	err = setDefaultValuesFromFlags(&userConnector, opts)
	if err != nil {
		return err
	}
	var connectorRequest connectormgmtclient.ConnectorRequest
	//if opts.f.IOStreams.CanPrompt() {
	connectorRequest, err = setValuesUsingInteractiveMode(&userConnector, opts)
	if err != nil {
		return err
	}
	// } else {
	// 	connectorRequest = userConnector
	// }

	var conn connection.Connection
	conn, err = f.Connection()
	if err != nil {
		return err
	}

	api := conn.API()

	a := api.ConnectorsMgmt().ConnectorsApi.CreateConnector(f.Context)
	a = a.ConnectorRequest(connectorRequest)
	a = a.Async(true)

	response, httpRes, err := a.Execute()
	if httpRes != nil {
		defer httpRes.Body.Close()
	}

	if apiErr := connectorerror.GetAPIError(err); apiErr != nil {
		if apiErr.GetCode() == connectorerror.ERROR_7 {
			return opts.f.Localizer.MustLocalizeError("connector.type.error.notFound", localize.NewEntry("Id", connectorRequest.ConnectorTypeId))
		}

		return opts.f.Localizer.MustLocalizeError("connector.type.create.error.other", localize.NewEntry("Error", apiErr.GetReason()))
	}

	if err != nil {
		return err
	}

	if err = dump.Formatted(f.IOStreams.Out, opts.outputFormat, response); err != nil {
		return err
	}

	f.Logger.Info(f.Localizer.MustLocalize("connector.create.info.success"))

	return nil
}

func createServiceAccount(opts *factory.Factory, shortDescription string) (*kafkamgmtclient.ServiceAccount, error) {
	conn, err := opts.Connection()
	if err != nil {
		return nil, err
	}
	serviceAccountPayload := kafkamgmtclient.ServiceAccountRequest{Name: shortDescription}

	serviceacct, httpRes, err := conn.API().
		ServiceAccountMgmt().
		CreateServiceAccount(opts.Context).
		ServiceAccountRequest(serviceAccountPayload).
		Execute()

	if httpRes != nil {
		defer httpRes.Body.Close()
	}

	if err != nil {
		return nil, err
	}
	opts.Logger.Info(opts.Localizer.MustLocalize("connector.sa.created",
		localize.NewEntry("ClientId", serviceacct.ClientId), localize.NewEntry("ClientSecret", serviceacct.ClientSecret)))
	return &serviceacct, nil
}

func setValuesUsingInteractiveMode(connector *connectormgmtclient.ConnectorRequest, opts *options) (connectormgmtclient.ConnectorRequest, error) {
	schemaOptions := surveyjson.JSONSchemaOptions{
		Out:                 os.Stdout,
		In:                  os.Stdin,
		OutErr:              os.Stderr,
		AskExisting:         false,
		AutoAcceptDefaults:  true,
		NoAsk:               false,
		IgnoreMissingValues: true,
	}

	// In future library will be able to
	// pick those values automatically
	initialValues := map[string]interface{}{
		"name":         connector.Name,
		"namespace_id": connector.NamespaceId,
		//"kafka":        connector.Kafka,
	}

	result, err := schemaOptions.GenerateValues(requestSchema, initialValues)
	if err != nil {
		return connectormgmtclient.ConnectorRequest{}, err
	}
	var connectorObj connectormgmtclient.ConnectorRequest
	err = json.Unmarshal(result, &connectorObj)
	if err != nil {
		return connectormgmtclient.ConnectorRequest{}, err
	}
	// We need to map values to the connector spec
	connectorObj.Connector = connector.Connector
	return connectorObj, nil
}

func setDefaultValuesFromFlags(connector *connectormgmtclient.ConnectorRequest, opts *options) error {
	if opts.kafkaId != "" {
		connector.Kafka = connectormgmtclient.KafkaConnectionSettings{
			Id: opts.kafkaId,
		}
	}

	if opts.namespace != "" {
		connector.NamespaceId = opts.namespace
	}

	if opts.name != "" {
		connector.Name = opts.name
	}

	if opts.serviceAccount {
		serviceAccount, err1 := createServiceAccount(opts.f, "connector-"+connector.Name)
		if err1 != nil {
			return err1
		}
		connector.ServiceAccount = *connectormgmtclient.NewServiceAccount(serviceAccount.GetClientId(), serviceAccount.GetClientSecret())
	}
	return nil
}

func readFileFromInput(opts *options) ([]byte, error) {
	var specifiedFile *os.File
	var err error
	if opts.file == "" {
		opts.f.Logger.Info(opts.f.Localizer.MustLocalize("common.message.reading.file"))
		specifiedFile, err = util.CreateFileFromStdin()
		if err != nil {
			return nil, errors.Wrap(err, opts.f.Localizer.MustLocalize("connector.message.reading.file.error"))
		}
	} else {
		if util.IsURL(opts.file) {
			specifiedFile, err = util.GetContentFromFileURL(opts.f.Context, opts.file)
		} else {
			specifiedFile, err = os.Open(opts.file)
		}
		if err != nil {
			return nil, errors.Wrap(err, opts.f.Localizer.MustLocalize("connector.message.reading.file.error"))
		}
	}
	defer specifiedFile.Close()
	byteValue, err := ioutil.ReadAll(specifiedFile)
	if err != nil {
		return nil, errors.Wrap(err, opts.f.Localizer.MustLocalize("connector.message.reading.file.error"))
	}
	return byteValue, nil
}
