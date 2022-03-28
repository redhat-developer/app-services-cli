package accountmgmtutil

import (
	"context"
	"errors"

	"github.com/redhat-developer/app-services-cli/pkg/shared/connection"

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

// QuotaSpec - contains quota name and remianing quota count
type QuotaSpec struct {
	name  string
	quota int
}

func GetUserSupportedInstanceTypes(ctx context.Context, spec remote.AmsConfig, conn connection.Connection) (quota []QuotaSpec, err error) {
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

	var quotas []QuotaSpec
	for _, quota := range quotaCostGet.GetItems() {
		for _, quotaResource := range quota.GetRelatedResources() {
			if quotaResource.GetResourceName() == spec.ResourceName {
				remainingQuota := int(quota.GetAllowed() - quota.GetConsumed())
				if quotaResource.GetProduct() == spec.TrialProductQuotaID {
					quotas = append(quotas, QuotaSpec{QuotaTrialType, remainingQuota})
				} else if quotaResource.GetProduct() == spec.InstanceQuotaID {
					quotas = append(quotas, QuotaSpec{QuotaStandardType, remainingQuota})
				}
			}
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
func PickInstanceType(amsTypes []QuotaSpec) (string, error) {
	if amsTypes == nil {
		// TODO better error
		return "", errors.New("Cannot pick the fight between AMS instance types. No one will win")
	}
	if len(amsTypes) == 0 {
		// TODO better error
		return "", errors.New("No fighters to pick the fight. Sorry")

	}

	for _, amsType := range amsTypes {
		if amsType.name == QuotaStandardType {
			return QuotaStandardType, nil
		}
	}

	return amsTypes[0].name, nil
}

func GetInstanceTypes(amsTypes []QuotaSpec) []string {
	var instanceTypes = make([]string, len(amsTypes))
	for i, amsType := range amsTypes {
		instanceTypes[i] = amsType.name
	}
	return instanceTypes
}
