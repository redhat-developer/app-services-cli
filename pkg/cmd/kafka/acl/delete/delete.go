package delete

import (
	"context"

	"github.com/AlecAivazis/survey/v2"
	"github.com/redhat-developer/app-services-cli/internal/config"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/factory"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/flag"
	flagset "github.com/redhat-developer/app-services-cli/pkg/cmd/kafka/acl/flags"
	"github.com/redhat-developer/app-services-cli/pkg/cmdutil"
	"github.com/redhat-developer/app-services-cli/pkg/connection"
	"github.com/redhat-developer/app-services-cli/pkg/dump"
	"github.com/redhat-developer/app-services-cli/pkg/icon"
	"github.com/redhat-developer/app-services-cli/pkg/iostreams"
	"github.com/redhat-developer/app-services-cli/pkg/ioutil/spinner"
	"github.com/redhat-developer/app-services-cli/pkg/kafka/aclutil"
	"github.com/redhat-developer/app-services-cli/pkg/localize"
	"github.com/redhat-developer/app-services-cli/pkg/logging"
	kafkainstanceclient "github.com/redhat-developer/app-services-sdk-go/kafkainstance/apiv1internal/client"
	"github.com/spf13/cobra"
)

var (
	serviceAccount string
	userID         string
	allAccounts    bool
	prefix         bool
)

type requestParams struct {
	principal    string
	resourceName string
	resourceType kafkainstanceclient.AclResourceTypeFilter
	patternType  kafkainstanceclient.AclPatternTypeFilter
	operation    kafkainstanceclient.AclOperationFilter
	permission   kafkainstanceclient.AclPermissionTypeFilter
}

type options struct {
	config     config.IConfig
	connection factory.ConnectionFunc
	logger     logging.Logger
	io         *iostreams.IOStreams
	localizer  localize.Localizer
	context    context.Context

	cluster         bool
	patternType     string
	resourceType    string
	resourceName    string
	permission      string
	operation       string
	group           string
	topic           string
	transactionalID string
	principal       string

	skipConfirm bool
	output      string
	instanceID  string
}

// NewDeleteCommand creates a new command to delete Kafka ACLs
func NewDeleteCommand(f *factory.Factory) *cobra.Command {
	opts := &options{
		config:     f.Config,
		connection: f.Connection,
		logger:     f.Logger,
		io:         f.IOStreams,
		localizer:  f.Localizer,
		context:    f.Context,
	}

	cmd := &cobra.Command{
		Use:     "delete",
		Short:   f.Localizer.MustLocalize("kafka.acl.delete.cmd.shortDescription"),
		Long:    f.Localizer.MustLocalize("kafka.acl.delete.cmd.longDescription"),
		Example: f.Localizer.MustLocalize("kafka.acl.delete.cmd.example"),
		Args:    cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			if !opts.io.CanPrompt() && !opts.skipConfirm {
				return flag.RequiredWhenNonInteractiveError("yes")
			}

			if err := validateAndSetOpts(opts); err != nil {
				return err
			}

			return runDelete(opts.instanceID, opts)
		},
	}

	fs := flagset.NewFlagSet(cmd, opts.localizer, opts.connection)

	_ = fs.AddPermission(&opts.permission).Required()
	_ = fs.AddOperation(&opts.operation).Required()

	fs.AddCluster(&opts.cluster)
	fs.AddPrefix(&prefix)
	fs.AddTopic(&opts.topic)
	fs.AddConsumerGroup(&opts.group)
	fs.AddTransactionalID(&opts.transactionalID)
	fs.AddOutput(&opts.output)
	fs.AddInstanceID(&opts.instanceID)
	fs.AddUser(&userID)
	fs.AddServiceAccount(&serviceAccount)
	fs.AddAllAccounts(&allAccounts)
	fs.AddYes(&opts.skipConfirm)

	return cmd
}

