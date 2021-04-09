# \DefaultApi

All URIs are relative to *https://api.openshift.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiAccountsMgmtV1AccessTokenPost**](DefaultApi.md#ApiAccountsMgmtV1AccessTokenPost) | **Post** /api/accounts_mgmt/v1/access_token | Return access token generated from registries in docker format
[**ApiAccountsMgmtV1AccountsGet**](DefaultApi.md#ApiAccountsMgmtV1AccountsGet) | **Get** /api/accounts_mgmt/v1/accounts | Returns a list of accounts
[**ApiAccountsMgmtV1AccountsIdGet**](DefaultApi.md#ApiAccountsMgmtV1AccountsIdGet) | **Get** /api/accounts_mgmt/v1/accounts/{id} | Get an account by id
[**ApiAccountsMgmtV1AccountsIdLabelsGet**](DefaultApi.md#ApiAccountsMgmtV1AccountsIdLabelsGet) | **Get** /api/accounts_mgmt/v1/accounts/{id}/labels | Returns a list of labels
[**ApiAccountsMgmtV1AccountsIdLabelsKeyDelete**](DefaultApi.md#ApiAccountsMgmtV1AccountsIdLabelsKeyDelete) | **Delete** /api/accounts_mgmt/v1/accounts/{id}/labels/{key} | Delete a label
[**ApiAccountsMgmtV1AccountsIdLabelsKeyGet**](DefaultApi.md#ApiAccountsMgmtV1AccountsIdLabelsKeyGet) | **Get** /api/accounts_mgmt/v1/accounts/{id}/labels/{key} | Get subscription labels by label key
[**ApiAccountsMgmtV1AccountsIdLabelsKeyPatch**](DefaultApi.md#ApiAccountsMgmtV1AccountsIdLabelsKeyPatch) | **Patch** /api/accounts_mgmt/v1/accounts/{id}/labels/{key} | Create a new label or update an existing label
[**ApiAccountsMgmtV1AccountsIdLabelsPost**](DefaultApi.md#ApiAccountsMgmtV1AccountsIdLabelsPost) | **Post** /api/accounts_mgmt/v1/accounts/{id}/labels | Create a new label or update an existing label
[**ApiAccountsMgmtV1AccountsIdPatch**](DefaultApi.md#ApiAccountsMgmtV1AccountsIdPatch) | **Patch** /api/accounts_mgmt/v1/accounts/{id} | Update an account
[**ApiAccountsMgmtV1AccountsPost**](DefaultApi.md#ApiAccountsMgmtV1AccountsPost) | **Post** /api/accounts_mgmt/v1/accounts | Create a new account
[**ApiAccountsMgmtV1ClusterAuthorizationsPost**](DefaultApi.md#ApiAccountsMgmtV1ClusterAuthorizationsPost) | **Post** /api/accounts_mgmt/v1/cluster_authorizations | Authorizes new cluster creation against an exsiting RH Subscription.
[**ApiAccountsMgmtV1ClusterRegistrationsPost**](DefaultApi.md#ApiAccountsMgmtV1ClusterRegistrationsPost) | **Post** /api/accounts_mgmt/v1/cluster_registrations | Finds or creates a cluster registration with a registy credential token and cluster ID
[**ApiAccountsMgmtV1CurrentAccountGet**](DefaultApi.md#ApiAccountsMgmtV1CurrentAccountGet) | **Get** /api/accounts_mgmt/v1/current_account | Get the authenticated account
[**ApiAccountsMgmtV1ErrorsGet**](DefaultApi.md#ApiAccountsMgmtV1ErrorsGet) | **Get** /api/accounts_mgmt/v1/errors | Returns a list of errors
[**ApiAccountsMgmtV1ErrorsIdGet**](DefaultApi.md#ApiAccountsMgmtV1ErrorsIdGet) | **Get** /api/accounts_mgmt/v1/errors/{id} | Get an error by id
[**ApiAccountsMgmtV1FeatureTogglesIdQueryPost**](DefaultApi.md#ApiAccountsMgmtV1FeatureTogglesIdQueryPost) | **Post** /api/accounts_mgmt/v1/feature_toggles/{id}/query | Query a feature toggle by id
[**ApiAccountsMgmtV1LabelsGet**](DefaultApi.md#ApiAccountsMgmtV1LabelsGet) | **Get** /api/accounts_mgmt/v1/labels | Returns a list of labels
[**ApiAccountsMgmtV1MetricsGet**](DefaultApi.md#ApiAccountsMgmtV1MetricsGet) | **Get** /api/accounts_mgmt/v1/metrics | Returns a list of metrics
[**ApiAccountsMgmtV1NotifyPost**](DefaultApi.md#ApiAccountsMgmtV1NotifyPost) | **Post** /api/accounts_mgmt/v1/notify | Notify the owner of cluster/subscription
[**ApiAccountsMgmtV1OrganizationsGet**](DefaultApi.md#ApiAccountsMgmtV1OrganizationsGet) | **Get** /api/accounts_mgmt/v1/organizations | Returns a list of organizations
[**ApiAccountsMgmtV1OrganizationsIdGet**](DefaultApi.md#ApiAccountsMgmtV1OrganizationsIdGet) | **Get** /api/accounts_mgmt/v1/organizations/{id} | Get an organization by id
[**ApiAccountsMgmtV1OrganizationsIdLabelsGet**](DefaultApi.md#ApiAccountsMgmtV1OrganizationsIdLabelsGet) | **Get** /api/accounts_mgmt/v1/organizations/{id}/labels | Returns a list of labels
[**ApiAccountsMgmtV1OrganizationsIdLabelsKeyDelete**](DefaultApi.md#ApiAccountsMgmtV1OrganizationsIdLabelsKeyDelete) | **Delete** /api/accounts_mgmt/v1/organizations/{id}/labels/{key} | Delete a label
[**ApiAccountsMgmtV1OrganizationsIdLabelsKeyGet**](DefaultApi.md#ApiAccountsMgmtV1OrganizationsIdLabelsKeyGet) | **Get** /api/accounts_mgmt/v1/organizations/{id}/labels/{key} | Get subscription labels by label key
[**ApiAccountsMgmtV1OrganizationsIdLabelsKeyPatch**](DefaultApi.md#ApiAccountsMgmtV1OrganizationsIdLabelsKeyPatch) | **Patch** /api/accounts_mgmt/v1/organizations/{id}/labels/{key} | Create a new label or update an existing label
[**ApiAccountsMgmtV1OrganizationsIdLabelsPost**](DefaultApi.md#ApiAccountsMgmtV1OrganizationsIdLabelsPost) | **Post** /api/accounts_mgmt/v1/organizations/{id}/labels | Create a new label or update an existing label
[**ApiAccountsMgmtV1OrganizationsIdPatch**](DefaultApi.md#ApiAccountsMgmtV1OrganizationsIdPatch) | **Patch** /api/accounts_mgmt/v1/organizations/{id} | Update an organization
[**ApiAccountsMgmtV1OrganizationsIdSummaryDashboardGet**](DefaultApi.md#ApiAccountsMgmtV1OrganizationsIdSummaryDashboardGet) | **Get** /api/accounts_mgmt/v1/organizations/{id}/summary_dashboard | Returns a summary of organizations clusters based on metrics
[**ApiAccountsMgmtV1OrganizationsOrgIdQuotaCostGet**](DefaultApi.md#ApiAccountsMgmtV1OrganizationsOrgIdQuotaCostGet) | **Get** /api/accounts_mgmt/v1/organizations/{orgId}/quota_cost | Returns a summary of quota cost
[**ApiAccountsMgmtV1OrganizationsOrgIdQuotaSummaryGet**](DefaultApi.md#ApiAccountsMgmtV1OrganizationsOrgIdQuotaSummaryGet) | **Get** /api/accounts_mgmt/v1/organizations/{orgId}/quota_summary | Returns a summary of resource quota
[**ApiAccountsMgmtV1OrganizationsOrgIdResourceQuotaGet**](DefaultApi.md#ApiAccountsMgmtV1OrganizationsOrgIdResourceQuotaGet) | **Get** /api/accounts_mgmt/v1/organizations/{orgId}/resource_quota | Returns a list of resource quota objects
[**ApiAccountsMgmtV1OrganizationsOrgIdResourceQuotaPost**](DefaultApi.md#ApiAccountsMgmtV1OrganizationsOrgIdResourceQuotaPost) | **Post** /api/accounts_mgmt/v1/organizations/{orgId}/resource_quota | Create a new resource quota
[**ApiAccountsMgmtV1OrganizationsOrgIdResourceQuotaQuotaIdDelete**](DefaultApi.md#ApiAccountsMgmtV1OrganizationsOrgIdResourceQuotaQuotaIdDelete) | **Delete** /api/accounts_mgmt/v1/organizations/{orgId}/resource_quota/{quotaId} | Delete a resource quota
[**ApiAccountsMgmtV1OrganizationsOrgIdResourceQuotaQuotaIdGet**](DefaultApi.md#ApiAccountsMgmtV1OrganizationsOrgIdResourceQuotaQuotaIdGet) | **Get** /api/accounts_mgmt/v1/organizations/{orgId}/resource_quota/{quotaId} | Get a resource quota by id
[**ApiAccountsMgmtV1OrganizationsOrgIdResourceQuotaQuotaIdPatch**](DefaultApi.md#ApiAccountsMgmtV1OrganizationsOrgIdResourceQuotaQuotaIdPatch) | **Patch** /api/accounts_mgmt/v1/organizations/{orgId}/resource_quota/{quotaId} | Update a resource quota
[**ApiAccountsMgmtV1OrganizationsPost**](DefaultApi.md#ApiAccountsMgmtV1OrganizationsPost) | **Post** /api/accounts_mgmt/v1/organizations | Create a new organization
[**ApiAccountsMgmtV1PlansGet**](DefaultApi.md#ApiAccountsMgmtV1PlansGet) | **Get** /api/accounts_mgmt/v1/plans | Get all plans
[**ApiAccountsMgmtV1PlansIdGet**](DefaultApi.md#ApiAccountsMgmtV1PlansIdGet) | **Get** /api/accounts_mgmt/v1/plans/{id} | Get a plan by id
[**ApiAccountsMgmtV1PullSecretsExternalResourceIdDelete**](DefaultApi.md#ApiAccountsMgmtV1PullSecretsExternalResourceIdDelete) | **Delete** /api/accounts_mgmt/v1/pull_secrets/{externalResourceId} | Delete a pull secret
[**ApiAccountsMgmtV1PullSecretsPost**](DefaultApi.md#ApiAccountsMgmtV1PullSecretsPost) | **Post** /api/accounts_mgmt/v1/pull_secrets | Return access token generated from registries in docker format
[**ApiAccountsMgmtV1RegistriesGet**](DefaultApi.md#ApiAccountsMgmtV1RegistriesGet) | **Get** /api/accounts_mgmt/v1/registries | Returns a list of registries
[**ApiAccountsMgmtV1RegistriesIdGet**](DefaultApi.md#ApiAccountsMgmtV1RegistriesIdGet) | **Get** /api/accounts_mgmt/v1/registries/{id} | Get an registry by id
[**ApiAccountsMgmtV1RegistryCredentialsGet**](DefaultApi.md#ApiAccountsMgmtV1RegistryCredentialsGet) | **Get** /api/accounts_mgmt/v1/registry_credentials | 
[**ApiAccountsMgmtV1RegistryCredentialsIdDelete**](DefaultApi.md#ApiAccountsMgmtV1RegistryCredentialsIdDelete) | **Delete** /api/accounts_mgmt/v1/registry_credentials/{id} | Delete a registry credential by id
[**ApiAccountsMgmtV1RegistryCredentialsIdGet**](DefaultApi.md#ApiAccountsMgmtV1RegistryCredentialsIdGet) | **Get** /api/accounts_mgmt/v1/registry_credentials/{id} | Get a registry credentials by id
[**ApiAccountsMgmtV1RegistryCredentialsIdPatch**](DefaultApi.md#ApiAccountsMgmtV1RegistryCredentialsIdPatch) | **Patch** /api/accounts_mgmt/v1/registry_credentials/{id} | Update a registry credential
[**ApiAccountsMgmtV1RegistryCredentialsPost**](DefaultApi.md#ApiAccountsMgmtV1RegistryCredentialsPost) | **Post** /api/accounts_mgmt/v1/registry_credentials | Request the creation of a registry credential
[**ApiAccountsMgmtV1ReservedResourcesGet**](DefaultApi.md#ApiAccountsMgmtV1ReservedResourcesGet) | **Get** /api/accounts_mgmt/v1/reserved_resources | Returns a list of reserved resources
[**ApiAccountsMgmtV1ResourceQuotaGet**](DefaultApi.md#ApiAccountsMgmtV1ResourceQuotaGet) | **Get** /api/accounts_mgmt/v1/resource_quota | Returns a list of resource quota objects
[**ApiAccountsMgmtV1RoleBindingsGet**](DefaultApi.md#ApiAccountsMgmtV1RoleBindingsGet) | **Get** /api/accounts_mgmt/v1/role_bindings | Returns a list of role bindings
[**ApiAccountsMgmtV1RoleBindingsIdDelete**](DefaultApi.md#ApiAccountsMgmtV1RoleBindingsIdDelete) | **Delete** /api/accounts_mgmt/v1/role_bindings/{id} | Delete a role binding
[**ApiAccountsMgmtV1RoleBindingsIdGet**](DefaultApi.md#ApiAccountsMgmtV1RoleBindingsIdGet) | **Get** /api/accounts_mgmt/v1/role_bindings/{id} | Get a role binding
[**ApiAccountsMgmtV1RoleBindingsIdPatch**](DefaultApi.md#ApiAccountsMgmtV1RoleBindingsIdPatch) | **Patch** /api/accounts_mgmt/v1/role_bindings/{id} | Update a role binding
[**ApiAccountsMgmtV1RoleBindingsPost**](DefaultApi.md#ApiAccountsMgmtV1RoleBindingsPost) | **Post** /api/accounts_mgmt/v1/role_bindings | Create a new role binding
[**ApiAccountsMgmtV1RolesGet**](DefaultApi.md#ApiAccountsMgmtV1RolesGet) | **Get** /api/accounts_mgmt/v1/roles | Returns a list of roles
[**ApiAccountsMgmtV1RolesIdGet**](DefaultApi.md#ApiAccountsMgmtV1RolesIdGet) | **Get** /api/accounts_mgmt/v1/roles/{id} | Get a role by id
[**ApiAccountsMgmtV1SkuRulesGet**](DefaultApi.md#ApiAccountsMgmtV1SkuRulesGet) | **Get** /api/accounts_mgmt/v1/sku_rules | Returns a list of UHC product SKU Rules
[**ApiAccountsMgmtV1SkuRulesIdGet**](DefaultApi.md#ApiAccountsMgmtV1SkuRulesIdGet) | **Get** /api/accounts_mgmt/v1/sku_rules/{id} | Get a sku rules by id
[**ApiAccountsMgmtV1SkusGet**](DefaultApi.md#ApiAccountsMgmtV1SkusGet) | **Get** /api/accounts_mgmt/v1/skus | Returns a list of UHC product SKUs
[**ApiAccountsMgmtV1SkusIdGet**](DefaultApi.md#ApiAccountsMgmtV1SkusIdGet) | **Get** /api/accounts_mgmt/v1/skus/{id} | Get a sku by id
[**ApiAccountsMgmtV1SubscriptionsGet**](DefaultApi.md#ApiAccountsMgmtV1SubscriptionsGet) | **Get** /api/accounts_mgmt/v1/subscriptions | Returns a list of subscriptions
[**ApiAccountsMgmtV1SubscriptionsIdDelete**](DefaultApi.md#ApiAccountsMgmtV1SubscriptionsIdDelete) | **Delete** /api/accounts_mgmt/v1/subscriptions/{id} | Deletes a subscription by id
[**ApiAccountsMgmtV1SubscriptionsIdGet**](DefaultApi.md#ApiAccountsMgmtV1SubscriptionsIdGet) | **Get** /api/accounts_mgmt/v1/subscriptions/{id} | Get a subscription by id
[**ApiAccountsMgmtV1SubscriptionsIdLabelsGet**](DefaultApi.md#ApiAccountsMgmtV1SubscriptionsIdLabelsGet) | **Get** /api/accounts_mgmt/v1/subscriptions/{id}/labels | Returns a list of labels
[**ApiAccountsMgmtV1SubscriptionsIdLabelsKeyDelete**](DefaultApi.md#ApiAccountsMgmtV1SubscriptionsIdLabelsKeyDelete) | **Delete** /api/accounts_mgmt/v1/subscriptions/{id}/labels/{key} | Delete a label
[**ApiAccountsMgmtV1SubscriptionsIdLabelsKeyGet**](DefaultApi.md#ApiAccountsMgmtV1SubscriptionsIdLabelsKeyGet) | **Get** /api/accounts_mgmt/v1/subscriptions/{id}/labels/{key} | Get subscription labels by label key
[**ApiAccountsMgmtV1SubscriptionsIdLabelsKeyPatch**](DefaultApi.md#ApiAccountsMgmtV1SubscriptionsIdLabelsKeyPatch) | **Patch** /api/accounts_mgmt/v1/subscriptions/{id}/labels/{key} | Create a new label or update an existing label
[**ApiAccountsMgmtV1SubscriptionsIdLabelsPost**](DefaultApi.md#ApiAccountsMgmtV1SubscriptionsIdLabelsPost) | **Post** /api/accounts_mgmt/v1/subscriptions/{id}/labels | Create a new label or update an existing label
[**ApiAccountsMgmtV1SubscriptionsIdNotifyPost**](DefaultApi.md#ApiAccountsMgmtV1SubscriptionsIdNotifyPost) | **Post** /api/accounts_mgmt/v1/subscriptions/{id}/notify | Notify the owner of a subscription
[**ApiAccountsMgmtV1SubscriptionsIdPatch**](DefaultApi.md#ApiAccountsMgmtV1SubscriptionsIdPatch) | **Patch** /api/accounts_mgmt/v1/subscriptions/{id} | Update a subscription
[**ApiAccountsMgmtV1SubscriptionsIdReservedResourcesGet**](DefaultApi.md#ApiAccountsMgmtV1SubscriptionsIdReservedResourcesGet) | **Get** /api/accounts_mgmt/v1/subscriptions/{id}/reserved_resources | Returns a list of reserved resources
[**ApiAccountsMgmtV1SubscriptionsIdSupportCasesGet**](DefaultApi.md#ApiAccountsMgmtV1SubscriptionsIdSupportCasesGet) | **Get** /api/accounts_mgmt/v1/subscriptions/{id}/support_cases | Returns a list of open support creates opened against the external cluster id of this subscrption
[**ApiAccountsMgmtV1SubscriptionsPost**](DefaultApi.md#ApiAccountsMgmtV1SubscriptionsPost) | **Post** /api/accounts_mgmt/v1/subscriptions | Create a new subscription
[**ApiAccountsMgmtV1SubscriptionsSubIdNotificationContactsAccountIdDelete**](DefaultApi.md#ApiAccountsMgmtV1SubscriptionsSubIdNotificationContactsAccountIdDelete) | **Delete** /api/accounts_mgmt/v1/subscriptions/{subId}/notification_contacts/{accountId} | Deletes a notification contact by subscription and account id
[**ApiAccountsMgmtV1SubscriptionsSubIdNotificationContactsGet**](DefaultApi.md#ApiAccountsMgmtV1SubscriptionsSubIdNotificationContactsGet) | **Get** /api/accounts_mgmt/v1/subscriptions/{subId}/notification_contacts | Returns a list of notification contacts for the given subscription
[**ApiAccountsMgmtV1SubscriptionsSubIdNotificationContactsPost**](DefaultApi.md#ApiAccountsMgmtV1SubscriptionsSubIdNotificationContactsPost) | **Post** /api/accounts_mgmt/v1/subscriptions/{subId}/notification_contacts | Add an account as a notification contact to this subscription
[**ApiAccountsMgmtV1SubscriptionsSubIdReservedResourcesReservedResourceIdGet**](DefaultApi.md#ApiAccountsMgmtV1SubscriptionsSubIdReservedResourcesReservedResourceIdGet) | **Get** /api/accounts_mgmt/v1/subscriptions/{subId}/reserved_resources/{reservedResourceId} | Get reserved resources by id
[**ApiAccountsMgmtV1SupportCasesCaseIdDelete**](DefaultApi.md#ApiAccountsMgmtV1SupportCasesCaseIdDelete) | **Delete** /api/accounts_mgmt/v1/support_cases/{caseId} | Delete a support case
[**ApiAccountsMgmtV1SupportCasesPost**](DefaultApi.md#ApiAccountsMgmtV1SupportCasesPost) | **Post** /api/accounts_mgmt/v1/support_cases | create a support case for the subscription
[**ApiAccountsMgmtV1TokenAuthorizationPost**](DefaultApi.md#ApiAccountsMgmtV1TokenAuthorizationPost) | **Post** /api/accounts_mgmt/v1/token_authorization | Finds the account owner of the provided token
[**ApiAuthorizationsV1AccessReviewPost**](DefaultApi.md#ApiAuthorizationsV1AccessReviewPost) | **Post** /api/authorizations/v1/access_review | Review an account&#39;s access to perform an action on a particular resource or resource type
[**ApiAuthorizationsV1CapabilityReviewPost**](DefaultApi.md#ApiAuthorizationsV1CapabilityReviewPost) | **Post** /api/authorizations/v1/capability_review | Review an account&#39;s capabilities
[**ApiAuthorizationsV1ExportControlReviewPost**](DefaultApi.md#ApiAuthorizationsV1ExportControlReviewPost) | **Post** /api/authorizations/v1/export_control_review | Determine whether a user is restricted from downloading Red Hat software based on export control compliance. 
[**ApiAuthorizationsV1ResourceReviewPost**](DefaultApi.md#ApiAuthorizationsV1ResourceReviewPost) | **Post** /api/authorizations/v1/resource_review | Obtain resource ids for resources an account may perform the specified action upon. Resource ids returned as [\&quot;*\&quot;] is shorthand for all ids.
[**ApiAuthorizationsV1SelfAccessReviewPost**](DefaultApi.md#ApiAuthorizationsV1SelfAccessReviewPost) | **Post** /api/authorizations/v1/self_access_review | Review your ability to perform an action on a particular resource or resource type
[**ApiAuthorizationsV1SelfResourceReviewPost**](DefaultApi.md#ApiAuthorizationsV1SelfResourceReviewPost) | **Post** /api/authorizations/v1/self_resource_review | Obtain resource ids for resources you may perform the specified action upon. Resource ids returned as [\&quot;*\&quot;] is shorthand for all ids.
[**ApiAuthorizationsV1SelfTermsReviewPost**](DefaultApi.md#ApiAuthorizationsV1SelfTermsReviewPost) | **Post** /api/authorizations/v1/self_terms_review | Review your status of Terms
[**ApiAuthorizationsV1TermsReviewPost**](DefaultApi.md#ApiAuthorizationsV1TermsReviewPost) | **Post** /api/authorizations/v1/terms_review | Review an account&#39;s status of Terms



## ApiAccountsMgmtV1AccessTokenPost

> AccessTokenCfg ApiAccountsMgmtV1AccessTokenPost(ctx).Execute()

Return access token generated from registries in docker format

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ApiAccountsMgmtV1AccessTokenPost(context.Background()).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ApiAccountsMgmtV1AccessTokenPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiAccountsMgmtV1AccessTokenPost`: AccessTokenCfg
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ApiAccountsMgmtV1AccessTokenPost`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiApiAccountsMgmtV1AccessTokenPostRequest struct via the builder pattern


### Return type

[**AccessTokenCfg**](AccessTokenCfg.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiAccountsMgmtV1AccountsGet

> AccountList ApiAccountsMgmtV1AccountsGet(ctx).Page(page).Size(size).Search(search).OrderBy(orderBy).Fields(fields).FetchLabels(fetchLabels).FetchCapabilities(fetchCapabilities).Execute()

Returns a list of accounts

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    page := int32(56) // int32 | Page number of record list when record list exceeds specified page size (optional) (default to 1)
    size := int32(56) // int32 | Maximum number of records to return (optional) (default to 100)
    search := "search_example" // string | Specifies the search criteria. The syntax of this parameter is similar to the syntax of the _where_ clause of an SQL statement, using the names of the json attributes / column names of the account. For example, in order to retrieve all the accounts with a username starting with `my`:  ```sql username like 'my%' ```  The search criteria can also be applied on related resource. For example, in order to retrieve all the subscriptions labeled by `foo=bar`,  ```sql labels.key = 'foo' and labels.value = 'bar' ```  If the parameter isn't provided, or if the value is empty, then all the accounts that the user has permission to see will be returned. (optional)
    orderBy := "orderBy_example" // string | Specifies the order by criteria. The syntax of this parameter is similar to the syntax of the _order by_ clause of an SQL statement, but using the names of the json attributes / column of the account. For example, in order to retrieve all accounts ordered by username:  ```sql username asc ```  Or in order to retrieve all accounts ordered by username _and_ first name:  ```sql username asc, firstName asc ```  If the parameter isn't provided, or if the value is empty, then no explicit ordering will be applied. (optional)
    fields := "fields_example" // string | Supplies a comma-separated list of fields to be returned. Fields of sub-structures and of arrays use <structure>.<field> notation. <stucture>.* means all field of a structure Example: For each Subscription to get id, href, plan(id and kind) and labels (all fields)  ``` ocm get subscriptions --parameter fields=id,href,plan.id,plan.kind,labels.* --parameter fetchLabels=true ``` (optional)
    fetchLabels := true // bool | If true, includes the labels on a subscription/organization/account in the output. Could slow request response time. (optional)
    fetchCapabilities := true // bool | If true, includes the capabilities on a subscription in the output. Could slow request response time. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ApiAccountsMgmtV1AccountsGet(context.Background()).Page(page).Size(size).Search(search).OrderBy(orderBy).Fields(fields).FetchLabels(fetchLabels).FetchCapabilities(fetchCapabilities).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ApiAccountsMgmtV1AccountsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiAccountsMgmtV1AccountsGet`: AccountList
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ApiAccountsMgmtV1AccountsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiAccountsMgmtV1AccountsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Page number of record list when record list exceeds specified page size | [default to 1]
 **size** | **int32** | Maximum number of records to return | [default to 100]
 **search** | **string** | Specifies the search criteria. The syntax of this parameter is similar to the syntax of the _where_ clause of an SQL statement, using the names of the json attributes / column names of the account. For example, in order to retrieve all the accounts with a username starting with &#x60;my&#x60;:  &#x60;&#x60;&#x60;sql username like &#39;my%&#39; &#x60;&#x60;&#x60;  The search criteria can also be applied on related resource. For example, in order to retrieve all the subscriptions labeled by &#x60;foo&#x3D;bar&#x60;,  &#x60;&#x60;&#x60;sql labels.key &#x3D; &#39;foo&#39; and labels.value &#x3D; &#39;bar&#39; &#x60;&#x60;&#x60;  If the parameter isn&#39;t provided, or if the value is empty, then all the accounts that the user has permission to see will be returned. | 
 **orderBy** | **string** | Specifies the order by criteria. The syntax of this parameter is similar to the syntax of the _order by_ clause of an SQL statement, but using the names of the json attributes / column of the account. For example, in order to retrieve all accounts ordered by username:  &#x60;&#x60;&#x60;sql username asc &#x60;&#x60;&#x60;  Or in order to retrieve all accounts ordered by username _and_ first name:  &#x60;&#x60;&#x60;sql username asc, firstName asc &#x60;&#x60;&#x60;  If the parameter isn&#39;t provided, or if the value is empty, then no explicit ordering will be applied. | 
 **fields** | **string** | Supplies a comma-separated list of fields to be returned. Fields of sub-structures and of arrays use &lt;structure&gt;.&lt;field&gt; notation. &lt;stucture&gt;.* means all field of a structure Example: For each Subscription to get id, href, plan(id and kind) and labels (all fields)  &#x60;&#x60;&#x60; ocm get subscriptions --parameter fields&#x3D;id,href,plan.id,plan.kind,labels.* --parameter fetchLabels&#x3D;true &#x60;&#x60;&#x60; | 
 **fetchLabels** | **bool** | If true, includes the labels on a subscription/organization/account in the output. Could slow request response time. | 
 **fetchCapabilities** | **bool** | If true, includes the capabilities on a subscription in the output. Could slow request response time. | 

### Return type

[**AccountList**](AccountList.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiAccountsMgmtV1AccountsIdGet

> Account ApiAccountsMgmtV1AccountsIdGet(ctx, id).FetchLabels(fetchLabels).FetchCapabilities(fetchCapabilities).Execute()

Get an account by id

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "id_example" // string | The id of record
    fetchLabels := true // bool | If true, includes the labels on a subscription/organization/account in the output. Could slow request response time. (optional)
    fetchCapabilities := true // bool | If true, includes the capabilities on a subscription in the output. Could slow request response time. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ApiAccountsMgmtV1AccountsIdGet(context.Background(), id).FetchLabels(fetchLabels).FetchCapabilities(fetchCapabilities).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ApiAccountsMgmtV1AccountsIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiAccountsMgmtV1AccountsIdGet`: Account
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ApiAccountsMgmtV1AccountsIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of record | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiAccountsMgmtV1AccountsIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fetchLabels** | **bool** | If true, includes the labels on a subscription/organization/account in the output. Could slow request response time. | 
 **fetchCapabilities** | **bool** | If true, includes the capabilities on a subscription in the output. Could slow request response time. | 

### Return type

[**Account**](Account.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiAccountsMgmtV1AccountsIdLabelsGet

> LabelList ApiAccountsMgmtV1AccountsIdLabelsGet(ctx, id).Page(page).Size(size).Search(search).OrderBy(orderBy).Execute()

Returns a list of labels

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "id_example" // string | The id of record
    page := int32(56) // int32 | Page number of record list when record list exceeds specified page size (optional) (default to 1)
    size := int32(56) // int32 | Maximum number of records to return (optional) (default to 100)
    search := "search_example" // string | Specifies the search criteria. The syntax of this parameter is similar to the syntax of the _where_ clause of an SQL statement, using the names of the json attributes / column names of the account. For example, in order to retrieve all the accounts with a username starting with `my`:  ```sql username like 'my%' ```  The search criteria can also be applied on related resource. For example, in order to retrieve all the subscriptions labeled by `foo=bar`,  ```sql labels.key = 'foo' and labels.value = 'bar' ```  If the parameter isn't provided, or if the value is empty, then all the accounts that the user has permission to see will be returned. (optional)
    orderBy := "orderBy_example" // string | Specifies the order by criteria. The syntax of this parameter is similar to the syntax of the _order by_ clause of an SQL statement, but using the names of the json attributes / column of the account. For example, in order to retrieve all accounts ordered by username:  ```sql username asc ```  Or in order to retrieve all accounts ordered by username _and_ first name:  ```sql username asc, firstName asc ```  If the parameter isn't provided, or if the value is empty, then no explicit ordering will be applied. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ApiAccountsMgmtV1AccountsIdLabelsGet(context.Background(), id).Page(page).Size(size).Search(search).OrderBy(orderBy).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ApiAccountsMgmtV1AccountsIdLabelsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiAccountsMgmtV1AccountsIdLabelsGet`: LabelList
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ApiAccountsMgmtV1AccountsIdLabelsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of record | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiAccountsMgmtV1AccountsIdLabelsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | Page number of record list when record list exceeds specified page size | [default to 1]
 **size** | **int32** | Maximum number of records to return | [default to 100]
 **search** | **string** | Specifies the search criteria. The syntax of this parameter is similar to the syntax of the _where_ clause of an SQL statement, using the names of the json attributes / column names of the account. For example, in order to retrieve all the accounts with a username starting with &#x60;my&#x60;:  &#x60;&#x60;&#x60;sql username like &#39;my%&#39; &#x60;&#x60;&#x60;  The search criteria can also be applied on related resource. For example, in order to retrieve all the subscriptions labeled by &#x60;foo&#x3D;bar&#x60;,  &#x60;&#x60;&#x60;sql labels.key &#x3D; &#39;foo&#39; and labels.value &#x3D; &#39;bar&#39; &#x60;&#x60;&#x60;  If the parameter isn&#39;t provided, or if the value is empty, then all the accounts that the user has permission to see will be returned. | 
 **orderBy** | **string** | Specifies the order by criteria. The syntax of this parameter is similar to the syntax of the _order by_ clause of an SQL statement, but using the names of the json attributes / column of the account. For example, in order to retrieve all accounts ordered by username:  &#x60;&#x60;&#x60;sql username asc &#x60;&#x60;&#x60;  Or in order to retrieve all accounts ordered by username _and_ first name:  &#x60;&#x60;&#x60;sql username asc, firstName asc &#x60;&#x60;&#x60;  If the parameter isn&#39;t provided, or if the value is empty, then no explicit ordering will be applied. | 

### Return type

[**LabelList**](LabelList.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiAccountsMgmtV1AccountsIdLabelsKeyDelete

> ApiAccountsMgmtV1AccountsIdLabelsKeyDelete(ctx, id, key).Execute()

Delete a label

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "id_example" // string | The id of record
    key := "key_example" // string | The key of the label

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ApiAccountsMgmtV1AccountsIdLabelsKeyDelete(context.Background(), id, key).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ApiAccountsMgmtV1AccountsIdLabelsKeyDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of record | 
**key** | **string** | The key of the label | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiAccountsMgmtV1AccountsIdLabelsKeyDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiAccountsMgmtV1AccountsIdLabelsKeyGet

> Label ApiAccountsMgmtV1AccountsIdLabelsKeyGet(ctx, id, key).Execute()

Get subscription labels by label key

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "id_example" // string | The id of record
    key := "key_example" // string | The key of the label

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ApiAccountsMgmtV1AccountsIdLabelsKeyGet(context.Background(), id, key).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ApiAccountsMgmtV1AccountsIdLabelsKeyGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiAccountsMgmtV1AccountsIdLabelsKeyGet`: Label
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ApiAccountsMgmtV1AccountsIdLabelsKeyGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of record | 
**key** | **string** | The key of the label | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiAccountsMgmtV1AccountsIdLabelsKeyGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Label**](Label.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiAccountsMgmtV1AccountsIdLabelsKeyPatch

> Label ApiAccountsMgmtV1AccountsIdLabelsKeyPatch(ctx, id, key).Label(label).Execute()

Create a new label or update an existing label

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "id_example" // string | The id of record
    key := "key_example" // string | The key of the label
    label := *openapiclient.NewLabel(false, "Key_example", "Value_example") // Label | Label data

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ApiAccountsMgmtV1AccountsIdLabelsKeyPatch(context.Background(), id, key).Label(label).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ApiAccountsMgmtV1AccountsIdLabelsKeyPatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiAccountsMgmtV1AccountsIdLabelsKeyPatch`: Label
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ApiAccountsMgmtV1AccountsIdLabelsKeyPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of record | 
**key** | **string** | The key of the label | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiAccountsMgmtV1AccountsIdLabelsKeyPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **label** | [**Label**](Label.md) | Label data | 

### Return type

[**Label**](Label.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiAccountsMgmtV1AccountsIdLabelsPost

> Label ApiAccountsMgmtV1AccountsIdLabelsPost(ctx, id).Label(label).Execute()

Create a new label or update an existing label

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "id_example" // string | The id of record
    label := *openapiclient.NewLabel(false, "Key_example", "Value_example") // Label | Label data

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ApiAccountsMgmtV1AccountsIdLabelsPost(context.Background(), id).Label(label).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ApiAccountsMgmtV1AccountsIdLabelsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiAccountsMgmtV1AccountsIdLabelsPost`: Label
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ApiAccountsMgmtV1AccountsIdLabelsPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of record | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiAccountsMgmtV1AccountsIdLabelsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **label** | [**Label**](Label.md) | Label data | 

### Return type

[**Label**](Label.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiAccountsMgmtV1AccountsIdPatch

> Account ApiAccountsMgmtV1AccountsIdPatch(ctx, id).AccountPatchRequest(accountPatchRequest).Execute()

Update an account

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "id_example" // string | The id of record
    accountPatchRequest := *openapiclient.NewAccountPatchRequest() // AccountPatchRequest | Updated account data

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ApiAccountsMgmtV1AccountsIdPatch(context.Background(), id).AccountPatchRequest(accountPatchRequest).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ApiAccountsMgmtV1AccountsIdPatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiAccountsMgmtV1AccountsIdPatch`: Account
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ApiAccountsMgmtV1AccountsIdPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of record | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiAccountsMgmtV1AccountsIdPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accountPatchRequest** | [**AccountPatchRequest**](AccountPatchRequest.md) | Updated account data | 

### Return type

[**Account**](Account.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiAccountsMgmtV1AccountsPost

> Account ApiAccountsMgmtV1AccountsPost(ctx).Account(account).DryRun(dryRun).Execute()

Create a new account

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    account := *openapiclient.NewAccount("Username_example") // Account | Account data
    dryRun := true // bool | If true, instructs API to avoid making any changes, but rather run through validations only. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ApiAccountsMgmtV1AccountsPost(context.Background()).Account(account).DryRun(dryRun).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ApiAccountsMgmtV1AccountsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiAccountsMgmtV1AccountsPost`: Account
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ApiAccountsMgmtV1AccountsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiAccountsMgmtV1AccountsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **account** | [**Account**](Account.md) | Account data | 
 **dryRun** | **bool** | If true, instructs API to avoid making any changes, but rather run through validations only. | 

### Return type

[**Account**](Account.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiAccountsMgmtV1ClusterAuthorizationsPost

> ClusterAuthorizationResponse ApiAccountsMgmtV1ClusterAuthorizationsPost(ctx).ClusterAuthorizationRequest(clusterAuthorizationRequest).Execute()

Authorizes new cluster creation against an exsiting RH Subscription.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    clusterAuthorizationRequest := *openapiclient.NewClusterAuthorizationRequest("AccountUsername_example", "ClusterId_example") // ClusterAuthorizationRequest | Cluster and authorization data

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ApiAccountsMgmtV1ClusterAuthorizationsPost(context.Background()).ClusterAuthorizationRequest(clusterAuthorizationRequest).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ApiAccountsMgmtV1ClusterAuthorizationsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiAccountsMgmtV1ClusterAuthorizationsPost`: ClusterAuthorizationResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ApiAccountsMgmtV1ClusterAuthorizationsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiAccountsMgmtV1ClusterAuthorizationsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clusterAuthorizationRequest** | [**ClusterAuthorizationRequest**](ClusterAuthorizationRequest.md) | Cluster and authorization data | 

### Return type

[**ClusterAuthorizationResponse**](ClusterAuthorizationResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiAccountsMgmtV1ClusterRegistrationsPost

> ClusterRegistrationResponse ApiAccountsMgmtV1ClusterRegistrationsPost(ctx).ClusterRegistrationRequest(clusterRegistrationRequest).Execute()

Finds or creates a cluster registration with a registy credential token and cluster ID

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    clusterRegistrationRequest := *openapiclient.NewClusterRegistrationRequest() // ClusterRegistrationRequest | Cluster and authorization data

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ApiAccountsMgmtV1ClusterRegistrationsPost(context.Background()).ClusterRegistrationRequest(clusterRegistrationRequest).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ApiAccountsMgmtV1ClusterRegistrationsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiAccountsMgmtV1ClusterRegistrationsPost`: ClusterRegistrationResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ApiAccountsMgmtV1ClusterRegistrationsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiAccountsMgmtV1ClusterRegistrationsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clusterRegistrationRequest** | [**ClusterRegistrationRequest**](ClusterRegistrationRequest.md) | Cluster and authorization data | 

### Return type

[**ClusterRegistrationResponse**](ClusterRegistrationResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiAccountsMgmtV1CurrentAccountGet

> Account ApiAccountsMgmtV1CurrentAccountGet(ctx).FetchLabels(fetchLabels).Execute()

Get the authenticated account

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    fetchLabels := true // bool | If true, includes the labels on a subscription/organization/account in the output. Could slow request response time. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ApiAccountsMgmtV1CurrentAccountGet(context.Background()).FetchLabels(fetchLabels).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ApiAccountsMgmtV1CurrentAccountGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiAccountsMgmtV1CurrentAccountGet`: Account
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ApiAccountsMgmtV1CurrentAccountGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiAccountsMgmtV1CurrentAccountGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fetchLabels** | **bool** | If true, includes the labels on a subscription/organization/account in the output. Could slow request response time. | 

### Return type

[**Account**](Account.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiAccountsMgmtV1ErrorsGet

> ErrorList ApiAccountsMgmtV1ErrorsGet(ctx).Page(page).Size(size).Search(search).Execute()

Returns a list of errors

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    page := int32(56) // int32 | Page number of record list when record list exceeds specified page size (optional) (default to 1)
    size := int32(56) // int32 | Maximum number of records to return (optional) (default to 100)
    search := "search_example" // string | Specifies the search criteria. The syntax of this parameter is similar to the syntax of the _where_ clause of an SQL statement, using the names of the json attributes / column names of the account. For example, in order to retrieve all the accounts with a username starting with `my`:  ```sql username like 'my%' ```  The search criteria can also be applied on related resource. For example, in order to retrieve all the subscriptions labeled by `foo=bar`,  ```sql labels.key = 'foo' and labels.value = 'bar' ```  If the parameter isn't provided, or if the value is empty, then all the accounts that the user has permission to see will be returned. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ApiAccountsMgmtV1ErrorsGet(context.Background()).Page(page).Size(size).Search(search).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ApiAccountsMgmtV1ErrorsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiAccountsMgmtV1ErrorsGet`: ErrorList
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ApiAccountsMgmtV1ErrorsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiAccountsMgmtV1ErrorsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Page number of record list when record list exceeds specified page size | [default to 1]
 **size** | **int32** | Maximum number of records to return | [default to 100]
 **search** | **string** | Specifies the search criteria. The syntax of this parameter is similar to the syntax of the _where_ clause of an SQL statement, using the names of the json attributes / column names of the account. For example, in order to retrieve all the accounts with a username starting with &#x60;my&#x60;:  &#x60;&#x60;&#x60;sql username like &#39;my%&#39; &#x60;&#x60;&#x60;  The search criteria can also be applied on related resource. For example, in order to retrieve all the subscriptions labeled by &#x60;foo&#x3D;bar&#x60;,  &#x60;&#x60;&#x60;sql labels.key &#x3D; &#39;foo&#39; and labels.value &#x3D; &#39;bar&#39; &#x60;&#x60;&#x60;  If the parameter isn&#39;t provided, or if the value is empty, then all the accounts that the user has permission to see will be returned. | 

### Return type

[**ErrorList**](ErrorList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiAccountsMgmtV1ErrorsIdGet

> Error ApiAccountsMgmtV1ErrorsIdGet(ctx, id).Execute()

Get an error by id

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "id_example" // string | The id of record

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ApiAccountsMgmtV1ErrorsIdGet(context.Background(), id).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ApiAccountsMgmtV1ErrorsIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiAccountsMgmtV1ErrorsIdGet`: Error
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ApiAccountsMgmtV1ErrorsIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of record | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiAccountsMgmtV1ErrorsIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Error**](Error.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiAccountsMgmtV1FeatureTogglesIdQueryPost

> FeatureToggle ApiAccountsMgmtV1FeatureTogglesIdQueryPost(ctx, id).FeatureToggleQueryRequest(featureToggleQueryRequest).Execute()

Query a feature toggle by id

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "id_example" // string | The id of record
    featureToggleQueryRequest := *openapiclient.NewFeatureToggleQueryRequest("OrganizationId_example") // FeatureToggleQueryRequest | The context of the query

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ApiAccountsMgmtV1FeatureTogglesIdQueryPost(context.Background(), id).FeatureToggleQueryRequest(featureToggleQueryRequest).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ApiAccountsMgmtV1FeatureTogglesIdQueryPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiAccountsMgmtV1FeatureTogglesIdQueryPost`: FeatureToggle
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ApiAccountsMgmtV1FeatureTogglesIdQueryPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of record | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiAccountsMgmtV1FeatureTogglesIdQueryPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **featureToggleQueryRequest** | [**FeatureToggleQueryRequest**](FeatureToggleQueryRequest.md) | The context of the query | 

### Return type

[**FeatureToggle**](FeatureToggle.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiAccountsMgmtV1LabelsGet

> LabelList ApiAccountsMgmtV1LabelsGet(ctx).Page(page).Size(size).Search(search).OrderBy(orderBy).Execute()

Returns a list of labels

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    page := int32(56) // int32 | Page number of record list when record list exceeds specified page size (optional) (default to 1)
    size := int32(56) // int32 | Maximum number of records to return (optional) (default to 100)
    search := "search_example" // string | Specifies the search criteria. The syntax of this parameter is similar to the syntax of the _where_ clause of an SQL statement, using the names of the json attributes / column names of the account. For example, in order to retrieve all the accounts with a username starting with `my`:  ```sql username like 'my%' ```  The search criteria can also be applied on related resource. For example, in order to retrieve all the subscriptions labeled by `foo=bar`,  ```sql labels.key = 'foo' and labels.value = 'bar' ```  If the parameter isn't provided, or if the value is empty, then all the accounts that the user has permission to see will be returned. (optional)
    orderBy := "orderBy_example" // string | Specifies the order by criteria. The syntax of this parameter is similar to the syntax of the _order by_ clause of an SQL statement, but using the names of the json attributes / column of the account. For example, in order to retrieve all accounts ordered by username:  ```sql username asc ```  Or in order to retrieve all accounts ordered by username _and_ first name:  ```sql username asc, firstName asc ```  If the parameter isn't provided, or if the value is empty, then no explicit ordering will be applied. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ApiAccountsMgmtV1LabelsGet(context.Background()).Page(page).Size(size).Search(search).OrderBy(orderBy).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ApiAccountsMgmtV1LabelsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiAccountsMgmtV1LabelsGet`: LabelList
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ApiAccountsMgmtV1LabelsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiAccountsMgmtV1LabelsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Page number of record list when record list exceeds specified page size | [default to 1]
 **size** | **int32** | Maximum number of records to return | [default to 100]
 **search** | **string** | Specifies the search criteria. The syntax of this parameter is similar to the syntax of the _where_ clause of an SQL statement, using the names of the json attributes / column names of the account. For example, in order to retrieve all the accounts with a username starting with &#x60;my&#x60;:  &#x60;&#x60;&#x60;sql username like &#39;my%&#39; &#x60;&#x60;&#x60;  The search criteria can also be applied on related resource. For example, in order to retrieve all the subscriptions labeled by &#x60;foo&#x3D;bar&#x60;,  &#x60;&#x60;&#x60;sql labels.key &#x3D; &#39;foo&#39; and labels.value &#x3D; &#39;bar&#39; &#x60;&#x60;&#x60;  If the parameter isn&#39;t provided, or if the value is empty, then all the accounts that the user has permission to see will be returned. | 
 **orderBy** | **string** | Specifies the order by criteria. The syntax of this parameter is similar to the syntax of the _order by_ clause of an SQL statement, but using the names of the json attributes / column of the account. For example, in order to retrieve all accounts ordered by username:  &#x60;&#x60;&#x60;sql username asc &#x60;&#x60;&#x60;  Or in order to retrieve all accounts ordered by username _and_ first name:  &#x60;&#x60;&#x60;sql username asc, firstName asc &#x60;&#x60;&#x60;  If the parameter isn&#39;t provided, or if the value is empty, then no explicit ordering will be applied. | 

### Return type

[**LabelList**](LabelList.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiAccountsMgmtV1MetricsGet

> MetricsList ApiAccountsMgmtV1MetricsGet(ctx).Search(search).Execute()

Returns a list of metrics

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    search := "search_example" // string | Specifies the search criteria. The syntax of this parameter is similar to the syntax of the _where_ clause of an SQL statement, using the names of the json attributes / column names of the account. For example, in order to retrieve all the accounts with a username starting with `my`:  ```sql username like 'my%' ```  The search criteria can also be applied on related resource. For example, in order to retrieve all the subscriptions labeled by `foo=bar`,  ```sql labels.key = 'foo' and labels.value = 'bar' ```  If the parameter isn't provided, or if the value is empty, then all the accounts that the user has permission to see will be returned. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ApiAccountsMgmtV1MetricsGet(context.Background()).Search(search).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ApiAccountsMgmtV1MetricsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiAccountsMgmtV1MetricsGet`: MetricsList
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ApiAccountsMgmtV1MetricsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiAccountsMgmtV1MetricsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **search** | **string** | Specifies the search criteria. The syntax of this parameter is similar to the syntax of the _where_ clause of an SQL statement, using the names of the json attributes / column names of the account. For example, in order to retrieve all the accounts with a username starting with &#x60;my&#x60;:  &#x60;&#x60;&#x60;sql username like &#39;my%&#39; &#x60;&#x60;&#x60;  The search criteria can also be applied on related resource. For example, in order to retrieve all the subscriptions labeled by &#x60;foo&#x3D;bar&#x60;,  &#x60;&#x60;&#x60;sql labels.key &#x3D; &#39;foo&#39; and labels.value &#x3D; &#39;bar&#39; &#x60;&#x60;&#x60;  If the parameter isn&#39;t provided, or if the value is empty, then all the accounts that the user has permission to see will be returned. | 

### Return type

[**MetricsList**](MetricsList.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiAccountsMgmtV1NotifyPost

> ApiAccountsMgmtV1NotifyPost(ctx).NotificationRequest(notificationRequest).Execute()

Notify the owner of cluster/subscription

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    notificationRequest := *openapiclient.NewNotificationRequest("Subject_example", "TemplateName_example") // NotificationRequest | The contents of the notification to send to the owner of a cluster/subscription in addition to the set of template parameters which are sent automatically ACCOUNT_USERNAME, FIRST_NAME, LAST_NAME, ORGANIZATION_NAME, ORGANIZATION_EXTERNAL_ID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ApiAccountsMgmtV1NotifyPost(context.Background()).NotificationRequest(notificationRequest).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ApiAccountsMgmtV1NotifyPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiAccountsMgmtV1NotifyPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **notificationRequest** | [**NotificationRequest**](NotificationRequest.md) | The contents of the notification to send to the owner of a cluster/subscription in addition to the set of template parameters which are sent automatically ACCOUNT_USERNAME, FIRST_NAME, LAST_NAME, ORGANIZATION_NAME, ORGANIZATION_EXTERNAL_ID | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiAccountsMgmtV1OrganizationsGet

> OrganizationList ApiAccountsMgmtV1OrganizationsGet(ctx).Page(page).Size(size).Search(search).OrderBy(orderBy).FetchLabels(fetchLabels).FetchCapabilities(fetchCapabilities).Fields(fields).Execute()

Returns a list of organizations

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    page := int32(56) // int32 | Page number of record list when record list exceeds specified page size (optional) (default to 1)
    size := int32(56) // int32 | Maximum number of records to return (optional) (default to 100)
    search := "search_example" // string | Specifies the search criteria. The syntax of this parameter is similar to the syntax of the _where_ clause of an SQL statement, using the names of the json attributes / column names of the account. For example, in order to retrieve all the accounts with a username starting with `my`:  ```sql username like 'my%' ```  The search criteria can also be applied on related resource. For example, in order to retrieve all the subscriptions labeled by `foo=bar`,  ```sql labels.key = 'foo' and labels.value = 'bar' ```  If the parameter isn't provided, or if the value is empty, then all the accounts that the user has permission to see will be returned. (optional)
    orderBy := "orderBy_example" // string | Specifies the order by criteria. The syntax of this parameter is similar to the syntax of the _order by_ clause of an SQL statement, but using the names of the json attributes / column of the account. For example, in order to retrieve all accounts ordered by username:  ```sql username asc ```  Or in order to retrieve all accounts ordered by username _and_ first name:  ```sql username asc, firstName asc ```  If the parameter isn't provided, or if the value is empty, then no explicit ordering will be applied. (optional)
    fetchLabels := true // bool | If true, includes the labels on a subscription/organization/account in the output. Could slow request response time. (optional)
    fetchCapabilities := true // bool | If true, includes the capabilities on a subscription in the output. Could slow request response time. (optional)
    fields := "fields_example" // string | Supplies a comma-separated list of fields to be returned. Fields of sub-structures and of arrays use <structure>.<field> notation. <stucture>.* means all field of a structure Example: For each Subscription to get id, href, plan(id and kind) and labels (all fields)  ``` ocm get subscriptions --parameter fields=id,href,plan.id,plan.kind,labels.* --parameter fetchLabels=true ``` (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ApiAccountsMgmtV1OrganizationsGet(context.Background()).Page(page).Size(size).Search(search).OrderBy(orderBy).FetchLabels(fetchLabels).FetchCapabilities(fetchCapabilities).Fields(fields).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ApiAccountsMgmtV1OrganizationsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiAccountsMgmtV1OrganizationsGet`: OrganizationList
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ApiAccountsMgmtV1OrganizationsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiAccountsMgmtV1OrganizationsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Page number of record list when record list exceeds specified page size | [default to 1]
 **size** | **int32** | Maximum number of records to return | [default to 100]
 **search** | **string** | Specifies the search criteria. The syntax of this parameter is similar to the syntax of the _where_ clause of an SQL statement, using the names of the json attributes / column names of the account. For example, in order to retrieve all the accounts with a username starting with &#x60;my&#x60;:  &#x60;&#x60;&#x60;sql username like &#39;my%&#39; &#x60;&#x60;&#x60;  The search criteria can also be applied on related resource. For example, in order to retrieve all the subscriptions labeled by &#x60;foo&#x3D;bar&#x60;,  &#x60;&#x60;&#x60;sql labels.key &#x3D; &#39;foo&#39; and labels.value &#x3D; &#39;bar&#39; &#x60;&#x60;&#x60;  If the parameter isn&#39;t provided, or if the value is empty, then all the accounts that the user has permission to see will be returned. | 
 **orderBy** | **string** | Specifies the order by criteria. The syntax of this parameter is similar to the syntax of the _order by_ clause of an SQL statement, but using the names of the json attributes / column of the account. For example, in order to retrieve all accounts ordered by username:  &#x60;&#x60;&#x60;sql username asc &#x60;&#x60;&#x60;  Or in order to retrieve all accounts ordered by username _and_ first name:  &#x60;&#x60;&#x60;sql username asc, firstName asc &#x60;&#x60;&#x60;  If the parameter isn&#39;t provided, or if the value is empty, then no explicit ordering will be applied. | 
 **fetchLabels** | **bool** | If true, includes the labels on a subscription/organization/account in the output. Could slow request response time. | 
 **fetchCapabilities** | **bool** | If true, includes the capabilities on a subscription in the output. Could slow request response time. | 
 **fields** | **string** | Supplies a comma-separated list of fields to be returned. Fields of sub-structures and of arrays use &lt;structure&gt;.&lt;field&gt; notation. &lt;stucture&gt;.* means all field of a structure Example: For each Subscription to get id, href, plan(id and kind) and labels (all fields)  &#x60;&#x60;&#x60; ocm get subscriptions --parameter fields&#x3D;id,href,plan.id,plan.kind,labels.* --parameter fetchLabels&#x3D;true &#x60;&#x60;&#x60; | 

### Return type

[**OrganizationList**](OrganizationList.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiAccountsMgmtV1OrganizationsIdGet

> Organization ApiAccountsMgmtV1OrganizationsIdGet(ctx, id).FetchLabels(fetchLabels).FetchCapabilities(fetchCapabilities).Execute()

Get an organization by id

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "id_example" // string | The id of record
    fetchLabels := true // bool | If true, includes the labels on a subscription/organization/account in the output. Could slow request response time. (optional)
    fetchCapabilities := true // bool | If true, includes the capabilities on a subscription in the output. Could slow request response time. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ApiAccountsMgmtV1OrganizationsIdGet(context.Background(), id).FetchLabels(fetchLabels).FetchCapabilities(fetchCapabilities).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ApiAccountsMgmtV1OrganizationsIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiAccountsMgmtV1OrganizationsIdGet`: Organization
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ApiAccountsMgmtV1OrganizationsIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of record | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiAccountsMgmtV1OrganizationsIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fetchLabels** | **bool** | If true, includes the labels on a subscription/organization/account in the output. Could slow request response time. | 
 **fetchCapabilities** | **bool** | If true, includes the capabilities on a subscription in the output. Could slow request response time. | 

### Return type

[**Organization**](Organization.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiAccountsMgmtV1OrganizationsIdLabelsGet

> LabelList ApiAccountsMgmtV1OrganizationsIdLabelsGet(ctx, id).Page(page).Size(size).Search(search).OrderBy(orderBy).Execute()

Returns a list of labels

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "id_example" // string | The id of record
    page := int32(56) // int32 | Page number of record list when record list exceeds specified page size (optional) (default to 1)
    size := int32(56) // int32 | Maximum number of records to return (optional) (default to 100)
    search := "search_example" // string | Specifies the search criteria. The syntax of this parameter is similar to the syntax of the _where_ clause of an SQL statement, using the names of the json attributes / column names of the account. For example, in order to retrieve all the accounts with a username starting with `my`:  ```sql username like 'my%' ```  The search criteria can also be applied on related resource. For example, in order to retrieve all the subscriptions labeled by `foo=bar`,  ```sql labels.key = 'foo' and labels.value = 'bar' ```  If the parameter isn't provided, or if the value is empty, then all the accounts that the user has permission to see will be returned. (optional)
    orderBy := "orderBy_example" // string | Specifies the order by criteria. The syntax of this parameter is similar to the syntax of the _order by_ clause of an SQL statement, but using the names of the json attributes / column of the account. For example, in order to retrieve all accounts ordered by username:  ```sql username asc ```  Or in order to retrieve all accounts ordered by username _and_ first name:  ```sql username asc, firstName asc ```  If the parameter isn't provided, or if the value is empty, then no explicit ordering will be applied. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ApiAccountsMgmtV1OrganizationsIdLabelsGet(context.Background(), id).Page(page).Size(size).Search(search).OrderBy(orderBy).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ApiAccountsMgmtV1OrganizationsIdLabelsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiAccountsMgmtV1OrganizationsIdLabelsGet`: LabelList
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ApiAccountsMgmtV1OrganizationsIdLabelsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of record | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiAccountsMgmtV1OrganizationsIdLabelsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | Page number of record list when record list exceeds specified page size | [default to 1]
 **size** | **int32** | Maximum number of records to return | [default to 100]
 **search** | **string** | Specifies the search criteria. The syntax of this parameter is similar to the syntax of the _where_ clause of an SQL statement, using the names of the json attributes / column names of the account. For example, in order to retrieve all the accounts with a username starting with &#x60;my&#x60;:  &#x60;&#x60;&#x60;sql username like &#39;my%&#39; &#x60;&#x60;&#x60;  The search criteria can also be applied on related resource. For example, in order to retrieve all the subscriptions labeled by &#x60;foo&#x3D;bar&#x60;,  &#x60;&#x60;&#x60;sql labels.key &#x3D; &#39;foo&#39; and labels.value &#x3D; &#39;bar&#39; &#x60;&#x60;&#x60;  If the parameter isn&#39;t provided, or if the value is empty, then all the accounts that the user has permission to see will be returned. | 
 **orderBy** | **string** | Specifies the order by criteria. The syntax of this parameter is similar to the syntax of the _order by_ clause of an SQL statement, but using the names of the json attributes / column of the account. For example, in order to retrieve all accounts ordered by username:  &#x60;&#x60;&#x60;sql username asc &#x60;&#x60;&#x60;  Or in order to retrieve all accounts ordered by username _and_ first name:  &#x60;&#x60;&#x60;sql username asc, firstName asc &#x60;&#x60;&#x60;  If the parameter isn&#39;t provided, or if the value is empty, then no explicit ordering will be applied. | 

### Return type

[**LabelList**](LabelList.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiAccountsMgmtV1OrganizationsIdLabelsKeyDelete

> ApiAccountsMgmtV1OrganizationsIdLabelsKeyDelete(ctx, id, key).Execute()

Delete a label

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "id_example" // string | The id of record
    key := "key_example" // string | The key of the label

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ApiAccountsMgmtV1OrganizationsIdLabelsKeyDelete(context.Background(), id, key).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ApiAccountsMgmtV1OrganizationsIdLabelsKeyDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of record | 
**key** | **string** | The key of the label | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiAccountsMgmtV1OrganizationsIdLabelsKeyDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiAccountsMgmtV1OrganizationsIdLabelsKeyGet

> Label ApiAccountsMgmtV1OrganizationsIdLabelsKeyGet(ctx, id, key).Execute()

Get subscription labels by label key

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "id_example" // string | The id of record
    key := "key_example" // string | The key of the label

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ApiAccountsMgmtV1OrganizationsIdLabelsKeyGet(context.Background(), id, key).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ApiAccountsMgmtV1OrganizationsIdLabelsKeyGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiAccountsMgmtV1OrganizationsIdLabelsKeyGet`: Label
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ApiAccountsMgmtV1OrganizationsIdLabelsKeyGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of record | 
**key** | **string** | The key of the label | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiAccountsMgmtV1OrganizationsIdLabelsKeyGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Label**](Label.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiAccountsMgmtV1OrganizationsIdLabelsKeyPatch

> Label ApiAccountsMgmtV1OrganizationsIdLabelsKeyPatch(ctx, id, key).Label(label).Execute()

Create a new label or update an existing label

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "id_example" // string | The id of record
    key := "key_example" // string | The key of the label
    label := *openapiclient.NewLabel(false, "Key_example", "Value_example") // Label | Label data

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ApiAccountsMgmtV1OrganizationsIdLabelsKeyPatch(context.Background(), id, key).Label(label).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ApiAccountsMgmtV1OrganizationsIdLabelsKeyPatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiAccountsMgmtV1OrganizationsIdLabelsKeyPatch`: Label
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ApiAccountsMgmtV1OrganizationsIdLabelsKeyPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of record | 
**key** | **string** | The key of the label | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiAccountsMgmtV1OrganizationsIdLabelsKeyPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **label** | [**Label**](Label.md) | Label data | 

### Return type

[**Label**](Label.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiAccountsMgmtV1OrganizationsIdLabelsPost

> Label ApiAccountsMgmtV1OrganizationsIdLabelsPost(ctx, id).Label(label).Execute()

Create a new label or update an existing label

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "id_example" // string | The id of record
    label := *openapiclient.NewLabel(false, "Key_example", "Value_example") // Label | Label data

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ApiAccountsMgmtV1OrganizationsIdLabelsPost(context.Background(), id).Label(label).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ApiAccountsMgmtV1OrganizationsIdLabelsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiAccountsMgmtV1OrganizationsIdLabelsPost`: Label
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ApiAccountsMgmtV1OrganizationsIdLabelsPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of record | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiAccountsMgmtV1OrganizationsIdLabelsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **label** | [**Label**](Label.md) | Label data | 

### Return type

[**Label**](Label.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiAccountsMgmtV1OrganizationsIdPatch

> Organization ApiAccountsMgmtV1OrganizationsIdPatch(ctx, id).OrganizationPatchRequest(organizationPatchRequest).Execute()

Update an organization

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "id_example" // string | The id of record
    organizationPatchRequest := *openapiclient.NewOrganizationPatchRequest() // OrganizationPatchRequest | Updated organization data

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ApiAccountsMgmtV1OrganizationsIdPatch(context.Background(), id).OrganizationPatchRequest(organizationPatchRequest).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ApiAccountsMgmtV1OrganizationsIdPatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiAccountsMgmtV1OrganizationsIdPatch`: Organization
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ApiAccountsMgmtV1OrganizationsIdPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of record | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiAccountsMgmtV1OrganizationsIdPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **organizationPatchRequest** | [**OrganizationPatchRequest**](OrganizationPatchRequest.md) | Updated organization data | 

### Return type

[**Organization**](Organization.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiAccountsMgmtV1OrganizationsIdSummaryDashboardGet

> Summary ApiAccountsMgmtV1OrganizationsIdSummaryDashboardGet(ctx, id).Execute()

Returns a summary of organizations clusters based on metrics

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "id_example" // string | The id of record

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ApiAccountsMgmtV1OrganizationsIdSummaryDashboardGet(context.Background(), id).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ApiAccountsMgmtV1OrganizationsIdSummaryDashboardGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiAccountsMgmtV1OrganizationsIdSummaryDashboardGet`: Summary
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ApiAccountsMgmtV1OrganizationsIdSummaryDashboardGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of record | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiAccountsMgmtV1OrganizationsIdSummaryDashboardGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Summary**](Summary.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiAccountsMgmtV1OrganizationsOrgIdQuotaCostGet

> QuotaCostList ApiAccountsMgmtV1OrganizationsOrgIdQuotaCostGet(ctx, orgId).Search(search).FetchRelatedResources(fetchRelatedResources).Execute()

Returns a summary of quota cost

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    orgId := "orgId_example" // string | The id of organization
    search := "search_example" // string | Specifies the search criteria. The syntax of this parameter is similar to the syntax of the _where_ clause of an SQL statement, using the names of the json attributes / column names of the account. For example, in order to retrieve all the accounts with a username starting with `my`:  ```sql username like 'my%' ```  The search criteria can also be applied on related resource. For example, in order to retrieve all the subscriptions labeled by `foo=bar`,  ```sql labels.key = 'foo' and labels.value = 'bar' ```  If the parameter isn't provided, or if the value is empty, then all the accounts that the user has permission to see will be returned. (optional)
    fetchRelatedResources := true // bool | If true, includes the related resources in the output. Could slow request response time. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ApiAccountsMgmtV1OrganizationsOrgIdQuotaCostGet(context.Background(), orgId).Search(search).FetchRelatedResources(fetchRelatedResources).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ApiAccountsMgmtV1OrganizationsOrgIdQuotaCostGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiAccountsMgmtV1OrganizationsOrgIdQuotaCostGet`: QuotaCostList
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ApiAccountsMgmtV1OrganizationsOrgIdQuotaCostGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | The id of organization | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiAccountsMgmtV1OrganizationsOrgIdQuotaCostGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **search** | **string** | Specifies the search criteria. The syntax of this parameter is similar to the syntax of the _where_ clause of an SQL statement, using the names of the json attributes / column names of the account. For example, in order to retrieve all the accounts with a username starting with &#x60;my&#x60;:  &#x60;&#x60;&#x60;sql username like &#39;my%&#39; &#x60;&#x60;&#x60;  The search criteria can also be applied on related resource. For example, in order to retrieve all the subscriptions labeled by &#x60;foo&#x3D;bar&#x60;,  &#x60;&#x60;&#x60;sql labels.key &#x3D; &#39;foo&#39; and labels.value &#x3D; &#39;bar&#39; &#x60;&#x60;&#x60;  If the parameter isn&#39;t provided, or if the value is empty, then all the accounts that the user has permission to see will be returned. | 
 **fetchRelatedResources** | **bool** | If true, includes the related resources in the output. Could slow request response time. | 

### Return type

[**QuotaCostList**](QuotaCostList.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiAccountsMgmtV1OrganizationsOrgIdQuotaSummaryGet

> QuotaSummaryList ApiAccountsMgmtV1OrganizationsOrgIdQuotaSummaryGet(ctx, orgId).Search(search).Execute()

Returns a summary of resource quota

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    orgId := "orgId_example" // string | The id of organization
    search := "search_example" // string | Specifies the search criteria. The syntax of this parameter is similar to the syntax of the _where_ clause of an SQL statement, using the names of the json attributes / column names of the account. For example, in order to retrieve all the accounts with a username starting with `my`:  ```sql username like 'my%' ```  The search criteria can also be applied on related resource. For example, in order to retrieve all the subscriptions labeled by `foo=bar`,  ```sql labels.key = 'foo' and labels.value = 'bar' ```  If the parameter isn't provided, or if the value is empty, then all the accounts that the user has permission to see will be returned. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ApiAccountsMgmtV1OrganizationsOrgIdQuotaSummaryGet(context.Background(), orgId).Search(search).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ApiAccountsMgmtV1OrganizationsOrgIdQuotaSummaryGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiAccountsMgmtV1OrganizationsOrgIdQuotaSummaryGet`: QuotaSummaryList
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ApiAccountsMgmtV1OrganizationsOrgIdQuotaSummaryGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | The id of organization | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiAccountsMgmtV1OrganizationsOrgIdQuotaSummaryGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **search** | **string** | Specifies the search criteria. The syntax of this parameter is similar to the syntax of the _where_ clause of an SQL statement, using the names of the json attributes / column names of the account. For example, in order to retrieve all the accounts with a username starting with &#x60;my&#x60;:  &#x60;&#x60;&#x60;sql username like &#39;my%&#39; &#x60;&#x60;&#x60;  The search criteria can also be applied on related resource. For example, in order to retrieve all the subscriptions labeled by &#x60;foo&#x3D;bar&#x60;,  &#x60;&#x60;&#x60;sql labels.key &#x3D; &#39;foo&#39; and labels.value &#x3D; &#39;bar&#39; &#x60;&#x60;&#x60;  If the parameter isn&#39;t provided, or if the value is empty, then all the accounts that the user has permission to see will be returned. | 

### Return type

[**QuotaSummaryList**](QuotaSummaryList.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiAccountsMgmtV1OrganizationsOrgIdResourceQuotaGet

> ResourceQuotaList ApiAccountsMgmtV1OrganizationsOrgIdResourceQuotaGet(ctx, orgId).Page(page).Size(size).Search(search).OrderBy(orderBy).Execute()

Returns a list of resource quota objects

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    orgId := "orgId_example" // string | The id of organization
    page := int32(56) // int32 | Page number of record list when record list exceeds specified page size (optional) (default to 1)
    size := int32(56) // int32 | Maximum number of records to return (optional) (default to 100)
    search := "search_example" // string | Specifies the search criteria. The syntax of this parameter is similar to the syntax of the _where_ clause of an SQL statement, using the names of the json attributes / column names of the account. For example, in order to retrieve all the accounts with a username starting with `my`:  ```sql username like 'my%' ```  The search criteria can also be applied on related resource. For example, in order to retrieve all the subscriptions labeled by `foo=bar`,  ```sql labels.key = 'foo' and labels.value = 'bar' ```  If the parameter isn't provided, or if the value is empty, then all the accounts that the user has permission to see will be returned. (optional)
    orderBy := "orderBy_example" // string | Specifies the order by criteria. The syntax of this parameter is similar to the syntax of the _order by_ clause of an SQL statement, but using the names of the json attributes / column of the account. For example, in order to retrieve all accounts ordered by username:  ```sql username asc ```  Or in order to retrieve all accounts ordered by username _and_ first name:  ```sql username asc, firstName asc ```  If the parameter isn't provided, or if the value is empty, then no explicit ordering will be applied. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ApiAccountsMgmtV1OrganizationsOrgIdResourceQuotaGet(context.Background(), orgId).Page(page).Size(size).Search(search).OrderBy(orderBy).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ApiAccountsMgmtV1OrganizationsOrgIdResourceQuotaGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiAccountsMgmtV1OrganizationsOrgIdResourceQuotaGet`: ResourceQuotaList
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ApiAccountsMgmtV1OrganizationsOrgIdResourceQuotaGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | The id of organization | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiAccountsMgmtV1OrganizationsOrgIdResourceQuotaGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | Page number of record list when record list exceeds specified page size | [default to 1]
 **size** | **int32** | Maximum number of records to return | [default to 100]
 **search** | **string** | Specifies the search criteria. The syntax of this parameter is similar to the syntax of the _where_ clause of an SQL statement, using the names of the json attributes / column names of the account. For example, in order to retrieve all the accounts with a username starting with &#x60;my&#x60;:  &#x60;&#x60;&#x60;sql username like &#39;my%&#39; &#x60;&#x60;&#x60;  The search criteria can also be applied on related resource. For example, in order to retrieve all the subscriptions labeled by &#x60;foo&#x3D;bar&#x60;,  &#x60;&#x60;&#x60;sql labels.key &#x3D; &#39;foo&#39; and labels.value &#x3D; &#39;bar&#39; &#x60;&#x60;&#x60;  If the parameter isn&#39;t provided, or if the value is empty, then all the accounts that the user has permission to see will be returned. | 
 **orderBy** | **string** | Specifies the order by criteria. The syntax of this parameter is similar to the syntax of the _order by_ clause of an SQL statement, but using the names of the json attributes / column of the account. For example, in order to retrieve all accounts ordered by username:  &#x60;&#x60;&#x60;sql username asc &#x60;&#x60;&#x60;  Or in order to retrieve all accounts ordered by username _and_ first name:  &#x60;&#x60;&#x60;sql username asc, firstName asc &#x60;&#x60;&#x60;  If the parameter isn&#39;t provided, or if the value is empty, then no explicit ordering will be applied. | 

### Return type

[**ResourceQuotaList**](ResourceQuotaList.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiAccountsMgmtV1OrganizationsOrgIdResourceQuotaPost

> ResourceQuota ApiAccountsMgmtV1OrganizationsOrgIdResourceQuotaPost(ctx, orgId).ResourceQuotaRequest(resourceQuotaRequest).Execute()

Create a new resource quota

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    orgId := "orgId_example" // string | The id of organization
    resourceQuotaRequest := *openapiclient.NewResourceQuotaRequest("Sku_example") // ResourceQuotaRequest | Resource quota data

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ApiAccountsMgmtV1OrganizationsOrgIdResourceQuotaPost(context.Background(), orgId).ResourceQuotaRequest(resourceQuotaRequest).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ApiAccountsMgmtV1OrganizationsOrgIdResourceQuotaPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiAccountsMgmtV1OrganizationsOrgIdResourceQuotaPost`: ResourceQuota
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ApiAccountsMgmtV1OrganizationsOrgIdResourceQuotaPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | The id of organization | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiAccountsMgmtV1OrganizationsOrgIdResourceQuotaPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **resourceQuotaRequest** | [**ResourceQuotaRequest**](ResourceQuotaRequest.md) | Resource quota data | 

### Return type

[**ResourceQuota**](ResourceQuota.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiAccountsMgmtV1OrganizationsOrgIdResourceQuotaQuotaIdDelete

> ApiAccountsMgmtV1OrganizationsOrgIdResourceQuotaQuotaIdDelete(ctx, orgId, quotaId).Execute()

Delete a resource quota

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    orgId := "orgId_example" // string | The id of organization
    quotaId := "quotaId_example" // string | The id of quota

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ApiAccountsMgmtV1OrganizationsOrgIdResourceQuotaQuotaIdDelete(context.Background(), orgId, quotaId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ApiAccountsMgmtV1OrganizationsOrgIdResourceQuotaQuotaIdDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | The id of organization | 
**quotaId** | **string** | The id of quota | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiAccountsMgmtV1OrganizationsOrgIdResourceQuotaQuotaIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiAccountsMgmtV1OrganizationsOrgIdResourceQuotaQuotaIdGet

> ResourceQuota ApiAccountsMgmtV1OrganizationsOrgIdResourceQuotaQuotaIdGet(ctx, orgId, quotaId).Execute()

Get a resource quota by id

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    orgId := "orgId_example" // string | The id of organization
    quotaId := "quotaId_example" // string | The id of quota

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ApiAccountsMgmtV1OrganizationsOrgIdResourceQuotaQuotaIdGet(context.Background(), orgId, quotaId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ApiAccountsMgmtV1OrganizationsOrgIdResourceQuotaQuotaIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiAccountsMgmtV1OrganizationsOrgIdResourceQuotaQuotaIdGet`: ResourceQuota
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ApiAccountsMgmtV1OrganizationsOrgIdResourceQuotaQuotaIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | The id of organization | 
**quotaId** | **string** | The id of quota | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiAccountsMgmtV1OrganizationsOrgIdResourceQuotaQuotaIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ResourceQuota**](ResourceQuota.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiAccountsMgmtV1OrganizationsOrgIdResourceQuotaQuotaIdPatch

> ResourceQuota ApiAccountsMgmtV1OrganizationsOrgIdResourceQuotaQuotaIdPatch(ctx, orgId, quotaId).ResourceQuotaRequest(resourceQuotaRequest).Execute()

Update a resource quota

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    orgId := "orgId_example" // string | The id of organization
    quotaId := "quotaId_example" // string | The id of quota
    resourceQuotaRequest := *openapiclient.NewResourceQuotaRequest("Sku_example") // ResourceQuotaRequest | Updated resource quota data

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ApiAccountsMgmtV1OrganizationsOrgIdResourceQuotaQuotaIdPatch(context.Background(), orgId, quotaId).ResourceQuotaRequest(resourceQuotaRequest).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ApiAccountsMgmtV1OrganizationsOrgIdResourceQuotaQuotaIdPatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiAccountsMgmtV1OrganizationsOrgIdResourceQuotaQuotaIdPatch`: ResourceQuota
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ApiAccountsMgmtV1OrganizationsOrgIdResourceQuotaQuotaIdPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | The id of organization | 
**quotaId** | **string** | The id of quota | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiAccountsMgmtV1OrganizationsOrgIdResourceQuotaQuotaIdPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **resourceQuotaRequest** | [**ResourceQuotaRequest**](ResourceQuotaRequest.md) | Updated resource quota data | 

### Return type

[**ResourceQuota**](ResourceQuota.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiAccountsMgmtV1OrganizationsPost

> Organization ApiAccountsMgmtV1OrganizationsPost(ctx).Organization(organization).Execute()

Create a new organization

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    organization := *openapiclient.NewOrganization() // Organization | Organization data

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ApiAccountsMgmtV1OrganizationsPost(context.Background()).Organization(organization).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ApiAccountsMgmtV1OrganizationsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiAccountsMgmtV1OrganizationsPost`: Organization
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ApiAccountsMgmtV1OrganizationsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiAccountsMgmtV1OrganizationsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **organization** | [**Organization**](Organization.md) | Organization data | 

### Return type

[**Organization**](Organization.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiAccountsMgmtV1PlansGet

> PlanList ApiAccountsMgmtV1PlansGet(ctx).Page(page).Size(size).Search(search).Execute()

Get all plans

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    page := int32(56) // int32 | Page number of record list when record list exceeds specified page size (optional) (default to 1)
    size := int32(56) // int32 | Maximum number of records to return (optional) (default to 100)
    search := "search_example" // string | Specifies the search criteria. The syntax of this parameter is similar to the syntax of the _where_ clause of an SQL statement, using the names of the json attributes / column names of the account. For example, in order to retrieve all the accounts with a username starting with `my`:  ```sql username like 'my%' ```  The search criteria can also be applied on related resource. For example, in order to retrieve all the subscriptions labeled by `foo=bar`,  ```sql labels.key = 'foo' and labels.value = 'bar' ```  If the parameter isn't provided, or if the value is empty, then all the accounts that the user has permission to see will be returned. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ApiAccountsMgmtV1PlansGet(context.Background()).Page(page).Size(size).Search(search).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ApiAccountsMgmtV1PlansGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiAccountsMgmtV1PlansGet`: PlanList
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ApiAccountsMgmtV1PlansGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiAccountsMgmtV1PlansGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Page number of record list when record list exceeds specified page size | [default to 1]
 **size** | **int32** | Maximum number of records to return | [default to 100]
 **search** | **string** | Specifies the search criteria. The syntax of this parameter is similar to the syntax of the _where_ clause of an SQL statement, using the names of the json attributes / column names of the account. For example, in order to retrieve all the accounts with a username starting with &#x60;my&#x60;:  &#x60;&#x60;&#x60;sql username like &#39;my%&#39; &#x60;&#x60;&#x60;  The search criteria can also be applied on related resource. For example, in order to retrieve all the subscriptions labeled by &#x60;foo&#x3D;bar&#x60;,  &#x60;&#x60;&#x60;sql labels.key &#x3D; &#39;foo&#39; and labels.value &#x3D; &#39;bar&#39; &#x60;&#x60;&#x60;  If the parameter isn&#39;t provided, or if the value is empty, then all the accounts that the user has permission to see will be returned. | 

### Return type

[**PlanList**](PlanList.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiAccountsMgmtV1PlansIdGet

> Plan ApiAccountsMgmtV1PlansIdGet(ctx, id).Execute()

Get a plan by id

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "id_example" // string | The id of record

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ApiAccountsMgmtV1PlansIdGet(context.Background(), id).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ApiAccountsMgmtV1PlansIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiAccountsMgmtV1PlansIdGet`: Plan
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ApiAccountsMgmtV1PlansIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of record | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiAccountsMgmtV1PlansIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Plan**](Plan.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiAccountsMgmtV1PullSecretsExternalResourceIdDelete

> ApiAccountsMgmtV1PullSecretsExternalResourceIdDelete(ctx, externalResourceId).Execute()

Delete a pull secret

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    externalResourceId := "externalResourceId_example" // string | The external resource id of record

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ApiAccountsMgmtV1PullSecretsExternalResourceIdDelete(context.Background(), externalResourceId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ApiAccountsMgmtV1PullSecretsExternalResourceIdDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**externalResourceId** | **string** | The external resource id of record | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiAccountsMgmtV1PullSecretsExternalResourceIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiAccountsMgmtV1PullSecretsPost

> AccessTokenCfg ApiAccountsMgmtV1PullSecretsPost(ctx).PullSecretRequest(pullSecretRequest).Execute()

Return access token generated from registries in docker format

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    pullSecretRequest := *openapiclient.NewPullSecretRequest("ExternalResourceId_example") // PullSecretRequest | Identifier of the resource in the external service that this pull secret relates to

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ApiAccountsMgmtV1PullSecretsPost(context.Background()).PullSecretRequest(pullSecretRequest).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ApiAccountsMgmtV1PullSecretsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiAccountsMgmtV1PullSecretsPost`: AccessTokenCfg
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ApiAccountsMgmtV1PullSecretsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiAccountsMgmtV1PullSecretsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pullSecretRequest** | [**PullSecretRequest**](PullSecretRequest.md) | Identifier of the resource in the external service that this pull secret relates to | 

### Return type

[**AccessTokenCfg**](AccessTokenCfg.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiAccountsMgmtV1RegistriesGet

> RegistryList ApiAccountsMgmtV1RegistriesGet(ctx).Page(page).Size(size).Search(search).OrderBy(orderBy).Execute()

Returns a list of registries

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    page := int32(56) // int32 | Page number of record list when record list exceeds specified page size (optional) (default to 1)
    size := int32(56) // int32 | Maximum number of records to return (optional) (default to 100)
    search := "search_example" // string | Specifies the search criteria. The syntax of this parameter is similar to the syntax of the _where_ clause of an SQL statement, using the names of the json attributes / column names of the account. For example, in order to retrieve all the accounts with a username starting with `my`:  ```sql username like 'my%' ```  The search criteria can also be applied on related resource. For example, in order to retrieve all the subscriptions labeled by `foo=bar`,  ```sql labels.key = 'foo' and labels.value = 'bar' ```  If the parameter isn't provided, or if the value is empty, then all the accounts that the user has permission to see will be returned. (optional)
    orderBy := "orderBy_example" // string | Specifies the order by criteria. The syntax of this parameter is similar to the syntax of the _order by_ clause of an SQL statement, but using the names of the json attributes / column of the account. For example, in order to retrieve all accounts ordered by username:  ```sql username asc ```  Or in order to retrieve all accounts ordered by username _and_ first name:  ```sql username asc, firstName asc ```  If the parameter isn't provided, or if the value is empty, then no explicit ordering will be applied. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ApiAccountsMgmtV1RegistriesGet(context.Background()).Page(page).Size(size).Search(search).OrderBy(orderBy).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ApiAccountsMgmtV1RegistriesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiAccountsMgmtV1RegistriesGet`: RegistryList
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ApiAccountsMgmtV1RegistriesGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiAccountsMgmtV1RegistriesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Page number of record list when record list exceeds specified page size | [default to 1]
 **size** | **int32** | Maximum number of records to return | [default to 100]
 **search** | **string** | Specifies the search criteria. The syntax of this parameter is similar to the syntax of the _where_ clause of an SQL statement, using the names of the json attributes / column names of the account. For example, in order to retrieve all the accounts with a username starting with &#x60;my&#x60;:  &#x60;&#x60;&#x60;sql username like &#39;my%&#39; &#x60;&#x60;&#x60;  The search criteria can also be applied on related resource. For example, in order to retrieve all the subscriptions labeled by &#x60;foo&#x3D;bar&#x60;,  &#x60;&#x60;&#x60;sql labels.key &#x3D; &#39;foo&#39; and labels.value &#x3D; &#39;bar&#39; &#x60;&#x60;&#x60;  If the parameter isn&#39;t provided, or if the value is empty, then all the accounts that the user has permission to see will be returned. | 
 **orderBy** | **string** | Specifies the order by criteria. The syntax of this parameter is similar to the syntax of the _order by_ clause of an SQL statement, but using the names of the json attributes / column of the account. For example, in order to retrieve all accounts ordered by username:  &#x60;&#x60;&#x60;sql username asc &#x60;&#x60;&#x60;  Or in order to retrieve all accounts ordered by username _and_ first name:  &#x60;&#x60;&#x60;sql username asc, firstName asc &#x60;&#x60;&#x60;  If the parameter isn&#39;t provided, or if the value is empty, then no explicit ordering will be applied. | 

### Return type

[**RegistryList**](RegistryList.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiAccountsMgmtV1RegistriesIdGet

> Registry ApiAccountsMgmtV1RegistriesIdGet(ctx, id).Execute()

Get an registry by id

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "id_example" // string | The id of record

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ApiAccountsMgmtV1RegistriesIdGet(context.Background(), id).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ApiAccountsMgmtV1RegistriesIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiAccountsMgmtV1RegistriesIdGet`: Registry
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ApiAccountsMgmtV1RegistriesIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of record | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiAccountsMgmtV1RegistriesIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Registry**](Registry.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiAccountsMgmtV1RegistryCredentialsGet

> RegistryCredentialList ApiAccountsMgmtV1RegistryCredentialsGet(ctx).Page(page).Size(size).Search(search).OrderBy(orderBy).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    page := int32(56) // int32 | Page number of record list when record list exceeds specified page size (optional) (default to 1)
    size := int32(56) // int32 | Maximum number of records to return (optional) (default to 100)
    search := "search_example" // string | Specifies the search criteria. The syntax of this parameter is similar to the syntax of the _where_ clause of an SQL statement, using the names of the json attributes / column names of the account. For example, in order to retrieve all the accounts with a username starting with `my`:  ```sql username like 'my%' ```  The search criteria can also be applied on related resource. For example, in order to retrieve all the subscriptions labeled by `foo=bar`,  ```sql labels.key = 'foo' and labels.value = 'bar' ```  If the parameter isn't provided, or if the value is empty, then all the accounts that the user has permission to see will be returned. (optional)
    orderBy := "orderBy_example" // string | Specifies the order by criteria. The syntax of this parameter is similar to the syntax of the _order by_ clause of an SQL statement, but using the names of the json attributes / column of the account. For example, in order to retrieve all accounts ordered by username:  ```sql username asc ```  Or in order to retrieve all accounts ordered by username _and_ first name:  ```sql username asc, firstName asc ```  If the parameter isn't provided, or if the value is empty, then no explicit ordering will be applied. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ApiAccountsMgmtV1RegistryCredentialsGet(context.Background()).Page(page).Size(size).Search(search).OrderBy(orderBy).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ApiAccountsMgmtV1RegistryCredentialsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiAccountsMgmtV1RegistryCredentialsGet`: RegistryCredentialList
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ApiAccountsMgmtV1RegistryCredentialsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiAccountsMgmtV1RegistryCredentialsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Page number of record list when record list exceeds specified page size | [default to 1]
 **size** | **int32** | Maximum number of records to return | [default to 100]
 **search** | **string** | Specifies the search criteria. The syntax of this parameter is similar to the syntax of the _where_ clause of an SQL statement, using the names of the json attributes / column names of the account. For example, in order to retrieve all the accounts with a username starting with &#x60;my&#x60;:  &#x60;&#x60;&#x60;sql username like &#39;my%&#39; &#x60;&#x60;&#x60;  The search criteria can also be applied on related resource. For example, in order to retrieve all the subscriptions labeled by &#x60;foo&#x3D;bar&#x60;,  &#x60;&#x60;&#x60;sql labels.key &#x3D; &#39;foo&#39; and labels.value &#x3D; &#39;bar&#39; &#x60;&#x60;&#x60;  If the parameter isn&#39;t provided, or if the value is empty, then all the accounts that the user has permission to see will be returned. | 
 **orderBy** | **string** | Specifies the order by criteria. The syntax of this parameter is similar to the syntax of the _order by_ clause of an SQL statement, but using the names of the json attributes / column of the account. For example, in order to retrieve all accounts ordered by username:  &#x60;&#x60;&#x60;sql username asc &#x60;&#x60;&#x60;  Or in order to retrieve all accounts ordered by username _and_ first name:  &#x60;&#x60;&#x60;sql username asc, firstName asc &#x60;&#x60;&#x60;  If the parameter isn&#39;t provided, or if the value is empty, then no explicit ordering will be applied. | 

### Return type

[**RegistryCredentialList**](RegistryCredentialList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiAccountsMgmtV1RegistryCredentialsIdDelete

> ApiAccountsMgmtV1RegistryCredentialsIdDelete(ctx, id).Execute()

Delete a registry credential by id

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "id_example" // string | The id of record

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ApiAccountsMgmtV1RegistryCredentialsIdDelete(context.Background(), id).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ApiAccountsMgmtV1RegistryCredentialsIdDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of record | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiAccountsMgmtV1RegistryCredentialsIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiAccountsMgmtV1RegistryCredentialsIdGet

> RegistryCredential ApiAccountsMgmtV1RegistryCredentialsIdGet(ctx, id).Execute()

Get a registry credentials by id

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "id_example" // string | The id of record

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ApiAccountsMgmtV1RegistryCredentialsIdGet(context.Background(), id).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ApiAccountsMgmtV1RegistryCredentialsIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiAccountsMgmtV1RegistryCredentialsIdGet`: RegistryCredential
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ApiAccountsMgmtV1RegistryCredentialsIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of record | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiAccountsMgmtV1RegistryCredentialsIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RegistryCredential**](RegistryCredential.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiAccountsMgmtV1RegistryCredentialsIdPatch

> RegistryCredential ApiAccountsMgmtV1RegistryCredentialsIdPatch(ctx, id).RegistryCredentialPatchRequest(registryCredentialPatchRequest).Execute()

Update a registry credential

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "id_example" // string | The id of record
    registryCredentialPatchRequest := *openapiclient.NewRegistryCredentialPatchRequest() // RegistryCredentialPatchRequest | Updated registry credential data

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ApiAccountsMgmtV1RegistryCredentialsIdPatch(context.Background(), id).RegistryCredentialPatchRequest(registryCredentialPatchRequest).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ApiAccountsMgmtV1RegistryCredentialsIdPatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiAccountsMgmtV1RegistryCredentialsIdPatch`: RegistryCredential
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ApiAccountsMgmtV1RegistryCredentialsIdPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of record | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiAccountsMgmtV1RegistryCredentialsIdPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **registryCredentialPatchRequest** | [**RegistryCredentialPatchRequest**](RegistryCredentialPatchRequest.md) | Updated registry credential data | 

### Return type

[**RegistryCredential**](RegistryCredential.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiAccountsMgmtV1RegistryCredentialsPost

> RegistryCredential ApiAccountsMgmtV1RegistryCredentialsPost(ctx).RegistryCredential(registryCredential).Execute()

Request the creation of a registry credential

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    registryCredential := *openapiclient.NewRegistryCredential() // RegistryCredential | Registry credential data

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ApiAccountsMgmtV1RegistryCredentialsPost(context.Background()).RegistryCredential(registryCredential).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ApiAccountsMgmtV1RegistryCredentialsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiAccountsMgmtV1RegistryCredentialsPost`: RegistryCredential
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ApiAccountsMgmtV1RegistryCredentialsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiAccountsMgmtV1RegistryCredentialsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **registryCredential** | [**RegistryCredential**](RegistryCredential.md) | Registry credential data | 

### Return type

[**RegistryCredential**](RegistryCredential.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiAccountsMgmtV1ReservedResourcesGet

> ReservedResourceList ApiAccountsMgmtV1ReservedResourcesGet(ctx).Page(page).Size(size).Search(search).OrderBy(orderBy).Execute()

Returns a list of reserved resources

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    page := int32(56) // int32 | Page number of record list when record list exceeds specified page size (optional) (default to 1)
    size := int32(56) // int32 | Maximum number of records to return (optional) (default to 100)
    search := "search_example" // string | Specifies the search criteria. The syntax of this parameter is similar to the syntax of the _where_ clause of an SQL statement, using the names of the json attributes / column names of the account. For example, in order to retrieve all the accounts with a username starting with `my`:  ```sql username like 'my%' ```  The search criteria can also be applied on related resource. For example, in order to retrieve all the subscriptions labeled by `foo=bar`,  ```sql labels.key = 'foo' and labels.value = 'bar' ```  If the parameter isn't provided, or if the value is empty, then all the accounts that the user has permission to see will be returned. (optional)
    orderBy := "orderBy_example" // string | Specifies the order by criteria. The syntax of this parameter is similar to the syntax of the _order by_ clause of an SQL statement, but using the names of the json attributes / column of the account. For example, in order to retrieve all accounts ordered by username:  ```sql username asc ```  Or in order to retrieve all accounts ordered by username _and_ first name:  ```sql username asc, firstName asc ```  If the parameter isn't provided, or if the value is empty, then no explicit ordering will be applied. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ApiAccountsMgmtV1ReservedResourcesGet(context.Background()).Page(page).Size(size).Search(search).OrderBy(orderBy).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ApiAccountsMgmtV1ReservedResourcesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiAccountsMgmtV1ReservedResourcesGet`: ReservedResourceList
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ApiAccountsMgmtV1ReservedResourcesGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiAccountsMgmtV1ReservedResourcesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Page number of record list when record list exceeds specified page size | [default to 1]
 **size** | **int32** | Maximum number of records to return | [default to 100]
 **search** | **string** | Specifies the search criteria. The syntax of this parameter is similar to the syntax of the _where_ clause of an SQL statement, using the names of the json attributes / column names of the account. For example, in order to retrieve all the accounts with a username starting with &#x60;my&#x60;:  &#x60;&#x60;&#x60;sql username like &#39;my%&#39; &#x60;&#x60;&#x60;  The search criteria can also be applied on related resource. For example, in order to retrieve all the subscriptions labeled by &#x60;foo&#x3D;bar&#x60;,  &#x60;&#x60;&#x60;sql labels.key &#x3D; &#39;foo&#39; and labels.value &#x3D; &#39;bar&#39; &#x60;&#x60;&#x60;  If the parameter isn&#39;t provided, or if the value is empty, then all the accounts that the user has permission to see will be returned. | 
 **orderBy** | **string** | Specifies the order by criteria. The syntax of this parameter is similar to the syntax of the _order by_ clause of an SQL statement, but using the names of the json attributes / column of the account. For example, in order to retrieve all accounts ordered by username:  &#x60;&#x60;&#x60;sql username asc &#x60;&#x60;&#x60;  Or in order to retrieve all accounts ordered by username _and_ first name:  &#x60;&#x60;&#x60;sql username asc, firstName asc &#x60;&#x60;&#x60;  If the parameter isn&#39;t provided, or if the value is empty, then no explicit ordering will be applied. | 

### Return type

[**ReservedResourceList**](ReservedResourceList.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiAccountsMgmtV1ResourceQuotaGet

> ResourceQuotaList ApiAccountsMgmtV1ResourceQuotaGet(ctx).Page(page).Size(size).Search(search).OrderBy(orderBy).Execute()

Returns a list of resource quota objects

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    page := int32(56) // int32 | Page number of record list when record list exceeds specified page size (optional) (default to 1)
    size := int32(56) // int32 | Maximum number of records to return (optional) (default to 100)
    search := "search_example" // string | Specifies the search criteria. The syntax of this parameter is similar to the syntax of the _where_ clause of an SQL statement, using the names of the json attributes / column names of the account. For example, in order to retrieve all the accounts with a username starting with `my`:  ```sql username like 'my%' ```  The search criteria can also be applied on related resource. For example, in order to retrieve all the subscriptions labeled by `foo=bar`,  ```sql labels.key = 'foo' and labels.value = 'bar' ```  If the parameter isn't provided, or if the value is empty, then all the accounts that the user has permission to see will be returned. (optional)
    orderBy := "orderBy_example" // string | Specifies the order by criteria. The syntax of this parameter is similar to the syntax of the _order by_ clause of an SQL statement, but using the names of the json attributes / column of the account. For example, in order to retrieve all accounts ordered by username:  ```sql username asc ```  Or in order to retrieve all accounts ordered by username _and_ first name:  ```sql username asc, firstName asc ```  If the parameter isn't provided, or if the value is empty, then no explicit ordering will be applied. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ApiAccountsMgmtV1ResourceQuotaGet(context.Background()).Page(page).Size(size).Search(search).OrderBy(orderBy).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ApiAccountsMgmtV1ResourceQuotaGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiAccountsMgmtV1ResourceQuotaGet`: ResourceQuotaList
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ApiAccountsMgmtV1ResourceQuotaGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiAccountsMgmtV1ResourceQuotaGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Page number of record list when record list exceeds specified page size | [default to 1]
 **size** | **int32** | Maximum number of records to return | [default to 100]
 **search** | **string** | Specifies the search criteria. The syntax of this parameter is similar to the syntax of the _where_ clause of an SQL statement, using the names of the json attributes / column names of the account. For example, in order to retrieve all the accounts with a username starting with &#x60;my&#x60;:  &#x60;&#x60;&#x60;sql username like &#39;my%&#39; &#x60;&#x60;&#x60;  The search criteria can also be applied on related resource. For example, in order to retrieve all the subscriptions labeled by &#x60;foo&#x3D;bar&#x60;,  &#x60;&#x60;&#x60;sql labels.key &#x3D; &#39;foo&#39; and labels.value &#x3D; &#39;bar&#39; &#x60;&#x60;&#x60;  If the parameter isn&#39;t provided, or if the value is empty, then all the accounts that the user has permission to see will be returned. | 
 **orderBy** | **string** | Specifies the order by criteria. The syntax of this parameter is similar to the syntax of the _order by_ clause of an SQL statement, but using the names of the json attributes / column of the account. For example, in order to retrieve all accounts ordered by username:  &#x60;&#x60;&#x60;sql username asc &#x60;&#x60;&#x60;  Or in order to retrieve all accounts ordered by username _and_ first name:  &#x60;&#x60;&#x60;sql username asc, firstName asc &#x60;&#x60;&#x60;  If the parameter isn&#39;t provided, or if the value is empty, then no explicit ordering will be applied. | 

### Return type

[**ResourceQuotaList**](ResourceQuotaList.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiAccountsMgmtV1RoleBindingsGet

> RoleBindingList ApiAccountsMgmtV1RoleBindingsGet(ctx).Page(page).Size(size).Search(search).OrderBy(orderBy).Execute()

Returns a list of role bindings

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    page := int32(56) // int32 | Page number of record list when record list exceeds specified page size (optional) (default to 1)
    size := int32(56) // int32 | Maximum number of records to return (optional) (default to 100)
    search := "search_example" // string | Specifies the search criteria. The syntax of this parameter is similar to the syntax of the _where_ clause of an SQL statement, using the names of the json attributes / column names of the account. For example, in order to retrieve all the accounts with a username starting with `my`:  ```sql username like 'my%' ```  The search criteria can also be applied on related resource. For example, in order to retrieve all the subscriptions labeled by `foo=bar`,  ```sql labels.key = 'foo' and labels.value = 'bar' ```  If the parameter isn't provided, or if the value is empty, then all the accounts that the user has permission to see will be returned. (optional)
    orderBy := "orderBy_example" // string | Specifies the order by criteria. The syntax of this parameter is similar to the syntax of the _order by_ clause of an SQL statement, but using the names of the json attributes / column of the account. For example, in order to retrieve all accounts ordered by username:  ```sql username asc ```  Or in order to retrieve all accounts ordered by username _and_ first name:  ```sql username asc, firstName asc ```  If the parameter isn't provided, or if the value is empty, then no explicit ordering will be applied. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ApiAccountsMgmtV1RoleBindingsGet(context.Background()).Page(page).Size(size).Search(search).OrderBy(orderBy).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ApiAccountsMgmtV1RoleBindingsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiAccountsMgmtV1RoleBindingsGet`: RoleBindingList
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ApiAccountsMgmtV1RoleBindingsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiAccountsMgmtV1RoleBindingsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Page number of record list when record list exceeds specified page size | [default to 1]
 **size** | **int32** | Maximum number of records to return | [default to 100]
 **search** | **string** | Specifies the search criteria. The syntax of this parameter is similar to the syntax of the _where_ clause of an SQL statement, using the names of the json attributes / column names of the account. For example, in order to retrieve all the accounts with a username starting with &#x60;my&#x60;:  &#x60;&#x60;&#x60;sql username like &#39;my%&#39; &#x60;&#x60;&#x60;  The search criteria can also be applied on related resource. For example, in order to retrieve all the subscriptions labeled by &#x60;foo&#x3D;bar&#x60;,  &#x60;&#x60;&#x60;sql labels.key &#x3D; &#39;foo&#39; and labels.value &#x3D; &#39;bar&#39; &#x60;&#x60;&#x60;  If the parameter isn&#39;t provided, or if the value is empty, then all the accounts that the user has permission to see will be returned. | 
 **orderBy** | **string** | Specifies the order by criteria. The syntax of this parameter is similar to the syntax of the _order by_ clause of an SQL statement, but using the names of the json attributes / column of the account. For example, in order to retrieve all accounts ordered by username:  &#x60;&#x60;&#x60;sql username asc &#x60;&#x60;&#x60;  Or in order to retrieve all accounts ordered by username _and_ first name:  &#x60;&#x60;&#x60;sql username asc, firstName asc &#x60;&#x60;&#x60;  If the parameter isn&#39;t provided, or if the value is empty, then no explicit ordering will be applied. | 

### Return type

[**RoleBindingList**](RoleBindingList.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiAccountsMgmtV1RoleBindingsIdDelete

> ApiAccountsMgmtV1RoleBindingsIdDelete(ctx, id).Execute()

Delete a role binding

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "id_example" // string | The id of record

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ApiAccountsMgmtV1RoleBindingsIdDelete(context.Background(), id).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ApiAccountsMgmtV1RoleBindingsIdDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of record | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiAccountsMgmtV1RoleBindingsIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiAccountsMgmtV1RoleBindingsIdGet

> RoleBinding ApiAccountsMgmtV1RoleBindingsIdGet(ctx, id).Execute()

Get a role binding

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "id_example" // string | The id of record

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ApiAccountsMgmtV1RoleBindingsIdGet(context.Background(), id).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ApiAccountsMgmtV1RoleBindingsIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiAccountsMgmtV1RoleBindingsIdGet`: RoleBinding
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ApiAccountsMgmtV1RoleBindingsIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of record | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiAccountsMgmtV1RoleBindingsIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RoleBinding**](RoleBinding.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiAccountsMgmtV1RoleBindingsIdPatch

> RoleBinding ApiAccountsMgmtV1RoleBindingsIdPatch(ctx, id).RoleBindingRequest(roleBindingRequest).Execute()

Update a role binding

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "id_example" // string | The id of record
    roleBindingRequest := *openapiclient.NewRoleBindingRequest() // RoleBindingRequest | Updated role binding data

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ApiAccountsMgmtV1RoleBindingsIdPatch(context.Background(), id).RoleBindingRequest(roleBindingRequest).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ApiAccountsMgmtV1RoleBindingsIdPatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiAccountsMgmtV1RoleBindingsIdPatch`: RoleBinding
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ApiAccountsMgmtV1RoleBindingsIdPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of record | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiAccountsMgmtV1RoleBindingsIdPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **roleBindingRequest** | [**RoleBindingRequest**](RoleBindingRequest.md) | Updated role binding data | 

### Return type

[**RoleBinding**](RoleBinding.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiAccountsMgmtV1RoleBindingsPost

> RoleBinding ApiAccountsMgmtV1RoleBindingsPost(ctx).RoleBindingCreateRequest(roleBindingCreateRequest).Execute()

Create a new role binding

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    roleBindingCreateRequest := *openapiclient.NewRoleBindingCreateRequest("AccountId_example", "RoleId_example", "Type_example") // RoleBindingCreateRequest | Role binding data

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ApiAccountsMgmtV1RoleBindingsPost(context.Background()).RoleBindingCreateRequest(roleBindingCreateRequest).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ApiAccountsMgmtV1RoleBindingsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiAccountsMgmtV1RoleBindingsPost`: RoleBinding
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ApiAccountsMgmtV1RoleBindingsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiAccountsMgmtV1RoleBindingsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **roleBindingCreateRequest** | [**RoleBindingCreateRequest**](RoleBindingCreateRequest.md) | Role binding data | 

### Return type

[**RoleBinding**](RoleBinding.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiAccountsMgmtV1RolesGet

> RoleList ApiAccountsMgmtV1RolesGet(ctx).Page(page).Size(size).Search(search).Execute()

Returns a list of roles

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    page := int32(56) // int32 | Page number of record list when record list exceeds specified page size (optional) (default to 1)
    size := int32(56) // int32 | Maximum number of records to return (optional) (default to 100)
    search := "search_example" // string | Specifies the search criteria. The syntax of this parameter is similar to the syntax of the _where_ clause of an SQL statement, using the names of the json attributes / column names of the account. For example, in order to retrieve all the accounts with a username starting with `my`:  ```sql username like 'my%' ```  The search criteria can also be applied on related resource. For example, in order to retrieve all the subscriptions labeled by `foo=bar`,  ```sql labels.key = 'foo' and labels.value = 'bar' ```  If the parameter isn't provided, or if the value is empty, then all the accounts that the user has permission to see will be returned. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ApiAccountsMgmtV1RolesGet(context.Background()).Page(page).Size(size).Search(search).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ApiAccountsMgmtV1RolesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiAccountsMgmtV1RolesGet`: RoleList
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ApiAccountsMgmtV1RolesGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiAccountsMgmtV1RolesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Page number of record list when record list exceeds specified page size | [default to 1]
 **size** | **int32** | Maximum number of records to return | [default to 100]
 **search** | **string** | Specifies the search criteria. The syntax of this parameter is similar to the syntax of the _where_ clause of an SQL statement, using the names of the json attributes / column names of the account. For example, in order to retrieve all the accounts with a username starting with &#x60;my&#x60;:  &#x60;&#x60;&#x60;sql username like &#39;my%&#39; &#x60;&#x60;&#x60;  The search criteria can also be applied on related resource. For example, in order to retrieve all the subscriptions labeled by &#x60;foo&#x3D;bar&#x60;,  &#x60;&#x60;&#x60;sql labels.key &#x3D; &#39;foo&#39; and labels.value &#x3D; &#39;bar&#39; &#x60;&#x60;&#x60;  If the parameter isn&#39;t provided, or if the value is empty, then all the accounts that the user has permission to see will be returned. | 

### Return type

[**RoleList**](RoleList.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiAccountsMgmtV1RolesIdGet

> Role ApiAccountsMgmtV1RolesIdGet(ctx, id).Execute()

Get a role by id

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "id_example" // string | The id of record

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ApiAccountsMgmtV1RolesIdGet(context.Background(), id).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ApiAccountsMgmtV1RolesIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiAccountsMgmtV1RolesIdGet`: Role
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ApiAccountsMgmtV1RolesIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of record | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiAccountsMgmtV1RolesIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Role**](Role.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiAccountsMgmtV1SkuRulesGet

> SkuRulesList ApiAccountsMgmtV1SkuRulesGet(ctx).Search(search).Execute()

Returns a list of UHC product SKU Rules

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    search := "search_example" // string | Specifies the search criteria. The syntax of this parameter is similar to the syntax of the _where_ clause of an SQL statement, using the names of the json attributes / column names of the account. For example, in order to retrieve all the accounts with a username starting with `my`:  ```sql username like 'my%' ```  The search criteria can also be applied on related resource. For example, in order to retrieve all the subscriptions labeled by `foo=bar`,  ```sql labels.key = 'foo' and labels.value = 'bar' ```  If the parameter isn't provided, or if the value is empty, then all the accounts that the user has permission to see will be returned. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ApiAccountsMgmtV1SkuRulesGet(context.Background()).Search(search).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ApiAccountsMgmtV1SkuRulesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiAccountsMgmtV1SkuRulesGet`: SkuRulesList
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ApiAccountsMgmtV1SkuRulesGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiAccountsMgmtV1SkuRulesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **search** | **string** | Specifies the search criteria. The syntax of this parameter is similar to the syntax of the _where_ clause of an SQL statement, using the names of the json attributes / column names of the account. For example, in order to retrieve all the accounts with a username starting with &#x60;my&#x60;:  &#x60;&#x60;&#x60;sql username like &#39;my%&#39; &#x60;&#x60;&#x60;  The search criteria can also be applied on related resource. For example, in order to retrieve all the subscriptions labeled by &#x60;foo&#x3D;bar&#x60;,  &#x60;&#x60;&#x60;sql labels.key &#x3D; &#39;foo&#39; and labels.value &#x3D; &#39;bar&#39; &#x60;&#x60;&#x60;  If the parameter isn&#39;t provided, or if the value is empty, then all the accounts that the user has permission to see will be returned. | 

### Return type

[**SkuRulesList**](SkuRulesList.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiAccountsMgmtV1SkuRulesIdGet

> SkuRules ApiAccountsMgmtV1SkuRulesIdGet(ctx, id).Execute()

Get a sku rules by id

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "id_example" // string | The id of record

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ApiAccountsMgmtV1SkuRulesIdGet(context.Background(), id).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ApiAccountsMgmtV1SkuRulesIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiAccountsMgmtV1SkuRulesIdGet`: SkuRules
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ApiAccountsMgmtV1SkuRulesIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of record | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiAccountsMgmtV1SkuRulesIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SkuRules**](SkuRules.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiAccountsMgmtV1SkusGet

> SkuList ApiAccountsMgmtV1SkusGet(ctx).Search(search).Execute()

Returns a list of UHC product SKUs

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    search := "search_example" // string | Specifies the search criteria. The syntax of this parameter is similar to the syntax of the _where_ clause of an SQL statement, using the names of the json attributes / column names of the account. For example, in order to retrieve all the accounts with a username starting with `my`:  ```sql username like 'my%' ```  The search criteria can also be applied on related resource. For example, in order to retrieve all the subscriptions labeled by `foo=bar`,  ```sql labels.key = 'foo' and labels.value = 'bar' ```  If the parameter isn't provided, or if the value is empty, then all the accounts that the user has permission to see will be returned. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ApiAccountsMgmtV1SkusGet(context.Background()).Search(search).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ApiAccountsMgmtV1SkusGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiAccountsMgmtV1SkusGet`: SkuList
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ApiAccountsMgmtV1SkusGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiAccountsMgmtV1SkusGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **search** | **string** | Specifies the search criteria. The syntax of this parameter is similar to the syntax of the _where_ clause of an SQL statement, using the names of the json attributes / column names of the account. For example, in order to retrieve all the accounts with a username starting with &#x60;my&#x60;:  &#x60;&#x60;&#x60;sql username like &#39;my%&#39; &#x60;&#x60;&#x60;  The search criteria can also be applied on related resource. For example, in order to retrieve all the subscriptions labeled by &#x60;foo&#x3D;bar&#x60;,  &#x60;&#x60;&#x60;sql labels.key &#x3D; &#39;foo&#39; and labels.value &#x3D; &#39;bar&#39; &#x60;&#x60;&#x60;  If the parameter isn&#39;t provided, or if the value is empty, then all the accounts that the user has permission to see will be returned. | 

### Return type

[**SkuList**](SkuList.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiAccountsMgmtV1SkusIdGet

> SKU ApiAccountsMgmtV1SkusIdGet(ctx, id).Execute()

Get a sku by id

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "id_example" // string | The id of record

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ApiAccountsMgmtV1SkusIdGet(context.Background(), id).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ApiAccountsMgmtV1SkusIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiAccountsMgmtV1SkusIdGet`: SKU
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ApiAccountsMgmtV1SkusIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of record | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiAccountsMgmtV1SkusIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SKU**](SKU.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiAccountsMgmtV1SubscriptionsGet

> SubscriptionList ApiAccountsMgmtV1SubscriptionsGet(ctx).Page(page).Size(size).Search(search).FetchAccounts(fetchAccounts).FetchLabels(fetchLabels).FetchCapabilities(fetchCapabilities).Fields(fields).OrderBy(orderBy).Labels(labels).Execute()

Returns a list of subscriptions

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    page := int32(56) // int32 | Page number of record list when record list exceeds specified page size (optional) (default to 1)
    size := int32(56) // int32 | Maximum number of records to return (optional) (default to 100)
    search := "search_example" // string | Specifies the search criteria. The syntax of this parameter is similar to the syntax of the _where_ clause of an SQL statement, using the names of the json attributes / column names of the account. For example, in order to retrieve all the accounts with a username starting with `my`:  ```sql username like 'my%' ```  The search criteria can also be applied on related resource. For example, in order to retrieve all the subscriptions labeled by `foo=bar`,  ```sql labels.key = 'foo' and labels.value = 'bar' ```  If the parameter isn't provided, or if the value is empty, then all the accounts that the user has permission to see will be returned. (optional)
    fetchAccounts := true // bool | If true, includes the account reference information in the output. Could slow request response time. (optional)
    fetchLabels := true // bool | If true, includes the labels on a subscription/organization/account in the output. Could slow request response time. (optional)
    fetchCapabilities := true // bool | If true, includes the capabilities on a subscription in the output. Could slow request response time. (optional)
    fields := "fields_example" // string | Supplies a comma-separated list of fields to be returned. Fields of sub-structures and of arrays use <structure>.<field> notation. <stucture>.* means all field of a structure Example: For each Subscription to get id, href, plan(id and kind) and labels (all fields)  ``` ocm get subscriptions --parameter fields=id,href,plan.id,plan.kind,labels.* --parameter fetchLabels=true ``` (optional)
    orderBy := "orderBy_example" // string | Specifies the order by criteria. The syntax of this parameter is similar to the syntax of the _order by_ clause of an SQL statement, but using the names of the json attributes / column of the account. For example, in order to retrieve all accounts ordered by username:  ```sql username asc ```  Or in order to retrieve all accounts ordered by username _and_ first name:  ```sql username asc, firstName asc ```  If the parameter isn't provided, or if the value is empty, then no explicit ordering will be applied. (optional)
    labels := "labels_example" // string | Specifies the criteria to filter the subscription resource based on their labels. A label is represented as a `key=value` pair,  ``` labels = \"foo=bar\" ```  and multiple labels are separated by comma,  ``` labels = \"foo=bar,fooz=barz\" ``` (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ApiAccountsMgmtV1SubscriptionsGet(context.Background()).Page(page).Size(size).Search(search).FetchAccounts(fetchAccounts).FetchLabels(fetchLabels).FetchCapabilities(fetchCapabilities).Fields(fields).OrderBy(orderBy).Labels(labels).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ApiAccountsMgmtV1SubscriptionsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiAccountsMgmtV1SubscriptionsGet`: SubscriptionList
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ApiAccountsMgmtV1SubscriptionsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiAccountsMgmtV1SubscriptionsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Page number of record list when record list exceeds specified page size | [default to 1]
 **size** | **int32** | Maximum number of records to return | [default to 100]
 **search** | **string** | Specifies the search criteria. The syntax of this parameter is similar to the syntax of the _where_ clause of an SQL statement, using the names of the json attributes / column names of the account. For example, in order to retrieve all the accounts with a username starting with &#x60;my&#x60;:  &#x60;&#x60;&#x60;sql username like &#39;my%&#39; &#x60;&#x60;&#x60;  The search criteria can also be applied on related resource. For example, in order to retrieve all the subscriptions labeled by &#x60;foo&#x3D;bar&#x60;,  &#x60;&#x60;&#x60;sql labels.key &#x3D; &#39;foo&#39; and labels.value &#x3D; &#39;bar&#39; &#x60;&#x60;&#x60;  If the parameter isn&#39;t provided, or if the value is empty, then all the accounts that the user has permission to see will be returned. | 
 **fetchAccounts** | **bool** | If true, includes the account reference information in the output. Could slow request response time. | 
 **fetchLabels** | **bool** | If true, includes the labels on a subscription/organization/account in the output. Could slow request response time. | 
 **fetchCapabilities** | **bool** | If true, includes the capabilities on a subscription in the output. Could slow request response time. | 
 **fields** | **string** | Supplies a comma-separated list of fields to be returned. Fields of sub-structures and of arrays use &lt;structure&gt;.&lt;field&gt; notation. &lt;stucture&gt;.* means all field of a structure Example: For each Subscription to get id, href, plan(id and kind) and labels (all fields)  &#x60;&#x60;&#x60; ocm get subscriptions --parameter fields&#x3D;id,href,plan.id,plan.kind,labels.* --parameter fetchLabels&#x3D;true &#x60;&#x60;&#x60; | 
 **orderBy** | **string** | Specifies the order by criteria. The syntax of this parameter is similar to the syntax of the _order by_ clause of an SQL statement, but using the names of the json attributes / column of the account. For example, in order to retrieve all accounts ordered by username:  &#x60;&#x60;&#x60;sql username asc &#x60;&#x60;&#x60;  Or in order to retrieve all accounts ordered by username _and_ first name:  &#x60;&#x60;&#x60;sql username asc, firstName asc &#x60;&#x60;&#x60;  If the parameter isn&#39;t provided, or if the value is empty, then no explicit ordering will be applied. | 
 **labels** | **string** | Specifies the criteria to filter the subscription resource based on their labels. A label is represented as a &#x60;key&#x3D;value&#x60; pair,  &#x60;&#x60;&#x60; labels &#x3D; \&quot;foo&#x3D;bar\&quot; &#x60;&#x60;&#x60;  and multiple labels are separated by comma,  &#x60;&#x60;&#x60; labels &#x3D; \&quot;foo&#x3D;bar,fooz&#x3D;barz\&quot; &#x60;&#x60;&#x60; | 

### Return type

[**SubscriptionList**](SubscriptionList.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiAccountsMgmtV1SubscriptionsIdDelete

> ApiAccountsMgmtV1SubscriptionsIdDelete(ctx, id).Execute()

Deletes a subscription by id

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "id_example" // string | The id of record

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ApiAccountsMgmtV1SubscriptionsIdDelete(context.Background(), id).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ApiAccountsMgmtV1SubscriptionsIdDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of record | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiAccountsMgmtV1SubscriptionsIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiAccountsMgmtV1SubscriptionsIdGet

> Subscription ApiAccountsMgmtV1SubscriptionsIdGet(ctx, id).FetchAccounts(fetchAccounts).FetchLabels(fetchLabels).FetchCapabilities(fetchCapabilities).FetchCpuAndSocket(fetchCpuAndSocket).Execute()

Get a subscription by id

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "id_example" // string | The id of record
    fetchAccounts := true // bool | If true, includes the account reference information in the output. Could slow request response time. (optional)
    fetchLabels := true // bool | If true, includes the labels on a subscription/organization/account in the output. Could slow request response time. (optional)
    fetchCapabilities := true // bool | If true, includes the capabilities on a subscription in the output. Could slow request response time. (optional)
    fetchCpuAndSocket := true // bool | If true, fetches, from the clusters service, the total numbers of CPU's and sockets under an obligation, and includes in the output. Could slow request response time. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ApiAccountsMgmtV1SubscriptionsIdGet(context.Background(), id).FetchAccounts(fetchAccounts).FetchLabels(fetchLabels).FetchCapabilities(fetchCapabilities).FetchCpuAndSocket(fetchCpuAndSocket).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ApiAccountsMgmtV1SubscriptionsIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiAccountsMgmtV1SubscriptionsIdGet`: Subscription
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ApiAccountsMgmtV1SubscriptionsIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of record | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiAccountsMgmtV1SubscriptionsIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fetchAccounts** | **bool** | If true, includes the account reference information in the output. Could slow request response time. | 
 **fetchLabels** | **bool** | If true, includes the labels on a subscription/organization/account in the output. Could slow request response time. | 
 **fetchCapabilities** | **bool** | If true, includes the capabilities on a subscription in the output. Could slow request response time. | 
 **fetchCpuAndSocket** | **bool** | If true, fetches, from the clusters service, the total numbers of CPU&#39;s and sockets under an obligation, and includes in the output. Could slow request response time. | 

### Return type

[**Subscription**](Subscription.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiAccountsMgmtV1SubscriptionsIdLabelsGet

> LabelList ApiAccountsMgmtV1SubscriptionsIdLabelsGet(ctx, id).Page(page).Size(size).Search(search).OrderBy(orderBy).Execute()

Returns a list of labels

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "id_example" // string | The id of record
    page := int32(56) // int32 | Page number of record list when record list exceeds specified page size (optional) (default to 1)
    size := int32(56) // int32 | Maximum number of records to return (optional) (default to 100)
    search := "search_example" // string | Specifies the search criteria. The syntax of this parameter is similar to the syntax of the _where_ clause of an SQL statement, using the names of the json attributes / column names of the account. For example, in order to retrieve all the accounts with a username starting with `my`:  ```sql username like 'my%' ```  The search criteria can also be applied on related resource. For example, in order to retrieve all the subscriptions labeled by `foo=bar`,  ```sql labels.key = 'foo' and labels.value = 'bar' ```  If the parameter isn't provided, or if the value is empty, then all the accounts that the user has permission to see will be returned. (optional)
    orderBy := "orderBy_example" // string | Specifies the order by criteria. The syntax of this parameter is similar to the syntax of the _order by_ clause of an SQL statement, but using the names of the json attributes / column of the account. For example, in order to retrieve all accounts ordered by username:  ```sql username asc ```  Or in order to retrieve all accounts ordered by username _and_ first name:  ```sql username asc, firstName asc ```  If the parameter isn't provided, or if the value is empty, then no explicit ordering will be applied. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ApiAccountsMgmtV1SubscriptionsIdLabelsGet(context.Background(), id).Page(page).Size(size).Search(search).OrderBy(orderBy).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ApiAccountsMgmtV1SubscriptionsIdLabelsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiAccountsMgmtV1SubscriptionsIdLabelsGet`: LabelList
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ApiAccountsMgmtV1SubscriptionsIdLabelsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of record | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiAccountsMgmtV1SubscriptionsIdLabelsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | Page number of record list when record list exceeds specified page size | [default to 1]
 **size** | **int32** | Maximum number of records to return | [default to 100]
 **search** | **string** | Specifies the search criteria. The syntax of this parameter is similar to the syntax of the _where_ clause of an SQL statement, using the names of the json attributes / column names of the account. For example, in order to retrieve all the accounts with a username starting with &#x60;my&#x60;:  &#x60;&#x60;&#x60;sql username like &#39;my%&#39; &#x60;&#x60;&#x60;  The search criteria can also be applied on related resource. For example, in order to retrieve all the subscriptions labeled by &#x60;foo&#x3D;bar&#x60;,  &#x60;&#x60;&#x60;sql labels.key &#x3D; &#39;foo&#39; and labels.value &#x3D; &#39;bar&#39; &#x60;&#x60;&#x60;  If the parameter isn&#39;t provided, or if the value is empty, then all the accounts that the user has permission to see will be returned. | 
 **orderBy** | **string** | Specifies the order by criteria. The syntax of this parameter is similar to the syntax of the _order by_ clause of an SQL statement, but using the names of the json attributes / column of the account. For example, in order to retrieve all accounts ordered by username:  &#x60;&#x60;&#x60;sql username asc &#x60;&#x60;&#x60;  Or in order to retrieve all accounts ordered by username _and_ first name:  &#x60;&#x60;&#x60;sql username asc, firstName asc &#x60;&#x60;&#x60;  If the parameter isn&#39;t provided, or if the value is empty, then no explicit ordering will be applied. | 

### Return type

[**LabelList**](LabelList.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiAccountsMgmtV1SubscriptionsIdLabelsKeyDelete

> ApiAccountsMgmtV1SubscriptionsIdLabelsKeyDelete(ctx, id, key).Execute()

Delete a label

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "id_example" // string | The id of record
    key := "key_example" // string | The key of the label

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ApiAccountsMgmtV1SubscriptionsIdLabelsKeyDelete(context.Background(), id, key).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ApiAccountsMgmtV1SubscriptionsIdLabelsKeyDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of record | 
**key** | **string** | The key of the label | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiAccountsMgmtV1SubscriptionsIdLabelsKeyDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiAccountsMgmtV1SubscriptionsIdLabelsKeyGet

> Label ApiAccountsMgmtV1SubscriptionsIdLabelsKeyGet(ctx, id, key).Execute()

Get subscription labels by label key

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "id_example" // string | The id of record
    key := "key_example" // string | The key of the label

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ApiAccountsMgmtV1SubscriptionsIdLabelsKeyGet(context.Background(), id, key).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ApiAccountsMgmtV1SubscriptionsIdLabelsKeyGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiAccountsMgmtV1SubscriptionsIdLabelsKeyGet`: Label
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ApiAccountsMgmtV1SubscriptionsIdLabelsKeyGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of record | 
**key** | **string** | The key of the label | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiAccountsMgmtV1SubscriptionsIdLabelsKeyGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Label**](Label.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiAccountsMgmtV1SubscriptionsIdLabelsKeyPatch

> Label ApiAccountsMgmtV1SubscriptionsIdLabelsKeyPatch(ctx, id, key).Label(label).Execute()

Create a new label or update an existing label

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "id_example" // string | The id of record
    key := "key_example" // string | The key of the label
    label := *openapiclient.NewLabel(false, "Key_example", "Value_example") // Label | Label data

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ApiAccountsMgmtV1SubscriptionsIdLabelsKeyPatch(context.Background(), id, key).Label(label).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ApiAccountsMgmtV1SubscriptionsIdLabelsKeyPatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiAccountsMgmtV1SubscriptionsIdLabelsKeyPatch`: Label
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ApiAccountsMgmtV1SubscriptionsIdLabelsKeyPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of record | 
**key** | **string** | The key of the label | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiAccountsMgmtV1SubscriptionsIdLabelsKeyPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **label** | [**Label**](Label.md) | Label data | 

### Return type

[**Label**](Label.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiAccountsMgmtV1SubscriptionsIdLabelsPost

> Label ApiAccountsMgmtV1SubscriptionsIdLabelsPost(ctx, id).Label(label).Execute()

Create a new label or update an existing label

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "id_example" // string | The id of record
    label := *openapiclient.NewLabel(false, "Key_example", "Value_example") // Label | Label data

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ApiAccountsMgmtV1SubscriptionsIdLabelsPost(context.Background(), id).Label(label).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ApiAccountsMgmtV1SubscriptionsIdLabelsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiAccountsMgmtV1SubscriptionsIdLabelsPost`: Label
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ApiAccountsMgmtV1SubscriptionsIdLabelsPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of record | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiAccountsMgmtV1SubscriptionsIdLabelsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **label** | [**Label**](Label.md) | Label data | 

### Return type

[**Label**](Label.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiAccountsMgmtV1SubscriptionsIdNotifyPost

> ApiAccountsMgmtV1SubscriptionsIdNotifyPost(ctx, id).NotificationRequest(notificationRequest).Execute()

Notify the owner of a subscription

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "id_example" // string | The id of record
    notificationRequest := *openapiclient.NewNotificationRequest("Subject_example", "TemplateName_example") // NotificationRequest | The contents of the notification to send to the owner of a subscription in addition to the set of template parameters which are sent automatically ACCOUNT_USERNAME, FIRST_NAME, LAST_NAME, ORGANIZATION_NAME, ORGANIZATION_EXTERNAL_ID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ApiAccountsMgmtV1SubscriptionsIdNotifyPost(context.Background(), id).NotificationRequest(notificationRequest).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ApiAccountsMgmtV1SubscriptionsIdNotifyPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of record | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiAccountsMgmtV1SubscriptionsIdNotifyPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **notificationRequest** | [**NotificationRequest**](NotificationRequest.md) | The contents of the notification to send to the owner of a subscription in addition to the set of template parameters which are sent automatically ACCOUNT_USERNAME, FIRST_NAME, LAST_NAME, ORGANIZATION_NAME, ORGANIZATION_EXTERNAL_ID | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiAccountsMgmtV1SubscriptionsIdPatch

> Subscription ApiAccountsMgmtV1SubscriptionsIdPatch(ctx, id).SubscriptionPatchRequest(subscriptionPatchRequest).Execute()

Update a subscription

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "id_example" // string | The id of record
    subscriptionPatchRequest := *openapiclient.NewSubscriptionPatchRequest() // SubscriptionPatchRequest | Updated subscription data

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ApiAccountsMgmtV1SubscriptionsIdPatch(context.Background(), id).SubscriptionPatchRequest(subscriptionPatchRequest).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ApiAccountsMgmtV1SubscriptionsIdPatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiAccountsMgmtV1SubscriptionsIdPatch`: Subscription
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ApiAccountsMgmtV1SubscriptionsIdPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of record | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiAccountsMgmtV1SubscriptionsIdPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **subscriptionPatchRequest** | [**SubscriptionPatchRequest**](SubscriptionPatchRequest.md) | Updated subscription data | 

### Return type

[**Subscription**](Subscription.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiAccountsMgmtV1SubscriptionsIdReservedResourcesGet

> ReservedResourceList ApiAccountsMgmtV1SubscriptionsIdReservedResourcesGet(ctx, id).Page(page).Size(size).Search(search).OrderBy(orderBy).Execute()

Returns a list of reserved resources

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "id_example" // string | The id of record
    page := int32(56) // int32 | Page number of record list when record list exceeds specified page size (optional) (default to 1)
    size := int32(56) // int32 | Maximum number of records to return (optional) (default to 100)
    search := "search_example" // string | Specifies the search criteria. The syntax of this parameter is similar to the syntax of the _where_ clause of an SQL statement, using the names of the json attributes / column names of the account. For example, in order to retrieve all the accounts with a username starting with `my`:  ```sql username like 'my%' ```  The search criteria can also be applied on related resource. For example, in order to retrieve all the subscriptions labeled by `foo=bar`,  ```sql labels.key = 'foo' and labels.value = 'bar' ```  If the parameter isn't provided, or if the value is empty, then all the accounts that the user has permission to see will be returned. (optional)
    orderBy := "orderBy_example" // string | Specifies the order by criteria. The syntax of this parameter is similar to the syntax of the _order by_ clause of an SQL statement, but using the names of the json attributes / column of the account. For example, in order to retrieve all accounts ordered by username:  ```sql username asc ```  Or in order to retrieve all accounts ordered by username _and_ first name:  ```sql username asc, firstName asc ```  If the parameter isn't provided, or if the value is empty, then no explicit ordering will be applied. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ApiAccountsMgmtV1SubscriptionsIdReservedResourcesGet(context.Background(), id).Page(page).Size(size).Search(search).OrderBy(orderBy).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ApiAccountsMgmtV1SubscriptionsIdReservedResourcesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiAccountsMgmtV1SubscriptionsIdReservedResourcesGet`: ReservedResourceList
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ApiAccountsMgmtV1SubscriptionsIdReservedResourcesGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of record | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiAccountsMgmtV1SubscriptionsIdReservedResourcesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | Page number of record list when record list exceeds specified page size | [default to 1]
 **size** | **int32** | Maximum number of records to return | [default to 100]
 **search** | **string** | Specifies the search criteria. The syntax of this parameter is similar to the syntax of the _where_ clause of an SQL statement, using the names of the json attributes / column names of the account. For example, in order to retrieve all the accounts with a username starting with &#x60;my&#x60;:  &#x60;&#x60;&#x60;sql username like &#39;my%&#39; &#x60;&#x60;&#x60;  The search criteria can also be applied on related resource. For example, in order to retrieve all the subscriptions labeled by &#x60;foo&#x3D;bar&#x60;,  &#x60;&#x60;&#x60;sql labels.key &#x3D; &#39;foo&#39; and labels.value &#x3D; &#39;bar&#39; &#x60;&#x60;&#x60;  If the parameter isn&#39;t provided, or if the value is empty, then all the accounts that the user has permission to see will be returned. | 
 **orderBy** | **string** | Specifies the order by criteria. The syntax of this parameter is similar to the syntax of the _order by_ clause of an SQL statement, but using the names of the json attributes / column of the account. For example, in order to retrieve all accounts ordered by username:  &#x60;&#x60;&#x60;sql username asc &#x60;&#x60;&#x60;  Or in order to retrieve all accounts ordered by username _and_ first name:  &#x60;&#x60;&#x60;sql username asc, firstName asc &#x60;&#x60;&#x60;  If the parameter isn&#39;t provided, or if the value is empty, then no explicit ordering will be applied. | 

### Return type

[**ReservedResourceList**](ReservedResourceList.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiAccountsMgmtV1SubscriptionsIdSupportCasesGet

> ApiAccountsMgmtV1SubscriptionsIdSupportCasesGet(ctx, id).Page(page).Size(size).Execute()

Returns a list of open support creates opened against the external cluster id of this subscrption

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "id_example" // string | The id of record
    page := int32(56) // int32 | Page number of record list when record list exceeds specified page size (optional) (default to 1)
    size := int32(56) // int32 | Maximum number of records to return (optional) (default to 100)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ApiAccountsMgmtV1SubscriptionsIdSupportCasesGet(context.Background(), id).Page(page).Size(size).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ApiAccountsMgmtV1SubscriptionsIdSupportCasesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of record | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiAccountsMgmtV1SubscriptionsIdSupportCasesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | Page number of record list when record list exceeds specified page size | [default to 1]
 **size** | **int32** | Maximum number of records to return | [default to 100]

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiAccountsMgmtV1SubscriptionsPost

> Subscription ApiAccountsMgmtV1SubscriptionsPost(ctx).SubscriptionCreateRequest(subscriptionCreateRequest).Execute()

Create a new subscription

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    subscriptionCreateRequest := *openapiclient.NewSubscriptionCreateRequest("ClusterUuid_example", "PlanId_example", "Status_example") // SubscriptionCreateRequest | Subscription Creation data

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ApiAccountsMgmtV1SubscriptionsPost(context.Background()).SubscriptionCreateRequest(subscriptionCreateRequest).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ApiAccountsMgmtV1SubscriptionsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiAccountsMgmtV1SubscriptionsPost`: Subscription
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ApiAccountsMgmtV1SubscriptionsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiAccountsMgmtV1SubscriptionsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **subscriptionCreateRequest** | [**SubscriptionCreateRequest**](SubscriptionCreateRequest.md) | Subscription Creation data | 

### Return type

[**Subscription**](Subscription.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiAccountsMgmtV1SubscriptionsSubIdNotificationContactsAccountIdDelete

> ApiAccountsMgmtV1SubscriptionsSubIdNotificationContactsAccountIdDelete(ctx, subId, accountId).Execute()

Deletes a notification contact by subscription and account id

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    subId := "subId_example" // string | The id of subscription
    accountId := "accountId_example" // string | The id of account

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ApiAccountsMgmtV1SubscriptionsSubIdNotificationContactsAccountIdDelete(context.Background(), subId, accountId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ApiAccountsMgmtV1SubscriptionsSubIdNotificationContactsAccountIdDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subId** | **string** | The id of subscription | 
**accountId** | **string** | The id of account | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiAccountsMgmtV1SubscriptionsSubIdNotificationContactsAccountIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiAccountsMgmtV1SubscriptionsSubIdNotificationContactsGet

> AccountList ApiAccountsMgmtV1SubscriptionsSubIdNotificationContactsGet(ctx, subId).Page(page).Size(size).Search(search).Fields(fields).OrderBy(orderBy).Execute()

Returns a list of notification contacts for the given subscription

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    subId := "subId_example" // string | The id of subscription
    page := int32(56) // int32 | Page number of record list when record list exceeds specified page size (optional) (default to 1)
    size := int32(56) // int32 | Maximum number of records to return (optional) (default to 100)
    search := "search_example" // string | Specifies the search criteria. The syntax of this parameter is similar to the syntax of the _where_ clause of an SQL statement, using the names of the json attributes / column names of the account. For example, in order to retrieve all the accounts with a username starting with `my`:  ```sql username like 'my%' ```  The search criteria can also be applied on related resource. For example, in order to retrieve all the subscriptions labeled by `foo=bar`,  ```sql labels.key = 'foo' and labels.value = 'bar' ```  If the parameter isn't provided, or if the value is empty, then all the accounts that the user has permission to see will be returned. (optional)
    fields := "fields_example" // string | Supplies a comma-separated list of fields to be returned. Fields of sub-structures and of arrays use <structure>.<field> notation. <stucture>.* means all field of a structure Example: For each Subscription to get id, href, plan(id and kind) and labels (all fields)  ``` ocm get subscriptions --parameter fields=id,href,plan.id,plan.kind,labels.* --parameter fetchLabels=true ``` (optional)
    orderBy := "orderBy_example" // string | Specifies the order by criteria. The syntax of this parameter is similar to the syntax of the _order by_ clause of an SQL statement, but using the names of the json attributes / column of the account. For example, in order to retrieve all accounts ordered by username:  ```sql username asc ```  Or in order to retrieve all accounts ordered by username _and_ first name:  ```sql username asc, firstName asc ```  If the parameter isn't provided, or if the value is empty, then no explicit ordering will be applied. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ApiAccountsMgmtV1SubscriptionsSubIdNotificationContactsGet(context.Background(), subId).Page(page).Size(size).Search(search).Fields(fields).OrderBy(orderBy).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ApiAccountsMgmtV1SubscriptionsSubIdNotificationContactsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiAccountsMgmtV1SubscriptionsSubIdNotificationContactsGet`: AccountList
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ApiAccountsMgmtV1SubscriptionsSubIdNotificationContactsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subId** | **string** | The id of subscription | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiAccountsMgmtV1SubscriptionsSubIdNotificationContactsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | Page number of record list when record list exceeds specified page size | [default to 1]
 **size** | **int32** | Maximum number of records to return | [default to 100]
 **search** | **string** | Specifies the search criteria. The syntax of this parameter is similar to the syntax of the _where_ clause of an SQL statement, using the names of the json attributes / column names of the account. For example, in order to retrieve all the accounts with a username starting with &#x60;my&#x60;:  &#x60;&#x60;&#x60;sql username like &#39;my%&#39; &#x60;&#x60;&#x60;  The search criteria can also be applied on related resource. For example, in order to retrieve all the subscriptions labeled by &#x60;foo&#x3D;bar&#x60;,  &#x60;&#x60;&#x60;sql labels.key &#x3D; &#39;foo&#39; and labels.value &#x3D; &#39;bar&#39; &#x60;&#x60;&#x60;  If the parameter isn&#39;t provided, or if the value is empty, then all the accounts that the user has permission to see will be returned. | 
 **fields** | **string** | Supplies a comma-separated list of fields to be returned. Fields of sub-structures and of arrays use &lt;structure&gt;.&lt;field&gt; notation. &lt;stucture&gt;.* means all field of a structure Example: For each Subscription to get id, href, plan(id and kind) and labels (all fields)  &#x60;&#x60;&#x60; ocm get subscriptions --parameter fields&#x3D;id,href,plan.id,plan.kind,labels.* --parameter fetchLabels&#x3D;true &#x60;&#x60;&#x60; | 
 **orderBy** | **string** | Specifies the order by criteria. The syntax of this parameter is similar to the syntax of the _order by_ clause of an SQL statement, but using the names of the json attributes / column of the account. For example, in order to retrieve all accounts ordered by username:  &#x60;&#x60;&#x60;sql username asc &#x60;&#x60;&#x60;  Or in order to retrieve all accounts ordered by username _and_ first name:  &#x60;&#x60;&#x60;sql username asc, firstName asc &#x60;&#x60;&#x60;  If the parameter isn&#39;t provided, or if the value is empty, then no explicit ordering will be applied. | 

### Return type

[**AccountList**](AccountList.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiAccountsMgmtV1SubscriptionsSubIdNotificationContactsPost

> Account ApiAccountsMgmtV1SubscriptionsSubIdNotificationContactsPost(ctx, subId).NotificationContactCreateRequest(notificationContactCreateRequest).Execute()

Add an account as a notification contact to this subscription

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    subId := "subId_example" // string | The id of subscription
    notificationContactCreateRequest := *openapiclient.NewNotificationContactCreateRequest() // NotificationContactCreateRequest | Add a notification contact by an account's username

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ApiAccountsMgmtV1SubscriptionsSubIdNotificationContactsPost(context.Background(), subId).NotificationContactCreateRequest(notificationContactCreateRequest).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ApiAccountsMgmtV1SubscriptionsSubIdNotificationContactsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiAccountsMgmtV1SubscriptionsSubIdNotificationContactsPost`: Account
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ApiAccountsMgmtV1SubscriptionsSubIdNotificationContactsPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subId** | **string** | The id of subscription | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiAccountsMgmtV1SubscriptionsSubIdNotificationContactsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **notificationContactCreateRequest** | [**NotificationContactCreateRequest**](NotificationContactCreateRequest.md) | Add a notification contact by an account&#39;s username | 

### Return type

[**Account**](Account.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiAccountsMgmtV1SubscriptionsSubIdReservedResourcesReservedResourceIdGet

> ReservedResource ApiAccountsMgmtV1SubscriptionsSubIdReservedResourcesReservedResourceIdGet(ctx, subId, reservedResourceId).Execute()

Get reserved resources by id

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    subId := "subId_example" // string | The id of subscription
    reservedResourceId := "reservedResourceId_example" // string | The id of reserved resource

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ApiAccountsMgmtV1SubscriptionsSubIdReservedResourcesReservedResourceIdGet(context.Background(), subId, reservedResourceId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ApiAccountsMgmtV1SubscriptionsSubIdReservedResourcesReservedResourceIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiAccountsMgmtV1SubscriptionsSubIdReservedResourcesReservedResourceIdGet`: ReservedResource
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ApiAccountsMgmtV1SubscriptionsSubIdReservedResourcesReservedResourceIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subId** | **string** | The id of subscription | 
**reservedResourceId** | **string** | The id of reserved resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiAccountsMgmtV1SubscriptionsSubIdReservedResourcesReservedResourceIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ReservedResource**](ReservedResource.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiAccountsMgmtV1SupportCasesCaseIdDelete

> ApiAccountsMgmtV1SupportCasesCaseIdDelete(ctx, caseId).Execute()

Delete a support case

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    caseId := "caseId_example" // string | The id of a support case

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ApiAccountsMgmtV1SupportCasesCaseIdDelete(context.Background(), caseId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ApiAccountsMgmtV1SupportCasesCaseIdDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**caseId** | **string** | The id of a support case | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiAccountsMgmtV1SupportCasesCaseIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiAccountsMgmtV1SupportCasesPost

> SupportCasesCreatedResponse ApiAccountsMgmtV1SupportCasesPost(ctx).SupportCasesRequest(supportCasesRequest).Execute()

create a support case for the subscription

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    supportCasesRequest := *openapiclient.NewSupportCasesRequest("Description_example", "Severity_example", "Summary_example") // SupportCasesRequest | The contents of the support case to be created

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ApiAccountsMgmtV1SupportCasesPost(context.Background()).SupportCasesRequest(supportCasesRequest).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ApiAccountsMgmtV1SupportCasesPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiAccountsMgmtV1SupportCasesPost`: SupportCasesCreatedResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ApiAccountsMgmtV1SupportCasesPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiAccountsMgmtV1SupportCasesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **supportCasesRequest** | [**SupportCasesRequest**](SupportCasesRequest.md) | The contents of the support case to be created | 

### Return type

[**SupportCasesCreatedResponse**](SupportCasesCreatedResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiAccountsMgmtV1TokenAuthorizationPost

> TokenAuthorizationResponse ApiAccountsMgmtV1TokenAuthorizationPost(ctx).TokenAuthorizationRequest(tokenAuthorizationRequest).Execute()

Finds the account owner of the provided token

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    tokenAuthorizationRequest := *openapiclient.NewTokenAuthorizationRequest() // TokenAuthorizationRequest | Token authorization data

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ApiAccountsMgmtV1TokenAuthorizationPost(context.Background()).TokenAuthorizationRequest(tokenAuthorizationRequest).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ApiAccountsMgmtV1TokenAuthorizationPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiAccountsMgmtV1TokenAuthorizationPost`: TokenAuthorizationResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ApiAccountsMgmtV1TokenAuthorizationPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiAccountsMgmtV1TokenAuthorizationPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tokenAuthorizationRequest** | [**TokenAuthorizationRequest**](TokenAuthorizationRequest.md) | Token authorization data | 

### Return type

[**TokenAuthorizationResponse**](TokenAuthorizationResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiAuthorizationsV1AccessReviewPost

> AccessReviewResponse ApiAuthorizationsV1AccessReviewPost(ctx).AccessReview(accessReview).Execute()

Review an account's access to perform an action on a particular resource or resource type

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accessReview := *openapiclient.NewAccessReview("AccountUsername_example", "Action_example", "ResourceType_example") // AccessReview | Access review data

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ApiAuthorizationsV1AccessReviewPost(context.Background()).AccessReview(accessReview).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ApiAuthorizationsV1AccessReviewPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiAuthorizationsV1AccessReviewPost`: AccessReviewResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ApiAuthorizationsV1AccessReviewPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiAuthorizationsV1AccessReviewPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accessReview** | [**AccessReview**](AccessReview.md) | Access review data | 

### Return type

[**AccessReviewResponse**](AccessReviewResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiAuthorizationsV1CapabilityReviewPost

> CapabilityReview ApiAuthorizationsV1CapabilityReviewPost(ctx).CapabilityReviewRequest(capabilityReviewRequest).Execute()

Review an account's capabilities

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    capabilityReviewRequest := *openapiclient.NewCapabilityReviewRequest("AccountUsername_example", "Capability_example", "Type_example") // CapabilityReviewRequest | Capability review data

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ApiAuthorizationsV1CapabilityReviewPost(context.Background()).CapabilityReviewRequest(capabilityReviewRequest).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ApiAuthorizationsV1CapabilityReviewPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiAuthorizationsV1CapabilityReviewPost`: CapabilityReview
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ApiAuthorizationsV1CapabilityReviewPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiAuthorizationsV1CapabilityReviewPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **capabilityReviewRequest** | [**CapabilityReviewRequest**](CapabilityReviewRequest.md) | Capability review data | 

### Return type

[**CapabilityReview**](CapabilityReview.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiAuthorizationsV1ExportControlReviewPost

> ExportControlReview ApiAuthorizationsV1ExportControlReviewPost(ctx).ExportControlReviewRequest(exportControlReviewRequest).Execute()

Determine whether a user is restricted from downloading Red Hat software based on export control compliance. 

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    exportControlReviewRequest := *openapiclient.NewExportControlReviewRequest("AccountUsername_example") // ExportControlReviewRequest | Export control review data

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ApiAuthorizationsV1ExportControlReviewPost(context.Background()).ExportControlReviewRequest(exportControlReviewRequest).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ApiAuthorizationsV1ExportControlReviewPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiAuthorizationsV1ExportControlReviewPost`: ExportControlReview
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ApiAuthorizationsV1ExportControlReviewPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiAuthorizationsV1ExportControlReviewPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **exportControlReviewRequest** | [**ExportControlReviewRequest**](ExportControlReviewRequest.md) | Export control review data | 

### Return type

[**ExportControlReview**](ExportControlReview.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiAuthorizationsV1ResourceReviewPost

> ResourceReview ApiAuthorizationsV1ResourceReviewPost(ctx).ResourceReviewRequest(resourceReviewRequest).Execute()

Obtain resource ids for resources an account may perform the specified action upon. Resource ids returned as [\"*\"] is shorthand for all ids.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    resourceReviewRequest := *openapiclient.NewResourceReviewRequest() // ResourceReviewRequest | Resource review data

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ApiAuthorizationsV1ResourceReviewPost(context.Background()).ResourceReviewRequest(resourceReviewRequest).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ApiAuthorizationsV1ResourceReviewPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiAuthorizationsV1ResourceReviewPost`: ResourceReview
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ApiAuthorizationsV1ResourceReviewPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiAuthorizationsV1ResourceReviewPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **resourceReviewRequest** | [**ResourceReviewRequest**](ResourceReviewRequest.md) | Resource review data | 

### Return type

[**ResourceReview**](ResourceReview.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiAuthorizationsV1SelfAccessReviewPost

> AccessReviewResponse ApiAuthorizationsV1SelfAccessReviewPost(ctx).SelfAccessReview(selfAccessReview).Execute()

Review your ability to perform an action on a particular resource or resource type

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    selfAccessReview := *openapiclient.NewSelfAccessReview("Action_example", "ResourceType_example") // SelfAccessReview | Self access review data

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ApiAuthorizationsV1SelfAccessReviewPost(context.Background()).SelfAccessReview(selfAccessReview).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ApiAuthorizationsV1SelfAccessReviewPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiAuthorizationsV1SelfAccessReviewPost`: AccessReviewResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ApiAuthorizationsV1SelfAccessReviewPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiAuthorizationsV1SelfAccessReviewPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **selfAccessReview** | [**SelfAccessReview**](SelfAccessReview.md) | Self access review data | 

### Return type

[**AccessReviewResponse**](AccessReviewResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiAuthorizationsV1SelfResourceReviewPost

> SelfResourceReview ApiAuthorizationsV1SelfResourceReviewPost(ctx).SelfResourceReviewRequest(selfResourceReviewRequest).Execute()

Obtain resource ids for resources you may perform the specified action upon. Resource ids returned as [\"*\"] is shorthand for all ids.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    selfResourceReviewRequest := *openapiclient.NewSelfResourceReviewRequest() // SelfResourceReviewRequest | Self resource review data

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ApiAuthorizationsV1SelfResourceReviewPost(context.Background()).SelfResourceReviewRequest(selfResourceReviewRequest).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ApiAuthorizationsV1SelfResourceReviewPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiAuthorizationsV1SelfResourceReviewPost`: SelfResourceReview
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ApiAuthorizationsV1SelfResourceReviewPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiAuthorizationsV1SelfResourceReviewPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **selfResourceReviewRequest** | [**SelfResourceReviewRequest**](SelfResourceReviewRequest.md) | Self resource review data | 

### Return type

[**SelfResourceReview**](SelfResourceReview.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiAuthorizationsV1SelfTermsReviewPost

> TermsReviewResponse ApiAuthorizationsV1SelfTermsReviewPost(ctx).Execute()

Review your status of Terms

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ApiAuthorizationsV1SelfTermsReviewPost(context.Background()).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ApiAuthorizationsV1SelfTermsReviewPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiAuthorizationsV1SelfTermsReviewPost`: TermsReviewResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ApiAuthorizationsV1SelfTermsReviewPost`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiApiAuthorizationsV1SelfTermsReviewPostRequest struct via the builder pattern


### Return type

[**TermsReviewResponse**](TermsReviewResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiAuthorizationsV1TermsReviewPost

> TermsReviewResponse ApiAuthorizationsV1TermsReviewPost(ctx).TermsReview(termsReview).Execute()

Review an account's status of Terms

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    termsReview := *openapiclient.NewTermsReview("AccountUsername_example") // TermsReview | Data to check terms for

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ApiAuthorizationsV1TermsReviewPost(context.Background()).TermsReview(termsReview).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ApiAuthorizationsV1TermsReviewPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiAuthorizationsV1TermsReviewPost`: TermsReviewResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ApiAuthorizationsV1TermsReviewPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiAuthorizationsV1TermsReviewPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **termsReview** | [**TermsReview**](TermsReview.md) | Data to check terms for | 

### Return type

[**TermsReviewResponse**](TermsReviewResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

