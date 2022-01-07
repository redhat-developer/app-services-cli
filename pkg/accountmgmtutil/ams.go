package accountmgmtutil

import (
	"context"
	"errors"

	"github.com/redhat-developer/app-services-cli/pkg/core/connection"

	amsclient "github.com/redhat-developer/app-services-sdk-go/accountmgmt/apiv1/client"

	"github.com/redhat-developer/app-services-cli/pkg/remote"
)

func CheckTermsAccepted(ctx context.Context, spec remote.AmsConfig, conn connection.Connection) (accepted bool, redirectURI string, err error) {
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

func GetUserSupportedInstanceTypes(ctx context.Context, spec remote.AmsConfig, conn connection.Connection) (quota []string, err error) {
	orgId, err := GetOrganizationId(ctx, conn)
	if err != nil {
		return nil, err
	}

	quotaCostGet, _, err := conn.API().AccountMgmt().
		ApiAccountsMgmtV1OrganizationsOrgIdQuotaCostGet(ctx, orgId).
		Search("").
		FetchRelatedResources(false).
		Execute()
	if err != nil {
		return nil, err
	}

	var quotas []string
	for _, quota := range quotaCostGet.GetItems() {
		if quota.Id == &spec.TrialQuotaID {
			quotas = append(quotas, QuotaTrialType)
		}
		if quota.Id == &spec.InstanceQuotaID {
			quotas = append(quotas, QuotaStandardType)
		}
	}
	if len(quotas) == 0 {
		return nil, errors.New("Your account missing quota for creating instance of specified type")
	}
	return quotas, err
}

func GetOrganizationId(ctx context.Context, conn connection.Connection) (accountId string, err error) {
	account, _, err := conn.API().AccountMgmt().ApiAccountsMgmtV1CurrentAccountGet(ctx).
		Execute()
	if err != nil {
		return "", err
	}

	if account.GetBanned() {
		return "", errors.New("Your account has been banned from using the App Services. If you believe this is an error, please contact our support team.")
	}

	return account.GetId(), nil
}
