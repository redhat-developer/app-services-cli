package accountmgmtutil

import (
	"context"
	"errors"

	"github.com/redhat-developer/app-services-cli/pkg/shared/connection"
	"github.com/redhat-developer/app-services-cli/pkg/shared/factory"

	amsclient "github.com/redhat-developer/app-services-sdk-go/accountmgmt/apiv1/client"

	"github.com/redhat-developer/app-services-cli/pkg/shared/remote"
)

func CheckTermsAccepted(ctx context.Context, spec *remote.AmsConfig, conn connection.Connection) (accepted bool, redirectURI string, err error) {
	termsReview, _, err := conn.API().AccountMgmt().
		ApiAuthorizationsV1SelfTermsReviewPost(ctx).
		SelfTermsReview(amsclient.SelfTermsReview{
			EventCode: &spec.TermsAndConditionsEventCode,
			SiteCode:  &spec.TermsAndConditionsSiteCode,
		}).Execute()
	if err != nil {
		return false, "", err
	}

	if !termsReview.GetTermsAvailable() && !termsReview.GetTermsRequired() {
		return true, "", nil
	}

	if !termsReview.HasRedirectUrl() {
		return false, "", errors.New("terms must be signed, but there is no terms URL")
	}

	return false, termsReview.GetRedirectUrl(), nil
}

// QuotaSpec - contains quota name and remaining quota count
type QuotaSpec struct {
	Name          string
	Quota         int
	BillingModel  string
	CloudAccounts *[]amsclient.CloudAccount
}

type MarketplaceInfo struct {
	BillingModel   string
	Provider       string
	CloudAccountID string
}

type OrgQuotas struct {
	StandardQuotas    []QuotaSpec
	MarketplaceQuotas []QuotaSpec
	TrialQuotas       []QuotaSpec
}

func fetchOrgQuotaCost(ctx context.Context, conn connection.Connection) (*amsclient.QuotaCostList, error) {
	orgId, err := GetOrganizationID(ctx, conn)
	if err != nil {
		return nil, err
	}

	quotaCostGet, _, err := conn.API().AccountMgmt().
		ApiAccountsMgmtV1OrganizationsOrgIdQuotaCostGet(ctx, orgId).
		FetchRelatedResources(true).
		FetchCloudAccounts(true).
		Execute()

	return &quotaCostGet, err

}

func GetOrgQuotas(f *factory.Factory, spec *remote.AmsConfig) (*OrgQuotas, error) {

	conn, err := f.Connection(connection.DefaultConfigSkipMasAuth)
	if err != nil {
		return nil, err
	}

	quotaCostGet, err := fetchOrgQuotaCost(f.Context, conn)
	if err != nil {
		return nil, err
	}

	var standardQuotas, marketplaceQuotas, trialQuotas []QuotaSpec
	for _, quota := range quotaCostGet.GetItems() {
		quotaResources := quota.GetRelatedResources()
		for i := range quotaResources {
			quotaResource := quotaResources[i]
			if quotaResource.GetResourceName() == spec.ResourceName {
				if quotaResource.GetProduct() == spec.TrialProductQuotaID {
					trialQuotas = append(trialQuotas, QuotaSpec{QuotaTrialType, 0, quotaResource.BillingModel, nil})
				} else if quotaResource.GetProduct() == spec.InstanceQuotaID {
					remainingQuota := int(quota.GetAllowed() - quota.GetConsumed())
					if quotaResource.BillingModel == QuotaStandardType {
						standardQuotas = append(standardQuotas, QuotaSpec{QuotaStandardType, remainingQuota, quotaResource.BillingModel, nil})
					} else if quotaResource.BillingModel == QuotaMarketplaceType {
						marketplaceQuotas = append(marketplaceQuotas, QuotaSpec{QuotaMarketplaceType, remainingQuota, quotaResource.BillingModel, quota.CloudAccounts})
					}
				}
			}
		}
	}

	availableOrgQuotas := &OrgQuotas{standardQuotas, marketplaceQuotas, trialQuotas}

	return availableOrgQuotas, nil
}

