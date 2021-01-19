package cloudproviderutil

import (
	"github.com/bf2fc6cc711aee1a0c2a/cli/pkg/api/managedservices"
)

// GetEnabledNames returns a list of cloud provider names from the enabled cloud providers
func GetEnabledNames(cloudProviders []managedservices.CloudProvider) []string {
	var cloudProviderNames = []string{}
	for _, provider := range cloudProviders {
		if provider.GetEnabled() {
			cloudProviderNames = append(cloudProviderNames, provider.GetName())
		}
	}
	return cloudProviderNames
}

// FindByName finds and returns a cloud provider item from the list by its name
func FindByName(cloudProviders []managedservices.CloudProvider, name string) *managedservices.CloudProvider {
	for _, p := range cloudProviders {
		if p.GetName() == name {
			return &p
		}
	}
	return nil
}
