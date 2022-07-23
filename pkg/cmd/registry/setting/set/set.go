package set

import (
	"context"
	"github.com/AlecAivazis/survey/v2"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/registry/registrycmdutil"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/registry/rule/rulecmdutil"
	"github.com/redhat-developer/app-services-cli/pkg/core/cmdutil/flagutil"
	"github.com/redhat-developer/app-services-cli/pkg/core/ioutil/icon"
	"github.com/redhat-developer/app-services-cli/pkg/core/ioutil/iostreams"
	"github.com/redhat-developer/app-services-cli/pkg/core/localize"
	"github.com/redhat-developer/app-services-cli/pkg/core/logging"
	"github.com/redhat-developer/app-services-cli/pkg/core/servicecontext"
	"github.com/redhat-developer/app-services-cli/pkg/shared/factory"
	registryinstanceclient "github.com/redhat-developer/app-services-sdk-go/registryinstance/apiv1internal/client"
	"github.com/spf13/cobra"
	"k8s.io/utils/strings/slices"

	"github.com/redhat-developer/app-services-cli/pkg/shared/contextutil"
)

type options struct {
	IO             *iostreams.IOStreams
	Connection     factory.ConnectionFunc
	Logger         logging.Logger
	localizer      localize.Localizer
	Context        context.Context
	ServiceContext servicecontext.IContext

	registryID     string
	settingName    string
	value          string
	resetToDefault bool
}

// NewSetCommand creates a new command to set a service registry setting
func NewSetCommand(f *factory.Factory) *cobra.Command {

	opts := &options{
		IO:             f.IOStreams,
		Connection:     f.Connection,
		Logger:         f.Logger,
		localizer:      f.Localizer,
		Context:        f.Context,
		ServiceContext: f.ServiceContext,
	}

	cmd := &cobra.Command{
		Use:     "set",
		Short:   f.Localizer.MustLocalize("setting.set.cmd.description.short"),
		Long:    f.Localizer.MustLocalize("setting.set.cmd.description.long"),
		Example: f.Localizer.MustLocalize("setting.set.cmd.example"),
		Args:    cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) (err error) {

			var missingFlags []string

			if opts.settingName == "" {
				missingFlags = append(missingFlags, "setting-name")
			}
			if opts.value == "" && !opts.resetToDefault {
				missingFlags = append(missingFlags, "value")
			}

			if !opts.IO.CanPrompt() && len(missingFlags) > 0 {
				return flagutil.RequiredWhenNonInteractiveError(missingFlags...)
			}

			if len(missingFlags) > 0 {
				err = runInteractivePrompt(opts, missingFlags)
				if err != nil {
					return err
				}
			}

			registryInstance, err := contextutil.GetCurrentRegistryInstance(f)
			if err != nil {
				return err
			}

			opts.registryID = registryInstance.GetId()

			return runSet(opts)
		},
	}

	flags := rulecmdutil.NewFlagSet(cmd, f)

	flags.AddRegistryInstance(&opts.registryID)

	flags.StringVarP(&opts.settingName, "setting-name", "n", "", f.Localizer.MustLocalize("setting.set.cmd.flag.settingName.description"))
	flags.StringVar(&opts.value, "value", "", f.Localizer.MustLocalize("setting.set.cmd.flag.value.description"))
	flags.BoolVar(&opts.resetToDefault, "default", false, f.Localizer.MustLocalize("setting.set.cmd.flag.default.description"))

	return cmd
}

func runSet(opts *options) error {
	conn, err := opts.Connection()
	if err != nil {
		return err
	}

	api := conn.API()

	a, _, err := api.ServiceRegistryInstance(opts.registryID)
	if err != nil {
		return err
	}

	if !opts.resetToDefault {
		request := a.AdminApi.UpdateConfigProperty(opts.Context, opts.settingName)

		request = request.UpdateConfigurationProperty(registryinstanceclient.UpdateConfigurationProperty{Value: opts.value})

		_, err = request.Execute()
		if err != nil {
			return registrycmdutil.TransformInstanceError(err)
		}

		opts.Logger.Info(icon.SuccessPrefix(), opts.localizer.MustLocalize("setting.set.log.info.settingSet"))
	} else {
		request := a.AdminApi.ResetConfigProperty(opts.Context, opts.settingName)

		_, err = request.Execute()
		if err != nil {
			return registrycmdutil.TransformInstanceError(err)
		}

		opts.Logger.Info(icon.SuccessPrefix(), opts.localizer.MustLocalize("setting.set.log.info.settingReset"))
	}
	return nil
}

func runInteractivePrompt(opts *options, missingFlags []string) (err error) {

	if slices.Contains(missingFlags, "setting-name") {
		settingNamePrompt := &survey.Input{
			Message: opts.localizer.MustLocalize("setting.set.input.settingName.message"),
		}

		err = survey.AskOne(settingNamePrompt, &opts.settingName)
		if err != nil {
			return err
		}
	}

	if slices.Contains(missingFlags, "value") {
		valuePrompt := &survey.Input{
			Message: opts.localizer.MustLocalize("setting.set.input.value.message"),
		}

		err = survey.AskOne(valuePrompt, &opts.value)
		if err != nil {
			return err
		}
	}

	return nil
}
