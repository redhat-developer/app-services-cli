package owner

import (
	"github.com/AlecAivazis/survey/v2"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/registry/registrycmdutil"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/registry/rule/rulecmdutil"
	"github.com/redhat-developer/app-services-cli/pkg/core/cmdutil/flagutil"
	"github.com/redhat-developer/app-services-cli/pkg/core/ioutil/icon"
	"github.com/redhat-developer/app-services-cli/pkg/shared/contextutil"
	"github.com/redhat-developer/app-services-cli/pkg/shared/factory"
	"github.com/spf13/cobra"
	"k8s.io/utils/strings/slices"
)

type OwnerGetOptions struct {
	artifact string
	group    string

	registryID string

	f *factory.Factory
}

// NewGetCommand creates a new command to get a service registry setting
func NewGetCommand(f *factory.Factory) *cobra.Command {

	opts := &OwnerGetOptions{
		f: f,
	}

	cmd := &cobra.Command{
		Use:     "owner-get",
		Short:   f.Localizer.MustLocalize("artifact.cmd.owner.get.description.short"),
		Long:    f.Localizer.MustLocalize("artifact.cmd.owner.get.description.long"),
		Example: f.Localizer.MustLocalize("artifact.cmd.owner.get.example"),
		Args:    cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) (err error) {

			var missingFlags []string

			if opts.artifact == "" {
				missingFlags = append(missingFlags, "artifact-id")
			}

			if !opts.f.IOStreams.CanPrompt() && len(missingFlags) > 0 {
				return flagutil.RequiredWhenNonInteractiveError(missingFlags...)
			}

			if len(missingFlags) > 0 {
				err = runGetInteractivePrompt(opts, missingFlags)
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
	flags.AddGroup(&opts.group)
	flags.AddArtifactID(&opts.artifact)

	return cmd
}

func runGet(opts *OwnerGetOptions) error {
	conn, err := opts.f.Connection()
	if err != nil {
		return err
	}

	api := conn.API()

	a, _, err := api.ServiceRegistryInstance(opts.registryID)
	if err != nil {
		return err
	}

	request := a.MetadataApi.GetArtifactOwner(opts.f.Context, opts.group, opts.artifact)

	artifactOwner, _, err := request.Execute()
	if err != nil {
		return registrycmdutil.TransformInstanceError(err)
	}

	opts.f.Logger.Info(icon.SuccessPrefix(), *artifactOwner.Owner)
	return nil
}

func runGetInteractivePrompt(opts *OwnerGetOptions, missingFlags []string) (err error) {

	if slices.Contains(missingFlags, "artifact-id") {
		artifactIdPrompt := &survey.Input{
			Message: opts.f.Localizer.MustLocalize("artifact.cmd.owner.get.input.artifactId.message"),
		}

		err = survey.AskOne(artifactIdPrompt, &opts.artifact)
		if err != nil {
			return err
		}
	}

	return nil
}
