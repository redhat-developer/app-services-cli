package get

import (
	"context"
	"fmt"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/registry/artifact/util"
	registryinstanceclient "github.com/redhat-developer/app-services-sdk-core/app-services-sdk-go/registryinstance/apiv1internal/client"
	"os"

	"github.com/redhat-developer/app-services-cli/pkg/cmd/registry/registrycmdutil"
	"github.com/redhat-developer/app-services-cli/pkg/core/cmdutil/flagutil"
	"github.com/redhat-developer/app-services-cli/pkg/core/ioutil/icon"
	"github.com/redhat-developer/app-services-cli/pkg/core/ioutil/iostreams"
	"github.com/redhat-developer/app-services-cli/pkg/core/localize"
	"github.com/redhat-developer/app-services-cli/pkg/core/logging"
	"github.com/redhat-developer/app-services-cli/pkg/core/servicecontext"
	"github.com/redhat-developer/app-services-cli/pkg/shared/contextutil"
	"github.com/redhat-developer/app-services-cli/pkg/shared/factory"

	"github.com/spf13/cobra"
)

type options struct {
	artifact   string
	group      string
	outputFile string

	registryID string
	version    string
	references bool

	IO             *iostreams.IOStreams
	Logger         logging.Logger
	Connection     factory.ConnectionFunc
	localizer      localize.Localizer
	Context        context.Context
	ServiceContext servicecontext.IContext
}

func NewGetCommand(f *factory.Factory) *cobra.Command {
	opts := &options{
		Connection:     f.Connection,
		IO:             f.IOStreams,
		localizer:      f.Localizer,
		Logger:         f.Logger,
		Context:        f.Context,
		ServiceContext: f.ServiceContext,
	}

	cmd := &cobra.Command{
		Use:     "get",
		Short:   f.Localizer.MustLocalize("artifact.cmd.get.description.short"),
		Long:    f.Localizer.MustLocalize("artifact.cmd.get.description.long"),
		Example: f.Localizer.MustLocalize("artifact.cmd.get.example"),
		Args:    cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			if opts.artifact == "" {
				return f.Localizer.MustLocalizeError("artifact.common.message.artifactIdRequired")
			}

			if opts.registryID != "" {
				return runGet(opts)
			}

			registryInstance, err := contextutil.GetCurrentRegistryInstance(f)
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
	cmd.Flags().StringVar(&opts.outputFile, "output-file", "", opts.localizer.MustLocalize("artifact.common.message.file.location"))
	cmd.Flags().StringVar(&opts.version, "version", "", opts.localizer.MustLocalize("artifact.common.version"))
	cmd.Flags().BoolVar(&opts.references, "references", false, opts.localizer.MustLocalize("artifact.cmd.get.references"))

	flagutil.EnableOutputFlagCompletion(cmd)

	return cmd
}

func runGet(opts *options) error {
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

	if !opts.references {

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
			return registrycmdutil.TransformInstanceError(err)
		}
		fileContent, err := os.ReadFile(dataFile.Name())
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

	} else {
		version, err := versionOrLatest(dataAPI, opts)
		if err != nil {
			return err
		}
		result, _, err := dataAPI.VersionsApi.GetArtifactVersionReferences(opts.Context, opts.group, opts.artifact, *version).Execute()
		if err != nil {
			return registrycmdutil.TransformInstanceError(err)
		}
		util.PrettyPrintReferences(os.Stdout, result)
	}

	return nil
}

func versionOrLatest(client *registryinstanceclient.APIClient, opts *options) (*string, error) {
	version := opts.version
	if version == "" {
		metaData, _, err := client.MetadataApi.GetArtifactMetaData(opts.Context, opts.group, opts.artifact).Execute()
		if err != nil {
			return nil, registrycmdutil.TransformInstanceError(err)
		}
		version = metaData.GetVersion()
		opts.Logger.Info(opts.localizer.MustLocalize("registry.common.message.version.usingLatest", localize.NewEntry("Version", version)))
	}
	return &version, nil
}
