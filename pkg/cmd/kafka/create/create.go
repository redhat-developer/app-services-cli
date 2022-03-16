package create

import (
	"context"
	"errors"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"time"

	kafkaFlagutil "github.com/redhat-developer/app-services-cli/pkg/cmd/kafka/flagutil"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/kafka/kafkacmdutil"
	"github.com/redhat-developer/app-services-cli/pkg/shared/accountmgmtutil"
	"github.com/redhat-developer/app-services-cli/pkg/shared/contextutil"
	"github.com/redhat-developer/app-services-cli/pkg/shared/remote"
	"github.com/redhat-developer/app-services-cli/pkg/shared/svcstatus"
	"k8s.io/utils/strings/slices"

	"github.com/redhat-developer/app-services-cli/pkg/core/cmdutil"
	"github.com/redhat-developer/app-services-cli/pkg/core/cmdutil/flagutil"
	"github.com/redhat-developer/app-services-cli/pkg/core/ioutil/color"
	"github.com/redhat-developer/app-services-cli/pkg/core/ioutil/dump"
	"github.com/redhat-developer/app-services-cli/pkg/core/ioutil/icon"
	"github.com/redhat-developer/app-services-cli/pkg/core/ioutil/iostreams"
	"github.com/redhat-developer/app-services-cli/pkg/core/ioutil/spinner"
	"github.com/redhat-developer/app-services-cli/pkg/core/localize"
	"github.com/redhat-developer/app-services-cli/pkg/core/logging"
	"github.com/redhat-developer/app-services-cli/pkg/core/servicecontext"
	"github.com/redhat-developer/app-services-cli/pkg/shared/connection"
	"github.com/redhat-developer/app-services-cli/pkg/shared/factory"
	"github.com/redhat-developer/app-services-cli/pkg/shared/kafkautil"
	pkgKafka "github.com/redhat-developer/app-services-cli/pkg/shared/kafkautil"

	kafkamgmtclient "github.com/redhat-developer/app-services-sdk-go/kafkamgmt/apiv1/client"
	kafkamgmtv1errors "github.com/redhat-developer/app-services-sdk-go/kafkamgmt/apiv1/error"

	"github.com/AlecAivazis/survey/v2"

	"github.com/spf13/cobra"
)

type options struct {
	name     string
	provider string
	region   string
	size     string

	multiAZ bool

	outputFormat string
	autoUse      bool

	interactive    bool
	wait           bool
	bypassAmsCheck bool

	IO                *iostreams.IOStreams
	Connection        factory.ConnectionFunc
	Logger            logging.Logger
	localizer         localize.Localizer
	Context           context.Context
	ServiceContext    servicecontext.IContext
	userInstanceTypes []string
}

const (
	// default Kafka instance values
	defaultMultiAZ  = true
	defaultRegion   = "us-east-1"
	defaultProvider = "aws"
)

