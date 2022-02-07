package enable

import (
	"context"
	"net/http"

	"github.com/redhat-developer/app-services-cli/pkg/cmd/registry/rule/rulecmdutil"
	"github.com/redhat-developer/app-services-cli/pkg/core/cmdutil"
	"github.com/redhat-developer/app-services-cli/pkg/core/config"
	"github.com/redhat-developer/app-services-cli/pkg/core/ioutil/iostreams"
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

	ruleType   string
	config     string
	registryID string

	artifactID string
	group      string
}

// NewEnableCommand creates a new command for enabling rule
func NewEnableCommand(f *factory.Factory) *cobra.Command {

	opts := &options{
		IO:         f.IOStreams,
		Config:     f.Config,
		Connection: f.Connection,
		Logger:     f.Logger,
		localizer:  f.Localizer,
		Context:    f.Context,
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

			cfg, err := opts.Config.Load()
			if err != nil {
				return err
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

			instanceID, ok := cfg.GetServiceRegistryIdOk()
			if !ok {
				return opts.localizer.MustLocalizeError("artifact.cmd.common.error.noServiceRegistrySelected")
			}

			opts.registryID = instanceID

			return runEnable(opts)
		},
	}

	flags := rulecmdutil.NewFlagSet(cmd, f)

	flags.AddRegistryInstance(&opts.registryID)

	flags.AddArtifactID(&opts.artifactID)
	flags.AddGroup(&opts.group)
	flags.AddConfig(&opts.config)
	flags.AddRuleType(&opts.ruleType)

	_ = cmd.MarkFlagRequired("rule-type")
	_ = cmd.MarkFlagRequired("config")

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

		req := dataAPI.AdminApi.CreateGlobalRule(opts.Context)

		req = req.Rule(rule)

		httpRes, newErr = req.Execute()
		if httpRes != nil {
			defer httpRes.Body.Close()
		}
	} else {

		req := dataAPI.ArtifactRulesApi.CreateArtifactRule(opts.Context, opts.group, opts.artifactID)

		req = req.Rule(rule)

		httpRes, newErr = req.Execute()
		if httpRes != nil {
			defer httpRes.Body.Close()
		}

	}

	if newErr != nil {
		if httpRes == nil {
			return newErr
		}

		operationTmplPair := localize.NewEntry("Operation", "enable")
		switch httpRes.StatusCode {
		case http.StatusUnauthorized:
			return opts.localizer.MustLocalizeError("registry.rule.common.error.unauthorized", operationTmplPair)
		case http.StatusForbidden:
			return opts.localizer.MustLocalizeError("registry.rule.common.error.forbidden", operationTmplPair)
		case http.StatusNotFound:
			return opts.localizer.MustLocalizeError("registry.rule.common.error.notFoundError", operationTmplPair)
		case http.StatusConflict:
			return opts.localizer.MustLocalizeError("registry.rule.common.error.conflict")
		case http.StatusInternalServerError:
			return opts.localizer.MustLocalizeError("registry.rule.common.error.internalServerError")
		default:
			return err
		}

	}

	opts.Logger.Info(opts.localizer.MustLocalize("registry.rule.enable.log.info.ruleEnabled"))

	return nil
}
