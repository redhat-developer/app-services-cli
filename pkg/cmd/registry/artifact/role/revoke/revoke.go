package revoke

import (
	"context"

	"github.com/redhat-developer/app-services-cli/pkg/cmd/registry/sdk"
	"github.com/redhat-developer/app-services-cli/pkg/core/cmdutil/factory"
	"github.com/redhat-developer/app-services-cli/pkg/core/config"
	"github.com/redhat-developer/app-services-cli/pkg/core/connection"
	"github.com/redhat-developer/app-services-cli/pkg/core/ioutil/iostreams"
	"github.com/redhat-developer/app-services-cli/pkg/core/localize"
	"github.com/redhat-developer/app-services-cli/pkg/core/logging"
	"github.com/spf13/cobra"
)

type options struct {
	user           string
	serviceAccount string

	principal  string
	registryID string

	IO         *iostreams.IOStreams
	Config     config.IConfig
	Connection factory.ConnectionFunc
	Logger     logging.Logger
	localizer  localize.Localizer
	Context    context.Context
}

// NewRevokeCommand creates command used to revoke access for the user
func NewRevokeCommand(f *factory.Factory) *cobra.Command {
	opts := &options{
		IO:         f.IOStreams,
		Config:     f.Config,
		Connection: f.Connection,
		Logger:     f.Logger,
		localizer:  f.Localizer,
		Context:    f.Context,
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

			cfg, err := opts.Config.Load()
			if err != nil {
				return err
			}

			instanceID, ok := cfg.GetServiceRegistryIdOk()
			if !ok {
				return opts.localizer.MustLocalizeError("artifact.cmd.common.error.noServiceRegistrySelected")
			}

			opts.registryID = instanceID
			return runRevoke(opts)
		},
	}

	cmd.Flags().StringVar(&opts.serviceAccount, "service-account", "", opts.localizer.MustLocalize("registry.role.cmd.flag.serviceAccount.description"))
	cmd.Flags().StringVar(&opts.user, "username", "", opts.localizer.MustLocalize("registry.role.cmd.flag.username.description"))
	cmd.Flags().StringVar(&opts.registryID, "instance-id", "", opts.localizer.MustLocalize("artifact.common.instance.id"))

	return cmd
}

func runRevoke(opts *options) error {
	conn, err := opts.Connection(connection.DefaultConfigRequireMasAuth)
	if err != nil {
		return err
	}

	dataAPI, _, err := conn.API().ServiceRegistryInstance(opts.registryID)
	if err != nil {
		return err
	}

	_, err = dataAPI.AdminApi.DeleteRoleMapping(opts.Context, opts.principal).Execute()
	if err != nil {
		return sdk.TransformInstanceError(err)
	}
	opts.Logger.Info(opts.localizer.MustLocalize("registry.role.cmd.revoke.success"))

	return nil

}
