package profileutil

import (
	"context"

	"github.com/redhat-developer/app-services-cli/pkg/core/localize"
	"github.com/redhat-developer/app-services-cli/pkg/core/servicecontext"
	kafkamgmtclient "github.com/redhat-developer/app-services-sdk-go/kafkamgmt/apiv1/client"
	kafkamgmtv1errors "github.com/redhat-developer/app-services-sdk-go/kafkamgmt/apiv1/error"
	registrymgmtclient "github.com/redhat-developer/app-services-sdk-go/registrymgmt/apiv1/client"
	srsmgmtv1errors "github.com/redhat-developer/app-services-sdk-go/registrymgmt/apiv1/error"
)

type ContextHandler struct {
	Context   *servicecontext.Context
	Localizer localize.Localizer
}

// GetCurrentKafkaInstance returns the Kafka instance set in the currently selected context
func (c *ContextHandler) GetCurrentKafkaInstance(api kafkamgmtclient.DefaultApi) (*kafkamgmtclient.KafkaRequest, error) {

	currentCtx, err := c.getCurrentContext()
	if err != nil {
		return nil, err
	}

	s, err := c.GetContext(currentCtx)
	if err != nil {
		return nil, err
	}

	if s.KafkaID == "" {
		return nil, c.Localizer.MustLocalizeError("context.common.error.noKafkaID")
	}

	kafkaInstance, _, err := api.GetKafkaById(context.Background(), s.KafkaID).Execute()
	if kafkamgmtv1errors.IsAPIError(err, kafkamgmtv1errors.ERROR_7) {
		return nil, c.Localizer.MustLocalizeError("context.common.error.registry.notFound")
	}

	return &kafkaInstance, err
}

// GetCurrentRegistryInstance returns the Service Registry instance set in the currently selected context
func (c *ContextHandler) GetCurrentRegistryInstance(api registrymgmtclient.RegistriesApi) (*registrymgmtclient.Registry, error) {

	currentCtx, err := c.getCurrentContext()
	if err != nil {
		return nil, err
	}

	s, err := c.GetContext(currentCtx)
	if err != nil {
		return nil, err
	}

	if s.ServiceRegistryID == "" {
		return nil, c.Localizer.MustLocalizeError("context.common.error.noRegistryID")
	}

	registryInstance, _, err := api.GetRegistry(context.Background(), s.ServiceRegistryID).Execute()

	if srsmgmtv1errors.IsAPIError(err, srsmgmtv1errors.ERROR_2) {
		return nil, c.Localizer.MustLocalizeError("context.common.error.registry.notFound")
	}

	return &registryInstance, nil
}

// GetContext returns the services associated with the context
func (c *ContextHandler) GetContext(ctxName string) (*servicecontext.ServiceConfig, error) {

	currCtx, ok := c.Context.Contexts[ctxName]
	if !ok {
		return nil, c.Localizer.MustLocalizeError("context.common.error.context.notFound", localize.NewEntry("Name", ctxName))
	}

	return &currCtx, nil
}

// getCurrentContext returns the name of the currently selected context
func (c *ContextHandler) getCurrentContext() (string, error) {

	if c.Context.CurrentContext == "" {
		return "", c.Localizer.MustLocalizeError("context.common.error.notSet")
	}

	return c.Context.CurrentContext, nil
}
