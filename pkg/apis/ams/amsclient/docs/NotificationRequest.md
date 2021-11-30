# NotificationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BccAddress** | Pointer to **string** |  | [optional] 
**ClusterId** | Pointer to **string** |  | [optional] 
**ClusterUuid** | Pointer to **string** |  | [optional] 
**IncludeRedHatAssociates** | Pointer to **bool** |  | [optional] 
**Subject** | **string** |  | 
**SubscriptionId** | Pointer to **string** |  | [optional] 
**TemplateName** | **string** |  | 
**TemplateParameters** | Pointer to [**[]TemplateParameter**](TemplateParameter.md) |  | [optional] 

## Methods

### NewNotificationRequest

`func NewNotificationRequest(subject string, templateName string, ) *NotificationRequest`

NewNotificationRequest instantiates a new NotificationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationRequestWithDefaults

`func NewNotificationRequestWithDefaults() *NotificationRequest`

NewNotificationRequestWithDefaults instantiates a new NotificationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBccAddress

`func (o *NotificationRequest) GetBccAddress() string`

GetBccAddress returns the BccAddress field if non-nil, zero value otherwise.

### GetBccAddressOk

`func (o *NotificationRequest) GetBccAddressOk() (*string, bool)`

GetBccAddressOk returns a tuple with the BccAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBccAddress

`func (o *NotificationRequest) SetBccAddress(v string)`

SetBccAddress sets BccAddress field to given value.

### HasBccAddress

`func (o *NotificationRequest) HasBccAddress() bool`

HasBccAddress returns a boolean if a field has been set.

### GetClusterId

`func (o *NotificationRequest) GetClusterId() string`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *NotificationRequest) GetClusterIdOk() (*string, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *NotificationRequest) SetClusterId(v string)`

SetClusterId sets ClusterId field to given value.

### HasClusterId

`func (o *NotificationRequest) HasClusterId() bool`

HasClusterId returns a boolean if a field has been set.

### GetClusterUuid

`func (o *NotificationRequest) GetClusterUuid() string`

GetClusterUuid returns the ClusterUuid field if non-nil, zero value otherwise.

### GetClusterUuidOk

`func (o *NotificationRequest) GetClusterUuidOk() (*string, bool)`

GetClusterUuidOk returns a tuple with the ClusterUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterUuid

`func (o *NotificationRequest) SetClusterUuid(v string)`

SetClusterUuid sets ClusterUuid field to given value.

### HasClusterUuid

`func (o *NotificationRequest) HasClusterUuid() bool`

HasClusterUuid returns a boolean if a field has been set.

### GetIncludeRedHatAssociates

`func (o *NotificationRequest) GetIncludeRedHatAssociates() bool`

GetIncludeRedHatAssociates returns the IncludeRedHatAssociates field if non-nil, zero value otherwise.

### GetIncludeRedHatAssociatesOk

`func (o *NotificationRequest) GetIncludeRedHatAssociatesOk() (*bool, bool)`

GetIncludeRedHatAssociatesOk returns a tuple with the IncludeRedHatAssociates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeRedHatAssociates

`func (o *NotificationRequest) SetIncludeRedHatAssociates(v bool)`

SetIncludeRedHatAssociates sets IncludeRedHatAssociates field to given value.

### HasIncludeRedHatAssociates

`func (o *NotificationRequest) HasIncludeRedHatAssociates() bool`

HasIncludeRedHatAssociates returns a boolean if a field has been set.

### GetSubject

`func (o *NotificationRequest) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *NotificationRequest) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *NotificationRequest) SetSubject(v string)`

SetSubject sets Subject field to given value.


### GetSubscriptionId

`func (o *NotificationRequest) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *NotificationRequest) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *NotificationRequest) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.

### HasSubscriptionId

`func (o *NotificationRequest) HasSubscriptionId() bool`

HasSubscriptionId returns a boolean if a field has been set.

### GetTemplateName

`func (o *NotificationRequest) GetTemplateName() string`

GetTemplateName returns the TemplateName field if non-nil, zero value otherwise.

### GetTemplateNameOk

`func (o *NotificationRequest) GetTemplateNameOk() (*string, bool)`

GetTemplateNameOk returns a tuple with the TemplateName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateName

`func (o *NotificationRequest) SetTemplateName(v string)`

SetTemplateName sets TemplateName field to given value.


### GetTemplateParameters

`func (o *NotificationRequest) GetTemplateParameters() []TemplateParameter`

GetTemplateParameters returns the TemplateParameters field if non-nil, zero value otherwise.

### GetTemplateParametersOk

`func (o *NotificationRequest) GetTemplateParametersOk() (*[]TemplateParameter, bool)`

GetTemplateParametersOk returns a tuple with the TemplateParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateParameters

`func (o *NotificationRequest) SetTemplateParameters(v []TemplateParameter)`

SetTemplateParameters sets TemplateParameters field to given value.

### HasTemplateParameters

`func (o *NotificationRequest) HasTemplateParameters() bool`

HasTemplateParameters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


