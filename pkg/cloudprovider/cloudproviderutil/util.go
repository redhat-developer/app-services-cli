package cloudproviderutil

import kafkamgmtv1 "github.com/redhat-developer/app-services-sdk-go/kafkamgmt/apiv1"

// GetEnabledNames returns a list of cloud provider names from the enabled cloud providers
func GetEnabledNames(cloudProviders []kafkamgmtv1.CloudProvider) []string {
	var cloudProviderNames = []string{}
	for _, provider := range cloudProviders {
		if provider.GetEnabled() {
			cloudProviderNames = append(cloudProviderNames, provider.GetName())
		}
	}
	return cloudProviderNames
}

// FindByName finds and returns a cloud provider item from the list by its name
func FindByName(cloudProviders []kafkamgmtv1.CloudProvider, name string) *kafkamgmtv1.CloudProvider {
	for _, p := range cloudProviders {
		if p.GetName() == name {
			return &p
		}
	}
	return nil
}
