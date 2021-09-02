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

// GetAllowedArtifactTypeEnumValuesAsString gets artifact types as string.
func GetAllowedArtifactTypeEnumValuesAsString() string {
	return strings.Join(AllowedArtifactTypeEnumValues, ", ")
}
