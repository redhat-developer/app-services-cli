package flagutil

import (
	"github.com/redhat-developer/app-services-cli/pkg/cmd/kafka/acl/sdk"
	"github.com/redhat-developer/app-services-cli/pkg/core/cmdutil/factory"
	"github.com/redhat-developer/app-services-cli/pkg/core/cmdutil/flagutil"
	"github.com/redhat-developer/app-services-cli/pkg/core/localize"
	"github.com/spf13/cobra"
)

const (
	ClusterFlagName         = "cluster"
	TopicFlagName           = "topic"
	GroupFlagName           = "group"
	TransactionalIDFlagName = "transactional-id"
)

var ResourceTypeFlagEntries = []*localize.TemplateEntry{
	localize.NewEntry("ClusterFlag", ClusterFlagName),
	localize.NewEntry("TopicFlag", TopicFlagName),
	localize.NewEntry("TransactionalIDFlag", TransactionalIDFlagName),
	localize.NewEntry("GroupFlag", GroupFlagName),
}

type flagSet struct {
	cmd     *cobra.Command
	factory *factory.Factory
	*flagutil.FlagSet
}

// NewFlagSet returns a new flag set with common Kafka ACL flags
func NewFlagSet(cmd *cobra.Command, f *factory.Factory) *flagSet {
	return &flagSet{
		cmd:     cmd,
		factory: f,
		FlagSet: flagutil.NewFlagSet(cmd, f.Localizer),
	}
}

// AddResourceType adds a flag for ACL resource type and registers completion options
func (fs *flagSet) AddResourceType(resourceType *string) *flagutil.FlagOptions {
	flagName := "resource-type"

	resourceTypeFilterMap := sdk.GetResourceTypeFilterMap()

	resourceTypes := make([]string, 0, len(resourceTypeFilterMap))
	for i := range resourceTypeFilterMap {
		resourceTypes = append(resourceTypes, i)
	}

	fs.StringVar(
		resourceType,
		flagName,
		sdk.ResourceTypeANY,
		flagutil.FlagDescription(fs.factory.Localizer, "kafka.acl.common.flag.resourceType", resourceTypes...),
	)

	_ = fs.cmd.RegisterFlagCompletionFunc(flagName, func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return resourceTypes, cobra.ShellCompDirectiveNoSpace
	})

	return flagutil.WithFlagOptions(fs.cmd, flagName)
}

// AddOperationFilter adds a flag for ACL operations filter and registers completion options
func (fs *flagSet) AddOperationFilter(operationType *string) *flagutil.FlagOptions {
	flagName := "operation"

	operationFilterMap := sdk.GetOperationFilterMap()

	operations := make([]string, 0, len(operationFilterMap))
	for i := range operationFilterMap {
		operations = append(operations, i)
	}

	fs.StringVar(
		operationType,
		flagName,
		"",
		flagutil.FlagDescription(fs.factory.Localizer, "kafka.acl.common.flag.operation.description", operations...),
	)

	_ = fs.cmd.RegisterFlagCompletionFunc(flagName, func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return operations, cobra.ShellCompDirectiveNoSpace
	})

	return flagutil.WithFlagOptions(fs.cmd, flagName)
}

// AddOperationCreate adds a flag for ACL operations and registers completion options
func (fs *flagSet) AddOperationCreate(operationType *string) *flagutil.FlagOptions {
	flagName := "operation"

	operationMap := sdk.GetOperationMap()

	operations := make([]string, 0, len(operationMap))
	for i := range operationMap {
		operations = append(operations, i)
	}

	fs.StringVar(
		operationType,
		flagName,
		"",
		flagutil.FlagDescription(fs.factory.Localizer, "kafka.acl.common.flag.operation.description", operations...),
	)

	_ = fs.cmd.RegisterFlagCompletionFunc(flagName, func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return operations, cobra.ShellCompDirectiveNoSpace
	})

	return flagutil.WithFlagOptions(fs.cmd, flagName)
}

// AddPermissionFilter adds a flag for ACL permissions filters
func (fs *flagSet) AddPermissionFilter(permission *string) *flagutil.FlagOptions {
	flagName := "permission"

	permissionTypeFilterMap := sdk.GetPermissionTypeFilterMap()

	permissions := make([]string, 0, len(permissionTypeFilterMap))
	for i := range permissionTypeFilterMap {
		permissions = append(permissions, i)
	}

	fs.StringVar(
		permission,
		flagName,
		sdk.PermissionANY,
		flagutil.FlagDescription(fs.factory.Localizer, "kafka.acl.common.flag.permission.description", permissions...),
	)

	_ = fs.cmd.RegisterFlagCompletionFunc(flagName, func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return permissions, cobra.ShellCompDirectiveNoSpace
	})

	return flagutil.WithFlagOptions(fs.cmd, flagName)
}

