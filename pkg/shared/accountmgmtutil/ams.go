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

func FetchQuotaCost(f *factory.Factory, billingModel string, cloudAccountID string, marketplace string, spec *remote.AmsConfig) (userQuotaSpec *QuotaSpec, err error) {

	var conn connection.Connection
	if conn, err = f.Connection(connection.DefaultConfigSkipMasAuth); err != nil {
		return nil, err
	}

	if billingModel == "" && (cloudAccountID != "" || marketplace != "") {
		billingModel = QuotaMarketplaceType
	} else if billingModel == "" && cloudAccountID == "" && marketplace == "" {
		billingModel = QuotaStandardType
	}

	quotaCostGet, err := fetchOrgQuotaCost(f.Context, conn)
	if err != nil {
		return nil, err
	}

	var userQuota amsclient.QuotaCost

	for _, quota := range quotaCostGet.GetItems() {
		relatedResources := quota.GetRelatedResources()
		for i := range relatedResources {
			if relatedResources[i].GetResourceName() == spec.ResourceName && relatedResources[i].GetProduct() == spec.InstanceQuotaID && relatedResources[i].GetBillingModel() == billingModel {
				userQuota = quota
			}
		}
	}

	if userQuota.GetQuotaId() == "" {
		return nil, errors.New("quota could not be found")
	}

	var cloudAccountExists bool

	if billingModel == QuotaMarketplaceType && marketplace != "" && cloudAccountID != "" {
		for _, cloudAccount := range userQuota.GetCloudAccounts() {
			if cloudAccount.GetCloudAccountId() == cloudAccountID && cloudAccount.GetCloudProviderId() == marketplace {
				cloudAccountExists = true
				break
			}
		}

		if !cloudAccountExists {
			return nil, errors.New("cloud account doesnt exist for the given marketplace provider")
		}
	}

	userQuotaSpec = &QuotaSpec{billingModel, int(userQuota.GetAllowed() - userQuota.GetConsumed()), billingModel, userQuota.CloudAccounts}

	return userQuotaSpec, nil
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
