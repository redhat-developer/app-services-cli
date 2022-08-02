package aclcmdutil

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/redhat-developer/app-services-cli/pkg/cmd/kafka/kafkacmdutil"
	"github.com/redhat-developer/app-services-cli/pkg/core/ioutil/color"
	"github.com/redhat-developer/app-services-cli/pkg/core/ioutil/iostreams"
	"github.com/redhat-developer/app-services-cli/pkg/core/localize"
	"github.com/redhat-developer/app-services-cli/pkg/core/logging"
	"github.com/redhat-developer/app-services-cli/pkg/core/servicecontext"
	"github.com/redhat-developer/app-services-cli/pkg/shared/factory"
	"github.com/spf13/cobra"

	kafkainstanceclient "github.com/redhat-developer/app-services-sdk-go/kafkainstance/apiv1/client"
)

// CrudOptions is the interface used for options of create and delete command
type CrudOptions struct {
	Connection     factory.ConnectionFunc
	Logger         logging.Logger
	IO             *iostreams.IOStreams
	Localizer      localize.Localizer
	Context        context.Context
	ServiceContext servicecontext.IContext

	Cluster         bool
	PatternType     string
	ResourceType    string
	ResourceName    string
	Permission      string
	Operation       string
	Group           string
	Topic           string
	TransactionalID string
	Principal       string

	SkipConfirm bool
	Output      string
	InstanceID  string
}

// When the value of the `--topic`, `--group`, `user` or `service-account` option is one of
// the keys of this map, it will be replaced by the corresponding value.
var commonArgAliases = map[string]string{
	AllAlias: Wildcard,
}

// ExecuteACLRuleCreate makes request to create an ACL rule
func ExecuteACLRuleCreate(req kafkainstanceclient.ApiCreateAclRequest, localizer localize.Localizer, kafkaInstanceName string) error {

	httpRes, err := req.Execute()
	if httpRes != nil {
		defer httpRes.Body.Close()
	}

	return ValidateAPIError(httpRes, localizer, err, "create", kafkaInstanceName)
}

// FormatPrincipal formats the provided principal ID to "User:principal"
func FormatPrincipal(userID string) string {
	return fmt.Sprintf("User:%s", userID)
}

// GetResourceName returns the name of the resource
// transformed into a server recognized format
func GetResourceName(resourceName string) string {
	if commonArgAliases[resourceName] == Wildcard {
		return Wildcard
	}
	return resourceName
}

// IsValidResourceOperation checks if the operation is valid, and returns the list valid operations when invalid
func IsValidResourceOperation(resourceType string, operation string, resourceOperationsMap map[string][]string) (bool, []string) {
	resourceTypeMapped := resourceTypeOperationKeyMap[resourceType]
	resourceOperations := resourceOperationsMap[resourceTypeMapped]

	for i, op := range resourceOperations {
		if operationMapped, ok := validOperationsResponseMap[op]; ok {
			resourceOperations[i] = operationMapped
		} else {
			resourceOperations[i] = op
		}
		if resourceOperations[i] == operation {
			return true, nil
		}
	}

	return false, resourceOperations
}

// SetACLResources sets resources options and returns number of changed resources
func SetACLResources(opts *CrudOptions) int {
	var selectedResourceTypeCount int

	if opts.Topic != "" {
		selectedResourceTypeCount++
		opts.ResourceType = ResourceTypeTOPIC
		opts.ResourceName = opts.Topic
	}
	if opts.Group != "" {
		selectedResourceTypeCount++
		opts.ResourceType = ResourceTypeGROUP
		opts.ResourceName = opts.Group
	}
	if opts.TransactionalID != "" {
		selectedResourceTypeCount++
		opts.ResourceType = ResourceTypeTRANSACTIONAL_ID
		opts.ResourceName = opts.TransactionalID
	}
	if opts.Cluster {
		selectedResourceTypeCount++
		opts.ResourceType = ResourceTypeCLUSTER
		opts.ResourceName = KafkaCluster
	}

	return selectedResourceTypeCount
}

// ValidateAPIError checks for a HTTP error and maps it to a user friendly error
func ValidateAPIError(httpRes *http.Response, localizer localize.Localizer, err error, operation string, instanceName string) error {
	if err == nil {
		return nil
	}

	if httpRes == nil {
		return err
	}

	operationTmplPair := localize.NewEntry("Operation", operation)

	switch httpRes.StatusCode {
	case http.StatusUnauthorized:
		return localizer.MustLocalizeError("kafka.acl.common.error.unauthorized", operationTmplPair)
	case http.StatusForbidden:
		return localizer.MustLocalizeError("kafka.acl.common.error.forbidden", operationTmplPair)
	case http.StatusInternalServerError:
		return localizer.MustLocalizeError("kafka.acl.common.error.internalServerError")
	case http.StatusServiceUnavailable:
		return localizer.MustLocalizeError("kafka.acl.common.error.unableToConnectToKafka", localize.NewEntry("Name", instanceName))
	default:
		return err
	}
}

// BuildInstructions accepts a slice of errors and creates a single formatted error object
func BuildInstructions(errorCollection []error) error {

	errString := "invalid or missing option(s):" + "\n"

	for _, err := range errorCollection {
		errString += fmt.Sprintf("   %s ", color.Error("*")) + err.Error() + "\n"
	}

	return errors.New(errString)
}

// RegisterTopicCompletionFunc enables dynamic autocompletion for topic flag
func RegisterTopicCompletionFunc(cmd *cobra.Command, f *factory.Factory) error {
	return cmd.RegisterFlagCompletionFunc("topic", func(cmd *cobra.Command, _ []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return kafkacmdutil.FilterValidTopicNameArgs(f, toComplete)
	})
}

// RegisterGroupCompletionFunc enables dynamic autocompletion for group flag
func RegisterGroupCompletionFunc(cmd *cobra.Command, f *factory.Factory) error {
	return cmd.RegisterFlagCompletionFunc("group", func(cmd *cobra.Command, _ []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return kafkacmdutil.FilterValidConsumerGroupIDs(f, toComplete)
	})
}
