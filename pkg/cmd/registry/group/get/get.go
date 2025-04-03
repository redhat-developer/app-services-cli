package get

import (
	"github.com/AlecAivazis/survey/v2"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/registry/registrycmdutil"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/registry/rule/rulecmdutil"
	"github.com/redhat-developer/app-services-cli/pkg/core/cmdutil/flagutil"
	"github.com/redhat-developer/app-services-cli/pkg/core/ioutil/dump"
	"github.com/redhat-developer/app-services-cli/pkg/shared/factory"
	"github.com/spf13/cobra"
	"k8s.io/utils/strings/slices"

	"github.com/redhat-developer/app-services-cli/pkg/shared/contextutil"
)

type options struct {
	registryID string

	groupId      string
	outputFormat string

	f *factory.Factory
}

// NewGetCommand creates a new command to get an artifacts group metadata
func NewGetCommand(f *factory.Factory) *cobra.Command {

	opts := &options{
		f: f,
	}

	cmd := &cobra.Command{
		Use:     "get",
		Short:   f.Localizer.MustLocalize("group.get.cmd.description.short"),
		Long:    f.Localizer.MustLocalize("group.get.cmd.description.long"),
		Example: f.Localizer.MustLocalize("group.get.cmd.example"),
		Args:    cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) (err error) {
			var missingFlags []string

			if opts.outputFormat != "" && !flagutil.IsValidInput(opts.outputFormat, flagutil.ValidOutputFormats...) {
				return flagutil.InvalidValueError("output", opts.outputFormat, flagutil.ValidOutputFormats...)
			}

			if opts.groupId == "" {
				missingFlags = append(missingFlags, "group-id")
			}

			if !opts.f.IOStreams.CanPrompt() && len(missingFlags) > 0 {
				return flagutil.RequiredWhenNonInteractiveError(missingFlags...)
			}

			if len(missingFlags) > 0 {
				err = runInteractivePrompt(opts, missingFlags)
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

	flags.StringVarP(&opts.groupId, "group-id", "g", "", opts.f.Localizer.MustLocalize("group.cmd.get.flag.group-id"))
	flags.AddOutput(&opts.outputFormat)

	flags.AddRegistryInstance(&opts.registryID)

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
	request := a.GroupsApi.GetGroupById(opts.f.Context, opts.groupId)

	groupMetaData, _, err := request.Execute()
	if err != nil {
		return registrycmdutil.TransformInstanceError(err)
	}

	return dump.Formatted(opts.f.IOStreams.Out, opts.outputFormat, groupMetaData)
}

func runInteractivePrompt(opts *options, missingFlags []string) (err error) {

	if slices.Contains(missingFlags, "group-id") {
		settingNamePrompt := &survey.Input{
			Message: opts.f.Localizer.MustLocalize("group.cmd.get.input.group-id.message"),
		}

		err = survey.AskOne(settingNamePrompt, &opts.groupId)
		if err != nil {
			return err
		}
	}

	return nil
}
