package get

import (
	"context"
	"errors"
	"fmt"
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

type Options struct {
	artifact   string
	group      string
	outputFile string

	registryID string
	version    string

	IO         *iostreams.IOStreams
	Config     config.IConfig
	Logger     func() (logging.Logger, error)
	Connection factory.ConnectionFunc
	localizer  localize.Localizer
}

func NewGetCommand(f *factory.Factory) *cobra.Command {
	opts := &Options{
		Config:     f.Config,
		Connection: f.Connection,
		IO:         f.IOStreams,
		localizer:  f.Localizer,
		Logger:     f.Logger,
	}

	cmd := &cobra.Command{
		Use:   "get",
		Short: "Get artifact by id and group",
		Long: `Get artifact by specifying id and group.
Command will fetch the latest artifact from the registry based on the artifactId and group.

When --version is specified command will fetch the specific version of the artifact.
Get command will fetch artifacts based on group and artifactId and version.
For fetching artifacts using global identifiers please use "service-registry download" command
`,
		Example: `
## Get latest artifact by name
rhoas service-registry artifacts get myschema

## Get latest artifact and save its content to file
rhoas service-registry artifacts get myschema myschema.json

## Get latest artifact and pipe it to other command 
rhoas service-registry artifacts get myschema | grep -i 'user'

## Get latest artifact by specifying custom group, registry and name as flag
rhoas service-registry artifacts get --group mygroup --instance-id=myregistry --artifact myartifact
`,
		Args: cobra.RangeArgs(0, 2),
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) > 0 {
				opts.artifact = args[0]
			}

			if opts.artifact == "" {
				return errors.New("Artifact is required. Please specify artifact as positional argument or by using --artifact flag")
			}

			if len(args) > 1 {
				opts.outputFile = args[1]
			}

			if opts.registryID != "" {
				return runGet(opts)
			}

			cfg, err := opts.Config.Load()
			if err != nil {
				return err
			}

			if !cfg.HasServiceRegistry() {
				return errors.New("No service Registry selected. Use 'rhoas service-registry use' to select your registry")
			}

			opts.registryID = fmt.Sprint(cfg.Services.ServiceRegistry.InstanceID)
			return runGet(opts)
		},
	}

	cmd.Flags().StringVarP(&opts.artifact, "artifact", "a", "", "Id of the artifact")
	cmd.Flags().StringVarP(&opts.group, "group", "g", "", "Group of the artifact")
	cmd.Flags().StringVarP(&opts.registryID, "instance-id", "", "", "Id of the registry to be used. By default uses currently selected registry")
	cmd.Flags().StringVarP(&opts.outputFile, "output-file", "", "", "Filename of the output file")
	cmd.Flags().StringVarP(&opts.version, "version", "", "", "Version of the artifact")

	flagutil.EnableOutputFlagCompletion(cmd)

	return cmd
}

func runGet(opts *Options) error {
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

	ctx := context.Background()
	var dataFile *os.File
	if opts.version != "" {
		logger.Info("Fetching artifact with version: " + opts.version)
		request := dataAPI.VersionsApi.GetArtifactVersion(ctx, opts.group, opts.artifact, opts.version)
		dataFile, _, err = request.Execute()
	} else {
		logger.Info("Fetching latest artifact")
		request := dataAPI.ArtifactsApi.GetLatestArtifact(ctx, opts.group, opts.artifact)
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

	logger.Info("Successfully fetched artifact")
	return nil
}
