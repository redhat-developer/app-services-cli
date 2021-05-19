package create

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"github.com/redhat-developer/app-services-cli/pkg/connection"
	"github.com/redhat-developer/app-services-cli/pkg/localize"
	"github.com/redhat-developer/app-services-cli/pkg/serviceregistry"

	"github.com/redhat-developer/app-services-cli/pkg/cmd/flag"

	"github.com/redhat-developer/app-services-cli/pkg/dump"
	"github.com/redhat-developer/app-services-cli/pkg/iostreams"
	"gopkg.in/yaml.v2"

	"github.com/redhat-developer/app-services-cli/internal/config"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/factory"
	"github.com/redhat-developer/app-services-cli/pkg/logging"

	"github.com/spf13/cobra"
)

type Options struct {
	artifactLocation string
	registryID       string
	interactive      bool
	outputFormat     string

	IO         *iostreams.IOStreams
	Config     config.IConfig
	Connection factory.ConnectionFunc
	Logger     func() (logging.Logger, error)
	localizer  localize.Localizer
}

// NewCreateCommand gets a new command for creating kafka topic.
func NewCreateCommand(f *factory.Factory) *cobra.Command {
	opts := &Options{
		Connection: f.Connection,
		Config:     f.Config,
		Logger:     f.Logger,
		IO:         f.IOStreams,
		localizer:  f.Localizer,
	}

	cmd := &cobra.Command{
		Use:     "create",
		Short:   "Create artifact",
		Long:    "",
		Example: "",
		Args:    cobra.RangeArgs(0, 1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			if !opts.IO.CanPrompt() && len(args) == 0 {
				return errors.New(opts.localizer.MustLocalize("argument.error.requiredWhenNonInteractive", localize.NewEntry("Argument", "name")))
			} else if len(args) == 0 {
				opts.interactive = true
			}

			if err = flag.ValidateOutput(opts.outputFormat); err != nil {
				return err
			}

			if !opts.interactive {
				opts.artifactLocation = args[0]
			}

			if opts.registryID != "" {
				return runCmd(opts)
			}

			cfg, err := opts.Config.Load()
			if err != nil {
				return err
			}

			if !cfg.HasServiceRegistry() {
				return fmt.Errorf("no service Registry selected. Use rhoas registry use to select your registry")
			}

			opts.registryID = fmt.Sprint(cfg.Services.ServiceRegistry.InstanceID)
			return runCmd(opts)
		},
	}

	cmd.Flags().StringVarP(&opts.outputFormat, "output", "o", "json", opts.localizer.MustLocalize("kafka.topic.common.flag.output.description"))

	return cmd
}

// nolint:funlen
func runCmd(opts *Options) error {

	if opts.interactive {
		// run the create command interactively
		err := runInteractivePrompt(opts)
		if err != nil {
			return err
		}
	}

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
	logger.Info("Opening file...")
	specifiedFile, err := os.Open(opts.artifactLocation) // For read access.
	if err != nil {
		return err
	}

	logger.Info("Validating content type supported by registry")
	logger.Info("Using default group")
	// TODO groupId of artifacts needs to be determined by flag or status?
	requestArtifact := dataAPI.CreateArtifact(ctx, "default")
	requestArtifact = requestArtifact.Body(specifiedFile)
	metadata, _, err := requestArtifact.Execute()
	if err != nil {
		return err
	}
	specifiedFile.Close()

	logger.Info("Artifact created")

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

func runInteractivePrompt(opts *Options) (err error) {
	return errors.New("Interactive prompt not implemented")
}
