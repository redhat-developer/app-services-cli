package contextutil

import (
	"context"
	"github.com/redhat-developer/app-services-cli/pkg/core/localize"
	"github.com/redhat-developer/app-services-cli/pkg/core/servicecontext"

	"github.com/redhat-developer/app-services-cli/pkg/shared/connection"
	"github.com/redhat-developer/app-services-cli/pkg/shared/factory"
	connectormgmtclient "github.com/redhat-developer/app-services-sdk-go/connectormgmt/apiv1/client"
	srsmgmtv1errors "github.com/redhat-developer/app-services-sdk-go/registrymgmt/apiv1/error"

	registrymgmtclient "github.com/redhat-developer/app-services-sdk-go/registrymgmt/apiv1/client"

	kafkamgmtclient "github.com/redhat-developer/app-services-sdk-go/kafkamgmt/apiv1/client"
	kafkamgmtv1errors "github.com/redhat-developer/app-services-sdk-go/kafkamgmt/apiv1/error"
)

// GetContext returns the services associated with the context
func GetContext(svcContext *servicecontext.Context, localizer localize.Localizer, ctxName string) (*servicecontext.ServiceConfig, error) {

	ctx, ok := svcContext.Contexts[ctxName]
	if !ok {
		return nil, localizer.MustLocalizeError("context.common.error.context.notFound", localize.NewEntry("Name", svcContext.CurrentContext))
	}

	return &ctx, nil

}

// GetCurrentContext returns the name of the currently selected context
func GetCurrentContext(svcContext *servicecontext.Context, localizer localize.Localizer) (*servicecontext.ServiceConfig, error) {

	if svcContext.CurrentContext == "" {
		return nil, localizer.MustLocalizeError("context.common.error.notSet")
	}

	currCtx, ok := svcContext.Contexts[svcContext.CurrentContext]
	if !ok {
		return nil, localizer.MustLocalizeError("context.common.error.context.notFound", localize.NewEntry("Name", svcContext.CurrentContext))
	}

	return &currCtx, nil
}

// GetCurrentKafkaInstance returns the Kafka instance set in the currently selected context
func GetCurrentKafkaInstance(f *factory.Factory) (*kafkamgmtclient.KafkaRequest, error) {

	svcContext, err := f.ServiceContext.Load()
	if err != nil {
		return nil, err
	}

	currCtx, err := GetCurrentContext(svcContext, f.Localizer)
	if err != nil {
		return nil, err
	}

	return GetKafkaForServiceConfig(currCtx, f)
}

func GetKafkaForServiceConfig(currCtx *servicecontext.ServiceConfig, f *factory.Factory) (*kafkamgmtclient.KafkaRequest, error) {
	conn, err := f.Connection()
	if err != nil {
		return nil, err
	}
	if currCtx.KafkaID == "" {
		return nil, f.Localizer.MustLocalizeError("context.common.error.noKafkaID")
	}

	kafkaInstance, _, err := conn.API().KafkaMgmt().GetKafkaById(context.Background(), currCtx.KafkaID).Execute()
	if kafkamgmtv1errors.IsAPIError(err, kafkamgmtv1errors.ERROR_7) {
		return nil, f.Localizer.MustLocalizeError("context.common.error.kafka.notFound")
	}

	return &kafkaInstance, nil
}

// GetCurrentRegistryInstance returns the Service Registry instance set in the currently selected context
func GetCurrentRegistryInstance(f *factory.Factory) (*registrymgmtclient.Registry, error) {

	svcContext, err := f.ServiceContext.Load()
	if err != nil {
		return nil, err
	}

	currCtx, err := GetCurrentContext(svcContext, f.Localizer)
	if err != nil {
		return nil, err
	}

	return GetRegistryForServiceConfig(currCtx, f)

}

