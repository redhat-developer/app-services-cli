package flagset

import (
	"github.com/redhat-developer/app-services-cli/pkg/cmd/factory"
	"github.com/redhat-developer/app-services-cli/pkg/cmdutil/flags"
	"github.com/redhat-developer/app-services-cli/pkg/kafka/aclutil"
	"github.com/redhat-developer/app-services-cli/pkg/localize"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

const (
	ClusterFlagName         = "cluster"
	TopicFlagName           = "topic"
	GroupFlagName           = "group"
	TransactionalIDFlagName = "transactional-id"
)

type flagSet struct {
	flags     *pflag.FlagSet
	cmd       *cobra.Command
	localizer localize.Localizer
	conn      factory.ConnectionFunc
	*flags.FlagSet
}

func NewFlagSet(cmd *cobra.Command, localizer localize.Localizer, conn factory.ConnectionFunc) *flagSet {
	return &flagSet{
		cmd:       cmd,
		flags:     cmd.Flags(),
		localizer: localizer,
		conn:      conn,
		FlagSet:   flags.NewFlagSet(cmd, localizer),
	}
}

// AddResourceType adds a flag for ACL resource type and registers completion options
func (fs *flagSet) AddResourceType(resourceType *string) *markRequiredOpt {
	flagName := "resource-type"

	resourceTypeFilterMap := aclutil.GetResourceTypeFilterMap()

	resourceTypes := make([]string, 0, len(resourceTypeFilterMap))
	for i := range resourceTypeFilterMap {
		resourceTypes = append(resourceTypes, i)
	}

	fs.flags.StringVar(
		resourceType,
		flagName,
		aclutil.ResourceTypeFilterANY,
		flags.FlagDescription(fs.localizer, "kafka.acl.common.flag.resourceType", resourceTypes...),
	)

	_ = fs.cmd.RegisterFlagCompletionFunc(flagName, func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return resourceTypes, cobra.ShellCompDirectiveNoSpace
	})

	return withMarkRequiredFunc(fs.cmd, flagName)
}

// AddOperation adds a flag for ACL operations and registers completion options
func (fs *flagSet) AddOperation(operationType *string) *markRequiredOpt {
	flagName := "operation"

	operationFilterMap := aclutil.GetOperationFilterMap()

	operations := make([]string, 0, len(operationFilterMap))
	for i := range operationFilterMap {
		operations = append(operations, i)
	}

	fs.flags.StringVar(
		operationType,
		flagName,
		aclutil.ResourceTypeFilterANY,
		flags.FlagDescription(fs.localizer, "kafka.acl.common.flag.operation.description", operations...),
	)

	_ = fs.cmd.RegisterFlagCompletionFunc(flagName, func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return operations, cobra.ShellCompDirectiveNoSpace
	})

	return withMarkRequiredFunc(fs.cmd, flagName)
}

// AddPermission adds a flag for ACL permissions
func (fs *flagSet) AddPermission(permission *string) *markRequiredOpt {
	flagName := "permission"

	permissionTypeFilterMap := aclutil.GetPermissionTypeFilterMap()

	permissions := make([]string, 0, len(permissionTypeFilterMap))
	for i := range permissionTypeFilterMap {
		permissions = append(permissions, i)
	}

	fs.flags.StringVar(
		permission,
		flagName,
		aclutil.ResourceTypeFilterANY,
		flags.FlagDescription(fs.localizer, "kafka.acl.common.flag.permission.description", permissions...),
	)

	_ = fs.cmd.RegisterFlagCompletionFunc(flagName, func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return permissions, cobra.ShellCompDirectiveNoSpace
	})

	return withMarkRequiredFunc(fs.cmd, flagName)
}

// AddTopic adds a flag for setting the topic name
func (fs *flagSet) AddTopic(topic *string) {
	flagName := TopicFlagName

	fs.flags.StringVar(
		topic,
		flagName,
		"",
		fs.localizer.MustLocalize("kafka.acl.common.flag.topic.description"),
	)
}

// AddConsumerGroup adds a flag for setting the consumer group ID
func (fs *flagSet) AddConsumerGroup(group *string) {
	flagName := GroupFlagName

	fs.flags.StringVar(
		group,
		flagName,
		"",
		fs.localizer.MustLocalize("kafka.acl.common.flag.group.description"),
	)
}

// AddTransactionalID adds a flag for setting the consumer group ID
func (fs *flagSet) AddTransactionalID(id *string) {
	flagName := TransactionalIDFlagName

	fs.flags.StringVar(
		id,
		flagName,
		"",
		fs.localizer.MustLocalize("kafka.acl.common.flag.transactionalID.description"),
	)
}

// AddPrefix adds a flag for enabling the PREFIX resource pattern type
func (fs *flagSet) AddPrefix(prefix *bool) {
	flagName := "prefix"

	fs.flags.BoolVar(
		prefix,
		flagName,
		false,
		fs.localizer.MustLocalize("kafka.acl.common.flag.prefix.description"),
	)
}

// AddPrefix adds a flag for sertting the "cluster" ACL resource type
func (fs *flagSet) AddCluster(prefix *bool) {
	flagName := ClusterFlagName

	fs.flags.BoolVar(
		prefix,
		flagName,
		false,
		fs.localizer.MustLocalize("kafka.acl.common.flag.cluster.description"),
	)
}

// AddUser adds a flag to pass a user ID principal
func (fs *flagSet) AddUser(userID *string) *markRequiredOpt {
	flagName := "user"

	fs.flags.StringVar(
		userID,
		flagName,
		"",
		fs.localizer.MustLocalize("kafka.acl.common.flag.user.description"),
	)

	_ = flags.RegisterUserCompletionFunc(fs.cmd, flagName, fs.conn)

	return withMarkRequiredFunc(fs.cmd, flagName)
}

// AddUser adds a flag to pass a user ID principal
func (fs *flagSet) AddServiceAccount(serviceAccountID *string) *markRequiredOpt {
	flagName := "service-account"

	fs.flags.StringVar(
		serviceAccountID,
		flagName,
		"",
		fs.localizer.MustLocalize("kafka.acl.common.flag.serviceAccount.description"),
	)

	_ = flags.RegisterServiceAccountCompletionFunc(fs.cmd, flagName, fs.conn)

	return withMarkRequiredFunc(fs.cmd, flagName)
}

// AddInstanceID adds a flag for the Kafka instance ID
func (fs *flagSet) AddInstanceID(id *string) {
	flagName := "instance-id"

	fs.flags.StringVar(
		id,
		flagName,
		"",
		fs.localizer.MustLocalize("kafka.common.flag.instanceID.description"),
	)
}

// AddAllAccounts adds a flag to set a wildcard for principals
func (fs *flagSet) AddAllAccounts(allAccounts *bool) {
	flagName := "all-accounts"

	fs.flags.BoolVar(
		allAccounts,
		flagName,
		false,
		fs.localizer.MustLocalize("kafka.acl.common.flag.allAccounts.description"),
	)
}

func withMarkRequiredFunc(cmd *cobra.Command, flagName string) *markRequiredOpt {
	return &markRequiredOpt{
		Required: func() error {
			return cmd.MarkFlagRequired(flagName)
		},
	}
}

type markRequiredOpt struct {
	Required func() error
}
