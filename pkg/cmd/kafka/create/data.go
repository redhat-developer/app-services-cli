package create

import (
	kafkamgmtclient "github.com/jackdelahunt/app-services-sdk-core/app-services-sdk-go/kafkamgmt/apiv1/client"
	"github.com/redhat-developer/app-services-cli/pkg/shared/accountmgmtutil"
	"github.com/redhat-developer/app-services-cli/pkg/shared/factory"
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
	case accountmgmtutil.QuotaStandardType:
		return StandardType
	case accountmgmtutil.QuotaMarketplaceType:
		return StandardType
	case accountmgmtutil.QuotaTrialType:
		return DeveloperType
	default:
		return DeveloperType
	}
}

func GetValidKafkaSizesLabels(sizes []kafkamgmtclient.SupportedKafkaSize) []string {
	var labels []string = make([]string, len(sizes))
	for i := range sizes {
		labels[i] = sizes[i].GetId()
	}

	return labels
}

func FetchValidKafkaSizesLabels(f *factory.Factory,
	providerID string, regionId string, amsType accountmgmtutil.QuotaSpec) ([]string, error) {
	sizes, err := FetchValidKafkaSizes(f, providerID, regionId, amsType)
	if err != nil {
		return nil, err
	}
	return GetValidKafkaSizesLabels(sizes), nil

}

func FetchSupportedBillingModels(userQuotas *accountmgmtutil.OrgQuotas, provider string) []string {

	billingModels := []string{}

	if len(userQuotas.StandardQuotas) > 0 {
		billingModels = append(billingModels, accountmgmtutil.QuotaStandardType)
	}

	if len(userQuotas.MarketplaceQuotas) > 0 {
		if provider != "" {
			for _, quota := range userQuotas.MarketplaceQuotas {
				for _, cloudAccount := range *quota.CloudAccounts {
					if cloudAccount.GetCloudProviderId() == provider || cloudAccount.GetCloudProviderId() == accountmgmtutil.RedHatMarketPlace {
						billingModels = append(billingModels, accountmgmtutil.QuotaMarketplaceType)
						return billingModels
					}
				}
			}
		} else {
			billingModels = append(billingModels, accountmgmtutil.QuotaMarketplaceType)
		}
	}

	return billingModels
}

func FetchValidMarketplaces(amsTypes []accountmgmtutil.QuotaSpec, provider string) []string {

	validMarketplaces := []string{}

	for _, quota := range amsTypes {
		if quota.CloudAccounts != nil {
			for _, cloudAccount := range *quota.CloudAccounts {
				if provider != "" {
					if *cloudAccount.CloudProviderId == provider || *cloudAccount.CloudProviderId == accountmgmtutil.RedHatMarketPlace {
						validMarketplaces = append(validMarketplaces, *cloudAccount.CloudProviderId)
					}
				} else {
					validMarketplaces = append(validMarketplaces, *cloudAccount.CloudProviderId)
				}
			}
		}
	}

	return unique(validMarketplaces)
}

func FetchValidMarketplaceAccounts(amsTypes []accountmgmtutil.QuotaSpec, marketplace string) []string {

	validAccounts := []string{}

	for _, quota := range amsTypes {
		if quota.CloudAccounts != nil {
			for _, cloudAccount := range *quota.CloudAccounts {
				if marketplace != "" {
					if cloudAccount.GetCloudProviderId() == marketplace {
						validAccounts = append(validAccounts, cloudAccount.GetCloudAccountId())
					}
				} else {
					validAccounts = append(validAccounts, cloudAccount.GetCloudAccountId())
				}
			}
		}
	}

	return unique(validAccounts)
}

func FetchInstanceTypes(f *factory.Factory, providerID string, regionId string) ([]kafkamgmtclient.SupportedKafkaInstanceType, error) {

	conn, err := f.Connection()
	if err != nil {
		return nil, err
	}

	instanceTypes, _, err := conn.API().
		KafkaMgmt().
		GetInstanceTypesByCloudProviderAndRegion(f.Context, providerID, regionId).
		Execute()
	if err != nil {
		return nil, err
	}

	return instanceTypes.GetInstanceTypes(), nil

}

// FetchValidKafkaSizes returns list of the valid instance sizes for the specified region and ams instance types
func FetchValidKafkaSizes(f *factory.Factory,
	providerID string, regionId string, amsType accountmgmtutil.QuotaSpec) ([]kafkamgmtclient.SupportedKafkaSize, error) {

	validSizes := []kafkamgmtclient.SupportedKafkaSize{}

	instanceTypes, err := FetchInstanceTypes(f, providerID, regionId)
	if err != nil {
		return nil, err
	}

	desiredInstanceType := mapAmsTypeToBackendType(&amsType)

	// Temporary workaround to be removed
	if desiredInstanceType == DeveloperType {
		for _, instanceType := range instanceTypes {
			if desiredInstanceType == instanceType.GetId() {
				instanceSizes := instanceType.GetSizes()
				for i := range instanceSizes {
					validSizes = append(validSizes, instanceSizes[i])
				}
			}
		}
	} else {
		for _, instanceType := range instanceTypes {
			if desiredInstanceType == instanceType.GetId() {
				instanceSizes := instanceType.GetSizes()
				for i := range instanceSizes {
					if instanceSizes[i].GetQuotaConsumed() <= int32(amsType.Quota) {
						validSizes = append(validSizes, instanceSizes[i])
					}
				}
			}
		}
	}

	return validSizes, nil
}

// GetEnabledCloudProviderNames returns a list of cloud provider names from the enabled cloud providers
func GetEnabledCloudProviderNames(f *factory.Factory) ([]string, error) {
	validProviders := []string{}
	conn, err := f.Connection()
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

// GetEnabledCloudRegionIDs extracts and returns a slice of the unique IDs of all enabled regions
func GetEnabledCloudRegionIDs(f *factory.Factory, providerID string, userAllowedAMSInstanceType *accountmgmtutil.QuotaSpec) ([]string, error) {
	validRegions := []string{}
	conn, err := f.Connection()
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

	regions := cloudProviderResponse.GetItems()

	var regionIDs []string
	for i, region := range regions {
		if region.GetEnabled() {
			if userAllowedAMSInstanceType != nil {
				if IsRegionAllowed(&regions[i], userAllowedAMSInstanceType) {
					regionIDs = append(regionIDs, region.GetId())
				}
			} else {
				regionIDs = append(regionIDs, region.GetId())
			}
		}
	}
	return regionIDs, err
}

func IsRegionAllowed(region *kafkamgmtclient.CloudRegion, userInstanceType *accountmgmtutil.QuotaSpec) bool {
	if len(region.Capacity) > 0 {
		backendInstanceType := mapAmsTypeToBackendType(userInstanceType)
		for _, capacityItem := range region.Capacity {
			if capacityItem.InstanceType == backendInstanceType {
				return true
			}
		}
	}
	return false
}

func unique(s []string) []string {
	inResult := make(map[string]bool)
	var result []string
	for _, str := range s {
		if _, ok := inResult[str]; !ok {
			inResult[str] = true
			result = append(result, str)
		}
	}
	return result
}
