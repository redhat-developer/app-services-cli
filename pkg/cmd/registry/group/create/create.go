package create

import (
	"github.com/AlecAivazis/survey/v2"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/registry/registrycmdutil"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/registry/rule/rulecmdutil"
	"github.com/redhat-developer/app-services-cli/pkg/core/cmdutil/flagutil"
	"github.com/redhat-developer/app-services-cli/pkg/core/ioutil/icon"
	"github.com/redhat-developer/app-services-cli/pkg/core/localize"
	"github.com/redhat-developer/app-services-cli/pkg/shared/factory"
	registryinstanceclient "github.com/redhat-developer/app-services-sdk-core/app-services-sdk-go/registryinstance/apiv1internal/client"
	"github.com/spf13/cobra"
	"k8s.io/utils/strings/slices"

	"github.com/redhat-developer/app-services-cli/pkg/shared/contextutil"
)

type options struct {
	registryID string

	groupId     string
	description string
	properties  map[string]string

	f *factory.Factory
}

// NewCreateCommand creates a new command to create a new artifact group
func NewCreateCommand(f *factory.Factory) *cobra.Command {

	opts := &options{
		f: f,
	}

	cmd := &cobra.Command{
		Use:     "create",
		Short:   f.Localizer.MustLocalize("group.create.cmd.description.short"),
		Long:    f.Localizer.MustLocalize("group.create.cmd.description.long"),
		Example: f.Localizer.MustLocalize("group.create.cmd.example"),
		Args:    cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) (err error) {
			var missingFlags []string

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
				return runCreate(opts)
			}

			registryInstance, err := contextutil.GetCurrentRegistryInstance(f)
			if err != nil {
				return err
			}

			opts.registryID = registryInstance.GetId()

			return runCreate(opts)
		},
	}

	flags := rulecmdutil.NewFlagSet(cmd, f)

	flags.StringVarP(&opts.groupId, "group-id", "g", "", opts.f.Localizer.MustLocalize("group.cmd.create.flag.group-id"))
	flags.StringVarP(&opts.description, "description", "d", "", opts.f.Localizer.MustLocalize("group.cmd.create.flag.description"))
	flags.StringToStringVarP(&opts.properties, "properties", "p", map[string]string{}, opts.f.Localizer.MustLocalize("group.cmd.create.flag.properties"))

	flags.AddRegistryInstance(&opts.registryID)

	return cmd

}

func runCreate(opts *options) error {
	conn, err := opts.f.Connection()
	if err != nil {
		return err
	}

	api := conn.API()

	a, _, err := api.ServiceRegistryInstance(opts.registryID)
	if err != nil {
		return err
	}
	request := a.GroupsApi.CreateGroup(opts.f.Context)

	createGroupMetaData := registryinstanceclient.CreateGroupMetaData{
		Id:          opts.groupId,
		Description: &opts.description,
		Properties:  &opts.properties,
	}

	request = request.CreateGroupMetaData(createGroupMetaData)

	_, _, err = request.Execute()
	if err != nil {
		return registrycmdutil.TransformInstanceError(err)
	}

	opts.f.Logger.Info(icon.SuccessPrefix(), opts.f.Localizer.MustLocalize("group.cmd.create.log.info.created", localize.NewEntry("GroupId", opts.groupId)))

	return nil
}

func runInteractivePrompt(opts *options, missingFlags []string) (err error) {

	if slices.Contains(missingFlags, "group-id") {
		settingNamePrompt := &survey.Input{
			Message: opts.f.Localizer.MustLocalize("group.cmd.create.input.group-id.message"),
		}

		err = survey.AskOne(settingNamePrompt, &opts.groupId)
		if err != nil {
			return err
		}
	}

	return nil
}
