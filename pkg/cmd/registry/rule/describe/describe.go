package describe

import (
	"context"
	"net/http"

	"github.com/redhat-developer/app-services-cli/pkg/cmd/registry/registrycmdutil"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/registry/rule/rulecmdutil"
	"github.com/redhat-developer/app-services-cli/pkg/core/ioutil/dump"
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
		Connection: f.Connection,
		Logger:     f.Logger,
		localizer:  f.Localizer,
		Context:    f.Context,
		Profiles:   f.Profile,
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

		opts.Logger.Info(opts.localizer.MustLocalize("registry.rule.describe.log.info.fetching.globalRule", localize.NewEntry("Type", opts.ruleType), localize.NewEntry("ID", opts.registryID)))

		req := dataAPI.AdminApi.GetGlobalRuleConfig(opts.Context, *rulecmdutil.GetMappedRuleType(opts.ruleType))

		rule, httpRes, err = req.Execute()
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

		opts.Logger.Info(opts.localizer.MustLocalize("registry.rule.describe.log.info.fetching.artifactRule", localize.NewEntry("Type", opts.ruleType)))

		ruleTypeParam := string(*rulecmdutil.GetMappedRuleType(opts.ruleType))

		req := dataAPI.ArtifactRulesApi.GetArtifactRuleConfig(opts.Context, opts.group, opts.artifactID, ruleTypeParam)

		rule, httpRes, err = req.Execute()
		if httpRes != nil {
			defer httpRes.Body.Close()
		}
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
