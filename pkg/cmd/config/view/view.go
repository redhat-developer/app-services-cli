package create

import (
	"context"
	"fmt"
	"os"

	"github.com/redhat-developer/app-services-cli/pkg/color"
	"github.com/redhat-developer/app-services-cli/pkg/icon"
	"github.com/redhat-developer/app-services-cli/pkg/ioutil/spinner"

	"github.com/redhat-developer/app-services-cli/pkg/localize"
	"github.com/redhat-developer/app-services-cli/pkg/serviceaccount/validation"
	kafkamgmtclient "github.com/redhat-developer/app-services-sdk-go/kafkamgmt/apiv1/client"

	"github.com/redhat-developer/app-services-cli/pkg/connection"

	"github.com/AlecAivazis/survey/v2"
	flagutil "github.com/redhat-developer/app-services-cli/pkg/cmdutil/flagutil"
	"github.com/redhat-developer/app-services-cli/pkg/iostreams"
	"github.com/redhat-developer/app-services-cli/pkg/serviceaccount/credentials"

	"github.com/redhat-developer/app-services-cli/internal/config"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/factory"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/flag"
	"github.com/redhat-developer/app-services-cli/pkg/logging"
	"github.com/spf13/cobra"
)

type options struct {
	IO         *iostreams.IOStreams
	Config     config.IConfig
	Connection factory.ConnectionFunc
	Logger     logging.Logger
	localizer  localize.Localizer
	Context    context.Context

	fileFormat       string
	overwrite        bool
	shortDescription string
	filename         string

	interactive bool
}

// NewCreateCommand creates a new command to create service accounts
func NewCreateCommand(f *factory.Factory) *cobra.Command {
	opts := &options{
		IO:         f.IOStreams,
		Config:     f.Config,
		Connection: f.Connection,
		Logger:     f.Logger,
		localizer:  f.Localizer,
		Context:    f.Context,
	}

	cmd := &cobra.Command{
		Use:     "manage",
		Short:   "manage configuration profiles. What services are enabled",
		Long:    heredoc.Doc`
		Manage command lets you select groups of services to be used at certain time.
		`,
		Example:  ,
		Args:    cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) (err error) {
			 		return runCreate(opts)
		},
	}

	cmd.Flags().StringVar(&opts.shortDescription, "short-description", "", opts.localizer.MustLocalize("serviceAccount.create.flag.shortDescription.description"))
	cmd.Flags().BoolVar(&opts.overwrite, "overwrite", false, opts.localizer.MustLocalize("serviceAccount.common.flag.overwrite.description"))
	cmd.Flags().StringVar(&opts.filename, "output-file", "", opts.localizer.MustLocalize("serviceAccount.common.flag.fileLocation.description"))
	cmd.Flags().StringVar(&opts.fileFormat, "file-format", "", opts.localizer.MustLocalize("serviceAccount.common.flag.fileFormat.description"))

	flagutil.EnableStaticFlagCompletion(cmd, "file-format", flagutil.CredentialsOutputFormats)

	return cmd
}

// nolint:funlen
func runCreate(opts *options) error {
	conn, err := opts.Connection(connection.DefaultConfigSkipMasAuth)
	if err != nil {
		return err
	}
	cfg, err := opts.Config.Load()
	if err != nil {
		return err
	}

	if opts.interactive {
		// run the create command interactively
		err = runInteractivePrompt(opts)
		if err != nil {
			return err
		}
	} else if opts.filename == "" {
		// obtain the absolute path to where credentials will be saved
		opts.filename = credentials.GetDefaultPath(opts.fileFormat)
	}

	// If the credentials file already exists, and the --overwrite flag is not set then return an error
	// indicating that the user should explicitly request overwriting of the file
	_, err = os.Stat(opts.filename)
	if err == nil && !opts.overwrite {
		return opts.localizer.MustLocalizeError("serviceAccount.common.error.credentialsFileAlreadyExists", localize.NewEntry("FilePath", opts.filename))
	}

	spinner := spinner.New(opts.IO.ErrOut, opts.localizer)
	spinner.SetSuffix(opts.localizer.MustLocalize("serviceAccount.create.log.info.creating"))
	spinner.Start()
	// create the service account
	serviceAccountPayload := kafkamgmtclient.ServiceAccountRequest{Name: opts.shortDescription}

	serviceacct, httpRes, err := conn.API().
		ServiceAccount().
		CreateServiceAccount(opts.Context).
		ServiceAccountRequest(serviceAccountPayload).
		Execute()
	spinner.Stop()

	if httpRes != nil {
		defer httpRes.Body.Close()
	}

	if err != nil {
		return err
	}

	opts.Logger.Info(icon.SuccessPrefix(), opts.localizer.MustLocalize("serviceAccount.create.log.info.createdSuccessfully", localize.NewEntry("ID", serviceacct.GetId())))

	creds := &credentials.Credentials{
		ClientID:     serviceacct.GetClientId(),
		ClientSecret: serviceacct.GetClientSecret(),
		TokenURL:     cfg.MasAuthURL + "/protocol/openid-connect/token",
	}

	// save the credentials to a file
	err = credentials.Write(opts.fileFormat, opts.filename, creds)
	if err != nil {
		return fmt.Errorf("%v: %w", opts.localizer.MustLocalize("serviceAccount.common.error.couldNotSaveCredentialsFile"), err)
	}

	opts.Logger.Info(icon.SuccessPrefix(), opts.localizer.MustLocalize("serviceAccount.common.log.info.credentialsSaved",
		localize.NewEntry("FilePath", color.CodeSnippet(opts.filename)),
		localize.NewEntry("ClientID", color.Success(creds.ClientID)),
	))

	return nil
}

func runInteractivePrompt(opts *options) (err error) {
	_, err = opts.Connection(connection.DefaultConfigSkipMasAuth)
	if err != nil {
		return err
	}

	validator := &validation.Validator{
		Localizer: opts.localizer,
	}

	opts.Logger.Debug(opts.localizer.MustLocalize("common.log.debug.startingInteractivePrompt"))

	promptName := &survey.Input{
		Message: opts.localizer.MustLocalize("serviceAccount.create.input.shortDescription.message"),
		Help:    opts.localizer.MustLocalize("serviceAccount.create.input.shortDescription.help"),
	}

	err = survey.AskOne(promptName, &opts.shortDescription, survey.WithValidator(survey.Required), survey.WithValidator(validator.ValidateShortDescription))
	if err != nil {
		return err
	}

	// if the --file-format flag was not used, ask in the prompt
	if opts.fileFormat == "" {
		opts.Logger.Debug(opts.localizer.MustLocalize("serviceAccount.common.log.debug.interactive.fileFormatNotSet"))

		fileFormatPrompt := &survey.Select{
			Message: opts.localizer.MustLocalize("serviceAccount.create.input.fileFormat.message"),
			Help:    opts.localizer.MustLocalize("serviceAccount.create.input.fileFormat.help"),
			Options: flagutil.CredentialsOutputFormats,
			Default: credentials.EnvFormat,
		}

		err = survey.AskOne(fileFormatPrompt, &opts.fileFormat)
		if err != nil {
			return err
		}
	}

	opts.filename, opts.overwrite, err = credentials.ChooseFileLocation(opts.fileFormat, opts.filename, opts.overwrite)
	if err != nil {
		return err
	}

	return nil
}
