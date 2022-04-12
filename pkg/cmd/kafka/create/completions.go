package create

import (
	"github.com/redhat-developer/app-services-cli/pkg/shared/factory"
	"github.com/spf13/cobra"
)

// GetCloudProviderCompletionValues returns the list of supported cloud providers for creating a Kafka instance
// This is used in the cmd.RegisterFlagCompletionFunc for dynamic completion of --provider
func GetCloudProviderCompletionValues(f *factory.Factory) (validProviders []string, directive cobra.ShellCompDirective) {
	validProviders, _ = GetEnabledCloudProviderNames(f)

	return validProviders, cobra.ShellCompDirectiveNoSpace
}

// GetCloudProviderCompletionValues returns the list of region IDs for a particular cloud provider
func GetCloudProviderRegionCompletionValues(f *factory.Factory, providerID string) (validRegions []string, directive cobra.ShellCompDirective) {
	if providerID == "" {
		return
	}

	validRegions, _ = GetEnabledCloudRegionIDs(f, providerID, nil)

	return validRegions, cobra.ShellCompDirectiveNoSpace
}
