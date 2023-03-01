package accountmgmtutil

import (
	"context"
	"errors"
	"github.com/redhat-developer/app-services-cli/pkg/shared/connection"
	"github.com/redhat-developer/app-services-cli/pkg/shared/factory"

	amsclient "github.com/redhat-developer/app-services-sdk-core/app-services-sdk-go/accountmgmt/apiv1/client"

	"github.com/redhat-developer/app-services-cli/pkg/shared/remote"
)

const (
	RedHatMarketPlace = "rhm"
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
	EvalQuotas        []QuotaSpec
	EnterpriseQuotas  []QuotaSpec
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

	conn, err := f.Connection()
	if err != nil {
		return nil, err
	}

	quotaCostGet, err := fetchOrgQuotaCost(f.Context, conn)
	if err != nil {
		return nil, err
	}

	//this should be refactored and base of the logic by the billing model information that's returned by KFM for each supported instance type
	var standardQuotas, marketplaceQuotas, trialQuotas, evalQuotas, enterpriseQuotas []QuotaSpec
	for _, quota := range quotaCostGet.GetItems() {
		quotaResources := quota.GetRelatedResources()
		for i := range quotaResources {
			quotaResource := quotaResources[i]
			if quotaResource.GetResourceName() == spec.ResourceName {
				switch quotaResource.GetProduct() {
				case spec.TrialProductQuotaID:
					trialQuotas = append(trialQuotas, QuotaSpec{QuotaTrialType, 0, quotaResource.BillingModel, nil})
				case spec.InstanceQuotaID:
					remainingQuota := int(quota.GetAllowed() - quota.GetConsumed())
					if quotaResource.BillingModel == QuotaStandardType {
						standardQuotas = append(standardQuotas, QuotaSpec{QuotaStandardType, remainingQuota, quotaResource.BillingModel, nil})
					} else if quotaResource.BillingModel == QuotaMarketplaceType {
						marketplaceQuotas = append(marketplaceQuotas, QuotaSpec{QuotaMarketplaceType, remainingQuota, quotaResource.BillingModel, quota.CloudAccounts})
					}
				case "RHOSAKEval":
					remainingQuota := int(quota.GetAllowed() - quota.GetConsumed())
					evalQuotas = append(evalQuotas, QuotaSpec{QuotaEvalType, remainingQuota, quotaResource.BillingModel, quota.CloudAccounts})
				case spec.EnterpriseProductQuotaID:
					remainingQuota := int(quota.GetAllowed() - quota.GetConsumed())
					enterpriseQuotas = append(enterpriseQuotas, QuotaSpec{QuotaEnterpriseType, remainingQuota, quotaResource.BillingModel, nil})
				}
			}
		}
	}

	availableOrgQuotas := &OrgQuotas{standardQuotas, marketplaceQuotas, trialQuotas, evalQuotas, enterpriseQuotas}

	return availableOrgQuotas, nil
}

