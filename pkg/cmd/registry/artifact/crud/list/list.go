package list

import (
	"context"
	"encoding/json"
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

	"gopkg.in/yaml.v2"
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

type Options struct {
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
}

// NewListCommand creates a new command for listing registry artifacts.
func NewListCommand(f *factory.Factory) *cobra.Command {
	opts := &Options{
		Config:     f.Config,
		Connection: f.Connection,
		Logger:     f.Logger,
		IO:         f.IOStreams,
		localizer:  f.Localizer,
	}

	cmd := &cobra.Command{
		Use:   "list",
		Short: "List artifacts",
		Long:  "List all artifacts for the group by specified output format (by default table)",
		Example: `
## List all artifacts for the "default" artifacts group
rhoas service-registry artifact list

## List all artifacts with "my-group" group 
rhoas service-registry artifact list --group=my-group

## List all artifacts with limit and group
rhoas service-registry artifact list --page=2 --limit=10
		`,
		Args: cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			if opts.outputFormat != "" && !flagutil.IsValidInput(opts.outputFormat, flagutil.ValidOutputFormats...) {
				return flag.InvalidValueError("output", opts.outputFormat, flagutil.ValidOutputFormats...)
			}

			if opts.page < 1 || opts.limit < 1 {
				return errors.New("page and limit values should be bigger than 1")
			}

			if opts.registryID != "" {
				return runList(opts)
			}

			cfg, err := opts.Config.Load()
			if err != nil {
				return err
			}

			if !cfg.HasServiceRegistry() {
				return errors.New("no service Registry selected. Use rhoas registry use to select your registry")
			}

			opts.registryID = cfg.Services.ServiceRegistry.InstanceID

			return runList(opts)
		},
	}

	cmd.Flags().StringVarP(&opts.group, "group", "g", util.DefaultArtifactGroup, "Artifact group")
	cmd.Flags().Int32VarP(&opts.page, "page", "", 1, "Page number")
	cmd.Flags().Int32VarP(&opts.limit, "limit", "", 100, "Page limit")

	cmd.Flags().StringVar(&opts.registryID, "instance-id", "", "Id of the registry to be used. By default uses currently selected registry")
	cmd.Flags().StringVarP(&opts.outputFormat, "output", "o", "", "Output format (json, yaml, yml)")

	flagutil.EnableOutputFlagCompletion(cmd)

	return cmd
}

func runList(opts *Options) error {
	if opts.group == util.DefaultArtifactGroup {
		opts.Logger.Info("Group was not specified. Using", util.DefaultArtifactGroup, "artifacts group.")
		opts.group = util.DefaultArtifactGroup
	}

	connection, err := opts.Connection(connection.DefaultConfigRequireMasAuth)
	if err != nil {
		return err
	}

	api := connection.API()

	a, _, err := api.ServiceRegistryInstance(opts.registryID)
	if err != nil {
		return err
	}
	request := a.ArtifactsApi.ListArtifactsInGroup(context.Background(), opts.group)

	request = request.Offset((opts.page - 1) * opts.limit)
	request = request.Limit(opts.limit)
	request = request.Orderby(registryinstanceclient.SORTBY_CREATED_ON)
	request = request.Order(registryinstanceclient.SORTORDER_ASC)

	response, _, err := request.Execute()
	if err != nil {
		return registryinstanceerror.TransformError(err)
	}

	if len(response.Artifacts) == 0 && opts.outputFormat == "" {
		opts.Logger.Info("No artifacts available for " + opts.group + " group and registry id " + opts.registryID)
		return nil
	}

	switch opts.outputFormat {
	case "json":
		data, _ := json.Marshal(response)
		_ = dump.JSON(opts.IO.Out, data)
	case "yaml", "yml":
		data, _ := yaml.Marshal(response)
		_ = dump.YAML(opts.IO.Out, data)
	default:
		rows := mapResponseItemsToRows(response.Artifacts)
		dump.Table(opts.IO.Out, rows)
		opts.Logger.Info("")
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
