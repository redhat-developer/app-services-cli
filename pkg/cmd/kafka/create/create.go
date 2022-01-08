package create

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"time"

	kafkaFlagutil "github.com/redhat-developer/app-services-cli/pkg/cmd/kafka/flagutil"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/kafka/kafkacmdutil"
	"github.com/redhat-developer/app-services-cli/pkg/svcstatus"

	"github.com/redhat-developer/app-services-cli/pkg/accountmgmtutil"
	"github.com/redhat-developer/app-services-cli/pkg/core/cmdutil"
	"github.com/redhat-developer/app-services-cli/pkg/core/cmdutil/factory"
	"github.com/redhat-developer/app-services-cli/pkg/core/cmdutil/flagutil"
	"github.com/redhat-developer/app-services-cli/pkg/core/config"
	"github.com/redhat-developer/app-services-cli/pkg/core/connection"
	"github.com/redhat-developer/app-services-cli/pkg/core/ioutil/color"
	"github.com/redhat-developer/app-services-cli/pkg/core/ioutil/dump"
	"github.com/redhat-developer/app-services-cli/pkg/core/ioutil/icon"
	"github.com/redhat-developer/app-services-cli/pkg/core/ioutil/iostreams"
	"github.com/redhat-developer/app-services-cli/pkg/core/ioutil/spinner"
	"github.com/redhat-developer/app-services-cli/pkg/core/localize"
	"github.com/redhat-developer/app-services-cli/pkg/core/logging"
	pkgKafka "github.com/redhat-developer/app-services-cli/pkg/kafkautil"
	kafkamgmtclient "github.com/redhat-developer/app-services-sdk-go/kafkamgmt/apiv1/client"

	"github.com/AlecAivazis/survey/v2"

	"github.com/spf13/cobra"
)

type options struct {
	name     string
	provider string
	region   string
	multiAZ  bool

	outputFormat string
	autoUse      bool

	interactive      bool
	wait             bool
	bypassTermsCheck bool

	IO         *iostreams.IOStreams
	Config     config.IConfig
	Connection factory.ConnectionFunc
	Logger     logging.Logger
	localizer  localize.Localizer
	Context    context.Context
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
		IO:         f.IOStreams,
		Config:     f.Config,
		Connection: f.Connection,
		Logger:     f.Logger,
		localizer:  f.Localizer,
		Context:    f.Context,

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
	flags.AddOutput(&opts.outputFormat)
	flags.BoolVar(&opts.autoUse, "use", true, opts.localizer.MustLocalize("kafka.create.flag.autoUse.description"))
	flags.BoolVarP(&opts.wait, "wait", "w", false, opts.localizer.MustLocalize("kafka.create.flag.wait.description"))
	flags.AddBypassTermsCheck(&opts.bypassTermsCheck)

	_ = cmd.RegisterFlagCompletionFunc(kafkaFlagutil.FlagProvider, func(cmd *cobra.Command, _ []string, _ string) ([]string, cobra.ShellCompDirective) {
		return pkgKafka.GetCloudProviderCompletionValues(f)
	})

	return cmd
}

// nolint:funlen
func runCreate(opts *options) error {
	cfg, err := opts.Config.Load()
	if err != nil {
		return err
	}

	var conn connection.Connection
	if conn, err = opts.Connection(connection.DefaultConfigSkipMasAuth); err != nil {
		return err
	}

	if !opts.bypassTermsCheck {
		opts.Logger.Debug("Checking if terms and conditions have been accepted")
		// the user must have accepted the terms and conditions from the provider
		// before they can create a kafka instance
		termsSpec := accountmgmtutil.GetRemoteTermsSpec(&opts.Context, opts.Logger)
		var termsAccepted bool
		var termsURL string
		termsAccepted, termsURL, err = accountmgmtutil.CheckTermsAccepted(opts.Context, termsSpec.Kafka, conn)
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

	if apiErr := pkgKafka.GetAPIError(err); apiErr != nil {
		switch apiErr.GetCode() {
		case pkgKafka.ErrorCode24:
			return opts.localizer.MustLocalizeError("kafka.create.error.oneinstance")
		case pkgKafka.ErrorCode36:
			return opts.localizer.MustLocalizeError("kafka.create.error.conflictError", localize.NewEntry("Name", payload.Name))
		}
	}

	if err != nil {
		return err
	}

	kafkaCfg := &config.KafkaConfig{
		ClusterID: response.GetId(),
	}

	if opts.autoUse {
		opts.Logger.Debug("Auto-use is set, updating the current instance")
		cfg.Services.Kafka = kafkaCfg
		if err = opts.Config.Save(cfg); err != nil {
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

	// set type to store the answers from the prompt with defaults
	answers := struct {
		Name          string
		Region        string
		MultiAZ       bool
		CloudProvider string
	}{
		MultiAZ: defaultMultiAZ,
	}

	promptName := &survey.Input{
		Message: opts.localizer.MustLocalize("kafka.create.input.name.message"),
		Help:    opts.localizer.MustLocalize("kafka.create.input.name.help"),
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

	regions := cloudRegionResponse.GetItems()
	regionIDs := pkgKafka.GetEnabledCloudRegionIDs(regions)

	regionPrompt := &survey.Select{
		Message: opts.localizer.MustLocalize("kafka.create.input.cloudRegion.message"),
		Options: regionIDs,
		Help:    opts.localizer.MustLocalize("kafka.create.input.cloudRegion.help"),
	}

	err = survey.AskOne(regionPrompt, &answers.Region)
	if err != nil {
		return nil, err
	}

	payload = &kafkamgmtclient.KafkaRequestPayload{
		Name:          answers.Name,
		Region:        &answers.Region,
		CloudProvider: &answers.CloudProvider,
		MultiAz:       &answers.MultiAZ,
	}

	return payload, nil
}
