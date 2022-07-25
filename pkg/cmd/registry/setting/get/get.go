package get

import (
	"github.com/AlecAivazis/survey/v2"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/registry/registrycmdutil"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/registry/rule/rulecmdutil"
	"github.com/redhat-developer/app-services-cli/pkg/core/cmdutil/flagutil"
	"github.com/redhat-developer/app-services-cli/pkg/core/ioutil/dump"
	"github.com/redhat-developer/app-services-cli/pkg/shared/factory"
	"github.com/spf13/cobra"

	"github.com/redhat-developer/app-services-cli/pkg/shared/contextutil"
)

type options struct {
	registryID  string
	settingName string
	output      string

	f *factory.Factory
}

// NewGetCommand creates a new command to get a service registry setting
func NewGetCommand(f *factory.Factory) *cobra.Command {

	opts := &options{
		f: f,
	}

	cmd := &cobra.Command{
		Use:     "get",
		Short:   f.Localizer.MustLocalize("setting.get.cmd.description.short"),
		Long:    f.Localizer.MustLocalize("setting.get.cmd.description.long"),
		Example: f.Localizer.MustLocalize("setting.get.cmd.example"),
		Args:    cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) (err error) {

			if opts.settingName == "" {
				if !opts.f.IOStreams.CanPrompt() {
					return flagutil.RequiredWhenNonInteractiveError("name")
				}
				err = runInteractivePrompt(opts)
				if err != nil {
					return err
				}
			}

			if opts.registryID != "" {
				return runGet(opts)
			}

			registryInstance, err := contextutil.GetCurrentRegistryInstance(f)
			if err != nil {
				return err
			}

			opts.registryID = registryInstance.GetId()

			return runGet(opts)
		},
	}

	flags := rulecmdutil.NewFlagSet(cmd, f)

	flags.AddRegistryInstance(&opts.registryID)

	flags.StringVarP(&opts.settingName, "name", "n", "", f.Localizer.MustLocalize("setting.get.cmd.flag.settingName.description"))

	flags.AddOutput(&opts.output)

	return cmd
}

func runGet(opts *options) error {
	conn, err := opts.f.Connection()
	if err != nil {
		return err
	}

	api := conn.API()

	a, _, err := api.ServiceRegistryInstance(opts.registryID)
	if err != nil {
		return err
	}

	request := a.AdminApi.GetConfigProperty(opts.f.Context, opts.settingName)

	configProperty, _, err := request.Execute()
	if err != nil {
		return registrycmdutil.TransformInstanceError(err)
	}

	return dump.Formatted(opts.f.IOStreams.Out, opts.output, configProperty)
}

func runInteractivePrompt(opts *options) (err error) {

	settingNamePrompt := &survey.Input{
		Message: opts.f.Localizer.MustLocalize("setting.get.input.settingName.message"),
	}

	err = survey.AskOne(settingNamePrompt, &opts.settingName)
	if err != nil {
		return err
	}

	return nil
}
