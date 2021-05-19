package cloudregionutil

import (
	kafkamgmtv1 "github.com/redhat-developer/app-services-sdk-go/kafka/mgmt/apiv1"
)

// GetEnabledIDs extracts and returns a slice of the unique IDs of all enabled regions
func GetEnabledIDs(regions []kafkamgmtv1.CloudRegion) []string {
	var regionIDs = []string{}
	for _, region := range regions {
		if region.GetEnabled() {
			regionIDs = append(regionIDs, region.GetId())
		}
	}
	return regionIDs
}
