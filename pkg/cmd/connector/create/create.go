package create

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/redhat-developer/app-services-cli/pkg/cmd/registry/artifact/util"
	"github.com/redhat-developer/app-services-cli/pkg/core/cmdutil/flagutil"
	"github.com/redhat-developer/app-services-cli/pkg/core/ioutil/dump"
	"github.com/redhat-developer/app-services-cli/pkg/core/localize"

	"github.com/redhat-developer/app-services-cli/pkg/shared/connection"
	"github.com/redhat-developer/app-services-cli/pkg/shared/factory"
	connectormgmtclient "github.com/redhat-developer/app-services-sdk-go/connectormgmt/apiv1/client"

	kafkamgmtclient "github.com/redhat-developer/app-services-sdk-go/kafkamgmt/apiv1/client"
	"github.com/spf13/cobra"
)

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

	var specifiedFile *os.File
	var err error
	if opts.file == "" {
		opts.f.Logger.Info(opts.f.Localizer.MustLocalize("common.message.reading.file"))
		file, newErr := util.CreateFileFromStdin()
		if newErr != nil {
			opts.f.Localizer.MustLocalizeError("common.message.reading.file.error", localize.NewEntry("ErrorMessage", err))
		}
		specifiedFile = file
	} else {
		if util.IsURL(opts.file) {
			specifiedFile, err = util.GetContentFromFileURL(f.Context, opts.file)
		} else {
			specifiedFile, err = os.Open(opts.file)
		}
		if err != nil {
			return opts.f.Localizer.MustLocalizeError("common.message.reading.file.error", localize.NewEntry("ErrorMessage", err))
		}
	}
	defer specifiedFile.Close()
	byteValue, err := ioutil.ReadAll(specifiedFile)
	if err != nil {
		return opts.f.Localizer.MustLocalizeError("common.message.reading.file.error", localize.NewEntry("ErrorMessage", err))
	}

	var connector connectormgmtclient.ConnectorRequest
	err = json.Unmarshal(byteValue, &connector)
	if err != nil {
		return opts.f.Localizer.MustLocalizeError("common.message.reading.file.error", localize.NewEntry("ErrorMessage", err))
	}

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

	serviceAccount, err := createServiceAccount(opts.f, "connector-"+connector.Name)
	if err != nil {
		return err
	}

	connector.ServiceAccount = *connectormgmtclient.NewServiceAccount(serviceAccount.GetClientId(), serviceAccount.GetClientSecret())

	var conn connection.Connection
	conn, err = f.Connection()
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
