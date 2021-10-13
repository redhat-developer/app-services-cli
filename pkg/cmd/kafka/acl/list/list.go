package list

import (
	"context"
	"net/http"

	"github.com/redhat-developer/app-services-cli/internal/build"
	"github.com/redhat-developer/app-services-cli/internal/config"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/factory"
	"github.com/redhat-developer/app-services-cli/pkg/cmdutil"
	flagutil "github.com/redhat-developer/app-services-cli/pkg/cmdutil/flags"
	"github.com/redhat-developer/app-services-cli/pkg/connection"
	"github.com/redhat-developer/app-services-cli/pkg/dump"
	"github.com/redhat-developer/app-services-cli/pkg/iostreams"
	"github.com/redhat-developer/app-services-cli/pkg/kafka/acl"
	"github.com/redhat-developer/app-services-cli/pkg/localize"
	"github.com/redhat-developer/app-services-cli/pkg/logging"
	"github.com/spf13/cobra"
)

type options struct {
	Config     config.IConfig
	Connection factory.ConnectionFunc
	Logger     logging.Logger
	IO         *iostreams.IOStreams
	localizer  localize.Localizer
	Context    context.Context

	page    int32
	size    int32
	kafkaID string
	output  string
}

// NewListACLCommand creates a new command to list Kafka ACL rules
func NewListACLCommand(f *factory.Factory) *cobra.Command {

	opts := &options{
		Config:     f.Config,
		Connection: f.Connection,
		Logger:     f.Logger,
		IO:         f.IOStreams,
		localizer:  f.Localizer,
		Context:    f.Context,
	}

	cmd := &cobra.Command{
		Use:     "list",
		Short:   f.Localizer.MustLocalize("kafka.acl.list.cmd.shortDescription"),
		Long:    f.Localizer.MustLocalize("kafka.acl.list.cmd.longDescription"),
		Example: f.Localizer.MustLocalize("kafka.acl.list.cmd.example"),
		Args:    cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {

			if opts.kafkaID != "" {
				return runList(opts)
			}

			cfg, err := opts.Config.Load()
			if err != nil {
				return err
			}

			if !cfg.HasKafka() {
				return opts.localizer.MustLocalizeError("kafka.acl.common.error.noKafkaSelected")
			}

			opts.kafkaID = cfg.Services.Kafka.ClusterID

			return runList(opts)
		},
	}

	cmd.Flags().Int32Var(&opts.page, "page", cmdutil.ConvertPageValueToInt32(build.DefaultPageNumber), opts.localizer.MustLocalize("kafka.acl.list.flag.page.description"))
	cmd.Flags().Int32Var(&opts.size, "size", cmdutil.ConvertSizeValueToInt32(build.DefaultPageSize), opts.localizer.MustLocalize("kafka.acl.list.flag.size.description"))
	cmd.Flags().StringVarP(&opts.output, "output", "o", dump.EmptyFormat, opts.localizer.MustLocalize("kafka.acl.list.flag.output.description"))
	cmd.Flags().StringVar(&opts.kafkaID, "instance-id", "", opts.localizer.MustLocalize("kafka.acl.common.flag.instance.id"))

	flagutil.EnableOutputFlagCompletion(cmd)

	return cmd
}

func runList(opts *options) (err error) {
	conn, err := opts.Connection(connection.DefaultConfigRequireMasAuth)
	if err != nil {
		return err
	}

	api, kafkaInstance, err := conn.API().KafkaAdmin(opts.kafkaID)
	if err != nil {
		return err
	}

	req := api.AclsApi.GetAcls(opts.Context)

	req = req.Page(float32(opts.page)).Size(float32(opts.size))
	req = req.Order("asc").OrderKey("principal")

	permissionsData, httpRes, err := req.Execute()
	if httpRes != nil {
		defer httpRes.Body.Close()
	}

	if err != nil {
		if httpRes == nil {
			return err
		}

		operationTmplPair := localize.NewEntry("Operation", "list")

		switch httpRes.StatusCode {
		case http.StatusUnauthorized:
			return opts.localizer.MustLocalizeError("kafka.acl.common.error.unauthorized", operationTmplPair)
		case http.StatusForbidden:
			return opts.localizer.MustLocalizeError("kafka.acl.common.error.forbidden", operationTmplPair)
		case http.StatusInternalServerError:
			return opts.localizer.MustLocalizeError("kafka.acl.common.error.internalServerError")
		case http.StatusServiceUnavailable:
			return opts.localizer.MustLocalizeError("kafka.acl.common.error.unableToConnectToKafka", localize.NewEntry("Name", kafkaInstance.GetName()))
		default:
			return err
		}
	}

	switch opts.output {
	case dump.EmptyFormat:
		opts.Logger.Info("")
		permissions := permissionsData.GetItems()
		rows := acl.MapPermissionListToTableFormat(permissions, opts.localizer)
		dump.Table(opts.IO.Out, rows)
	default:
		return dump.Formatted(opts.IO.Out, opts.output, permissionsData)
	}

	return nil
}
