package create

import (
	"fmt"
	"os"
	"os/signal"
	"time"

	kafkaFlagutil "github.com/redhat-developer/app-services-cli/pkg/cmd/kafka/flagutil"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/kafka/kafkacmdutil"
	"github.com/redhat-developer/app-services-cli/pkg/shared/accountmgmtutil"
	"github.com/redhat-developer/app-services-cli/pkg/shared/contextutil"
	"github.com/redhat-developer/app-services-cli/pkg/shared/remote"
	"github.com/redhat-developer/app-services-cli/pkg/shared/svcstatus"

	"github.com/redhat-developer/app-services-cli/pkg/core/cmdutil"
	"github.com/redhat-developer/app-services-cli/pkg/core/cmdutil/flagutil"
	"github.com/redhat-developer/app-services-cli/pkg/core/ioutil/color"
	"github.com/redhat-developer/app-services-cli/pkg/core/ioutil/dump"
	"github.com/redhat-developer/app-services-cli/pkg/core/ioutil/icon"
	"github.com/redhat-developer/app-services-cli/pkg/core/ioutil/spinner"
	"github.com/redhat-developer/app-services-cli/pkg/core/localize"
	"github.com/redhat-developer/app-services-cli/pkg/shared/connection"
	"github.com/redhat-developer/app-services-cli/pkg/shared/factory"

	kafkamgmtclient "github.com/redhat-developer/app-services-sdk-go/kafkamgmt/apiv1/client"
	kafkamgmtv1errors "github.com/redhat-developer/app-services-sdk-go/kafkamgmt/apiv1/error"

	"github.com/AlecAivazis/survey/v2"

	"github.com/spf13/cobra"
)

const (
	// FlagProvider is a flag representing an provider ID
	FlagProvider = "provider"
	// FlagRegion is a flag representing an region ID
	FlagRegion = "region"
	// FlagSize is a flag representing an size ID
	FlagSize = "size"
	// FlagMarketPlaceAcctID is a flag representing a marketplace cloud account ID used to purchase the instance
	FlagMarketPlaceAcctID = "marketplace-account-id"
	// FlagMarketPlace is a flag representing marketplace where the instance is purchased on
	FlagMarketPlace = "marketplace"
)

type options struct {
	name     string
	provider string
	region   string
	size     string

	marketplaceAcctId string
	marketplace       string

	outputFormat string
	autoUse      bool

	interactive  bool
	wait         bool
	bypassChecks bool
	dryRun       bool

	f *factory.Factory
}

var (
	defaultRegion   = "us-east-1"
	defaultProvider = "aws"
)

