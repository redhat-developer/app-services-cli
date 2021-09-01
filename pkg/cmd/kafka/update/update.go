package update

import (
	"context"
	"errors"
	"fmt"
	"reflect"

	"github.com/AlecAivazis/survey/v2"
	"github.com/redhat-developer/app-services-cli/internal/config"
	"github.com/redhat-developer/app-services-cli/pkg/api/kas"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/factory"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/flag"
	flagutil "github.com/redhat-developer/app-services-cli/pkg/cmdutil/flags"
	"github.com/redhat-developer/app-services-cli/pkg/color"
	"github.com/redhat-developer/app-services-cli/pkg/connection"
	"github.com/redhat-developer/app-services-cli/pkg/dump"
	"github.com/redhat-developer/app-services-cli/pkg/iostreams"
	"github.com/redhat-developer/app-services-cli/pkg/kafka"
	kafkacmdutil "github.com/redhat-developer/app-services-cli/pkg/kafka/cmdutil"
	"github.com/redhat-developer/app-services-cli/pkg/localize"
	"github.com/redhat-developer/app-services-cli/pkg/logging"
	kafkamgmtclient "github.com/redhat-developer/app-services-sdk-go/kafkamgmt/apiv1/client"
	"github.com/spf13/cobra"
)

type Options struct {
	name        string
	id          string
	owner       string
	skipConfirm bool

	outputFormat string

	IO         *iostreams.IOStreams
	Config     config.IConfig
	Connection factory.ConnectionFunc
	logger     logging.Logger
	localizer  localize.Localizer
}

func NewUpdateCommand(f *factory.Factory) *cobra.Command {
	opts := Options{
		IO:         f.IOStreams,
		Config:     f.Config,
		localizer:  f.Localizer,
		logger:     f.Logger,
		Connection: f.Connection,
	}

	cmd := &cobra.Command{
		Use:     "update",
		Short:   opts.localizer.MustLocalize("kafka.update.cmd.shortDescription"),
		Long:    opts.localizer.MustLocalize("kafka.update.cmd.longDescription"),
		Example: opts.localizer.MustLocalize("kafka.update.cmd.examples"),
		Args:    cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			if !opts.IO.CanPrompt() {
				var missingFlags []string
				if opts.owner == "" {
					missingFlags = append(missingFlags, "owner")
				}
				if !opts.skipConfirm {
					missingFlags = append(missingFlags, "yes")
				}
				if len(missingFlags) > 0 {
					return flag.RequiredWhenNonInteractiveError(missingFlags...)
				}
			}

			validOutputFormats := flagutil.ValidOutputFormats
			if opts.outputFormat != "" && !flagutil.IsValidInput(opts.outputFormat, validOutputFormats...) {
				return flag.InvalidValueError("output", opts.outputFormat, validOutputFormats...)
			}

			if opts.name != "" && opts.id != "" {
				return errors.New(opts.localizer.MustLocalize("service.error.idAndNameCannotBeUsed"))
			}

			if opts.id != "" || opts.name != "" {
				return run(&opts)
			}

			cfg, err := opts.Config.Load()
			if err != nil {
				return err
			}

			var kafkaConfig *config.KafkaConfig
			if cfg.Services.Kafka == kafkaConfig || cfg.Services.Kafka.ClusterID == "" {
				return errors.New(opts.localizer.MustLocalize("kafka.common.error.noKafkaSelected"))
			}

			opts.id = cfg.Services.Kafka.ClusterID

			return run(&opts)
		},
	}

	cmd.Flags().StringVarP(&opts.outputFormat, "output", "o", dump.JSONFormat, opts.localizer.MustLocalize("kafka.common.flag.output.description"))
	cmd.Flags().StringVar(&opts.id, "id", "", opts.localizer.MustLocalize("kafka.update.flag.id"))
	cmd.Flags().StringVar(&opts.owner, "owner", "", opts.localizer.MustLocalize("kafka.update.flag.owner"))
	cmd.Flags().BoolVarP(&opts.skipConfirm, "yes", "y", false, opts.localizer.MustLocalize("kafka.update.flag.yes"))
	cmd.Flags().StringVar(&opts.name, "name", "", opts.localizer.MustLocalize("kafka.update.flag.name"))

	_ = cmd.MarkFlagRequired("owner")

	_ = kafkacmdutil.RegisterNameFlagCompletionFunc(cmd, f)

	return cmd
}

