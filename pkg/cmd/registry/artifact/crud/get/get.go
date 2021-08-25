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
	Logger     logging.Logger
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
Command will fetch the latest artifact from the registry based on the artifact-id and group.

When --version is specified command will fetch the specific version of the artifact.
Get command will fetch artifacts based on --group and --artifact-id and --version.
For fetching artifacts using global identifiers please use "service-registry download" command
`,
		Example: `
## Get latest artifact with name "my-artifact" and print it out to standard out
rhoas service-registry artifact get --artifact-id=my-artifact

## Get latest artifact with name "my-artifact" from group "my-group" and save it to artifact.json file
rhoas service-registry artifact get --artifact-id=my-artifact --group=my-group --file-location=artifact.json

## Get latest artifact and pipe it to other command 
rhoas service-registry artifact get --artifact-id=my-artifact | grep -i 'user'

## Get artifact with custom version and print it out to standard out
rhoas service-registry artifact get --artifact-id=myartifact --version=4
`,
		Args: cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			if opts.artifact == "" {
				return errors.New("artifact id is required. Please specify artifact by using --artifact-id flag")
			}

			if opts.registryID != "" {
				return runGet(opts)
			}

			cfg, err := opts.Config.Load()
			if err != nil {
				return err
			}

			if !cfg.HasServiceRegistry() {
				return errors.New("no service registry selected. Please specify registry by using --instance-id flag")
			}

			opts.registryID = fmt.Sprint(cfg.Services.ServiceRegistry.InstanceID)
			return runGet(opts)
		},
	}

	cmd.Flags().StringVarP(&opts.artifact, "artifact-id", "a", "", "Id of the artifact")
	cmd.Flags().StringVarP(&opts.group, "group", "g", util.DefaultArtifactGroup, "Artifact group")
	cmd.Flags().StringVar(&opts.registryID, "instance-id", "", "Id of the registry to be used. By default uses currently selected registry")
	cmd.Flags().StringVar(&opts.outputFile, "file-location", "", "Location of the output file")
	cmd.Flags().StringVar(&opts.version, "version", "", "Version of the artifact")

	flagutil.EnableOutputFlagCompletion(cmd)

	return cmd
}

func runGet(opts *Options) error {
	conn, err := opts.Connection(connection.DefaultConfigRequireMasAuth)
	if err != nil {
		return err
	}

	dataAPI, _, err := conn.API().ServiceRegistryInstance(opts.registryID)
	if err != nil {
		return err
	}

	if opts.group == util.DefaultArtifactGroup {
		opts.Logger.Info("Group was not specified. Using", util.DefaultArtifactGroup, "artifacts group.")
		opts.group = util.DefaultArtifactGroup
	}

	ctx := context.Background()
	var dataFile *os.File
	if opts.version != "" {
		opts.Logger.Info("Fetching artifact with version: " + opts.version)
		request := dataAPI.VersionsApi.GetArtifactVersion(ctx, opts.group, opts.artifact, opts.version)
		dataFile, _, err = request.Execute()
	} else {
		opts.Logger.Info("Fetching latest artifact")
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

	opts.Logger.Info("Successfully fetched artifact")
	return nil
}
