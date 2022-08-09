package update

import (
	"encoding/json"
	"fmt"
	"github.com/jackdelahunt/survey-json-schema/pkg/surveyjson"
	"github.com/redhat-developer/app-services-cli/pkg/core/cmdutil/flagutil"
	"github.com/redhat-developer/app-services-cli/pkg/shared/connection"
	"github.com/redhat-developer/app-services-cli/pkg/shared/connection/api"
	"github.com/redhat-developer/app-services-cli/pkg/shared/connectorutil"
	"github.com/redhat-developer/app-services-cli/pkg/shared/contextutil"
	"github.com/redhat-developer/app-services-cli/pkg/shared/kafkautil"
	"github.com/spf13/cobra"
	"os"

	"github.com/redhat-developer/app-services-cli/pkg/core/ioutil/dump"
	"github.com/redhat-developer/app-services-cli/pkg/core/localize"
	connectorerror "github.com/redhat-developer/app-services-sdk-go/connectormgmt/apiv1/error"
	"gopkg.in/AlecAivazis/survey.v1"

	"github.com/redhat-developer/app-services-cli/pkg/shared/factory"
	connectormgmtclient "github.com/redhat-developer/app-services-sdk-go/connectormgmt/apiv1/client"
)

type options struct {
	namespaceID string
	kafkaID     string
	name        string

	outputFormat string
	f            *factory.Factory
}

func NewUpdateCommand(f *factory.Factory) *cobra.Command {
	opts := &options{
		f: f,
	}

	cmd := &cobra.Command{
		Use:     "update",
		Short:   f.Localizer.MustLocalize("connector.update.cmd.shortDescription"),
		Long:    f.Localizer.MustLocalize("connector.update.cmd.longDescription"),
		Example: f.Localizer.MustLocalize("connector.update.cmd.example"),
		Args:    cobra.NoArgs,
		Hidden:  true,
		RunE: func(cmd *cobra.Command, args []string) error {
			validOutputFormats := flagutil.ValidOutputFormats
			if opts.outputFormat != "" && !flagutil.IsValidInput(opts.outputFormat, validOutputFormats...) {
				return flagutil.InvalidValueError("output", opts.outputFormat, validOutputFormats...)
			}

			return runUpdate(opts)
		},
	}
	flags := flagutil.NewFlagSet(cmd, f.Localizer)
	flags.StringVar(&opts.name, "name", "", f.Localizer.MustLocalize("connector.flag.name.description"))
	flags.StringVar(&opts.namespaceID, "namespace-id", "", f.Localizer.MustLocalize("connector.flag.kafkaID.description"))
	flags.StringVar(&opts.kafkaID, "kafka-id", "", f.Localizer.MustLocalize("connector.flag.namespaceID.description"))
	flags.AddOutput(&opts.outputFormat)

	return cmd
}

func runUpdate(opts *options) error {

	conn, err1 := opts.f.Connection()
	if err1 != nil {
		return err1
	}

	api := conn.API()

	connector, err2 := contextutil.GetCurrentConnectorInstance(&conn, opts.f)
	if err2 != nil {
		connector, _ = connectorutil.InteractiveSelect(conn, opts.f)
	}

	patch, err3 := setValuesInInteractiveMode(connector, opts, &conn)
	if err3 != nil {
		return err3
	}

	connectorSpecification, err4 := askForConnectorDetails(&api, connector, opts)
	if err4 != nil {
		return err4
	}

	patch["connector"] = connectorSpecification

	a := api.ConnectorsMgmt().ConnectorsApi.PatchConnector(opts.f.Context, connector.GetId())
	a = a.Body(patch)
	updated, httpRes, err := a.Execute()
	if httpRes != nil {
		defer httpRes.Body.Close()
	}

	if apiErr := connectorerror.GetAPIError(err); apiErr != nil {
		return opts.f.Localizer.MustLocalizeError("connector.type.update.error.other", localize.NewEntry("Error", apiErr.GetReason()))
	}

	if err != nil {
		return err
	}

	if err = dump.Formatted(opts.f.IOStreams.Out, opts.outputFormat, updated); err != nil {
		return err
	}

	opts.f.Logger.Info(opts.f.Localizer.MustLocalize("connector.update.info.success"))

	return nil
}

func setValuesInInteractiveMode(connector *connectormgmtclient.Connector, opts *options, conn *connection.Connection) (map[string]interface{}, error) {

	patch := make(map[string]interface{})

	if opts.name == "" {
		name, err := askForValue("Name", connector.GetName(), opts)
		if err != nil {
			return nil, err
		}

		opts.name = name
	}
	patch["name"] = opts.name

	if opts.kafkaID == "" {
		if opts.f.IOStreams.CanPrompt() {
			kafka, err := kafkautil.InteractiveSelect(opts.f.Context, *conn, opts.f.Logger, opts.f.Localizer)
			if err != nil {
				return nil, err
			}

			patch["kafka"] = *kafka
		} else {
			return nil, opts.f.Localizer.MustLocalizeError("connector.interactive.error", localize.NewEntry("Field", "kafka.id"))
		}
	} else {
		kafka, _, err := kafkautil.GetKafkaByID(opts.f.Context, (*conn).API().KafkaMgmt(), opts.kafkaID)
		if err != nil {
			return nil, err
		}

		patch["kafka"] = *kafka
	}

	if opts.namespaceID == "" {
		namespaceId, err := askForValue("Namespace ID", connector.GetNamespaceId(), opts)
		if err != nil {
			return nil, err
		}
		opts.namespaceID = namespaceId
	}

	patch["namespace_id"] = opts.namespaceID

	for key, value := range connector.Connector {
		result, err := askForValue(key, fmt.Sprint(value), opts)
		if err != nil {
			return nil, err
		}

		connector.Connector[key] = result
	}
	patch["connector"] = connector.Connector

	return patch, nil
}

func askForValue(field string, current string, opts *options) (string, error) {

	if !opts.f.IOStreams.CanPrompt() {
		return "", opts.f.Localizer.MustLocalizeError("connector.interactive.error", localize.NewEntry("Field", field))
	}

	var value string
	prompt := &survey.Input{
		Message: field,
		Default: current,
	}
	err := survey.AskOne(prompt, &value, nil)
	if err != nil {
		return value, err
	}
	return value, nil
}

func askForConnectorDetails(api *api.API, connector *connectormgmtclient.Connector, opts *options) (*map[string]interface{}, error) {
	response, _, err := (*api).ConnectorsMgmt().ConnectorTypesApi.GetConnectorTypeByID(opts.f.Context, connector.GetConnectorTypeId()).Execute()
	if err != nil {
		return nil, err
	}

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
		return nil, err
	}

	connectorValues, err := schemaOptions.GenerateValues(schemaBytes, initialValues)
	if err != nil {
		return nil, err
	}

	var connectorSpecification map[string]interface{}
	err = json.Unmarshal(connectorValues, &connectorSpecification)
	if err != nil {
		return nil, err
	}

	return &connectorSpecification, nil
}
