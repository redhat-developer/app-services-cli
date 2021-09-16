package metadata

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/redhat-developer/app-services-cli/pkg/icon"

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
	Context    context.Context
}

// NewSetMetadataCommand creates a new command for updating metadata for registry artifacts.
func NewSetMetadataCommand(f *factory.Factory) *cobra.Command {
	opts := &SetOptions{
		Config:     f.Config,
		Connection: f.Connection,
		IO:         f.IOStreams,
		localizer:  f.Localizer,
		Logger:     f.Logger,
		Context:    f.Context,
	}

	cmd := &cobra.Command{
		Use:     "metadata-set",
		Short:   f.Localizer.MustLocalize("artifact.cmd.metadata.set.description.short"),
		Long:    f.Localizer.MustLocalize("artifact.cmd.metadata.set.description.long"),
		Example: f.Localizer.MustLocalize("artifact.cmd.metadata.set.example"),
		Args:    cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			if opts.name == "" && opts.description == "" && !opts.IO.CanPrompt() {
				return f.Localizer.MustLocalizeError("artifact.cmd.common.error.no.editor.mode.in.non.interactive")
			}

			if opts.artifact == "" {
				return errors.New(f.Localizer.MustLocalize("artifact.common.message.artifactIdRequired"))
			}

			if opts.registryID != "" {
				return runSet(opts)
			}

			cfg, err := opts.Config.Load()
			if err != nil {
				return err
			}

			if !cfg.HasServiceRegistry() {
				return opts.localizer.MustLocalizeError("registry.no.service.selected.use.instance.id.flag")
			}

			opts.registryID = cfg.Services.ServiceRegistry.InstanceID
			return runSet(opts)
		},
	}

	cmd.Flags().StringVar(&opts.artifact, "artifact-id", "", opts.localizer.MustLocalize("artifact.common.id"))
	cmd.Flags().StringVarP(&opts.group, "group", "g", util.DefaultArtifactGroup, opts.localizer.MustLocalize("artifact.common.group"))
	cmd.Flags().StringVar(&opts.registryID, "instance-id", "", opts.localizer.MustLocalize("artifact.common.instance.id"))
	cmd.Flags().StringVarP(&opts.outputFormat, "output", "o", "", opts.localizer.MustLocalize("artifact.common.message.output.format"))

	cmd.Flags().StringVar(&opts.name, "name", "", opts.localizer.MustLocalize("artifact.common.custom.name"))
	cmd.Flags().StringVar(&opts.description, "description", "", opts.localizer.MustLocalize("artifact.common.custom.description"))

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
		opts.Logger.Info(opts.localizer.MustLocalize("artifact.common.message.no.group", localize.NewEntry("DefaultArtifactGroup", util.DefaultArtifactGroup)))
		opts.group = util.DefaultArtifactGroup
	}

	opts.Logger.Info(opts.localizer.MustLocalize("artifact.common.message.artifact.metadata.fetching"))

	request := dataAPI.MetadataApi.GetArtifactMetaData(opts.Context, opts.group, opts.artifact)
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
		opts.Logger.Info(opts.localizer.MustLocalize("artifact.common.message.running.editor.with.editable.metadata"))
		editableMedata, err = runEditor(editableMedata)
		if err != nil {
			return err
		}
	}

	opts.Logger.Info(opts.localizer.MustLocalize("artifact.common.message.artifact.metadata.updating"))

	editRequest := dataAPI.MetadataApi.UpdateArtifactMetaData(opts.Context, opts.group, opts.artifact)
	_, err = editRequest.EditableMetaData(*editableMedata).Execute()
	if err != nil {
		return registryinstanceerror.TransformError(err)
	}

	opts.Logger.Info(icon.SuccessPrefix(), opts.localizer.MustLocalize("artifact.common.message.artifact.metadata.updated"))
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
	systemEditor := editor.New(metadataJson, "metadata.json")
	output, err := systemEditor.Run()
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