func SelectQuotaForUser(f *factory.Factory, orgQuota *OrgQuotas, marketplaceInfo MarketplaceInfo) (*QuotaSpec, error) {

	if len(orgQuota.StandardQuotas) == 0 && len(orgQuota.MarketplaceQuotas) == 0 {
		if marketplaceInfo.BillingModel != "" || marketplaceInfo.Provider != "" {
			return nil, f.Localizer.MustLocalizeError("kafka.create.quota.error.onlyTrialAvailable")
		}
		// select a trial quota as all other types are missing
		return &orgQuota.TrialQuotas[0], nil
	}

	if len(orgQuota.MarketplaceQuotas) == 0 && len(orgQuota.StandardQuotas) > 0 {
		if marketplaceInfo.BillingModel == QuotaMarketplaceType || marketplaceInfo.Provider != "" || marketplaceInfo.CloudAccountID != "" {
			return nil, f.Localizer.MustLocalizeError("kafka.create.quota.error.noMarketplace")
		}
		// select a standard quota
		return &orgQuota.StandardQuotas[0], nil
	}

	if len(orgQuota.StandardQuotas) == 0 && len(orgQuota.MarketplaceQuotas) > 0 {

		if marketplaceInfo.BillingModel == QuotaStandardType {
			return nil, f.Localizer.MustLocalizeError("kafka.create.quota.error.noStandard")
		}

		marketplaceQuota, err := getMarketplaceQuota(f, orgQuota.MarketplaceQuotas, marketplaceInfo)
		if err != nil {
			return nil, err
		}

		marketplaceQuota.CloudAccounts, err = pickCloudAccount(f, marketplaceQuota.CloudAccounts, marketplaceInfo)
		if err != nil {
			return nil, err
		}

		return marketplaceQuota, nil
	}

	if len(orgQuota.StandardQuotas) > 0 && len(orgQuota.MarketplaceQuotas) > 0 {

		if marketplaceInfo.BillingModel == QuotaStandardType {
			return &orgQuota.StandardQuotas[0], nil
		} else if marketplaceInfo.BillingModel == QuotaMarketplaceType || marketplaceInfo.Provider != "" || marketplaceInfo.CloudAccountID != "" {
			marketplaceQuota, err := getMarketplaceQuota(f, orgQuota.MarketplaceQuotas, marketplaceInfo)
			if err != nil {
				return nil, err
			}

			marketplaceQuota.CloudAccounts, err = pickCloudAccount(f, marketplaceQuota.CloudAccounts, marketplaceInfo)
			if err != nil {
				return nil, err
			}

			return marketplaceQuota, nil
		}
		return nil, f.Localizer.MustLocalizeError("kafka.create.quota.error.noBillingModel")
	}

	return &orgQuota.TrialQuotas[0], nil
}

func getMarketplaceQuota(f *factory.Factory, marketplaceQuotas []QuotaSpec, marketplace MarketplaceInfo) (*QuotaSpec, error) {
	if len(marketplaceQuotas) == 1 {
		if marketplace.Provider != "" && marketplace.CloudAccountID != "" {
			marketplaceQuota, err := pickMarketplaceQuota(f, marketplaceQuotas, marketplace)
			if err != nil {
				return nil, err
			}
			return marketplaceQuota, nil
		}
		return &marketplaceQuotas[0], nil
	} else if len(marketplaceQuotas) > 1 && marketplace.Provider == "" && marketplace.CloudAccountID == "" {
		return nil, f.Localizer.MustLocalizeError("kafka.create.quota.error.multipleMarketplaceQuotas")
	}

	marketplaceQuota, err := pickMarketplaceQuota(f, marketplaceQuotas, marketplace)
	if err != nil {
		return nil, err
	}
	return marketplaceQuota, nil

}

func pickMarketplaceQuota(f *factory.Factory, marketplaceQuotas []QuotaSpec, marketplace MarketplaceInfo) (*QuotaSpec, error) {

	matchedQuotas := []QuotaSpec{}

	for _, quota := range marketplaceQuotas {
		cloudAccounts := *quota.CloudAccounts
		for _, cloudAccount := range cloudAccounts {
			if *cloudAccount.CloudProviderId == marketplace.Provider && *cloudAccount.CloudAccountId == marketplace.CloudAccountID {
				matchedQuotas = append(matchedQuotas, quota)
			}
		}
	}

	if len(matchedQuotas) == 0 {
		return nil, f.Localizer.MustLocalizeError("kafka.create.quota.error.cloudAccountNotFound")
	}

	return &matchedQuotas[0], nil
}

func pickCloudAccount(f *factory.Factory, cloudAccounts *[]amsclient.CloudAccount, market MarketplaceInfo) (*[]amsclient.CloudAccount, error) {

	if len(*cloudAccounts) == 1 {
		return cloudAccounts, nil
	}

	if len(*cloudAccounts) > 2 && market.Provider == "" && market.CloudAccountID == "" {
		return nil, f.Localizer.MustLocalizeError("kafka.create.quota.error.multipleCloudAccounts")
	}

	var matchedAccounts []amsclient.CloudAccount

	for _, cloudAccount := range *cloudAccounts {
		if *cloudAccount.CloudProviderId == market.Provider || *cloudAccount.CloudAccountId == market.CloudAccountID {
			matchedAccounts = append(matchedAccounts, cloudAccount)
		}
	}

	return &matchedAccounts, nil
}

// FetchValidMarketplaces returns the marketplaces available to the user to create Kafka Instance
func FetchValidMarketplaces(amsTypes []QuotaSpec) []string {

	validMarketplaces := []string{}

	for _, quota := range amsTypes {
		if quota.CloudAccounts != nil {
			for _, cloudAccount := range *quota.CloudAccounts {
				validMarketplaces = append(validMarketplaces, *cloudAccount.CloudProviderId)
			}
		}
	}

	return unique(validMarketplaces)
}

// FetchValidMarketplaceAccounts returns the cloud accounts available for the specified marketplace
func FetchValidMarketplaceAccounts(amsTypes []QuotaSpec, marketplace string) []string {

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

func GetOrganizationID(ctx context.Context, conn connection.Connection) (accountID string, err error) {
	account, _, err := conn.API().AccountMgmt().ApiAccountsMgmtV1CurrentAccountGet(ctx).
		Execute()
	if err != nil {
		return "", err
	}

	return account.Organization.GetId(), nil
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
