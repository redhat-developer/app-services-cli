package migrate

import (
	"context"
	"os"

	"github.com/redhat-developer/app-services-cli/pkg/cmd/registry/registrycmdutil"
	"github.com/redhat-developer/app-services-cli/pkg/core/config"
	"github.com/redhat-developer/app-services-cli/pkg/core/ioutil/iostreams"
	"github.com/redhat-developer/app-services-cli/pkg/core/localize"
	"github.com/redhat-developer/app-services-cli/pkg/core/logging"
	"github.com/redhat-developer/app-services-cli/pkg/shared/connection"
	"github.com/redhat-developer/app-services-cli/pkg/shared/factory"

	"github.com/spf13/cobra"
)

type ImportOptions struct {
	file       string
	registryID string

	IO         *iostreams.IOStreams
	Config     config.IConfig
	Connection factory.ConnectionFunc
	Logger     logging.Logger
	localizer  localize.Localizer
	Context    context.Context
}

func NewImportCommand(f *factory.Factory) *cobra.Command {
	opts := &ImportOptions{
		IO:         f.IOStreams,
		Config:     f.Config,
		Connection: f.Connection,
		Logger:     f.Logger,
		localizer:  f.Localizer,
		Context:    f.Context,
	}

	cmd := &cobra.Command{
		Use:     "import",
		Short:   f.Localizer.MustLocalize("artifact.cmd.import.description.short"),
		Long:    f.Localizer.MustLocalize("artifact.cmd.import.description.long"),
		Example: f.Localizer.MustLocalize("artifact.cmd.import.example"),
		Args:    cobra.RangeArgs(0, 1),
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) > 0 {
				opts.file = args[0]
			}

			if opts.registryID != "" {
				return runImport(opts)
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
			return runImport(opts)
		},
	}
	cmd.Flags().StringVar(&opts.file, "file", "", opts.localizer.MustLocalize("artifact.common.file.location"))
	cmd.Flags().StringVar(&opts.registryID, "instance-id", "", opts.localizer.MustLocalize("artifact.common.instance.id"))

	return cmd
}

func runImport(opts *ImportOptions) error {
	conn, err := opts.Connection(connection.DefaultConfigRequireMasAuth)
	if err != nil {
		return err
	}

	dataAPI, _, err := conn.API().ServiceRegistryInstance(opts.registryID)
	if err != nil {
		return err
	}

	opts.Logger.Info(opts.localizer.MustLocalize("artifact.common.message.opening.file", localize.NewEntry("FileName", opts.file)))
	specifiedFile, err := os.Open(opts.file)
	if err != nil {
		return err
	}

	request := dataAPI.AdminApi.ImportData(opts.Context)
	_, err = request.Body(specifiedFile).Execute()
	if err != nil {
		return registrycmdutil.TransformInstanceError(err)
	}

	opts.Logger.Info(opts.localizer.MustLocalize("artifact.import.success"))

	return nil
}
