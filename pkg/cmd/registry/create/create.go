package create

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/redhat-developer/app-services-cli/pkg/api/ams/amsclient"
	"github.com/redhat-developer/app-services-cli/pkg/localize"

	"github.com/redhat-developer/app-services-cli/pkg/cmd/flag"
	flagutil "github.com/redhat-developer/app-services-cli/pkg/cmdutil/flags"
	"github.com/redhat-developer/app-services-cli/pkg/connection"
	"github.com/redhat-developer/app-services-cli/pkg/kafka"

	"github.com/redhat-developer/app-services-cli/internal/build"

	srsclient "github.com/redhat-developer/app-services-cli/pkg/api/srs/client"

	"github.com/AlecAivazis/survey/v2"
	"github.com/redhat-developer/app-services-cli/pkg/dump"
	"github.com/redhat-developer/app-services-cli/pkg/iostreams"

	pkgKafka "github.com/redhat-developer/app-services-cli/pkg/kafka"
	"github.com/redhat-developer/app-services-cli/pkg/logging"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"

	"github.com/redhat-developer/app-services-cli/internal/config"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/factory"
	"github.com/redhat-developer/app-services-cli/pkg/cmdutil"
)

type Options struct {
	name string

	outputFormat string
	autoUse      bool

	interactive bool

	IO         *iostreams.IOStreams
	Config     config.IConfig
	Connection factory.ConnectionFunc
	Logger     func() (logging.Logger, error)
	localizer  localize.Localizer
}

// NewCreateCommand creates a new command for creating kafkas.
func NewCreateCommand(f *factory.Factory) *cobra.Command {
	opts := &Options{
		IO:         f.IOStreams,
		Config:     f.Config,
		Connection: f.Connection,
		Logger:     f.Logger,
		localizer:  f.Localizer,
	}

	cmd := &cobra.Command{
		Use:     "create",
		Short:   "Create service registry",
		Long:    "",
		Example: "",
		Args:    cobra.RangeArgs(0, 1),
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) > 0 {
				opts.name = args[0]

				if err := kafka.ValidateName(opts.name); err != nil {
					return err
				}
			}

			if !opts.IO.CanPrompt() && opts.name == "" {
				return errors.New(opts.localizer.MustLocalize("kafka.create.argument.name.error.requiredWhenNonInteractive"))
			} else if opts.name == "" {
				opts.interactive = true
			}

			validOutputFormats := flagutil.ValidOutputFormats
			if opts.outputFormat != "" && !flagutil.IsValidInput(opts.outputFormat, validOutputFormats...) {
				return flag.InvalidValueError("output", opts.outputFormat, validOutputFormats...)
			}

			return runCreate(opts)
		},
	}

	cmd.Flags().StringVarP(&opts.outputFormat, "output", "o", "json", opts.localizer.MustLocalize("kafka.common.flag.output.description"))
	cmd.Flags().BoolVar(&opts.autoUse, "use", true, opts.localizer.MustLocalize("kafka.create.flag.autoUse.description"))

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

	api := connection.API()

	// the user must have accepted the terms and conditions from the provider
	// before they can create a instance
	termsAccepted, termsURL, err := checkTermsAccepted(opts.Connection)
	if err != nil {
		return err
	}
	if !termsAccepted && termsURL != "" {
		// FIXME Common i18n between kafka and registry?
		logger.Info(opts.localizer.MustLocalize("kafka.create.log.info.termsCheck", localize.NewEntry("TermsURL", termsURL)))
		return nil
	}

	var payload *srsclient.RegistryCreate
	if opts.interactive {
		logger.Debug()

		payload, err = promptPayload(opts)
		if err != nil {
			return err
		}

	} else {
		payload = &srsclient.RegistryCreate{
			Name: &opts.name,
		}
	}

	logger.Info(opts.localizer.MustLocalize("kafka.create.log.debug.creatingKafka", localize.NewEntry("Name", opts.name)))

	a := api.ServiceRegistry().CreateRegistry(context.Background())
	a = a.RegistryCreate(*payload)
	response, _, err := a.Execute()

	if err != nil {
		return err
	}

	logger.Info(opts.localizer.MustLocalize("kafka.create.info.successMessage", localize.NewEntry("Name", response.GetName())))

	switch opts.outputFormat {
	case "json":
		data, _ := json.MarshalIndent(response, "", cmdutil.DefaultJSONIndent)
		_ = dump.JSON(opts.IO.Out, data)
	case "yaml", "yml":
		data, _ := yaml.Marshal(response)
		_ = dump.YAML(opts.IO.Out, data)
	}

	registryConfig := &config.ServiceRegistryConfig{
		InstanceID: response.GetId(),
		Name:       *response.Name,
	}

	if opts.autoUse {
		// FIXME Use as helper?
		logger.Debug(opts.localizer.MustLocalize("kafka.create.debug.autoUseSetMessage"))
		cfg.Services.ServiceRegistry = registryConfig
		if err := opts.Config.Save(cfg); err != nil {
			return fmt.Errorf("%v: %w", opts.localizer.MustLocalize("kafka.common.error.couldNotUseKafka"), err)
		}
	} else {
		logger.Debug(opts.localizer.MustLocalize("kafka.create.debug.autoUseNotSetMessage"))
	}

	return nil
}

// Show a prompt to allow the user to interactively insert the data for their Kafka
func promptPayload(opts *Options) (payload *srsclient.RegistryCreate, err error) {
	if err != nil {
		return nil, err
	}

	// set type to store the answers from the prompt with defaults
	answers := struct {
		Name string
	}{}

	promptName := &survey.Input{
		Message: opts.localizer.MustLocalize("kafka.create.input.name.message"),
		Help:    opts.localizer.MustLocalize("kafka.create.input.name.help"),
	}

	err = survey.AskOne(promptName, &answers.Name, survey.WithValidator(pkgKafka.ValidateName))
	if err != nil {
		return nil, err
	}

	payload = &srsclient.RegistryCreate{
		Name: &answers.Name,
	}

	return payload, nil
}

func checkTermsAccepted(connFunc factory.ConnectionFunc) (accepted bool, redirectURI string, err error) {
	conn, err := connFunc(connection.DefaultConfigSkipMasAuth)
	if err != nil {
		return false, "", err
	}

	termsReview, _, err := conn.API().AccountMgmt().
		ApiAuthorizationsV1SelfTermsReviewPost(context.Background()).
		SelfTermsReview(amsclient.SelfTermsReview{
			EventCode: &build.TermsReviewEventCode,
			SiteCode:  &build.TermsReviewSiteCode,
		}).
		Execute()
	if err != nil {
		return false, "", err
	}

	if !termsReview.GetTermsAvailable() && !termsReview.GetTermsRequired() {
		return true, "", nil
	}

	if !termsReview.HasRedirectUrl() {
		return false, "", errors.New("terms must be signed, but there is no terms URL")
	}

	return false, termsReview.GetRedirectUrl(), nil
}
