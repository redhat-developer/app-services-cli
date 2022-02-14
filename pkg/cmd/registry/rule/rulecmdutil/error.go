package rulecmdutil

import (
	"github.com/redhat-developer/app-services-cli/pkg/core/localize"
)

type RegistryRuleError struct {
	Localizer  localize.Localizer
	InstanceID string
}

func (r *RegistryRuleError) ConflictError(ruleType string) error {
	return r.Localizer.MustLocalizeError("registry.rule.common.error.conflict", localize.NewEntry("Type", ruleType))
}

func (r *RegistryRuleError) ArtifactNotFoundError(artifactID string) error {
	return r.Localizer.MustLocalizeError("registry.rule.common.error.artifactNotFound", localize.NewEntry("ID", artifactID))
}

func (r *RegistryRuleError) RuleNotEnabled(ruleTYpe string) error {
	return r.Localizer.MustLocalizeError("registry.rule.common.error.notEnabled", localize.NewEntry("Type", ruleTYpe))
}
