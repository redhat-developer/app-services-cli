package delete

import (
	"github.com/AlecAivazis/survey/v2"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/kafka/acl/aclcmdutil"
	aclFlagUtil "github.com/redhat-developer/app-services-cli/pkg/cmd/kafka/acl/flagutil"
	"github.com/redhat-developer/app-services-cli/pkg/core/cmdutil"
	"github.com/redhat-developer/app-services-cli/pkg/core/cmdutil/flagutil"
	"github.com/redhat-developer/app-services-cli/pkg/core/ioutil/dump"
	"github.com/redhat-developer/app-services-cli/pkg/core/ioutil/icon"
	"github.com/redhat-developer/app-services-cli/pkg/core/ioutil/spinner"
	"github.com/redhat-developer/app-services-cli/pkg/core/localize"
	"github.com/redhat-developer/app-services-cli/pkg/shared/connection"
	"github.com/redhat-developer/app-services-cli/pkg/shared/factory"
	kafkainstanceclient "github.com/redhat-developer/app-services-sdk-go/kafkainstance/apiv1internal/client"
	"github.com/spf13/cobra"
)

var (
	serviceAccount  string
	userID          string
	allAccounts     bool
	prefix          bool
	patternTypeFlag string
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
	opts := &aclcmdutil.CrudOptions{
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
				return flagutil.RequiredWhenNonInteractiveError("yes")
			}

			var errorCollection []error

			selectedResourceTypeCount := aclcmdutil.SetACLResources(opts)
			if selectedResourceTypeCount > 1 {
				errorCollection = append(errorCollection,
					opts.Localizer.MustLocalizeError("kafka.acl.common.error.oneResourceTypeAllowed", aclFlagUtil.ResourceTypeFlagEntries...))
			}

			if principalErrors := validateAndSetOpts(opts); principalErrors != nil {
				errorCollection = append(errorCollection, principalErrors)
			}

			if len(errorCollection) > 0 {
				return aclcmdutil.BuildInstructions(errorCollection)
			}

			return runDelete(opts.InstanceID, opts)
		},
	}

	flags := aclFlagUtil.NewFlagSet(cmd, f)

	flags.AddPermissionFilter(&opts.Permission)
	flags.AddOperationFilter(&opts.Operation)

	flags.AddCluster(&opts.Cluster)

	flags.AddTopic(&opts.Topic)
	flags.AddConsumerGroup(&opts.Group)
	flags.AddTransactionalID(&opts.TransactionalID)
	flags.AddOutput(&opts.Output)
	flags.AddInstanceID(&opts.InstanceID)
	flags.AddUser(&userID)
	flags.AddServiceAccount(&serviceAccount)
	flags.AddAllAccounts(&allAccounts)
	flags.AddYes(&opts.SkipConfirm)

	cmd.Flags().BoolVar(
		&prefix,
		"prefix",
		false,
		flagutil.DeprecateFlag(opts.Localizer.MustLocalize("kafka.acl.common.flag.delete.prefix.description")),
	)

	cmd.Flags().StringVar(
		&patternTypeFlag,
		"pattern-type",
		aclcmdutil.PatternTypeLITERAL,
		opts.Localizer.MustLocalize("kafka.acl.common.flag.patterntypes.description",
			localize.NewEntry("Types", aclcmdutil.PatternTypes)),
	)

	cmd.RegisterFlagCompletionFunc("pattern-type", func(cmd *cobra.Command, _ []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return aclcmdutil.PatternTypes, cobra.ShellCompDirectiveNoSpace
	})

	return cmd
}

