package create

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/redhat-developer/app-services-cli/pkg/connection"
	"github.com/redhat-developer/app-services-cli/pkg/dump"
	"github.com/redhat-developer/app-services-cli/pkg/localize"
	"github.com/redhat-developer/app-services-cli/pkg/serviceregistry/registryinstanceerror"
	registryinstanceclient "github.com/redhat-developer/app-services-sdk-go/registryinstance/apiv1internal/client"

	"github.com/redhat-developer/app-services-cli/pkg/cmd/flag"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/registry/artifact/util"
	flagutil "github.com/redhat-developer/app-services-cli/pkg/cmdutil/flags"

	"github.com/redhat-developer/app-services-cli/pkg/iostreams"

	"github.com/redhat-developer/app-services-cli/pkg/logging"

	"github.com/spf13/cobra"

	"github.com/redhat-developer/app-services-cli/internal/config"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/factory"
)

type Options struct {
	artifact string
	group    string

	file         string
	artifactType string

	registryID   string
	outputFormat string

	IO         *iostreams.IOStreams
	Config     config.IConfig
	Connection factory.ConnectionFunc
	Logger     func() (logging.Logger, error)
	localizer  localize.Localizer
}

var longDescription = `
Creates a new artifact by posting the artifact content to the registry.

Artifacts can be typically in JSON format for most of the supported types, but may be in another format for a few (for example, PROTOBUF).
The registry attempts to figure out what kind of artifact is being added from the following supported list:

- Avro (AVRO)
- Protobuf (PROTOBUF)
- JSON Schema (JSON)
- Kafka Connect (KCONNECT)
- OpenAPI (OPENAPI)
- AsyncAPI (ASYNCAPI)
- GraphQL (GRAPHQL)
- Web Services Description Language (WSDL)
- XML Schema (XSD)

An artifact is created using the content provided in the body of the request.  
This content is created under a unique artifact ID that can be provided by user.
If not provided in the request, the server generates a unique ID for the artifact. 
It is typically recommended that callers provide the ID, because this is a meaningful identifier, and for most use cases should be supplied by the caller.
If an artifact with the provided artifact ID already exists command will fail with error.


When --group parameter is missing the command will create a new artifact under the "default" group.
when --registry is missing the command will create a new artifact for currently active service registry (visible in rhoas service-registry describe)
`

func NewCreateCommand(f *factory.Factory) *cobra.Command {
	opts := &Options{
		IO:         f.IOStreams,
		Config:     f.Config,
		Connection: f.Connection,
		Logger:     f.Logger,
		localizer:  f.Localizer,
	}

	cmd := &cobra.Command{
		Use:   "create",
		Short: "Creates new artifact from file or standard input",
		Long:  longDescription,
		Example: `
# Create an artifact in default group
rhoas service-registry artifact create my-artifact.json

# Create an artifact with specified type
rhoas service-registry artifact create --type=JSON my-artifact.json
		`,
		Args: cobra.RangeArgs(0, 1),
		RunE: func(cmd *cobra.Command, args []string) error {
			validOutputFormats := flagutil.ValidOutputFormats
			if opts.outputFormat != "" && !flagutil.IsValidInput(opts.outputFormat, validOutputFormats...) {
				return flag.InvalidValueError("output", opts.outputFormat, validOutputFormats...)
			}

			if len(args) > 0 {
				opts.file = args[0]
			}

			if opts.registryID != "" {
				return runCreate(opts)
			}

			cfg, err := opts.Config.Load()
			if err != nil {
				return err
			}

			if opts.artifactType != "" {
				if _, err = registryinstanceclient.NewArtifactTypeFromValue(opts.artifactType); err != nil {
					return errors.New("Invalid artifact type. Please use one of following values: " + util.GetAllowedArtifactTypeEnumValuesAsString())
				}
			}

			if !cfg.HasServiceRegistry() {
				return errors.New("No service Registry selected. Use 'rhoas service-registry use' use to select your registry")
			}

			opts.registryID = fmt.Sprint(cfg.Services.ServiceRegistry.InstanceID)
			return runCreate(opts)
		},
	}

	cmd.Flags().StringVarP(&opts.outputFormat, "output", "o", "json", opts.localizer.MustLocalize("registry.cmd.flag.output.description"))
	cmd.Flags().StringVarP(&opts.file, "file", "f", "", "File location of the artifact")

	cmd.Flags().StringVarP(&opts.artifact, "artifact-id", "a", "", "Id of the artifact")
	cmd.Flags().StringVarP(&opts.group, "group", "g", "", "Group of the artifact")
	cmd.Flags().StringVarP(&opts.artifactType, "type", "t", "", "Type of artifact. Choose from:  "+util.GetAllowedArtifactTypeEnumValuesAsString())
	cmd.Flags().StringVarP(&opts.registryID, "instance-id", "", "", "Id of the registry to be used. By default uses currently selected registry")

	flagutil.EnableOutputFlagCompletion(cmd)
	_ = cmd.RegisterFlagCompletionFunc("type", func(cmd *cobra.Command, _ []string, _ string) ([]string, cobra.ShellCompDirective) {
		return util.AllowedArtifactTypeEnumValues, cobra.ShellCompDirectiveNoSpace
	})

	return cmd
}

func runCreate(opts *Options) error {

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

	if opts.group == "" {
		logger.Info("Group was not specified. Using " + util.DefaultArtifactGroup + " artifacts group.")
		opts.group = util.DefaultArtifactGroup
	}

	var specifiedFile *os.File
	if opts.file != "" {
		logger.Info("Opening file: " + opts.file)
		specifiedFile, err = os.Open(opts.file)
		if err != nil {
			return err
		}
	} else {
		logger.Info("Reading file content from stdin")
		specifiedFile, err = util.CreateFileFromStdin()
		if err != nil {
			return err
		}
	}

	ctx := context.Background()
	request := dataAPI.ArtifactsApi.CreateArtifact(ctx, opts.group)
	if opts.artifactType != "" {
		request = request.XRegistryArtifactType(registryinstanceclient.ArtifactType(opts.artifactType))
	}
	if opts.artifact != "" {
		request = request.XRegistryArtifactId(opts.artifact)
	}

	request = request.Body(specifiedFile)
	metadata, _, err := request.Execute()
	if err != nil {
		return registryinstanceerror.TransformError(err)
	}
	logger.Info("Artifact created")

	dump.PrintDataInFormat(opts.outputFormat, metadata, opts.IO.Out)

	return nil
}
