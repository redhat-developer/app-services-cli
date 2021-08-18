package delete

import (
	"context"
	"errors"
	"fmt"

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

type Options struct {
	artifact string
	group    string

	registryID string

	outputFormat string
	force        bool

	IO         *iostreams.IOStreams
	Config     config.IConfig
	Connection factory.ConnectionFunc
	Logger     func() (logging.Logger, error)
	localizer  localize.Localizer
}

func NewDeleteCommand(f *factory.Factory) *cobra.Command {
	opts := &Options{
		Config:     f.Config,
		Connection: f.Connection,
		Logger:     f.Logger,
		IO:         f.IOStreams,
		localizer:  f.Localizer,
	}

	cmd := &cobra.Command{
		Use:   "delete",
		Short: "Deletes single or all artifacts in a given group",
		Long: `
Deletes single or all artifacts in a given group. 

Delete command works in two modes:

	- When --artifact-id argument is missing delete will delete all artifacts in the group
	- When --artifact-id is specified delete deletes only single artifact and its version

When --group parameter is missing the command will create a new artifact under the "default" group.
		`,
		Example: `
## Delete all artifacts in the group "default"
rhoas service-registry artifact delete 

## Delete artifact in the group "default" with name "my-artifact"
rhoas service-registry artifact delete my-artifact
		`,
		Args: cobra.RangeArgs(0, 1),
		RunE: func(cmd *cobra.Command, args []string) error {
			if !opts.IO.CanPrompt() && !opts.force {
				return flag.RequiredWhenNonInteractiveError("yes")
			}

			validOutputFormats := flagutil.ValidOutputFormats
			if opts.outputFormat != "" && !flagutil.IsValidInput(opts.outputFormat, validOutputFormats...) {
				return flag.InvalidValueError("output", opts.outputFormat, validOutputFormats...)
			}

			if len(args) > 0 {
				opts.artifact = args[0]
			}

			if opts.registryID != "" {
				return runDelete(opts)
			}

			cfg, err := opts.Config.Load()
			if err != nil {
				return err
			}

			if !cfg.HasServiceRegistry() {
				return errors.New("no service Registry selected. Use 'rhoas service-registry use' use to select your registry")
			}

			opts.registryID = fmt.Sprint(cfg.Services.ServiceRegistry.InstanceID)
			return runDelete(opts)
		},
	}

	cmd.Flags().BoolVarP(&opts.force, "yes", "y", false, "Delete without prompt")
	cmd.Flags().StringVarP(&opts.outputFormat, "output", "o", "json", opts.localizer.MustLocalize("registry.cmd.flag.output.description"))

	cmd.Flags().StringVarP(&opts.artifact, "artifact-id", "a", "", "Id of the artifact")
	cmd.Flags().StringVarP(&opts.group, "group", "g", util.DefaultArtifactGroup, "Group of the artifact")
	cmd.Flags().StringVar(&opts.registryID, "instance-id", "", "Id of the registry to be used. By default uses currently selected registry")
	flagutil.EnableOutputFlagCompletion(cmd)

	return cmd
}

func runDelete(opts *Options) error {
	logger, err := opts.Logger()
	if err != nil {
		return err
	}

	conn, err := opts.Connection(connection.DefaultConfigRequireMasAuth)
	if err != nil {
		return err
	}

	dataAPI, _, err := conn.API().ServiceRegistryInstance(opts.registryID)
	if err != nil {
		return err
	}

	if opts.group == util.DefaultArtifactGroup {
		logger.Info("Group was not specified. Using", util.DefaultArtifactGroup, "artifacts group.")
		opts.group = util.DefaultArtifactGroup
	}

	ctx := context.Background()
	if opts.artifact == "" {
		logger.Info("Artifact was not specified. Command will delete all artifacts in the group")
		err = confirmDelete(opts, "Do you want to delete ALL ARTIFACTS from group "+opts.group)
		if err != nil {
			return err
		}
		request := dataAPI.ArtifactsApi.DeleteArtifactsInGroup(ctx, opts.group)
		_, err = request.Execute()
		if err != nil {
			return registryinstanceerror.TransformError(err)
		}
		logger.Info("Artifacts in group " + opts.group + " deleted")
	} else {
		_, _, err := dataAPI.MetadataApi.GetArtifactMetaData(ctx, opts.group, opts.artifact).Execute()
		if err != nil {
			return errors.New("artifact " + opts.artifact + " not found")
		}
		logger.Info("Deleting artifact " + opts.artifact)
		err = confirmDelete(opts, "Do you want to delete artifact "+opts.artifact+" from group "+opts.group)
		if err != nil {
			return err
		}
		request := dataAPI.ArtifactsApi.DeleteArtifact(ctx, opts.group, opts.artifact)

		_, err = request.Execute()
		if err != nil {
			return registryinstanceerror.TransformError(err)
		}
		logger.Info("Artifact deleted: " + opts.artifact)
	}

	return nil
}

func confirmDelete(opts *Options, message string) error {
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
