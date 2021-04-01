# SupportCasesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountNumber** | Pointer to **string** |  | [optional] 
**CaseLanguage** | Pointer to **string** |  | [optional] 
**ClusterId** | Pointer to **string** |  | [optional] 
**ClusterUuid** | Pointer to **string** |  | [optional] 
**ContactSsoName** | Pointer to **string** |  | [optional] 
**Description** | **string** |  | 
**EventStreamId** | Pointer to **string** |  | [optional] 
**OpenshiftClusterId** | Pointer to **string** |  | [optional] 
**Product** | Pointer to **string** |  | [optional] [default to "OpenShift Container Platform"]
**Severity** | **string** |  | 
**SubscriptionId** | Pointer to **string** |  | [optional] 
**Summary** | **string** |  | 
**Version** | Pointer to **string** |  | [optional] [default to "4.4"]

## Methods

### NewSupportCasesRequest

`func NewSupportCasesRequest(description string, severity string, summary string, ) *SupportCasesRequest`

NewSupportCasesRequest instantiates a new SupportCasesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSupportCasesRequestWithDefaults

`func NewSupportCasesRequestWithDefaults() *SupportCasesRequest`

NewSupportCasesRequestWithDefaults instantiates a new SupportCasesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountNumber

`func (o *SupportCasesRequest) GetAccountNumber() string`

GetAccountNumber returns the AccountNumber field if non-nil, zero value otherwise.

### GetAccountNumberOk

`func (o *SupportCasesRequest) GetAccountNumberOk() (*string, bool)`

GetAccountNumberOk returns a tuple with the AccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountNumber

`func (o *SupportCasesRequest) SetAccountNumber(v string)`

SetAccountNumber sets AccountNumber field to given value.

### HasAccountNumber

`func (o *SupportCasesRequest) HasAccountNumber() bool`

HasAccountNumber returns a boolean if a field has been set.

### GetCaseLanguage

`func (o *SupportCasesRequest) GetCaseLanguage() string`

GetCaseLanguage returns the CaseLanguage field if non-nil, zero value otherwise.

### GetCaseLanguageOk

`func (o *SupportCasesRequest) GetCaseLanguageOk() (*string, bool)`

GetCaseLanguageOk returns a tuple with the CaseLanguage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaseLanguage

`func (o *SupportCasesRequest) SetCaseLanguage(v string)`

SetCaseLanguage sets CaseLanguage field to given value.

### HasCaseLanguage

`func (o *SupportCasesRequest) HasCaseLanguage() bool`

HasCaseLanguage returns a boolean if a field has been set.

### GetClusterId

`func (o *SupportCasesRequest) GetClusterId() string`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *SupportCasesRequest) GetClusterIdOk() (*string, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *SupportCasesRequest) SetClusterId(v string)`

SetClusterId sets ClusterId field to given value.

### HasClusterId

`func (o *SupportCasesRequest) HasClusterId() bool`

HasClusterId returns a boolean if a field has been set.

### GetClusterUuid

`func (o *SupportCasesRequest) GetClusterUuid() string`

GetClusterUuid returns the ClusterUuid field if non-nil, zero value otherwise.

### GetClusterUuidOk

`func (o *SupportCasesRequest) GetClusterUuidOk() (*string, bool)`

GetClusterUuidOk returns a tuple with the ClusterUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterUuid

`func (o *SupportCasesRequest) SetClusterUuid(v string)`

SetClusterUuid sets ClusterUuid field to given value.

### HasClusterUuid

`func (o *SupportCasesRequest) HasClusterUuid() bool`

HasClusterUuid returns a boolean if a field has been set.

### GetContactSsoName

`func (o *SupportCasesRequest) GetContactSsoName() string`

GetContactSsoName returns the ContactSsoName field if non-nil, zero value otherwise.

### GetContactSsoNameOk

`func (o *SupportCasesRequest) GetContactSsoNameOk() (*string, bool)`

GetContactSsoNameOk returns a tuple with the ContactSsoName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactSsoName

`func (o *SupportCasesRequest) SetContactSsoName(v string)`

SetContactSsoName sets ContactSsoName field to given value.

### HasContactSsoName

`func (o *SupportCasesRequest) HasContactSsoName() bool`

HasContactSsoName returns a boolean if a field has been set.

### GetDescription

`func (o *SupportCasesRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SupportCasesRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SupportCasesRequest) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetEventStreamId

`func (o *SupportCasesRequest) GetEventStreamId() string`

GetEventStreamId returns the EventStreamId field if non-nil, zero value otherwise.

### GetEventStreamIdOk

`func (o *SupportCasesRequest) GetEventStreamIdOk() (*string, bool)`

GetEventStreamIdOk returns a tuple with the EventStreamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventStreamId

`func (o *SupportCasesRequest) SetEventStreamId(v string)`

SetEventStreamId sets EventStreamId field to given value.

### HasEventStreamId

`func (o *SupportCasesRequest) HasEventStreamId() bool`

HasEventStreamId returns a boolean if a field has been set.

### GetOpenshiftClusterId

`func (o *SupportCasesRequest) GetOpenshiftClusterId() string`

GetOpenshiftClusterId returns the OpenshiftClusterId field if non-nil, zero value otherwise.

### GetOpenshiftClusterIdOk

`func (o *SupportCasesRequest) GetOpenshiftClusterIdOk() (*string, bool)`

GetOpenshiftClusterIdOk returns a tuple with the OpenshiftClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenshiftClusterId

`func (o *SupportCasesRequest) SetOpenshiftClusterId(v string)`

SetOpenshiftClusterId sets OpenshiftClusterId field to given value.

### HasOpenshiftClusterId

`func (o *SupportCasesRequest) HasOpenshiftClusterId() bool`

HasOpenshiftClusterId returns a boolean if a field has been set.

### GetProduct

`func (o *SupportCasesRequest) GetProduct() string`

GetProduct returns the Product field if non-nil, zero value otherwise.

### GetProductOk

`func (o *SupportCasesRequest) GetProductOk() (*string, bool)`

GetProductOk returns a tuple with the Product field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduct

`func (o *SupportCasesRequest) SetProduct(v string)`

SetProduct sets Product field to given value.

### HasProduct

`func (o *SupportCasesRequest) HasProduct() bool`

HasProduct returns a boolean if a field has been set.

### GetSeverity

`func (o *SupportCasesRequest) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *SupportCasesRequest) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *SupportCasesRequest) SetSeverity(v string)`

SetSeverity sets Severity field to given value.


### GetSubscriptionId

`func (o *SupportCasesRequest) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *SupportCasesRequest) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *SupportCasesRequest) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.

### HasSubscriptionId

`func (o *SupportCasesRequest) HasSubscriptionId() bool`

HasSubscriptionId returns a boolean if a field has been set.

### GetSummary

`func (o *SupportCasesRequest) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *SupportCasesRequest) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *SupportCasesRequest) SetSummary(v string)`

SetSummary sets Summary field to given value.


### GetVersion

`func (o *SupportCasesRequest) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *SupportCasesRequest) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *SupportCasesRequest) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *SupportCasesRequest) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


