package disable

import (
	"context"
	"net/http"

	"github.com/redhat-developer/app-services-cli/pkg/cmd/registry/registrycmdutil"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/registry/rule/rulecmdutil"
	"github.com/redhat-developer/app-services-cli/pkg/core/config"
	"github.com/redhat-developer/app-services-cli/pkg/core/ioutil/icon"
	"github.com/redhat-developer/app-services-cli/pkg/core/ioutil/iostreams"
	"github.com/redhat-developer/app-services-cli/pkg/core/ioutil/spinner"
	"github.com/redhat-developer/app-services-cli/pkg/core/localize"
	"github.com/redhat-developer/app-services-cli/pkg/core/logging"
	"github.com/redhat-developer/app-services-cli/pkg/shared/connection"
	"github.com/redhat-developer/app-services-cli/pkg/shared/factory"
	"github.com/spf13/cobra"

	registryinstanceclient "github.com/redhat-developer/app-services-sdk-go/registryinstance/apiv1internal/client"
)

type options struct {
	IO         *iostreams.IOStreams
	Config     config.IConfig
	Connection factory.ConnectionFunc
	Logger     logging.Logger
	localizer  localize.Localizer
	Context    context.Context

	ruleType string

	registryID string
	artifactID string
	group      string
}

// NewDisableCommand creates a new command for disabling rule
func NewDisableCommand(f *factory.Factory) *cobra.Command {

	opts := &options{
		IO:         f.IOStreams,
		Config:     f.Config,
		Connection: f.Connection,
		Logger:     f.Logger,
		localizer:  f.Localizer,
		Context:    f.Context,
	}

	cmd := &cobra.Command{
		Use:     "disable",
		Short:   f.Localizer.MustLocalize("registry.rule.disable.cmd.description.short"),
		Long:    f.Localizer.MustLocalize("registry.rule.disable.cmd.description.long"),
		Example: f.Localizer.MustLocalize("registry.rule.disable.cmd.example"),
		Args:    cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) (err error) {

			validator := rulecmdutil.Validator{
				Localizer: opts.localizer,
			}

			cfg, err := opts.Config.Load()
			if err != nil {
				return err
			}

			if err = validator.ValidateRuleType(opts.ruleType); err != nil && opts.ruleType != "" {
				return err
			}

			instanceID, ok := cfg.GetServiceRegistryIdOk()
			if !ok {
				return opts.localizer.MustLocalizeError("artifact.cmd.common.error.noServiceRegistrySelected")
			}

			opts.registryID = instanceID

			return runDisable(opts)
		},
	}

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

	conn, err := opts.Connection(connection.DefaultConfigRequireMasAuth)
	if err != nil {
		return err
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

		s := spinner.New(opts.IO.ErrOut, opts.localizer)
		s.SetLocalizedSuffix("registry.rule.disable.log.info.disabling.globalRule", localize.NewEntry("RuleType", opts.ruleType))
		s.Start()

		req := dataAPI.AdminApi.DeleteGlobalRule(opts.Context, *rulecmdutil.GetMappedRuleType(opts.ruleType))

		httpRes, err = req.Execute()

		s.Stop()
	} else {

		s := spinner.New(opts.IO.ErrOut, opts.localizer)
		s.SetLocalizedSuffix("registry.rule.disable.log.info.disabling.globalRules", localize.NewEntry("RuleType", opts.ruleType))
		s.Start()

		req := dataAPI.AdminApi.DeleteAllGlobalRules(opts.Context)

		httpRes, err = req.Execute()

		s.Stop()
	}

	return httpRes, err
}

func disableArtifactRule(opts *options, dataAPI *registryinstanceclient.APIClient) (httpRes *http.Response, err error) {
	if opts.ruleType != "" {

		s := spinner.New(opts.IO.ErrOut, opts.localizer)
		s.SetLocalizedSuffix(
			"registry.rule.disable.log.info.disabling.artifactRule",
			localize.NewEntry("RuleType", opts.ruleType),
			localize.NewEntry("ArtifactID", opts.artifactID),
		)
		s.Start()

		req := dataAPI.ArtifactRulesApi.DeleteArtifactRule(opts.Context, opts.group, opts.artifactID, string(*rulecmdutil.GetMappedRuleType(opts.ruleType)))

		httpRes, err = req.Execute()
		if httpRes != nil {
			defer httpRes.Body.Close()
		}

		s.Stop()
	} else {

		s := spinner.New(opts.IO.ErrOut, opts.localizer)
		s.SetLocalizedSuffix("registry.rule.disable.log.info.disabling.artifactRules", localize.NewEntry("ArtifactID", opts.artifactID))
		s.Start()

		req := dataAPI.ArtifactRulesApi.DeleteArtifactRules(opts.Context, opts.group, opts.artifactID)

		httpRes, err = req.Execute()
		if httpRes != nil {
			defer httpRes.Body.Close()
		}

		s.Stop()
	}

	return httpRes, err
}
