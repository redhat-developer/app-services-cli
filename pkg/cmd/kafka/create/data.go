package create

import (
	"github.com/redhat-developer/app-services-cli/pkg/shared/accountmgmtutil"
	"github.com/redhat-developer/app-services-cli/pkg/shared/connection"
	"github.com/redhat-developer/app-services-cli/pkg/shared/factory"
	kafkamgmtclient "github.com/redhat-developer/app-services-sdk-go/kafkamgmt/apiv1/client"
)

// Types we use on backend to map AMS Quotas
type CloudProviderId = string

// Additional types that are used in backend to match
// AMS Quota instance types
const (
	// Matches QuotaTrialType = trial
	DeveloperType CloudProviderId = "developer"

	// Matches QuotaStandardType = standard
	StandardType CloudProviderId = "standard"

	// Deprecated by DeveloperType
	TrialType CloudProviderId = "eval"
)

// mapAmsTypeToBackendType - Cloud providers API is not using AMS types but some other values (CloudProviderValues)
func mapAmsTypeToBackendType(amsType *accountmgmtutil.QuotaSpec) CloudProviderId {
	switch amsType.Name {
	case accountmgmtutil.QuotaTrialType:
		return StandardType
	case accountmgmtutil.QuotaStandardType:
		return DeveloperType
	default:
		return TrialType
	}
}

// return list of the valid instance sizes for the specifed region and ams instance types
func GetValidKafkaSizes(f *factory.Factory,
	providerID string, regionId string, amsType *accountmgmtutil.QuotaSpec) ([]string, error) {

	conn, err := f.Connection(connection.DefaultConfigSkipMasAuth)
	if err != nil {
		return nil, err
	}

	validSizes := []string{}

	instanceTypes, _, err := conn.API().
		KafkaMgmt().
		GetInstanceTypesByCloudProviderAndRegion(f.Context, providerID, regionId).
		Execute()
	if err != nil {
		return nil, err
	}

	var desiredInstanceType string
	if amsType != nil {
		desiredInstanceType = mapAmsTypeToBackendType(amsType)
	} else {
		// If we do not know instance type from AMS we should assume standard instance type
		// This is because trials do not have sizes so it is best to suggest all standard sizes instead
		desiredInstanceType = mapAmsTypeToBackendType(
			&accountmgmtutil.QuotaSpec{
				Name:  accountmgmtutil.QuotaStandardType,
				Quota: 1,
			})
	}

	for _, instanceType := range instanceTypes.GetItems() {
		if desiredInstanceType == instanceType.GetId() {
			for _, instanceSize := range instanceType.GetSizes() {
				validSizes = append(validSizes, instanceSize.GetId())
			}
		}
	}
	return validSizes, nil
}

// GetEnabledCloudProviderNames returns a list of cloud provider names from the enabled cloud providers
func GetEnabledCloudProviderNames(f *factory.Factory) ([]string, error) {
	validProviders := []string{}
	conn, err := f.Connection(connection.DefaultConfigSkipMasAuth)
	if err != nil {
		return validProviders, err
	}

	cloudProviderResponse, _, err := conn.API().KafkaMgmt().GetCloudProviders(f.Context).Execute()
	if err != nil {
		return validProviders, err
	}

	cloudProviders := cloudProviderResponse.GetItems()
	cloudProviderNames := []string{}
	for _, provider := range cloudProviders {
		if provider.GetEnabled() {
			cloudProviderNames = append(cloudProviderNames, provider.GetName())
		}
	}
	return cloudProviderNames, err
}

// FindCloudProviderByName finds and returns a cloud provider item from the list by its name
func FindCloudProviderByName(cloudProviders []kafkamgmtclient.CloudProvider, name string) *kafkamgmtclient.CloudProvider {
	for _, p := range cloudProviders {
		if p.GetName() == name {
			return &p
		}
	}
	return nil
}

// GetEnabledCloudRegionIDs extracts and returns a slice of the unique IDs of all enabled regions
func GetEnabledCloudRegionIDs(f *factory.Factory, providerID string, userAllowedAMSInstanceType *accountmgmtutil.QuotaSpec) ([]string, error) {
	validRegions := []string{}
	conn, err := f.Connection(connection.DefaultConfigSkipMasAuth)
	if err != nil {
		return validRegions, err
	}

	cloudProviderResponse, _, err := conn.API().
		KafkaMgmt().
		GetCloudProviderRegions(f.Context, providerID).
		Execute()
	if err != nil {
		return validRegions, err
	}

	var serverInstanceType string
	if userAllowedAMSInstanceType != nil {
		serverInstanceType = mapAmsTypeToBackendType(userAllowedAMSInstanceType)
	}

	regions := cloudProviderResponse.GetItems()

	var regionIDs []string
	for i, region := range regions {
		if region.GetEnabled() {
			if serverInstanceType != "" {
				if IsRegionAllowed(&regions[i], serverInstanceType) {
					regionIDs = append(regionIDs, region.GetId())
				}
			} else {
				regionIDs = append(regionIDs, region.GetId())
			}
		}
	}
	return regionIDs, err
}

func IsRegionAllowed(region *kafkamgmtclient.CloudRegion, userInstanceType string) bool {
	if region.GetSupportedInstanceTypes() != nil {
		for _, instanceType := range region.GetSupportedInstanceTypes() {
			if instanceType == userInstanceType {
				return true
			}
		}
	}
	return false
}
