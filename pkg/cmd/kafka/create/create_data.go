package create

import (
	"context"

	"github.com/redhat-developer/app-services-cli/pkg/shared/connection"
	"github.com/redhat-developer/app-services-cli/pkg/shared/factory"
	kafkamgmtclient "github.com/redhat-developer/app-services-sdk-go/kafkamgmt/apiv1/client"
	"github.com/spf13/cobra"
)

// GetCloudProviderCompletionValues returns the list of supported cloud providers for creating a Kafka instance
// This is used in the cmd.RegisterFlagCompletionFunc for dynamic completion of --provider
func GetCloudProviderCompletionValues(f *factory.Factory) (validProviders []string, directive cobra.ShellCompDirective) {
	validProviders = []string{}
	directive = cobra.ShellCompDirectiveNoSpace

	conn, err := f.Connection(connection.DefaultConfigSkipMasAuth)
	if err != nil {
		return validProviders, directive
	}

	cloudProviderResponse, _, err := conn.API().KafkaMgmt().GetCloudProviders(f.Context).Execute()
	if err != nil {
		return validProviders, directive
	}

	cloudProviders := cloudProviderResponse.GetItems()
	validProviders = GetEnabledCloudProviderNames(cloudProviders)

	return validProviders, directive
}

// GetCloudProviderCompletionValues returns the list of region IDs for a particular cloud provider
func GetCloudProviderRegionCompletionValues(f *factory.Factory, providerID string) (validRegions []string, directive cobra.ShellCompDirective) {
	validRegions = []string{}
	directive = cobra.ShellCompDirectiveNoSpace

	if providerID == "" {
		return
	}

	conn, err := f.Connection(connection.DefaultConfigSkipMasAuth)
	if err != nil {
		return validRegions, directive
	}

	cloudProviderResponse, _, err := conn.API().
		KafkaMgmt().
		GetCloudProviderRegions(f.Context, providerID).
		Execute()
	if err != nil {
		return validRegions, directive
	}

	cloudProviders := cloudProviderResponse.GetItems()
	validRegions = GetEnabledCloudRegionIDs(cloudProviders, nil)

	return validRegions, directive
}

// GetCloudProviderSizeValues returns the list of region IDs for a particular cloud provider
func GetCloudProviderSizeValues(f *factory.Factory, providerID string, regionId string) (validSizes []string, directive cobra.ShellCompDirective) {
	directive = cobra.ShellCompDirectiveNoSpace

	if providerID == "" {
		return
	}
	conn, err := f.Connection(connection.DefaultConfigSkipMasAuth)
	if err != nil {
		return []string{}, directive
	}

	// ignores amsType like standard, developer etc.
	validSizes, _ = GetValidSizes(conn, f.Context, providerID, regionId, nil)

	return validSizes, directive
}

// return list of the valid instance sizes for the specifed region and ams instance types
func GetValidSizes(conn connection.Connection, context context.Context,
	providerID string, regionId string, amsType *string) ([]string, error) {
	validSizes := []string{}

	instanceTypes, _, err := conn.API().
		KafkaMgmt().
		GetInstanceTypesByCloudProviderAndRegion(context, providerID, regionId).
		Execute()
	if err != nil {
		return nil, err
	}

	for _, instanceType := range instanceTypes.GetItems() {
		if amsType != nil && amsType != instanceType.Id {
			continue
		}
		for _, instanceSize := range instanceType.GetSizes() {
			validSizes = append(validSizes, instanceSize.GetId())
		}
	}
	return validSizes, nil
}

// GetEnabledCloudProviderNames returns a list of cloud provider names from the enabled cloud providers
func GetEnabledCloudProviderNames(cloudProviders []kafkamgmtclient.CloudProvider) []string {
	cloudProviderNames := []string{}
	for _, provider := range cloudProviders {
		if provider.GetEnabled() {
			cloudProviderNames = append(cloudProviderNames, provider.GetName())
		}
	}
	return cloudProviderNames
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
func GetEnabledCloudRegionIDs(regions []kafkamgmtclient.CloudRegion, userAllowedInstanceTypes *[]string) []string {
	var regionIDs []string
	for i, region := range regions {
		if region.GetEnabled() {
			if userAllowedInstanceTypes != nil {
				if IsRegionAllowed(&regions[i], userAllowedInstanceTypes) {
					regionIDs = append(regionIDs, region.GetId())
				}
			} else {
				regionIDs = append(regionIDs, region.GetId())
			}
		}
	}
	return regionIDs
}

func IsRegionAllowed(region *kafkamgmtclient.CloudRegion, userAllowedInstanceTypes *[]string) bool {
	for _, userInstanceType := range *userAllowedInstanceTypes {
		if region.GetSupportedInstanceTypes() != nil {
			for _, instanceType := range region.GetSupportedInstanceTypes() {
				if instanceType == userInstanceType {
					return true
				}
			}
		}
	}
	return false
}
