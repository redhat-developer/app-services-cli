package list

import (
	"context"

	"github.com/redhat-developer/app-services-cli/pkg/cmd/kafka/acl/aclcmdutil"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/kafka/acl/flagutil"
	kafkacmdutil "github.com/redhat-developer/app-services-cli/pkg/shared/kafkautil"
	"github.com/redhat-developer/app-services-cli/pkg/shared/profileutil"

	"github.com/redhat-developer/app-services-cli/pkg/core/ioutil/dump"
	"github.com/redhat-developer/app-services-cli/pkg/core/ioutil/iostreams"
	"github.com/redhat-developer/app-services-cli/pkg/core/localize"
	"github.com/redhat-developer/app-services-cli/pkg/core/logging"
	"github.com/redhat-developer/app-services-cli/pkg/core/servicecontext"
	"github.com/redhat-developer/app-services-cli/pkg/shared/connection"
	"github.com/redhat-developer/app-services-cli/pkg/shared/factory"
	"github.com/spf13/cobra"
)

var (
	serviceAccount string
	userID         string
	allAccounts    bool
)

type options struct {
	connection     factory.ConnectionFunc
	logger         logging.Logger
	io             *iostreams.IOStreams
	localizer      localize.Localizer
	context        context.Context
	serviceContext servicecontext.IContext

	page      int32
	size      int32
	kafkaID   string
	principal string

	topic   string
	group   string
	cluster bool

	output string
}

// nolint:funlen
// NewListACLCommand creates a new command to list Kafka ACL rules
func NewListACLCommand(f *factory.Factory) *cobra.Command {

	opts := &options{
		connection:     f.Connection,
		logger:         f.Logger,
		io:             f.IOStreams,
		localizer:      f.Localizer,
		context:        f.Context,
		serviceContext: f.ServiceContext,
	}

	cmd := &cobra.Command{
		Use:     "list",
		Short:   f.Localizer.MustLocalize("kafka.acl.list.cmd.shortDescription"),
		Long:    f.Localizer.MustLocalize("kafka.acl.list.cmd.longDescription"),
		Example: f.Localizer.MustLocalize("kafka.acl.list.cmd.example"),
		Args:    cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {

			if opts.page < 1 {
				return opts.localizer.MustLocalizeError("kafka.common.validation.page.error.invalid.minValue", localize.NewEntry("Page", opts.page))
			}

			if opts.size < 1 {
				return opts.localizer.MustLocalizeError("kafka.common.validation.page.error.invalid.minValue", localize.NewEntry("Size", opts.size))
			}

			if opts.kafkaID == "" {

				svcContext, err := opts.serviceContext.Load()
				if err != nil {
					return err
				}

				profileHandler := &profileutil.ContextHandler{
					Context:   svcContext,
					Localizer: opts.localizer,
				}

				conn, err := opts.connection(connection.DefaultConfigRequireMasAuth)
				if err != nil {
					return err
				}

				kafkaInstance, err := profileHandler.GetCurrentKafkaInstance(conn.API().KafkaMgmt())
				if err != nil {
					return err
				}

				opts.kafkaID = kafkaInstance.GetId()
			}

			// user and service account can't be along with "--all-accounts" flag
			if allAccounts && (serviceAccount != "" || userID != "") {
				return opts.localizer.MustLocalizeError("kafka.acl.common.error.allAccountsCannotBeUsedWithUserFlag")
			}

			// user and service account should not allow wildcard
			if userID == aclcmdutil.Wildcard || serviceAccount == aclcmdutil.Wildcard || userID == aclcmdutil.AllAlias || serviceAccount == aclcmdutil.AllAlias {
				return opts.localizer.MustLocalizeError("kafka.acl.common.error.useAllAccountsFlag")
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

			return runList(opts)
		},
	}

	flags := flagutil.NewFlagSet(cmd, f)

	flags.AddInstanceID(&opts.kafkaID)
	flags.AddOutput(&opts.output)
	flags.AddPage(&opts.page)
	flags.AddSize(&opts.size)
	flags.AddUser(&userID)
	flags.AddServiceAccount(&serviceAccount)
	flags.AddAllAccounts(&allAccounts)
	flags.BoolVar(&opts.cluster, "cluster", false, opts.localizer.MustLocalize("kafka.acl.list.flag.cluster.description"))
	flags.StringVar(&opts.topic, "topic", "", opts.localizer.MustLocalize("kafka.acl.list.flag.topic.description"))
	flags.StringVar(&opts.group, "group", "", opts.localizer.MustLocalize("kafka.acl.list.flag.group.description"))

	_ = cmd.RegisterFlagCompletionFunc("topic", func(cmd *cobra.Command, _ []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return kafkacmdutil.FilterValidTopicNameArgs(f, toComplete)
	})

	_ = cmd.RegisterFlagCompletionFunc("group", func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return kafkacmdutil.FilterValidConsumerGroupIDs(f, toComplete)
	})

	return cmd
}

func runList(opts *options) (err error) {
	conn, err := opts.connection(connection.DefaultConfigRequireMasAuth)
	if err != nil {
		return err
	}

	api, kafkaInstance, err := conn.API().KafkaAdmin(opts.kafkaID)
	if err != nil {
		return err
	}

	req := api.AclsApi.GetAcls(opts.context)

	req = req.Page(opts.page).Size(opts.size)
	req = req.Order("asc").OrderKey("principal")

	if opts.principal != "" {
		principalQuery := aclcmdutil.FormatPrincipal(opts.principal)
		req = req.Principal(principalQuery)
	}

	var selectedResourceTypeCount int
	var resourceType string
	var resourceName string

	if opts.topic != "" {
		selectedResourceTypeCount++
		resourceType = aclcmdutil.ResourceTypeTOPIC
		resourceName = opts.topic
	}

	if opts.group != "" {
		selectedResourceTypeCount++
		resourceType = aclcmdutil.ResourceTypeGROUP
		resourceName = opts.group
	}

	if opts.cluster {
		selectedResourceTypeCount++
		resourceType = aclcmdutil.ResourceTypeCLUSTER
		resourceName = aclcmdutil.KafkaCluster
	}

	if selectedResourceTypeCount > 1 {
		return opts.localizer.MustLocalizeError("kafka.acl.list.error.oneResourceTypeAllowed", flagutil.ResourceTypeFlagEntries...)
	}

	if resourceType != "" {
		req = req.ResourceType(aclcmdutil.GetMappedResourceTypeFilterValue(resourceType))
	}

	if resourceName != "" {
		req = req.ResourceName(aclcmdutil.GetResourceName(resourceName))
	}

	permissionsData, httpRes, err := req.Execute()
	if httpRes != nil {
		defer httpRes.Body.Close()
	}

	if err = aclcmdutil.ValidateAPIError(httpRes, opts.localizer, err, "list", kafkaInstance.GetName()); err != nil {
		return err
	}

	if permissionsData.GetTotal() == 0 && opts.output == "" {
		opts.logger.Info(opts.localizer.MustLocalize("kafka.acl.list.log.info.noACLs", localize.NewEntry("InstanceName", kafkaInstance.GetName())))

		return nil
	}

	switch opts.output {
	case dump.EmptyFormat:
		opts.logger.Info("")
		permissions := permissionsData.GetItems()
		rows := aclcmdutil.MapACLsToTableRows(permissions, opts.localizer)
		dump.Table(opts.io.Out, rows)
	default:
		return dump.Formatted(opts.io.Out, opts.output, permissionsData)
	}

	return nil
}
