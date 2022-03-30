package generate

import (
	"fmt"
	"time"

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
	TokenURL     string
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

	cfg, err := opts.Config.Load()
	if err != nil {
		return err
	}

	configurations := &configValues{}

	var serviceAvailable bool

	if svcConfig.KafkaID != "" {
		kafkaInstance, newErr := contextutil.GetCurrentKafkaInstance(factory)
		if newErr != nil {
			return newErr
		}

		serviceAvailable = true
		configurations.KafkaURL = kafkaInstance.GetBootstrapServerHost()
	}

	if svcConfig.ServiceRegistryID != "" {
		registryInstance, newErr := contextutil.GetCurrentRegistryInstance(factory)
		if newErr != nil {
			return newErr
		}

		serviceAvailable = true
		configurations.RegistryURL = registryInstance.GetRegistryUrl()
	}

	if !serviceAvailable {
		return opts.localizer.MustLocalizeError("generate.log.info.noSevices")
	}

	serviceAccount, err := createServiceAccount(opts, fmt.Sprintf("%s-%v", opts.name, time.Now().Unix()))
	if err != nil {
		return err
	}

	opts.Logger.Info(
		icon.SuccessPrefix(),
		opts.localizer.MustLocalize("serviceAccount.create.log.info.createdSuccessfully", localize.NewEntry("ID", serviceAccount.GetId())),
	)

	configurations.ClientID = serviceAccount.GetClientId()
	configurations.ClientSecret = serviceAccount.GetClientSecret()
	configurations.TokenURL = cfg.MasAuthURL + "/protocol/openid-connect/token"

	if err = WriteConfig(opts.configType, configurations); err != nil {
		return err
	}

	opts.Logger.Info(icon.SuccessPrefix(), opts.localizer.MustLocalize("generate.log.info.credentialsSaved"))

	return nil
}
