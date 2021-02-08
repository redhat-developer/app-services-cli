package create

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/api/kas"
	kasclient "github.com/bf2fc6cc711aee1a0c2a/cli/pkg/api/kas/client"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/color"

	"github.com/MakeNowJust/heredoc"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/cloudprovider/cloudproviderutil"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/cloudregion/cloudregionutil"

	"github.com/AlecAivazis/survey/v2"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/dump"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/iostreams"

	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/connection"
	pkgKafka "github.com/bf2fc6cc711aee1a0c2a/cli/pkg/kafka"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/logging"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"

	"github.com/bf2fc6cc711aee1a0c2a/cli/internal/config"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/cmd/factory"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/cmd/kafka/flags"
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/cmdutil"
)

type Options struct {
	name     string
	provider string
	region   string
	multiAZ  bool

	outputFormat string

	interactive bool

	IO         *iostreams.IOStreams
	Config     config.IConfig
	Connection func() (connection.Connection, error)
	Logger     func() (logging.Logger, error)
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

		multiAZ: defaultMultiAZ,
	}

	cmd := &cobra.Command{
		Use:   "create",
		Short: "Create a Kafka instance",
		Long: heredoc.Doc(`
			Create a Kafka instance.

			The created instance is set in the config as the current Kafka instance.

		  The created instance can be viewed from the command-line and the web UI.
		`),
		Example: heredoc.Doc(`
			# start an interactive prompt to fill out the configuration values for the instance
			$ rhoas kafka create

			# create a Kafka instance
			$ rhoas kafka create my-kafka-instance

			# create a Kafka instance and output the result in YAML
			$ rhoas kafka create -o yaml
		`),
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) > 0 {
				opts.name = args[0]
			}

			if !opts.IO.CanPrompt() && opts.name == "" {
				return fmt.Errorf("name required when not running interactively. Run 'rhoas kafka create <instance-name>' to create a Kafka instance")
			} else if opts.name == "" && opts.provider == "" && opts.region == "" {
				opts.interactive = true
			}

			if opts.outputFormat != "json" && opts.outputFormat != "yaml" && opts.outputFormat != "yml" {
				return fmt.Errorf("Invalid output format '%v'", opts.outputFormat)
			}

			return runCreate(opts)
		},
	}

	cmd.Flags().StringVar(&opts.provider, flags.FlagProvider, "", "Cloud provider ID")
	cmd.Flags().StringVar(&opts.region, flags.FlagRegion, "", "Cloud provider Region ID")
	cmd.Flags().StringVarP(&opts.outputFormat, "output", "o", "json", "Format to display the Kafka instance. Choose from: \"json\", \"yaml\", \"yml\"")

	return cmd
}

func runCreate(opts *Options) error {
	logger, err := opts.Logger()
	if err != nil {
		return err
	}

	cfg, err := opts.Config.Load()
	if err != nil {
		return err
	}

	connection, err := opts.Connection()
	if err != nil {
		return err
	}

	api := connection.API()

	var payload *kasclient.KafkaRequestPayload
	if opts.interactive {
		logger.Debug("Creating Kafka instance in interactive mode")

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

		payload = &kasclient.KafkaRequestPayload{
			Name:          opts.name,
			Region:        &opts.region,
			CloudProvider: &opts.provider,
			MultiAz:       &opts.multiAZ,
		}
	}

	logger.Info("Creating Kafka instance")

	a := api.Kafka().CreateKafka(context.Background())
	a = a.KafkaRequestPayload(*payload)
	a = a.Async(true)
	response, _, apiErr := a.Execute()

	if kas.IsErr(apiErr, kas.ErrorDuplicateKafkaClusterName) {
		return fmt.Errorf("Error: could not create Kafka instance %v: name already exists", color.Info(opts.name))
	}

	if apiErr.Error() != "" {
		return fmt.Errorf("Unable to create Kafka instance: %w", apiErr)
	}

	logger.Info("Kafka instance created:")
	switch opts.outputFormat {
	case "json":
		data, _ := json.MarshalIndent(response, "", cmdutil.DefaultJSONIndent)
		_ = dump.JSON(opts.IO.Out, data)
	case "yaml", "yml":
		data, _ := yaml.Marshal(response)
		_ = dump.YAML(opts.IO.Out, data)
	}

	kafkaCfg := &config.KafkaConfig{
		ClusterID: *response.Id,
	}

	cfg.Services.Kafka = kafkaCfg
	if err := opts.Config.Save(cfg); err != nil {
		return fmt.Errorf("Unable to use Kafka instance: %w", err)
	}

	return nil
}

// Show a prompt to allow the user to interactively insert the data for their Kafka
func promptKafkaPayload(opts *Options) (payload *kasclient.KafkaRequestPayload, err error) {
	connection, err := opts.Connection()
	if err != nil {
		return nil, err
	}

	api := connection.API()

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
		Message: "Name:",
		Help:    "The name of the Kafka instance",
	}

	err = survey.AskOne(promptName, &answers.Name, survey.WithValidator(pkgKafka.ValidateName))
	if err != nil {
		return nil, err
	}

	// fetch all cloud available providers
	cloudProviderResponse, _, apiErr := api.Kafka().ListCloudProviders(context.Background()).Execute()
	if apiErr.Error() != "" {
		return nil, apiErr
	}

	cloudProviders := cloudProviderResponse.GetItems()
	cloudProviderNames := cloudproviderutil.GetEnabledNames(cloudProviders)

	cloudProviderPrompt := &survey.Select{
		Message: "Cloud Provider:",
		Options: cloudProviderNames,
	}

	err = survey.AskOne(cloudProviderPrompt, &answers.CloudProvider)
	if err != nil {
		return nil, err
	}

	// get the selected provider type from the name selected
	selectedCloudProvider := cloudproviderutil.FindByName(cloudProviders, answers.CloudProvider)

	// nolint
	cloudRegionResponse, _, apiErr := api.Kafka().ListCloudProviderRegions(context.Background(), selectedCloudProvider.GetId()).Execute()
	if apiErr.Error() != "" {
		return nil, apiErr
	}

	regions := cloudRegionResponse.GetItems()
	regionIDs := cloudregionutil.GetEnabledIDs(regions)

	regionPrompt := &survey.Select{
		Message: "Cloud Region:",
		Options: regionIDs,
		Help:    "Geographical region where the Kafka instance will be deployed",
	}

	err = survey.AskOne(regionPrompt, &answers.Region)
	if err != nil {
		return nil, err
	}

	payload = &kasclient.KafkaRequestPayload{
		Name:          answers.Name,
		Region:        &answers.Region,
		CloudProvider: &answers.CloudProvider,
		MultiAz:       &answers.MultiAZ,
	}

	return payload, nil
}
