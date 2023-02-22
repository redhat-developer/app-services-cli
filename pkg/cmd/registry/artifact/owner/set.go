package owner

import (
	"github.com/AlecAivazis/survey/v2"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/registry/registrycmdutil"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/registry/rule/rulecmdutil"
	"github.com/redhat-developer/app-services-cli/pkg/core/cmdutil/flagutil"
	"github.com/redhat-developer/app-services-cli/pkg/core/ioutil/icon"
	"github.com/redhat-developer/app-services-cli/pkg/core/localize"
	"github.com/redhat-developer/app-services-cli/pkg/shared/contextutil"
	"github.com/redhat-developer/app-services-cli/pkg/shared/factory"
	registryinstanceclient "github.com/redhat-developer/app-services-sdk-core/app-services-sdk-go/registryinstance/apiv1internal/client"
	"github.com/spf13/cobra"
	"k8s.io/utils/strings/slices"
)

type options struct {
	artifact string
	group    string
	owner    string

	registryID string

	f *factory.Factory
}

// NewSetCommand creates a new command to set owner of an artifact
func NewSetCommand(f *factory.Factory) *cobra.Command {

	opts := &options{
		f: f,
	}

	cmd := &cobra.Command{
		Use:     "owner-set",
		Short:   f.Localizer.MustLocalize("artifact.cmd.owner.set.description.short"),
		Long:    f.Localizer.MustLocalize("artifact.cmd.owner.set.description.long"),
		Example: f.Localizer.MustLocalize("artifact.cmd.owner.set.example"),
		Args:    cobra.NoArgs,
		Hidden:  true,
		RunE: func(cmd *cobra.Command, _ []string) (err error) {

			var missingFlags []string

			if opts.artifact == "" {
				missingFlags = append(missingFlags, "artifact-id")
			}

			if opts.owner == "" {
				missingFlags = append(missingFlags, "owner")
			}

			if !opts.f.IOStreams.CanPrompt() && len(missingFlags) > 0 {
				return flagutil.RequiredWhenNonInteractiveError(missingFlags...)
			}

			if len(missingFlags) > 0 {
				err = runSetInteractivePrompt(opts, missingFlags)
				if err != nil {
					return err
				}
			}

			if opts.registryID != "" {
				return runSet(opts)
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
	flags.AddGroup(&opts.group)
	flags.AddArtifactID(&opts.artifact)

	flags.StringVar(&opts.owner, "owner", "", f.Localizer.MustLocalize("setting.set.cmd.flag.owner.description"))

	return cmd
}

func runSet(opts *options) error {
	conn, err := opts.f.Connection()
	if err != nil {
		return err
	}

	api := conn.API()

	a, _, err := api.ServiceRegistryInstance(opts.registryID)
	if err != nil {
		return err
	}

	request := a.MetadataApi.UpdateArtifactOwner(opts.f.Context, opts.group, opts.artifact)

	request = request.ArtifactOwner(registryinstanceclient.ArtifactOwner{
		Owner: &opts.owner,
	})

	_, err = request.Execute()
	if err != nil {
		return registrycmdutil.TransformInstanceError(err)
	}

	opts.f.Logger.Info(icon.SuccessPrefix(), opts.f.Localizer.MustLocalize("artifact.cmd.owner.set.success", localize.NewEntry("Name", opts.artifact)))
	return nil
}

func runSetInteractivePrompt(opts *options, missingFlags []string) (err error) {

	if slices.Contains(missingFlags, "artifact-id") {
		artifactIdPrompt := &survey.Input{
			Message: opts.f.Localizer.MustLocalize("artifact.cmd.owner.set.input.artifactId.message"),
		}

		err = survey.AskOne(artifactIdPrompt, &opts.artifact)
		if err != nil {
			return err
		}
	}
	if slices.Contains(missingFlags, "owner") {
		ownerIdPrompt := &survey.Input{
			Message: opts.f.Localizer.MustLocalize("artifact.cmd.owner.set.input.owner.message"),
		}

		err = survey.AskOne(ownerIdPrompt, &opts.owner)
		if err != nil {
			return err
		}
	}

	return nil
}
