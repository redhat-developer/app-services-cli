package util

import (
	"github.com/redhat-developer/app-services-cli/pkg/core/ioutil/dump"
	registryinstanceclient "github.com/redhat-developer/app-services-sdk-core/app-services-sdk-go/registryinstance/apiv1internal/client"
	"io"
)

type referenceRow struct {
	Name       string `json:"name" header:"Reference Name"`
	GroupId    string `json:"group" header:"Group"`
	ArtifactId string `json:"artifactId" header:"Artifact ID"`
	Version    string `json:"version" header:"Version"`
}

func PrettyPrintReferences(out io.Writer, references []registryinstanceclient.ArtifactReference) {
	rows := make([]referenceRow, len(references))
	version := "-"
	groupId := "default"
	for i, v := range references {
		if v.Version != nil {
			version = *v.Version
		}
		if v.GroupId != "" {
			groupId = v.GroupId
		}
		row := referenceRow{
			Name:       v.Name,
			GroupId:    groupId,
			ArtifactId: v.ArtifactId,
			Version:    version,
		}
		rows[i] = row
	}
	_, _ = out.Write([]byte("\n"))
	dump.Table(out, rows)
	_, _ = out.Write([]byte("\n"))
}
