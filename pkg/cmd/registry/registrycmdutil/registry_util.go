package registrycmdutil

import (
	"errors"
	"regexp"

	coreErrors "github.com/redhat-developer/app-services-cli/pkg/core/errors"
)

var validNameRegexp = regexp.MustCompile(`^[a-z]([-a-z0-9]*[a-z0-9])?$`)

// CompatibilityEndpoints - Service Registry API paths for various clients
type CompatibilityEndpoints struct {
	CoreRegistry       string `json:"coreRegistryAPI"`
	SchemaRegistry     string `json:"schemaRegistryCompatAPI"`
	CncfSchemaRegistry string `json:"cncfSchemaRegistryAPI"`
}

const REGISTRY_CORE_PATH = "/apis/registry/v2"
const REGISTRY_COMPAT_PATH = "/apis/ccompat/v6"
const REGISTRY_CNCF_PATH = "/apis/cncf/v0"

// ValidateName validates the proposed name of a Kafka instance
func ValidateName(val interface{}) error {
	name, ok := val.(string)

	if !ok {
		return coreErrors.NewCastError(val, "string")
	}

	if len(name) < 1 || len(name) > 32 {
		return errors.New("service registry instance name must be between 1 and 32 characters")
	}

	matched := validNameRegexp.MatchString(name)

	if matched {
		return nil
	}

	return errors.New("invalid service registry name: " + name)
}

// GetCompatibilityEndpoints returns the compatible API endpoints
func GetCompatibilityEndpoints(url string) *CompatibilityEndpoints {
	endpoints := &CompatibilityEndpoints{
		CoreRegistry:       url + REGISTRY_CORE_PATH,
		SchemaRegistry:     url + REGISTRY_COMPAT_PATH,
		CncfSchemaRegistry: url + REGISTRY_CNCF_PATH,
	}

	return endpoints
}
