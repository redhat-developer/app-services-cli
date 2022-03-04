package create

import (
	"context"

	"github.com/redhat-developer/app-services-cli/pkg/cmd/kafka/flagutil"
	"github.com/redhat-developer/app-services-cli/pkg/core/ioutil/icon"
	"github.com/redhat-developer/app-services-cli/pkg/core/ioutil/iostreams"
	"github.com/redhat-developer/app-services-cli/pkg/core/localize"
	"github.com/redhat-developer/app-services-cli/pkg/core/logging"
	"github.com/redhat-developer/app-services-cli/pkg/core/profile"
	"github.com/redhat-developer/app-services-cli/pkg/shared/factory"
	"github.com/redhat-developer/app-services-cli/pkg/shared/profileutil"
	"github.com/spf13/cobra"
)

type options struct {
	IO         *iostreams.IOStreams
	Logger     logging.Logger
	Connection factory.ConnectionFunc
	localizer  localize.Localizer
	Context    context.Context
	Profiles   profile.IContext

	name       string
	kafkaID    string
	registryID string
}

// NewCreateCommand creates a new command to create contexts
func NewCreateCommand(f *factory.Factory) *cobra.Command {

	opts := &options{
		Connection: f.Connection,
		IO:         f.IOStreams,
		Logger:     f.Logger,
		localizer:  f.Localizer,
		Profiles:   f.Profile,
	}

	cmd := &cobra.Command{
		Use:     "create",
		Short:   f.Localizer.MustLocalize("context.create.cmd.shortDescription"),
		Long:    f.Localizer.MustLocalize("context.create.cmd.longDescription"),
		Example: f.Localizer.MustLocalize("context.create.cmd.example"),
		Args:    cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return runCreate(opts)
		},
	}

	flags := flagutil.NewFlagSet(cmd, opts.localizer)

	flags.StringVar(&opts.name, "name", "", opts.localizer.MustLocalize("context.common.flag.name"))
	flags.StringVar(&opts.kafkaID, "kafka-id", "", opts.localizer.MustLocalize("context.common.flag.name"))
	flags.StringVar(&opts.registryID, "registry-id", "", opts.localizer.MustLocalize("context.common.flag.name"))

	return cmd

}

func runCreate(opts *options) error {

	context, err := opts.Profiles.Load()
	if err != nil {
		return err
	}

	profileHandler := &profileutil.ContextHandler{
		Context:   context,
		Localizer: opts.localizer,
	}

	profiles := context.Contexts

	if profiles == nil {
		profiles = map[string]profile.ServiceConfig{}
	}

	currentCtx, _ := profileHandler.GetContext(opts.name)
	if currentCtx != nil {
		return opts.localizer.MustLocalizeError("context.create.log.alreadyExists", localize.NewEntry("Name", opts.name))
	}

	services := profile.ServiceConfig{
		KafkaID:           opts.kafkaID,
		ServiceRegistryID: opts.registryID,
	}

	profiles[opts.name] = services

	context.Contexts = profiles

	err = opts.Profiles.Save(context)
	if err != nil {
		return err
	}

	opts.Logger.Info(icon.SuccessPrefix(), opts.localizer.MustLocalize("context.create.log.successMessage"))

	return nil
}
