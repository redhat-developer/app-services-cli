# ServiceAccountList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | **string** |  | 
**Items** | [**[]ServiceAccountListItem**](ServiceAccountListItem.md) |  | 

## Methods

### NewServiceAccountList

`func NewServiceAccountList(kind string, items []ServiceAccountListItem, ) *ServiceAccountList`

NewServiceAccountList instantiates a new ServiceAccountList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceAccountListWithDefaults

`func NewServiceAccountListWithDefaults() *ServiceAccountList`

NewServiceAccountListWithDefaults instantiates a new ServiceAccountList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *ServiceAccountList) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ServiceAccountList) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ServiceAccountList) SetKind(v string)`

SetKind sets Kind field to given value.


### GetItems

`func (o *ServiceAccountList) GetItems() []ServiceAccountListItem`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *ServiceAccountList) GetItemsOk() (*[]ServiceAccountListItem, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *ServiceAccountList) SetItems(v []ServiceAccountListItem)`

SetItems sets Items field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