// NewCreateCommand creates a new command for creating kafkas.
func NewCreateCommand(f *factory.Factory) *cobra.Command {
	opts := &options{
		IO:             f.IOStreams,
		Connection:     f.Connection,
		Logger:         f.Logger,
		localizer:      f.Localizer,
		Context:        f.Context,
		ServiceContext: f.ServiceContext,

		multiAZ: defaultMultiAZ,
	}

	cmd := &cobra.Command{
		Use:     "create",
		Short:   opts.localizer.MustLocalize("kafka.create.cmd.shortDescription"),
		Long:    opts.localizer.MustLocalize("kafka.create.cmd.longDescription"),
		Example: opts.localizer.MustLocalize("kafka.create.cmd.example"),
		Args:    cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			if opts.name != "" {
				validator := &kafkacmdutil.Validator{
					Localizer:  opts.localizer,
					Connection: opts.Connection,
				}
				if err := validator.ValidateName(opts.name); err != nil {
					return err
				}
			}

			if !opts.IO.CanPrompt() && opts.name == "" {
				return opts.localizer.MustLocalizeError("kafka.create.argument.name.error.requiredWhenNonInteractive")
			} else if opts.name == "" {
				if opts.provider != "" || opts.region != "" {
					return opts.localizer.MustLocalizeError("kafka.create.argument.name.error.requiredWhenNonInteractive")
				}
				opts.interactive = true
			}

			validOutputFormats := flagutil.ValidOutputFormats
			if opts.outputFormat != "" && !flagutil.IsValidInput(opts.outputFormat, validOutputFormats...) {
				return flagutil.InvalidValueError("output", opts.outputFormat, validOutputFormats...)
			}

			return runCreate(opts)
		},
	}

	flags := kafkaFlagutil.NewFlagSet(cmd, opts.localizer)

	flags.StringVar(&opts.name, "name", "", opts.localizer.MustLocalize("kafka.create.flag.name.description"))
	flags.StringVar(&opts.provider, kafkaFlagutil.FlagProvider, "", opts.localizer.MustLocalize("kafka.create.flag.cloudProvider.description"))
	flags.StringVar(&opts.region, kafkaFlagutil.FlagRegion, "", opts.localizer.MustLocalize("kafka.create.flag.cloudRegion.description"))
	flags.StringVar(&opts.size, "size", "", opts.localizer.MustLocalize("kafka.create.flag.size.description"))
	flags.AddOutput(&opts.outputFormat)
	flags.BoolVar(&opts.autoUse, "use", true, opts.localizer.MustLocalize("kafka.create.flag.autoUse.description"))
	flags.BoolVarP(&opts.wait, "wait", "w", false, opts.localizer.MustLocalize("kafka.create.flag.wait.description"))
	flags.AddBypassTermsCheck(&opts.bypassAmsCheck)

	_ = cmd.RegisterFlagCompletionFunc(kafkaFlagutil.FlagProvider, func(cmd *cobra.Command, _ []string, _ string) ([]string, cobra.ShellCompDirective) {
		return pkgKafka.GetCloudProviderCompletionValues(f)
	})

	_ = cmd.RegisterFlagCompletionFunc(kafkaFlagutil.FlagRegion, func(cmd *cobra.Command, _ []string, _ string) ([]string, cobra.ShellCompDirective) {
		return kafkautil.GetCloudProviderRegionCompletionValues(f, opts.provider)
	})

	_ = cmd.RegisterFlagCompletionFunc(kafkaFlagutil.FlagRegion, func(cmd *cobra.Command, _ []string, _ string) ([]string, cobra.ShellCompDirective) {
		return kafkautil.GetCloudProviderSizeValues(f, opts.provider, opts.region)
	})

	return cmd
}