func GetRegistryForServiceConfig(currCtx *servicecontext.ServiceConfig, f *factory.Factory) (*registrymgmtclient.Registry, error) {
	conn, err := f.Connection()
	if err != nil {
		return nil, err
	}

	if currCtx.ServiceRegistryID == "" {
		return nil, f.Localizer.MustLocalizeError("context.common.error.noRegistryID")
	}

	registryInstance, _, err := conn.API().ServiceRegistryMgmt().GetRegistry(context.Background(), currCtx.ServiceRegistryID).Execute()
	if srsmgmtv1errors.IsAPIError(err, srsmgmtv1errors.ERROR_2) {
		return nil, f.Localizer.MustLocalizeError("context.common.error.registry.notFound")
	}

	return &registryInstance, nil
}

// GetCurrentConnectorInstance returns the connector instance set in the currently selected context
func GetCurrentConnectorInstance(conn *connection.Connection, f *factory.Factory) (*connectormgmtclient.Connector, error) {

	svcContext, err := f.ServiceContext.Load()
	if err != nil {
		return nil, err
	}

	currCtx, err := GetCurrentContext(svcContext, f.Localizer)
	if err != nil {
		return nil, err
	}

	return GetConnectorForServiceConfig(currCtx, conn, f)
}

func GetConnectorForServiceConfig(currCtx *servicecontext.ServiceConfig, conn *connection.Connection, f *factory.Factory) (*connectormgmtclient.Connector, error) {

	if currCtx.ConnectorID == "" {
		return nil, f.Localizer.MustLocalizeError("context.common.error.noConnectorID")
	}

	connectorInstance, _, err := (*conn).API().ConnectorsMgmt().ConnectorsApi.GetConnector(f.Context, currCtx.ConnectorID).Execute()
	if err != nil {
		return nil, err
	}

	return &connectorInstance, nil
}

func SetCurrentConnectorInstance(connector *connectormgmtclient.Connector, conn *connection.Connection, f *factory.Factory) error {

	svcContext, err := f.ServiceContext.Load()
	if err != nil {
		return err
	}

	currCtx, err := GetCurrentContext(svcContext, f.Localizer)
	if err != nil {
		return err
	}

	currCtx.ConnectorID = connector.GetId()
	svcContext.Contexts[svcContext.CurrentContext] = *currCtx

	if err = f.ServiceContext.Save(svcContext); err != nil {
		return err
	}

	return nil
}

// GetCurrentNamespaceInstance returns the namespace set in the currently selected context
func GetCurrentNamespaceInstance(conn *connection.Connection, f *factory.Factory) (*connectormgmtclient.ConnectorNamespace, error) {

	svcContext, err := f.ServiceContext.Load()
	if err != nil {
		return nil, err
	}

	currCtx, err := GetCurrentContext(svcContext, f.Localizer)
	if err != nil {
		return nil, err
	}

	return GetNamespaceForServiceConfig(currCtx, conn, f)
}

func GetNamespaceForServiceConfig(currCtx *servicecontext.ServiceConfig, conn *connection.Connection, f *factory.Factory) (*connectormgmtclient.ConnectorNamespace, error) {

	if currCtx.NamespaceID == "" {
		return nil, f.Localizer.MustLocalizeError("context.common.error.noConnectorID")
	}

	namespace, _, err := (*conn).API().ConnectorsMgmt().ConnectorNamespacesApi.GetConnectorNamespace(f.Context, currCtx.NamespaceID).Execute()
	if err != nil {
		return nil, err
	}

	return &namespace, err
}

func SetCurrentNamespaceInstance(namespace *connectormgmtclient.ConnectorNamespace, conn *connection.Connection, f *factory.Factory) error {

	svcContext, err := f.ServiceContext.Load()
	if err != nil {
		return err
	}

	currCtx, err := GetCurrentContext(svcContext, f.Localizer)
	if err != nil {
		return err
	}

	currCtx.ConnectorID = namespace.GetId()
	svcContext.Contexts[svcContext.CurrentContext] = *currCtx

	if err = f.ServiceContext.Save(svcContext); err != nil {
		return err
	}

	return nil
}
