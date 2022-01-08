package grant

import (
	"context"

	"github.com/redhat-developer/app-services-cli/pkg/cmd/kafka/acl/aclcmdutil"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/kafka/acl/flagutil"

	"github.com/AlecAivazis/survey/v2"
	"github.com/redhat-developer/app-services-cli/pkg/core/cmdutil/factory"
	"github.com/redhat-developer/app-services-cli/pkg/core/config"
	"github.com/redhat-developer/app-services-cli/pkg/core/connection"
	"github.com/redhat-developer/app-services-cli/pkg/core/ioutil/dump"
	"github.com/redhat-developer/app-services-cli/pkg/core/ioutil/icon"
	"github.com/redhat-developer/app-services-cli/pkg/core/ioutil/iostreams"
	"github.com/redhat-developer/app-services-cli/pkg/core/localize"
	"github.com/redhat-developer/app-services-cli/pkg/core/logging"
	"github.com/spf13/cobra"

	kafkainstanceclient "github.com/redhat-developer/app-services-sdk-go/kafkainstance/apiv1internal/client"
)

var (
	serviceAccount string
	userID         string
	allAccounts    bool
)

type options struct {
	Config     config.IConfig
	Connection factory.ConnectionFunc
	Logger     logging.Logger
	IO         *iostreams.IOStreams
	localizer  localize.Localizer
	Context    context.Context

	kafkaID     string
	topic       string
	principal   string
	group       string
	producer    bool
	consumer    bool
	topicPrefix string
	groupPrefix string
	force       bool
}

// NewGrantPermissionsACLCommand creates a series of ACL rules
func NewGrantPermissionsACLCommand(f *factory.Factory) *cobra.Command {

	opts := &options{
		Config:     f.Config,
		Connection: f.Connection,
		Logger:     f.Logger,
		IO:         f.IOStreams,
		localizer:  f.Localizer,
		Context:    f.Context,
	}

	cmd := &cobra.Command{
		Use:     "grant-access",
		Short:   f.Localizer.MustLocalize("kafka.acl.grantPermissions.cmd.shortDescription"),
		Long:    f.Localizer.MustLocalize("kafka.acl.grantPermissions.cmd.longDescription"),
		Example: f.Localizer.MustLocalize("kafka.acl.grantPermissions.cmd.example"),
		Args:    cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {

			if opts.kafkaID != "" {
				return runGrantPermissions(opts)
			}

			cfg, err := opts.Config.Load()
			if err != nil {
				return err
			}

			instanceID, ok := cfg.GetKafkaIdOk()
			if !ok {
				return opts.localizer.MustLocalizeError("kafka.acl.common.error.noKafkaSelected")
			}

			opts.kafkaID = instanceID

			if err = validateFlagInputCombination(opts); err != nil {
				return err
			}

			return runGrantPermissions(opts)
		},
	}

	flags := flagutil.NewFlagSet(cmd, f)
	flags.AddInstanceID(&opts.kafkaID)
	flags.AddYes(&opts.force)

	flags.StringVar(&userID, "user", "", opts.localizer.MustLocalize("kafka.acl.common.flag.user.description"))
	flags.StringVar(&serviceAccount, "service-account", "", opts.localizer.MustLocalize("kafka.acl.common.flag.serviceAccount.description"))
	flags.StringVar(&opts.topic, "topic", "", opts.localizer.MustLocalize("kafka.acl.grantPermissions.common.flag.topic.description"))
	flags.StringVar(&opts.group, "group", "", opts.localizer.MustLocalize("kafka.acl.grantPermissions.common.flag.group.description"))
	flags.BoolVar(&opts.consumer, "consumer", false, opts.localizer.MustLocalize("kafka.acl.grantPermissions.flag.consumer.description"))
	flags.BoolVar(&opts.producer, "producer", false, opts.localizer.MustLocalize("kafka.acl.grantPermissions.flag.producer.description"))
	flags.StringVar(&opts.topicPrefix, "topic-prefix", "", opts.localizer.MustLocalize("kafka.acl.grantPermissions.common.flag.topicPrefix.description"))
	flags.StringVar(&opts.groupPrefix, "group-prefix", "", opts.localizer.MustLocalize("kafka.acl.grantPermissions.common.flag.groupPrefix.description"))
	flags.BoolVar(&allAccounts, "all-accounts", false, opts.localizer.MustLocalize("kafka.acl.common.flag.allAccounts.description"))

	return cmd
}

