package status

import (
	"context"
	"fmt"
	"io"
	"reflect"
	"strings"
	"text/tabwriter"

	"github.com/redhat-developer/app-services-cli/pkg/kafkautil"
	"github.com/redhat-developer/app-services-cli/pkg/svcstatus"

	"github.com/redhat-developer/app-services-cli/pkg/core/config"
	"github.com/redhat-developer/app-services-cli/pkg/core/connection"
	"github.com/redhat-developer/app-services-cli/pkg/core/localize"
	"github.com/redhat-developer/app-services-cli/pkg/core/logging"

	"github.com/redhat-developer/app-services-cli/pkg/shared"

	"github.com/openconfig/goyang/pkg/indent"
)

const tagTitle = "title"

type serviceStatus struct {
	Kafka    *kafkaStatus    `json:"kafka,omitempty" title:"Kafka"`
	Registry *registryStatus `json:"registry,omitempty" title:"Service Registry"`
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

type clientConfig struct {
	config     config.IConfig
	Logger     logging.Logger
	connection connection.Connection
	localizer  localize.Localizer
}

type statusClient struct {
	config    config.IConfig
	Logger    logging.Logger
	conn      connection.Connection
	localizer localize.Localizer
}

// newStatusClient returns a new client to fetch service statuses
// and build it into a service status config object
func newStatusClient(cfg *clientConfig) *statusClient {
	return &statusClient{
		config:    cfg.config,
		Logger:    cfg.Logger,
		conn:      cfg.connection,
		localizer: cfg.localizer,
	}
}

// BuildStatus gets the status of all services currently set in the user config
func (c *statusClient) BuildStatus(ctx context.Context, services []string) (status *serviceStatus, ok bool, err error) {
	cfg, err := c.config.Load()
	if err != nil {
		return nil, false, err
	}

	status = &serviceStatus{}

	if stringInSlice(shared.ServiceRegistryServiceName, services) {
		if instanceID, exists := cfg.GetKafkaIdOk(); exists {
			// nolint:govet
			kafkaStatus, err := c.getKafkaStatus(ctx, instanceID)
			if err != nil {
				if kafkautil.IsErr(err, kafkautil.ErrorCode7) {
					err = kafkautil.NotFoundByIDError(instanceID)
					c.Logger.Error(err)
					c.Logger.Info(c.localizer.MustLocalize("status.log.info.rhoasKafkaUse"))
				}
			} else {
				status.Kafka = kafkaStatus
				ok = true
			}
		} else {
			c.Logger.Debug("No Kafka instance is currently used, skipping status check")
		}
	}

	if stringInSlice(shared.ServiceRegistryServiceName, services) {
		registryCfg := cfg.Services.ServiceRegistry
		if registryCfg != nil && registryCfg.InstanceID != "" {
			// nolint:govet
			registry, newErr := c.getRegistryStatus(ctx, registryCfg.InstanceID)
			if newErr != nil {
				return status, ok, err
			}
			status.Registry = registry
			ok = true
		} else {
			c.Logger.Debug("No service registry is currently used, skipping status check")
		}
	}

	return status, ok, err
}

func (c *statusClient) getKafkaStatus(ctx context.Context, id string) (status *kafkaStatus, err error) {
	kafkaResponse, _, err := c.conn.API().KafkaMgmt().GetKafkaById(ctx, id).Execute()
	if err != nil {
		return nil, err
	}

	status = &kafkaStatus{
		ID:                  kafkaResponse.GetId(),
		Name:                kafkaResponse.GetName(),
		Status:              kafkaResponse.GetStatus(),
		BootstrapServerHost: kafkaResponse.GetBootstrapServerHost(),
	}

	if kafkaResponse.GetStatus() == svcstatus.StatusFailed {
		status.FailedReason = kafkaResponse.GetFailedReason()
	}

	return status, err
}

func (c *statusClient) getRegistryStatus(ctx context.Context, id string) (status *registryStatus, err error) {
	registry, _, err := c.conn.API().ServiceRegistryMgmt().GetRegistry(ctx, id).Execute()
	if err != nil {
		return nil, err
	}

	status = &registryStatus{
		ID:          registry.GetId(),
		Name:        registry.GetName(),
		RegistryUrl: registry.GetRegistryUrl(),
		Status:      string(registry.GetStatus()),
	}

	return status, err
}

// Print prints the status information of all set services
func Print(w io.Writer, status *serviceStatus) {
	v := reflect.ValueOf(status).Elem()

	indirectVal := reflect.Indirect(v)
	for i := 0; i < indirectVal.NumField(); i++ {
		fieldType := indirectVal.Type().Field(i)
		fieldVal := indirectVal.Field(i)

		if !fieldVal.IsNil() {
			title := getTitle(&fieldType)
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

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}
