package disable

import (
	"context"
	"net/http"

	"github.com/AlecAivazis/survey/v2"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/registry/registrycmdutil"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/registry/rule/rulecmdutil"
	"github.com/redhat-developer/app-services-cli/pkg/core/cmdutil/flagutil"
	"github.com/redhat-developer/app-services-cli/pkg/core/ioutil/icon"
	"github.com/redhat-developer/app-services-cli/pkg/core/ioutil/iostreams"
	"github.com/redhat-developer/app-services-cli/pkg/core/localize"
	"github.com/redhat-developer/app-services-cli/pkg/core/logging"
	"github.com/redhat-developer/app-services-cli/pkg/core/servicecontext"
	"github.com/redhat-developer/app-services-cli/pkg/shared/contextutil"
	"github.com/redhat-developer/app-services-cli/pkg/shared/factory"
	registryinstanceclient "github.com/redhat-developer/app-services-sdk-core/app-services-sdk-go/registryinstance/apiv1internal/client"
	"github.com/spf13/cobra"
)

type options struct {
	IO             *iostreams.IOStreams
	Connection     factory.ConnectionFunc
	Logger         logging.Logger
	localizer      localize.Localizer
	Context        context.Context
	ServiceContext servicecontext.IContext

	ruleType string

	registryID string
	artifactID string
	group      string

	skipConfirm bool
}

// NewDisableCommand creates a new command for disabling rule
func NewDisableCommand(f *factory.Factory) *cobra.Command {

	opts := &options{
		IO:             f.IOStreams,
		Connection:     f.Connection,
		Logger:         f.Logger,
		localizer:      f.Localizer,
		Context:        f.Context,
		ServiceContext: f.ServiceContext,
	}

	cmd := &cobra.Command{
		Use:     "disable",
		Short:   f.Localizer.MustLocalize("registry.rule.disable.cmd.description.short"),
		Long:    f.Localizer.MustLocalize("registry.rule.disable.cmd.description.long"),
		Example: f.Localizer.MustLocalize("registry.rule.disable.cmd.example"),
		Args:    cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) (err error) {

			if !opts.IO.CanPrompt() && !opts.skipConfirm {
				return flagutil.RequiredWhenNonInteractiveError("yes")
			}

			registryInstance, err := contextutil.GetCurrentRegistryInstance(f)
			if err != nil {
				return err
			}

			opts.registryID = registryInstance.GetId()

			return runDisable(opts)
		},
	}

	cmd.Flags().BoolVarP(&opts.skipConfirm, "yes", "y", false, opts.localizer.MustLocalize("registry.rule.disable.flag.yes"))

	flags := rulecmdutil.NewFlagSet(cmd, f)

	flags.AddRegistryInstance(&opts.registryID)

	flags.AddArtifactID(&opts.artifactID)
	flags.AddGroup(&opts.group)
	flags.AddRuleType(&opts.ruleType)

	return cmd
}

func runDisable(opts *options) error {

	var httpRes *http.Response
	var newErr error

	conn, err := opts.Connection()
	if err != nil {
		return err
	}

	if !opts.skipConfirm {
		prompt := &survey.Confirm{
			Message: opts.localizer.MustLocalize("registry.rule.disable.confirm"),
		}
		if promptErr := survey.AskOne(prompt, &opts.skipConfirm); err != nil {
			return promptErr
		}
	}

	dataAPI, _, err := conn.API().ServiceRegistryInstance(opts.registryID)
	if err != nil {
		return err
	}

	if opts.artifactID == "" {

		httpRes, newErr = disableGlobalRule(opts, dataAPI)
		if httpRes != nil {
			defer httpRes.Body.Close()
		}

	} else {

		if opts.group == registrycmdutil.DefaultArtifactGroup {
			opts.Logger.Info(opts.localizer.MustLocalize("registry.artifact.common.message.no.group", localize.NewEntry("DefaultArtifactGroup", registrycmdutil.DefaultArtifactGroup)))
		}

		request := dataAPI.ArtifactsApi.GetLatestArtifact(opts.Context, opts.group, opts.artifactID)
		_, httpRes, err = request.Execute()
		if httpRes != nil {
			defer httpRes.Body.Close()
		}

		if err != nil {
			return registrycmdutil.TransformInstanceError(err)
		}

		httpRes, newErr = disableArtifactRule(opts, dataAPI)
		if httpRes != nil {
			defer httpRes.Body.Close()
		}

	}

	if newErr != nil {
		if httpRes == nil {
			return registrycmdutil.TransformInstanceError(newErr)
		}

		ruleErr := &rulecmdutil.RuleErrHandler{
			Localizer: opts.localizer,
		}

		switch httpRes.StatusCode {
		case http.StatusNotFound:
			return ruleErr.RuleNotEnabled(opts.ruleType)
		default:
			return registrycmdutil.TransformInstanceError(newErr)
		}
	}

	opts.Logger.Info(icon.SuccessPrefix(), opts.localizer.MustLocalize("registry.rule.disable.log.info.success"))

	return nil

}

func disableGlobalRule(opts *options, dataAPI *registryinstanceclient.APIClient) (httpRes *http.Response, err error) {

	if opts.ruleType != "" {

		opts.Logger.Info(opts.localizer.MustLocalize("registry.rule.disable.log.info.disabling.globalRule", localize.NewEntry("RuleType", opts.ruleType), localize.NewEntry("ID", opts.registryID)))

		req := dataAPI.AdminApi.DeleteGlobalRule(opts.Context, *rulecmdutil.GetMappedRuleType(opts.ruleType))

		httpRes, err = req.Execute()
	} else {

		opts.Logger.Info(opts.localizer.MustLocalize("registry.rule.disable.log.info.disabling.globalRules", localize.NewEntry("ID", opts.registryID)))

		req := dataAPI.AdminApi.DeleteAllGlobalRules(opts.Context)

		httpRes, err = req.Execute()
	}

	return httpRes, err
}

func disableArtifactRule(opts *options, dataAPI *registryinstanceclient.APIClient) (httpRes *http.Response, err error) {
	if opts.ruleType != "" {

		opts.Logger.Info(
			opts.localizer.MustLocalize("registry.rule.disable.log.info.disabling.artifactRule",
				localize.NewEntry("RuleType", opts.ruleType),
				localize.NewEntry("ArtifactID", opts.artifactID),
			),
		)

		req := dataAPI.ArtifactRulesApi.DeleteArtifactRule(opts.Context, opts.group, opts.artifactID, string(*rulecmdutil.GetMappedRuleType(opts.ruleType)))

		httpRes, err = req.Execute()
		if httpRes != nil {
			defer httpRes.Body.Close()
		}

	} else {
		opts.Logger.Info(opts.localizer.MustLocalize("registry.rule.disable.log.info.disabling.artifactRules", localize.NewEntry("ArtifactID", opts.artifactID)))

		req := dataAPI.ArtifactRulesApi.DeleteArtifactRules(opts.Context, opts.group, opts.artifactID)

		httpRes, err = req.Execute()
		if httpRes != nil {
			defer httpRes.Body.Close()
		}
	}

	return httpRes, err
}
