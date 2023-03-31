package migrate

import (
	"context"
	"io"
	"os"

	"github.com/redhat-developer/app-services-cli/pkg/cmd/registry/registrycmdutil"
	"github.com/redhat-developer/app-services-cli/pkg/core/ioutil/iostreams"
	"github.com/redhat-developer/app-services-cli/pkg/core/localize"
	"github.com/redhat-developer/app-services-cli/pkg/core/logging"
	"github.com/redhat-developer/app-services-cli/pkg/core/servicecontext"
	"github.com/redhat-developer/app-services-cli/pkg/shared/contextutil"
	"github.com/redhat-developer/app-services-cli/pkg/shared/factory"

	"github.com/spf13/cobra"
)

type ExportOptions struct {
	file       string
	registryID string

	IO             *iostreams.IOStreams
	Connection     factory.ConnectionFunc
	Logger         logging.Logger
	localizer      localize.Localizer
	Context        context.Context
	ServiceContext servicecontext.IContext
}

func NewExportCommand(f *factory.Factory) *cobra.Command {
	opts := &ExportOptions{
		IO:             f.IOStreams,
		Connection:     f.Connection,
		Logger:         f.Logger,
		localizer:      f.Localizer,
		Context:        f.Context,
		ServiceContext: f.ServiceContext,
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

			registryInstance, err := contextutil.GetCurrentRegistryInstance(f)
			if err != nil {
				return err
			}

			opts.registryID = registryInstance.GetId()
			return runExport(opts)
		},
	}
	cmd.Flags().StringVar(&opts.file, "output-file", "", opts.localizer.MustLocalize("artifact.common.file.location"))
	cmd.Flags().StringVar(&opts.registryID, "instance-id", "", opts.localizer.MustLocalize("registry.common.flag.instance.id"))
	_ = cmd.MarkFlagRequired("output-file")

	return cmd
}

func runExport(opts *ExportOptions) error {
	conn, err := opts.Connection()
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

	request := dataAPI.AdminApi.ExportData(opts.Context).ForBrowser(false)
	file, _, err := request.Execute()
	if err != nil {
		return registrycmdutil.TransformInstanceError(err)
	}
	_, err = io.Copy(fileContent, file)
	if err != nil {
		return err
	}
	opts.Logger.Info(opts.localizer.MustLocalize("artifact.export.success"))

	return nil
}
