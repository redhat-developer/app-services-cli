package versions

import (
	"context"

	"github.com/redhat-developer/app-services-cli/pkg/cmd/kafka/flagutil"
	cmdFlagUtil "github.com/redhat-developer/app-services-cli/pkg/cmdutil/flagutil"
	"github.com/redhat-developer/app-services-cli/pkg/connection"
	"github.com/redhat-developer/app-services-cli/pkg/dump"
	"github.com/redhat-developer/app-services-cli/pkg/icon"
	"github.com/redhat-developer/app-services-cli/pkg/iostreams"
	"github.com/redhat-developer/app-services-cli/pkg/localize"
	"github.com/redhat-developer/app-services-cli/pkg/serviceregistry/registryinstanceerror"
	"github.com/spf13/cobra"

	"github.com/redhat-developer/app-services-cli/internal/config"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/factory"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/registry/artifact/util"

	"github.com/redhat-developer/app-services-cli/pkg/logging"
)

type options struct {
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

func NewVersionsCommand(f *factory.Factory) *cobra.Command {
	opts := &options{
		Config:     f.Config,
		Connection: f.Connection,
		IO:         f.IOStreams,
		localizer:  f.Localizer,
		Logger:     f.Logger,
		Context:    f.Context,
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

	flags := flagutil.NewFlagSet(cmd, opts.localizer)
	flags.StringVar(&opts.artifact, "artifact-id", "", opts.localizer.MustLocalize("artifact.common.id"))
	flags.StringVarP(&opts.group, "group", "g", util.DefaultArtifactGroup, opts.localizer.MustLocalize("artifact.common.group"))
	flags.StringVar(&opts.registryID, "instance-id", "", opts.localizer.MustLocalize("artifact.common.instance.id"))
	flags.StringVarP(&opts.outputFormat, "output", "o", "", opts.localizer.MustLocalize("artifact.common.message.output.format"))

	cmdFlagUtil.EnableOutputFlagCompletion(cmd)

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

	if opts.group == util.DefaultArtifactGroup {
		opts.Logger.Info(opts.localizer.MustLocalize("artifact.common.message.no.group", localize.NewEntry("DefaultArtifactGroup", util.DefaultArtifactGroup)))
		opts.group = util.DefaultArtifactGroup
	}

	opts.Logger.Info(opts.localizer.MustLocalize("artifact.common.message.artifact.versions.fetching"))

	request := dataAPI.VersionsApi.ListArtifactVersions(opts.Context, opts.group, opts.artifact)
	response, _, err := request.Execute()
	if err != nil {
		return registryinstanceerror.TransformError(err)
	}

	opts.Logger.Info(icon.SuccessPrefix(), opts.localizer.MustLocalize("artifact.common.message.artifact.versions.fetched"))

	return dump.Formatted(opts.IO.Out, opts.outputFormat, response)
}
