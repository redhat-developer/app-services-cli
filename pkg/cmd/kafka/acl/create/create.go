package create

import (
	"github.com/AlecAivazis/survey/v2"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/kafka/acl/aclcmdutil"
	aclFlagUtil "github.com/redhat-developer/app-services-cli/pkg/cmd/kafka/acl/flagutil"
	"github.com/redhat-developer/app-services-cli/pkg/core/cmdutil"
	"github.com/redhat-developer/app-services-cli/pkg/core/cmdutil/flagutil"
	"github.com/redhat-developer/app-services-cli/pkg/core/ioutil/dump"
	"github.com/redhat-developer/app-services-cli/pkg/core/ioutil/spinner"
	"github.com/redhat-developer/app-services-cli/pkg/core/localize"
	"github.com/redhat-developer/app-services-cli/pkg/shared/connection"
	"github.com/redhat-developer/app-services-cli/pkg/shared/contextutil"
	"github.com/redhat-developer/app-services-cli/pkg/shared/factory"
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
	resourceType kafkainstanceclient.AclResourceType
	patternType  kafkainstanceclient.AclPatternType
	operation    kafkainstanceclient.AclOperation
	permission   kafkainstanceclient.AclPermissionType
}

// NewCreateCommand creates a new command to add Kafka ACLs
func NewCreateCommand(f *factory.Factory) *cobra.Command {
	opts := &aclcmdutil.CrudOptions{
		Connection:     f.Connection,
		Logger:         f.Logger,
		IO:             f.IOStreams,
		Localizer:      f.Localizer,
		Context:        f.Context,
		ServiceContext: f.ServiceContext,
	}

	cmd := &cobra.Command{
		Use:     "create",
		Short:   f.Localizer.MustLocalize("kafka.acl.create.cmd.shortDescription"),
		Long:    f.Localizer.MustLocalize("kafka.acl.create.cmd.longDescription"),
		Example: f.Localizer.MustLocalize("kafka.acl.create.cmd.example"),
		Args:    cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			if !opts.IO.CanPrompt() && !opts.SkipConfirm {
				return flagutil.RequiredWhenNonInteractiveError("yes")
			}

			var errorCollection []error

			if opts.Permission == "" {
				errorCollection = append(errorCollection, opts.Localizer.MustLocalizeError("kafka.acl.common.flag.permission.required"))
			}

			if opts.Operation == "" {
				errorCollection = append(errorCollection, opts.Localizer.MustLocalizeError("kafka.acl.common.flag.operation.required"))
			}

			selectedResourceTypeCount := aclcmdutil.SetACLResources(opts)

			if selectedResourceTypeCount != 1 {
				errorCollection = append(errorCollection, opts.Localizer.MustLocalizeError("kafka.acl.common.error.oneResourceTypeAllowed", aclFlagUtil.ResourceTypeFlagEntries...))
			}

			if principalErrors := validateAndSetOpts(opts); principalErrors != nil {
				errorCollection = append(errorCollection, principalErrors)
			}

			if len(errorCollection) > 0 {
				return aclcmdutil.BuildInstructions(errorCollection)
			}

			if opts.InstanceID == "" {

				kafkaInstance, err := contextutil.GetCurrentKafkaInstance(f)
				if err != nil {
					return err
				}

				opts.InstanceID = kafkaInstance.GetId()
			}

			return runAdd(opts.InstanceID, opts)
		},
	}

	flags := aclFlagUtil.NewFlagSet(cmd, f)

	flags.AddPermissionCreate(&opts.Permission)
	flags.AddOperationCreate(&opts.Operation)

	flags.AddCluster(&opts.Cluster)
	flags.AddPrefix(&prefix)
	flags.AddTopic(&opts.Topic)
	flags.AddConsumerGroup(&opts.Group)
	flags.AddTransactionalID(&opts.TransactionalID)
	flags.AddInstanceID(&opts.InstanceID)
	flags.AddUser(&userID)
	flags.AddServiceAccount(&serviceAccount)
	flags.AddAllAccounts(&allAccounts)
	flags.AddYes(&opts.SkipConfirm)

	return cmd
}

