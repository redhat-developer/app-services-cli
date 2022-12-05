package update

import (
	"encoding/json"

	"github.com/redhat-developer/app-services-cli/pkg/core/cmdutil/flagutil"
	"github.com/redhat-developer/app-services-cli/pkg/shared/connectorutil"
	"github.com/redhat-developer/app-services-cli/pkg/shared/contextutil"
	"github.com/redhat-developer/app-services-cli/pkg/shared/kafkautil"
	"github.com/spf13/cobra"

	"github.com/redhat-developer/app-services-cli/pkg/core/ioutil/dump"
	"github.com/redhat-developer/app-services-cli/pkg/core/ioutil/editor"
	"github.com/redhat-developer/app-services-cli/pkg/core/localize"
	connectormgmtclient "github.com/redhat-developer/app-services-sdk-go/connectormgmt/apiv1/client"
	connectorerror "github.com/redhat-developer/app-services-sdk-go/connectormgmt/apiv1/error"

	"github.com/redhat-developer/app-services-cli/pkg/shared/factory"
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
		Hidden:  false,
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
	flags.StringVar(&opts.namespaceID, "namespace-id", "", f.Localizer.MustLocalize("connector.flag.namespaceID.description"))
	flags.StringVar(&opts.kafkaID, "kafka-id", "", f.Localizer.MustLocalize("connector.flag.kafkaID.description"))
	flags.AddOutput(&opts.outputFormat)

	return cmd
}

// nolint:funlen
func runUpdate(opts *options) error {

	conn, err := opts.f.Connection()
	if err != nil {
		return err
	}

	api := conn.API()

	connector, err := contextutil.GetCurrentConnectorInstance(&conn, opts.f)
	if err != nil || connector == nil {
		if connector, err = connectorutil.InteractiveSelect(conn, opts.f); err != nil {
			return err
		}
	}

	connectorChanged := false
	if opts.name != "" {
		connector.SetName(opts.name)
		connectorChanged = true
	}
	if opts.namespaceID != "" {
		connector.SetNamespaceId(opts.namespaceID)
		connectorChanged = true
	}
	if opts.kafkaID != "" {
		kafkaInstance, _, kafkaErr := kafkautil.GetKafkaByID(opts.f.Context, api.KafkaMgmt(), opts.kafkaID)
		if kafkaErr != nil {
			return kafkaErr
		}
		connector.Kafka.SetId(kafkaInstance.GetId())
		connector.Kafka.SetUrl(kafkaInstance.GetBootstrapServerHost())
		connectorChanged = true
	}

	if !connectorChanged {
		opts.f.Logger.Info(opts.f.Localizer.MustLocalize("connector.update.info.editor.open"))
		err = runEditor(connector)
		if err != nil {
			return err
		}
	}

	var patchData map[string]interface{}
	data, err := json.Marshal(connector)
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, &patchData)
	if err != nil {
		return err
	}

	a := api.ConnectorsMgmt().ConnectorsApi.PatchConnector(opts.f.Context, connector.GetId())
	a = a.Body(patchData)
	updated, httpRes, err := a.Execute()
	if httpRes != nil {
		defer httpRes.Body.Close()
	}

	if apiErr := connectorerror.GetAPIError(err); apiErr != nil {

		if apiErr.GetCode() == connectorerror.ERROR_11 {
			return opts.f.Localizer.MustLocalizeError("connector.update.error.authTokenInvalid")
		}

		if apiErr.GetCode() == connectorerror.ERROR_7 {
			return opts.f.Localizer.MustLocalizeError("connector.update.error.noMatchingResource")
		}

		if apiErr.GetCode() == connectorerror.ERROR_25 {
			return opts.f.Localizer.MustLocalizeError("connector.update.error.doesNotExistAnymore")
		}

		if apiErr.GetCode() == connectorerror.ERROR_9 {
			return opts.f.Localizer.MustLocalizeError("connector.update.error.unexpectedError")
		}

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

func runEditor(currentMetadata *connectormgmtclient.Connector) error {
	metadataJson, err := json.MarshalIndent(currentMetadata, "", " ")
	if err != nil {
		return err
	}
	systemEditor := editor.New(metadataJson, "metadata.json")
	output, err := systemEditor.Run()
	if err != nil {
		return err
	}

	var resultData *connectormgmtclient.Connector
	err = json.Unmarshal(output, &resultData)
	if err != nil {
		return err
	}
	*currentMetadata = *resultData
	return nil
}
