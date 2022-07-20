package generate

import (
	"fmt"
	"time"

	"github.com/redhat-developer/app-services-cli/pkg/cmd/serviceaccount/svcaccountcmdutil"
	"github.com/redhat-developer/app-services-cli/pkg/core/ioutil/icon"
	"github.com/redhat-developer/app-services-cli/pkg/core/localize"
	"github.com/redhat-developer/app-services-cli/pkg/core/servicecontext"
	"github.com/redhat-developer/app-services-cli/pkg/shared/contextutil"
	"github.com/redhat-developer/app-services-cli/pkg/shared/factory"

	kafkamgmtclient "github.com/redhat-developer/app-services-sdk-go/kafkamgmt/apiv1/client"
)

type configValues struct {
	KafkaHost    string
	RegistryURL  string
	ClientID     string
	ClientSecret string
	TokenURL     string

	// Optional
	Name string
}

func createServiceAccount(opts *options, shortDescription string) (*kafkamgmtclient.ServiceAccount, error) {
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
		kafkaInstance, newErr := contextutil.GetCurrentKafkaInstance(factory)
		if newErr != nil {
			return newErr
		}

		serviceAvailable = true
		configurations.KafkaHost = kafkaInstance.GetBootstrapServerHost()
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
	configInstanceName := fmt.Sprintf("%s-%v", opts.name, time.Now().Unix())
	serviceAccount, err := createServiceAccount(opts, configInstanceName)
	if err != nil {
		return err
	}

	opts.Logger.Info(
		icon.SuccessPrefix(),
		opts.localizer.MustLocalize("serviceAccount.create.log.info.createdSuccessfully", localize.NewEntry("ID", serviceAccount.GetId())),
	)

	conn, err := opts.Connection()
	if err != nil {
		return err
	}

	providerUrls, err := svcaccountcmdutil.GetProvidersDetails(conn, opts.Context)
	if err != nil {
		return err
	}

	configurations.ClientID = serviceAccount.GetClientId()
	configurations.ClientSecret = serviceAccount.GetClientSecret()
	configurations.TokenURL = providerUrls.GetTokenUrl()
	configurations.Name = configInstanceName

	var fileName string
	if fileName, err = WriteConfig(opts, configurations); err != nil {
		return err
	}

	opts.Logger.Info(icon.SuccessPrefix(), opts.localizer.MustLocalize("generate.log.info.credentialsSaved", localize.NewEntry("FilePath", fileName)))

	return nil
}