// NewCreateCommand creates a new command for creating kafkas.
// nolint: funlen
func NewCreateCommand(f *factory.Factory) *cobra.Command {
	opts := &options{
		f: f,
	}

	cmd := &cobra.Command{
		Use:     "create",
		Short:   f.Localizer.MustLocalize("kafka.create.cmd.shortDescription"),
		Long:    f.Localizer.MustLocalize("kafka.create.cmd.longDescription"),
		Example: f.Localizer.MustLocalize("kafka.create.cmd.example"),
		Args:    cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			if opts.name != "" {
				validator := &kafkacmdutil.Validator{
					Localizer:  f.Localizer,
					Connection: f.Connection,
				}
				if err := validator.ValidateName(opts.name); err != nil {
					return err
				}
			}

			if !f.IOStreams.CanPrompt() && opts.name == "" {
				return f.Localizer.MustLocalizeError("kafka.create.argument.name.error.requiredWhenNonInteractive")
			} else if opts.name == "" {
				if opts.provider != "" || opts.region != "" {
					return f.Localizer.MustLocalizeError("kafka.create.argument.name.error.requiredWhenNonInteractive")
				}
				opts.interactive = true
			}

			validOutputFormats := flagutil.ValidOutputFormats
			if opts.outputFormat != "" && !flagutil.IsValidInput(opts.outputFormat, validOutputFormats...) {
				return flagutil.InvalidValueError("output", opts.outputFormat, validOutputFormats...)
			}

			validMarketplaceAcctIDs, err := accountmgmtutil.GetValidMarketplaceAcctIDs(f.Context, f.Connection)
			if err != nil {
				return err
			}

			if opts.marketplaceAcctId != "" && !flagutil.IsValidInput(opts.marketplaceAcctId, validMarketplaceAcctIDs...) {
				return flagutil.InvalidValueError(FlagMarketPlaceAcctID, opts.marketplaceAcctId, validMarketplaceAcctIDs...)
			}

			validMarketplaces, err := accountmgmtutil.GetValidMarketplaces(f.Context, f.Connection)
			if err != nil {
				return err
			}

			if opts.marketplace != "" && !flagutil.IsValidInput(opts.marketplace, validMarketplaces...) {
				return flagutil.InvalidValueError(FlagMarketPlace, opts.marketplace, validMarketplaces...)
			}

			return runCreate(opts)
		},
	}

	flags := kafkaFlagutil.NewFlagSet(cmd, f.Localizer)

	flags.StringVar(&opts.name, "name", "", f.Localizer.MustLocalize("kafka.create.flag.name.description"))
	flags.StringVar(&opts.provider, FlagProvider, "", f.Localizer.MustLocalize("kafka.create.flag.cloudProvider.description"))
	flags.StringVar(&opts.region, FlagRegion, "", f.Localizer.MustLocalize("kafka.create.flag.cloudRegion.description"))
	flags.StringVar(&opts.size, FlagSize, "", f.Localizer.MustLocalize("kafka.create.flag.size.description"))
	flags.StringVar(&opts.marketplaceAcctId, FlagMarketPlaceAcctID, "", f.Localizer.MustLocalize("kafka.create.flag.marketplaceId.description"))
	flags.StringVar(&opts.marketplace, FlagMarketPlace, "", f.Localizer.MustLocalize("kafka.create.flag.marketplaceType.description"))
	flags.AddOutput(&opts.outputFormat)
	flags.BoolVar(&opts.autoUse, "use", true, f.Localizer.MustLocalize("kafka.create.flag.autoUse.description"))
	flags.BoolVarP(&opts.wait, "wait", "w", false, f.Localizer.MustLocalize("kafka.create.flag.wait.description"))
	flags.BoolVarP(&opts.dryRun, "dry-run", "", false, f.Localizer.MustLocalize("kafka.create.flag.dryrun.description"))
	flags.AddBypassTermsCheck(&opts.bypassChecks)

	_ = cmd.RegisterFlagCompletionFunc(FlagProvider, func(cmd *cobra.Command, _ []string, _ string) ([]string, cobra.ShellCompDirective) {
		return GetCloudProviderCompletionValues(f)
	})

	_ = cmd.RegisterFlagCompletionFunc(FlagRegion, func(cmd *cobra.Command, _ []string, _ string) ([]string, cobra.ShellCompDirective) {
		return GetCloudProviderRegionCompletionValues(f, opts.provider)
	})

	_ = cmd.RegisterFlagCompletionFunc(FlagSize, func(cmd *cobra.Command, _ []string, _ string) ([]string, cobra.ShellCompDirective) {
		return GetKafkaSizeCompletionValues(f, opts.provider, opts.region)
	})

	_ = cmd.RegisterFlagCompletionFunc(FlagMarketPlaceAcctID, func(cmd *cobra.Command, _ []string, _ string) ([]string, cobra.ShellCompDirective) {
		return GetMarketplaceAcctIdCompletionValues(f)
	})

	_ = cmd.RegisterFlagCompletionFunc(FlagMarketPlace, func(cmd *cobra.Command, _ []string, _ string) ([]string, cobra.ShellCompDirective) {
		return GetMarketplaceCompletionValues(f)
	})

	return cmd
}

