package cloudregionutil

import (
	kafkamgmtclient "github.com/redhat-developer/app-services-sdk-go/kafkamgmt/apiv1/client"
)

// GetEnabledIDs extracts and returns a slice of the unique IDs of all enabled regions
func GetEnabledIDs(regions []kafkamgmtclient.CloudRegion) []string {
	regionIDs := []string{}
	for _, region := range regions {
		if region.GetEnabled() {
			regionIDs = append(regionIDs, region.GetId())
		}
	}
	return regionIDs
}
