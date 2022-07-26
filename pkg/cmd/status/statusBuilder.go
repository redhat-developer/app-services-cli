package status

import (
	"fmt"
	"io"
	"reflect"
	"strings"
	"text/tabwriter"

	"github.com/redhat-developer/app-services-cli/pkg/core/cmdutil/flagutil"
	"github.com/redhat-developer/app-services-cli/pkg/shared/contextutil"
	"github.com/redhat-developer/app-services-cli/pkg/shared/factory"
	"github.com/redhat-developer/app-services-cli/pkg/shared/servicespec"
	"github.com/redhat-developer/app-services-cli/pkg/shared/svcstatus"
	registrymgmtclient "github.com/redhat-developer/app-services-sdk-go/registrymgmt/apiv1/client"

	connectormgmtclient "github.com/redhat-developer/app-services-sdk-go/connectormgmt/apiv1/client"
	kafkamgmtclient "github.com/redhat-developer/app-services-sdk-go/kafkamgmt/apiv1/client"

	"github.com/redhat-developer/app-services-cli/pkg/core/servicecontext"

	"github.com/openconfig/goyang/pkg/indent"
)

const tagTitle = "title"

type serviceStatus struct {
	Name      string           `json:"name,omitempty" title:"Service Context Name"`
	Location  string           `json:"location,omitempty" title:"Context File Location"`
	Kafka     *kafkaStatus     `json:"kafka,omitempty" title:"Kafka"`
	Registry  *registryStatus  `json:"registry,omitempty" title:"Service Registry"`
	Namespace *NamespaceStatus `json:"namespace,omitempty" title:"Namespace"`
	Connector *ConnectorStatus `json:"connector,omitempty" title:"Connector"`
}

func (s serviceStatus) hasStatus() bool {
	return s.Kafka != nil || s.Registry != nil || s.Connector != nil || s.Namespace != nil
}

type kafkaStatus struct {
	ID                  string `json:"id,omitempty"`
	Name                string `json:"name,omitempty"`
	Status              string `json:"status,omitempty"`
	BootstrapServerHost string `json:"bootstrap_server_host,omitempty" title:"Bootstrap URL"`
	FailedReason        string `json:"failed_reason,omitempty" title:"Failed Reason"`
}

