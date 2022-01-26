package create

import (
	"context"
	"os"

	"github.com/redhat-developer/app-services-cli/pkg/cmd/registry/artifact/util"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/registry/registrycmdutil"

	"github.com/redhat-developer/app-services-cli/pkg/core/cmdutil/factory"
	"github.com/redhat-developer/app-services-cli/pkg/core/cmdutil/flagutil"
	"github.com/redhat-developer/app-services-cli/pkg/core/config"
	"github.com/redhat-developer/app-services-cli/pkg/core/connection"
	"github.com/redhat-developer/app-services-cli/pkg/core/ioutil/color"
	"github.com/redhat-developer/app-services-cli/pkg/core/ioutil/dump"
	"github.com/redhat-developer/app-services-cli/pkg/core/ioutil/iostreams"
	"github.com/redhat-developer/app-services-cli/pkg/core/localize"
	"github.com/redhat-developer/app-services-cli/pkg/core/logging"
	"github.com/redhat-developer/app-services-cli/pkg/serviceregistryutil"
	registryinstanceclient "github.com/redhat-developer/app-services-sdk-go/registryinstance/apiv1internal/client"

	"github.com/spf13/cobra"
)

type options struct {
	artifact string
	group    string

	file         string
	artifactType string

	version     string
	name        string
	description string

	registryID   string
	outputFormat string

	IO         *iostreams.IOStreams
	Config     config.IConfig
	Connection factory.ConnectionFunc
	Logger     logging.Logger
	localizer  localize.Localizer
	Context    context.Context
}

func NewCreateCommand(f *factory.Factory) *cobra.Command {
	opts := &options{
		IO:         f.IOStreams,
		Config:     f.Config,
		Connection: f.Connection,
		Logger:     f.Logger,
		localizer:  f.Localizer,
		Context:    f.Context,
	}

	cmd := &cobra.Command{
		Use:     "create",
		Short:   f.Localizer.MustLocalize("artifact.cmd.create.description.short"),
		Long:    f.Localizer.MustLocalize("artifact.cmd.create.description.long"),
		Example: f.Localizer.MustLocalize("artifact.cmd.create.example"),
		Args:    cobra.RangeArgs(0, 1),
		RunE: func(cmd *cobra.Command, args []string) error {
			validOutputFormats := flagutil.ValidOutputFormats
			if opts.outputFormat != "" && !flagutil.IsValidInput(opts.outputFormat, validOutputFormats...) {
				return flagutil.InvalidValueError("output", opts.outputFormat, validOutputFormats...)
			}

			if len(args) > 0 {
				opts.file = args[0]
			}

			if opts.artifactType != "" {
				if _, err := registryinstanceclient.NewArtifactTypeFromValue(opts.artifactType); err != nil {
					return opts.localizer.MustLocalizeError("artifact.cmd.create.error.invalidArtifactType", localize.NewEntry("AllowedTypes", util.GetAllowedArtifactTypeEnumValuesAsString()))
				}
			}

			if opts.registryID != "" {
				return runCreate(opts)
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
			return runCreate(opts)
		},
	}

	cmd.Flags().StringVarP(&opts.outputFormat, "output", "o", "json", opts.localizer.MustLocalize("artifact.common.message.output.format"))
	cmd.Flags().StringVar(&opts.file, "file", "", opts.localizer.MustLocalize("artifact.common.file.location"))

	cmd.Flags().StringVar(&opts.artifact, "artifact-id", "", opts.localizer.MustLocalize("artifact.common.id"))
	cmd.Flags().StringVarP(&opts.group, "group", "g", util.DefaultArtifactGroup, opts.localizer.MustLocalize("artifact.common.group"))

	cmd.Flags().StringVar(&opts.version, "version", "", opts.localizer.MustLocalize("artifact.common.custom.version"))
	cmd.Flags().StringVar(&opts.name, "name", "", opts.localizer.MustLocalize("artifact.common.custom.name"))
	cmd.Flags().StringVar(&opts.description, "description", "", opts.localizer.MustLocalize("artifact.common.custom.description"))

	cmd.Flags().StringVarP(&opts.artifactType, "type", "t", "", opts.localizer.MustLocalize("artifact.common.type", localize.NewEntry("AllowedTypes", util.GetAllowedArtifactTypeEnumValuesAsString())))
	cmd.Flags().StringVar(&opts.registryID, "instance-id", "", opts.localizer.MustLocalize("artifact.common.instance.id"))

	flagutil.EnableOutputFlagCompletion(cmd)

	_ = cmd.RegisterFlagCompletionFunc("type", func(cmd *cobra.Command, _ []string, _ string) ([]string, cobra.ShellCompDirective) {
		return util.AllowedArtifactTypeEnumValues, cobra.ShellCompDirectiveNoSpace
	})

	return cmd
}

func runCreate(opts *options) error {
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

	var specifiedFile *os.File
	if opts.file != "" {
		if util.IsURL(opts.file) {
			opts.Logger.Info(opts.localizer.MustLocalize("artifact.common.message.loading.file", localize.NewEntry("FileName", opts.file)))
			specifiedFile, err = util.GetContentFromFileURL(opts.Context, opts.file)
			if err != nil {
				return err
			}
		} else {
			opts.Logger.Info(opts.localizer.MustLocalize("artifact.common.message.opening.file", localize.NewEntry("FileName", opts.file)))
			specifiedFile, err = os.Open(opts.file)
			if err != nil {
				return err
			}
		}
	} else {
		opts.Logger.Info(opts.localizer.MustLocalize("artifact.common.message.reading.file"))
		specifiedFile, err = util.CreateFileFromStdin()
		if err != nil {
			return err
		}
	}

	request := dataAPI.ArtifactsApi.CreateArtifact(opts.Context, opts.group)
	if opts.artifactType != "" {
		request = request.XRegistryArtifactType(registryinstanceclient.ArtifactType(opts.artifactType))
	}
	if opts.artifact != "" {
		request = request.XRegistryArtifactId(opts.artifact)
	}
	if opts.version != "" {
		request = request.XRegistryVersion(opts.version)
	}

	if opts.name != "" {
		request = request.XRegistryName(opts.name)
	}

	request = request.XRegistryDescription(opts.description)

	request = request.Body(specifiedFile)
	metadata, _, err := request.Execute()
	if err != nil {
		return registrycmdutil.TransformInstanceError(err)
	}
	opts.Logger.Info(opts.localizer.MustLocalize("artifact.common.message.created"))

	artifactURL, ok := util.GetArtifactURL(registry, &metadata)

	if ok {
		opts.Logger.Info(opts.localizer.MustLocalize("artifact.common.webURL", localize.NewEntry("URL", color.Info(artifactURL))))
	}

	return dump.Formatted(opts.IO.Out, opts.outputFormat, metadata)
}