// nolint:funlen
func runCreate(opts *options) error {
	f := opts.f
	svcContext, err := f.ServiceContext.Load()
	if err != nil {
		return err
	}

	currCtx, err := contextutil.GetCurrentContext(svcContext, opts.f.Localizer)
	if err != nil {
		return err
	}

	var conn connection.Connection
	if conn, err = f.Connection(connection.DefaultConfigSkipMasAuth); err != nil {
		return err
	}

	err, constants := remote.GetRemoteServiceConstants(f.Context, f.Logger)
	if err != nil {
		return err
	}

	var userInstanceType *accountmgmtutil.QuotaSpec
	if !opts.bypassChecks {
		f.Logger.Debug("Checking if terms and conditions have been accepted")
		// the user must have accepted the terms and conditions from the provider
		// before they can create a kafka instance
		var termsAccepted bool
		var termsURL string
		termsAccepted, termsURL, err = accountmgmtutil.CheckTermsAccepted(f.Context, &constants.Kafka.Ams, conn)

		if err != nil {
			return err
		}
		if !termsAccepted && termsURL != "" {
			f.Logger.Info(f.Localizer.MustLocalize("service.info.termsCheck", localize.NewEntry("TermsURL", termsURL)))
			return nil
		}

		userInstanceType, err = accountmgmtutil.GetUserSupportedInstanceType(f.Context, &constants.Kafka.Ams, conn)
		if err != nil || userInstanceType == nil {
			return f.Localizer.MustLocalizeError("kafka.create.error.userInstanceType.notFound")
		}
	}

	var payload *kafkamgmtclient.KafkaRequestPayload
	if opts.interactive {
		f.Logger.Debug()
		if userInstanceType == nil {
			return f.Localizer.MustLocalizeError("kafka.create.error.noInteractiveMode")
		}

		payload, err = promptKafkaPayload(opts, *userInstanceType)
		if err != nil {
			return err
		}
	} else {
		if opts.provider == "" {
			opts.provider = defaultProvider
		}

		if opts.region == "" {
			opts.region = defaultRegion
		}

		payload = &kafkamgmtclient.KafkaRequestPayload{
			Name:          opts.name,
			Region:        &opts.region,
			CloudProvider: &opts.provider,
		}

		if opts.marketplaceAcctId != "" && opts.marketplace != "" {
			payload.Marketplace = kafkamgmtclient.NullableString{}
			payload.Marketplace.Set(&opts.marketplace)
			payload.BillingCloudAccountId = kafkamgmtclient.NullableString{}
			payload.BillingCloudAccountId.Set(&opts.marketplaceAcctId)
		}

		if !opts.bypassChecks {
			validator := ValidatorInput{
				provider:            opts.provider,
				region:              opts.region,
				size:                opts.size,
				userAMSInstanceType: userInstanceType,
				f:                   f,
				constants:           constants,
				conn:                conn,
			}
			err1 := validator.ValidateProviderAndRegion()
			if err1 != nil {
				return err1
			}

			err1 = validator.ValidateSize()
			if err1 != nil {
				return err1
			}
			if opts.size != "" {
				sizes, err1 := FetchValidKafkaSizes(opts.f, opts.provider, opts.region, *userInstanceType)
				if err1 != nil {
					return err1
				}
				printSizeWarningIfNeeded(opts.f, opts.size, sizes)
				payload.SetPlan(mapAmsTypeToBackendType(userInstanceType) + "." + opts.size)
			}
		}

	}

	f.Logger.Debug("Creating kafka instance", payload.Name, payload.GetPlan())

	if opts.dryRun {
		f.Logger.Info(f.Localizer.MustLocalize("kafka.create.log.info.dryRun.success"))
		return nil
	}

	api := conn.API()

	a := api.KafkaMgmt().CreateKafka(f.Context)
	a = a.KafkaRequestPayload(*payload)
	a = a.Async(true)

	response, httpRes, err := a.Execute()
	if httpRes != nil {
		defer httpRes.Body.Close()
	}

	if apiErr := kafkamgmtv1errors.GetAPIError(err); apiErr != nil {
		switch apiErr.GetCode() {
		case kafkamgmtv1errors.ERROR_120:
			// For standard instances
			return f.Localizer.MustLocalizeError("kafka.create.error.quota.exceeded")
		case kafkamgmtv1errors.ERROR_24:
			// For dev instances
			return f.Localizer.MustLocalizeError("kafka.create.error.instance.limit")
		case kafkamgmtv1errors.ERROR_36:
			return f.Localizer.MustLocalizeError("kafka.create.error.conflictError", localize.NewEntry("Name", payload.Name))
		case kafkamgmtv1errors.ERROR_41:
			return f.Localizer.MustLocalizeError("kafka.create.error.notsupported", localize.NewEntry("Name", payload.Name))
		case kafkamgmtv1errors.ERROR_42:
			return f.Localizer.MustLocalizeError("kafka.create.error.plan.notsupported", localize.NewEntry("Plan", payload.Plan))
		case kafkamgmtv1errors.ERROR_43:
			return f.Localizer.MustLocalizeError("kafka.create.error.billing.invalid", localize.NewEntry("Billing", payload.BillingCloudAccountId))
		}
	}

	if err != nil {
		return err
	}

	if opts.autoUse {
		f.Logger.Debug("Auto-use is set, updating the current instance")
		currCtx.KafkaID = response.GetId()
		svcContext.Contexts[svcContext.CurrentContext] = *currCtx

		if err = f.ServiceContext.Save(svcContext); err != nil {
			return fmt.Errorf("%v: %w", f.Localizer.MustLocalize("kafka.common.error.couldNotUseKafka"), err)
		}
	} else {
		f.Logger.Debug("Auto-use is not set, skipping updating the current instance")
	}

	nameTemplateEntry := localize.NewEntry("Name", response.GetName())

	if opts.wait {
		f.Logger.Debug("--wait flag is enabled, waiting for Kafka to finish creating")
		s := spinner.New(f.IOStreams.ErrOut, f.Localizer)
		s.SetLocalizedSuffix("kafka.create.log.info.creatingKafka", nameTemplateEntry)
		s.Start()

		// when there is a SIGINT, display a message informing the user that this does not cancel the creation
		// and that it is being created in the background
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt)
		go func() {
			for range c {
				f.Logger.Info()
				f.Logger.Info(f.Localizer.MustLocalize("kafka.create.log.info.creatingKafkaSyncSigint"))
				os.Exit(0)
			}
		}()

		for svcstatus.IsInstanceCreating(response.GetStatus()) {
			time.Sleep(cmdutil.DefaultPollTime)

			response, httpRes, err = api.KafkaMgmt().GetKafkaById(f.Context, response.GetId()).Execute()
			if err != nil {
				return err
			}
			defer httpRes.Body.Close()
			f.Logger.Debug("Checking Kafka status:", response.GetStatus())

			s.SetLocalizedSuffix("kafka.create.log.info.creationInProgress",
				localize.NewEntry("Name", response.GetName()),
				localize.NewEntry("Status", color.Info(response.GetStatus())),
			)

		}
		s.Stop()
		f.Logger.Info()
		f.Logger.Info(icon.SuccessPrefix(), f.Localizer.MustLocalize("kafka.create.info.successSync", nameTemplateEntry))
	}

	if err = dump.Formatted(f.IOStreams.Out, opts.outputFormat, response); err != nil {
		return err
	}

	if !opts.wait {
		f.Logger.Info()
		f.Logger.Info(f.Localizer.MustLocalize("kafka.create.info.successAsync", nameTemplateEntry))
	}

	return nil
}

