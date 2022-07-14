package revoke

import (
	"context"

	"github.com/redhat-developer/app-services-cli/pkg/cmd/registry/registrycmdutil"
	"github.com/redhat-developer/app-services-cli/pkg/core/ioutil/iostreams"
	"github.com/redhat-developer/app-services-cli/pkg/core/localize"
	"github.com/redhat-developer/app-services-cli/pkg/core/logging"
	"github.com/redhat-developer/app-services-cli/pkg/core/servicecontext"
	"github.com/redhat-developer/app-services-cli/pkg/shared/contextutil"
	"github.com/redhat-developer/app-services-cli/pkg/shared/factory"
	"github.com/spf13/cobra"
)

type options struct {
	user           string
	serviceAccount string

	principal  string
	registryID string

	IO             *iostreams.IOStreams
	Connection     factory.ConnectionFunc
	Logger         logging.Logger
	localizer      localize.Localizer
	Context        context.Context
	ServiceContext servicecontext.IContext
}

// NewRevokeCommand creates command used to revoke access for the user
func NewRevokeCommand(f *factory.Factory) *cobra.Command {
	opts := &options{
		IO:             f.IOStreams,
		Connection:     f.Connection,
		Logger:         f.Logger,
		localizer:      f.Localizer,
		Context:        f.Context,
		ServiceContext: f.ServiceContext,
	}

	cmd := &cobra.Command{
		Use:     "revoke",
		Short:   f.Localizer.MustLocalize("registry.role.cmd.revoke.shortDescription"),
		Long:    f.Localizer.MustLocalize("registry.role.cmd.revoke.longDescription"),
		Example: f.Localizer.MustLocalize("registry.role.cmd.revoke.example"),
		Args:    cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			if opts.user != "" && opts.serviceAccount != "" {
				return opts.localizer.MustLocalizeError("artifact.cmd.common.error.useSaOrUserOnly")
			}

			//nolint:gocritic
			if opts.user != "" {
				opts.principal = opts.user
			} else if opts.serviceAccount != "" {
				opts.principal = opts.serviceAccount
			} else {
				return opts.localizer.MustLocalizeError("artifact.cmd.common.error.missingUserOrSA")
			}

			if opts.registryID != "" {
				return runRevoke(opts)
			}

			registryInstance, err := contextutil.GetCurrentRegistryInstance(f)
			if err != nil {
				return err
			}

			opts.registryID = registryInstance.GetId()
			return runRevoke(opts)
		},
	}

	cmd.Flags().StringVar(&opts.serviceAccount, "service-account", "", opts.localizer.MustLocalize("registry.role.cmd.flag.serviceAccount.description"))
	cmd.Flags().StringVar(&opts.user, "username", "", opts.localizer.MustLocalize("registry.role.cmd.flag.username.description"))
	cmd.Flags().StringVar(&opts.registryID, "instance-id", "", opts.localizer.MustLocalize("registry.common.flag.instance.id"))

	return cmd
}

func runRevoke(opts *options) error {
	conn, err := opts.Connection()
	if err != nil {
		return err
	}

	dataAPI, _, err := conn.API().ServiceRegistryInstance(opts.registryID)
	if err != nil {
		return err
	}

	_, err = dataAPI.AdminApi.DeleteRoleMapping(opts.Context, opts.principal).Execute()
	if err != nil {
		return registrycmdutil.TransformInstanceError(err)
	}
	opts.Logger.Info(opts.localizer.MustLocalize("registry.role.cmd.revoke.success"))

	return nil

}
