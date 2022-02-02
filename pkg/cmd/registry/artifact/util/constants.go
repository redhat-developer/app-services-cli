package util

import (
	"strings"
)

const DefaultArtifactGroup = "default"

var AllowedArtifactTypeEnumValues = []string{
	"AVRO",
	"PROTOBUF",
	"JSON",
	"OPENAPI",
	"ASYNCAPI",
	"GRAPHQL",
	"KCONNECT",
	"WSDL",
	"XSD",
	"XML",
}

var AllowedArtifactStateEnumValues = []string{
	"ENABLED",
	"DISABLED",
	"DEPRECATED",
}

const (
	ViewerRole  = "Viewer"
	ManagerRole = "Manager"
	AdminRole   = "Admin"
)

var AllowedRoleTypeEnumValues = []string{
	ViewerRole,
	ManagerRole,
	AdminRole,
}

// GetAllowedArtifactTypeEnumValuesAsString gets artifact types as string.
func GetAllowedArtifactTypeEnumValuesAsString() string {
	return strings.Join(AllowedArtifactTypeEnumValues, ", ")
}

// GetAllowedArtifactTypeEnumValuesAsString gets artifact types as string.
func GetAllowedArtifactStateEnumValuesAsString() string {
	return strings.Join(AllowedArtifactStateEnumValues, ", ")
}

// GetAllowedRoleTypeEnumValuesAsString gets types of roles as string.
func GetAllowedRoleTypeEnumValuesAsString() string {
	return strings.Join(AllowedRoleTypeEnumValues, ", ")
}
