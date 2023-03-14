package create

import (
	"context"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/registry/artifact/types"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/registry/artifact/util"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/registry/registrycmdutil"
	"github.com/redhat-developer/app-services-cli/pkg/core/cmdutil/flagutil"
	"github.com/redhat-developer/app-services-cli/pkg/core/ioutil/color"
	"github.com/redhat-developer/app-services-cli/pkg/core/ioutil/iostreams"
	"github.com/redhat-developer/app-services-cli/pkg/core/localize"
	"github.com/redhat-developer/app-services-cli/pkg/core/logging"
	"github.com/redhat-developer/app-services-cli/pkg/core/servicecontext"
	"github.com/redhat-developer/app-services-cli/pkg/shared/contextutil"
	"github.com/redhat-developer/app-services-cli/pkg/shared/factory"
	"github.com/redhat-developer/app-services-cli/pkg/shared/serviceregistryutil"
	registryinstanceclient "github.com/redhat-developer/app-services-sdk-core/app-services-sdk-go/registryinstance/apiv1internal/client"
	registrymgmtclient "github.com/redhat-developer/app-services-sdk-core/app-services-sdk-go/registrymgmt/apiv1/client"
	"io"
	"os"
	"strings"

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

	downloadOnServer    bool
	references          []string
	referenceSeparators string

	IO             *iostreams.IOStreams
	Connection     factory.ConnectionFunc
	Logger         logging.Logger
	localizer      localize.Localizer
	Context        context.Context
	ServiceContext servicecontext.IContext
}

func NewCreateCommand(f *factory.Factory) *cobra.Command {
	opts := &options{
		IO:             f.IOStreams,
		Connection:     f.Connection,
		Logger:         f.Logger,
		localizer:      f.Localizer,
		Context:        f.Context,
		ServiceContext: f.ServiceContext,
	}

	cmd := &cobra.Command{
		Use:     "create",
		Short:   f.Localizer.MustLocalize("artifact.cmd.create.description.short"),
		Long:    f.Localizer.MustLocalize("artifact.cmd.create.description.long"),
		Example: f.Localizer.MustLocalize("artifact.cmd.create.example"),
		Args:    cobra.RangeArgs(0, 1),
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) > 0 {
				opts.file = args[0]
			}

			if opts.registryID != "" {
				return runCreate(opts)
			}

			registryInstance, err := contextutil.GetCurrentRegistryInstance(f)
			if err != nil {
				return err
			}

			opts.registryID = registryInstance.GetId()
			return runCreate(opts)
		},
	}

	cmd.Flags().StringVarP(&opts.outputFormat, "output", "o", "json", opts.localizer.MustLocalize("artifact.common.message.output.formatNoTable"))
	cmd.Flags().StringVar(&opts.file, "file", "", opts.localizer.MustLocalize("artifact.common.file.location"))

	cmd.Flags().StringVar(&opts.artifact, "artifact-id", "", opts.localizer.MustLocalize("artifact.common.id"))
	cmd.Flags().StringVarP(&opts.group, "group", "g", registrycmdutil.DefaultArtifactGroup, opts.localizer.MustLocalize("artifact.common.group"))

	cmd.Flags().StringVar(&opts.version, "version", "", opts.localizer.MustLocalize("artifact.common.custom.version"))
	cmd.Flags().StringVar(&opts.name, "name", "", opts.localizer.MustLocalize("artifact.common.custom.name"))
	cmd.Flags().StringVar(&opts.description, "description", "", opts.localizer.MustLocalize("artifact.common.custom.description"))

	cmd.Flags().StringVarP(&opts.artifactType, "type", "t", "", opts.localizer.MustLocalize("artifact.common.type"))
	cmd.Flags().StringVar(&opts.registryID, "instance-id", "", opts.localizer.MustLocalize("registry.common.flag.instance.id"))

	cmd.Flags().BoolVar(&opts.downloadOnServer, "download-on-server", false, opts.localizer.MustLocalize("artifact.common.downloadOnServer"))

	cmd.Flags().StringArrayVarP(&opts.references, "reference", "r", []string{}, opts.localizer.MustLocalize("registry.common.flag.reference.gav"))
	cmd.Flags().StringVar(&opts.referenceSeparators, "reference-separators", "=:", opts.localizer.MustLocalize("registry.common.flag.reference.separators"))

	flagutil.EnableOutputFlagCompletion(cmd)

	return cmd
}

