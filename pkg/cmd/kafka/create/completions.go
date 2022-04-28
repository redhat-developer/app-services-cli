package create

import (
	"github.com/redhat-developer/app-services-cli/pkg/shared/accountmgmtutil"
	"github.com/redhat-developer/app-services-cli/pkg/shared/connection"
	"github.com/redhat-developer/app-services-cli/pkg/shared/factory"
	"github.com/redhat-developer/app-services-cli/pkg/shared/remote"
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

// GetEnabledCloudProviderNames returns a list of valid kafka sizes for the specifed region and ams instance types
func GetKafkaSizeCompletionValues(f *factory.Factory, providerID string, regionId string) (validRegions []string, directive cobra.ShellCompDirective) {
	directive = cobra.ShellCompDirectiveNoSpace

	// We need both values to provide a valid list of sizes
	if providerID == "" || regionId == "" {
		return nil, directive
	}

	err, constants := remote.GetRemoteServiceConstants(f.Context, f.Logger)
	if err != nil {
		return nil, directive
	}

	conn, err := f.Connection(connection.DefaultConfigSkipMasAuth)
	if err != nil {
		return nil, directive
	}

	userInstanceType, err := accountmgmtutil.GetUserSupportedInstanceType(f.Context, constants.Kafka.Ams, conn)

	// Not including quota in this request as it takes very long time to list quota for all regions in suggestion mode
	validRegions, _ = GetValidKafkaSizes(f, providerID, regionId, *userInstanceType)

	return validRegions, cobra.ShellCompDirectiveNoSpace
}