// set type to store the answers from the prompt with defaults
type promptAnswers struct {
	Name              string
	Size              string
	Region            string
	CloudProvider     string
	MarketplaceAcctID string
	Marketplace       string
}

// Show a prompt to allow the user to interactively insert the data for their Kafka
// nolint:funlen
func promptKafkaPayload(opts *options, userQuotaType accountmgmtutil.QuotaSpec) (*kafkamgmtclient.KafkaRequestPayload, error) {
	f := opts.f

	validator := &kafkacmdutil.Validator{
		Localizer:  f.Localizer,
		Connection: f.Connection,
	}

	promptName := &survey.Input{
		Message: f.Localizer.MustLocalize("kafka.create.input.name.message"),
		Help:    f.Localizer.MustLocalize("kafka.create.input.name.help"),
	}

	answers := &promptAnswers{}

	err := survey.AskOne(promptName, &answers.Name, survey.WithValidator(validator.ValidateName), survey.WithValidator(validator.ValidateNameIsAvailable))
	if err != nil {
		return nil, err
	}

	cloudProviderNames, err := GetEnabledCloudProviderNames(opts.f)
	if err != nil {
		return nil, err
	}

	cloudProviderPrompt := &survey.Select{
		Message: f.Localizer.MustLocalize("kafka.create.input.cloudProvider.message"),
		Options: cloudProviderNames,
	}

	err = survey.AskOne(cloudProviderPrompt, &answers.CloudProvider)
	if err != nil {
		return nil, err
	}

	regionIDs, err := GetEnabledCloudRegionIDs(opts.f, answers.CloudProvider, &userQuotaType)
	if err != nil {
		return nil, err
	}

	regionPrompt := &survey.Select{
		Message: f.Localizer.MustLocalize("kafka.create.input.cloudRegion.message"),
		Options: regionIDs,
		Help:    f.Localizer.MustLocalize("kafka.create.input.cloudRegion.help"),
	}

	err = survey.AskOne(regionPrompt, &answers.Region)
	if err != nil {
		return nil, err
	}

	sizes, err := FetchValidKafkaSizes(opts.f, answers.CloudProvider, answers.Region, userQuotaType)
	if err != nil {
		return nil, err
	}

	if len(sizes) == 1 {
		answers.Size = sizes[0].GetId()
	} else {
		sizeLabels := GetValidKafkaSizesLabels(sizes)
		planPrompt := &survey.Select{
			Message: f.Localizer.MustLocalize("kafka.create.input.plan.message"),
			Options: sizeLabels,
		}

		err = survey.AskOne(planPrompt, &answers.Size)
		if err != nil {
			return nil, err
		}
	}

	marketplaces, err := accountmgmtutil.GetValidMarketplaces(f.Context, f.Connection)
	if err != nil {
		return nil, err
	}

	if len(marketplaces) > 0 {
		if err = promptMarketplaceSelect(f.Localizer, marketplaces, answers); err != nil {
			return nil, err
		}
	}

	marketplaceAcctIDs, err := accountmgmtutil.GetValidMarketplaceAcctIDs(f.Context, f.Connection)
	if err != nil {
		return nil, err
	}

	if len(marketplaceAcctIDs) > 0 {
		if err = promptMarketplaceAcctIDSelect(f.Localizer, marketplaceAcctIDs, answers); err != nil {
			return nil, err
		}
	}

	accountIDNullable := kafkamgmtclient.NullableString{}
	accountIDNullable.Set(&answers.MarketplaceAcctID)

	cloudProviderIDNullable := kafkamgmtclient.NullableString{}
	cloudProviderIDNullable.Set(&answers.Marketplace)

	payload := &kafkamgmtclient.KafkaRequestPayload{
		Name:                  answers.Name,
		Region:                &answers.Region,
		CloudProvider:         &answers.CloudProvider,
		BillingCloudAccountId: accountIDNullable,
		Marketplace:           cloudProviderIDNullable,
	}
	printSizeWarningIfNeeded(opts.f, answers.Size, sizes)
	payload.SetPlan(mapAmsTypeToBackendType(&userQuotaType) + "." + answers.Size)

	return payload, nil
}

