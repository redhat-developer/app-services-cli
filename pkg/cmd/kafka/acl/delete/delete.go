package delete

import (
	"github.com/AlecAivazis/survey/v2"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/factory"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/flag"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/kafka/acl/flagutil"
	"github.com/redhat-developer/app-services-cli/pkg/cmdutil"
	"github.com/redhat-developer/app-services-cli/pkg/connection"
	"github.com/redhat-developer/app-services-cli/pkg/dump"
	"github.com/redhat-developer/app-services-cli/pkg/icon"
	"github.com/redhat-developer/app-services-cli/pkg/ioutil/spinner"
	"github.com/redhat-developer/app-services-cli/pkg/kafka/aclutil"
	"github.com/redhat-developer/app-services-cli/pkg/localize"
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

// NewDeleteCommand creates a new command to delete Kafka ACLs
func NewDeleteCommand(f *factory.Factory) *cobra.Command {
	opts := &aclutil.CrudOptions{
		Config:     f.Config,
		Connection: f.Connection,
		Logger:     f.Logger,
		IO:         f.IOStreams,
		Localizer:  f.Localizer,
		Context:    f.Context,
	}

	cmd := &cobra.Command{
		Use:     "delete",
		Short:   f.Localizer.MustLocalize("kafka.acl.delete.cmd.shortDescription"),
		Long:    f.Localizer.MustLocalize("kafka.acl.delete.cmd.longDescription"),
		Example: f.Localizer.MustLocalize("kafka.acl.delete.cmd.example"),
		Args:    cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			if !opts.IO.CanPrompt() && !opts.SkipConfirm {
				return flag.RequiredWhenNonInteractiveError("yes")
			}

			if err := aclutil.ValidateAndSetResources(opts, flagutil.ResourceTypeFlagEntries); err != nil {
				return err
			}

			if err := validateAndSetOpts(opts); err != nil {
				return err
			}

			return runDelete(opts.InstanceID, opts)
		},
	}

	flags := flagutil.NewFlagSet(cmd, f)

	_ = flags.AddPermissionFilter(&opts.Permission).Required()
	_ = flags.AddOperationFilter(&opts.Operation).Required()

	flags.AddCluster(&opts.Cluster)
	flags.AddPrefix(&prefix)
	flags.AddTopic(&opts.Topic)
	flags.AddConsumerGroup(&opts.Group)
	flags.AddTransactionalID(&opts.TransactionalID)
	flags.AddOutput(&opts.Output)
	flags.AddInstanceID(&opts.InstanceID)
	flags.AddUser(&userID)
	flags.AddServiceAccount(&serviceAccount)
	flags.AddAllAccounts(&allAccounts)
	flags.AddYes(&opts.SkipConfirm)

	return cmd
}

