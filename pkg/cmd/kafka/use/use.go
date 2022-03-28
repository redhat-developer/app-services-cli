package use

import (
	"context"
	"fmt"

	"github.com/redhat-developer/app-services-cli/pkg/cmd/kafka/flagutil"

	"github.com/redhat-developer/app-services-cli/pkg/core/ioutil/icon"
	"github.com/redhat-developer/app-services-cli/pkg/core/ioutil/iostreams"
	"github.com/redhat-developer/app-services-cli/pkg/core/localize"
	"github.com/redhat-developer/app-services-cli/pkg/core/logging"
	"github.com/redhat-developer/app-services-cli/pkg/core/servicecontext"
	"github.com/redhat-developer/app-services-cli/pkg/shared/connection"
	"github.com/redhat-developer/app-services-cli/pkg/shared/contextutil"
	"github.com/redhat-developer/app-services-cli/pkg/shared/factory"
	"github.com/redhat-developer/app-services-cli/pkg/shared/kafkautil"
	kafkamgmtclient "github.com/redhat-developer/app-services-sdk-go/kafkamgmt/apiv1/client"

	"github.com/spf13/cobra"
)

type options struct {
	id          string
	name        string
	interactive bool

	IO             *iostreams.IOStreams
	Connection     factory.ConnectionFunc
	Logger         logging.Logger
	localizer      localize.Localizer
	Context        context.Context
	ServiceContext servicecontext.IContext
}

func NewUseCommand(f *factory.Factory) *cobra.Command {
	opts := &options{
		Connection:     f.Connection,
		Logger:         f.Logger,
		IO:             f.IOStreams,
		localizer:      f.Localizer,
		Context:        f.Context,
		ServiceContext: f.ServiceContext,
	}

	cmd := &cobra.Command{
		Use:     "use",
		Short:   opts.localizer.MustLocalize("kafka.use.cmd.shortDescription"),
		Long:    opts.localizer.MustLocalize("kafka.use.cmd.longDescription"),
		Example: opts.localizer.MustLocalize("kafka.use.cmd.example"),
		Args:    cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			if opts.id == "" && opts.name == "" {
				if !opts.IO.CanPrompt() {
					return opts.localizer.MustLocalizeError("kafka.use.error.idOrNameRequired")
				}
				opts.interactive = true
			}

			if opts.name != "" && opts.id != "" {
				return opts.localizer.MustLocalizeError("service.error.idAndNameCannotBeUsed")
			}

			return runUse(opts)
		},
	}

	flags := flagutil.NewFlagSet(cmd, opts.localizer)

	flags.StringVar(&opts.id, "id", "", opts.localizer.MustLocalize("kafka.use.flag.id"))
	flags.StringVar(&opts.name, "name", "", opts.localizer.MustLocalize("kafka.use.flag.name"))

	if err := kafkautil.RegisterNameFlagCompletionFunc(cmd, f); err != nil {
		opts.Logger.Debug(opts.localizer.MustLocalize("kafka.common.error.load.completions.name.flag"), err)
	}

	return cmd
}

func runUse(opts *options) error {
	if opts.interactive {
		// run the use command interactively
		err := runInteractivePrompt(opts)
		if err != nil {
			return err
		}
		// no Kafka was selected, exit program
		if opts.name == "" {
			return nil
		}
	}

	svcContext, err := opts.ServiceContext.Load()
	if err != nil {
		return err
	}

	profileHandler := &contextutil.ContextHandler{
		Context:   svcContext,
		Localizer: opts.localizer,
	}

	currCtx, err := profileHandler.GetCurrentContext()
	if err != nil {
		return err
	}

	svcConfig, err := profileHandler.GetContext(currCtx)
	if err != nil {
		return err
	}

	conn, err := opts.Connection(connection.DefaultConfigSkipMasAuth)
	if err != nil {
		return err
	}

	api := conn.API()

	var res *kafkamgmtclient.KafkaRequest
	if opts.name != "" {
		res, _, err = kafkautil.GetKafkaByName(opts.Context, api.KafkaMgmt(), opts.name)
		if err != nil {
			return err
		}
	} else {
		res, _, err = kafkautil.GetKafkaByID(opts.Context, api.KafkaMgmt(), opts.id)
		if err != nil {
			return err
		}
	}

	nameTmplEntry := localize.NewEntry("Name", res.GetName())
	svcConfig.KafkaID = res.GetId()
	svcContext.Contexts[svcContext.CurrentContext] = *svcConfig

	if err := opts.ServiceContext.Save(svcContext); err != nil {
		saveErrMsg := opts.localizer.MustLocalize("kafka.use.error.saveError", nameTmplEntry)
		return fmt.Errorf("%v: %w", saveErrMsg, err)
	}

	opts.Logger.Info(icon.SuccessPrefix(), opts.localizer.MustLocalize("kafka.use.log.info.useSuccess", nameTmplEntry))

	return nil
}

func runInteractivePrompt(opts *options) error {
	conn, err := opts.Connection(connection.DefaultConfigSkipMasAuth)
	if err != nil {
		return err
	}

	opts.Logger.Debug(opts.localizer.MustLocalize("common.log.debug.startingInteractivePrompt"))

	selectedKafka, err := kafkautil.InteractiveSelect(opts.Context, conn, opts.Logger, opts.localizer)
	if err != nil {
		return err
	}

	opts.name = selectedKafka.GetName()

	return nil
}
