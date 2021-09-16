package delete

import (
	"context"
	"errors"

	flagutil "github.com/redhat-developer/app-services-cli/pkg/cmdutil/flags"

	"github.com/redhat-developer/app-services-cli/pkg/connection"
	"github.com/redhat-developer/app-services-cli/pkg/localize"

	"github.com/redhat-developer/app-services-cli/pkg/iostreams"
	"github.com/redhat-developer/app-services-cli/pkg/serviceregistry/registryinstanceerror"

	"github.com/redhat-developer/app-services-cli/pkg/logging"

	"github.com/AlecAivazis/survey/v2"
	"github.com/redhat-developer/app-services-cli/internal/config"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/factory"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/flag"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/registry/artifact/util"
	"github.com/spf13/cobra"
)

type options struct {
	artifact string
	group    string

	registryID string
	force      bool

	IO         *iostreams.IOStreams
	Config     config.IConfig
	Connection factory.ConnectionFunc
	Logger     logging.Logger
	localizer  localize.Localizer
	Context    context.Context
}

func NewDeleteCommand(f *factory.Factory) *cobra.Command {
	opts := &options{
		Config:     f.Config,
		Connection: f.Connection,
		Logger:     f.Logger,
		IO:         f.IOStreams,
		localizer:  f.Localizer,
		Context:    f.Context,
	}

	cmd := &cobra.Command{
		Use:     "delete",
		Short:   f.Localizer.MustLocalize("artifact.cmd.delete.description.short"),
		Long:    f.Localizer.MustLocalize("artifact.cmd.delete.description.long"),
		Example: f.Localizer.MustLocalize("artifact.cmd.delete.example"),
		Args:    cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			if !opts.IO.CanPrompt() && !opts.force {
				return flag.RequiredWhenNonInteractiveError("yes")
			}

			if opts.registryID != "" {
				return runDelete(opts)
			}

			cfg, err := opts.Config.Load()
			if err != nil {
				return err
			}

			if !cfg.HasServiceRegistry() {
				return opts.localizer.MustLocalizeError("artifact.cmd.common.error.noServiceRegistrySelected")
			}

			opts.registryID = cfg.Services.ServiceRegistry.InstanceID
			return runDelete(opts)
		},
	}

	cmd.Flags().BoolVarP(&opts.force, "yes", "y", false, opts.localizer.MustLocalize("artifact.common.delete.without.prompt"))
	cmd.Flags().StringVar(&opts.artifact, "artifact-id", "", opts.localizer.MustLocalize("artifact.common.id"))
	cmd.Flags().StringVarP(&opts.group, "group", "g", util.DefaultArtifactGroup, opts.localizer.MustLocalize("artifact.common.group"))
	cmd.Flags().StringVar(&opts.registryID, "instance-id", "", opts.localizer.MustLocalize("artifact.common.registryIdToUse"))
	flagutil.EnableOutputFlagCompletion(cmd)

	return cmd
}

func runDelete(opts *options) error {
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

	if opts.artifact == "" {
		opts.Logger.Info(opts.localizer.MustLocalize("artifact.common.message.deleteAllArtifactsInGroup"))
		err = confirmDelete(opts, opts.localizer.MustLocalize("artifact.common.message.deleteAllArtifactsFromGroup", localize.NewEntry("GroupName", opts.group)))
		if err != nil {
			return err
		}
		request := dataAPI.ArtifactsApi.DeleteArtifactsInGroup(opts.Context, opts.group)
		_, err = request.Execute()
		if err != nil {
			return registryinstanceerror.TransformError(err)
		}
		opts.Logger.Info(opts.localizer.MustLocalize("artifact.common.message.AllArtifactsInGroupDeleted", localize.NewEntry("GroupName", opts.group)))
	} else {
		_, _, err := dataAPI.MetadataApi.GetArtifactMetaData(opts.Context, opts.group, opts.artifact).Execute()
		if err != nil {
			return opts.localizer.MustLocalizeError("artifact.common.error.artifact.notFound", localize.NewEntry("Name", opts.artifact))
		}
		opts.Logger.Info(opts.localizer.MustLocalize("artifact.common.message.deleting.artifact", localize.NewEntry("Name", opts.artifact)))
		err = confirmDelete(opts, opts.localizer.MustLocalize("artifact.common.message.deleting.artifactFromGroup", localize.NewEntry("Name", opts.artifact), localize.NewEntry("Group", opts.group)))
		if err != nil {
			return err
		}
		request := dataAPI.ArtifactsApi.DeleteArtifact(opts.Context, opts.group, opts.artifact)

		_, err = request.Execute()
		if err != nil {
			return registryinstanceerror.TransformError(err)
		}
		opts.Logger.Info(opts.localizer.MustLocalize("artifact.common.message.deleted", localize.NewEntry("Name", opts.artifact)))
	}

	return nil
}

func confirmDelete(opts *options, message string) error {
	if !opts.force {
		var shouldContinue bool
		confirm := &survey.Confirm{
			Message: message,
		}
		err := survey.AskOne(confirm, &shouldContinue)
		if err != nil {
			return err
		}

		if !shouldContinue {
			return errors.New("command stopped by user")
		}
	}
	return nil
}
