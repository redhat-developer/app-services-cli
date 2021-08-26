package delete

import (
	"context"
	"errors"
	"net/http"

	"github.com/AlecAivazis/survey/v2"
	"github.com/redhat-developer/app-services-cli/internal/config"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/factory"
	"github.com/redhat-developer/app-services-cli/pkg/cmdutil"
	"github.com/redhat-developer/app-services-cli/pkg/connection"
	"github.com/redhat-developer/app-services-cli/pkg/iostreams"
	"github.com/redhat-developer/app-services-cli/pkg/localize"
	"github.com/redhat-developer/app-services-cli/pkg/logging"
	"github.com/spf13/cobra"
)

type Options struct {
	kafkaID     string
	id          string
	skipConfirm bool

	IO         *iostreams.IOStreams
	Config     config.IConfig
	Connection factory.ConnectionFunc
	Logger     logging.Logger
	localizer  localize.Localizer
}

// NewDeleteConsumerGroupCommand gets a new command for deleting a consumer group.
func NewDeleteConsumerGroupCommand(f *factory.Factory) *cobra.Command {
	opts := &Options{
		Connection: f.Connection,
		Config:     f.Config,
		IO:         f.IOStreams,
		Logger:     f.Logger,
		localizer:  f.Localizer,
	}

	cmd := &cobra.Command{
		Use:     opts.localizer.MustLocalize("kafka.consumerGroup.delete.cmd.use"),
		Short:   opts.localizer.MustLocalize("kafka.consumerGroup.delete.cmd.shortDescription"),
		Long:    opts.localizer.MustLocalize("kafka.consumerGroup.delete.cmd.longDescription"),
		Example: opts.localizer.MustLocalize("kafka.consumerGroup.delete.cmd.example"),
		Args:    cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			if opts.kafkaID != "" {
				return runCmd(opts)
			}

			cfg, err := opts.Config.Load()
			if err != nil {
				return err
			}

			if !cfg.HasKafka() {
				return errors.New(opts.localizer.MustLocalize("kafka.consumerGroup.common.error.noKafkaSelected"))
			}

			opts.kafkaID = cfg.Services.Kafka.ClusterID

			return runCmd(opts)
		},
	}

	opts.localizer.MustLocalize("kafka.consumerGroup.common.flag.id.description", localize.NewEntry("Action", "delete"))
	cmd.Flags().BoolVarP(&opts.skipConfirm, "yes", "y", false, opts.localizer.MustLocalize("kafka.consumerGroup.delete.flag.yes.description"))
	cmd.Flags().StringVar(&opts.id, "id", "", opts.localizer.MustLocalize("kafka.consumerGroup.common.flag.id.description", localize.NewEntry("Action", "delete")))
	_ = cmd.MarkFlagRequired("id")

	// flag based completions for ID
	_ = cmd.RegisterFlagCompletionFunc("id", func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return cmdutil.FilterValidConsumerGroupIDs(f, toComplete)
	})

	return cmd
}

// nolint:funlen
func runCmd(opts *Options) error {
	conn, err := opts.Connection(connection.DefaultConfigRequireMasAuth)
	if err != nil {
		return err
	}

	api, kafkaInstance, err := conn.API().KafkaAdmin(opts.kafkaID)
	if err != nil {
		return err
	}

	ctx := context.Background()

	_, httpRes, err := api.GroupsApi.GetConsumerGroupById(ctx, opts.id).Execute()
	defer httpRes.Body.Close()

	cgIDPair := localize.NewEntry("ID", opts.id)
	kafkaNameTmplPair := localize.NewEntry("InstanceName", kafkaInstance.GetName())
	if err != nil {
		if httpRes == nil {
			return err
		}
		if httpRes.StatusCode == http.StatusNotFound {
			return errors.New(opts.localizer.MustLocalize("kafka.consumerGroup.common.error.notFoundError", cgIDPair, kafkaNameTmplPair))
		}
	}

	if !opts.skipConfirm {
		var confirmedID string
		promptConfirmDelete := &survey.Input{
			Message: opts.localizer.MustLocalize("kafka.consumerGroup.delete.input.name.message"),
		}

		err = survey.AskOne(promptConfirmDelete, &confirmedID)
		if err != nil {
			return err
		}

		if confirmedID != opts.id {
			return errors.New(opts.localizer.MustLocalize("kafka.consumerGroup.delete.error.mismatchedIDConfirmation", localize.NewEntry("ConfirmedID", confirmedID), cgIDPair))
		}
	}

	httpRes, err = api.GroupsApi.DeleteConsumerGroupById(ctx, opts.id).Execute()

	if err != nil {
		if httpRes == nil {
			return err
		}

		defer httpRes.Body.Close()

		operationTmplPair := localize.NewEntry("Operation", "delete")

		switch httpRes.StatusCode {
		case http.StatusUnauthorized:
			return errors.New(opts.localizer.MustLocalize("kafka.consumerGroup.common.error.unauthorized", operationTmplPair))
		case http.StatusForbidden:
			return errors.New(opts.localizer.MustLocalize("kafka.consumerGroup.common.error.forbidden", operationTmplPair))
		case http.StatusLocked:
			return errors.New(opts.localizer.MustLocalize("kafka.consumerGroup.delete.error.locked"))
		case http.StatusInternalServerError:
			return errors.New(opts.localizer.MustLocalize("kafka.consumerGroup.common.error.internalServerError"))
		case http.StatusServiceUnavailable:
			return errors.New(opts.localizer.MustLocalize("kafka.consumerGroup.common.error.unableToConnectToKafka", localize.NewEntry("Name", kafkaInstance.GetName())))
		default:
			return err
		}
	}

	opts.Logger.Info(opts.localizer.MustLocalize("kafka.consumerGroup.delete.log.info.consumerGroupDeleted", localize.NewEntry("ConsumerGroupID", opts.id), kafkaNameTmplPair))

	return nil
}
