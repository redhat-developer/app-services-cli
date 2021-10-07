package delete

import (
	"context"

	"github.com/redhat-developer/app-services-cli/internal/config"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/factory"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/kafka/acl/common"
	"github.com/redhat-developer/app-services-cli/pkg/cmdutil"
	"github.com/redhat-developer/app-services-cli/pkg/connection"
	"github.com/redhat-developer/app-services-cli/pkg/iostreams"
	"github.com/redhat-developer/app-services-cli/pkg/localize"
	"github.com/redhat-developer/app-services-cli/pkg/logging"
	"github.com/spf13/cobra"
)

type options struct {
	Config     config.IConfig
	Connection factory.ConnectionFunc
	Logger     logging.Logger
	IO         *iostreams.IOStreams
	localizer  localize.Localizer
	Context    context.Context

	resourceType    string
	permission      string
	operation       string
	group           string
	topic           string
	transactionalID string
	userID          string
	serviceAccount  string
	patternType     string

	output string
}

// NewDeleteCommand creates a new command to list Kafka ACL rules
func NewDeleteCommand(f *factory.Factory) *cobra.Command {
	opts := &options{
		Config:     f.Config,
		Connection: f.Connection,
		Logger:     f.Logger,
		IO:         f.IOStreams,
		localizer:  f.Localizer,
		Context:    f.Context,
	}

	cmd := &cobra.Command{
		Use:     "delete",
		Short:   f.Localizer.MustLocalize("kafka.acl.list.cmd.shortDescription"),
		Long:    f.Localizer.MustLocalize("kafka.acl.list.cmd.longDescription"),
		Example: f.Localizer.MustLocalize("kafka.acl.list.cmd.example"),
		Args:    cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			if opts.resourceType == common.ResourceTypeFilterGROUP && opts.group == "" {
				return opts.localizer.MustLocalizeError("kafka.acl.common.error.resourceTypeMissingOrIncorrect",
					localize.NewEntry("ResourceType", common.ResourceTypeFilterGROUP), localize.NewEntry("Flag", common.TopicFlagName))
			}
			if opts.resourceType == common.ResourceTypeFilterTOPIC && opts.topic == "" {
				return opts.localizer.MustLocalizeError("kafka.acl.common.error.resourceTypeMissingOrIncorrect",
					localize.NewEntry("ResourceType", common.ResourceTypeFilterTOPIC), localize.NewEntry("Flag", common.TopicFlagName))
			}
			if opts.resourceType == common.ResourceTypeFilterTRANSACTIONAL_ID && opts.transactionalID == "" {
				return opts.localizer.MustLocalizeError("kafka.acl.common.error.resourceTypeMissingOrIncorrect",
					localize.NewEntry("ResourceType", common.ResourceTypeFilterTRANSACTIONAL_ID), localize.NewEntry("Flag", common.TransactionalIDFlag))
			}

			cfg, err := opts.Config.Load()
			if err != nil {
				return err
			}

			if !cfg.HasKafka() {
				return opts.localizer.MustLocalizeError("kafka.acl.common.error.noKafkaSelected")
			}

			instanceID := cfg.Services.Kafka.ClusterID

			return runDelete(instanceID, opts)
		},
	}

	fs := common.NewFlagSet(cmd, opts.localizer)

	// TODO: Do not use a resource-type flag - infer it from the usage of the resource type flags
	_ = fs.AddResourceType(&opts.resourceType).Required()
	_ = fs.AddPermission(&opts.permission).Required()
	// TODO: Should not be required when resource type is cluster
	_ = fs.AddPatternType(&opts.patternType).Required()

	fs.AddOperation(&opts.operation)
	fs.AddTopic(&opts.topic)
	fs.AddConsumerGroup(&opts.group)
	fs.AddTransactionalID(&opts.transactionalID)
	fs.AddOutput(&opts.output)
	fs.AddUser(&opts.userID)
	fs.AddServiceAccount(&opts.serviceAccount)

	return cmd
}

func runDelete(instanceID string, opts *options) error {
	conn, err := opts.Connection(connection.DefaultConfigRequireMasAuth)
	if err != nil {
		return err
	}

	adminAPI, _, err := conn.API().KafkaAdmin(instanceID)
	if err != nil {
		return err
	}

	resourceOperations, httpRes, err := adminAPI.AclsApi.GetAclResourceOperations(opts.Context).Execute()
	if httpRes != nil {
		defer httpRes.Body.Close()
	}
	if err != nil {
		return err
	}

	validResourceOperations := common.GetValidResourceOperations(opts.resourceType, resourceOperations)
	if !common.IsValidOperation(opts.operation, validResourceOperations) {
		return opts.localizer.MustLocalizeError("kafka.acl.common.error.invalidResourceOperation",
			localize.NewEntry("ResourceType", opts.resourceType),
			localize.NewEntry("Operation", opts.operation),
			localize.NewEntry("ValidOperationList", cmdutil.StringSliceToListStringWithQuotes(validResourceOperations)),
		)
	}

	// TODO: Fetch all ACLs that would be deleted and display to user before deleting

	bindingList, httpRes, err := adminAPI.AclsApi.DeleteAcls(opts.Context).
		ResourceType(common.GetResourceTypeFilter(opts.resourceType)).
		// TODO: Allow adding of principal via arguments
		Principal("User:*").
		PatternType(common.GetPatternTypeFilter(opts.patternType)).
		ResourceName(opts.topic).
		Operation(common.GetOperationFilter(opts.operation)).
		Permission(common.GetPermissionFilter(opts.permission)).
		Execute()

	if httpRes != nil {
		defer httpRes.Body.Close()
	}

	if err != nil {
		return err
	}

	opts.Logger.Info("Deleted", bindingList.GetSize(), "ACLs")

	return nil
}
