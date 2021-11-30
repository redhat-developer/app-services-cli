package rbacutil

import (
	"context"

	"github.com/redhat-developer/app-services-cli/pkg/apis/rbac"
)

// FetchAllUsers retrieves and returns every user within the current user's organization with the applied filters
func FetchAllUsers(ctx context.Context, rbacAPI func() rbac.PrincipalAPI, queryParams ...rbac.QueryParam) ([]rbac.Principal, error) {
	principals, resp, err := rbacAPI().GetPrincipals(ctx, rbac.WithIntQueryParam("limit", 0))
	if resp != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		return nil, err
	}

	totalUsers := principals.Meta.Count

	// no need to re-fetch if the first call contains all items
	if totalUsers <= len(principals.Data) {
		return principals.Data, nil
	}

	queryParams = append(queryParams, rbac.WithIntQueryParam("limit", totalUsers))
	principals, resp, err = rbacAPI().GetPrincipals(ctx, queryParams...)

	if resp != nil {
		defer resp.Body.Close()
	}

	if err != nil {
		return nil, err
	}
	return principals.Data, nil
}
