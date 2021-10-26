package create

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/redhat-developer/app-services-cli/pkg/api/kas"
	"github.com/redhat-developer/app-services-cli/pkg/icon"
	"github.com/redhat-developer/app-services-cli/pkg/ioutil/spinner"

	kafkamgmtclient "github.com/redhat-developer/app-services-sdk-go/kafkamgmt/apiv1/client"

	"github.com/redhat-developer/app-services-cli/pkg/color"
	"github.com/redhat-developer/app-services-cli/pkg/dump"
	kafkacmdutil "github.com/redhat-developer/app-services-cli/pkg/kafka/cmdutil"
	"github.com/redhat-developer/app-services-cli/pkg/localize"

	"github.com/redhat-developer/app-services-cli/pkg/ams"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/flag"
	cmdFlagUtil "github.com/redhat-developer/app-services-cli/pkg/cmdutil/flagutil"
	"github.com/redhat-developer/app-services-cli/pkg/connection"
	svcstatus "github.com/redhat-developer/app-services-cli/pkg/service/status"

	"github.com/AlecAivazis/survey/v2"

	"github.com/redhat-developer/app-services-cli/pkg/iostreams"
	pkgKafka "github.com/redhat-developer/app-services-cli/pkg/kafka"
	"github.com/redhat-developer/app-services-cli/pkg/logging"

	"github.com/spf13/cobra"

	"github.com/redhat-developer/app-services-cli/internal/config"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/factory"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/kafka/flagutil"
	"github.com/redhat-developer/app-services-cli/pkg/cmdutil"
)

type options struct {
	name     string
	provider string
	region   string
	multiAZ  bool

	outputFormat string
	autoUse      bool

	interactive bool
	wait        bool

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
				validator := &pkgKafka.Validator{
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

			validOutputFormats := cmdFlagUtil.ValidOutputFormats
			if opts.outputFormat != "" && !cmdFlagUtil.IsValidInput(opts.outputFormat, validOutputFormats...) {
				return flag.InvalidValueError("output", opts.outputFormat, validOutputFormats...)
			}

			return runCreate(opts)
		},
	}

	flags := flagutil.NewFlagSet(cmd, opts.localizer)

	flags.StringVar(&opts.name, "name", "", opts.localizer.MustLocalize("kafka.create.flag.name.description"))
	flags.StringVar(&opts.provider, flagutil.FlagProvider, "", opts.localizer.MustLocalize("kafka.create.flag.cloudProvider.description"))
	flags.StringVar(&opts.region, flagutil.FlagRegion, "", opts.localizer.MustLocalize("kafka.create.flag.cloudRegion.description"))
	flags.AddOutput(&opts.outputFormat)
	flags.BoolVar(&opts.autoUse, "use", true, opts.localizer.MustLocalize("kafka.create.flag.autoUse.description"))
	flags.BoolVarP(&opts.wait, "wait", "w", false, opts.localizer.MustLocalize("kafka.create.flag.wait.description"))

	_ = cmd.RegisterFlagCompletionFunc(flagutil.FlagProvider, func(cmd *cobra.Command, _ []string, _ string) ([]string, cobra.ShellCompDirective) {
		return kafkacmdutil.GetCloudProviderCompletionValues(f)
	})

	cmdFlagUtil.EnableOutputFlagCompletion(cmd)

	return cmd
}

// nolint:funlen
func runCreate(opts *options) error {
	cfg, err := opts.Config.Load()
	if err != nil {
		return err
	}

	conn, err := opts.Connection(connection.DefaultConfigSkipMasAuth)
	if err != nil {
		return err
	}

	// the user must have accepted the terms and conditions from the provider
	// before they can create a kafka instance
	termsAccepted, termsURL, err := ams.CheckTermsAccepted(opts.Context, conn)
	if err != nil {
		return err
	}
	if !termsAccepted && termsURL != "" {
		opts.Logger.Info(opts.localizer.MustLocalize("service.info.termsCheck", localize.NewEntry("TermsURL", termsURL)))
		return nil
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

	a := api.Kafka().CreateKafka(opts.Context)
	a = a.KafkaRequestPayload(*payload)
	a = a.Async(true)
	response, httpRes, err := a.Execute()
	if httpRes != nil {
		defer httpRes.Body.Close()
	}

	if apiErr := kas.GetAPIError(err); apiErr != nil {
		switch apiErr.GetCode() {
		case kas.ErrorCode24:
			return opts.localizer.MustLocalizeError("kafka.create.error.oneinstance")
		case kas.ErrorCode36:
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

		for svcstatus.IsCreating(response.GetStatus()) {
			time.Sleep(cmdutil.DefaultPollTime)

			response, httpRes, err = api.Kafka().GetKafkaById(opts.Context, response.GetId()).Execute()
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

	validator := &pkgKafka.Validator{
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
	cloudProviderResponse, httpRes, err := api.Kafka().GetCloudProviders(opts.Context).Execute()
	if httpRes != nil {
		defer httpRes.Body.Close()
	}

	if err != nil {
		return nil, err
	}

	cloudProviders := cloudProviderResponse.GetItems()
	cloudProviderNames := kafkacmdutil.GetEnabledCloudProviderNames(cloudProviders)

	cloudProviderPrompt := &survey.Select{
		Message: opts.localizer.MustLocalize("kafka.create.input.cloudProvider.message"),
		Options: cloudProviderNames,
	}

	err = survey.AskOne(cloudProviderPrompt, &answers.CloudProvider)
	if err != nil {
		return nil, err
	}

	// get the selected provider type from the name selected
	selectedCloudProvider := kafkacmdutil.FindCloudProviderByName(cloudProviders, answers.CloudProvider)

	// nolint
	cloudRegionResponse, _, err := api.Kafka().GetCloudProviderRegions(opts.Context, selectedCloudProvider.GetId()).Execute()
	if err != nil {
		return nil, err
	}

	regions := cloudRegionResponse.GetItems()
	regionIDs := kafkacmdutil.GetEnabledCloudRegionIDs(regions)

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