// nolint:funlen
func runDelete(instanceID string, opts *aclcmdutil.CrudOptions) error {
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

	// Validate only when both are present
	if opts.ResourceType != "" && opts.Operation != "" {
		if isValidOp, validResourceOperations := aclcmdutil.IsValidResourceOperation(opts.ResourceType, opts.Operation, resourceOperations); !isValidOp {
			return opts.Localizer.MustLocalizeError("kafka.acl.common.error.invalidResourceOperation",
				localize.NewEntry("ResourceType", opts.ResourceType),
				localize.NewEntry("Operation", opts.Operation),
				localize.NewEntry("ValidOperationList", cmdutil.StringSliceToListStringWithQuotes(validResourceOperations)),
			)
		}
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

	requestDeleteAcls := adminAPI.AclsApi.DeleteAcls(ctx)
	if requestParams.resourceType != "" {
		requestDeleteAcls = requestDeleteAcls.ResourceType(requestParams.resourceType)
	}

	if requestParams.principal != "" {
		requestDeleteAcls = requestDeleteAcls.Principal(requestParams.principal)
	}

	if requestParams.resourceName != "" {
		requestDeleteAcls = requestDeleteAcls.ResourceName(requestParams.resourceName)
	}
	if requestParams.patternType != "" {
		requestDeleteAcls = requestDeleteAcls.PatternType(requestParams.patternType)
	}
	if requestParams.operation != "" {
		requestDeleteAcls = requestDeleteAcls.Operation(requestParams.operation)
	}
	if requestParams.permission != "" {
		requestDeleteAcls = requestDeleteAcls.Permission(requestParams.permission)
	}

	deletedACLs, httpRes, err := requestDeleteAcls.Execute()

	if httpRes != nil {
		defer httpRes.Body.Close()
	}

	err = aclcmdutil.ValidateAPIError(httpRes, opts.Localizer, err, "delete", kafkaInstance.GetName())
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

	rows := aclcmdutil.MapACLsToTableRows(*deletedACLs.Items, opts.Localizer)
	opts.Logger.Info(opts.Localizer.MustLocalizePlural("kafka.acl.grantPermissions.log.delete.info.aclsPreview", len(rows)))
	opts.Logger.Info()

	dump.Table(opts.IO.Out, rows)
	opts.Logger.Info()

	return nil
}

func getRequestParams(opts *aclcmdutil.CrudOptions) *requestParams {
	return &requestParams{
		resourceType: aclcmdutil.GetMappedResourceTypeFilterValue(opts.ResourceType),
		principal:    aclcmdutil.FormatPrincipal(opts.Principal),
		resourceName: aclcmdutil.GetResourceName(opts.ResourceName),
		patternType:  aclcmdutil.GetMappedPatternTypeFilterValue(opts.PatternType),
		operation:    aclcmdutil.GetMappedOperationFilterValue(opts.Operation),
		permission:   aclcmdutil.GetMappedPermissionTypeFilterValue(opts.Permission),
	}
}

func validateAndSetOpts(opts *aclcmdutil.CrudOptions) error {

	// user and service account should not be provided together
	if userID != "" && serviceAccount != "" {
		return opts.Localizer.MustLocalizeError("kafka.acl.common.error.bothPrincipalsSelected")
	}

	if userID == aclcmdutil.Wildcard || serviceAccount == aclcmdutil.Wildcard || userID == aclcmdutil.AllAlias || serviceAccount == aclcmdutil.AllAlias {
		return opts.Localizer.MustLocalizeError("kafka.acl.common.error.useAllAccountsFlag")
	}

	if allAccounts {
		if userID != "" || serviceAccount != "" {
			return opts.Localizer.MustLocalizeError("kafka.acl.common.error.allAccountsCannotBeUsedWithUserFlag")
		}
		opts.Principal = aclcmdutil.Wildcard
	}

	// check if principal is provided
	if !allAccounts && (userID == "" && serviceAccount == "") {
		return opts.Localizer.MustLocalizeError("kafka.acl.common.error.noPrincipalsSelected")
	}

	// Backwards compatibility:
	if prefix {
		opts.PatternType = aclcmdutil.PatternTypePREFIX
	} else if patternTypeFlag == aclcmdutil.PatternTypeANY {
		opts.PatternType = aclcmdutil.PatternTypeANY
	} else {
		opts.PatternType = aclcmdutil.PatternTypeLITERAL
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
