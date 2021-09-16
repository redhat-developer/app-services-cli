package list

import (
	"context"
	"errors"

	flagutil "github.com/redhat-developer/app-services-cli/pkg/cmdutil/flags"
	"github.com/redhat-developer/app-services-cli/pkg/connection"
	"github.com/redhat-developer/app-services-cli/pkg/iostreams"
	"github.com/redhat-developer/app-services-cli/pkg/localize"
	"github.com/redhat-developer/app-services-cli/pkg/serviceregistry/registryinstanceerror"
	registryinstanceclient "github.com/redhat-developer/app-services-sdk-go/registryinstance/apiv1internal/client"

	"github.com/redhat-developer/app-services-cli/pkg/dump"

	"github.com/spf13/cobra"

	"github.com/redhat-developer/app-services-cli/internal/config"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/factory"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/flag"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/registry/artifact/util"
	"github.com/redhat-developer/app-services-cli/pkg/logging"
)

// row is the details of a Service Registry instance needed to print to a table
type artifactRow struct {
	// The ID of a single artifact.
	Id string `json:"id" header:"ID"`

	Name string `json:"name,omitempty" header:"Name"`

	CreatedOn string `json:"createdOn" header:"Created on"`

	CreatedBy string `json:"createdBy" header:"Created By"`

	Type registryinstanceclient.ArtifactType `json:"type" header:"Type"`

	State registryinstanceclient.ArtifactState `json:"state" header:"State"`
}

type options struct {
	group string

	registryID   string
	outputFormat string

	page  int32
	limit int32

	IO         *iostreams.IOStreams
	Config     config.IConfig
	Connection factory.ConnectionFunc
	Logger     logging.Logger
	localizer  localize.Localizer
	Context    context.Context
}

// NewListCommand creates a new command for listing registry artifacts.
func NewListCommand(f *factory.Factory) *cobra.Command {
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
		Short:   f.Localizer.MustLocalize("artifact.cmd.list.description.short"),
		Long:    f.Localizer.MustLocalize("artifact.cmd.list.description.long"),
		Example: f.Localizer.MustLocalize("artifact.cmd.list.example"),
		Args:    cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			if opts.outputFormat != "" && !flagutil.IsValidInput(opts.outputFormat, flagutil.ValidOutputFormats...) {
				return flag.InvalidValueError("output", opts.outputFormat, flagutil.ValidOutputFormats...)
			}

			if opts.page < 1 || opts.limit < 1 {
				return errors.New(opts.localizer.MustLocalize("artifact.common.error.page.and.limit.too.small"))
			}

			if opts.registryID != "" {
				return runList(opts)
			}

			cfg, err := opts.Config.Load()
			if err != nil {
				return err
			}

			if !cfg.HasServiceRegistry() {
				return errors.New(opts.localizer.MustLocalize("artifact.cmd.common.error.noServiceRegistrySelected"))
			}

			opts.registryID = cfg.Services.ServiceRegistry.InstanceID

			return runList(opts)
		},
	}

	cmd.Flags().StringVarP(&opts.group, "group", "g", util.DefaultArtifactGroup, opts.localizer.MustLocalize("artifact.common.group"))
	cmd.Flags().Int32VarP(&opts.page, "page", "", 1, opts.localizer.MustLocalize("artifact.common.page.number"))
	cmd.Flags().Int32VarP(&opts.limit, "limit", "", 100, opts.localizer.MustLocalize("artifact.common.page.limit"))

	cmd.Flags().StringVar(&opts.registryID, "instance-id", "", opts.localizer.MustLocalize("artifact.common.instance.id"))
	cmd.Flags().StringVarP(&opts.outputFormat, "output", "o", "", opts.localizer.MustLocalize("artifact.common.message.output.format"))

	flagutil.EnableOutputFlagCompletion(cmd)

	return cmd
}

func runList(opts *options) error {
	if opts.group == util.DefaultArtifactGroup {
		opts.Logger.Info(opts.localizer.MustLocalize("artifact.common.message.no.group", localize.NewEntry("DefaultArtifactGroup", util.DefaultArtifactGroup)))
		opts.group = util.DefaultArtifactGroup
	}

	conn, err := opts.Connection(connection.DefaultConfigRequireMasAuth)
	if err != nil {
		return err
	}

	api := conn.API()

	a, _, err := api.ServiceRegistryInstance(opts.registryID)
	if err != nil {
		return err
	}
	request := a.ArtifactsApi.ListArtifactsInGroup(opts.Context, opts.group)

	request = request.Offset((opts.page - 1) * opts.limit)
	request = request.Limit(opts.limit)
	request = request.Orderby(registryinstanceclient.SORTBY_CREATED_ON)
	request = request.Order(registryinstanceclient.SORTORDER_ASC)

	response, _, err := request.Execute()
	if err != nil {
		return registryinstanceerror.TransformError(err)
	}

	if len(response.Artifacts) == 0 && opts.outputFormat == "" {
		opts.Logger.Info(opts.localizer.MustLocalize("artifact.common.message.no.artifact.available.for.group.and.registry", localize.NewEntry("Group", opts.group), localize.NewEntry("Registry", opts.registryID)))
		return nil
	}

	switch opts.outputFormat {
	case dump.TableFormat:
		rows := mapResponseItemsToRows(response.Artifacts)
		dump.Table(opts.IO.Out, rows)
		opts.Logger.Info("")
	default:
		dump.PrintDataInFormat(opts.outputFormat, response, opts.IO.Out)
	}

	return nil
}

func mapResponseItemsToRows(artifacts []registryinstanceclient.SearchedArtifact) []artifactRow {
	rows := []artifactRow{}

	for i := range artifacts {
		k := (artifacts)[i]
		row := artifactRow{
			Id:        k.GetId(),
			Name:      k.GetName(),
			CreatedOn: k.GetCreatedOn(),
			CreatedBy: k.GetCreatedBy(),
			Type:      k.GetType(),
			State:     k.GetState(),
		}

		rows = append(rows, row)
	}

	return rows
}
