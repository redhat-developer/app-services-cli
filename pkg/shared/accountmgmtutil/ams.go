package accountmgmtutil

import (
	"context"
	"errors"
	"sort"

	"github.com/redhat-developer/app-services-cli/pkg/shared/connection"

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

func GetUserSupportedInstanceType(ctx context.Context, spec *remote.AmsConfig, conn connection.Connection, billingModel string) (quota *QuotaSpec, err error) {
	userInstanceTypes, err := GetUserSupportedInstanceTypes(ctx, spec, conn, billingModel)
	if err != nil {
		return nil, err
	}

	amsType := PickInstanceType(userInstanceTypes)

	return amsType, nil
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

func GetUserSupportedInstanceTypes(ctx context.Context, spec *remote.AmsConfig, conn connection.Connection, billingModel string) (quota []QuotaSpec, err error) {

	quotaCostGet, err := fetchOrgQuotaCost(ctx, conn)
	if err != nil {
		return nil, err
	}

	var quotas []QuotaSpec
	for _, quota := range quotaCostGet.GetItems() {
		quotaResources := quota.GetRelatedResources()
		for i := range quotaResources {
			quotaResource := quotaResources[i]
			if quotaResource.GetResourceName() == spec.ResourceName {
				// nolint
				if quotaResource.GetProduct() == spec.TrialProductQuotaID {
					quotas = append(quotas, QuotaSpec{QuotaTrialType, 0, quotaResource.BillingModel, nil})
				} else if quotaResource.GetProduct() == spec.InstanceQuotaID && quotaResource.GetBillingModel() == "standard" {
					remainingQuota := int(quota.GetAllowed() - quota.GetConsumed())
					quotas = append(quotas, QuotaSpec{QuotaStandardType, remainingQuota, quotaResource.BillingModel, nil})
				} else if quotaResource.GetProduct() == spec.InstanceQuotaID && quotaResource.GetBillingModel() == "marketplace" {
					remainingQuota := int(quota.GetAllowed() - quota.GetConsumed())
					quotas = append(quotas, QuotaSpec{QuotaMarketplaceTYpe, remainingQuota, quotaResource.BillingModel, quota.CloudAccounts})
				}
			}
		}
	}

	return BattleOfInstanceBillingModels(quotas, billingModel), err
}

// This function selects the billing model that should be used
// It represents some requirement to always use the same standard billing models
// This function should not exist but it does represents some requirement that we cannot do on backend
func BattleOfInstanceBillingModels(quotas []QuotaSpec, alwaysWinsBillingModel string) []QuotaSpec {
	var betterQuotas []QuotaSpec

	if alwaysWinsBillingModel == "" {
		alwaysWinsBillingModel = "standard"
	}

	for i := 0; i < len(quotas); i++ {
		if quotas[i].BillingModel == alwaysWinsBillingModel {
			betterQuotas = append(betterQuotas, quotas[i])
		}
	}

	return betterQuotas
}

// PickInstanceType - Standard instance always wins!
// This function should not exist but it does represents some requirement
// from business to only pick one instance type when two are presented.
// When standard instance type is present in user instances it should always take precedence
func PickInstanceType(amsTypes []QuotaSpec) *QuotaSpec {
	if amsTypes == nil || len(amsTypes) == 0 {
		return nil
	}

	sort.Slice(amsTypes, func(i, j int) bool {
		return amsTypes[i].Quota > amsTypes[j].Quota
	})

	// There is chance of having multiple instances in the future
	// We will pick the first one as we do not know which one to pick
	return &amsTypes[0]
}

func GetOrganizationID(ctx context.Context, conn connection.Connection) (accountID string, err error) {
	account, _, err := conn.API().AccountMgmt().ApiAccountsMgmtV1CurrentAccountGet(ctx).
		Execute()
	if err != nil {
		return "", err
	}

	return account.Organization.GetId(), nil
}

func GetValidMarketplaceAcctIDs(userQuotaType *QuotaSpec, marketplace string) (marketplaceAcctIDs []string, err error) {

	for _, cloudAccount := range *userQuotaType.CloudAccounts {
		if marketplace != "" {
			if cloudAccount.GetCloudProviderId() == marketplace {
				marketplaceAcctIDs = append(marketplaceAcctIDs, cloudAccount.GetCloudAccountId())
			}
		}
	}

	return unique(marketplaceAcctIDs), err
}

func GetValidMarketplaces(userQuotaType *QuotaSpec) (marketplaces []string, err error) {

	for _, cloudAccount := range *userQuotaType.CloudAccounts {
		marketplaces = append(marketplaces, cloudAccount.GetCloudProviderId())
	}

	return unique(marketplaces), err
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