// nolint:funlen
func runCreate(opts *options) error {
	svcContext, err := opts.ServiceContext.Load()
	if err != nil {
		return err
	}

	currCtx, err := contextutil.GetCurrentContext(svcContext, opts.localizer)
	if err != nil {
		return err
	}

	var conn connection.Connection
	if conn, err = opts.Connection(connection.DefaultConfigSkipMasAuth); err != nil {
		return err
	}

	err, constants := remote.GetRemoteServiceConstants(opts.Context, opts.Logger)
	if err != nil {
		return err
	}

	if !opts.bypassAmsCheck {
		opts.Logger.Debug("Checking if terms and conditions have been accepted")
		// the user must have accepted the terms and conditions from the provider
		// before they can create a kafka instance
		var termsAccepted bool
		var termsURL string
		termsAccepted, termsURL, err = accountmgmtutil.CheckTermsAccepted(opts.Context, constants.Kafka.Ams, conn)

		if err != nil {
			return err
		}
		if !termsAccepted && termsURL != "" {
			opts.Logger.Info(opts.localizer.MustLocalize("service.info.termsCheck", localize.NewEntry("TermsURL", termsURL)))
			return nil
		}
	}

	var payload *kafkamgmtclient.KafkaRequestPayload
	if opts.interactive {
		opts.Logger.Debug()

		payload, err = promptKafkaPayload(opts)
		if err != nil {
			return err
		}

		opts.provider = payload.GetCloudProvider()
		opts.region = payload.GetRegion()
	} else {
		if opts.provider == "" {
			opts.provider = defaultProvider
		}
		if opts.region == "" {
			opts.region = defaultRegion
		}

		if !opts.bypassAmsCheck {
			opts.userInstanceTypes, err = accountmgmtutil.GetUserSupportedInstanceTypes(opts.Context, constants.Kafka.Ams, conn)
			if err != nil {
				opts.Logger.Debug("Cannot retrieve user supported instance types. Skipping validation", err)
				return err
			}

			err = validateProviderAndRegion(opts, constants, conn)
			if err != nil {
				return err
			}

			err = validateSize(opts, constants, conn)
			if err != nil {
				return err
			}
		}

		payload = &kafkamgmtclient.KafkaRequestPayload{
			Name:          opts.name,
			Region:        &opts.region,
			CloudProvider: &opts.provider,
			Plan:          &opts.size,
			MultiAz:       &opts.multiAZ,
		}
	}

	api := conn.API()

	a := api.KafkaMgmt().CreateKafka(opts.Context)
	a = a.KafkaRequestPayload(*payload)
	a = a.Async(true)

	response, httpRes, err := a.Execute()
	if httpRes != nil {
		defer httpRes.Body.Close()
	}

	if apiErr := kafkamgmtv1errors.GetAPIError(err); apiErr != nil {
		switch apiErr.GetCode() {
		case kafkamgmtv1errors.ERROR_120:
			return opts.localizer.MustLocalizeError("kafka.create.error.oneinstance")
		case kafkamgmtv1errors.ERROR_24:
			return opts.localizer.MustLocalizeError("kafka.create.error.temporary.unavailable")
		case kafkamgmtv1errors.ERROR_36:
			return opts.localizer.MustLocalizeError("kafka.create.error.conflictError", localize.NewEntry("Name", payload.Name))
		case kafkamgmtv1errors.ERROR_41:
			return opts.localizer.MustLocalizeError("kafka.create.error.notsupported", localize.NewEntry("Name", payload.Name))
		}
	}

	if err != nil {
		return err
	}

	if opts.autoUse {
		opts.Logger.Debug("Auto-use is set, updating the current instance")
		currCtx.KafkaID = response.GetId()
		svcContext.Contexts[svcContext.CurrentContext] = *currCtx

		if err = opts.ServiceContext.Save(svcContext); err != nil {
			return fmt.Errorf("%v: %w", opts.localizer.MustLocalize("kafka.common.error.couldNotUseKafka"), err)
		}
	} else {
		opts.Logger.Debug("Auto-use is not set, skipping updating the current instance")
	}

	nameTemplateEntry := localize.NewEntry("Name", response.GetName())

	if opts.wait {
		opts.Logger.Debug("--wait flag is enabled, waiting for Kafka to finish creating")
		s := spinner.New(opts.IO.ErrOut, opts.localizer)
		s.SetLocalizedSuffix("kafka.create.log.info.creatingKafka", nameTemplateEntry)
		s.Start()

		// when there is a SIGINT, display a message informing the user that this does not cancel the creation
		// and that it is being created in the background
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt)
		go func() {
			for range c {
				opts.Logger.Info()
				opts.Logger.Info(opts.localizer.MustLocalize("kafka.create.log.info.creatingKafkaSyncSigint"))
				os.Exit(0)
			}
		}()

		for svcstatus.IsInstanceCreating(response.GetStatus()) {
			time.Sleep(cmdutil.DefaultPollTime)

			response, httpRes, err = api.KafkaMgmt().GetKafkaById(opts.Context, response.GetId()).Execute()
			if err != nil {
				return err
			}
			defer httpRes.Body.Close()
			opts.Logger.Debug("Checking Kafka status:", response.GetStatus())

			s.SetLocalizedSuffix("kafka.create.log.info.creationInProgress",
				localize.NewEntry("Name", response.GetName()),
				localize.NewEntry("Status", color.Info(response.GetStatus())),
			)

		}
		s.Stop()
		opts.Logger.Info()
		opts.Logger.Info(icon.SuccessPrefix(), opts.localizer.MustLocalize("kafka.create.info.successSync", nameTemplateEntry))
	}

	if err = dump.Formatted(opts.IO.Out, opts.outputFormat, response); err != nil {
		return err
	}

	if !opts.wait {
		opts.Logger.Info()
		opts.Logger.Info(opts.localizer.MustLocalize("kafka.create.info.successAsync", nameTemplateEntry))
	}

	return nil
}