type registryStatus struct {
	ID          string `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	Status      string `json:"status,omitempty"`
	RegistryUrl string `json:"registryUrl,omitempty" title:"Registry URL"`
}

type NamespaceStatus struct {
	Name string `json:"name,omitempty"`
	ID   string `json:"id,omitempty"`
}

type ConnectorStatus struct {
	ID      string `json:"id,omitempty"`
	Name    string `json:"name,omitempty"`
	Type    string `json:"type,omitempty"`
	KafkaID string `json:"kafka_id,omitempty"`
	Status  string `json:"status,omitempty"`
}

type clientConfig struct {
	f             *factory.Factory
	serviceConfig *servicecontext.ServiceConfig
}

type statusClient struct {
	f             *factory.Factory
	serviceConfig *servicecontext.ServiceConfig
}

// newStatusClient returns a new client to fetch service statuses
// and build it into a service status config object
func newStatusClient(cfg *clientConfig) *statusClient {
	return &statusClient{
		f:             cfg.f,
		serviceConfig: cfg.serviceConfig,
	}
}

// BuildStatus gets the status of all services currently set in the service context
func (c *statusClient) BuildStatus(services []string) (status *serviceStatus, err error) {
	factory := c.f

	status = &serviceStatus{}

	if flagutil.StringInSlice(servicespec.KafkaServiceName, services) && c.serviceConfig.KafkaID != "" {
		kafkaResponse, err1 := contextutil.GetKafkaForServiceConfig(c.serviceConfig, factory)
		if err1 != nil {
			return status, err1
		}
		kafkaStatus := c.getKafkaStatus(kafkaResponse)
		status.Kafka = kafkaStatus
	}

	if flagutil.StringInSlice(servicespec.ServiceRegistryServiceName, services) && c.serviceConfig.ServiceRegistryID != "" {
		registryResponse, err1 := contextutil.GetRegistryForServiceConfig(c.serviceConfig, factory)
		if err1 != nil {
			return status, err1
		}
		registry := c.getRegistryStatus(registryResponse)
		status.Registry = registry
	}

	if flagutil.StringInSlice(servicespec.ConnectorServiceName, services) && c.serviceConfig.ConnectorID != "" {

		conn, err1 := factory.Connection()
		if err1 != nil {
			return nil, err1
		}

		connectorResponse, err2 := contextutil.GetConnectorForServiceConfig(c.serviceConfig, &conn, factory)
		if err2 != nil {
			return status, err2
		}

		factory.Logger.Info(connectorResponse)

		connectorStatus := c.getConnectorStatus(connectorResponse)
		status.Connector = connectorStatus
	}

	if flagutil.StringInSlice(servicespec.NamespaceServiceName, services) && c.serviceConfig.NamespaceID != "" {

		conn, err1 := factory.Connection()
		if err1 != nil {
			return nil, err1
		}

		namespaceResponse, err2 := contextutil.GetNamespaceForServiceConfig(c.serviceConfig, &conn, factory)
		if err2 != nil {
			return status, err2
		}

		factory.Logger.Info(namespaceResponse)

		namespaceStatus := c.getNamespaceStatus(namespaceResponse)
		status.Namespace = namespaceStatus
	}

	return status, err
}

func (c *statusClient) getKafkaStatus(kafkaResponse *kafkamgmtclient.KafkaRequest) (status *kafkaStatus) {
	status = &kafkaStatus{
		ID:                  kafkaResponse.GetId(),
		Name:                kafkaResponse.GetName(),
		Status:              kafkaResponse.GetStatus(),
		BootstrapServerHost: kafkaResponse.GetBootstrapServerHost(),
	}

	if kafkaResponse.GetStatus() == svcstatus.StatusFailed {
		status.FailedReason = kafkaResponse.GetFailedReason()
	}

	return status
}

func (c *statusClient) getRegistryStatus(registry *registrymgmtclient.Registry) (status *registryStatus) {
	status = &registryStatus{
		ID:          registry.GetId(),
		Name:        registry.GetName(),
		RegistryUrl: registry.GetRegistryUrl(),
		Status:      string(registry.GetStatus()),
	}

	return status
}

func (c *statusClient) getNamespaceStatus(namespace *connectormgmtclient.ConnectorNamespace) (status *NamespaceStatus) {
	status = &NamespaceStatus{
		ID:   namespace.GetId(),
		Name: namespace.GetName(),
	}

	return status
}

func (c *statusClient) getConnectorStatus(connector *connectormgmtclient.Connector) (status *ConnectorStatus) {
	status = &ConnectorStatus{
		ID:      *connector.Id,
		Name:    connector.Name,
		Type:    connector.ConnectorTypeId,
		KafkaID: connector.Kafka.Id,
		Status:  string(*connector.Status.State),
	}

	return status
}

// Print prints the status information of all set services
func Print(w io.Writer, status *serviceStatus) {
	v := reflect.ValueOf(status).Elem()

	indirectVal := reflect.Indirect(v)
	for i := 0; i < indirectVal.NumField(); i++ {
		fieldType := indirectVal.Type().Field(i)
		fieldVal := indirectVal.Field(i)
		title := getTitle(&fieldType)

		if fieldVal.Kind() == reflect.String {
			fmt.Fprintf(w, "%v:\t%v\n", title, fieldVal)
		}

		if fieldVal.Kind() == reflect.Ptr && !fieldVal.IsNil() {
			fmt.Fprintln(w, "")
			printServiceStatus(w, title, fieldVal)
		}
	}
}

// print the status of service v
func printServiceStatus(w io.Writer, name string, v reflect.Value) {
	indentWriter := indent.NewWriter(w, "  ")

	// set table padding
	padding := 5

	// create a new tabwriter
	tw := tabwriter.NewWriter(indentWriter, 0, 0, padding, ' ', tabwriter.TabIndent)

	// tracks the longest row in chars so we can set an equal length divider
	maxRowLen := 0
	// iterate over every field in the type

	indirectV := reflect.Indirect(v)
	for i := 0; i < indirectV.NumField(); i++ {
		// get field type metadata
		fieldType := indirectV.Type().Field(i)

		// get value of the field
		fieldValue := indirectV.Field(i)

		// get the title to use for the field
		title := getTitle(&fieldType)

		if !getOmitEmpty(&fieldType) || !fieldValue.IsZero() {
			// print the row and take note of its character length
			charLen, _ := fmt.Fprintf(tw, "%v:\t\t%v\n", title, fieldValue)
			if charLen > maxRowLen {
				maxRowLen = charLen
			}
		}
	}
	// print the service header
	fmt.Fprintln(indentWriter, name)
	// print the title divider
	fmt.Fprintln(indentWriter, createDivider(maxRowLen+padding))

	tw.Flush()
}

// get title of the field from the "title" tag
// If the tag does not exist, use the name of the field
func getTitle(f *reflect.StructField) string {
	// Get the field tag value
	tag := f.Tag.Get(tagTitle)
	if tag == "" {
		tag = f.Name
	}

	return tag
}

// check if omitempty is set
func getOmitEmpty(f *reflect.StructField) bool {
	var omitempty bool
	tag := f.Tag.Get("json")
	if tag != "" {
		omitempty = strings.Contains(tag, "omitempty")
	}

	return omitempty
}

// create a divider for the top of the table of n length
func createDivider(n int) string {
	b := "-"
	for i := 0; i <= n; i++ {
		b += "-"
	}

	return b
}
