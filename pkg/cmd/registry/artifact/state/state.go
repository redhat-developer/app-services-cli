package state

import (
	"context"

	"github.com/redhat-developer/app-services-cli/pkg/connection"
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

type options struct {
	artifact string
	group    string

	registryID string

	state string

	IO         *iostreams.IOStreams
	Config     config.IConfig
	Logger     logging.Logger
	Connection factory.ConnectionFunc
	localizer  localize.Localizer
	context    context.Context
}

// NewSetMetadataCommand creates a new command for updating metadata for registry artifacts.
func NewSetStateCommand(f *factory.Factory) *cobra.Command {
	opts := &options{
		Config:     f.Config,
		Connection: f.Connection,
		IO:         f.IOStreams,
		localizer:  f.Localizer,
		Logger:     f.Logger,
		context:    f.Context,
	}

	cmd := &cobra.Command{
		Use:     "state-set",
		Short:   f.Localizer.MustLocalize("artifact.cmd.stateset.description.short"),
		Long:    f.Localizer.MustLocalize("artifact.cmd.stateset.description.long"),
		Example: f.Localizer.MustLocalize("artifact.cmd.stateset.example"),
		Args:    cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			if opts.artifact == "" {
				return f.Localizer.MustLocalizeError("artifact.common.message.artifactIdRequired")
			}

			if _, err := registryinstanceclient.NewArtifactStateFromValue(opts.state); err != nil {
				return opts.localizer.MustLocalizeError("artifact.cmd.state.error.invalidArtifactState", localize.NewEntry("AllowedTypes", util.GetAllowedArtifactStateEnumValuesAsString()))
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
	cmd.Flags().StringVar(&opts.state, "state", "", opts.localizer.MustLocalize("artifact.flag.state.description"))

	_ = cmd.RegisterFlagCompletionFunc("state", func(cmd *cobra.Command, _ []string, _ string) ([]string, cobra.ShellCompDirective) {
		return util.AllowedArtifactStateEnumValues, cobra.ShellCompDirectiveNoSpace
	})
	return cmd
}

func runSet(opts *options) error {
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

	updateState, err := registryinstanceclient.NewArtifactStateFromValue(opts.state)
	if err != nil {
		return err
	}

	request := dataAPI.ArtifactsApi.UpdateArtifactState(opts.context, opts.group, opts.artifact)
	_, err = request.UpdateState(*registryinstanceclient.NewUpdateState(*updateState)).Execute()
	if err != nil {
		return registryinstanceerror.TransformError(err)
	}

	opts.Logger.Info(opts.localizer.MustLocalize("artifact.common.message.artifact.state.updated"))
	return nil
}
