package generate

import (
	"github.com/redhat-developer/app-services-cli/pkg/core/ioutil/icon"
	"github.com/redhat-developer/app-services-cli/pkg/core/localize"
	"github.com/redhat-developer/app-services-cli/pkg/core/servicecontext"
	"github.com/redhat-developer/app-services-cli/pkg/shared/connection"
	"github.com/redhat-developer/app-services-cli/pkg/shared/contextutil"
	"github.com/redhat-developer/app-services-cli/pkg/shared/factory"

	kafkamgmtclient "github.com/redhat-developer/app-services-sdk-go/kafkamgmt/apiv1/client"
)

type configValues struct {
	KafkaURL     string
	RegistryURL  string
	ClientID     string
	ClientSecret string
}

func createServiceAccount(opts *options, shortDescription string) (*kafkamgmtclient.ServiceAccount, error) {

	conn, err := opts.Connection(connection.DefaultConfigSkipMasAuth)
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

	return &serviceacct, nil
}

// BuildConfiguration builds the configs for the service context
func BuildConfiguration(svcConfig *servicecontext.ServiceConfig, opts *options) error {

	factory := &factory.Factory{
		IOStreams:      opts.IO,
		Logger:         opts.Logger,
		Context:        opts.Context,
		Localizer:      opts.localizer,
		Connection:     opts.Connection,
		ServiceContext: opts.ServiceContext,
	}

	configurations := &configValues{}

	var serviceAvailable bool

	if svcConfig.KafkaID != "" {
		kafkaInstance, err := contextutil.GetCurrentKafkaInstance(factory)
		if err != nil {
			return err
		}

		serviceAvailable = true
		configurations.KafkaURL = kafkaInstance.GetBootstrapServerHost()
	}

	if svcConfig.ServiceRegistryID != "" {
		registryInstance, err := contextutil.GetCurrentRegistryInstance(factory)
		if err != nil {
			return err
		}

		serviceAvailable = true
		configurations.RegistryURL = registryInstance.GetRegistryUrl()
	}

	if !serviceAvailable {
		return opts.localizer.MustLocalizeError("context.generate.log.info.noSevices")
	}

	serviceAccount, err := createServiceAccount(opts, "to-do-generate-description")
	if err != nil {
		return err
	}

	opts.Logger.Info(
		icon.SuccessPrefix(),
		opts.localizer.MustLocalize("serviceAccount.create.log.info.createdSuccessfully", localize.NewEntry("ID", serviceAccount.GetId())),
	)

	configurations.ClientID = serviceAccount.GetClientId()
	configurations.ClientSecret = serviceAccount.GetClientSecret()

	if err = WriteConfig(opts.configType, configurations); err != nil {
		return err
	}

	opts.Logger.Info(icon.SuccessPrefix(), opts.localizer.MustLocalize("context.generate.log.info.credentialsSaved"))

	return nil
}
