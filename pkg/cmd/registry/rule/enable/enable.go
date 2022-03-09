package enable

import (
	"context"
	"net/http"

	"github.com/AlecAivazis/survey/v2"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/registry/registrycmdutil"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/registry/rule/rulecmdutil"
	"github.com/redhat-developer/app-services-cli/pkg/core/cmdutil"
	"github.com/redhat-developer/app-services-cli/pkg/core/cmdutil/flagutil"
	"github.com/redhat-developer/app-services-cli/pkg/core/ioutil/icon"
	"github.com/redhat-developer/app-services-cli/pkg/core/ioutil/iostreams"
	"github.com/redhat-developer/app-services-cli/pkg/core/localize"
	"github.com/redhat-developer/app-services-cli/pkg/core/logging"
	"github.com/redhat-developer/app-services-cli/pkg/core/servicecontext"
	"github.com/redhat-developer/app-services-cli/pkg/shared/connection"
	"github.com/redhat-developer/app-services-cli/pkg/shared/factory"
	"github.com/redhat-developer/app-services-cli/pkg/shared/profileutil"
	"github.com/spf13/cobra"

	registryinstanceclient "github.com/redhat-developer/app-services-sdk-go/registryinstance/apiv1internal/client"
)

type options struct {
	IO             *iostreams.IOStreams
	Connection     factory.ConnectionFunc
	Logger         logging.Logger
	localizer      localize.Localizer
	Context        context.Context
	ServiceContext servicecontext.IContext

	ruleType   string
	config     string
	registryID string

	artifactID string
	group      string
}

// NewEnableCommand creates a new command for enabling rule
// nolint:funlen
func NewEnableCommand(f *factory.Factory) *cobra.Command {

	opts := &options{
		IO:             f.IOStreams,
		Connection:     f.Connection,
		Logger:         f.Logger,
		localizer:      f.Localizer,
		Context:        f.Context,
		ServiceContext: f.ServiceContext,
	}

	cmd := &cobra.Command{
		Use:     "enable",
		Short:   f.Localizer.MustLocalize("registry.rule.enable.cmd.description.short"),
		Long:    f.Localizer.MustLocalize("registry.rule.enable.cmd.description.long"),
		Example: f.Localizer.MustLocalize("registry.rule.enable.cmd.example"),
		Args:    cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) (err error) {

			validator := rulecmdutil.Validator{
				Localizer: opts.localizer,
			}

			var missingFlags []string

			if opts.ruleType == "" {
				missingFlags = append(missingFlags, "rule-type")
			}
			if opts.config == "" {
				missingFlags = append(missingFlags, "config")
			}

			if !opts.IO.CanPrompt() && len(missingFlags) > 0 {
				return flagutil.RequiredWhenNonInteractiveError(missingFlags...)
			}

			if len(missingFlags) == 2 {
				err = runInteractivePrompt(opts)
				if err != nil {
					return err
				}
			} else if len(missingFlags) > 0 {
				return flagutil.RequiredWhenNonInteractiveError(missingFlags...)
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

			svcContext, err := opts.ServiceContext.Load()
			if err != nil {
				return err
			}

			profileHandler := &profileutil.ContextHandler{
				Context:   svcContext,
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

			return runEnable(opts)
		},
	}

	flags := rulecmdutil.NewFlagSet(cmd, f)

	flags.AddRegistryInstance(&opts.registryID)

	flags.AddArtifactID(&opts.artifactID)
	flags.AddGroup(&opts.group)
	flags.AddConfig(&opts.config)
	flags.AddRuleType(&opts.ruleType)

	return cmd

}

func runEnable(opts *options) error {

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

	rule := registryinstanceclient.Rule{
		Config: rulecmdutil.GetMappedConfigValue(opts.config),
		Type:   rulecmdutil.GetMappedRuleType(opts.ruleType),
	}

	if opts.artifactID == "" {

		opts.Logger.Info(
			opts.localizer.MustLocalize("registry.rule.enable.log.info.enabling.globalRules",
				localize.NewEntry("RuleType", opts.ruleType),
				localize.NewEntry("Configuration", opts.config),
			),
		)

		req := dataAPI.AdminApi.CreateGlobalRule(opts.Context)

		req = req.Rule(rule)

		httpRes, newErr = req.Execute()
		if httpRes != nil {
			defer httpRes.Body.Close()
		}
	} else {

		if opts.group == registrycmdutil.DefaultArtifactGroup {
			opts.Logger.Info(opts.localizer.MustLocalize("registry.artifact.common.message.no.group", localize.NewEntry("DefaultArtifactGroup", registrycmdutil.DefaultArtifactGroup)))
		}

		opts.Logger.Info(
			opts.localizer.MustLocalize("registry.rule.enable.log.info.enabling.artifactRules",
				localize.NewEntry("RuleType", opts.ruleType),
				localize.NewEntry("Configuration", opts.config),
				localize.NewEntry("ArtifactID", opts.artifactID),
			),
		)

		req := dataAPI.ArtifactRulesApi.CreateArtifactRule(opts.Context, opts.group, opts.artifactID)

		req = req.Rule(rule)

		httpRes, newErr = req.Execute()
		if httpRes != nil {
			defer httpRes.Body.Close()
		}
	}

	ruleErrHandler := &rulecmdutil.RuleErrHandler{
		Localizer: opts.localizer,
	}

	if newErr != nil {
		if httpRes == nil {
			return registrycmdutil.TransformInstanceError(newErr)
		}

		switch httpRes.StatusCode {
		case http.StatusNotFound:
			return ruleErrHandler.ArtifactNotFoundError(opts.artifactID)
		case http.StatusConflict:
			return ruleErrHandler.ConflictError(opts.ruleType)
		default:
			return registrycmdutil.TransformInstanceError(newErr)
		}

	}

	opts.Logger.Info(icon.SuccessPrefix(), opts.localizer.MustLocalize("registry.rule.enable.log.info.ruleEnabled"))

	return nil
}

func runInteractivePrompt(opts *options) (err error) {

	ruleTypePrompt := &survey.Select{
		Message: opts.localizer.MustLocalize("registry.rule.enable.input.ruleType.message"),
		Help:    opts.localizer.MustLocalize("registry.rule.common.flag.ruleType"),
		Options: rulecmdutil.ValidRuleTypes,
	}

	err = survey.AskOne(ruleTypePrompt, &opts.ruleType)
	if err != nil {
		return err
	}

	configOptions := rulecmdutil.ValidRuleConfigs[opts.ruleType]

	configPrompt := &survey.Select{
		Message: opts.localizer.MustLocalize("registry.rule.enable.input.config.message"),
		Help:    opts.localizer.MustLocalize("registry.rule.common.flag.config"),
		Options: configOptions,
	}

	err = survey.AskOne(configPrompt, &opts.config)
	if err != nil {
		return err
	}

	artifactIDPrompt := &survey.Input{
		Message: opts.localizer.MustLocalize("registry.rule.enable.input.artifactID.message"),
		Help:    opts.localizer.MustLocalize("registry.rule.enable.input.artifactID.help"),
	}

	err = survey.AskOne(artifactIDPrompt, &opts.artifactID)
	if err != nil {
		return err
	}

	groupPrompt := &survey.Input{
		Message: opts.localizer.MustLocalize("registry.rule.enable.input.group.message"),
		Help:    opts.localizer.MustLocalize("registry.rule.enable.input.group.help"),
		Default: registrycmdutil.DefaultArtifactGroup,
	}

	err = survey.AskOne(groupPrompt, &opts.group)
	if err != nil {
		return err
	}

	return nil
}
