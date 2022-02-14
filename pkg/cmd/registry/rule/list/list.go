package list

import (
	"context"
	"net/http"

	"github.com/redhat-developer/app-services-cli/pkg/cmd/registry/registrycmdutil"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/registry/rule/rulecmdutil"
	"github.com/redhat-developer/app-services-cli/pkg/core/config"
	"github.com/redhat-developer/app-services-cli/pkg/core/ioutil/dump"
	"github.com/redhat-developer/app-services-cli/pkg/core/ioutil/iostreams"
	"github.com/redhat-developer/app-services-cli/pkg/core/localize"
	"github.com/redhat-developer/app-services-cli/pkg/core/logging"
	"github.com/redhat-developer/app-services-cli/pkg/shared/connection"
	"github.com/redhat-developer/app-services-cli/pkg/shared/factory"
	"github.com/spf13/cobra"

	registryinstanceclient "github.com/redhat-developer/app-services-sdk-go/registryinstance/apiv1internal/client"
)

const (
	ruleValidity      = "VALIDITY"
	ruleCompatibility = "COMPATIBILITY"
)

// ruleRow is the details of a Service Registry rules needed to print to a table
type ruleRow struct {
	RuleType string `json:"ruleType" header:"Rule Type"`

	Description string `json:"description,omitempty" header:"Description"`

	Status string `json:"status" header:"Status"`
}

type options struct {
	IO         *iostreams.IOStreams
	Config     config.IConfig
	Connection factory.ConnectionFunc
	Logger     logging.Logger
	localizer  localize.Localizer
	Context    context.Context

	registryID string
	artifactID string
	group      string
}

var compatibilityRuleStatus = ruleRow{
	RuleType:    ruleCompatibility,
	Description: "Enforce a compatibility level when updating this artifact (for example, Backwards Compatibility).",
	Status:      "DISABLED",
}

var validityRuleStatus = ruleRow{
	RuleType:    ruleValidity,
	Description: "Ensure that content is valid when updating this artifact.",
	Status:      "DISABLED",
}

// NewListCommand creates a new command to view status of rules
func NewListCommand(f *factory.Factory) *cobra.Command {

	opts := &options{
		IO:         f.IOStreams,
		Config:     f.Config,
		Connection: f.Connection,
		Logger:     f.Logger,
		localizer:  f.Localizer,
		Context:    f.Context,
	}

	cmd := &cobra.Command{
		Use:     "list",
		Short:   f.Localizer.MustLocalize("registry.rule.list.cmd.description.short"),
		Long:    f.Localizer.MustLocalize("registry.rule.list.cmd.description.long"),
		Example: f.Localizer.MustLocalize("registry.rule.list.cmd.example"),
		Args:    cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) (err error) {

			cfg, err := opts.Config.Load()
			if err != nil {
				return err
			}

			instanceID, ok := cfg.GetServiceRegistryIdOk()
			if !ok {
				return opts.localizer.MustLocalizeError("artifact.cmd.common.error.noServiceRegistrySelected")
			}

			opts.registryID = instanceID

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

	var ruleErr = &rulecmdutil.RegistryRuleError{
		Localizer:  opts.localizer,
		InstanceID: opts.registryID,
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
			return ruleErr.ArtifactNotFoundError(opts.artifactID)
		default:
			return registrycmdutil.TransformInstanceError(newErr)
		}
	}

	for _, rule := range enabledRules {
		if rule == ruleValidity {
			validityRuleStatus.Status = "ENABLED"
		}
		if rule == ruleCompatibility {
			compatibilityRuleStatus.Status = "ENABLED"
		}
	}

	var adminRules = []ruleRow{validityRuleStatus, compatibilityRuleStatus}

	opts.Logger.Info()
	dump.Table(opts.IO.Out, adminRules)

	opts.Logger.Info()
	opts.Logger.Info(opts.localizer.MustLocalize("registry.rule.list.log.info.describeHint"))

	return nil
}
