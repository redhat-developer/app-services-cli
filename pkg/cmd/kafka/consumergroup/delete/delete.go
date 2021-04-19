package delete

import (
	"context"
	"errors"
	"fmt"

	"github.com/AlecAivazis/survey/v2"
	"github.com/redhat-developer/app-services-cli/internal/config"
	"github.com/redhat-developer/app-services-cli/internal/localizer"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/factory"
	"github.com/redhat-developer/app-services-cli/pkg/cmdutil"
	"github.com/redhat-developer/app-services-cli/pkg/connection"
	"github.com/redhat-developer/app-services-cli/pkg/iostreams"
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
	Logger     func() (logging.Logger, error)
}

// NewDeleteConsumerGroupCommand gets a new command for deleting a consumer group.
func NewDeleteConsumerGroupCommand(f *factory.Factory) *cobra.Command {
	opts := &Options{
		Connection: f.Connection,
		Config:     f.Config,
		IO:         f.IOStreams,
		Logger:     f.Logger,
	}

	cmd := &cobra.Command{
		Use:     localizer.MustLocalizeFromID("kafka.consumerGroup.delete.cmd.use"),
		Short:   localizer.MustLocalizeFromID("kafka.consumerGroup.delete.cmd.shortDescription"),
		Long:    localizer.MustLocalizeFromID("kafka.consumerGroup.delete.cmd.longDescription"),
		Example: localizer.MustLocalizeFromID("kafka.consumerGroup.delete.cmd.example"),
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
				return errors.New(localizer.MustLocalizeFromID("kafka.consumerGroup.common.error.noKafkaSelected"))
			}

			opts.kafkaID = cfg.Services.Kafka.ClusterID

			return runCmd(opts)
		},
	}

	cmd.Flags().BoolVarP(&opts.skipConfirm, "yes", "y", false, localizer.MustLocalizeFromID("kafka.consumerGroup.delete.flag.yes.description"))
	cmd.Flags().StringVar(&opts.id, "id", "", localizer.MustLocalizeFromID("kafka.consumerGroup.common.flag.id.description"))
	_ = cmd.MarkFlagRequired("id")

	// flag based completions for ID
	_ = cmd.RegisterFlagCompletionFunc("id", func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return cmdutil.FilterValidConsumerGroupIDs(f, toComplete)
	})

	return cmd
}

// nolint:funlen
func runCmd(opts *Options) error {

	logger, err := opts.Logger()
	if err != nil {
		return err
	}

	conn, err := opts.Connection(connection.DefaultConfigRequireMasAuth)
	if err != nil {
		return err
	}

	api, kafkaInstance, err := conn.API().TopicAdmin(opts.kafkaID)
	if err != nil {
		return err
	}

	ctx := context.Background()

	_, httpRes, err := api.GetConsumerGroupById(ctx, opts.id).Execute()

	if err != nil {
		if httpRes == nil {
			return err
		}
		if httpRes.StatusCode == 404 {
			return errors.New(localizer.MustLocalize(&localizer.Config{
				MessageID: "kafka.consumerGroup.common.error.notFoundError",
				TemplateData: map[string]interface{}{
					"ID":           opts.id,
					"InstanceName": kafkaInstance.GetName(),
				},
			}))
		}
	}

	if !opts.skipConfirm {
		var confirmedID string
		promptConfirmDelete := &survey.Input{
			Message: localizer.MustLocalizeFromID("kafka.consumerGroup.delete.input.name.message"),
		}

		err = survey.AskOne(promptConfirmDelete, &confirmedID)
		if err != nil {
			return err
		}

		if confirmedID != opts.id {
			return errors.New(localizer.MustLocalize(&localizer.Config{
				MessageID: "kafka.consumerGroup.delete.error.mismatchedIDConfirmation",
				TemplateData: map[string]interface{}{
					"ConfirmedID": confirmedID,
					"ID":          opts.id,
				},
			}))
		}
	}

	httpRes, err = api.DeleteConsumerGroupById(ctx, opts.id).Execute()

	if err != nil {
		if httpRes == nil {
			return err
		}

		switch httpRes.StatusCode {
		case 401:
			return errors.New(localizer.MustLocalize(&localizer.Config{
				MessageID: "kafka.consumerGroup.common.error.unauthorized",
				TemplateData: map[string]interface{}{
					"Operation": "delete",
				},
			}))
		case 403:
			return errors.New(localizer.MustLocalize(&localizer.Config{
				MessageID: "kafka.consumerGroup.common.error.forbidden",
				TemplateData: map[string]interface{}{
					"Operation": "delete",
				},
			}))
		case 423:
			return errors.New(localizer.MustLocalizeFromID("kafka.consumerGroup.delete.error.locked"))
		case 500:
			return errors.New(localizer.MustLocalizeFromID("kafka.consumerGroup.common.error.internalServerError"))
		case 503:
			return fmt.Errorf("%v: %w", localizer.MustLocalize(&localizer.Config{
				MessageID: "kafka.consumerGroup.common.error.unableToConnectToKafka",
				TemplateData: map[string]interface{}{
					"Name": kafkaInstance.GetName(),
				},
			}), err)
		default:
			return err
		}
	}

	logger.Info(localizer.MustLocalize(&localizer.Config{
		MessageID: "kafka.consumerGroup.delete.log.info.consumerGroupDeleted",
		TemplateData: map[string]interface{}{
			"ConsumerGroupID": opts.id,
			"InstanceName":    kafkaInstance.GetName(),
		},
	}))

	return nil
}
