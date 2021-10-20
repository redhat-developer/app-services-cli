package admin

import (
	"context"
	"fmt"

	"github.com/redhat-developer/app-services-cli/internal/config"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/factory"
	"github.com/redhat-developer/app-services-cli/pkg/connection"
	"github.com/redhat-developer/app-services-cli/pkg/icon"
	"github.com/redhat-developer/app-services-cli/pkg/iostreams"
	"github.com/redhat-developer/app-services-cli/pkg/kafka/aclutil"
	"github.com/redhat-developer/app-services-cli/pkg/localize"
	"github.com/redhat-developer/app-services-cli/pkg/logging"
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

	kafkaID   string
	principal string
}

// NewAdminACLCommand creates ACL rule to aloow user to add and delete ACL rules
func NewAdminACLCommand(f *factory.Factory) *cobra.Command {

	opts := &options{
		Config:     f.Config,
		Connection: f.Connection,
		Logger:     f.Logger,
		IO:         f.IOStreams,
		localizer:  f.Localizer,
		Context:    f.Context,
	}

	cmd := &cobra.Command{
		Use:     "admin",
		Short:   f.Localizer.MustLocalize("kafka.acl.admin.cmd.shortDescription"),
		Long:    f.Localizer.MustLocalize("kafka.acl.admin.cmd.longDescription"),
		Example: f.Localizer.MustLocalize("kafka.acl.admin.cmd.example"),
		Args:    cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {

			cfg, err := opts.Config.Load()
			if err != nil {
				return err
			}

			if !cfg.HasKafka() {
				return opts.localizer.MustLocalizeError("kafka.acl.common.error.noKafkaSelected")
			}

			opts.kafkaID = cfg.Services.Kafka.ClusterID

			// check if priincipal is provided
			if userID == "" && serviceAccount == "" && !allAccounts {
				return opts.localizer.MustLocalizeError("kafka.acl.grantPermissions.error.noPrincipalsSelected")
			}

			// user and service account can't be along with "--all-accounts" flag
			if allAccounts && (serviceAccount != "" || userID != "") {
				return opts.localizer.MustLocalizeError("kafka.acl.grantPermissions.allPrinciapls.error.notAllowed")
			}

			return runAdmin(opts)
		},
	}

	cmd.Flags().StringVar(&userID, "user", "", opts.localizer.MustLocalize("kafka.acl.common.flag.user.description"))
	cmd.Flags().StringVar(&serviceAccount, "service-account", "", opts.localizer.MustLocalize("kafka.acl.common.flag.serviceAccount.description"))
	cmd.Flags().BoolVar(&allAccounts, "all-accounts", false, opts.localizer.MustLocalize("kafka.acl.common.flag.allAccounts.description"))

	return cmd
}

func runAdmin(opts *options) (err error) {

	conn, err := opts.Connection(connection.DefaultConfigRequireMasAuth)
	if err != nil {
		return err
	}

	api, kafkaInstance, err := conn.API().KafkaAdmin(opts.kafkaID)
	if err != nil {
		return err
	}

	kafkaName := kafkaInstance.GetName()

	if userID != "" {
		opts.principal = userID
	}

	if serviceAccount != "" {
		opts.principal = serviceAccount
	}

	if allAccounts {
		opts.principal = aclutil.Wildcard
	}

	req := api.AclsApi.CreateAcl(opts.Context)

	aclBindClusterAlter := *kafkainstanceclient.NewAclBinding(
		kafkainstanceclient.ACLRESOURCETYPE_CLUSTER,
		aclutil.KafkaCluster,
		kafkainstanceclient.ACLPATTERNTYPE_LITERAL,
		buildPrincipal(opts.principal),
		kafkainstanceclient.ACLOPERATION_ALTER,
		kafkainstanceclient.ACLPERMISSIONTYPE_ALLOW,
	)

	req = req.AclBinding(aclBindClusterAlter)

	err = aclutil.ExecuteACLRuleCreate(req, opts.localizer, kafkaName)
	if err != nil {
		return err
	}

	opts.Logger.Info(icon.SuccessPrefix(), opts.localizer.MustLocalize("kafka.acl.admin.log.info.successful", localize.NewEntry("Account", opts.principal), localize.NewEntry("InstanceName", kafkaName)))

	return nil
}

func buildPrincipal(user string) string {
	return fmt.Sprintf("User:%s", user)
}
