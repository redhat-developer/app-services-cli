package flagutil

import (
	"context"

	"github.com/redhat-developer/app-services-cli/pkg/api/rbac"
	"github.com/redhat-developer/app-services-cli/pkg/api/rbac/rbacutil"
	"github.com/redhat-developer/app-services-cli/pkg/shared/factory"
	"github.com/spf13/cobra"
)

var (
	cachedServiceAccounts []string
)

// EnableStaticFlagCompletion enables autocompletion for flags with predefined valid values
func EnableStaticFlagCompletion(cmd *cobra.Command, flagName string, validValues []string) {
	_ = cmd.RegisterFlagCompletionFunc(flagName, func(cmd *cobra.Command, _ []string, _ string) ([]string, cobra.ShellCompDirective) {
		return validValues, cobra.ShellCompDirectiveNoSpace
	})
}

// EnableOutputFlagCompletion enables autocompletion for output flag
func EnableOutputFlagCompletion(cmd *cobra.Command) {
	_ = cmd.RegisterFlagCompletionFunc("output", func(cmd *cobra.Command, _ []string, _ string) ([]string, cobra.ShellCompDirective) {
		return ValidOutputFormats, cobra.ShellCompDirectiveNoSpace
	})
}

// RegisterUserCompletionFunc adds the user list to flag dynamic completion
func RegisterUserCompletionFunc(cmd *cobra.Command, flagName string, f *factory.Factory) error {
	return cmd.RegisterFlagCompletionFunc(flagName, func(cmd *cobra.Command, _ []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		var usernames []string
		directive := cobra.ShellCompDirectiveNoSpace

		conn, err := f.Connection()
		if err != nil {
			return usernames, directive
		}

		queryParams := []rbac.QueryParam{rbac.WithQueryParam("match_criteria", "partial"), rbac.WithQueryParam("usernames", toComplete)}
		principals, err := rbacutil.FetchAllUsers(context.Background(), conn.API().RBAC().PrincipalAPI, queryParams...)
		if err != nil || len(principals) == 0 {
			return usernames, directive
		}

		for _, p := range principals {
			usernames = append(usernames, p.Username)
		}

		return usernames, directive
	})
}

// RegisterServiceAccountCompletionFunc adds the service account list to flag dynamic completion
func RegisterServiceAccountCompletionFunc(cmd *cobra.Command, f *factory.Factory) error {
	return cmd.RegisterFlagCompletionFunc("service-account", func(cmd *cobra.Command, _ []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		var emptyList []string
		directive := cobra.ShellCompDirectiveNoSpace

		conn, err := f.Connection()
		if err != nil {
			return emptyList, directive
		}

		// There is no server-side pagination for service accounts, so we will always
		// have the full list after the first time
		if len(cachedServiceAccounts) > 0 {
			return cachedServiceAccounts, directive
		}

		serviceAccountResults, _, err := conn.API().ServiceAccountMgmt().GetServiceAccounts(cmd.Context()).Execute()
		if err != nil || len(serviceAccountResults.GetItems()) == 0 {
			return emptyList, directive
		}

		serviceAccounts := serviceAccountResults.GetItems()

		for _, serviceAcct := range serviceAccounts {
			cachedServiceAccounts = append(cachedServiceAccounts, serviceAcct.GetClientId())
		}

		return cachedServiceAccounts, directive
	})
}
