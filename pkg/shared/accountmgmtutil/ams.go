package accountmgmtutil

import (
	"context"
	"errors"
	"strings"

	"github.com/redhat-developer/app-services-cli/pkg/shared/connection"
	"k8s.io/utils/strings/slices"

	amsclient "github.com/redhat-developer/app-services-sdk-go/accountmgmt/apiv1/client"

	"github.com/redhat-developer/app-services-cli/pkg/shared/remote"
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
	orgId, err := GetOrganizationID(ctx, conn)
	if err != nil {
		return nil, err
	}

	quotaCostGet, _, err := conn.API().AccountMgmt().
		ApiAccountsMgmtV1OrganizationsOrgIdQuotaCostGet(ctx, orgId).
		FetchRelatedResources(true).
		Execute()
	if err != nil {
		return nil, err
	}

	var quotas []string
	for _, quota := range quotaCostGet.GetItems() {
		quotaId := strings.TrimSpace(quota.GetQuotaId())

		if quotaId == spec.TrialQuotaID {
			quotas = append(quotas, QuotaTrialType)
		}
		if quotaId == spec.InstanceQuotaID {
			quotas = append(quotas, QuotaStandardType)
		}
	}

	return quotas, err
}

func GetOrganizationID(ctx context.Context, conn connection.Connection) (accountID string, err error) {
	account, _, err := conn.API().AccountMgmt().ApiAccountsMgmtV1CurrentAccountGet(ctx).
		Execute()
	if err != nil {
		return "", err
	}

	return account.Organization.GetId(), nil
}

// PickInstanceType - Standard instance always wins!!!
// This function should not exist but it does represents some requirement
// from business to only pick one instance type when two are presented.
// When standard instance type is present in user instances it should always take precedence
func PickInstanceType(amsType *[]string) (string, error) {
	if amsType == nil {
		// TODO better error
		return "", errors.New("Cannot pick the fight between AMS instance types. No one will win")
	}
	if len(*amsType) == 0 {
		// TODO better error
		return "", errors.New("No fighters to pick the fight. Sorry")

	}

	if slices.Contains(*amsType, QuotaStandardType) {
		return QuotaStandardType, nil
	}

	return (*amsType)[0], nil
}
