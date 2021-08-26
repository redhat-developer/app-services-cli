package create

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	kafkamgmtclient "github.com/redhat-developer/app-services-sdk-go/kafkamgmt/apiv1/client"
	"gopkg.in/yaml.v2"

	"github.com/redhat-developer/app-services-cli/pkg/color"
	"github.com/redhat-developer/app-services-cli/pkg/dump"
	kafkacmdutil "github.com/redhat-developer/app-services-cli/pkg/kafka/cmdutil"
	"github.com/redhat-developer/app-services-cli/pkg/localize"

	"github.com/redhat-developer/app-services-cli/pkg/ams"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/flag"
	flagutil "github.com/redhat-developer/app-services-cli/pkg/cmdutil/flags"
	"github.com/redhat-developer/app-services-cli/pkg/connection"
	svcstatus "github.com/redhat-developer/app-services-cli/pkg/service/status"

	"github.com/AlecAivazis/survey/v2"
	"github.com/redhat-developer/app-services-cli/pkg/iostreams"
	pkgKafka "github.com/redhat-developer/app-services-cli/pkg/kafka"
	"github.com/redhat-developer/app-services-cli/pkg/logging"

	"github.com/spf13/cobra"

	"github.com/redhat-developer/app-services-cli/internal/config"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/factory"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/kafka/flags"
	"github.com/redhat-developer/app-services-cli/pkg/cmdutil"
)

type Options struct {
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
	Logger     func() (logging.Logger, error)
	localizer  localize.Localizer
}

const (
	// default Kafka instance values
	defaultMultiAZ  = true
	defaultRegion   = "us-east-1"
	defaultProvider = "aws"
)

// NewCreateCommand creates a new command for creating kafkas.
func NewCreateCommand(f *factory.Factory) *cobra.Command {
	opts := &Options{
		IO:         f.IOStreams,
		Config:     f.Config,
		Connection: f.Connection,
		Logger:     f.Logger,
		localizer:  f.Localizer,

		multiAZ: defaultMultiAZ,
	}

	cmd := &cobra.Command{
		Use:     opts.localizer.MustLocalize("kafka.create.cmd.use"),
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
				return errors.New(opts.localizer.MustLocalize("kafka.create.argument.name.error.requiredWhenNonInteractive"))
			} else if opts.name == "" {
				if opts.provider != "" || opts.region != "" {
					return errors.New(opts.localizer.MustLocalize("kafka.create.argument.name.error.requiredWhenNonInteractive"))
				}
				opts.interactive = true
			}

			validOutputFormats := flagutil.ValidOutputFormats
			if opts.outputFormat != "" && !flagutil.IsValidInput(opts.outputFormat, validOutputFormats...) {
				return flag.InvalidValueError("output", opts.outputFormat, validOutputFormats...)
			}

			return runCreate(opts)
		},
	}

	cmd.Flags().StringVar(&opts.name, "name", "", opts.localizer.MustLocalize("kafka.create.flag.name.description"))
	cmd.Flags().StringVar(&opts.provider, flags.FlagProvider, "", opts.localizer.MustLocalize("kafka.create.flag.cloudProvider.description"))
	cmd.Flags().StringVar(&opts.region, flags.FlagRegion, "", opts.localizer.MustLocalize("kafka.create.flag.cloudRegion.description"))
	cmd.Flags().StringVarP(&opts.outputFormat, "output", "o", "json", opts.localizer.MustLocalize("kafka.common.flag.output.description"))
	cmd.Flags().BoolVar(&opts.autoUse, "use", true, opts.localizer.MustLocalize("kafka.create.flag.autoUse.description"))
	cmd.Flags().BoolVarP(&opts.wait, "wait", "w", false, opts.localizer.MustLocalize("kafka.create.flag.wait.description"))

	_ = cmd.RegisterFlagCompletionFunc(flags.FlagProvider, func(cmd *cobra.Command, _ []string, _ string) ([]string, cobra.ShellCompDirective) {
		return kafkacmdutil.GetCloudProviderCompletionValues(f)
	})

	flagutil.EnableOutputFlagCompletion(cmd)

	return cmd
}

