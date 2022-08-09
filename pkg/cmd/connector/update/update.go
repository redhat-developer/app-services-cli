package update

import (
	"encoding/json"

	"github.com/redhat-developer/app-services-cli/pkg/core/cmdutil/flagutil"
	"github.com/redhat-developer/app-services-cli/pkg/shared/connectorutil"
	"github.com/redhat-developer/app-services-cli/pkg/shared/contextutil"
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
		connector.Kafka.SetId(opts.kafkaID)
		connectorChanged = true
	}

	if !connectorChanged {
		err := runEditor(connector)
		if err != nil {
			return err
		}
	}
	var patchData map[string]interface{}
	data, _ := json.Marshal(connector)
	json.Unmarshal(data, &patchData)

	a := api.ConnectorsMgmt().ConnectorsApi.PatchConnector(opts.f.Context, connector.GetId())
	a = a.Body(patchData)
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
	currentMetadata = resultData
	return nil
}
