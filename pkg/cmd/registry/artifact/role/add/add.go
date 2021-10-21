package add

import (
	"context"

	"github.com/redhat-developer/app-services-cli/pkg/localize"
	"github.com/redhat-developer/app-services-cli/pkg/serviceregistry/registryinstanceerror"

	"github.com/redhat-developer/app-services-cli/pkg/cmd/kafka/flagutil"
	"github.com/redhat-developer/app-services-cli/pkg/connection"
	registryinstanceclient "github.com/redhat-developer/app-services-sdk-go/registryinstance/apiv1internal/client"

	"github.com/redhat-developer/app-services-cli/pkg/iostreams"

	"github.com/redhat-developer/app-services-cli/pkg/logging"

	"github.com/spf13/cobra"

	"github.com/redhat-developer/app-services-cli/internal/config"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/factory"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/registry/artifact/util"
)

type options struct {
	user           string
	serviceAccount string
	role           string

	principal string

	registryID string

	IO         *iostreams.IOStreams
	Config     config.IConfig
	Connection factory.ConnectionFunc
	Logger     logging.Logger
	localizer  localize.Localizer
	Context    context.Context
}

// NewAddCommand creates a new command for creating new service registry access rules
func NewAddCommand(f *factory.Factory) *cobra.Command {
	opts := &options{
		IO:         f.IOStreams,
		Config:     f.Config,
		Connection: f.Connection,
		Logger:     f.Logger,
		localizer:  f.Localizer,
		Context:    f.Context,
	}

	cmd := &cobra.Command{
		Use:     "add",
		Short:   f.Localizer.MustLocalize("registry.role.cmd.add.shortDescription"),
		Long:    f.Localizer.MustLocalize("registry.role.cmd.add.longDescription"),
		Example: f.Localizer.MustLocalize("registry.role.cmd.add.example"),
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

			if opts.role != "" {
				if _, err := registryinstanceclient.NewRoleTypeFromValue(opts.role); err != nil {
					return opts.localizer.MustLocalizeError("artifact.cmd.common.error.invalidRole",
						localize.NewEntry("AllowedRoles", util.GetAllowedRoleTypeEnumValuesAsString()))
				}
			}

			if opts.registryID != "" {
				return runAdd(opts)
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
			return runAdd(opts)
		},
	}

	flags := flagutil.NewFlagSet(cmd, opts.localizer)
	flags.StringVar(&opts.serviceAccount, "service-account", "", opts.localizer.MustLocalize("registry.role.cmd.flag.serviceAccount.description"))
	flags.StringVar(&opts.user, "username", "", opts.localizer.MustLocalize("registry.role.cmd.flag.username.description"))
	flags.StringVar(&opts.role, "role", "", opts.localizer.MustLocalize("registry.role.cmd.flag.role.description"))
	flags.StringVar(&opts.registryID, "instance-id", "", opts.localizer.MustLocalize("artifact.common.instance.id"))

	_ = cmd.MarkFlagRequired("role")
	_ = cmd.RegisterFlagCompletionFunc("role", func(cmd *cobra.Command, _ []string, _ string) ([]string, cobra.ShellCompDirective) {
		return util.AllowedRoleTypeEnumValues, cobra.ShellCompDirectiveNoSpace
	})

	return cmd
}

func runAdd(opts *options) error {
	conn, err := opts.Connection(connection.DefaultConfigRequireMasAuth)
	if err != nil {
		return err
	}

	dataAPI, _, err := conn.API().ServiceRegistryInstance(opts.registryID)
	if err != nil {
		return err
	}

	role, _ := registryinstanceclient.NewRoleTypeFromValue(opts.role)

	if principalHasRole(opts, dataAPI.AdminApi) {
		opts.Logger.Info(opts.localizer.MustLocalize("registry.role.cmd.updating"))
		request := dataAPI.AdminApi.UpdateRoleMapping(opts.Context, opts.principal)
		_, err = request.UpdateRole(registryinstanceclient.UpdateRole{
			Role: *role,
		}).Execute()
		if err != nil {
			return registryinstanceerror.TransformError(err)
		}
	} else {
		opts.Logger.Info(opts.localizer.MustLocalize("registry.role.cmd.creating"))
		roleMapping := registryinstanceclient.RoleMapping{
			PrincipalId: opts.principal,
			Role:        *role,
		}
		request := dataAPI.AdminApi.CreateRoleMapping(opts.Context)
		_, err = request.RoleMapping(roleMapping).Execute()
		if err != nil {
			return registryinstanceerror.TransformError(err)
		}
	}

	opts.Logger.Info(opts.localizer.MustLocalize("registry.role.cmd.add.success"))

	return nil

}

func principalHasRole(opts *options, admin registryinstanceclient.AdminApi) bool {
	_, _, err := admin.GetRoleMapping(opts.Context, opts.principal).Execute()
	if err != nil {
		apiError, _ := registryinstanceerror.GetAPIError(err)
		return apiError.GetErrorCode() != 404
	}
	return true
}
