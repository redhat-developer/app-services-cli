package update

import (
	"context"
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/redhat-developer/app-services-cli/pkg/auth/token"

	"github.com/AlecAivazis/survey/v2"
	"github.com/redhat-developer/app-services-cli/pkg/core/cmdutil/flagutil"
	"github.com/redhat-developer/app-services-cli/pkg/core/config"
	"github.com/redhat-developer/app-services-cli/pkg/core/ioutil/color"
	"github.com/redhat-developer/app-services-cli/pkg/core/ioutil/icon"
	"github.com/redhat-developer/app-services-cli/pkg/core/ioutil/iostreams"
	"github.com/redhat-developer/app-services-cli/pkg/core/ioutil/spinner"
	"github.com/redhat-developer/app-services-cli/pkg/core/localize"
	"github.com/redhat-developer/app-services-cli/pkg/core/logging"
	"github.com/redhat-developer/app-services-cli/pkg/shared/connection"
	"github.com/redhat-developer/app-services-cli/pkg/shared/factory"
	"github.com/redhat-developer/app-services-cli/pkg/shared/kafkautil"
	kafkamgmtclient "github.com/redhat-developer/app-services-sdk-go/kafkamgmt/apiv1/client"
	"github.com/spf13/cobra"

	"github.com/redhat-developer/app-services-cli/internal/build"
	"github.com/redhat-developer/app-services-cli/pkg/api/rbac"
	"github.com/redhat-developer/app-services-cli/pkg/api/rbac/rbacutil"
)

type options struct {
	name        string
	id          string
	owner       string
	skipConfirm bool

	interactive    bool
	userIsOrgAdmin bool
	reauth         flagutil.Tribool

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
			_, err := opts.Connection(connection.DefaultConfigSkipMasAuth)
			if err != nil {
				return err
			}
			cfg, err := opts.Config.Load()
			if err != nil {
				return err
			}

			opts.userIsOrgAdmin = token.IsOrgAdmin(cfg.AccessToken)
			if !opts.userIsOrgAdmin {
				opts.logger.Info(opts.localizer.MustLocalize("kafka.update.log.info.onlyOrgAdminsCanUpdate"))
				return nil
			}

			if !opts.IO.CanPrompt() {
				var missingFlags []string
				if !opts.skipConfirm {
					missingFlags = append(missingFlags, "yes")
				}
				if len(missingFlags) > 0 {
					return flagutil.RequiredWhenNonInteractiveError(missingFlags...)
				}
			}
			if opts.owner == "" && opts.reauth == "" {
				opts.interactive = true
			}

			if opts.name != "" && opts.id != "" {
				return opts.localizer.MustLocalizeError("service.error.idAndNameCannotBeUsed")
			}

			if opts.id != "" || opts.name != "" {
				return run(&opts)
			}

			instanceID, ok := cfg.GetKafkaIdOk()
			if !ok {
				return opts.localizer.MustLocalizeError("kafka.common.error.noKafkaSelected")
			}
			opts.id = instanceID

			return run(&opts)
		},
	}

	flags := flagutil.NewFlagSet(cmd, opts.localizer)

	flags.StringVar(&opts.id, "id", "", opts.localizer.MustLocalize("kafka.update.flag.id"))
	flags.StringVar(&opts.owner, "owner", "", opts.localizer.MustLocalize("kafka.update.flag.owner"))
	flags.TriBoolVar(&opts.reauth, "reauthentication", flagutil.TRIBOOL_DEFAULT, opts.localizer.MustLocalize("kafka.update.flag.reauthentication"))
	flags.AddYes(&opts.skipConfirm)
	flags.StringVar(&opts.name, "name", "", opts.localizer.MustLocalize("kafka.update.flag.name"))

	_ = kafkautil.RegisterNameFlagCompletionFunc(cmd, f)
	_ = flagutil.RegisterUserCompletionFunc(cmd, "owner", f)

	return cmd
}