// nolint:funlen
func runDelete(instanceID string, opts *options) error {
	ctx := opts.context

	conn, err := opts.connection(connection.DefaultConfigRequireMasAuth)
	if err != nil {
		return err
	}

	adminAPI, kafkaInstance, err := conn.API().KafkaAdmin(instanceID)
	if err != nil {
		return err
	}

	resourceOperations, httpRes, err := adminAPI.AclsApi.GetAclResourceOperations(ctx).Execute()
	if httpRes != nil {
		defer httpRes.Body.Close()
	}
	if err != nil {
		return err
	}

	if isValidOp, validResourceOperations := aclutil.IsValidResourceOperation(opts.resourceType, opts.operation, resourceOperations); !isValidOp {
		return opts.localizer.MustLocalizeError("kafka.acl.common.error.invalidResourceOperation",
			localize.NewEntry("ResourceType", opts.resourceType),
			localize.NewEntry("Operation", opts.operation),
			localize.NewEntry("ValidOperationList", cmdutil.StringSliceToListStringWithQuotes(validResourceOperations)),
		)
	}

	opts.patternType = aclutil.PatternTypeFilterLITERAL
	if prefix {
		opts.patternType = aclutil.PatternTypeFilterPREFIX
	}

	requestParams := getRequestParams(opts)

	aclList, httpRes, err := adminAPI.AclsApi.GetAcls(ctx).
		ResourceType(requestParams.resourceType).
		Principal(requestParams.principal).
		PatternType(requestParams.patternType).
		ResourceName(requestParams.resourceName).
		Operation(requestParams.operation).
		Permission(requestParams.permission).
		Execute()

	if httpRes != nil {
		defer httpRes.Body.Close()
	}

	if err = aclutil.ValidateAPIError(httpRes, opts.localizer, err, "list", kafkaInstance.GetName()); err != nil {
		return err
	}

	kafkaNameTmplEntry := localize.NewEntry("Name", kafkaInstance.GetName())
	if aclList.GetTotal() == 0 {
		opts.logger.Info(opts.localizer.MustLocalize("kafka.acl.common.log.info.noACLsMatchingFilters", kafkaNameTmplEntry))
		return nil
	}

	rows := aclutil.MapACLsToTableRows(aclList.GetItems(), opts.localizer)
	opts.logger.Info(icon.Warning(), opts.localizer.MustLocalize("kafka.acl.delete.log.info.theFollowingACLSwillBeDeleted", kafkaNameTmplEntry))
	opts.logger.Info()
	dump.Table(opts.io.ErrOut, rows)
	opts.logger.Info()

	if !opts.skipConfirm {
		prompt := &survey.Confirm{
			Message: opts.localizer.MustLocalize("kafka.acl.delete.input.confirmDeleteMessage"),
		}
		if err = survey.AskOne(prompt, &opts.skipConfirm); err != nil {
			return err
		}

		if !opts.skipConfirm {
			opts.logger.Debug("User has chosen to not delete ACLs")
			return nil
		}
	}

	opts.logger.Info()
	spinnr := spinner.New(opts.io.ErrOut, opts.localizer)
	spinnr.SetLocalizedSuffix("kafka.acl.delete.log.info.deletingACLs", kafkaNameTmplEntry)
	spinnr.Start()

	deletedACLs, httpRes, err := adminAPI.AclsApi.DeleteAcls(ctx).
		ResourceType(requestParams.resourceType).
		Principal(requestParams.principal).
		PatternType(requestParams.patternType).
		ResourceName(requestParams.resourceName).
		Operation(requestParams.operation).
		Permission(requestParams.permission).
		Execute()

	if httpRes != nil {
		defer httpRes.Body.Close()
	}

	if err = aclutil.ValidateAPIError(httpRes, opts.localizer, err, "delete", kafkaInstance.GetName()); err != nil {
		return err
	}

	spinnr.Stop()

	opts.logger.Info(icon.SuccessPrefix(), opts.localizer.MustLocalizePlural("kafka.acl.delete.successMessage",
		int(deletedACLs.GetTotal()),
		kafkaNameTmplEntry,
		localize.NewEntry("Count", aclList.GetTotal()),
	))

	return nil
}

func getRequestParams(opts *options) *requestParams {
	return &requestParams{
		resourceType: aclutil.GetMappedResourceTypeFilterValue(opts.resourceType),
		principal:    aclutil.FormatPrincipal(opts.principal),
		resourceName: opts.resourceName,
		patternType:  aclutil.GetMappedPatternTypeFilterValue(opts.patternType),
		operation:    aclutil.GetMappedOperationFilterValue(opts.operation),
		permission:   aclutil.GetMappedPermissionTypeFilterValue(opts.permission),
	}
}

func validateAndSetOpts(opts *options) error {
	var selectedResourceTypeCount int

	if opts.topic != "" {
		selectedResourceTypeCount++
		opts.resourceType = aclutil.ResourceTypeFilterTOPIC
		opts.resourceName = opts.topic
	}
	if opts.group != "" {
		selectedResourceTypeCount++
		opts.resourceType = aclutil.ResourceTypeFilterGROUP
		opts.resourceName = opts.group
	}
	if opts.transactionalID != "" {
		selectedResourceTypeCount++
		opts.resourceType = aclutil.ResourceTypeFilterTRANSACTIONAL_ID
		opts.resourceName = opts.transactionalID
	}
	if opts.cluster {
		selectedResourceTypeCount++
		opts.resourceType = aclutil.ResourceTypeFilterCLUSTER
		opts.resourceName = aclutil.KafkaCluster
	}

	resourceTypeFlagEntries := []*localize.TemplateEntry{
		localize.NewEntry("ClusterFlag", flagset.ClusterFlagName),
		localize.NewEntry("TopicFlag", flagset.TopicFlagName),
		localize.NewEntry("TransactionalIDFlag", flagset.TransactionalIDFlagName),
		localize.NewEntry("GroupFlag", flagset.GroupFlagName),
	}

	if selectedResourceTypeCount != 1 {
		return opts.localizer.MustLocalizeError("kafka.acl.common.error.oneResourceTypeAllowed", resourceTypeFlagEntries...)
	}

	// user and service account should not be provided together
	if userID != "" && serviceAccount != "" {
		return opts.localizer.MustLocalizeError("kafka.acl.common.error.bothPrincipalsSelected")
	}

	if userID == aclutil.Wildcard || serviceAccount == aclutil.Wildcard {
		return opts.localizer.MustLocalizeError("kafka.acl.common.error.useAllAccountsFlag")
	}

	if allAccounts {
		if userID != "" || serviceAccount != "" {
			return opts.localizer.MustLocalizeError("kafka.acl.common.error.allAccountsCannotBeUsedWithUserFlag")
		}
		opts.principal = aclutil.Wildcard
	}

	// check if priincipal is provided
	if !allAccounts && (userID == "" && serviceAccount == "") {
		return opts.localizer.MustLocalizeError("kafka.acl.common.error.noPrincipalsSelected")
	}

	if userID != "" {
		opts.principal = userID
	} else if serviceAccount != "" {
		opts.principal = serviceAccount
	}

	if opts.instanceID == "" {
		cfg, err := opts.config.Load()
		if err != nil {
			return err
		}

		instanceID, ok := cfg.GetKafkaIdOk()

		if !ok {
			return opts.localizer.MustLocalizeError("kafka.acl.common.error.noKafkaSelected")
		}

		opts.instanceID = instanceID
	}
	return nil
}
