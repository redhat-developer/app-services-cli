package cloudregionutil

import (
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/api/kas/client"
)

// GetEnabledIDs extracts and returns a slice of the unique IDs of all enabled regions
func GetEnabledIDs(regions []kasclient.CloudRegion) []string {
	var regionIDs = []string{}
	for _, region := range regions {
		if region.GetEnabled() {
			regionIDs = append(regionIDs, region.GetId())
		}
	}
	return regionIDs
}
