package resetcredentials

import (
	"context"
	kafkamgmtv1 "github.com/redhat-developer/app-services-sdk-go/apis/kafka/kafkamgmt/v1"
	"errors"
	"fmt"
	"os"

	"github.com/redhat-developer/app-services-cli/pkg/connection"
	"github.com/redhat-developer/app-services-cli/pkg/localize"

	"github.com/AlecAivazis/survey/v2"
	flagutil "github.com/redhat-developer/app-services-cli/pkg/cmdutil/flags"
	"github.com/redhat-developer/app-services-cli/pkg/iostreams"
	"github.com/redhat-developer/app-services-cli/pkg/serviceaccount/credentials"

	"github.com/redhat-developer/app-services-cli/internal/config"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/factory"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/flag"
	"github.com/redhat-developer/app-services-cli/pkg/logging"
	"github.com/spf13/cobra"
)

type Options struct {
	IO         *iostreams.IOStreams
	Config     config.IConfig
	Connection factory.ConnectionFunc
	Logger     func() (logging.Logger, error)
	localizer  localize.Localizer

	id         string
	fileFormat string
	overwrite  bool
	filename   string

	interactive bool
	force       bool
}

// NewResetCredentialsCommand creates a new command to delete a service account
func NewResetCredentialsCommand(f *factory.Factory) *cobra.Command {
	opts := &Options{
		IO:         f.IOStreams,
		Config:     f.Config,
		Connection: f.Connection,
		Logger:     f.Logger,
		localizer:  f.Localizer,
	}

	cmd := &cobra.Command{
		Use:     opts.localizer.MustLocalize("serviceAccount.resetCredentials.cmd.use"),
		Short:   opts.localizer.MustLocalize("serviceAccount.resetCredentials.cmd.shortDescription"),
		Long:    opts.localizer.MustLocalize("serviceAccount.resetCredentials.cmd.longDescription"),
		Example: opts.localizer.MustLocalize("serviceAccount.resetCredentials.cmd.example"),
		RunE: func(cmd *cobra.Command, _ []string) error {
			if !opts.IO.CanPrompt() && opts.id == "" {
				return errors.New(opts.localizer.MustLocalize("flag.error.requiredWhenNonInteractive", localize.NewEntry("Flag", "id")))
			} else if opts.id == "" {
				opts.interactive = true
			}

			if !opts.interactive && opts.fileFormat == "" {
				return errors.New(opts.localizer.MustLocalize("flag.error.requiredWhenNonInteractive", localize.NewEntry("Flag", "file-format")))
			}

			validOutput := flagutil.IsValidInput(opts.fileFormat, flagutil.CredentialsOutputFormats...)
			if !validOutput && opts.fileFormat != "" {
				return flag.InvalidValueError("file-format", opts.fileFormat, flagutil.CredentialsOutputFormats...)
			}

			return runResetCredentials(opts)
		},
	}

	cmd.Flags().StringVar(&opts.id, "id", "", opts.localizer.MustLocalize("serviceAccount.resetCredentials.flag.id.description"))
	cmd.Flags().BoolVar(&opts.overwrite, "overwrite", false, opts.localizer.MustLocalize("serviceAccount.common.flag.overwrite.description"))
	cmd.Flags().StringVar(&opts.filename, "file-location", "", opts.localizer.MustLocalize("serviceAccount.common.flag.fileLocation.description"))
	cmd.Flags().StringVar(&opts.fileFormat, "file-format", "", opts.localizer.MustLocalize("serviceAccount.common.flag.fileFormat.description"))
	cmd.Flags().BoolVarP(&opts.force, "yes", "y", false, opts.localizer.MustLocalize("serviceAccount.resetCredentials.flag.yes.description"))

	return cmd
}

