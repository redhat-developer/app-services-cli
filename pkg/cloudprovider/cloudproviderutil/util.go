package cloudproviderutil

import (
	kafkamgmtclient "github.com/redhat-developer/app-services-sdk-go/kafkamgmt/apiv1/client"
)

// GetEnabledNames returns a list of cloud provider names from the enabled cloud providers
func GetEnabledNames(cloudProviders []kafkamgmtclient.CloudProvider) []string {
	var cloudProviderNames = []string{}
	for _, provider := range cloudProviders {
		if provider.GetEnabled() {
			cloudProviderNames = append(cloudProviderNames, provider.GetName())
		}
	}
	return cloudProviderNames
}

// FindByName finds and returns a cloud provider item from the list by its name
func FindByName(cloudProviders []kafkamgmtclient.CloudProvider, name string) *kafkamgmtclient.CloudProvider {
	for _, p := range cloudProviders {
		if p.GetName() == name {
			return &p
		}
	}
	return nil
}
