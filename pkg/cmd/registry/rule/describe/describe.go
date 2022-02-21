package describe

import (
	"context"
	"net/http"

	"github.com/redhat-developer/app-services-cli/pkg/cmd/registry/registrycmdutil"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/registry/rule/rulecmdutil"
	"github.com/redhat-developer/app-services-cli/pkg/core/config"
	"github.com/redhat-developer/app-services-cli/pkg/core/ioutil/dump"
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
	output     string
}

// NewDescribeCommand creates a new command for viewing configuration details of a rule
func NewDescribeCommand(f *factory.Factory) *cobra.Command {

	opts := &options{
		IO:         f.IOStreams,
		Config:     f.Config,
		Connection: f.Connection,
		Logger:     f.Logger,
		localizer:  f.Localizer,
		Context:    f.Context,
	}

	cmd := &cobra.Command{
		Use:     "describe",
		Short:   f.Localizer.MustLocalize("registry.rule.describe.cmd.description.short"),
		Long:    f.Localizer.MustLocalize("registry.rule.describe.cmd.description.long"),
		Example: f.Localizer.MustLocalize("registry.rule.describe.cmd.example"),
		Args:    cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) (err error) {

			validator := rulecmdutil.Validator{
				Localizer: opts.localizer,
			}

			err = validator.ValidateRuleType(opts.ruleType)
			if err != nil {
				return err
			}

			cfg, err := opts.Config.Load()
			if err != nil {
				return err
			}

			instanceID, ok := cfg.GetServiceRegistryIdOk()
			if !ok {
				return opts.localizer.MustLocalizeError("artifact.cmd.common.error.noServiceRegistrySelected")
			}

			opts.registryID = instanceID

			return runDescribe(opts)
		},
	}

	flags := rulecmdutil.NewFlagSet(cmd, f)

	_ = flags.AddRuleType(&opts.ruleType).Required()
	flags.AddRegistryInstance(&opts.registryID)

	flags.AddArtifactID(&opts.artifactID)
	flags.AddGroup(&opts.group)
	flags.AddOutput(&opts.output)

	return cmd
}

func runDescribe(opts *options) error {

	var httpRes *http.Response
	var err error
	var rule registryinstanceclient.Rule

	validator := rulecmdutil.Validator{
		Localizer: opts.localizer,
	}

	ruleErrHandler := &rulecmdutil.RuleErrHandler{
		Localizer: opts.localizer,
	}

	conn, err := opts.Connection(connection.DefaultConfigRequireMasAuth)
	if err != nil {
		return err
	}

	err = validator.ValidateRuleType(opts.ruleType)
	if err != nil {
		return err
	}

	dataAPI, _, err := conn.API().ServiceRegistryInstance(opts.registryID)
	if err != nil {
		return err
	}

	if opts.artifactID == "" {

		s := spinner.New(opts.IO.ErrOut, opts.localizer)
		s.SetLocalizedSuffix("registry.rule.describe.log.info.fetching.globalRule", localize.NewEntry("Type", opts.ruleType))
		s.Start()

		req := dataAPI.AdminApi.GetGlobalRuleConfig(opts.Context, *rulecmdutil.GetMappedRuleType(opts.ruleType))

		rule, httpRes, err = req.Execute()
		if httpRes != nil {
			defer httpRes.Body.Close()
		}

		s.Stop()
	} else {

		request := dataAPI.ArtifactsApi.GetLatestArtifact(opts.Context, opts.group, opts.artifactID)
		_, httpRes, err = request.Execute()
		if httpRes != nil {
			defer httpRes.Body.Close()
		}

		if err != nil {
			return registrycmdutil.TransformInstanceError(err)
		}

		s := spinner.New(opts.IO.ErrOut, opts.localizer)
		s.SetLocalizedSuffix("registry.rule.describe.log.info.fetching.artifactRule", localize.NewEntry("Type", opts.ruleType))
		s.Start()

		ruleTypeParam := string(*rulecmdutil.GetMappedRuleType(opts.ruleType))

		req := dataAPI.ArtifactRulesApi.GetArtifactRuleConfig(opts.Context, opts.group, opts.artifactID, ruleTypeParam)

		rule, httpRes, err = req.Execute()
		if httpRes != nil {
			defer httpRes.Body.Close()
		}

		s.Stop()
	}

	if err != nil {
		if httpRes == nil {
			return registrycmdutil.TransformInstanceError(err)
		}

		switch httpRes.StatusCode {
		case http.StatusNotFound:
			return ruleErrHandler.RuleNotEnabled(opts.ruleType)
		default:
			return registrycmdutil.TransformInstanceError(err)
		}
	}

	return dump.Formatted(opts.IO.Out, opts.output, rule)

}
