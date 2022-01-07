package download

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/redhat-developer/app-services-cli/pkg/cmd/registry/artifact/util"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/registry/sdk"
	"github.com/redhat-developer/app-services-cli/pkg/core/cmdutil/factory"
	"github.com/redhat-developer/app-services-cli/pkg/core/cmdutil/flagutil"
	"github.com/redhat-developer/app-services-cli/pkg/core/config"
	"github.com/redhat-developer/app-services-cli/pkg/core/connection"
	"github.com/redhat-developer/app-services-cli/pkg/core/ioutil/icon"
	"github.com/redhat-developer/app-services-cli/pkg/core/ioutil/iostreams"
	"github.com/redhat-developer/app-services-cli/pkg/core/localize"
	"github.com/redhat-developer/app-services-cli/pkg/core/logging"

	"github.com/spf13/cobra"
)

var unusedFlagIdValue int64 = -1

type options struct {
	group string

	contentId  int64
	globalId   int64
	hash       string
	outputFile string

	registryID string

	IO         *iostreams.IOStreams
	Config     config.IConfig
	Logger     logging.Logger
	Connection factory.ConnectionFunc
	localizer  localize.Localizer
	Context    context.Context
}

// NewDownloadCommand creates a new command for downloading binary content for registry artifacts.
func NewDownloadCommand(f *factory.Factory) *cobra.Command {
	opts := &options{
		Config:     f.Config,
		Connection: f.Connection,
		IO:         f.IOStreams,
		localizer:  f.Localizer,
		Logger:     f.Logger,
		Context:    f.Context,
	}

	cmd := &cobra.Command{
		Use:     "download",
		Short:   f.Localizer.MustLocalize("artifact.cmd.download.description.short"),
		Long:    f.Localizer.MustLocalize("artifact.cmd.download.description.long"),
		Example: f.Localizer.MustLocalize("artifact.cmd.download.example"),
		Args:    cobra.RangeArgs(0, 1),
		RunE: func(cmd *cobra.Command, args []string) error {
			if opts.registryID != "" {
				return runGet(opts)
			}

			cfg, err := opts.Config.Load()
			if err != nil {
				return err
			}

			instanceID, ok := cfg.GetServiceRegistryIdOk()
			if !ok {
				return opts.localizer.MustLocalizeError("registry.no.service.selected.use.instance.id.flag")
			}

			opts.registryID = instanceID
			return runGet(opts)
		},
	}

	cmd.Flags().StringVarP(&opts.group, "group", "g", util.DefaultArtifactGroup, opts.localizer.MustLocalize("artifact.common.group"))
	cmd.Flags().StringVar(&opts.hash, "hash", "", opts.localizer.MustLocalize("artifact.common.sha"))
	cmd.Flags().Int64VarP(&opts.globalId, "global-id", "", unusedFlagIdValue, opts.localizer.MustLocalize("artifact.common.global.id"))
	cmd.Flags().Int64VarP(&opts.contentId, "content-id", "", unusedFlagIdValue, opts.localizer.MustLocalize("artifact.common.content.id"))

	cmd.Flags().StringVarP(&opts.outputFile, "output-file", "", "", opts.localizer.MustLocalize("artifact.common.message.file.location"))
	cmd.Flags().StringVar(&opts.registryID, "instance-id", "", opts.localizer.MustLocalize("artifact.common.registryIdToUse"))

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

	opts.Logger.Info(opts.localizer.MustLocalize("artifact.common.message.fetching.artifact"))

	var dataFile *os.File
	// nolint
	if opts.contentId != unusedFlagIdValue {
		request := dataAPI.ArtifactsApi.GetContentById(opts.Context, opts.contentId)
		dataFile, _, err = request.Execute()
	} else if opts.globalId != unusedFlagIdValue {
		request := dataAPI.ArtifactsApi.GetContentByGlobalId(opts.Context, opts.globalId)
		dataFile, _, err = request.Execute()
	} else if opts.hash != "" {
		request := dataAPI.ArtifactsApi.GetContentByHash(opts.Context, opts.hash)
		dataFile, _, err = request.Execute()
	} else {
		return opts.localizer.MustLocalizeError("artifact.cmd.common.error.specify.contentId.globalId.hash")
	}

	if err != nil {
		return sdk.TransformInstanceError(err)
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
