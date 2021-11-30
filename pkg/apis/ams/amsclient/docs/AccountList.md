# AccountList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | **string** |  | 
**Page** | **int32** |  | 
**Size** | **int32** |  | 
**Total** | **int32** |  | 
**Items** | [**[]Account**](Account.md) |  | 

## Methods

### NewAccountList

`func NewAccountList(kind string, page int32, size int32, total int32, items []Account, ) *AccountList`

NewAccountList instantiates a new AccountList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountListWithDefaults

`func NewAccountListWithDefaults() *AccountList`

NewAccountListWithDefaults instantiates a new AccountList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *AccountList) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *AccountList) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *AccountList) SetKind(v string)`

SetKind sets Kind field to given value.


### GetPage

`func (o *AccountList) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *AccountList) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *AccountList) SetPage(v int32)`

SetPage sets Page field to given value.


### GetSize

`func (o *AccountList) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *AccountList) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *AccountList) SetSize(v int32)`

SetSize sets Size field to given value.


### GetTotal

`func (o *AccountList) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *AccountList) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *AccountList) SetTotal(v int32)`

SetTotal sets Total field to given value.


### GetItems

`func (o *AccountList) GetItems() []Account`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *AccountList) GetItemsOk() (*[]Account, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *AccountList) SetItems(v []Account)`

SetItems sets Items field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