func validateSize(opts *options, constants *remote.DynamicServiceConstants, conn connection.Connection) error {
	amsType, err := accountmgmtutil.PickInstanceType(&opts.userInstanceTypes)
	if err != nil {
		return err
	}
	sizes, err := kafkautil.GetValidSizes(conn, opts.Context, opts.provider, opts.region, &amsType)
	if err != nil {
		return err
	}
	if !slices.Contains(sizes, opts.size) {
		// TODO error message
		return errors.New("Whatever") //opts.localizer.MustLocalizeError("")
	}
	return nil
}

func validateProviderAndRegion(opts *options, constants *remote.DynamicServiceConstants, conn connection.Connection) error {
	opts.Logger.Debug("Validating provider and region")
	cloudProviders, _, err := conn.API().
		KafkaMgmt().
		GetCloudProviders(opts.Context).
		Execute()

	if err != nil {
		return err
	}

	var selectedProvider kafkamgmtclient.CloudProvider

	providerNames := make([]string, 0)
	for _, item := range cloudProviders.Items {
		if !item.GetEnabled() {
			continue
		}
		if item.GetId() == opts.provider {
			selectedProvider = item
		}
		providerNames = append(providerNames, item.GetId())
	}
	opts.Logger.Debug("Validating cloud provider", opts.provider, ". Enabled providers: ", providerNames)

	if !selectedProvider.Enabled {
		providers := strings.Join(providerNames, ",")
		providerEntry := localize.NewEntry("Provider", opts.provider)
		validProvidersEntry := localize.NewEntry("Providers", providers)
		return opts.localizer.MustLocalizeError("kafka.create.provider.error.invalidProvider", providerEntry, validProvidersEntry)
	}
	// Temporary disabled due to breaking changes in the API
	return nil // validateProviderRegion(conn, opts, selectedProvider, constants)
}

func ValidateProviderRegion(conn connection.Connection, opts *options, selectedProvider kafkamgmtclient.CloudProvider, constants *remote.DynamicServiceConstants) error {
	cloudRegion, _, err := conn.API().
		KafkaMgmt().
		GetCloudProviderRegions(opts.Context, selectedProvider.GetId()).
		Execute()

	if err != nil {
		return err
	}

	var selectedRegion kafkamgmtclient.CloudRegion
	regionNames := make([]string, 0)
	for _, item := range cloudRegion.Items {
		if !item.GetEnabled() {
			continue
		}
		regionNames = append(regionNames, item.GetId())
		if item.GetId() == opts.region {
			selectedRegion = item
		}
	}

	if len(regionNames) != 0 {
		opts.Logger.Debug("Validating region", opts.region, ". Enabled providers: ", regionNames)
		regionsString := strings.Join(regionNames, ", ")
		if !selectedRegion.Enabled {
			regionEntry := localize.NewEntry("Region", opts.region)
			validRegionsEntry := localize.NewEntry("Regions", regionsString)
			providerEntry := localize.NewEntry("Provider", opts.provider)
			return opts.localizer.MustLocalizeError("kafka.create.region.error.invalidRegion", regionEntry, providerEntry, validRegionsEntry)
		}

		if err != nil {
			return err
		}

		regionInstanceTypes := selectedRegion.GetSupportedInstanceTypes()

		for _, item := range regionInstanceTypes {
			if slices.Contains(opts.userInstanceTypes, item) {
				return nil
			}
		}

		regionEntry := localize.NewEntry("Region", opts.region)
		userTypesEntry := localize.NewEntry("MyTypes", strings.Join(opts.userInstanceTypes, ", "))
		cloudTypesEntry := localize.NewEntry("CloudTypes", strings.Join(regionInstanceTypes, ", "))

		return opts.localizer.MustLocalizeError("kafka.create.region.error.regionNotSupported", regionEntry, userTypesEntry, cloudTypesEntry)

	}
	opts.Logger.Debug("No regions found for provider. Skipping provider validation", opts.provider)

	return nil
}