func runCreate(opts *options) error {
	format := util.OutputFormatFromString(opts.outputFormat)
	if format == util.UnknownOutputFormat || format == util.TableOutputFormat {
		return opts.localizer.MustLocalizeError("artifact.common.error.invalidOutputFormat")
	}
	separators := []rune(opts.referenceSeparators)
	if len(separators) != 2 || separators[0] == separators[1] {
		return opts.localizer.MustLocalizeError("artifact.cmd.create.error.invalidReferenceSeparator", localize.NewEntry("Separator", opts.referenceSeparators))
	}
	conn, err := opts.Connection()
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
	if opts.group == registrycmdutil.DefaultArtifactGroup {
		opts.Logger.Info(opts.localizer.MustLocalize("registry.artifact.common.message.no.group", localize.NewEntry("DefaultArtifactGroup", registrycmdutil.DefaultArtifactGroup)))
	}

	executeExtended := false
	var specifiedFile *os.File
	if opts.downloadOnServer {
		if !util.IsURL(opts.file) {
			return opts.localizer.MustLocalizeError("artifact.common.error.fileNotUrl", localize.NewEntry("FileName", opts.file))
		}
		executeExtended = true
	} else {
		specifiedFile, err = loadLocalFile(opts)
		if err != nil {
			return err
		}
	}
	request := dataAPI.ArtifactsApi.CreateArtifact(opts.Context, opts.group)
	if opts.artifactType != "" {
		artifactTypes, err2 := types.GetArtifactTypes(dataAPI, opts.Context)
		if err2 != nil {
			return err2
		}
		valid := false
		for _, v := range artifactTypes {
			if opts.artifactType == v {
				valid = true
				break
			}
		}
		if !valid {
			return opts.localizer.MustLocalizeError("artifact.cmd.create.error.invalidArtifactType", localize.NewEntry("AllowedTypes", strings.Join(artifactTypes, ", ")))
		}
		request = request.XRegistryArtifactType(opts.artifactType)
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
	metadata, err := executeRequest(executeExtended, opts, dataAPI, &request, specifiedFile)
	if err != nil {
		return registrycmdutil.TransformInstanceError(err)
	}
	return printCreateResult(opts, registry, metadata, format)
}

func printCreateResult(opts *options, registry *registrymgmtclient.Registry,
	metadata *registryinstanceclient.ArtifactMetaData, format util.OutputFormat) error {

	opts.Logger.Info(opts.localizer.MustLocalize("artifact.common.message.created"))
	artifactURL, ok := util.GetArtifactURL(registry, metadata)
	if ok {
		opts.Logger.Info(opts.localizer.MustLocalize("artifact.common.webURL", localize.NewEntry("URL", color.Info(artifactURL))))
	}
	return util.Dump(opts.IO.Out, format, metadata, nil)
}

func executeRequest(executeExtended bool, opts *options,
	dataAPI *registryinstanceclient.APIClient, requestP *registryinstanceclient.ApiCreateArtifactRequest,
	specifiedFile *os.File) (*registryinstanceclient.ArtifactMetaData, error) {
	request := *requestP
	if len(opts.references) > 0 {
		executeExtended = true
	}
	if executeExtended {
		// Content
		var content string
		if opts.downloadOnServer {
			content = opts.file
		} else {
			bytes, err := io.ReadAll(specifiedFile)
			if err != nil {
				return nil, err
			}
			content = string(bytes)
		}
		// References
		references, err := loadReferences(dataAPI, opts)
		if err != nil {
			return nil, err
		}
		body := registryinstanceclient.ContentCreateRequest{
			Content:    content,
			References: references,
		}
		bytes, err := body.MarshalJSON()
		if err != nil {
			return nil, err
		}
		file, err := util.GetFileFromBytes(bytes)
		if err != nil {
			return nil, err
		}
		request = request.
			ContentType("application/create.extended+json").
			Body(file)
	} else {
		request = request.
			ContentType("").
			Body(specifiedFile)
	}
	metadata, _, err := request.Execute()
	return &metadata, err
}

func loadLocalFile(opts *options) (*os.File, error) {
	var specifiedFile *os.File
	var err error
	if opts.file != "" {
		if util.IsURL(opts.file) {
			opts.Logger.Info(opts.localizer.MustLocalize("artifact.common.message.loading.file", localize.NewEntry("FileName", opts.file)))
			specifiedFile, err = util.GetContentFromFileURL(opts.Context, opts.file)
			if err != nil {
				return nil, err
			}
		} else {
			opts.Logger.Info(opts.localizer.MustLocalize("artifact.common.message.opening.file", localize.NewEntry("FileName", opts.file)))
			specifiedFile, err = os.Open(opts.file)
			if err != nil {
				return nil, err
			}
		}
	} else {
		opts.Logger.Info(opts.localizer.MustLocalize("common.message.reading.file"))
		specifiedFile, err = util.CreateFileFromStdin()
		if err != nil {
			return nil, err
		}
	}
	return specifiedFile, nil
}

func loadReferences(dataAPI *registryinstanceclient.APIClient, opts *options) ([]registryinstanceclient.ArtifactReference, error) {
	separators := []rune(opts.referenceSeparators)
	separatorMain := separators[0]
	separatorGAV := separators[1]
	result := make([]registryinstanceclient.ArtifactReference, len(opts.references))
	for i, v := range opts.references {
		ref := registryinstanceclient.ArtifactReference{}
		parts := strings.Split(v, string(separatorMain))
		if len(parts) != 2 {
			return nil, opts.localizer.MustLocalizeError("artifact.cmd.create.error.invalidReferenceFormatGAV", localize.NewEntry("Input", v))
		}
		ref.Name = parts[0]
		gavParts := strings.Split(parts[1], string(separatorGAV))
		if len(gavParts) == 3 {
			ref.GroupId = gavParts[0]
			ref.ArtifactId = gavParts[1]
			ref.Version = &gavParts[2]
			if ref.GroupId == "" {
				ref.GroupId = "default"
			}
			if *ref.Version == "" {
				metaData, _, err := dataAPI.MetadataApi.GetArtifactMetaData(opts.Context, ref.GroupId, ref.ArtifactId).Execute()
				if err != nil {
					return nil, registrycmdutil.TransformInstanceError(err)
				}
				*ref.Version = metaData.GetVersion()
			}
		} else {
			return nil, opts.localizer.MustLocalizeError("artifact.cmd.create.error.invalidReferenceFormatGAV", localize.NewEntry("Input", v))
		}
		result[i] = ref
	}
	return result, nil
}