// nolint:funlen
func runAdd(instanceID string, opts *aclcmdutil.CrudOptions) error {
	ctx := opts.Context

	conn, err := opts.Connection(connection.DefaultConfigRequireMasAuth)
	if err != nil {
		return err
	}

	adminAPI, kafkaInstance, err := conn.API().KafkaAdmin(instanceID)
	if err != nil {
		return err
	}

	kafkaName := kafkaInstance.GetName()

	resourceOperations, httpRes, err := adminAPI.AclsApi.GetAclResourceOperations(ctx).Execute()
	if httpRes != nil {
		defer httpRes.Body.Close()
	}
	if err != nil {
		return err
	}

	if isValidOp, validResourceOperations := aclcmdutil.IsValidResourceOperation(opts.ResourceType, opts.Operation, resourceOperations); !isValidOp {
		return opts.Localizer.MustLocalizeError("kafka.acl.common.error.invalidResourceOperation",
			localize.NewEntry("ResourceType", opts.ResourceType),
			localize.NewEntry("Operation", opts.Operation),
			localize.NewEntry("ValidOperationList", cmdutil.StringSliceToListStringWithQuotes(validResourceOperations)),
		)
	}

	requestParams := getRequestParams(opts)

	newAclBinding := kafkainstanceclient.NewAclBinding(
		kafkainstanceclient.AclResourceType(requestParams.resourceType),
		requestParams.resourceName,
		kafkainstanceclient.AclPatternType(requestParams.patternType),
		aclcmdutil.FormatPrincipal(opts.Principal),
		kafkainstanceclient.AclOperation(requestParams.operation),
		kafkainstanceclient.AclPermissionType(requestParams.permission),
	)

	rows := aclcmdutil.MapACLsToTableRows([]kafkainstanceclient.AclBinding{*newAclBinding}, opts.Localizer)

	opts.Logger.Info(opts.Localizer.MustLocalizePlural("kafka.acl.grantPermissions.log.info.aclsPreview", len(rows)))
	opts.Logger.Info()

	dump.Table(opts.IO.Out, rows)
	opts.Logger.Info()

	if !opts.SkipConfirm {
		prompt := &survey.Confirm{
			Message: opts.Localizer.MustLocalize("kafka.acl.create.input.confirmCreateMessage"),
		}
		if err = survey.AskOne(prompt, &opts.SkipConfirm); err != nil {
			return err
		}

		if !opts.SkipConfirm {
			opts.Logger.Debug("User has chosen to not create ACL")
			return nil
		}
	}

	kafkaNameTmplEntry := localize.NewEntry("Name", kafkaInstance.GetName())

	opts.Logger.Info()
	spinnr := spinner.New(opts.IO.ErrOut, opts.Localizer)
	spinnr.SetLocalizedSuffix("kafka.acl.create.log.info.creatingACL", kafkaNameTmplEntry)
	spinnr.Start()

	req := adminAPI.AclsApi.CreateAcl(opts.Context)

	req = req.AclBinding(*newAclBinding)

	err = aclcmdutil.ExecuteACLRuleCreate(req, opts.Localizer, kafkaName)
	spinnr.Stop()
	if err != nil {
		return err
	}

	return nil
}

func getRequestParams(opts *aclcmdutil.CrudOptions) *requestParams {
	return &requestParams{
		resourceType: kafkainstanceclient.AclResourceType(aclcmdutil.GetMappedResourceTypeFilterValue(opts.ResourceType)),
		principal:    aclcmdutil.FormatPrincipal(opts.Principal),
		resourceName: aclcmdutil.GetResourceName(opts.ResourceName),
		patternType:  aclcmdutil.GetMappedPatternTypeValue(opts.PatternType),
		operation:    aclcmdutil.GetMappedOperationValue(opts.Operation),
		permission:   aclcmdutil.GetMappedPermissionTypeValue(opts.Permission),
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

	opts.PatternType = aclcmdutil.PatternTypeLITERAL
	if prefix {
		opts.PatternType = aclcmdutil.PatternTypePREFIX
	}

	if userID != "" {
		opts.Principal = userID
	} else if serviceAccount != "" {
		opts.Principal = serviceAccount
	}

	return nil
}
