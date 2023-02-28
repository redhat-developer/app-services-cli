package util

import (
	"strings"
)

var AllowedArtifactStateEnumValues = []string{
	"ENABLED",
	"DISABLED",
	"DEPRECATED",
}

const (
	ViewerRole  = "viewer"
	ManagerRole = "manager"
	AdminRole   = "admin"
)

var AllowedRoleTypeEnumValues = []string{
	ViewerRole,
	ManagerRole,
	AdminRole,
}

func GetAllowedArtifactStateEnumValuesAsString() string {
	return strings.Join(AllowedArtifactStateEnumValues, ", ")
}

// GetAllowedRoleTypeEnumValuesAsString gets types of roles as string.
func GetAllowedRoleTypeEnumValuesAsString() string {
	return strings.Join(AllowedRoleTypeEnumValues, ", ")
}
