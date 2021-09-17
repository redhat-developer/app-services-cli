package update

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"strconv"

	"github.com/AlecAivazis/survey/v2"
	"github.com/redhat-developer/app-services-cli/internal/build"
	"github.com/redhat-developer/app-services-cli/internal/config"
	"github.com/redhat-developer/app-services-cli/pkg/api/kas"
	"github.com/redhat-developer/app-services-cli/pkg/api/rbac"
	"github.com/redhat-developer/app-services-cli/pkg/api/rbac/rbacutil"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/factory"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/flag"
	flagutil "github.com/redhat-developer/app-services-cli/pkg/cmdutil/flags"
	"github.com/redhat-developer/app-services-cli/pkg/color"
	"github.com/redhat-developer/app-services-cli/pkg/connection"
	"github.com/redhat-developer/app-services-cli/pkg/dump"
	"github.com/redhat-developer/app-services-cli/pkg/icon"
	"github.com/redhat-developer/app-services-cli/pkg/iostreams"
	"github.com/redhat-developer/app-services-cli/pkg/kafka"
	kafkacmdutil "github.com/redhat-developer/app-services-cli/pkg/kafka/cmdutil"
	"github.com/redhat-developer/app-services-cli/pkg/localize"
	"github.com/redhat-developer/app-services-cli/pkg/logging"
	kafkamgmtclient "github.com/redhat-developer/app-services-sdk-go/kafkamgmt/apiv1/client"
	"github.com/spf13/cobra"
)

type options struct {
	name        string
	id          string
	owner       string
	skipConfirm bool

	outputFormat string

	interactive bool

	IO         *iostreams.IOStreams
	Config     config.IConfig
	Connection factory.ConnectionFunc
	logger     logging.Logger
	localizer  localize.Localizer
	Context    context.Context
}

func NewUpdateCommand(f *factory.Factory) *cobra.Command {
	opts := options{
		IO:         f.IOStreams,
		Config:     f.Config,
		localizer:  f.Localizer,
		logger:     f.Logger,
		Connection: f.Connection,
		Context:    f.Context,
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
			if opts.owner == "" {
				opts.interactive = true
			}

			validOutputFormats := flagutil.ValidOutputFormats
			if opts.outputFormat != "" && !flagutil.IsValidInput(opts.outputFormat, validOutputFormats...) {
				return flag.InvalidValueError("output", opts.outputFormat, validOutputFormats...)
			}

			if opts.name != "" && opts.id != "" {
				return opts.localizer.MustLocalizeError("service.error.idAndNameCannotBeUsed")
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
				return opts.localizer.MustLocalizeError("kafka.common.error.noKafkaSelected")
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

	_ = kafkacmdutil.RegisterNameFlagCompletionFunc(cmd, f)
	_ = kafkacmdutil.RegisterOwnerFlagCompletionFunc(cmd, f)

	return cmd
}

func run(opts *options) error {
	conn, err := opts.Connection(connection.DefaultConfigSkipMasAuth)
	if err != nil {
		return err
	}

	api := conn.API()

	var kafkaInstance *kafkamgmtclient.KafkaRequest
	if opts.name != "" {
		kafkaInstance, _, err = kafka.GetKafkaByName(opts.Context, api.Kafka(), opts.name)
		if err != nil {
			return err
		}
		opts.id = kafkaInstance.GetName()
	} else {
		kafkaInstance, _, err = kafka.GetKafkaByID(opts.Context, api.Kafka(), opts.id)
		if err != nil {
			return err
		}
		opts.name = kafkaInstance.GetName()
	}

	if opts.interactive {
		opts.owner, err = selectOwnerInteractive(opts.Context, opts)
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
 %v %v

   %v`,
		color.Underline(color.Bold(opts.localizer.MustLocalize("kafka.update.summaryTitle"))),
		icon.Emoji("\U0001f50e", ""),
		updateSummary,
	)

	if !opts.skipConfirm {
		//nolint:govet
		confirm, err := promptConfirmUpdate(opts)
		if err != nil {
			return err
		}
		if !confirm {
			opts.logger.Debug("User has chosen to not update Kafka instance")
			return nil
		}
	}

	spinner := opts.IO.NewSpinner()
	spinner.Suffix = " " + opts.localizer.MustLocalize("kafka.update.log.info.updating", localize.NewEntry("Name", kafkaInstance.GetName()))
	spinner.Start()

	response, httpRes, err := api.Kafka().
		UpdateKafkaById(opts.Context, kafkaInstance.GetId()).
		KafkaUpdateRequest(*updateObj).
		Execute()

	if httpRes != nil {
		defer httpRes.Body.Close()
	}

	spinner.Stop()

	if err != nil {
		opts.logger.Info("\n") // Needed to ensure there is a newline after the spinner has stopped
		if apiError, ok := kas.GetAPIError(err); ok {
			return errors.New(apiError.GetReason())
		}
		return err
	}

	opts.logger.Infof(`

%v`, opts.localizer.MustLocalize("kafka.update.log.info.updateSuccess", localize.NewEntry("Name", response.GetName())))

	return dump.Formatted(opts.IO.Out, opts.outputFormat, response)
}

func promptOwnerSelect(localizer localize.Localizer, users []rbac.Principal) (string, error) {
	var usernames []string
	var displayNameUsernameMap = make(map[string]string)

	if len(users) > 0 {
		for _, p := range users {
			displayName := fmt.Sprintf("%v (%v %v)", p.Username, p.FirstName, p.LastName)
			displayNameUsernameMap[displayName] = p.Username
			usernames = append(usernames, displayName)
		}
	}
	prompt := survey.Select{
		Message: localizer.MustLocalize("kafka.update.input.message.selectOwner"),
		Options: usernames,
	}

	var response string
	pageSize, err := strconv.Atoi(build.DefaultPageSize)
	if err != nil {
		pageSize = 10
	}
	if err = survey.AskOne(&prompt, &response, survey.WithPageSize(pageSize)); err != nil {
		return "", err
	}

	username := displayNameUsernameMap[response]

	return username, err
}

func selectOwnerInteractive(ctx context.Context, opts *options) (string, error) {
	conn, err := opts.Connection(connection.DefaultConfigSkipMasAuth)
	if err != nil {
		return "", err
	}
	spinner := opts.IO.NewSpinner()
	spinner.Suffix = " " + opts.localizer.MustLocalize("kafka.update.log.info.loadingUsers")
	spinner.Start()

	//nolint:govet
	users, err := rbacutil.FetchAllUsers(ctx, conn.API().RBAC.PrincipalAPI)

	spinner.Stop()
	opts.logger.Info()
	if err != nil {
		return "", fmt.Errorf("%v: %w", opts.localizer.MustLocalize("kafka.update.error.loadUsersError"), err)
	}
	opts.owner, err = promptOwnerSelect(opts.localizer, users)

	return opts.owner, err
}

func promptConfirmUpdate(opts *options) (bool, error) {
	promptConfirm := survey.Confirm{
		Message: opts.localizer.MustLocalize("kafka.update.confirmDialog.message", localize.NewEntry("Name", opts.name)),
	}

	var confirmUpdate bool
	if err := survey.AskOne(&promptConfirm, &confirmUpdate); err != nil {
		return false, err
	}
	return confirmUpdate, nil
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

		summary += fmt.Sprintf("%v: %v    %v    %v\n", color.Bold(fieldTag), oldVal, icon.Emoji("\u27A1", "=>"), newVal)
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