// nolint:funlen
func runDelete(instanceID string, opts *aclutil.CrudOptions) error {
	ctx := opts.Context

	conn, err := opts.Connection(connection.DefaultConfigRequireMasAuth)
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

	if isValidOp, validResourceOperations := aclutil.IsValidResourceOperation(opts.ResourceType, opts.Operation, resourceOperations); !isValidOp {
		return opts.Localizer.MustLocalizeError("kafka.acl.common.error.invalidResourceOperation",
			localize.NewEntry("ResourceType", opts.ResourceType),
			localize.NewEntry("Operation", opts.Operation),
			localize.NewEntry("ValidOperationList", cmdutil.StringSliceToListStringWithQuotes(validResourceOperations)),
		)
	}

	kafkaNameTmplEntry := localize.NewEntry("Name", kafkaInstance.GetName())

	if !opts.SkipConfirm {
		prompt := &survey.Confirm{
			Message: opts.Localizer.MustLocalize("kafka.acl.delete.input.confirmDeleteMessage", kafkaNameTmplEntry),
		}
		if err = survey.AskOne(prompt, &opts.SkipConfirm); err != nil {
			return err
		}

		if !opts.SkipConfirm {
			opts.Logger.Debug("User has chosen to not delete ACLs")
			return nil
		}
	}

	opts.Logger.Info()
	spinnr := spinner.New(opts.IO.ErrOut, opts.Localizer)
	spinnr.SetLocalizedSuffix("kafka.acl.delete.log.info.deletingACLs", kafkaNameTmplEntry)
	spinnr.Start()

	requestParams := getRequestParams(opts)

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

	err = aclutil.ValidateAPIError(httpRes, opts.Localizer, err, "delete", kafkaInstance.GetName())
	spinnr.Stop()

	if err != nil {
		return err
	}

	deletedCount := int(deletedACLs.GetTotal())

	if deletedCount == 0 {
		opts.Logger.Info(icon.InfoPrefix(), opts.Localizer.MustLocalize("kafka.acl.delete.noACLsDeleted", kafkaNameTmplEntry))
		return nil
	}

	opts.Logger.Info(icon.SuccessPrefix(), opts.Localizer.MustLocalizePlural("kafka.acl.delete.successMessage",
		deletedCount,
		kafkaNameTmplEntry,
		localize.NewEntry("Count", deletedCount),
	))

	opts.Logger.Info(opts.Localizer.MustLocalize("kafka.acl.grantPermissions.log.delete.info.aclsPreview"))
	opts.Logger.Info()
	rows := aclutil.MapACLsToTableRows(*deletedACLs.Items, opts.Localizer)
	dump.Table(opts.IO.Out, rows)
	opts.Logger.Info()

	return nil
}

func getRequestParams(opts *aclutil.CrudOptions) *requestParams {
	return &requestParams{
		resourceType: aclutil.GetMappedResourceTypeFilterValue(opts.ResourceType),
		principal:    aclutil.FormatPrincipal(opts.Principal),
		resourceName: aclutil.GetResourceName(opts.ResourceName),
		patternType:  aclutil.GetMappedPatternTypeFilterValue(opts.PatternType),
		operation:    aclutil.GetMappedOperationFilterValue(opts.Operation),
		permission:   aclutil.GetMappedPermissionTypeFilterValue(opts.Permission),
	}
}

func validateAndSetOpts(opts *aclutil.CrudOptions) error {

	// user and service account should not be provided together
	if userID != "" && serviceAccount != "" {
		return opts.Localizer.MustLocalizeError("kafka.acl.common.error.bothPrincipalsSelected")
	}

	if userID == aclutil.Wildcard || serviceAccount == aclutil.Wildcard || userID == aclutil.AllAlias || serviceAccount == aclutil.AllAlias {
		return opts.Localizer.MustLocalizeError("kafka.acl.common.error.useAllAccountsFlag")
	}

	if allAccounts {
		if userID != "" || serviceAccount != "" {
			return opts.Localizer.MustLocalizeError("kafka.acl.common.error.allAccountsCannotBeUsedWithUserFlag")
		}
		opts.Principal = aclutil.Wildcard
	}

	// check if principal is provided
	if !allAccounts && (userID == "" && serviceAccount == "") {
		return opts.Localizer.MustLocalizeError("kafka.acl.common.error.noPrincipalsSelected")
	}

	opts.PatternType = aclutil.PatternTypeLITERAL
	if prefix {
		opts.PatternType = aclutil.PatternTypePREFIX
	}

	if userID != "" {
		opts.Principal = userID
	} else if serviceAccount != "" {
		opts.Principal = serviceAccount
	}

	if opts.InstanceID == "" {
		cfg, err := opts.Config.Load()
		if err != nil {
			return err
		}

		instanceID, ok := cfg.GetKafkaIdOk()

		if !ok {
			return opts.Localizer.MustLocalizeError("kafka.acl.common.error.noKafkaSelected")
		}

		opts.InstanceID = instanceID
	}

	return nil
}
