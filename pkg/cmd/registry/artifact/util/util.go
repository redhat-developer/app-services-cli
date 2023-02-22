package util

import (
	"fmt"

	"github.com/redhat-developer/app-services-cli/pkg/cmd/registry/registrycmdutil"
	registryinstanceclient "github.com/redhat-developer/app-services-sdk-core/app-services-sdk-go/registryinstance/apiv1internal/client"
	registrymgmtclient "github.com/redhat-developer/app-services-sdk-core/app-services-sdk-go/registrymgmt/apiv1/client"
)

// GetArtifactURL takes registry and artifact metadata to build URL to artifact in console
func GetArtifactURL(registry *registrymgmtclient.Registry, metadata *registryinstanceclient.ArtifactMetaData) (artifactURL string, ok bool) {

	group := metadata.GetGroupId()

	if group == "" {
		group = registrycmdutil.DefaultArtifactGroup
	}

	homeURL, ok := registry.GetBrowserUrlOk()

	if !ok {
		return "", false
	}

	artifactURL = fmt.Sprintf("%s/artifacts/%s/%s/versions/%s", *homeURL, group, metadata.Id, metadata.Version)

	return artifactURL, true
}
