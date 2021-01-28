package cloudproviderutil

import (
	kasclient "github.com/bf2fc6cc711aee1a0c2a/cli/pkg/api/kas/client"
)

// GetEnabledNames returns a list of cloud provider names from the enabled cloud providers
func GetEnabledNames(cloudProviders []kasclient.CloudProvider) []string {
	var cloudProviderNames = []string{}
	for _, provider := range cloudProviders {
		if provider.GetEnabled() {
			cloudProviderNames = append(cloudProviderNames, provider.GetName())
		}
	}
	return cloudProviderNames
}

// FindByName finds and returns a cloud provider item from the list by its name
func FindByName(cloudProviders []kasclient.CloudProvider, name string) *kasclient.CloudProvider {
	for _, p := range cloudProviders {
		if p.GetName() == name {
			return &p
		}
	}
	return nil
}
