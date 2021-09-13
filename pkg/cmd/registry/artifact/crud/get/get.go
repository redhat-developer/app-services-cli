package get

import (
	"context"
	"errors"
	"fmt"
	"github.com/redhat-developer/app-services-cli/pkg/icon"
	"io/ioutil"
	"os"

	flagutil "github.com/redhat-developer/app-services-cli/pkg/cmdutil/flags"
	"github.com/redhat-developer/app-services-cli/pkg/connection"
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
	artifact   string
	group      string
	outputFile string

	registryID string
	version    string

	IO         *iostreams.IOStreams
	Config     config.IConfig
	Logger     logging.Logger
	Connection factory.ConnectionFunc
	localizer  localize.Localizer
	Context    context.Context
}

func NewGetCommand(f *factory.Factory) *cobra.Command {
	opts := &options{
		Config:     f.Config,
		Connection: f.Connection,
		IO:         f.IOStreams,
		localizer:  f.Localizer,
		Logger:     f.Logger,
		Context:    f.Context,
	}

	cmd := &cobra.Command{
		Use:     "get",
		Short:   f.Localizer.MustLocalize("artifact.cmd.get.description.short"),
		Long:    f.Localizer.MustLocalize("artifact.cmd.get.description.long"),
		Example: f.Localizer.MustLocalize("artifact.cmd.get.example"),
		Args:    cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			if opts.artifact == "" {
				return errors.New(f.Localizer.MustLocalize("artifact.common.message.artifactIdRequired"))
			}

			if opts.registryID != "" {
				return runGet(opts)
			}

			cfg, err := opts.Config.Load()
			if err != nil {
				return err
			}

			if !cfg.HasServiceRegistry() {
				return errors.New(opts.localizer.MustLocalize("registry.no.service.selected.use.instance.id.flag"))
			}

			opts.registryID = cfg.Services.ServiceRegistry.InstanceID
			return runGet(opts)
		},
	}

	cmd.Flags().StringVar(&opts.artifact, "artifact-id", "", opts.localizer.MustLocalize("artifact.common.id"))
	cmd.Flags().StringVarP(&opts.group, "group", "g", util.DefaultArtifactGroup, opts.localizer.MustLocalize("artifact.common.group"))
	cmd.Flags().StringVar(&opts.registryID, "instance-id", "", opts.localizer.MustLocalize("artifact.common.instance.id"))
	cmd.Flags().StringVar(&opts.outputFile, "output-file", "", opts.localizer.MustLocalize("artifact.common.message.file.location"))
	cmd.Flags().StringVar(&opts.version, "version", "", opts.localizer.MustLocalize("artifact.common.version"))

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

	if opts.group == util.DefaultArtifactGroup {
		opts.Logger.Info(opts.localizer.MustLocalize("artifact.common.message.no.group", localize.NewEntry("DefaultArtifactGroup", util.DefaultArtifactGroup)))
		opts.group = util.DefaultArtifactGroup
	}

	var dataFile *os.File
	if opts.version != "" {
		opts.Logger.Info(opts.localizer.MustLocalize("artifact.common.message.fetching.with.version", localize.NewEntry("Version", opts.version)))
		request := dataAPI.VersionsApi.GetArtifactVersion(opts.Context, opts.group, opts.artifact, opts.version)
		dataFile, _, err = request.Execute()
	} else {
		opts.Logger.Info(opts.localizer.MustLocalize("artifact.common.message.fetching.latest"))
		request := dataAPI.ArtifactsApi.GetLatestArtifact(opts.Context, opts.group, opts.artifact)
		dataFile, _, err = request.Execute()
	}
	if err != nil {
		return registryinstanceerror.TransformError(err)
	}
	fileContent, err := ioutil.ReadFile(dataFile.Name())
	if err != nil {
		return err
	}
	if opts.outputFile != "" {
		err := os.WriteFile(opts.outputFile, fileContent, 0600)
		if err != nil {
			return err
		}
	} else {
		// Print to stdout
		fmt.Fprintf(os.Stdout, "%v\n", string(fileContent))
	}

	opts.Logger.Info(icon.SuccessPrefix(), opts.localizer.MustLocalize("artifact.common.message.fetched.successfully"))
	return nil
}
