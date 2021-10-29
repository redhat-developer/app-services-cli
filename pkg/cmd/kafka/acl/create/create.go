package create

import (
	"github.com/AlecAivazis/survey/v2"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/factory"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/flag"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/kafka/acl/flagutil"
	"github.com/redhat-developer/app-services-cli/pkg/cmdutil"
	"github.com/redhat-developer/app-services-cli/pkg/connection"
	"github.com/redhat-developer/app-services-cli/pkg/dump"
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
	resourceType kafkainstanceclient.AclResourceType
	patternType  kafkainstanceclient.AclPatternType
	operation    kafkainstanceclient.AclOperation
	permission   kafkainstanceclient.AclPermissionType
}

// NewCreateCommand creates a new command to add Kafka ACLs
func NewCreateCommand(f *factory.Factory) *cobra.Command {
	opts := &aclutil.CrudOptions{
		Config:     f.Config,
		Connection: f.Connection,
		Logger:     f.Logger,
		IO:         f.IOStreams,
		Localizer:  f.Localizer,
		Context:    f.Context,
	}

	cmd := &cobra.Command{
		Use:     "create",
		Short:   f.Localizer.MustLocalize("kafka.acl.create.cmd.shortDescription"),
		Long:    f.Localizer.MustLocalize("kafka.acl.create.cmd.longDescription"),
		Example: f.Localizer.MustLocalize("kafka.acl.create.cmd.example"),
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

			return runAdd(opts.InstanceID, opts)
		},
	}

	flags := flagutil.NewFlagSet(cmd, opts.Localizer, opts.Connection)

	_ = flags.AddPermissionCreate(&opts.Permission).Required()
	_ = flags.AddOperationCreate(&opts.Operation).Required()

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
func runAdd(instanceID string, opts *aclutil.CrudOptions) error {
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

	if isValidOp, validResourceOperations := aclutil.IsValidResourceOperation(opts.ResourceType, opts.Operation, resourceOperations); !isValidOp {
		return opts.Localizer.MustLocalizeError("kafka.acl.common.error.invalidResourceOperation",
			localize.NewEntry("ResourceType", opts.ResourceType),
			localize.NewEntry("Operation", opts.Operation),
			localize.NewEntry("ValidOperationList", cmdutil.StringSliceToListStringWithQuotes(validResourceOperations)),
		)
	}

	opts.PatternType = aclutil.PatternTypeLITERAL
	if prefix {
		opts.PatternType = aclutil.PatternTypePREFIX
	}

	requestParams := getRequestParams(opts)

	newAclBinding := kafkainstanceclient.NewAclBinding(
		kafkainstanceclient.AclResourceType(requestParams.resourceType),
		requestParams.resourceName,
		kafkainstanceclient.AclPatternType(requestParams.patternType),
		aclutil.FormatPrincipal(opts.Principal),
		kafkainstanceclient.AclOperation(requestParams.operation),
		kafkainstanceclient.AclPermissionType(requestParams.permission),
	)

	opts.Logger.Info(opts.Localizer.MustLocalize("kafka.acl.grantPermissions.log.info.aclsPreview"))
	opts.Logger.Info()

	rows := aclutil.MapACLsToTableRows([]kafkainstanceclient.AclBinding{*newAclBinding}, opts.Localizer)
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

	if err = aclutil.ExecuteACLRuleCreate(req, opts.Localizer, kafkaName); err != nil {
		return err
	}

	spinnr.Stop()

	return nil
}

func getRequestParams(opts *aclutil.CrudOptions) *requestParams {
	return &requestParams{
		resourceType: kafkainstanceclient.AclResourceType(aclutil.GetMappedResourceTypeFilterValue(opts.ResourceType)),
		principal:    aclutil.FormatPrincipal(opts.Principal),
		resourceName: opts.ResourceName,
		patternType:  aclutil.GetMappedPatternTypeValue(opts.PatternType),
		operation:    aclutil.GetMappedOperationValue(opts.Operation),
		permission:   aclutil.GetMappedPermissionTypeValue(opts.Permission),
	}
}

func validateAndSetOpts(opts *aclutil.CrudOptions) error {

	// user and service account should not be provided together
	if userID != "" && serviceAccount != "" {
		return opts.Localizer.MustLocalizeError("kafka.acl.common.error.bothPrincipalsSelected")
	}

	if userID == aclutil.Wildcard || serviceAccount == aclutil.Wildcard {
		return opts.Localizer.MustLocalizeError("kafka.acl.common.error.useAllAccountsFlag")
	}

	if allAccounts {
		if userID != "" || serviceAccount != "" {
			return opts.Localizer.MustLocalizeError("kafka.acl.common.error.allAccountsCannotBeUsedWithUserFlag")
		}
		opts.Principal = aclutil.Wildcard
	}

	// check if priincipal is provided
	if !allAccounts && (userID == "" && serviceAccount == "") {
		return opts.Localizer.MustLocalizeError("kafka.acl.common.error.noPrincipalsSelected")
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
