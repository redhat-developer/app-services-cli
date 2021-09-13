package update

import (
	"context"
	"errors"
	"os"

	"github.com/redhat-developer/app-services-cli/pkg/connection"
	"github.com/redhat-developer/app-services-cli/pkg/dump"
	"github.com/redhat-developer/app-services-cli/pkg/localize"

	"github.com/redhat-developer/app-services-cli/pkg/cmd/flag"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/registry/artifact/util"
	flagutil "github.com/redhat-developer/app-services-cli/pkg/cmdutil/flags"

	"github.com/redhat-developer/app-services-cli/pkg/iostreams"

	"github.com/redhat-developer/app-services-cli/pkg/logging"

	"github.com/spf13/cobra"

	"github.com/redhat-developer/app-services-cli/internal/config"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/factory"
)

type options struct {
	artifact string
	group    string

	file string

	registryID   string
	outputFormat string

	version     string
	name        string
	description string

	IO         *iostreams.IOStreams
	Config     config.IConfig
	Connection factory.ConnectionFunc
	Logger     logging.Logger
	localizer  localize.Localizer
	Context    context.Context
}

// NewUpdateCommand creates a new command for updating binary content of registry artifacts.
func NewUpdateCommand(f *factory.Factory) *cobra.Command {
	opts := &options{
		IO:         f.IOStreams,
		Config:     f.Config,
		Connection: f.Connection,
		Logger:     f.Logger,
		localizer:  f.Localizer,
		Context:    f.Context,
	}

	cmd := &cobra.Command{
		Use:     "update",
		Short:   f.Localizer.MustLocalize("artifact.cmd.update.description.short"),
		Long:    f.Localizer.MustLocalize("artifact.cmd.update.description.long"),
		Example: f.Localizer.MustLocalize("artifact.cmd.update.example"),
		Args:    cobra.RangeArgs(0, 1),
		RunE: func(cmd *cobra.Command, args []string) error {
			validOutputFormats := flagutil.ValidOutputFormats
			if opts.outputFormat != "" && !flagutil.IsValidInput(opts.outputFormat, validOutputFormats...) {
				return flag.InvalidValueError("output", opts.outputFormat, validOutputFormats...)
			}

			if opts.artifact == "" {
				return errors.New(opts.localizer.MustLocalize("artifact.common.error.artifact.id.required"))
			}

			if len(args) > 0 {
				opts.file = args[0]
			}

			if opts.registryID != "" {
				return runUpdate(opts)
			}

			cfg, err := opts.Config.Load()
			if err != nil {
				return err
			}

			if !cfg.HasServiceRegistry() {
				return errors.New(opts.localizer.MustLocalize("artifact.cmd.common.error.noServiceRegistrySelected"))
			}

			opts.registryID = cfg.Services.ServiceRegistry.InstanceID
			return runUpdate(opts)
		},
	}

	cmd.Flags().StringVarP(&opts.outputFormat, "output", "o", "json", opts.localizer.MustLocalize("registry.cmd.flag.output.description"))
	cmd.Flags().StringVarP(&opts.file, "file", "f", "", opts.localizer.MustLocalize("artifact.common.file.location"))

	cmd.Flags().StringVarP(&opts.artifact, "artifact-id", "", "", opts.localizer.MustLocalize("artifact.common.id"))
	cmd.Flags().StringVarP(&opts.group, "group", "g", util.DefaultArtifactGroup, opts.localizer.MustLocalize("artifact.common.group"))
	cmd.Flags().StringVar(&opts.registryID, "instance-id", "", opts.localizer.MustLocalize("artifact.common.instance.id"))

	cmd.Flags().StringVar(&opts.version, "version", "", opts.localizer.MustLocalize("artifact.common.custom.version"))
	cmd.Flags().StringVar(&opts.name, "name", "", opts.localizer.MustLocalize("artifact.common.custom.name"))
	cmd.Flags().StringVar(&opts.description, "description", "", opts.localizer.MustLocalize("artifact.common.custom.description"))

	flagutil.EnableOutputFlagCompletion(cmd)

	return cmd
}

func runUpdate(opts *options) error {
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

	var specifiedFile *os.File
	if opts.file != "" {
		opts.Logger.Info(opts.localizer.MustLocalize("artifact.common.message.opening.file", localize.NewEntry("FileName", opts.file)))
		specifiedFile, err = os.Open(opts.file)
		if err != nil {
			return err
		}
	} else {
		opts.Logger.Info(opts.localizer.MustLocalize("artifact.common.message.reading.file"))
		specifiedFile, err = util.CreateFileFromStdin()
		if err != nil {
			return err
		}
	}

	request := dataAPI.ArtifactsApi.UpdateArtifact(opts.Context, opts.group, opts.artifact)
	if opts.version != "" {
		request = request.XRegistryVersion(opts.version)
	}
	if opts.name != "" {
		request = request.XRegistryName(opts.name)
	}

	if opts.description != "" {
		request = request.XRegistryDescription(opts.description)
	}

	request = request.Body(specifiedFile)
	metadata, _, err := request.Execute()
	if err != nil {
		return err
	}
	opts.Logger.Info(opts.localizer.MustLocalize("artifact.common.message.updated"))

	dump.PrintDataInFormat(opts.outputFormat, metadata, opts.IO.Out)

	return nil
}