// nolint:funlen
func runCreate(opts *Options) error {
	logger, err := opts.Logger()
	if err != nil {
		return err
	}

	cfg, err := opts.Config.Load()
	if err != nil {
		return err
	}

	connection, err := opts.Connection(connection.DefaultConfigSkipMasAuth)
	if err != nil {
		return err
	}

	// the user must have accepted the terms and conditions from the provider
	// before they can create a kafka instance
	termsAccepted, termsURL, err := ams.CheckTermsAccepted(connection)
	if err != nil {
		return err
	}
	if !termsAccepted && termsURL != "" {
		logger.Info(opts.localizer.MustLocalize("service.info.termsCheck", localize.NewEntry("TermsURL", termsURL)))
		return nil
	}

	var payload *kafkamgmtclient.KafkaRequestPayload
	if opts.interactive {
		logger.Debug()

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

	api := connection.API()

	a := api.Kafka().CreateKafka(context.Background())
	a = a.KafkaRequestPayload(*payload)
	a = a.Async(true)
	response, httpRes, err := a.Execute()
	defer httpRes.Body.Close()

	if httpRes.StatusCode == http.StatusBadRequest {
		return errors.New(opts.localizer.MustLocalize("kafka.create.error.conflictError", localize.NewEntry("Name", payload.Name)))
	}

	if err != nil {
		return err
	}

	kafkaCfg := &config.KafkaConfig{
		ClusterID: response.GetId(),
	}

	if opts.autoUse {
		logger.Debug("Auto-use is set, updating the current instance")
		cfg.Services.Kafka = kafkaCfg
		if err = opts.Config.Save(cfg); err != nil {
			return fmt.Errorf("%v: %w", opts.localizer.MustLocalize("kafka.common.error.couldNotUseKafka"), err)
		}
	} else {
		logger.Debug("Auto-use is not set, skipping updating the current instance")
	}

	nameTemplateEntry := localize.NewEntry("Name", response.GetName())

	if opts.wait {
		logger.Debug("--wait flag is enabled, waiting for Kafka to finish creating")
		s := opts.IO.NewSpinner()
		s.Suffix = fmt.Sprintf(" %v", opts.localizer.MustLocalize("kafka.create.log.info.creatingKafka", nameTemplateEntry))
		s.Start()

		// when there is a SIGINT, display a message informing the user that this does not cancel the creation
		// and that it is being created in the background
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt)
		go func() {
			for range c {
				logger.Info()
				logger.Info(opts.localizer.MustLocalize("kafka.create.log.info.creatingKafkaSyncSigint"))
				os.Exit(0)
			}
		}()

		for svcstatus.IsCreating(response.GetStatus()) {
			time.Sleep(cmdutil.DefaultPollTime)

			response, httpRes, err = api.Kafka().GetKafkaById(context.Background(), response.GetId()).Execute()
			if err != nil {
				return err
			}
			defer httpRes.Body.Close()
			logger.Debug("Checking Kafka status:", response.GetStatus())

			s.Suffix = createSpinnerSuffix(opts.localizer, response.GetName(), response.GetStatus())

		}
		s.Stop()
		logger.Info("\n")
		logger.Info(opts.localizer.MustLocalize("kafka.create.info.successSync", nameTemplateEntry))
	}

	switch opts.outputFormat {
	case dump.JSONFormat:
		data, _ := json.Marshal(response)
		_ = dump.JSON(opts.IO.Out, data)
	case dump.YAMLFormat, dump.YMLFormat:
		data, _ := yaml.Marshal(response)
		_ = dump.YAML(opts.IO.Out, data)
	}

	if !opts.wait {
		logger.Info()
		logger.Info(opts.localizer.MustLocalize("kafka.create.info.successAsync", nameTemplateEntry))
	}

	return nil
}

// Show a prompt to allow the user to interactively insert the data for their Kafka
func promptKafkaPayload(opts *Options) (payload *kafkamgmtclient.KafkaRequestPayload, err error) {
	connection, err := opts.Connection(connection.DefaultConfigSkipMasAuth)
	if err != nil {
		return nil, err
	}

	api := connection.API()

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
	cloudProviderResponse, httpRes, err := api.Kafka().GetCloudProviders(context.Background()).Execute()
	defer httpRes.Body.Close()
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
	cloudRegionResponse, _, err := api.Kafka().GetCloudProviderRegions(context.Background(), selectedCloudProvider.GetId()).Execute()
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

func createSpinnerSuffix(localizer localize.Localizer, name string, status string) string {
	return " " + localizer.MustLocalize("kafka.create.log.info.creationInProgress",
		localize.NewEntry("Name", name),
		localize.NewEntry("Status", color.Info(status)))
}