// AddPermissionCreate adds a flag for ACL permissions
func (fs *flagSet) AddPermissionCreate(permission *string) *flagutil.FlagOptions {
	flagName := "permission"

	permissionTypeMap := sdk.GetPermissionTypeMap()

	permissions := make([]string, 0, len(permissionTypeMap))
	for i := range permissionTypeMap {
		permissions = append(permissions, i)
	}

	fs.StringVar(
		permission,
		flagName,
		"",
		flagutil.FlagDescription(fs.factory.Localizer, "kafka.acl.common.flag.permission.description", permissions...),
	)

	_ = fs.cmd.RegisterFlagCompletionFunc(flagName, func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return permissions, cobra.ShellCompDirectiveNoSpace
	})

	return flagutil.WithFlagOptions(fs.cmd, flagName)
}

// AddTopic adds a flag for setting the topic name
func (fs *flagSet) AddTopic(topic *string) {
	flagName := TopicFlagName

	fs.StringVar(
		topic,
		flagName,
		"",
		fs.factory.Localizer.MustLocalize("kafka.acl.common.flag.topic.description"),
	)

	_ = flagutil.RegisterTopicCompletionFunc(fs.cmd, fs.factory)
}

// AddConsumerGroup adds a flag for setting the consumer group ID
func (fs *flagSet) AddConsumerGroup(group *string) {
	flagName := GroupFlagName

	fs.StringVar(
		group,
		flagName,
		"",
		fs.factory.Localizer.MustLocalize("kafka.acl.common.flag.group.description"),
	)

	_ = flagutil.RegisterGroupCompletionFunc(fs.cmd, fs.factory)
}

// AddTransactionalID adds a flag for setting the consumer group ID
func (fs *flagSet) AddTransactionalID(id *string) {
	flagName := TransactionalIDFlagName

	fs.StringVar(
		id,
		flagName,
		"",
		fs.factory.Localizer.MustLocalize("kafka.acl.common.flag.transactionalID.description"),
	)
}

// AddPrefix adds a flag for enabling the PREFIX resource pattern type
func (fs *flagSet) AddPrefix(prefix *bool) {
	flagName := "prefix"

	fs.BoolVar(
		prefix,
		flagName,
		false,
		fs.factory.Localizer.MustLocalize("kafka.acl.common.flag.prefix.description"),
	)
}

// AddCluster adds a flag for setting the "cluster" ACL resource type
func (fs *flagSet) AddCluster(prefix *bool) {
	flagName := ClusterFlagName

	fs.BoolVar(
		prefix,
		flagName,
		false,
		fs.factory.Localizer.MustLocalize("kafka.acl.common.flag.cluster.description"),
	)
}

// AddUser adds a flag to pass a user ID principal
func (fs *flagSet) AddUser(userID *string) *flagutil.FlagOptions {
	flagName := "user"

	fs.StringVar(
		userID,
		flagName,
		"",
		fs.factory.Localizer.MustLocalize("kafka.acl.common.flag.user.description"),
	)

	_ = flagutil.RegisterUserCompletionFunc(fs.cmd, flagName, fs.factory)

	return flagutil.WithFlagOptions(fs.cmd, flagName)
}

// AddServiceAccount adds a flag to pass a service account client ID principal
func (fs *flagSet) AddServiceAccount(serviceAccountID *string) *flagutil.FlagOptions {
	flagName := "service-account"

	fs.StringVar(
		serviceAccountID,
		flagName,
		"",
		fs.factory.Localizer.MustLocalize("kafka.acl.common.flag.serviceAccount.description"),
	)

	_ = flagutil.RegisterServiceAccountCompletionFunc(fs.cmd, fs.factory)

	return flagutil.WithFlagOptions(fs.cmd, flagName)
}

// AddInstanceID adds a flag for the Kafka instance ID
func (fs *flagSet) AddInstanceID(id *string) {
	flagName := "instance-id"

	fs.StringVar(
		id,
		flagName,
		"",
		fs.factory.Localizer.MustLocalize("kafka.common.flag.instanceID.description"),
	)
}

// AddAllAccounts adds a flag to set a wildcard for principals
func (fs *flagSet) AddAllAccounts(allAccounts *bool) {
	flagName := "all-accounts"

	fs.BoolVar(
		allAccounts,
		flagName,
		false,
		fs.factory.Localizer.MustLocalize("kafka.acl.common.flag.allAccounts.description"),
	)
}
