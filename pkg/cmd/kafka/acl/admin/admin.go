package admin

import (
	"context"

	"github.com/AlecAivazis/survey/v2"
	"github.com/redhat-developer/app-services-cli/internal/config"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/factory"
	"github.com/redhat-developer/app-services-cli/pkg/connection"
	"github.com/redhat-developer/app-services-cli/pkg/dump"
	"github.com/redhat-developer/app-services-cli/pkg/icon"
	"github.com/redhat-developer/app-services-cli/pkg/iostreams"
	"github.com/redhat-developer/app-services-cli/pkg/kafka/aclutil"
	"github.com/redhat-developer/app-services-cli/pkg/localize"
	"github.com/redhat-developer/app-services-cli/pkg/logging"
	"github.com/spf13/cobra"

	flagset "github.com/redhat-developer/app-services-cli/pkg/cmd/kafka/acl/flagutil"

	kafkainstanceclient "github.com/redhat-developer/app-services-sdk-go/kafkainstance/apiv1internal/client"
)

var (
	serviceAccount string
	userID         string
	allAccounts    bool
)

type options struct {
	config     config.IConfig
	connection factory.ConnectionFunc
	logger     logging.Logger
	io         *iostreams.IOStreams
	localizer  localize.Localizer
	context    context.Context

	kafkaID     string
	principal   string
	skipConfirm bool
}

// NewAdminACLCommand creates ACL rule to aloow user to add and delete ACL rules
func NewAdminACLCommand(f *factory.Factory) *cobra.Command {

	opts := &options{
		config:     f.Config,
		connection: f.Connection,
		logger:     f.Logger,
		io:         f.IOStreams,
		localizer:  f.Localizer,
		context:    f.Context,
	}

	cmd := &cobra.Command{
		Use:     "grant-admin",
		Short:   f.Localizer.MustLocalize("kafka.acl.grantAdmin.cmd.shortDescription"),
		Long:    f.Localizer.MustLocalize("kafka.acl.grantAdmin.cmd.longDescription"),
		Example: f.Localizer.MustLocalize("kafka.acl.grantAdmin.cmd.example"),
		Args:    cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {

			cfg, err := opts.config.Load()
			if err != nil {
				return err
			}

			instanceID, ok := cfg.GetKafkaIdOk()

			if !ok {
				return opts.localizer.MustLocalizeError("kafka.acl.common.error.noKafkaSelected")
			}

			opts.kafkaID = instanceID

			// check if principal is provided
			if userID == "" && serviceAccount == "" && !allAccounts {
				return opts.localizer.MustLocalizeError("kafka.acl.common.error.noPrincipalsSelected")
			}

			// user and service account can't be along with "--all-accounts" flag
			if allAccounts && (serviceAccount != "" || userID != "") {
				return opts.localizer.MustLocalizeError("kafka.acl.common.error.allAccountsCannotBeUsedWithUserFlag")
			}

			// user and service account should not allow wildcard
			if userID == aclutil.Wildcard || serviceAccount == aclutil.Wildcard {
				return opts.localizer.MustLocalizeError("kafka.acl.common.error.useAllAccountsFlag")
			}

			if userID != "" {
				opts.principal = userID
			}

			if serviceAccount != "" {
				opts.principal = serviceAccount
			}

			if allAccounts {
				opts.principal = aclutil.Wildcard
			}

			return runAdmin(opts)
		},
	}

	fs := flagset.NewFlagSet(cmd, opts.localizer, opts.connection)

	fs.AddUser(&userID)
	fs.AddServiceAccount(&serviceAccount)
	fs.AddAllAccounts(&allAccounts)
	fs.AddYes(&opts.skipConfirm)

	return cmd
}

func runAdmin(opts *options) (err error) {

	conn, err := opts.connection(connection.DefaultConfigRequireMasAuth)
	if err != nil {
		return err
	}

	api, kafkaInstance, err := conn.API().KafkaAdmin(opts.kafkaID)
	if err != nil {
		return err
	}

	kafkaName := kafkaInstance.GetName()

	req := api.AclsApi.CreateAcl(opts.context)

	aclBindClusterAlter := kafkainstanceclient.NewAclBinding(
		kafkainstanceclient.ACLRESOURCETYPE_CLUSTER,
		aclutil.KafkaCluster,
		kafkainstanceclient.ACLPATTERNTYPE_LITERAL,
		aclutil.FormatPrincipal(opts.principal),
		kafkainstanceclient.ACLOPERATION_ALTER,
		kafkainstanceclient.ACLPERMISSIONTYPE_ALLOW,
	)

	opts.logger.Info(opts.localizer.MustLocalize("kafka.acl.grantPermissions.log.info.aclsPreview"))
	opts.logger.Info()

	rows := aclutil.MapACLsToTableRows(&[]kafkainstanceclient.AclBinding{*aclBindClusterAlter}, opts.localizer)
	dump.Table(opts.io.Out, rows)
	opts.logger.Info()

	if !opts.skipConfirm {
		var confirmGrant bool
		promptConfirmGrant := &survey.Confirm{
			Message: opts.localizer.MustLocalize("kafka.acl.common.input.confirmGrant.message"),
		}

		err = survey.AskOne(promptConfirmGrant, &confirmGrant)
		if err != nil {
			return err
		}

		if !confirmGrant {
			opts.logger.Debug(opts.localizer.MustLocalize("kafka.acl.grantAdmin.log.debug.grantNotConfirmed"))
			return nil
		}
	}

	req = req.AclBinding(*aclBindClusterAlter)

	err = aclutil.ExecuteACLRuleCreate(req, opts.localizer, kafkaName)
	if err != nil {
		return err
	}

	opts.logger.Info(icon.SuccessPrefix(), opts.localizer.MustLocalize("kafka.acl.grantAdmin.log.info.successful", localize.NewEntry("Account", opts.principal), localize.NewEntry("InstanceName", kafkaName)))

	return nil
}
