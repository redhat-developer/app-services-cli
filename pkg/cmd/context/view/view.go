package view

import (
	"context"
	"errors"

	"github.com/MakeNowJust/heredoc"

	"github.com/redhat-developer/app-services-cli/pkg/localize"

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

	fileFormat       string
	overwrite        bool
	shortDescription string
	filename         string

	interactive bool
}

func NewViewCommand(f *factory.Factory) *cobra.Command {
	opts := &options{
		IO:         f.IOStreams,
		Config:     f.Config,
		Connection: f.Connection,
		Logger:     f.Logger,
		localizer:  f.Localizer,
		Context:    f.Context,
	}

	cmd := &cobra.Command{
		Use:   "view",
		Short: "view configuration profiles",
		Long: heredoc.Doc(`
		view configuration profiles available for that 
		`),
		Example: "",
		Args:    cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) (err error) {
			return runCreate(opts)
		},
	}

	return cmd
}

// nolint:funlen
func runCreate(opts *options) error {
	cfg, err := opts.Config.Load()
	if err != nil {
		return err
	}

	if cfg.Services.Kafka.ClusterID == "" || cfg.Services.ServiceRegistry.InstanceID == "" {
		// For demo purpose
		return errors.New("No profile data available")
	}
	// TODO Hardcoded values because we missing data in current config
	// We need to store profile configuration using new config
	profilePlaceholder := heredoc.Doc(`
	{
	  "profiles": [{
			"name": "default",
			"services": {
				"kafka": {
					"id": "c694a89bbdnrh1nhmhj0",
					"name": "target",
					"bootstrapHostUrl": "target-1isy6rq3jki8q0otmjqfd3ocfrg.apps.mk-bttg0jn170hp.x5u8.s1.devshift.org",
				},
				"service-registry": {
					"id": "c4b2efb1-7360-4ef1-bc15-b5a5c13c93f7",
					"name": "test",
					"url": "https://registry.apps.example.com/t/5213600b-afc9-487e-8cc3-339f4248d706"
				}
			}
		}]
	}`)

	opts.Logger.Info(profilePlaceholder)
	return nil
}
