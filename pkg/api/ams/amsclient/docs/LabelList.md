# LabelList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | **string** |  | 
**Page** | **int32** |  | 
**Size** | **int32** |  | 
**Total** | **int32** |  | 
**Items** | [**[]Label**](Label.md) |  | 

## Methods

### NewLabelList

`func NewLabelList(kind string, page int32, size int32, total int32, items []Label, ) *LabelList`

NewLabelList instantiates a new LabelList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLabelListWithDefaults

`func NewLabelListWithDefaults() *LabelList`

NewLabelListWithDefaults instantiates a new LabelList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *LabelList) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *LabelList) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *LabelList) SetKind(v string)`

SetKind sets Kind field to given value.


### GetPage

`func (o *LabelList) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *LabelList) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *LabelList) SetPage(v int32)`

SetPage sets Page field to given value.


### GetSize

`func (o *LabelList) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *LabelList) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *LabelList) SetSize(v int32)`

SetSize sets Size field to given value.


### GetTotal

`func (o *LabelList) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *LabelList) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *LabelList) SetTotal(v int32)`

SetTotal sets Total field to given value.


### GetItems

`func (o *LabelList) GetItems() []Label`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *LabelList) GetItemsOk() (*[]Label, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *LabelList) SetItems(v []Label)`

SetItems sets Items field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