// nolint:funlen
func runResetCredentials(opts *Options) (err error) {
	connection, err := opts.Connection(connection.DefaultConfigSkipMasAuth)
	if err != nil {
		return err
	}

	logger, err := opts.Logger()
	if err != nil {
		return err
	}

	api := connection.API()

	serviceacct, _, err := api.Kafka().GetServiceAccountById(context.Background(), opts.id).Execute()
	if err != nil {
		return err
	}
	serviceAcctName := serviceacct.GetName()

	if opts.interactive {
		err = runInteractivePrompt(opts)
		if err != nil {
			return err
		}
	} else if opts.filename == "" {
		// obtain the default absolute path to where credentials will be saved
		opts.filename = credentials.GetDefaultPath(opts.fileFormat)
	}

	// If the credentials file already exists, and the --overwrite flag is not set then return an error
	// indicating that the user should explicitly request overwriting of the file
	if _, err = os.Stat(opts.filename); err == nil && !opts.overwrite {
		return errors.New(opts.localizer.MustLocalize("serviceAccount.common.error.credentialsFileAlreadyExists", localize.NewEntry("FilePath", opts.filename)))
	}

	if !opts.force {
		// prompt the user to confirm their wish to proceed with this action
		var confirmReset bool
		opts.localizer.MustLocalize("serviceAccount.resetCredentials.input.confirmReset.message", localize.NewEntry("ID", opts.id))
		promptConfirmDelete := &survey.Confirm{
			Message: opts.localizer.MustLocalize("serviceAccount.resetCredentials.input.confirmReset.message", localize.NewEntry("ID", opts.id)),
		}

		if err = survey.AskOne(promptConfirmDelete, &confirmReset); err != nil {
			return err
		}
		if !confirmReset {
			logger.Debug(opts.localizer.MustLocalize("serviceAccount.resetCredentials.log.debug.cancelledReset"))
			return nil
		}
	}

	updatedServiceAccount, err := resetCredentials(serviceAcctName, opts)

	if err != nil {
		return fmt.Errorf("%v: %w", opts.localizer.MustLocalize("serviceAccount.resetCredentials.error.resetError", localize.NewEntry("Name", updatedServiceAccount.GetName())), err)
	}

	logger.Info(opts.localizer.MustLocalize("serviceAccount.resetCredentials.log.info.resetSuccess", localize.NewEntry("Name", updatedServiceAccount.GetName())))

	creds := &credentials.Credentials{
		ClientID:     updatedServiceAccount.GetClientID(),
		ClientSecret: updatedServiceAccount.GetClientSecret(),
	}

	// save the credentials to a file
	err = credentials.Write(opts.fileFormat, opts.filename, creds)
	if err != nil {
		return err
	}

	logger.Info(opts.localizer.MustLocalize("serviceAccount.common.log.info.credentialsSaved", localize.NewEntry("FilePath", opts.filename)))

	return nil
}

func resetCredentials(name string, opts *Options) (*kafkamgmtv1.ServiceAccount, error) {
	connection, err := opts.Connection(connection.DefaultConfigSkipMasAuth)
	if err != nil {
		return nil, err
	}

	// check if the service account exists
	api := connection.API()

	logger, err := opts.Logger()
	if err != nil {
		return nil, err
	}

	logger.Debug(opts.localizer.MustLocalize("serviceAccount.resetCredentials.log.debug.resettingCredentials", localize.NewEntry("Name", name)))

	serviceacct, httpRes, err := api.Kafka().ResetServiceAccountCreds(context.Background(), opts.id).Execute()

	if err != nil {
		if httpRes == nil {
			return nil, err
		}

		switch httpRes.StatusCode {
		case 403:
			opts.localizer.MustLocalize("serviceAccount.common.error.forbidden", localize.NewEntry("Operation", "update"))
			return nil, fmt.Errorf("%v: %w", opts.localizer.MustLocalize("serviceAccount.common.error.forbidden", localize.NewEntry("Operation", "update")), err)
		case 500:
			return nil, errors.New(opts.localizer.MustLocalize("serviceAccount.common.error.internalServerError"))
		default:
			return nil, err
		}
	}

	return &serviceacct, nil
}

func runInteractivePrompt(opts *Options) (err error) {
	_, err = opts.Connection(connection.DefaultConfigSkipMasAuth)
	if err != nil {
		return err
	}

	logger, err := opts.Logger()
	if err != nil {
		return err
	}

	logger.Debug(opts.localizer.MustLocalize("common.log.debug.startingInteractivePrompt"))

	promptID := &survey.Input{
		Message: opts.localizer.MustLocalize("serviceAccount.resetCredentials.input.id.message"),
		Help:    opts.localizer.MustLocalize("serviceAccount.resetCredentials.input.id.help"),
	}

	err = survey.AskOne(promptID, &opts.id, survey.WithValidator(survey.Required))
	if err != nil {
		return err
	}

	// if the --output flag was not used, ask in the prompt
	if opts.fileFormat == "" {
		logger.Debug(opts.localizer.MustLocalize("serviceAccount.common.log.debug.interactive.fileFormatNotSet"))

		fileFormatPrompt := &survey.Select{
			Message: opts.localizer.MustLocalize("serviceAccount.resetCredentials.input.fileFormat.message"),
			Help:    opts.localizer.MustLocalize("serviceAccount.resetCredentials.input.fileFormat.help"),
			Options: flagutil.CredentialsOutputFormats,
			Default: "env",
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

	return err
}
