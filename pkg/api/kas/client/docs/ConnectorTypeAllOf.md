# ConnectorTypeAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name of the connector type. | [optional] 
**Version** | Pointer to **string** | Version of the connector type. | [optional] 
**Description** | Pointer to **string** | A description of the connector. | [optional] 
**JsonSchema** | Pointer to **map[string]interface{}** | A json schema that can be used to validate a connectors connector_spec field. | [optional] 

## Methods

### NewConnectorTypeAllOf

`func NewConnectorTypeAllOf() *ConnectorTypeAllOf`

NewConnectorTypeAllOf instantiates a new ConnectorTypeAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectorTypeAllOfWithDefaults

`func NewConnectorTypeAllOfWithDefaults() *ConnectorTypeAllOf`

NewConnectorTypeAllOfWithDefaults instantiates a new ConnectorTypeAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ConnectorTypeAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ConnectorTypeAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ConnectorTypeAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ConnectorTypeAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetVersion

`func (o *ConnectorTypeAllOf) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ConnectorTypeAllOf) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ConnectorTypeAllOf) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ConnectorTypeAllOf) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetDescription

`func (o *ConnectorTypeAllOf) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ConnectorTypeAllOf) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ConnectorTypeAllOf) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ConnectorTypeAllOf) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetJsonSchema

`func (o *ConnectorTypeAllOf) GetJsonSchema() map[string]interface{}`

GetJsonSchema returns the JsonSchema field if non-nil, zero value otherwise.

### GetJsonSchemaOk

`func (o *ConnectorTypeAllOf) GetJsonSchemaOk() (*map[string]interface{}, bool)`

GetJsonSchemaOk returns a tuple with the JsonSchema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJsonSchema

`func (o *ConnectorTypeAllOf) SetJsonSchema(v map[string]interface{})`

SetJsonSchema sets JsonSchema field to given value.

### HasJsonSchema

`func (o *ConnectorTypeAllOf) HasJsonSchema() bool`

HasJsonSchema returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


