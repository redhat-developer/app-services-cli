package state

import (
	"context"

	"github.com/redhat-developer/app-services-cli/pkg/cmd/registry/artifact/util"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/registry/registrycmdutil"
	"github.com/redhat-developer/app-services-cli/pkg/core/ioutil/iostreams"
	"github.com/redhat-developer/app-services-cli/pkg/core/localize"
	"github.com/redhat-developer/app-services-cli/pkg/core/logging"
	"github.com/redhat-developer/app-services-cli/pkg/core/servicecontext"
	"github.com/redhat-developer/app-services-cli/pkg/shared/contextutil"
	"github.com/redhat-developer/app-services-cli/pkg/shared/factory"
	"github.com/spf13/cobra"

	registryinstanceclient "github.com/jackdelahunt/app-services-sdk-core/app-services-sdk-go/registryinstance/apiv1internal/client"
)

type options struct {
	artifact string
	group    string

	registryID string

	state string

	IO             *iostreams.IOStreams
	Logger         logging.Logger
	Connection     factory.ConnectionFunc
	localizer      localize.Localizer
	context        context.Context
	ServiceContext servicecontext.IContext
}

// NewSetMetadataCommand creates a new command for updating metadata for registry artifacts.
func NewSetStateCommand(f *factory.Factory) *cobra.Command {
	opts := &options{
		Connection:     f.Connection,
		IO:             f.IOStreams,
		localizer:      f.Localizer,
		Logger:         f.Logger,
		context:        f.Context,
		ServiceContext: f.ServiceContext,
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

			registryInstance, err := contextutil.GetCurrentRegistryInstance(f)
			if err != nil {
				return err
			}

			opts.registryID = registryInstance.GetId()
			return runSet(opts)
		},
	}

	cmd.Flags().StringVar(&opts.artifact, "artifact-id", "", opts.localizer.MustLocalize("artifact.common.id"))
	cmd.Flags().StringVarP(&opts.group, "group", "g", registrycmdutil.DefaultArtifactGroup, opts.localizer.MustLocalize("artifact.common.group"))
	cmd.Flags().StringVar(&opts.registryID, "instance-id", "", opts.localizer.MustLocalize("registry.common.flag.instance.id"))
	cmd.Flags().StringVar(&opts.state, "state", "", opts.localizer.MustLocalize("artifact.flag.state.description"))

	_ = cmd.RegisterFlagCompletionFunc("state", func(cmd *cobra.Command, _ []string, _ string) ([]string, cobra.ShellCompDirective) {
		return util.AllowedArtifactStateEnumValues, cobra.ShellCompDirectiveNoSpace
	})
	return cmd
}

func runSet(opts *options) error {
	conn, err := opts.Connection()
	if err != nil {
		return err
	}

	dataAPI, _, err := conn.API().ServiceRegistryInstance(opts.registryID)
	if err != nil {
		return err
	}

	if opts.group == registrycmdutil.DefaultArtifactGroup {
		opts.Logger.Info(opts.localizer.MustLocalize("registry.artifact.common.message.no.group", localize.NewEntry("DefaultArtifactGroup", registrycmdutil.DefaultArtifactGroup)))
	}

	updateState, err := registryinstanceclient.NewArtifactStateFromValue(opts.state)
	if err != nil {
		return err
	}

	request := dataAPI.ArtifactsApi.UpdateArtifactState(opts.context, opts.group, opts.artifact)
	_, err = request.UpdateState(*registryinstanceclient.NewUpdateState(*updateState)).Execute()
	if err != nil {
		return registrycmdutil.TransformInstanceError(err)
	}

	opts.Logger.Info(opts.localizer.MustLocalize("artifact.common.message.artifact.state.updated"))
	return nil
}
