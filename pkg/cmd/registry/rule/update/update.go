package update

import (
	"context"
	"net/http"

	"github.com/redhat-developer/app-services-cli/pkg/cmd/registry/registrycmdutil"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/registry/rule/rulecmdutil"
	"github.com/redhat-developer/app-services-cli/pkg/core/cmdutil"
	"github.com/redhat-developer/app-services-cli/pkg/core/ioutil/icon"
	"github.com/redhat-developer/app-services-cli/pkg/core/ioutil/iostreams"
	"github.com/redhat-developer/app-services-cli/pkg/core/localize"
	"github.com/redhat-developer/app-services-cli/pkg/core/logging"
	"github.com/redhat-developer/app-services-cli/pkg/core/profile"
	"github.com/redhat-developer/app-services-cli/pkg/shared/connection"
	"github.com/redhat-developer/app-services-cli/pkg/shared/factory"
	"github.com/redhat-developer/app-services-cli/pkg/shared/profileutil"
	"github.com/spf13/cobra"

	registryinstanceclient "github.com/redhat-developer/app-services-sdk-go/registryinstance/apiv1internal/client"
)

type options struct {
	IO         *iostreams.IOStreams
	Connection factory.ConnectionFunc
	Logger     logging.Logger
	localizer  localize.Localizer
	Context    context.Context
	Profiles   profile.IContext

	ruleType   string
	config     string
	registryID string

	artifactID string
	group      string
}

// NewUpdateCommand creates a new command for updating rule
func NewUpdateCommand(f *factory.Factory) *cobra.Command {

	opts := &options{
		IO:         f.IOStreams,
		Connection: f.Connection,
		Logger:     f.Logger,
		localizer:  f.Localizer,
		Context:    f.Context,
		Profiles:   f.Profile,
	}

	cmd := &cobra.Command{
		Use:     "update",
		Short:   f.Localizer.MustLocalize("registry.rule.update.cmd.description.short"),
		Long:    f.Localizer.MustLocalize("registry.rule.update.cmd.description.long"),
		Example: f.Localizer.MustLocalize("registry.rule.update.cmd.example"),
		Args:    cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) (err error) {

			validator := rulecmdutil.Validator{
				Localizer: opts.localizer,
			}

			err = validator.ValidateRuleType(opts.ruleType)
			if err != nil {
				return err
			}

			isValid, configs := validator.IsValidRuleConfig(opts.ruleType, opts.config)
			if !isValid {
				return opts.localizer.MustLocalizeError("registry.rule.common.error.invalidRuleConfig",
					localize.NewEntry("RuleType", opts.ruleType),
					localize.NewEntry("Config", opts.config),
					localize.NewEntry("ValidConfigList", cmdutil.StringSliceToListStringWithQuotes(configs)),
				)
			}

			context, err := opts.Profiles.Load()
			if err != nil {
				return err
			}

			profileHandler := &profileutil.ContextHandler{
				Context:   context,
				Localizer: opts.localizer,
			}

			conn, err := opts.Connection(connection.DefaultConfigRequireMasAuth)
			if err != nil {
				return err
			}

			registryInstance, err := profileHandler.GetCurrentRegistryInstance(conn.API().ServiceRegistryMgmt())
			if err != nil {
				return err
			}

			opts.registryID = registryInstance.GetId()

			return runUpdate(opts)
		},
	}

	flags := rulecmdutil.NewFlagSet(cmd, f)

	flags.AddRegistryInstance(&opts.registryID)

	flags.AddArtifactID(&opts.artifactID)
	flags.AddGroup(&opts.group)

	_ = flags.AddConfig(&opts.config).Required()
	_ = flags.AddRuleType(&opts.ruleType).Required()

	return cmd
}

func runUpdate(opts *options) error {

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

		httpRes, newErr = updateGlobalRule(opts, dataAPI)
		if httpRes != nil {
			defer httpRes.Body.Close()
		}
	} else {

		request := dataAPI.ArtifactsApi.GetLatestArtifact(opts.Context, opts.group, opts.artifactID)
		_, httpRes, err = request.Execute()
		if httpRes != nil {
			defer httpRes.Body.Close()
		}

		if err != nil {
			return registrycmdutil.TransformInstanceError(err)
		}

		httpRes, newErr = updateArtifactRule(opts, dataAPI)
		if httpRes != nil {
			defer httpRes.Body.Close()
		}
	}

	ruleErr := &rulecmdutil.RuleErrHandler{
		Localizer: opts.localizer,
	}

	if newErr != nil {
		if httpRes == nil {
			return registrycmdutil.TransformInstanceError(newErr)
		}

		switch httpRes.StatusCode {
		case http.StatusNotFound:
			return ruleErr.RuleNotEnabled(opts.ruleType)
		case http.StatusConflict:
			return ruleErr.ConflictError(opts.ruleType)
		default:
			return registrycmdutil.TransformInstanceError(newErr)
		}
	}

	opts.Logger.Info(icon.SuccessPrefix(), opts.localizer.MustLocalize("registry.rule.update.log.info.ruleUpdated"))

	return nil

}

func updateGlobalRule(opts *options, dataAPI *registryinstanceclient.APIClient) (httpRes *http.Response, err error) {

	opts.Logger.Info(
		opts.localizer.MustLocalize("registry.rule.update.log.info.updating.globalRule",
			localize.NewEntry("RuleType", opts.ruleType),
			localize.NewEntry("Configuration", opts.config),
			localize.NewEntry("ID", opts.registryID),
		),
	)

	req := dataAPI.AdminApi.UpdateGlobalRuleConfig(opts.Context, *rulecmdutil.GetMappedRuleType(opts.ruleType))

	rule := registryinstanceclient.Rule{
		Config: rulecmdutil.GetMappedConfigValue(opts.config),
		Type:   rulecmdutil.GetMappedRuleType(opts.ruleType),
	}

	req = req.Rule2(rule)

	_, httpRes, err = req.Execute()

	return httpRes, err
}

func updateArtifactRule(opts *options, dataAPI *registryinstanceclient.APIClient) (httpRes *http.Response, err error) {

	opts.Logger.Info(
		opts.localizer.MustLocalize("registry.rule.update.log.info.updating.artifactRule",
			localize.NewEntry("RuleType", opts.ruleType),
			localize.NewEntry("Configuration", opts.config),
			localize.NewEntry("ArtifactID", opts.artifactID),
		),
	)

	req := dataAPI.ArtifactRulesApi.UpdateArtifactRuleConfig(opts.Context, opts.group, opts.artifactID, string(*rulecmdutil.GetMappedRuleType(opts.ruleType)))

	rule := registryinstanceclient.Rule{
		Config: rulecmdutil.GetMappedConfigValue(opts.config),
		Type:   rulecmdutil.GetMappedRuleType(opts.ruleType),
	}

	req = req.Rule2(rule)

	_, httpRes, err = req.Execute()
	if httpRes != nil {
		defer httpRes.Body.Close()
	}

	return httpRes, err
}