func run(opts *Options) error {
	conn, err := opts.Connection(connection.DefaultConfigRequireMasAuth)
	if err != nil {
		return err
	}

	api := conn.API()

	var kafkaInstance *kafkamgmtclient.KafkaRequest
	ctx := context.Background()
	if opts.name != "" {
		kafkaInstance, _, err = kafka.GetKafkaByName(ctx, api.Kafka(), opts.name)
		if err != nil {
			return err
		}
	} else {
		kafkaInstance, _, err = kafka.GetKafkaByID(ctx, api.Kafka(), opts.id)
		if err != nil {
			return err
		}
	}

	if opts.owner == kafkaInstance.GetOwner() {
		opts.logger.Info(opts.localizer.MustLocalize("kafka.update.log.info.sameOwnerNoChanges", localize.NewEntry("Owner", opts.owner), localize.NewEntry("InstanceName", kafkaInstance.GetName())))
		return nil
	}

	updateObj := kafkamgmtclient.NewKafkaUpdateRequest(opts.owner)

	// create a text block with a summary of what is being updated
	updateSummary := generateUpdateSummary(reflect.ValueOf(*updateObj), reflect.ValueOf(*kafkaInstance))

	opts.logger.Infof(`
%v 🗒️

 %v`, color.Underline(color.Bold(opts.localizer.MustLocalize("kafka.update.summaryTitle"))), updateSummary)

	if !opts.skipConfirm {
		promptConfirm := survey.Confirm{
			Message: opts.localizer.MustLocalize("kafka.update.confirmDialog.message", localize.NewEntry("Name", kafkaInstance.GetName())),
		}
		var confirmUpdate bool
		if err = survey.AskOne(&promptConfirm, &confirmUpdate); err != nil {
			return err
		}
		if !confirmUpdate {
			opts.logger.Debug("User has chosen to not update Kafka instance")
			return nil
		}
	}

	spinner := opts.IO.NewSpinner()
	spinner.Suffix = " " + opts.localizer.MustLocalize("kafka.update.log.info.updating", localize.NewEntry("Name", kafkaInstance.GetName()))
	spinner.Start()

	response, httpRes, err := api.Kafka().
		UpdateKafkaById(context.Background(), kafkaInstance.GetId()).
		KafkaUpdateRequest(*updateObj).
		Execute()

	spinner.Stop()

	if err != nil {
		opts.logger.Info("\n") // Needed to ensure there is a newline after the spinner has stopped
		if apiError, ok := kas.GetAPIError(err); ok {
			return errors.New(apiError.GetReason())
		}
		return err
	}

	defer httpRes.Body.Close()

	opts.logger.Infof(`

%v`, opts.localizer.MustLocalize("kafka.update.log.info.updateSuccess", localize.NewEntry("Name", response.GetName())))

	dump.PrintDataInFormat(opts.outputFormat, response, opts.IO.Out)

	return nil
}

// creates a summary of what values will be changed in this update
// returns a formatted string. Example:
// owner: foo_user	➡️	bar_user
func generateUpdateSummary(new reflect.Value, current reflect.Value) string {
	var summary string

	newT := new.Type()

	for i := 0; i < new.NumField(); i++ {
		field := newT.Field(i)
		fieldTag := field.Tag.Get("json")
		fieldName := field.Name

		oldVal := getElementValue(current.FieldByName(fieldName)).String()
		newVal := getElementValue(new.FieldByName(fieldName)).String()

		summary += fmt.Sprintf("%v: %v   ➡️ ️	%v\n", color.Bold(fieldTag), oldVal, newVal)
	}

	return summary
}

// get the true value from a reflect.Value
// if it is a pointer, extract the true value
func getElementValue(v reflect.Value) reflect.Value {
	if v.Kind() == reflect.Ptr {
		return v.Elem()
	}
	return v
}
