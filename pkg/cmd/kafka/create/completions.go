package create

import (
	"github.com/redhat-developer/app-services-cli/pkg/shared/accountmgmtutil"
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

// GetCloudProviderRegionCompletionValues returns the list of region IDs for a particular cloud provider
func GetCloudProviderRegionCompletionValues(f *factory.Factory, providerID string) (validRegions []string, directive cobra.ShellCompDirective) {
	if providerID == "" {
		return
	}

	validRegions, _ = GetEnabledCloudRegionIDs(f, providerID, nil)

	return validRegions, cobra.ShellCompDirectiveNoSpace
}

// GetKafkaSizeCompletionValues returns a list of valid kafka sizes for the specified region and ams instance types
func GetKafkaSizeCompletionValues(f *factory.Factory, providerID string, regionId string) (validSizes []string, directive cobra.ShellCompDirective) {
	directive = cobra.ShellCompDirectiveNoSpace

	// We need both values to provide a valid list of sizes
	if providerID == "" || regionId == "" {
		return nil, directive
	}

	err, constants := remote.GetRemoteServiceConstants(f.Context, f.Logger)
	if err != nil {
		return nil, directive
	}

	orgQuota, err := accountmgmtutil.GetOrgQuotas(f, &constants.Kafka.Ams)
	if err != nil {
		return nil, directive
	}

	userInstanceType, _ := accountmgmtutil.SelectQuotaForUser(f, orgQuota, accountmgmtutil.MarketplaceInfo{}, "")

	// Not including quota in this request as it takes very long time to list quota for all regions in suggestion mode
	validSizes, _ = FetchValidKafkaSizesLabels(f, providerID, regionId, *userInstanceType)

	return validSizes, cobra.ShellCompDirectiveNoSpace
}

func GetMarketplaceCompletionValues(f *factory.Factory) (validSizes []string, directive cobra.ShellCompDirective) {

	directive = cobra.ShellCompDirectiveNoSpace

	err, constants := remote.GetRemoteServiceConstants(f.Context, f.Logger)
	if err != nil {
		return nil, directive
	}

	orgQuota, err := accountmgmtutil.GetOrgQuotas(f, &constants.Kafka.Ams)
	if err != nil {
		return nil, directive
	}

	validMarketPlaces := FetchValidMarketplaces(orgQuota.MarketplaceQuotas, "")

	return validMarketPlaces, cobra.ShellCompDirectiveNoSpace
}

func GetMarketplaceAccountCompletionValues(f *factory.Factory, marketplace string) (validMarketplaceAcctIDs []string, directive cobra.ShellCompDirective) {

	directive = cobra.ShellCompDirectiveNoSpace

	if marketplace == "" {
		return validMarketplaceAcctIDs, directive
	}

	err, constants := remote.GetRemoteServiceConstants(f.Context, f.Logger)
	if err != nil {
		return nil, directive
	}

	orgQuota, err := accountmgmtutil.GetOrgQuotas(f, &constants.Kafka.Ams)
	if err != nil {
		return nil, directive
	}

	validMarketplaceAcctIDs = FetchValidMarketplaceAccounts(orgQuota.MarketplaceQuotas, marketplace)

	return validMarketplaceAcctIDs, cobra.ShellCompDirectiveNoSpace
}

func GetBillingModelCompletionValues(f *factory.Factory) (availableBillingModels []string, directive cobra.ShellCompDirective) {

	directive = cobra.ShellCompDirectiveNoSpace

	err, constants := remote.GetRemoteServiceConstants(f.Context, f.Logger)
	if err != nil {
		return nil, directive
	}

	orgQuota, err := accountmgmtutil.GetOrgQuotas(f, &constants.Kafka.Ams)
	if err != nil {
		return nil, directive
	}

	availableBillingModels = FetchSupportedBillingModels(orgQuota, "")

	return availableBillingModels, directive
}
