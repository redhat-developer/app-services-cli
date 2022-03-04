package list

import (
	"context"
	"net/http"
	"strings"

	"github.com/redhat-developer/app-services-cli/pkg/cmd/registry/registrycmdutil"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/registry/rule/rulecmdutil"
	"github.com/redhat-developer/app-services-cli/pkg/core/ioutil/dump"
	"github.com/redhat-developer/app-services-cli/pkg/core/ioutil/iostreams"
	"github.com/redhat-developer/app-services-cli/pkg/core/localize"
	"github.com/redhat-developer/app-services-cli/pkg/core/logging"
	"github.com/redhat-developer/app-services-cli/pkg/shared/connection"
	"github.com/redhat-developer/app-services-cli/pkg/shared/factory"
	"github.com/spf13/cobra"

	"github.com/redhat-developer/app-services-cli/pkg/core/profile"
	"github.com/redhat-developer/app-services-cli/pkg/shared/profileutil"
	registryinstanceclient "github.com/redhat-developer/app-services-sdk-go/registryinstance/apiv1internal/client"
)

const (
	ruleValidity      = "validity"
	ruleCompatibility = "compatibility"
)

const (
	ruleDisabled = "disabled"
	ruleEnabled  = "enabled"
)

// ruleRow is the details of a Service Registry rules needed to print to a table
type ruleRow struct {
	RuleType string `json:"ruleType" header:"Rule Type"`

	Description string `json:"description,omitempty" header:"Description"`

	Status string `json:"status" header:"Status"`
}

type options struct {
	IO         *iostreams.IOStreams
	Connection factory.ConnectionFunc
	Logger     logging.Logger
	localizer  localize.Localizer
	Context    context.Context
	Profiles   profile.IContext

	registryID string
	artifactID string
	group      string
}

// NewListCommand creates a new command to view status of rules
func NewListCommand(f *factory.Factory) *cobra.Command {

	opts := &options{
		IO:         f.IOStreams,
		Connection: f.Connection,
		Logger:     f.Logger,
		localizer:  f.Localizer,
		Context:    f.Context,
		Profiles:   f.Profile,
	}

	cmd := &cobra.Command{
		Use:     "list",
		Short:   f.Localizer.MustLocalize("registry.rule.list.cmd.description.short"),
		Long:    f.Localizer.MustLocalize("registry.rule.list.cmd.description.long"),
		Example: f.Localizer.MustLocalize("registry.rule.list.cmd.example"),
		Args:    cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) (err error) {

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

			return runList(opts)
		},
	}

	flags := rulecmdutil.NewFlagSet(cmd, f)

	flags.AddRegistryInstance(&opts.registryID)

	flags.AddArtifactID(&opts.artifactID)
	flags.AddGroup(&opts.group)

	return cmd

}

func runList(opts *options) error {

	var httpRes *http.Response
	var newErr error
	var enabledRules []registryinstanceclient.RuleType

	ruleErrHandler := &rulecmdutil.RuleErrHandler{
		Localizer: opts.localizer,
	}

	conn, err := opts.Connection(connection.DefaultConfigRequireMasAuth)
	if err != nil {
		return err
	}

	dataAPI, _, err := conn.API().ServiceRegistryInstance(opts.registryID)
	if err != nil {
		return err
	}

	if opts.artifactID == "" {

		opts.Logger.Info(opts.localizer.MustLocalize("registry.rule.list.log.info.fetching.globalRules"))

		req := dataAPI.AdminApi.ListGlobalRules(opts.Context)

		enabledRules, httpRes, newErr = req.Execute()
		if httpRes != nil {
			defer httpRes.Body.Close()
		}
	} else {
		if opts.group == registrycmdutil.DefaultArtifactGroup {
			opts.Logger.Info(opts.localizer.MustLocalize("registry.artifact.common.message.no.group", localize.NewEntry("DefaultArtifactGroup", registrycmdutil.DefaultArtifactGroup)))
		}

		opts.Logger.Info(opts.localizer.MustLocalize("registry.rule.list.log.info.fetching.artifactRules"))

		req := dataAPI.ArtifactRulesApi.ListArtifactRules(opts.Context, opts.group, opts.artifactID)

		enabledRules, httpRes, newErr = req.Execute()
		if httpRes != nil {
			defer httpRes.Body.Close()
		}
	}

	if newErr != nil {
		if httpRes == nil {
			return registrycmdutil.TransformInstanceError(newErr)
		}

		switch httpRes.StatusCode {
		case http.StatusNotFound:
			return ruleErrHandler.ArtifactNotFoundError(opts.artifactID)
		default:
			return registrycmdutil.TransformInstanceError(newErr)
		}
	}

	compatibilityRuleStatus := ruleRow{
		RuleType:    ruleCompatibility,
		Description: opts.localizer.MustLocalize("registry.rule.list.compatibilityRule.description"),
		Status:      ruleDisabled,
	}

	validityRuleStatus := ruleRow{
		RuleType:    ruleValidity,
		Description: opts.localizer.MustLocalize("registry.rule.list.validityRule.description"),
		Status:      ruleDisabled,
	}

	for _, rule := range enabledRules {
		if strings.EqualFold(string(rule), ruleValidity) {
			validityRuleStatus.Status = ruleEnabled
		}
		if strings.EqualFold(string(rule), ruleCompatibility) {
			compatibilityRuleStatus.Status = ruleEnabled
		}
	}

	adminRules := []ruleRow{validityRuleStatus, compatibilityRuleStatus}

	opts.Logger.Info()
	dump.Table(opts.IO.Out, adminRules)

	opts.Logger.Info()
	opts.Logger.Info(opts.localizer.MustLocalize("registry.rule.list.log.info.describeHint"))

	return nil
}
