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
	Name         string
	Quota        int
	BillingModel string
}

func GetUserSupportedInstanceType(ctx context.Context, spec *remote.AmsConfig, conn connection.Connection) (quota *QuotaSpec, err error) {
	userInstanceTypes, err := GetUserSupportedInstanceTypes(ctx, spec, conn)
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

func GetUserSupportedInstanceTypes(ctx context.Context, spec *remote.AmsConfig, conn connection.Connection) (quota []QuotaSpec, err error) {

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
				if quotaResource.GetProduct() == spec.TrialProductQuotaID {
					quotas = append(quotas, QuotaSpec{QuotaTrialType, 0, quotaResource.BillingModel})
				} else if quotaResource.GetProduct() == spec.InstanceQuotaID {
					remainingQuota := int(quota.GetAllowed() - quota.GetConsumed())
					quotas = append(quotas, QuotaSpec{QuotaStandardType, remainingQuota, quotaResource.BillingModel})
				}
			}
		}
	}

	return BattleOfInstanceBillingModels(quotas), err
}

// This function selects the billing model that should be used
// It represents some requirement to always use the same standard billing models
// This function should not exist but it does represents some requirement that we cannot do on backend
func BattleOfInstanceBillingModels(quotas []QuotaSpec) []QuotaSpec {
	var betterQuotasMap map[string]*QuotaSpec = make(map[string]*QuotaSpec)
	alwaysWinsBillingModel := "standard"
	for i := 0; i < len(quotas); i++ {
		if quotas[i].BillingModel == alwaysWinsBillingModel {
			betterQuotasMap[quotas[i].Name] = &quotas[i]
		} else if betterQuotasMap[quotas[i].Name] == nil {
			betterQuotasMap[quotas[i].Name] = &quotas[i]

		}
	}
	var betterQuotas []QuotaSpec
	for _, v := range betterQuotasMap {
		betterQuotas = append(betterQuotas, *v)
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

	for _, amsType := range amsTypes {
		if amsType.Name == QuotaStandardType {
			return &amsType
		}
	}

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

func GetValidMarketplaceAcctIDs(ctx context.Context, connectionFunc factory.ConnectionFunc) (marketplaceAcctIDs []string, err error) {

	conn, err := connectionFunc(connection.DefaultConfigSkipMasAuth)
	if err != nil {
		return nil, err
	}

	quotaCostGet, err := fetchOrgQuotaCost(ctx, conn)
	if err != nil {
		return nil, err
	}

	for _, quota := range quotaCostGet.GetItems() {
		if len(quota.GetCloudAccounts()) > 0 {
			for _, cloudAccount := range quota.GetCloudAccounts() {
				marketplaceAcctIDs = append(marketplaceAcctIDs, cloudAccount.GetCloudAccountId())
			}
		}
	}

	return unique(marketplaceAcctIDs), err
}

func GetValidMarketplaces(ctx context.Context, connectionFunc factory.ConnectionFunc) (marketplaces []string, err error) {

	conn, err := connectionFunc(connection.DefaultConfigSkipMasAuth)
	if err != nil {
		return nil, err
	}

	quotaCostGet, err := fetchOrgQuotaCost(ctx, conn)
	if err != nil {
		return nil, err
	}

	for _, quota := range quotaCostGet.GetItems() {
		if len(quota.GetCloudAccounts()) > 0 {
			for _, cloudAccount := range quota.GetCloudAccounts() {
				marketplaces = append(marketplaces, cloudAccount.GetCloudProviderId())
			}
		}
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