// nolint:funlen
func runGrantPermissions(opts *options) (err error) {

	conn, err := opts.Connection(connection.DefaultConfigRequireMasAuth)
	if err != nil {
		return err
	}

	api, kafkaInstance, err := conn.API().KafkaAdmin(opts.kafkaID)
	if err != nil {
		return err
	}

	kafkaName := kafkaInstance.GetName()

	var topicNameArg string
	var groupIdArg string
	var topicPatternArg = kafkainstanceclient.ACLPATTERNTYPE_LITERAL
	var groupPatternArg = kafkainstanceclient.ACLPATTERNTYPE_LITERAL

	if opts.topic != "" {
		topicNameArg = aclcmdutil.GetResourceName(opts.topic)
	}

	if opts.topicPrefix != "" {
		topicNameArg = opts.topicPrefix
		topicPatternArg = kafkainstanceclient.ACLPATTERNTYPE_PREFIXED
	}

	if opts.group != "" {
		groupIdArg = aclcmdutil.GetResourceName(opts.group)
	}

	if opts.groupPrefix != "" {
		groupIdArg = opts.groupPrefix
		groupPatternArg = kafkainstanceclient.ACLPATTERNTYPE_PREFIXED
	}

	if userID != "" {
		opts.principal = userID
	}

	if serviceAccount != "" {
		opts.principal = serviceAccount
	}

	if allAccounts {
		opts.principal = aclcmdutil.Wildcard
	}

	var aclBindRequests []kafkainstanceclient.ApiCreateAclRequest
	var aclBindingList []kafkainstanceclient.AclBinding

	userArg := aclcmdutil.FormatPrincipal(opts.principal)

	req := api.AclsApi.CreateAcl(opts.Context)

	aclBindTopicDescribe := kafkainstanceclient.NewAclBinding(
		kafkainstanceclient.ACLRESOURCETYPE_TOPIC,
		topicNameArg,
		topicPatternArg,
		userArg,
		kafkainstanceclient.ACLOPERATION_DESCRIBE,
		kafkainstanceclient.ACLPERMISSIONTYPE_ALLOW,
	)

	aclBindingList = append(aclBindingList, *aclBindTopicDescribe)

	aclBindRequests = append(aclBindRequests, req.AclBinding(*aclBindTopicDescribe))

	if opts.consumer {

		aclBindTopicRead := kafkainstanceclient.NewAclBinding(
			kafkainstanceclient.ACLRESOURCETYPE_TOPIC,
			topicNameArg,
			topicPatternArg,
			userArg,
			kafkainstanceclient.ACLOPERATION_READ,
			kafkainstanceclient.ACLPERMISSIONTYPE_ALLOW,
		)

		aclBindingList = append(aclBindingList, *aclBindTopicRead)

		aclBindRequests = append(aclBindRequests, req.AclBinding(*aclBindTopicRead))

		aclBindGroupRead := kafkainstanceclient.NewAclBinding(
			kafkainstanceclient.ACLRESOURCETYPE_GROUP,
			groupIdArg,
			groupPatternArg,
			userArg,
			kafkainstanceclient.ACLOPERATION_READ,
			kafkainstanceclient.ACLPERMISSIONTYPE_ALLOW,
		)

		aclBindingList = append(aclBindingList, *aclBindGroupRead)

		aclBindRequests = append(aclBindRequests, req.AclBinding(*aclBindGroupRead))

	}

	if opts.producer {

		aclBindTopicWrite := kafkainstanceclient.NewAclBinding(
			kafkainstanceclient.ACLRESOURCETYPE_TOPIC,
			topicNameArg,
			topicPatternArg,
			userArg,
			kafkainstanceclient.ACLOPERATION_WRITE,
			kafkainstanceclient.ACLPERMISSIONTYPE_ALLOW,
		)

		aclBindingList = append(aclBindingList, *aclBindTopicWrite)

		aclBindRequests = append(aclBindRequests, req.AclBinding(*aclBindTopicWrite))

		aclBindTopicCreate := kafkainstanceclient.NewAclBinding(
			kafkainstanceclient.ACLRESOURCETYPE_TOPIC,
			topicNameArg,
			topicPatternArg,
			userArg,
			kafkainstanceclient.ACLOPERATION_CREATE,
			kafkainstanceclient.ACLPERMISSIONTYPE_ALLOW,
		)

		aclBindingList = append(aclBindingList, *aclBindTopicCreate)

		aclBindRequests = append(aclBindRequests, req.AclBinding(*aclBindTopicCreate))

		// Add ACLs for transactional IDs
		aclBindTransactionIDWrite := kafkainstanceclient.NewAclBinding(
			kafkainstanceclient.ACLRESOURCETYPE_TRANSACTIONAL_ID,
			aclcmdutil.Wildcard,
			kafkainstanceclient.ACLPATTERNTYPE_LITERAL,
			userArg,
			kafkainstanceclient.ACLOPERATION_WRITE,
			kafkainstanceclient.ACLPERMISSIONTYPE_ALLOW,
		)

		aclBindingList = append(aclBindingList, *aclBindTransactionIDWrite)

		aclBindRequests = append(aclBindRequests, req.AclBinding(*aclBindTransactionIDWrite))

		aclBindTransactionIDDescribe := kafkainstanceclient.NewAclBinding(
			kafkainstanceclient.ACLRESOURCETYPE_TRANSACTIONAL_ID,
			aclcmdutil.Wildcard,
			kafkainstanceclient.ACLPATTERNTYPE_LITERAL,
			userArg,
			kafkainstanceclient.ACLOPERATION_DESCRIBE,
			kafkainstanceclient.ACLPERMISSIONTYPE_ALLOW,
		)

		aclBindingList = append(aclBindingList, *aclBindTransactionIDDescribe)

		aclBindRequests = append(aclBindRequests, req.AclBinding(*aclBindTransactionIDDescribe))
	}

	opts.Logger.Info(opts.localizer.MustLocalize("kafka.acl.grantPermissions.log.info.aclsPreview"))
	opts.Logger.Info()

	rows := aclcmdutil.MapACLsToTableRows(aclBindingList, opts.localizer)
	dump.Table(opts.IO.Out, rows)
	opts.Logger.Info()

	if !opts.force {
		var confirmGrant bool
		promptConfirmGrant := &survey.Confirm{
			Message: opts.localizer.MustLocalize("kafka.acl.common.input.confirmGrant.message"),
		}

		err = survey.AskOne(promptConfirmGrant, &confirmGrant)
		if err != nil {
			return err
		}

		if !confirmGrant {
			opts.Logger.Debug(opts.localizer.MustLocalize("kafka.acl.grantPermissions.log.debug.grantNotConfirmed"))
			return nil
		}
	}

	// Execute ACL rule creations
	for _, req := range aclBindRequests {
		if err = aclcmdutil.ExecuteACLRuleCreate(req, opts.localizer, kafkaName); err != nil {
			return err
		}
	}

	opts.Logger.Info(icon.SuccessPrefix(), opts.localizer.MustLocalize("kafka.acl.grantPermissions.log.info.aclsCreated", localize.NewEntry("InstanceName", kafkaName)))

	return nil

}

