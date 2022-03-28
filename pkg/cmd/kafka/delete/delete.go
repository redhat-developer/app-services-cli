package delete

import (
	"context"
	"fmt"

	"github.com/redhat-developer/app-services-cli/pkg/core/cmdutil/flagutil"
	"github.com/redhat-developer/app-services-cli/pkg/core/servicecontext"

	"github.com/redhat-developer/app-services-cli/pkg/core/ioutil/iostreams"
	"github.com/redhat-developer/app-services-cli/pkg/core/localize"
	"github.com/redhat-developer/app-services-cli/pkg/core/logging"
	"github.com/redhat-developer/app-services-cli/pkg/shared/connection"
	"github.com/redhat-developer/app-services-cli/pkg/shared/contextutil"
	"github.com/redhat-developer/app-services-cli/pkg/shared/factory"
	"github.com/redhat-developer/app-services-cli/pkg/shared/kafkautil"
	kafkamgmtclient "github.com/redhat-developer/app-services-sdk-go/kafkamgmt/apiv1/client"

	"github.com/AlecAivazis/survey/v2"
	"github.com/spf13/cobra"
)

type options struct {
	id          string
	name        string
	skipConfirm bool

	IO             *iostreams.IOStreams
	Connection     factory.ConnectionFunc
	Logger         logging.Logger
	localizer      localize.Localizer
	Context        context.Context
	ServiceContext servicecontext.IContext
}

// NewDeleteCommand command for deleting kafkas.
func NewDeleteCommand(f *factory.Factory) *cobra.Command {
	opts := &options{
		Connection:     f.Connection,
		Logger:         f.Logger,
		IO:             f.IOStreams,
		localizer:      f.Localizer,
		Context:        f.Context,
		ServiceContext: f.ServiceContext,
	}

	cmd := &cobra.Command{
		Use:     "delete",
		Short:   opts.localizer.MustLocalize("kafka.delete.cmd.shortDescription"),
		Long:    opts.localizer.MustLocalize("kafka.delete.cmd.longDescription"),
		Example: opts.localizer.MustLocalize("kafka.delete.cmd.example"),
		Args:    cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			if !opts.IO.CanPrompt() && !opts.skipConfirm {
				return flagutil.RequiredWhenNonInteractiveError("yes")
			}

			if opts.name != "" && opts.id != "" {
				return opts.localizer.MustLocalizeError("service.error.idAndNameCannotBeUsed")
			}

			if opts.id != "" || opts.name != "" {
				return runDelete(opts)
			}

			svcContext, err := opts.ServiceContext.Load()
			if err != nil {
				return err
			}

			profileHandler := &contextutil.ContextHandler{
				Context:   svcContext,
				Localizer: opts.localizer,
			}

			conn, err := opts.Connection(connection.DefaultConfigRequireMasAuth)
			if err != nil {
				return err
			}

			kafkaInstance, err := profileHandler.GetCurrentKafkaInstance(conn.API().KafkaMgmt())
			if err != nil {
				return err
			}

			opts.id = kafkaInstance.GetId()

			return runDelete(opts)
		},
	}

	flags := flagutil.NewFlagSet(cmd, opts.localizer)

	flags.AddYes(&opts.skipConfirm)
	flags.StringVar(&opts.id, "id", "", opts.localizer.MustLocalize("kafka.delete.flag.id"))
	flags.StringVar(&opts.name, "name", "", opts.localizer.MustLocalize("kafka.delete.flag.name"))

	if err := kafkautil.RegisterNameFlagCompletionFunc(cmd, f); err != nil {
		opts.Logger.Debug(opts.localizer.MustLocalize("kafka.common.error.load.completions.name.flag"), err)
	}

	return cmd
}

func runDelete(opts *options) error {
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

	var response *kafkamgmtclient.KafkaRequest
	if opts.name != "" {
		response, _, err = kafkautil.GetKafkaByName(opts.Context, api.KafkaMgmt(), opts.name)
		if err != nil {
			return err
		}
	} else {
		response, _, err = kafkautil.GetKafkaByID(opts.Context, api.KafkaMgmt(), opts.id)
		if err != nil {
			return err
		}
	}

	kafkaName := response.GetName()

	if !opts.skipConfirm {
		promptConfirmName := &survey.Input{
			Message: opts.localizer.MustLocalize("kafka.delete.input.confirmName.message", localize.NewEntry("Name", kafkaName)),
		}

		var confirmedKafkaName string
		err = survey.AskOne(promptConfirmName, &confirmedKafkaName)
		if err != nil {
			return err
		}

		if confirmedKafkaName != kafkaName {
			opts.Logger.Info(opts.localizer.MustLocalize("kafka.delete.log.info.incorrectNameConfirmation"))
			return nil
		}
	}

	// delete the Kafka
	opts.Logger.Debug(opts.localizer.MustLocalize("kafka.delete.log.debug.deletingKafka"), fmt.Sprintf("\"%s\"", kafkaName))
	a := api.KafkaMgmt().DeleteKafkaById(opts.Context, response.GetId())
	a = a.Async(true)
	_, _, err = a.Execute()

	if err != nil {
		return err
	}

	opts.Logger.Info(opts.localizer.MustLocalize("kafka.delete.log.info.deleting", localize.NewEntry("Name", kafkaName)))

	currentKafka := svcConfig.KafkaID
	// this is not the current instance, our work here is done
	if currentKafka != response.GetId() {
		return nil
	}

	// the Kafka that was deleted is set as the user's current cluster
	// since it was deleted it should be removed from the context
	svcConfig.KafkaID = ""
	svcContext.Contexts[svcContext.CurrentContext] = *svcConfig

	if err := opts.ServiceContext.Save(svcContext); err != nil {
		return err
	}

	return nil
}
