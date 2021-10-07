package common

import (
	"github.com/redhat-developer/app-services-cli/pkg/cmdutil/flags"
	"github.com/redhat-developer/app-services-cli/pkg/localize"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

const (
	ResourceTypeFilterANY              = "any"
	ResourceTypeFilterTOPIC            = "topic"
	ResourceTypeFilterCLUSTER          = "cluster"
	ResourceTypeFilterGROUP            = "group"
	ResourceTypeFilterTRANSACTIONAL_ID = "transactional-id"
)

const (
	PermissionALLOW = "allow"
	PermissionDENY  = "deny"
	PermissionANY   = "any"
)

const (
	OperationFilterALL              = "all"
	OperationFilterANY              = "any"
	OperationFilterREAD             = "read"
	OperationFilterWRITE            = "write"
	OperationFilterCREATE           = "create"
	OperationFilterDELETE           = "delete"
	OperationFilterALTER            = "alter"
	OperationFilterDESCRIBE         = "describe"
	OperationFilterDESCRIBE_CONFIGS = "describe-config"
	OperationFilterALTER_CONFIGS    = "alter-configs"
)

const (
	PatternTypeFilterLITERAL = "literal"
	PatternTypeFilterPREFIX  = "prefix"
	PatternTypeFilterANY     = "any"
)

const (
	TopicFlagName       = "topic"
	GroupFlagName       = "group"
	TransactionalIDFlag = "transactional-id"
)

type flagSet struct {
	flags     *pflag.FlagSet
	cmd       *cobra.Command
	localizer localize.Localizer
	*flags.FlagSet
}

func NewFlagSet(cmd *cobra.Command, localizer localize.Localizer) *flagSet {
	return &flagSet{
		cmd:       cmd,
		flags:     cmd.Flags(),
		localizer: localizer,
		FlagSet:   flags.NewFlagSet(cmd, localizer),
	}
}

// AddResourceType adds a flag for ACL resource type and registers completion options
func (fs *flagSet) AddResourceType(resourceType *string) *markRequiredOpt {
	flagName := "resource-type"

	resourceTypes := make([]string, 0, len(resourceTypeFilterMap))
	for i := range resourceTypeFilterMap {
		resourceTypes = append(resourceTypes, i)
	}

	fs.flags.StringVar(
		resourceType,
		flagName,
		ResourceTypeFilterANY,
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

	operations := make([]string, 0, len(operationFilterMap))
	for i := range operationFilterMap {
		operations = append(operations, i)
	}

	fs.flags.StringVar(
		operationType,
		flagName,
		ResourceTypeFilterANY,
		flags.FlagDescription(fs.localizer, "kafka.acl.common.flag.operation", operations...),
	)

	_ = fs.cmd.RegisterFlagCompletionFunc(flagName, func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return operations, cobra.ShellCompDirectiveNoSpace
	})

	return withMarkRequiredFunc(fs.cmd, flagName)
}

// AddPermission adds a flag for ACL permissions
func (fs *flagSet) AddPermission(permission *string) *markRequiredOpt {
	flagName := "permission"

	permissions := make([]string, 0, len(permissionTypeFilterMap))
	for i := range permissionTypeFilterMap {
		permissions = append(permissions, i)
	}

	fs.flags.StringVar(
		permission,
		flagName,
		ResourceTypeFilterANY,
		flags.FlagDescription(fs.localizer, "kafka.acl.common.flag.permission", permissions...),
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
		"Topic name",
	)
}

// AddConsumerGroup adds a flag for setting the consumer group ID
func (fs *flagSet) AddConsumerGroup(group *string) {
	flagName := GroupFlagName

	fs.flags.StringVar(
		group,
		flagName,
		"",
		"Consumer group ID",
	)
}

// AddTransactionalID adds a flag for setting the consumer group ID
func (fs *flagSet) AddTransactionalID(id *string) {
	flagName := TransactionalIDFlag

	fs.flags.StringVar(
		id,
		flagName,
		"",
		"Transactional ID",
	)
}

// AddPrefix adds a flag for enabling the PREFIX resource pattern type
func (fs *flagSet) AddPrefix(prefix *bool) {
	flagName := "prefix"

	fs.flags.BoolVar(
		prefix,
		flagName,
		false,
		"Determine if the resource should be exact match or prefix",
	)
}

// AddPatternType adds a flag to choose the ACL resource pattern type
func (fs *flagSet) AddPatternType(patternType *string) *markRequiredOpt {
	flagName := "pattern-type"

	// TODO: Add available options to description
	fs.flags.StringVar(
		patternType,
		flagName,
		"",
		"Determine the pattern type of the ACL resource",
	)

	patternTypes := make([]string, 0, len(patternTypeFilterMap))
	for i := range patternTypeFilterMap {
		patternTypes = append(patternTypes, i)
	}

	_ = fs.cmd.RegisterFlagCompletionFunc(flagName, func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return patternTypes, cobra.ShellCompDirectiveNoSpace
	})

	return withMarkRequiredFunc(fs.cmd, flagName)
}

// AddUser adds a flag to pass a user ID principal
func (fs *flagSet) AddUser(userID *string) *markRequiredOpt {
	flagName := "user"

	fs.flags.StringVar(
		userID,
		flagName,
		"",
		"User account principal for this operation",
	)

	return withMarkRequiredFunc(fs.cmd, flagName)
}

// AddUser adds a flag to pass a user ID principal
func (fs *flagSet) AddServiceAccount(serviceAccountID *string) *markRequiredOpt {
	flagName := "service-account"

	fs.flags.StringVar(
		serviceAccountID,
		flagName,
		"",
		"Service account principal for this operation",
	)

	return withMarkRequiredFunc(fs.cmd, flagName)
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
