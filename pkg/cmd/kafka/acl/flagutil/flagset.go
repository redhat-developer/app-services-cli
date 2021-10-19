package flagutil

import (
	"github.com/redhat-developer/app-services-cli/pkg/cmd/factory"
	"github.com/redhat-developer/app-services-cli/pkg/cmdutil/flagutil"
	"github.com/redhat-developer/app-services-cli/pkg/kafka/aclutil"
	"github.com/redhat-developer/app-services-cli/pkg/localize"
	"github.com/spf13/cobra"
)

const (
	ClusterFlagName         = "cluster"
	TopicFlagName           = "topic"
	GroupFlagName           = "group"
	TransactionalIDFlagName = "transactional-id"
)

type flagSet struct {
	cmd       *cobra.Command
	localizer localize.Localizer
	conn      factory.ConnectionFunc
	*flagutil.FlagSet
}

// NewFlagSet returns a new flag set with common Kafka ACL flags
func NewFlagSet(cmd *cobra.Command, localizer localize.Localizer, conn factory.ConnectionFunc) *flagSet {
	return &flagSet{
		cmd:       cmd,
		localizer: localizer,
		conn:      conn,
		FlagSet:   flagutil.NewFlagSet(cmd, localizer),
	}
}

// AddResourceType adds a flag for ACL resource type and registers completion options
func (fs *flagSet) AddResourceType(resourceType *string) *flagutil.FlagOptions {
	flagName := "resource-type"

	resourceTypeFilterMap := aclutil.GetResourceTypeFilterMap()

	resourceTypes := make([]string, 0, len(resourceTypeFilterMap))
	for i := range resourceTypeFilterMap {
		resourceTypes = append(resourceTypes, i)
	}

	fs.StringVar(
		resourceType,
		flagName,
		aclutil.ResourceTypeFilterANY,
		flagutil.FlagDescription(fs.localizer, "kafka.acl.common.flag.resourceType", resourceTypes...),
	)

	_ = fs.cmd.RegisterFlagCompletionFunc(flagName, func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return resourceTypes, cobra.ShellCompDirectiveNoSpace
	})

	return flagutil.WithFlagOptions(fs.cmd, flagName)
}

// AddOperation adds a flag for ACL operations and registers completion options
func (fs *flagSet) AddOperation(operationType *string) *flagutil.FlagOptions {
	flagName := "operation"

	operationFilterMap := aclutil.GetOperationFilterMap()

	operations := make([]string, 0, len(operationFilterMap))
	for i := range operationFilterMap {
		operations = append(operations, i)
	}

	fs.StringVar(
		operationType,
		flagName,
		aclutil.ResourceTypeFilterANY,
		flagutil.FlagDescription(fs.localizer, "kafka.acl.common.flag.operation.description", operations...),
	)

	_ = fs.cmd.RegisterFlagCompletionFunc(flagName, func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return operations, cobra.ShellCompDirectiveNoSpace
	})

	return flagutil.WithFlagOptions(fs.cmd, flagName)
}

// AddPermission adds a flag for ACL permissions
func (fs *flagSet) AddPermission(permission *string) *flagutil.FlagOptions {
	flagName := "permission"

	permissionTypeFilterMap := aclutil.GetPermissionTypeFilterMap()

	permissions := make([]string, 0, len(permissionTypeFilterMap))
	for i := range permissionTypeFilterMap {
		permissions = append(permissions, i)
	}

	fs.StringVar(
		permission,
		flagName,
		aclutil.ResourceTypeFilterANY,
		flagutil.FlagDescription(fs.localizer, "kafka.acl.common.flag.permission.description", permissions...),
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
		fs.localizer.MustLocalize("kafka.acl.common.flag.topic.description"),
	)
}

// AddConsumerGroup adds a flag for setting the consumer group ID
func (fs *flagSet) AddConsumerGroup(group *string) {
	flagName := GroupFlagName

	fs.StringVar(
		group,
		flagName,
		"",
		fs.localizer.MustLocalize("kafka.acl.common.flag.group.description"),
	)
}

// AddTransactionalID adds a flag for setting the consumer group ID
func (fs *flagSet) AddTransactionalID(id *string) {
	flagName := TransactionalIDFlagName

	fs.StringVar(
		id,
		flagName,
		"",
		fs.localizer.MustLocalize("kafka.acl.common.flag.transactionalID.description"),
	)
}

// AddPrefix adds a flag for enabling the PREFIX resource pattern type
func (fs *flagSet) AddPrefix(prefix *bool) {
	flagName := "prefix"

	fs.BoolVar(
		prefix,
		flagName,
		false,
		fs.localizer.MustLocalize("kafka.acl.common.flag.prefix.description"),
	)
}

// AddCluster adds a flag for setting the "cluster" ACL resource type
func (fs *flagSet) AddCluster(prefix *bool) {
	flagName := ClusterFlagName

	fs.BoolVar(
		prefix,
		flagName,
		false,
		fs.localizer.MustLocalize("kafka.acl.common.flag.cluster.description"),
	)
}

// AddUser adds a flag to pass a user ID principal
func (fs *flagSet) AddUser(userID *string) *flagutil.FlagOptions {
	flagName := "user"

	fs.StringVar(
		userID,
		flagName,
		"",
		fs.localizer.MustLocalize("kafka.acl.common.flag.user.description"),
	)

	_ = flagutil.RegisterUserCompletionFunc(fs.cmd, flagName, fs.conn)

	return flagutil.WithFlagOptions(fs.cmd, flagName)
}

// AddServiceAccount adds a flag to pass a service account client ID principal
func (fs *flagSet) AddServiceAccount(serviceAccountID *string) *flagutil.FlagOptions {
	flagName := "service-account"

	fs.StringVar(
		serviceAccountID,
		flagName,
		"",
		fs.localizer.MustLocalize("kafka.acl.common.flag.serviceAccount.description"),
	)

	_ = flagutil.RegisterServiceAccountCompletionFunc(fs.cmd, flagName, fs.conn)

	return flagutil.WithFlagOptions(fs.cmd, flagName)
}

// AddInstanceID adds a flag for the Kafka instance ID
func (fs *flagSet) AddInstanceID(id *string) {
	flagName := "instance-id"

	fs.StringVar(
		id,
		flagName,
		"",
		fs.localizer.MustLocalize("kafka.common.flag.instanceID.description"),
	)
}

// AddAllAccounts adds a flag to set a wildcard for principals
func (fs *flagSet) AddAllAccounts(allAccounts *bool) {
	flagName := "all-accounts"

	fs.BoolVar(
		allAccounts,
		flagName,
		false,
		fs.localizer.MustLocalize("kafka.acl.common.flag.allAccounts.description"),
	)
}