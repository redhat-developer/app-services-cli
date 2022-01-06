package migrate

import (
	"context"
	"io"
	"os"

	"github.com/redhat-developer/app-services-cli/pkg/cmd/registry/registryutil"
	"github.com/redhat-developer/app-services-cli/pkg/core/cmdutil/factory"
	"github.com/redhat-developer/app-services-cli/pkg/core/config"
	"github.com/redhat-developer/app-services-cli/pkg/core/connection"
	"github.com/redhat-developer/app-services-cli/pkg/core/ioutil/iostreams"
	"github.com/redhat-developer/app-services-cli/pkg/core/localize"
	"github.com/redhat-developer/app-services-cli/pkg/core/logging"

	"github.com/spf13/cobra"
)

type ExportOptions struct {
	file       string
	registryID string

	IO         *iostreams.IOStreams
	Config     config.IConfig
	Connection factory.ConnectionFunc
	Logger     logging.Logger
	localizer  localize.Localizer
	Context    context.Context
}

func NewExportCommand(f *factory.Factory) *cobra.Command {
	opts := &ExportOptions{
		IO:         f.IOStreams,
		Config:     f.Config,
		Connection: f.Connection,
		Logger:     f.Logger,
		localizer:  f.Localizer,
		Context:    f.Context,
	}

	cmd := &cobra.Command{
		Use:     "export",
		Short:   f.Localizer.MustLocalize("artifact.cmd.export.description.short"),
		Long:    f.Localizer.MustLocalize("artifact.cmd.export.description.long"),
		Example: f.Localizer.MustLocalize("artifact.cmd.export.example"),
		Args:    cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			if opts.registryID != "" {
				return runExport(opts)
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
			return runExport(opts)
		},
	}
	cmd.Flags().StringVar(&opts.file, "output-file", "", opts.localizer.MustLocalize("artifact.common.file.location"))
	cmd.Flags().StringVar(&opts.registryID, "instance-id", "", opts.localizer.MustLocalize("artifact.common.instance.id"))
	_ = cmd.MarkFlagRequired("output-file")

	return cmd
}

func runExport(opts *ExportOptions) error {
	conn, err := opts.Connection(connection.DefaultConfigRequireMasAuth)
	if err != nil {
		return err
	}

	fileContent, err := os.Create(opts.file)
	if err != nil {
		return err
	}

	dataAPI, _, err := conn.API().ServiceRegistryInstance(opts.registryID)
	if err != nil {
		return err
	}

	request := dataAPI.AdminApi.ExportData(opts.Context)
	file, _, err := request.Execute()
	if err != nil {
		return registryutil.TransformInstanceError(err)
	}
	_, err = io.Copy(fileContent, file)
	if err != nil {
		return err
	}
	opts.Logger.Info(opts.localizer.MustLocalize("artifact.export.success"))

	return nil
}
