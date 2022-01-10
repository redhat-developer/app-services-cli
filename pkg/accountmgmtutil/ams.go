package accountmgmtutil

import (
	"context"
	"errors"

	"github.com/redhat-developer/app-services-cli/pkg/core/connection"

	amsclient "github.com/redhat-developer/app-services-sdk-go/accountmgmt/apiv1/client"
)

func CheckTermsAccepted(ctx context.Context, spec ServiceTermsSpec, conn connection.Connection) (accepted bool, redirectURI string, err error) {
	termsReview, _, err := conn.API().AccountMgmt().
		ApiAuthorizationsV1SelfTermsReviewPost(ctx).
		SelfTermsReview(amsclient.SelfTermsReview{
			EventCode: &spec.EventCode,
			SiteCode:  &spec.SiteCode,
		}).
		Execute()
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
