package create

import (
	"context"
	"errors"

	"github.com/redhat-developer/app-services-cli/pkg/cmd/kafka/flagutil"
	"github.com/redhat-developer/app-services-cli/pkg/core/config"
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
	Config     config.IConfig
	Logger     logging.Logger
	Connection factory.ConnectionFunc
	localizer  localize.Localizer
	Context    context.Context
	Profiles   profile.IContext

	name       string
	kafkaID    string
	registryID string
}

func NewCreateCommand(f *factory.Factory) *cobra.Command {

	opts := &options{
		Config:     f.Config,
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

	currentCtx, _ := profileHandler.GetContext(context.CurrentContext)
	if currentCtx != nil {
		return errors.New("context already exists")
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

	return nil
}
