package metadata

import (
	"context"
	"encoding/json"
	"errors"

	flagutil "github.com/redhat-developer/app-services-cli/pkg/cmdutil/flags"
	"github.com/redhat-developer/app-services-cli/pkg/connection"
	"github.com/redhat-developer/app-services-cli/pkg/editor"
	"github.com/redhat-developer/app-services-cli/pkg/iostreams"
	"github.com/redhat-developer/app-services-cli/pkg/localize"
	"github.com/redhat-developer/app-services-cli/pkg/serviceregistry/registryinstanceerror"
	"github.com/spf13/cobra"

	"github.com/redhat-developer/app-services-cli/internal/config"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/factory"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/registry/artifact/util"
	"github.com/redhat-developer/app-services-cli/pkg/logging"
	registryinstanceclient "github.com/redhat-developer/app-services-sdk-go/registryinstance/apiv1internal/client"
)

type SetOptions struct {
	artifact     string
	group        string
	outputFormat string

	registryID string

	name        string
	description string

	IO         *iostreams.IOStreams
	Config     config.IConfig
	Logger     logging.Logger
	Connection factory.ConnectionFunc
	localizer  localize.Localizer
}

// NewSetMetadataCommand creates a new command for updating metadata for registry artifacts.
func NewSetMetadataCommand(f *factory.Factory) *cobra.Command {
	opts := &SetOptions{
		Config:     f.Config,
		Connection: f.Connection,
		IO:         f.IOStreams,
		localizer:  f.Localizer,
		Logger:     f.Logger,
	}

	cmd := &cobra.Command{
		Use:   "metadata-set",
		Short: "Update artifact metadata",
		Long: `
Updates the metadata for an artifact in the service registry. 
Editable metadata includes fields like name and description
`,
		Example: `
##  Update the metadata for an artifact
rhoas service-registry artifact metadata-set
		`,
		Args: cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			if opts.name == "" && opts.description == "" && !opts.IO.CanPrompt() {
				return errors.New("Editor mode cannot be started in non-interactive mode. Please use --name and --description flags")
			}

			if opts.artifact == "" {
				return errors.New("artifact id is required. Please specify artifact by using --artifact-id flag")
			}

			if opts.registryID != "" {
				return runSet(opts)
			}

			cfg, err := opts.Config.Load()
			if err != nil {
				return err
			}

			if !cfg.HasServiceRegistry() {
				return errors.New("no service registry selected. Please specify registry by using --instance-id flag")
			}

			opts.registryID = cfg.Services.ServiceRegistry.InstanceID
			return runSet(opts)
		},
	}

	cmd.Flags().StringVarP(&opts.artifact, "artifact-id", "a", "", "Id of the artifact")
	cmd.Flags().StringVarP(&opts.group, "group", "g", util.DefaultArtifactGroup, "Artifact group")
	cmd.Flags().StringVar(&opts.registryID, "instance-id", "", "Id of the registry to be used. By default uses currently selected registry")
	cmd.Flags().StringVarP(&opts.outputFormat, "output", "o", "", "Output format (json, yaml, yml)")

	cmd.Flags().StringVar(&opts.name, "name", "", "Custom name of the artifact")
	cmd.Flags().StringVar(&opts.description, "description", "", "Custom description of the artifact")

	flagutil.EnableOutputFlagCompletion(cmd)

	return cmd
}

func runSet(opts *SetOptions) error {
	conn, err := opts.Connection(connection.DefaultConfigRequireMasAuth)
	if err != nil {
		return err
	}

	dataAPI, _, err := conn.API().ServiceRegistryInstance(opts.registryID)
	if err != nil {
		return err
	}

	if opts.group == util.DefaultArtifactGroup {
		opts.Logger.Info("Group was not specified. Using 'default' artifacts group.")
		opts.group = util.DefaultArtifactGroup
	}

	opts.Logger.Info("Fetching current artifact metadata")

	ctx := context.Background()
	request := dataAPI.MetadataApi.GetArtifactMetaData(ctx, opts.group, opts.artifact)
	currentMetadata, _, err := request.Execute()
	if err != nil {
		return registryinstanceerror.TransformError(err)
	}

	editableMedata := &registryinstanceclient.EditableMetaData{
		Name:        currentMetadata.Name,
		Description: currentMetadata.Description,
		Labels:      currentMetadata.Labels,
		Properties:  currentMetadata.Properties,
	}

	if opts.name != "" || opts.description != "" {
		if opts.name != "" {
			editableMedata.Name = &opts.name
		}

		if opts.description != "" {
			editableMedata.Description = &opts.description
		}
	} else {
		opts.Logger.Info("Running edito to let you edit metadata. Please close editor to continue...")
		editableMedata, err = runEditor(editableMedata)
		if err != nil {
			return err
		}
	}

	opts.Logger.Info("Updating artifact metadata")

	editRequest := dataAPI.MetadataApi.UpdateArtifactMetaData(ctx, opts.group, opts.artifact)
	_, err = editRequest.EditableMetaData(*editableMedata).Execute()
	if err != nil {
		return registryinstanceerror.TransformError(err)
	}

	opts.Logger.Info("Successfully updated artifact metadata")
	return nil
}

func runEditor(currentMetadata *registryinstanceclient.EditableMetaData) (*registryinstanceclient.EditableMetaData, error) {
	// Fill defaults for json fields
	if currentMetadata.Labels == nil {
		currentMetadata.Labels = &[]string{}
	}
	if currentMetadata.Properties == nil {
		currentMetadata.Properties = &map[string]string{}
	}
	metadataJson, err := json.MarshalIndent(currentMetadata, "", " ")
	if err != nil {
		return nil, err
	}
	editor := editor.New(metadataJson, "metadata.json")
	output, err := editor.Run()
	if err != nil {
		return nil, err
	}
	var resultData registryinstanceclient.EditableMetaData
	err = json.Unmarshal(output, &resultData)

	if err != nil {
		return nil, err
	}
	return &resultData, nil
}
