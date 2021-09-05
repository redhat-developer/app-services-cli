package delete

import (
	"context"
	"errors"
	"github.com/redhat-developer/app-services-cli/internal/build"
	"net/http"

	"github.com/AlecAivazis/survey/v2"
	"github.com/redhat-developer/app-services-cli/internal/config"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/factory"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/flag"
	"github.com/redhat-developer/app-services-cli/pkg/connection"
	"github.com/redhat-developer/app-services-cli/pkg/iostreams"
	"github.com/redhat-developer/app-services-cli/pkg/localize"
	"github.com/redhat-developer/app-services-cli/pkg/logging"
	"github.com/redhat-developer/app-services-cli/pkg/serviceaccount/validation"
	"github.com/spf13/cobra"
)

type Options struct {
	IO         *iostreams.IOStreams
	Config     config.IConfig
	Connection factory.ConnectionFunc
	Logger     logging.Logger
	localizer  localize.Localizer

	id    string
	force bool
}

// NewDeleteCommand creates a new command to delete a service account
func NewDeleteCommand(f *factory.Factory) *cobra.Command {
	opts := &Options{
		Config:     f.Config,
		Connection: f.Connection,
		Logger:     f.Logger,
		IO:         f.IOStreams,
		localizer:  f.Localizer,
	}

	cmd := &cobra.Command{
		Use:     opts.localizer.MustLocalize("serviceAccount.delete.cmd.use"),
		Short:   opts.localizer.MustLocalize("serviceAccount.delete.cmd.shortDescription"),
		Long:    opts.localizer.MustLocalize("serviceAccount.delete.cmd.longDescription"),
		Example: opts.localizer.MustLocalize("serviceAccount.delete.cmd.example"),
		Args:    cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			if !opts.IO.CanPrompt() && !opts.force {
				return flag.RequiredWhenNonInteractiveError("yes")
			}

			validator := &validation.Validator{
				Localizer: opts.localizer,
			}

			validID := validator.ValidateUUID(opts.id)
			if validID != nil {
				return validID
			}

			return runDelete(opts)
		},
	}

	cmd.Flags().StringVar(&opts.id, "id", "", opts.localizer.MustLocalize("serviceAccount.delete.flag.id.description"))
	cmd.Flags().BoolVarP(&opts.force, "yes", "y", false, opts.localizer.MustLocalize("serviceAccount.delete.flag.yes.description"))

	_ = cmd.MarkFlagRequired("id")

	return cmd
}

func runDelete(opts *Options) (err error) {
	conn, err := opts.Connection(connection.DefaultConfigSkipMasAuth)
	if err != nil {
		return err
	}

	_, httpRes, err := conn.API().ServiceAccount().GetServiceAccountById(context.Background(), opts.id).Execute()
	if httpRes != nil {
		defer httpRes.Body.Close()
	}
	if err != nil {
		if httpRes == nil {
			return err
		}

		if httpRes.StatusCode == http.StatusNotFound {
			return errors.New(opts.localizer.MustLocalize("serviceAccount.common.error.notFoundError", localize.NewEntry("ID", opts.id)))
		}
	}

	if !opts.force {
		var confirmDelete bool
		promptConfirmDelete := &survey.Confirm{
			Message: opts.localizer.MustLocalize("serviceAccount.delete.input.confirmDelete.message", localize.NewEntry("ID", opts.id)),
		}

		err = survey.AskOne(promptConfirmDelete, &confirmDelete)
		if err != nil {
			return err
		}

		if !confirmDelete {
			opts.Logger.Debug(opts.localizer.MustLocalize("serviceAccount.delete.log.debug.deleteNotConfirmed"))
			return nil
		}
	}

	return deleteServiceAccount(opts)
}

func deleteServiceAccount(opts *Options) error {
	conn, err := opts.Connection(connection.DefaultConfigSkipMasAuth)
	if err != nil {
		return err
	}

	_, httpRes, err := conn.API().ServiceAccount().DeleteServiceAccountById(context.Background(), opts.id).Execute()
	if httpRes != nil {
		defer httpRes.Body.Close()
	}
	if err != nil {
		if httpRes == nil {
			return err
		}

		switch httpRes.StatusCode {
		case http.StatusForbidden:
			return errors.New(opts.localizer.MustLocalize("serviceAccount.common.error.forbidden", localize.NewEntry("Operation", "delete")))
		case http.StatusInternalServerError:
			return errors.New(opts.localizer.MustLocalize("serviceAccount.common.error.internalServerError"))
		default:
			return err
		}
	}

	opts.Logger.Info(build.GetEmojiSuccess(), opts.localizer.MustLocalize("serviceAccount.delete.log.info.deleteSuccess"))

	return nil
}