// nolint:funlen
func SelectQuotaForUser(f *factory.Factory, orgQuota *OrgQuotas, marketplaceInfo MarketplaceInfo, provider string) (*QuotaSpec, error) {

	if len(orgQuota.StandardQuotas) == 0 && len(orgQuota.MarketplaceQuotas) == 0 && len(orgQuota.EvalQuotas) == 0 && len(orgQuota.EnterpriseQuotas) == 0 {
		if marketplaceInfo.BillingModel != "" || marketplaceInfo.Provider != "" {
			return nil, f.Localizer.MustLocalizeError("kafka.create.quota.error.onlyTrialAvailable")
		}
		// select a trial quota as all other types are missing
		return &orgQuota.TrialQuotas[0], nil
	}

	if len(orgQuota.MarketplaceQuotas) == 0 && len(orgQuota.StandardQuotas) > 0 && len(orgQuota.EvalQuotas) == 0 && len(orgQuota.EnterpriseQuotas) == 0 {
		if marketplaceInfo.BillingModel == QuotaMarketplaceType || marketplaceInfo.Provider != "" || marketplaceInfo.CloudAccountID != "" {
			return nil, f.Localizer.MustLocalizeError("kafka.create.quota.error.noMarketplace")
		}

		if marketplaceInfo.BillingModel == QuotaEvalType {
			return nil, f.Localizer.MustLocalizeError("kafka.create.quota.error.noEval")
		}
		// select a standard quota
		return &orgQuota.StandardQuotas[0], nil
	}

	if len(orgQuota.MarketplaceQuotas) == 0 && len(orgQuota.StandardQuotas) == 0 && len(orgQuota.EvalQuotas) > 0 && len(orgQuota.EnterpriseQuotas) == 0 {
		if marketplaceInfo.BillingModel == QuotaMarketplaceType || marketplaceInfo.Provider != "" || marketplaceInfo.CloudAccountID != "" {
			return nil, f.Localizer.MustLocalizeError("kafka.create.quota.error.noMarketplace")
		}

		if marketplaceInfo.BillingModel == QuotaStandardType {
			return nil, f.Localizer.MustLocalizeError("kafka.create.quota.error.noStandard")
		}

		return &orgQuota.EvalQuotas[0], nil
	}

	if len(orgQuota.StandardQuotas) == 0 && len(orgQuota.MarketplaceQuotas) > 0 && len(orgQuota.EvalQuotas) == 0 && len(orgQuota.EnterpriseQuotas) == 0 {

		if marketplaceInfo.BillingModel == QuotaStandardType {
			return nil, f.Localizer.MustLocalizeError("kafka.create.quota.error.noStandard")
		}

		if marketplaceInfo.BillingModel == QuotaEvalType {
			return nil, f.Localizer.MustLocalizeError("kafka.create.quota.error.noEval")
		}

		var filteredMarketPlaceQuotas []QuotaSpec

		if provider != "" {
			for _, quota := range orgQuota.MarketplaceQuotas {
				for _, cloudAccount := range *quota.CloudAccounts {
					if cloudAccount.GetCloudProviderId() == provider || cloudAccount.GetCloudProviderId() == RedHatMarketPlace {
						filteredMarketPlaceQuotas = append(filteredMarketPlaceQuotas, quota)
						break
					}
				}
			}

			orgQuota.MarketplaceQuotas = uniqueQuotaSpec(filteredMarketPlaceQuotas)
		}

		if len(orgQuota.MarketplaceQuotas) == 0 {
			return nil, f.Localizer.MustLocalizeError("kafka.create.provider.error.noMarketplaceQuota")
		}

		marketplaceQuota, err := getMarketplaceQuota(f, orgQuota.MarketplaceQuotas, marketplaceInfo)
		if err != nil {
			return nil, err
		}

		marketplaceQuota.CloudAccounts, err = pickCloudAccount(f, marketplaceQuota.CloudAccounts, marketplaceInfo, provider)
		if err != nil {
			return nil, err
		}

		return marketplaceQuota, nil
	}

	if len(orgQuota.StandardQuotas) == 0 && len(orgQuota.MarketplaceQuotas) == 0 && len(orgQuota.EvalQuotas) == 0 && len(orgQuota.EnterpriseQuotas) > 0 {
		return nil, f.Localizer.MustLocalizeError("kafka.create.provider.error.onlyEnterpriseQuota")
	}

	quotaTypeCount := 0

	if len(orgQuota.StandardQuotas) > 0 {
		quotaTypeCount++
	}
	if len(orgQuota.MarketplaceQuotas) > 0 {
		quotaTypeCount++
	}
	if len(orgQuota.EvalQuotas) > 0 {
		quotaTypeCount++
	}

	if quotaTypeCount > 1 {

		switch marketplaceInfo.BillingModel {
		case QuotaStandardType:
			return &orgQuota.StandardQuotas[0], nil
		case QuotaEvalType:
			return &orgQuota.EvalQuotas[0], nil
		}

		if marketplaceInfo.BillingModel == QuotaMarketplaceType || marketplaceInfo.Provider != "" || marketplaceInfo.CloudAccountID != "" {

			var filteredMarketPlaceQuotas []QuotaSpec

			if provider != "" {
				for _, quota := range orgQuota.MarketplaceQuotas {
					for _, cloudAccount := range *quota.CloudAccounts {
						if cloudAccount.GetCloudProviderId() == provider || cloudAccount.GetCloudProviderId() == RedHatMarketPlace {
							filteredMarketPlaceQuotas = append(filteredMarketPlaceQuotas, quota)
							break
						}
					}
				}

				orgQuota.MarketplaceQuotas = uniqueQuotaSpec(filteredMarketPlaceQuotas)
			}

			if len(orgQuota.MarketplaceQuotas) == 0 {
				return nil, f.Localizer.MustLocalizeError("kafka.create.provider.error.noMarketplaceQuota")
			}

			marketplaceQuota, err := getMarketplaceQuota(f, orgQuota.MarketplaceQuotas, marketplaceInfo)
			if err != nil {
				return nil, err
			}

			marketplaceQuota.CloudAccounts, err = pickCloudAccount(f, marketplaceQuota.CloudAccounts, marketplaceInfo, provider)
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

func pickCloudAccount(f *factory.Factory, cloudAccounts *[]amsclient.CloudAccount, market MarketplaceInfo, provider string) (*[]amsclient.CloudAccount, error) {

	// filter cloud accounts according to provider
	var filteredCloudAccounts []amsclient.CloudAccount

	if provider != "" {
		for _, cloudAccount := range *cloudAccounts {
			if *cloudAccount.CloudProviderId == provider || *cloudAccount.CloudProviderId == RedHatMarketPlace {
				filteredCloudAccounts = append(filteredCloudAccounts, cloudAccount)
			}
		}

		*cloudAccounts = filteredCloudAccounts
	}

	if len(*cloudAccounts) == 1 {
		return cloudAccounts, nil
	}

	if len(*cloudAccounts) > 1 && market.Provider == "" && market.CloudAccountID == "" {
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

// uniqueQuotaSpec accepts a list of QuotaSpec objects and returns the unique QuotaSpecs
func uniqueQuotaSpec(s []QuotaSpec) []QuotaSpec {
	inResult := make(map[QuotaSpec]bool)
	var result []QuotaSpec
	for _, quota := range s {
		if _, ok := inResult[quota]; !ok {
			inResult[quota] = true
			result = append(result, quota)
		}
	}
	return result
}
