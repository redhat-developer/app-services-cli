# SubscriptionCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterUuid** | **string** |  | 
**ConsoleUrl** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**PlanId** | **string** |  | 
**Status** | **string** |  | 

## Methods

### NewSubscriptionCreateRequest

`func NewSubscriptionCreateRequest(clusterUuid string, planId string, status string, ) *SubscriptionCreateRequest`

NewSubscriptionCreateRequest instantiates a new SubscriptionCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriptionCreateRequestWithDefaults

`func NewSubscriptionCreateRequestWithDefaults() *SubscriptionCreateRequest`

NewSubscriptionCreateRequestWithDefaults instantiates a new SubscriptionCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusterUuid

`func (o *SubscriptionCreateRequest) GetClusterUuid() string`

GetClusterUuid returns the ClusterUuid field if non-nil, zero value otherwise.

### GetClusterUuidOk

`func (o *SubscriptionCreateRequest) GetClusterUuidOk() (*string, bool)`

GetClusterUuidOk returns a tuple with the ClusterUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterUuid

`func (o *SubscriptionCreateRequest) SetClusterUuid(v string)`

SetClusterUuid sets ClusterUuid field to given value.


### GetConsoleUrl

`func (o *SubscriptionCreateRequest) GetConsoleUrl() string`

GetConsoleUrl returns the ConsoleUrl field if non-nil, zero value otherwise.

### GetConsoleUrlOk

`func (o *SubscriptionCreateRequest) GetConsoleUrlOk() (*string, bool)`

GetConsoleUrlOk returns a tuple with the ConsoleUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsoleUrl

`func (o *SubscriptionCreateRequest) SetConsoleUrl(v string)`

SetConsoleUrl sets ConsoleUrl field to given value.

### HasConsoleUrl

`func (o *SubscriptionCreateRequest) HasConsoleUrl() bool`

HasConsoleUrl returns a boolean if a field has been set.

### GetDisplayName

`func (o *SubscriptionCreateRequest) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *SubscriptionCreateRequest) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *SubscriptionCreateRequest) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *SubscriptionCreateRequest) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetPlanId

`func (o *SubscriptionCreateRequest) GetPlanId() string`

GetPlanId returns the PlanId field if non-nil, zero value otherwise.

### GetPlanIdOk

`func (o *SubscriptionCreateRequest) GetPlanIdOk() (*string, bool)`

GetPlanIdOk returns a tuple with the PlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanId

`func (o *SubscriptionCreateRequest) SetPlanId(v string)`

SetPlanId sets PlanId field to given value.


### GetStatus

`func (o *SubscriptionCreateRequest) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SubscriptionCreateRequest) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SubscriptionCreateRequest) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


