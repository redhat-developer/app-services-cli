package download

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

var unusedFlagIdValue int64 = -1

type Options struct {
	group string

	contentId  int64
	globalId   int64
	hash       string
	outputFile string

	registryID string

	IO         *iostreams.IOStreams
	Config     config.IConfig
	Logger     func() (logging.Logger, error)
	Connection factory.ConnectionFunc
	localizer  localize.Localizer
}

// NewDownloadCommand creates a new command for downloading binary content for registry artifacts.
func NewDownloadCommand(f *factory.Factory) *cobra.Command {
	opts := &Options{
		Config:     f.Config,
		Connection: f.Connection,
		IO:         f.IOStreams,
		localizer:  f.Localizer,
		Logger:     f.Logger,
	}

	cmd := &cobra.Command{
		Use:   "download",
		Short: "Download artifacts from registry by using global identifiers",
		Long: `Get single or more artifacts by group, content, hash or globalIds. 
		NOTE: Use "service-registry get" command if you wish to download artifact by artifactId.

		Flags are used to specify the artifact to download:

		--contentId - id if the content from metadata
		--globalId - globalId of the content from metadata
		--hash - SHA-256 hash of the content`,
		Example: `
## Get latest artifact by content id
rhoas service-registry artifact download --content-id=183282932983

## Get latest artifact by content id to specific file
rhoas service-registry artifact download --content-id=183282932983 schema.json

## Get latest artifact by global id
rhoas service-registry artifact download --global-id=383282932983

## Get latest artifact by hash
rhoas service-registry artifact download --hash=c71d239df91726fc519c6eb72d318ec65820627232b2f796219e87dcf35d0ab4
`,
		Args: cobra.RangeArgs(0, 1),
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) > 0 {
				opts.outputFile = args[0]
			}

			if opts.registryID != "" {
				return runGet(opts)
			}

			cfg, err := opts.Config.Load()
			if err != nil {
				return err
			}

			if !cfg.HasServiceRegistry() {
				return errors.New("no service Registry selected. Use 'rhoas service-registry use' to select your registry")
			}

			opts.registryID = fmt.Sprint(cfg.Services.ServiceRegistry.InstanceID)
			return runGet(opts)
		},
	}

	cmd.Flags().StringVarP(&opts.group, "group", "g", "", "Group of the artifact")
	cmd.Flags().StringVarP(&opts.hash, "hash", "", "", "SHA-256 hash of the artifact")
	cmd.Flags().Int64VarP(&opts.globalId, "global-id", "", unusedFlagIdValue, "Global ID of the artifact")
	cmd.Flags().Int64VarP(&opts.contentId, "content-id", "", unusedFlagIdValue, "ContentId of the artifact")

	cmd.Flags().StringVarP(&opts.outputFile, "output-file", "", "", "Filename of the output file")
	cmd.Flags().StringVarP(&opts.registryID, "instance-id", "", "", "Id of the registry to be used. By default uses currently selected registry")

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
		logger.Info("Group was not specified. Using 'default' artifacts group.")
		opts.group = util.DefaultArtifactGroup
	}

	logger.Info("Fetching artifact")

	ctx := context.Background()
	var dataFile *os.File
	// nolint
	if opts.contentId != -1 {
		request := dataAPI.ArtifactsApi.GetContentById(ctx, opts.contentId)
		dataFile, _, err = request.Execute()
	} else if opts.globalId != -1 {
		request := dataAPI.ArtifactsApi.GetContentByGlobalId(ctx, opts.globalId)
		dataFile, _, err = request.Execute()
	} else if opts.hash != "" {
		request := dataAPI.ArtifactsApi.GetContentByHash(ctx, opts.hash)
		dataFile, _, err = request.Execute()
	} else {
		return errors.New("please specify at least one flag: [contentId, global-id, hash]")
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
