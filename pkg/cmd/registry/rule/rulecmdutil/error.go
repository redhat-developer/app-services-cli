package rulecmdutil

import (
	"github.com/redhat-developer/app-services-cli/pkg/core/localize"
)

type RuleErrHandler struct {
	Localizer  localize.Localizer
	InstanceID string
}

func (r *RuleErrHandler) ConflictError(ruleType string) error {
	return r.Localizer.MustLocalizeError("registry.rule.common.error.conflict", localize.NewEntry("Type", ruleType))
}

func (r *RuleErrHandler) ArtifactNotFoundError(artifactID string) error {
	return r.Localizer.MustLocalizeError("registry.rule.common.error.artifactNotFound", localize.NewEntry("ID", artifactID))
}

func (r *RuleErrHandler) RuleNotEnabled(ruleTYpe string) error {
	return r.Localizer.MustLocalizeError("registry.rule.common.error.notEnabled", localize.NewEntry("Type", ruleTYpe))
}
