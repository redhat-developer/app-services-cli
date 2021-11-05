package list

import (
	"context"

	"github.com/spf13/cobra"

	"github.com/redhat-developer/app-services-cli/internal/config"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/factory"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/kafka/acl/flagutil"
	"github.com/redhat-developer/app-services-cli/pkg/connection"
	"github.com/redhat-developer/app-services-cli/pkg/dump"
	"github.com/redhat-developer/app-services-cli/pkg/iostreams"
	"github.com/redhat-developer/app-services-cli/pkg/kafka/aclutil"
	"github.com/redhat-developer/app-services-cli/pkg/localize"
	"github.com/redhat-developer/app-services-cli/pkg/logging"
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

	page      int32
	size      int32
	kafkaID   string
	principal string
	output    string
}

// NewListACLCommand creates a new command to list Kafka ACL rules
func NewListACLCommand(f *factory.Factory) *cobra.Command {

	opts := &options{
		config:     f.Config,
		connection: f.Connection,
		logger:     f.Logger,
		io:         f.IOStreams,
		localizer:  f.Localizer,
		context:    f.Context,
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

			if opts.kafkaID != "" {
				return runList(opts)
			}

			cfg, err := opts.config.Load()
			if err != nil {
				return err
			}

			instanceID, ok := cfg.GetKafkaIdOk()
			if !ok {
				return opts.localizer.MustLocalizeError("kafka.acl.common.error.noKafkaSelected")
			}

			opts.kafkaID = instanceID

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

			return runList(opts)
		},
	}

	flags := flagutil.NewFlagSet(cmd, opts.localizer, opts.connection)

	flags.AddInstanceID(&opts.kafkaID)
	flags.AddOutput(&opts.output)
	flags.AddPage(&opts.page)
	flags.AddSize(&opts.size)
	flags.AddUser(&userID)
	flags.AddServiceAccount(&serviceAccount)
	flags.AddAllAccounts(&allAccounts)

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

	req = req.Page(float32(opts.page)).Size(float32(opts.size))
	req = req.Order("asc").OrderKey("principal")

	if opts.principal != "" {
		principalQuery := aclutil.FormatPrincipal(opts.principal)
		req = req.Principal(principalQuery)
	}

	permissionsData, httpRes, err := req.Execute()
	if httpRes != nil {
		defer httpRes.Body.Close()
	}

	if err = aclutil.ValidateAPIError(httpRes, opts.localizer, err, "list", kafkaInstance.GetName()); err != nil {
		return err
	}

	switch opts.output {
	case dump.EmptyFormat:
		opts.logger.Info("")
		permissions := permissionsData.GetItems()
		rows := aclutil.MapACLsToTableRows(permissions, opts.localizer)
		dump.Table(opts.io.Out, rows)
	default:
		return dump.Formatted(opts.io.Out, opts.output, permissionsData)
	}

	return nil
}
