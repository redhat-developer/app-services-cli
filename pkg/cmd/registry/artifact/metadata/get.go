package metadata

import (
	"context"

	"github.com/redhat-developer/app-services-cli/pkg/cmd/registry/artifact/util"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/registry/sdk"
	"github.com/redhat-developer/app-services-cli/pkg/core/cmdutil/factory"
	"github.com/redhat-developer/app-services-cli/pkg/core/cmdutil/flagutil"
	"github.com/redhat-developer/app-services-cli/pkg/core/config"
	"github.com/redhat-developer/app-services-cli/pkg/core/connection"
	"github.com/redhat-developer/app-services-cli/pkg/core/ioutil/color"
	"github.com/redhat-developer/app-services-cli/pkg/core/ioutil/dump"
	"github.com/redhat-developer/app-services-cli/pkg/core/ioutil/icon"
	"github.com/redhat-developer/app-services-cli/pkg/core/ioutil/iostreams"
	"github.com/redhat-developer/app-services-cli/pkg/core/localize"
	"github.com/redhat-developer/app-services-cli/pkg/core/logging"
	"github.com/redhat-developer/app-services-cli/pkg/serviceregistryutil"
	"github.com/spf13/cobra"
)

type GetOptions struct {
	artifact     string
	group        string
	outputFormat string

	registryID string

	IO         *iostreams.IOStreams
	Config     config.IConfig
	Logger     logging.Logger
	Connection factory.ConnectionFunc
	localizer  localize.Localizer
	Context    context.Context
}

// NewGetMetadataCommand creates a new command for fetching metadata for registry artifacts.
func NewGetMetadataCommand(f *factory.Factory) *cobra.Command {
	opts := &GetOptions{
		Config:     f.Config,
		Connection: f.Connection,
		IO:         f.IOStreams,
		localizer:  f.Localizer,
		Logger:     f.Logger,
		Context:    f.Context,
	}

	cmd := &cobra.Command{
		Use:     "metadata-get",
		Short:   f.Localizer.MustLocalize("artifact.cmd.metadata.get.description.short"),
		Long:    f.Localizer.MustLocalize("artifact.cmd.metadata.get.description.long"),
		Example: f.Localizer.MustLocalize("artifact.cmd.metadata.get.example"),
		Args:    cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			if opts.artifact == "" {
				return f.Localizer.MustLocalizeError("artifact.common.message.artifactIdRequired")
			}

			if opts.registryID != "" {
				return runGet(opts)
			}

			cfg, err := opts.Config.Load()
			if err != nil {
				return err
			}

			instanceID, ok := cfg.GetServiceRegistryIdOk()
			if !ok {
				return opts.localizer.MustLocalizeError("registry.no.service.selected.use.instance.id.flag")
			}

			opts.registryID = instanceID
			return runGet(opts)
		},
	}

	cmd.Flags().StringVar(&opts.artifact, "artifact-id", "", opts.localizer.MustLocalize("artifact.common.id"))
	cmd.Flags().StringVarP(&opts.group, "group", "g", util.DefaultArtifactGroup, opts.localizer.MustLocalize("artifact.common.group"))
	cmd.Flags().StringVar(&opts.registryID, "instance-id", "", opts.localizer.MustLocalize("artifact.common.instance.id"))
	cmd.Flags().StringVarP(&opts.outputFormat, "output", "o", "", opts.localizer.MustLocalize("artifact.common.message.output.format"))

	flagutil.EnableOutputFlagCompletion(cmd)

	return cmd
}

func runGet(opts *GetOptions) error {
	conn, err := opts.Connection(connection.DefaultConfigRequireMasAuth)
	if err != nil {
		return err
	}

	registry, _, err := serviceregistryutil.GetServiceRegistryByID(opts.Context, conn.API().ServiceRegistryMgmt(), opts.registryID)
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
	response, _, err := request.Execute()
	if err != nil {
		return sdk.TransformInstanceError(err)
	}

	artifactURL, ok := util.GetArtifactURL(registry, &response)

	opts.Logger.Info(icon.SuccessPrefix(), opts.localizer.MustLocalize("artifact.common.message.artifact.metadata.fetched"))

	if ok {
		opts.Logger.Info(opts.localizer.MustLocalize("artifact.common.webURL", localize.NewEntry("URL", color.Info(artifactURL))))
	}

	return dump.Formatted(opts.IO.Out, opts.outputFormat, response)
}
