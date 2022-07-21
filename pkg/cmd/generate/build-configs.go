package generate

import (
	"fmt"
	"time"

	"github.com/redhat-developer/app-services-cli/pkg/core/ioutil/icon"
	"github.com/redhat-developer/app-services-cli/pkg/core/localize"
	"github.com/redhat-developer/app-services-cli/pkg/core/servicecontext"
	"github.com/redhat-developer/app-services-cli/pkg/shared/contextutil"
	"github.com/redhat-developer/app-services-cli/pkg/shared/factory"
)

type configValues struct {
	KafkaHost   string
	RegistryURL string

	// Optional
	Name string
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
	var err error

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

	configurations.Name = configInstanceName

	var fileName string
	if fileName, err = WriteConfig(opts, configurations); err != nil {
		return err
	}

	opts.Logger.Info(icon.SuccessPrefix(), opts.localizer.MustLocalize("generate.log.info.credentialsSaved", localize.NewEntry("FilePath", fileName)))

	return nil
}
