package manage

import (
	"context"

	"github.com/MakeNowJust/heredoc"

	"github.com/redhat-developer/app-services-cli/pkg/kafka"
	"github.com/redhat-developer/app-services-cli/pkg/localize"
	"github.com/redhat-developer/app-services-cli/pkg/serviceregistry"

	"github.com/redhat-developer/app-services-cli/pkg/connection"

	"github.com/redhat-developer/app-services-cli/pkg/iostreams"

	"github.com/redhat-developer/app-services-cli/internal/config"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/factory"
	"github.com/redhat-developer/app-services-cli/pkg/logging"
	"github.com/spf13/cobra"
)

type options struct {
	IO         *iostreams.IOStreams
	Config     config.IConfig
	Connection factory.ConnectionFunc
	Logger     logging.Logger
	localizer  localize.Localizer
	Context    context.Context

	profile string
}

func NewManageCommand(f *factory.Factory) *cobra.Command {
	opts := &options{
		IO:         f.IOStreams,
		Config:     f.Config,
		Connection: f.Connection,
		Logger:     f.Logger,
		localizer:  f.Localizer,
		Context:    f.Context,
	}

	cmd := &cobra.Command{
		Use:   "manage",
		Short: "manage configuration profiles. What services are enabled",
		Long: heredoc.Doc(`
		Manage command lets you select groups of services. 
		Selected services will be used in the RHOAS CLI to generate config and 
		to perform specific actions like creation of topics.
		`),
		Example: "",
		Args:    cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) (err error) {
			return runManage(opts)
		},
	}

	cmd.Flags().StringVar(&opts.profile, "profile", "default", "name of the profile to be use or create")

	return cmd
}

// nolint:funlen
func runManage(opts *options) error {
	conn, err := opts.Connection(connection.DefaultConfigSkipMasAuth)
	if err != nil {
		return err
	}
	cfg, err := opts.Config.Load()
	if err != nil {
		return err
	}

	opts.Logger.Info("Setting active services profile " + opts.profile)
	opts.Logger.Info("Select Kafka")
	selectedService, err := kafka.InteractiveSelect(opts.Context, conn, opts.Logger, opts.localizer)
	if err != nil {
		return err
	}
	if selectedService == nil {
		return nil
	}
	// TODO use external config
	cfg.Services.Kafka.ClusterID = selectedService.GetId()

	opts.Logger.Info("Select Registry service")
	selectedRegistry, err := serviceregistry.InteractiveSelect(opts.Context, conn, opts.Logger)
	if err != nil {
		return err
	}
	if selectedRegistry == nil {
		return nil
	}

	cfg.Services.ServiceRegistry.InstanceID = selectedRegistry.GetId()
	err = opts.Config.Save(cfg)
	if err != nil {
		return err
	}
	opts.Logger.Info("Profile updated. Run 'rhoas config status' to view your active services")
	opts.Logger.Info("")
	opts.Logger.Info("You can now generate your configuration based on your profile using")
	opts.Logger.Info("rhoas config generate")
	return nil
}