// set type to store the answers from the prompt with defaults
type promptAnswers struct {
	Name          string
	Plan          string
	Region        string
	MultiAZ       bool
	CloudProvider string
}

// Show a prompt to allow the user to interactively insert the data for their Kafka
func promptKafkaPayload(opts *options) (payload *kafkamgmtclient.KafkaRequestPayload, err error) {
	conn, err := opts.Connection(connection.DefaultConfigSkipMasAuth)
	if err != nil {
		return nil, err
	}

	api := conn.API()

	validator := &kafkacmdutil.Validator{
		Localizer:  opts.localizer,
		Connection: opts.Connection,
	}

	promptName := &survey.Input{
		Message: opts.localizer.MustLocalize("kafka.create.input.name.message"),
		Help:    opts.localizer.MustLocalize("kafka.create.input.name.help"),
	}

	answers := &promptAnswers{
		MultiAZ: defaultMultiAZ,
	}

	err = survey.AskOne(promptName, &answers.Name, survey.WithValidator(validator.ValidateName), survey.WithValidator(validator.ValidateNameIsAvailable))
	if err != nil {
		return nil, err
	}

	// fetch all cloud available providers
	cloudProviderResponse, httpRes, err := api.KafkaMgmt().GetCloudProviders(opts.Context).Execute()
	if httpRes != nil {
		defer httpRes.Body.Close()
	}

	if err != nil {
		return nil, err
	}

	cloudProviders := cloudProviderResponse.GetItems()
	cloudProviderNames := pkgKafka.GetEnabledCloudProviderNames(cloudProviders)

	cloudProviderPrompt := &survey.Select{
		Message: opts.localizer.MustLocalize("kafka.create.input.cloudProvider.message"),
		Options: cloudProviderNames,
	}

	err = survey.AskOne(cloudProviderPrompt, &answers.CloudProvider)
	if err != nil {
		return nil, err
	}

	// get the selected provider type from the name selected
	selectedCloudProvider := pkgKafka.FindCloudProviderByName(cloudProviders, answers.CloudProvider)

	// nolint
	cloudRegionResponse, _, err := api.KafkaMgmt().GetCloudProviderRegions(opts.Context, selectedCloudProvider.GetId()).Execute()
	if err != nil {
		return nil, err
	}

	// Temporary disabled due to breaking changes in the API
	// userInstanceTypes, err := accountmgmtutil.GetUserSupportedInstanceTypes(opts.Context, constants.Kafka.Ams, conn)
	// if err != nil {
	// 	opts.Logger.Debug("Cannot retrieve user supported instance types. Skipping validation", err)
	// 	return payload, err
	// }

	regions := cloudRegionResponse.GetItems()
	regionIDs := pkgKafka.GetEnabledCloudRegionIDs(regions, nil)

	regionPrompt := &survey.Select{
		Message: opts.localizer.MustLocalize("kafka.create.input.cloudRegion.message"),
		Options: regionIDs,
		Help:    opts.localizer.MustLocalize("kafka.create.input.cloudRegion.help"),
	}

	err = survey.AskOne(regionPrompt, &answers.Region)
	if err != nil {
		return nil, err
	}

	amsType, err := accountmgmtutil.PickInstanceType(&opts.userInstanceTypes)
	if err != nil {
		return nil, err
	}
	sizes, err := kafkautil.GetValidSizes(conn, opts.Context, opts.provider, opts.region, &amsType)
	if err != nil {
		return nil, err
	}

	planPrompt := &survey.Select{
		Message: "What type of instance we should create",
		Options: sizes,
		Help:    "",
	}

	err = survey.AskOne(planPrompt, &answers.Plan)
	if err != nil {
		return nil, err
	}

	payload = &kafkamgmtclient.KafkaRequestPayload{
		Name:          answers.Name,
		Region:        &answers.Region,
		CloudProvider: &answers.CloudProvider,
		Plan:          &answers.Plan,
		MultiAz:       &answers.MultiAZ,
	}

	return payload, nil
}
