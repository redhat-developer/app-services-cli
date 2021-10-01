package v1alpha

import (
	"github.com/redhat-developer/app-services-cli/internal/config"
	"github.com/redhat-developer/app-services-cli/pkg/connection"
	"github.com/redhat-developer/app-services-cli/pkg/iostreams"
	"github.com/redhat-developer/app-services-cli/pkg/localize"
	"github.com/redhat-developer/app-services-cli/pkg/logging"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// TODO too generic names
type InputOptions struct {
	Connection connection.Connection
	Config     config.IConfig
	Logger     logging.Logger
	IO         *iostreams.IOStreams
	Localizer  localize.Localizer
}

// CustomResourceOptions object contains the data required to create a custom connection
type CustomResourceOptions struct {
	Path        string
	CRName      string
	ServiceName string
	CRJSON      []byte
	Resource    schema.GroupVersionResource
}