// validateFlagInputCombination checks if appropriate flags are provided for specified operation
func validateFlagInputCombination(opts *options) error {

	var errorCollection []error
	// check if any operation is specified
	if !opts.consumer && !opts.producer {
		errorCollection = append(errorCollection, opts.localizer.MustLocalizeError("kafka.acl.common.error.noOperationSpecified"))
	}

	// check if principal is provided
	if userID == "" && serviceAccount == "" && !allAccounts {
		errorCollection = append(errorCollection, opts.localizer.MustLocalizeError("kafka.acl.common.error.noPrincipalsSelected"))
	}

	// user and service account should not be provided together
	if userID != "" && serviceAccount != "" {
		errorCollection = append(errorCollection, opts.localizer.MustLocalizeError("kafka.acl.common.error.bothPrincipalsSelected"))
	}

	// user and service account can't be along with "--all-accounts" flag
	if allAccounts && (serviceAccount != "" || userID != "") {
		errorCollection = append(errorCollection, opts.localizer.MustLocalizeError("kafka.acl.common.error.allAccountsCannotBeUsedWithUserFlag"))
	}

	// checks if group resource name is provided when operation is not consumer
	if !opts.consumer && (opts.group != "" || opts.groupPrefix != "") {
		errorCollection = append(errorCollection, opts.localizer.MustLocalizeError("kafka.acl.grantPermissions.group.error.notAllowed"))

	}

	// checks if topic flag is provided
	if opts.topic == "" && opts.topicPrefix == "" {
		errorCollection = append(errorCollection, opts.localizer.MustLocalizeError("kafka.acl.grantPermissions.topic.error.required"))
	}

	// checks if group resource name is provided for consumer operation
	if opts.consumer && opts.group == "" && opts.groupPrefix == "" {
		errorCollection = append(errorCollection, opts.localizer.MustLocalizeError("kafka.acl.grantPermissions.group.error.required"))
	}

	if (!opts.consumer && !opts.producer) && (opts.group == "" && opts.groupPrefix == "") {
		errorCollection = append(errorCollection, opts.localizer.MustLocalizeError("kafka.acl.grantPermissions.group.error.required"))
	}

	// checks if "--topic" and "--topic-prefix" are provided together
	if opts.topicPrefix != "" && opts.topic != "" {
		topicErr := opts.localizer.MustLocalizeError("kafka.acl.grantPermissions.prefix.error.notAllowed",
			localize.NewEntry("Resource", "topic"),
		)
		errorCollection = append(errorCollection, topicErr)
	}

	// checks if "--group" and "--group-prefix" are provided together
	if opts.groupPrefix != "" && opts.group != "" {
		groupErr := opts.localizer.MustLocalizeError("kafka.acl.grantPermissions.prefix.error.notAllowed",
			localize.NewEntry("Resource", "group"),
		)
		errorCollection = append(errorCollection, groupErr)
	}

	// user and service account should not allow wildcard
	if userID == aclcmdutil.Wildcard || serviceAccount == aclcmdutil.Wildcard || userID == aclcmdutil.AllAlias || serviceAccount == aclcmdutil.AllAlias {
		errorCollection = append(errorCollection, opts.localizer.MustLocalizeError("kafka.acl.common.error.useAllAccountsFlag"))
	}

	if len(errorCollection) > 0 {
		return aclcmdutil.BuildInstructions(errorCollection)
	}

	return nil
}
