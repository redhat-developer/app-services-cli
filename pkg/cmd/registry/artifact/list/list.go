package list

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/redhat-developer/app-services-cli/pkg/localize"
	"github.com/redhat-developer/app-services-cli/pkg/serviceregistry"

	"github.com/redhat-developer/app-services-cli/pkg/cmd/flag"
	"github.com/redhat-developer/app-services-cli/pkg/connection"

	"gopkg.in/yaml.v2"

	"github.com/redhat-developer/app-services-cli/internal/config"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/factory"
	"github.com/redhat-developer/app-services-cli/pkg/dump"
	"github.com/redhat-developer/app-services-cli/pkg/iostreams"
	"github.com/redhat-developer/app-services-cli/pkg/logging"
	"github.com/spf13/cobra"
)

type Options struct {
	Config     config.IConfig
	IO         *iostreams.IOStreams
	Connection factory.ConnectionFunc
	Logger     func() (logging.Logger, error)
	localizer  localize.Localizer

	outputFormat string
	registryID   string
}

// NewListCommand gets a new command for getting kafkas.
func NewListCommand(f *factory.Factory) *cobra.Command {
	opts := &Options{
		Config:     f.Config,
		Connection: f.Connection,
		Logger:     f.Logger,
		IO:         f.IOStreams,
		localizer:  f.Localizer,
	}

	cmd := &cobra.Command{
		Use:     "list",
		Short:   "List artifacts",
		Long:    "",
		Example: "",
		Args:    cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			if opts.outputFormat != "" {
				if err := flag.ValidateOutput(opts.outputFormat); err != nil {
					return err
				}
			}

			if opts.registryID != "" {
				return runCmd(opts)
			}

			cfg, err := opts.Config.Load()
			if err != nil {
				return err
			}

			if !cfg.HasServiceRegistry() {
				// TODO unify common messages
				return fmt.Errorf("No Service registry selected")
			}

			opts.registryID = fmt.Sprint(cfg.Services.ServiceRegistry.InstanceID)

			return runCmd(opts)
		},
	}

	cmd.Flags().StringVarP(&opts.outputFormat, "output", "o", "", "Output format of the list command")

	return cmd
}

func runCmd(opts *Options) error {

	conn, err := opts.Connection(connection.DefaultConfigRequireMasAuth)
	if err != nil {
		return err
	}

	logger, err := opts.Logger()
	if err != nil {
		return err
	}

	api := conn.API().ServiceRegistry()
	ctx := context.Background()
	remoteRegistry, _, err := serviceregistry.GetServiceRegistryByID(ctx, api, opts.registryID)

	if err != nil {
		return err
	}

	dataAPI := conn.API().ServiceRegistryData(remoteRegistry.RegistryUrl)
	if err != nil {
		return err
	}

	// TODO groupId of artifacts needs to be determined by flag or status?
	requestArtifactList := dataAPI.ListArtifactsInGroup(ctx, "default")
	metadata, _, err := requestArtifactList.Execute()
	if err != nil {
		return err
	}

	logger.Info("List of artifacts for 'default' group:")

	// TODO support table by default
	if opts.outputFormat == "" {
		opts.outputFormat = "json"
	}

	switch opts.outputFormat {
	case "json":
		data, _ := json.Marshal(metadata)
		_ = dump.JSON(opts.IO.Out, data)
	case "yaml", "yml":
		data, _ := yaml.Marshal(metadata)
		_ = dump.YAML(opts.IO.Out, data)
	}

	return nil
}
