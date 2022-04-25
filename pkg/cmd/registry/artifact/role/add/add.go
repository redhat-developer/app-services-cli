package add

import (
	"context"

	"github.com/redhat-developer/app-services-cli/pkg/cmd/registry/artifact/util"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/registry/registrycmdutil"
	"github.com/redhat-developer/app-services-cli/pkg/core/ioutil/iostreams"
	"github.com/redhat-developer/app-services-cli/pkg/core/localize"
	"github.com/redhat-developer/app-services-cli/pkg/core/logging"
	"github.com/redhat-developer/app-services-cli/pkg/core/servicecontext"
	"github.com/redhat-developer/app-services-cli/pkg/shared/connection"
	"github.com/redhat-developer/app-services-cli/pkg/shared/contextutil"
	"github.com/redhat-developer/app-services-cli/pkg/shared/factory"
	registryinstanceclient "github.com/redhat-developer/app-services-sdk-go/registryinstance/apiv1internal/client"

	"github.com/spf13/cobra"
)

type options struct {
	user           string
	serviceAccount string
	role           string

	principal string

	registryID string

	IO             *iostreams.IOStreams
	Connection     factory.ConnectionFunc
	Logger         logging.Logger
	localizer      localize.Localizer
	Context        context.Context
	ServiceContext servicecontext.IContext
}

// NewAddCommand creates a new command for creating new service registry access rules
func NewAddCommand(f *factory.Factory) *cobra.Command {
	opts := &options{
		IO:             f.IOStreams,
		Connection:     f.Connection,
		Logger:         f.Logger,
		localizer:      f.Localizer,
		Context:        f.Context,
		ServiceContext: f.ServiceContext,
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
				if role := util.GetRoleEnum(opts.role); role == "" {
					return opts.localizer.MustLocalizeError("artifact.cmd.common.error.invalidRole",
						localize.NewEntry("AllowedRoles", util.GetAllowedRoleTypeEnumValuesAsString()))
				}
			}

			if opts.registryID != "" {
				return runAdd(opts)
			}

			registryInstance, err := contextutil.GetCurrentRegistryInstance(f)
			if err != nil {
				return err
			}

			opts.registryID = registryInstance.GetId()
			return runAdd(opts)
		},
	}

	cmd.Flags().StringVar(&opts.serviceAccount, "service-account", "", opts.localizer.MustLocalize("registry.role.cmd.flag.serviceAccount.description"))
	cmd.Flags().StringVar(&opts.user, "username", "", opts.localizer.MustLocalize("registry.role.cmd.flag.username.description"))
	cmd.Flags().StringVar(&opts.role, "role", "", opts.localizer.MustLocalize("registry.role.cmd.flag.role.description"))
	cmd.Flags().StringVar(&opts.registryID, "instance-id", "", opts.localizer.MustLocalize("registry.common.flag.instance.id"))

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

	role := util.GetRoleEnum(opts.role)

	if principalHasRole(opts, dataAPI.AdminApi) {
		opts.Logger.Info(opts.localizer.MustLocalize("registry.role.cmd.updating"))
		request := dataAPI.AdminApi.UpdateRoleMapping(opts.Context, opts.principal)
		_, err = request.UpdateRole(registryinstanceclient.UpdateRole{
			Role: role,
		}).Execute()
		if err != nil {
			return registrycmdutil.TransformInstanceError(err)
		}
	} else {
		opts.Logger.Info(opts.localizer.MustLocalize("registry.role.cmd.creating"))
		roleMapping := registryinstanceclient.RoleMapping{
			PrincipalId: opts.principal,
			Role:        role,
		}
		request := dataAPI.AdminApi.CreateRoleMapping(opts.Context)
		_, err = request.RoleMapping(roleMapping).Execute()
		if err != nil {
			return registrycmdutil.TransformInstanceError(err)
		}
	}

	opts.Logger.Info(opts.localizer.MustLocalize("registry.role.cmd.add.success"))

	return nil

}

func principalHasRole(opts *options, admin registryinstanceclient.AdminApi) bool {
	_, _, err := admin.GetRoleMapping(opts.Context, opts.principal).Execute()
	if err != nil {
		apiError, _ := registrycmdutil.GetInstanceAPIError(err)
		return apiError.GetErrorCode() != 404
	}
	return true
}