func run(opts *options) error {
	conn, err := opts.Connection(connection.DefaultConfigSkipMasAuth)
	if err != nil {
		return err
	}

	api := conn.API()

	kafkaInstance, err := getCurrentKafkaInstance(opts, api.KafkaMgmt())
	if err != nil {
		return err
	}

	if opts.interactive {
		err = runInteractivePrompt(opts, kafkaInstance)
		if err != nil {
			return err
		}
	}

	updateObj, err := getUpdateObj(opts, kafkaInstance)
	if err != nil {
		return err
	}

	// create a text block with a summary of what is being updated
	updateSummary := generateUpdateSummary(reflect.ValueOf(*updateObj), reflect.ValueOf(*kafkaInstance))

	opts.logger.Infof(`
 %v %v

   %v`,
		color.Underline(color.Bold(opts.localizer.MustLocalize("kafka.update.summaryTitle"))),
		icon.Emoji("\U0001f50e", ""),
		updateSummary,
	)

	if opts.reauth == flagutil.TRIBOOL_FALSE {
		opts.logger.Info(opts.localizer.MustLocalize("kafka.update.reauthentication.disclaimer"))
	}

	if !opts.skipConfirm {
		confirm, promptErr := promptConfirmUpdate(opts)
		if promptErr != nil {
			return promptErr
		}
		if !confirm {
			opts.logger.Debug("User has chosen to not update Kafka instance")
			return nil
		}
	}

	s := spinner.New(opts.IO.ErrOut, opts.localizer)
	s.SetLocalizedSuffix("kafka.update.log.info.updating", localize.NewEntry("Name", kafkaInstance.GetName()))
	s.Start()

	response, httpRes, err := api.KafkaMgmt().
		UpdateKafkaById(opts.Context, kafkaInstance.GetId()).
		KafkaUpdateRequest(*updateObj).
		Execute()

	if httpRes != nil {
		defer httpRes.Body.Close()
	}

	s.Stop()

	if err != nil {
		if apiError := kafkautil.GetAPIError(err); apiError != nil {
			return opts.localizer.MustLocalizeError("kafka.update.log.info.updateFailed", localize.NewEntry("Reason", apiError.GetReason()))
		}
		return err
	}

	opts.logger.Info()
	opts.logger.Info(opts.localizer.MustLocalize("kafka.update.log.info.updateSuccess", localize.NewEntry("Name", response.GetName())))

	return nil
}

func runInteractivePrompt(opts *options, kafkaInstance *kafkamgmtclient.KafkaRequest) (err error) {
	opts.owner, err = selectOwnerInteractive(opts.Context, opts)
	if err != nil {
		return err
	}

	reauthenticationPrompt := &survey.Select{
		Message: opts.localizer.MustLocalize("kafka.update.input.message.reauthentication"),
		Options: flagutil.ValidTribools,
		Default: strconv.FormatBool(kafkaInstance.GetReauthenticationEnabled()),
	}

	var reauthStr string

	err = survey.AskOne(reauthenticationPrompt, &reauthStr)
	if err != nil {
		return err
	}

	opts.reauth = flagutil.Tribool(reauthStr)

	return nil
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
	s := spinner.New(opts.IO.ErrOut, opts.localizer)
	s.SetLocalizedSuffix("kafka.update.log.info.loadingUsers")
	s.Start()

	//nolint:govet
	users, err := rbacutil.FetchAllUsers(ctx, conn.API().RBAC().PrincipalAPI)

	s.Stop()
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

		oldVal := getElementValue(current.FieldByName(fieldName))
		newVal := getUpdateObjValue(new.FieldByName(fieldName))

		fieldTagDisp := strings.Split(fieldTag, ",")[0]

		if newVal != reflect.ValueOf(nil) {
			summary += fmt.Sprintf("%v: %v    %v    %v\n", color.Bold(fieldTagDisp), oldVal, icon.Emoji("\u27A1", "=>"), newVal)
		}
	}

	return summary
}

func getCurrentKafkaInstance(opts *options, api kafkamgmtclient.DefaultApi) (kafkaInstance *kafkamgmtclient.KafkaRequest, err error) {

	if opts.name != "" {
		kafkaInstance, _, err = kafkautil.GetKafkaByName(opts.Context, api, opts.name)
		if err != nil {
			return nil, err
		}
		opts.id = kafkaInstance.GetName()
	} else {
		kafkaInstance, _, err = kafkautil.GetKafkaByID(opts.Context, api, opts.id)
		if err != nil {
			return nil, err
		}
		opts.name = kafkaInstance.GetName()
	}

	return kafkaInstance, nil
}

func getUpdateObj(opts *options, kafkaInstance *kafkamgmtclient.KafkaRequest) (*kafkamgmtclient.KafkaUpdateRequest, error) {

	// track if values have been changed
	var needsUpdate bool

	updateObj := kafkamgmtclient.NewKafkaUpdateRequest()

	if opts.owner != "" && opts.owner != kafkaInstance.GetOwner() {
		updateObj.SetOwner(opts.owner)
		needsUpdate = true
	}

	if opts.reauth != flagutil.TRIBOOL_DEFAULT && string(opts.reauth) != strconv.FormatBool(kafkaInstance.GetReauthenticationEnabled()) {
		enableBool, newErr := strconv.ParseBool(string(opts.reauth))
		if newErr != nil {
			return nil, newErr
		}

		updateObj.SetReauthenticationEnabled(enableBool)
		needsUpdate = true
	}

	if !needsUpdate {
		return nil, opts.localizer.MustLocalizeError("kafka.update.log.info.nothingToUpdate")
	}

	return updateObj, nil

}

// get the true value from a reflect.Value
// if it is a pointer, extract the true value
func getElementValue(v reflect.Value) reflect.Value {
	if v.Kind() == reflect.Ptr {
		return v.Elem()
	}
	return v
}

// get the true value from a reflect.Value for KafkaUpdateRequest struct
func getUpdateObjValue(v reflect.Value) reflect.Value {
	if v.Kind() == reflect.Struct {
		vstruct := v.FieldByName("value")
		return vstruct.Elem()
	}
	return v
}
