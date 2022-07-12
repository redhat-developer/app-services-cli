package accountmgmtutil

import (
	"context"
	"errors"
	"fmt"

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

func GetUserSupportedInstanceType(f *factory.Factory, spec *remote.AmsConfig, marketplace MarketplaceInfo) (*QuotaSpec, error) {

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

	fmt.Println("trial quotas - ", len(trialQuotas))
	fmt.Println("standard quotas - ", len(standardQuotas))
	fmt.Println("marketplace quotas - ", len(marketplaceQuotas))

	availableOrgQuotas := &OrgQuotas{standardQuotas, marketplaceQuotas, trialQuotas}

	return SelectQuotaForUser(f, availableOrgQuotas, marketplace)
}

func SelectQuotaForUser(f *factory.Factory, orgQuota *OrgQuotas, marketplaceInfo MarketplaceInfo) (*QuotaSpec, error) {
	if len(orgQuota.StandardQuotas) == 0 && len(orgQuota.MarketplaceQuotas) == 0 {
		if len(orgQuota.TrialQuotas) == 0 {
			return nil, errors.New("no quotas available")
		} else if marketplaceInfo.BillingModel != "" {
			return nil, errors.New("only trial quotas are available")
		} else {
			// select a trial quota as all others are missing
			return &orgQuota.TrialQuotas[0], nil
		}
	}

	if len(orgQuota.MarketplaceQuotas) == 0 && len(orgQuota.StandardQuotas) > 0 {

		if marketplaceInfo.BillingModel == QuotaMarketplaceType || marketplaceInfo.Provider != "" || marketplaceInfo.CloudAccountID != "" {
			return nil, errors.New("no marketplace quotas available")
		}
		// select a standard quota
		return &orgQuota.StandardQuotas[0], nil
	}

	if len(orgQuota.StandardQuotas) == 0 && len(orgQuota.MarketplaceQuotas) > 0 {

		if marketplaceInfo.BillingModel == QuotaStandardType {
			return nil, errors.New("no standard quotas available")
		}

		marketplaceQuota, err := getMarketplaceQuota(orgQuota.MarketplaceQuotas, marketplaceInfo.Provider, marketplaceInfo.CloudAccountID)
		if err != nil {
			return nil, err
		}

		marketplaceQuota.CloudAccounts, err = pickCloudAccount(marketplaceQuota.CloudAccounts, marketplaceInfo)
		if err != nil {
			return nil, err
		}

		return marketplaceQuota, nil
	}

	if len(orgQuota.StandardQuotas) > 0 && len(orgQuota.MarketplaceQuotas) > 0 {
		if marketplaceInfo.BillingModel == QuotaStandardType {
			return &orgQuota.StandardQuotas[0], nil
		} else if marketplaceInfo.BillingModel == QuotaMarketplaceType || marketplaceInfo.Provider != "" || marketplaceInfo.CloudAccountID != "" {
			marketplaceQuota, err := getMarketplaceQuota(orgQuota.MarketplaceQuotas, marketplaceInfo.Provider, marketplaceInfo.CloudAccountID)
			if err != nil {
				return nil, err
			}

			marketplaceQuota.CloudAccounts, err = pickCloudAccount(marketplaceQuota.CloudAccounts, marketplaceInfo)
			if err != nil {
				return nil, err
			}

			return marketplaceQuota, nil
		} else {
			return nil, errors.New("you must specify a billing model")
		}
	}

	return &orgQuota.StandardQuotas[0], nil
}

func getMarketplaceQuota(marketplaceQuotas []QuotaSpec, provider string, accountID string) (*QuotaSpec, error) {
	if len(marketplaceQuotas) == 1 {
		if provider != "" || accountID != "" {
			marketplaceQuota, err := pickMarketplaceQuota(marketplaceQuotas, provider, accountID)
			if err != nil {
				return nil, err
			}
			return marketplaceQuota, nil
		}
		return &marketplaceQuotas[0], nil
	} else if len(marketplaceQuotas) > 1 && provider == "" && accountID == "" {
		return nil, errors.New("more than one marketplace quota is available. Please specify marketplace and account ID")
	} else {
		marketplaceQuota, err := pickMarketplaceQuota(marketplaceQuotas, provider, accountID)
		if err != nil {
			return nil, err
		}
		return marketplaceQuota, nil
	}

}

func pickMarketplaceQuota(marketplaceQuotas []QuotaSpec, provider string, accountID string) (*QuotaSpec, error) {

	matchedQuotas := []QuotaSpec{}

	for _, quota := range marketplaceQuotas {
		cloudAccounts := *quota.CloudAccounts
		for _, cloudAccount := range cloudAccounts {
			if provider != "" && accountID != "" {
				if *cloudAccount.CloudProviderId == provider && *cloudAccount.CloudAccountId == accountID {
					matchedQuotas = append(matchedQuotas, quota)
				}
			} else if provider != "" || accountID != "" {
				if *cloudAccount.CloudProviderId == provider || *cloudAccount.CloudAccountId == accountID {
					matchedQuotas = append(matchedQuotas, quota)
				}
			}
		}
	}

	if len(matchedQuotas) == 0 {
		return nil, errors.New("no quota found with the given cloud account")
	} else if len(matchedQuotas) > 1 {
		return nil, errors.New("multiple quota objects were found")
	}

	return &matchedQuotas[0], nil
}

func pickCloudAccount(cloudAccounts *[]amsclient.CloudAccount, market MarketplaceInfo) (*[]amsclient.CloudAccount, error) {

	if len(*cloudAccounts) == 1 {
		return cloudAccounts, nil
	}

	if len(*cloudAccounts) > 2 && market.Provider == "" && market.CloudAccountID == "" {
		return nil, errors.New("multiple cloud accounts found, specify provider and cloud account ID")
	}

	var matchedAccounts []amsclient.CloudAccount

	for _, cloudAccount := range *cloudAccounts {
		if market.Provider != "" || market.CloudAccountID != "" {
			if *cloudAccount.CloudProviderId == market.Provider || *cloudAccount.CloudAccountId == market.CloudAccountID {
				matchedAccounts = append(matchedAccounts, cloudAccount)
			}
		}
	}

	if len(matchedAccounts) > 1 {
		return nil, errors.New("multiple cloud accounts found")
	}

	return &matchedAccounts, nil
}

func GetOrganizationID(ctx context.Context, conn connection.Connection) (accountID string, err error) {
	account, _, err := conn.API().AccountMgmt().ApiAccountsMgmtV1CurrentAccountGet(ctx).
		Execute()
	if err != nil {
		return "", err
	}

	return account.Organization.GetId(), nil
}