func promptMarketplaceSelect(localizer localize.Localizer, marketplaceAcctIDs []string, answers *promptAnswers) error {

	marketplacePrompt := &survey.Select{
		Message: localizer.MustLocalize("kafka.create.input.marketPlace.message"),
		Options: marketplaceAcctIDs,
	}

	if err := survey.AskOne(marketplacePrompt, &answers.Marketplace); err != nil {
		return err
	}

	return nil
}

func promptMarketplaceAcctIDSelect(localizer localize.Localizer, accountIDs []string, answers *promptAnswers) error {

	accountIDPrompt := &survey.Select{
		Message: localizer.MustLocalize("kafka.create.input.accountID.message"),
		Options: accountIDs,
	}

	if err := survey.AskOne(accountIDPrompt, &answers.MarketplaceAcctID); err != nil {
		return err
	}

	return nil
}

func printSizeWarningIfNeeded(f *factory.Factory, selectedSize string, sizes []kafkamgmtclient.SupportedKafkaSize) {
	for i := range sizes {
		if sizes[i].GetId() == selectedSize {
			f.Logger.Info(f.Localizer.MustLocalize("kafka.create.log.info.sizeUnit",
				localize.NewEntry("DisplaySize", sizes[i].GetDisplayName()), localize.NewEntry("Size", sizes[i].GetId())))
			if sizes[i].GetMaturityStatus() == "preview" {
				f.Logger.Info(f.Localizer.MustLocalize("kafka.create.log.info.sizePreview"))
			}
		}
	}
}
