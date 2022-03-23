package versions

import (
	"context"

	"github.com/redhat-developer/app-services-cli/pkg/cmd/registry/registrycmdutil"
	"github.com/redhat-developer/app-services-cli/pkg/core/cmdutil/flagutil"
	"github.com/redhat-developer/app-services-cli/pkg/core/ioutil/dump"
	"github.com/redhat-developer/app-services-cli/pkg/core/ioutil/icon"
	"github.com/redhat-developer/app-services-cli/pkg/core/ioutil/iostreams"
	"github.com/redhat-developer/app-services-cli/pkg/core/localize"
	"github.com/redhat-developer/app-services-cli/pkg/core/logging"
	"github.com/redhat-developer/app-services-cli/pkg/core/servicecontext"
	"github.com/redhat-developer/app-services-cli/pkg/shared/connection"
	"github.com/redhat-developer/app-services-cli/pkg/shared/factory"
	"github.com/redhat-developer/app-services-cli/pkg/shared/profileutil"
	"github.com/spf13/cobra"
)

type options struct {
	artifact     string
	group        string
	outputFormat string

	registryID string

	IO             *iostreams.IOStreams
	Logger         logging.Logger
	Connection     factory.ConnectionFunc
	localizer      localize.Localizer
	Context        context.Context
	ServiceContext servicecontext.IContext
}

func NewVersionsCommand(f *factory.Factory) *cobra.Command {
	opts := &options{
		Connection:     f.Connection,
		IO:             f.IOStreams,
		localizer:      f.Localizer,
		Logger:         f.Logger,
		Context:        f.Context,
		ServiceContext: f.ServiceContext,
	}

	cmd := &cobra.Command{
		Use:     "versions",
		Short:   f.Localizer.MustLocalize("artifact.cmd.versions.description.short"),
		Long:    f.Localizer.MustLocalize("artifact.cmd.versions.description.long"),
		Example: f.Localizer.MustLocalize("artifact.cmd.versions.example"),
		Args:    cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {

			if opts.artifact == "" {
				return f.Localizer.MustLocalizeError("artifact.common.message.artifactIdRequired")
			}

			if opts.registryID != "" {
				return runGet(opts)
			}

			svcContext, err := opts.ServiceContext.Load()
			if err != nil {
				return err
			}

			profileHandler := &profileutil.ContextHandler{
				Context:   svcContext,
				Localizer: opts.localizer,
			}

			conn, err := opts.Connection(connection.DefaultConfigRequireMasAuth)
			if err != nil {
				return err
			}

			registryInstance, err := profileHandler.GetCurrentRegistryInstance(conn.API().ServiceRegistryMgmt())
			if err != nil {
				return err
			}

			opts.registryID = registryInstance.GetId()
			return runGet(opts)
		},
	}

	cmd.Flags().StringVar(&opts.artifact, "artifact-id", "", opts.localizer.MustLocalize("artifact.common.id"))
	cmd.Flags().StringVarP(&opts.group, "group", "g", registrycmdutil.DefaultArtifactGroup, opts.localizer.MustLocalize("artifact.common.group"))
	cmd.Flags().StringVar(&opts.registryID, "instance-id", "", opts.localizer.MustLocalize("registry.common.flag.instance.id"))
	cmd.Flags().StringVarP(&opts.outputFormat, "output", "o", "", opts.localizer.MustLocalize("artifact.common.message.output.format"))

	flagutil.EnableOutputFlagCompletion(cmd)

	return cmd
}

func runGet(opts *options) error {
	conn, err := opts.Connection(connection.DefaultConfigRequireMasAuth)
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

	opts.Logger.Info(opts.localizer.MustLocalize("artifact.common.message.artifact.versions.fetching"))

	request := dataAPI.VersionsApi.ListArtifactVersions(opts.Context, opts.group, opts.artifact)
	response, _, err := request.Execute()
	if err != nil {
		return registrycmdutil.TransformInstanceError(err)
	}

	opts.Logger.Info(icon.SuccessPrefix(), opts.localizer.MustLocalize("artifact.common.message.artifact.versions.fetched"))

	return dump.Formatted(opts.IO.Out, opts.outputFormat, response)
}
