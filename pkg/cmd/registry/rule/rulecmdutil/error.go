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

func (r *RegistryRuleError) UnathorizedError(operation string) error {
	return r.Localizer.MustLocalizeError("registry.rule.common.error.unauthorized", localize.NewEntry("Operation", operation))
}

func (r *RegistryRuleError) ForbiddenError(operation string) error {
	return r.Localizer.MustLocalizeError("registry.rule.common.error.forbidden", localize.NewEntry("Operation", operation))
}

func (r *RegistryRuleError) ServerError() error {
	return r.Localizer.MustLocalizeError("registry.rule.common.error.internalServerError")
}

func (r *RegistryRuleError) ArtifactNotFoundError(artifactID string) error {
	return r.Localizer.MustLocalizeError("registry.rule.common.error.artifactNotFound", localize.NewEntry("ID", artifactID))
}
