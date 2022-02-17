package delete

import (
	"context"
	"errors"

	"github.com/AlecAivazis/survey/v2"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/registry/registrycmdutil"
	"github.com/redhat-developer/app-services-cli/pkg/core/cmdutil/flagutil"
	"github.com/redhat-developer/app-services-cli/pkg/core/config"
	"github.com/redhat-developer/app-services-cli/pkg/core/ioutil/iostreams"
	"github.com/redhat-developer/app-services-cli/pkg/core/localize"
	"github.com/redhat-developer/app-services-cli/pkg/core/logging"
	"github.com/redhat-developer/app-services-cli/pkg/shared/connection"
	"github.com/redhat-developer/app-services-cli/pkg/shared/factory"
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
				return flagutil.RequiredWhenNonInteractiveError("yes")
			}

			if opts.registryID != "" {
				return runDelete(opts)
			}

			cfg, err := opts.Config.Load()
			if err != nil {
				return err
			}

			instanceID, ok := cfg.GetServiceRegistryIdOk()
			if !ok {
				return opts.localizer.MustLocalizeError("artifact.cmd.common.error.noServiceRegistrySelected")
			}

			opts.registryID = instanceID
			return runDelete(opts)
		},
	}

	cmd.Flags().BoolVarP(&opts.force, "yes", "y", false, opts.localizer.MustLocalize("artifact.common.delete.without.prompt"))
	cmd.Flags().StringVar(&opts.artifact, "artifact-id", "", opts.localizer.MustLocalize("artifact.common.id"))
	cmd.Flags().StringVarP(&opts.group, "group", "g", registrycmdutil.DefaultArtifactGroup, opts.localizer.MustLocalize("artifact.common.group"))
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

	if opts.group == registrycmdutil.DefaultArtifactGroup {
		opts.Logger.Info(opts.localizer.MustLocalize("registry.artifact.common.message.no.group", localize.NewEntry("DefaultArtifactGroup", registrycmdutil.DefaultArtifactGroup)))
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
			return registrycmdutil.TransformInstanceError(err)
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
			return registrycmdutil.TransformInstanceError(err)
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
