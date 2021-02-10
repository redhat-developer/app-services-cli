# ConnectorAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metadata** | Pointer to [**ConnectorAllOfMetadata**](Connector_allOf_metadata.md) |  | [optional] 
**DeploymentLocation** | Pointer to [**ConnectorAllOfDeploymentLocation**](Connector_allOf_deployment_location.md) |  | [optional] 
**ConnectorTypeId** | Pointer to **string** |  | [optional] 
**ConnectorSpec** | Pointer to **map[string]interface{}** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 

## Methods

### NewConnectorAllOf

`func NewConnectorAllOf() *ConnectorAllOf`

NewConnectorAllOf instantiates a new ConnectorAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectorAllOfWithDefaults

`func NewConnectorAllOfWithDefaults() *ConnectorAllOf`

NewConnectorAllOfWithDefaults instantiates a new ConnectorAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadata

`func (o *ConnectorAllOf) GetMetadata() ConnectorAllOfMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ConnectorAllOf) GetMetadataOk() (*ConnectorAllOfMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ConnectorAllOf) SetMetadata(v ConnectorAllOfMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ConnectorAllOf) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetDeploymentLocation

`func (o *ConnectorAllOf) GetDeploymentLocation() ConnectorAllOfDeploymentLocation`

GetDeploymentLocation returns the DeploymentLocation field if non-nil, zero value otherwise.

### GetDeploymentLocationOk

`func (o *ConnectorAllOf) GetDeploymentLocationOk() (*ConnectorAllOfDeploymentLocation, bool)`

GetDeploymentLocationOk returns a tuple with the DeploymentLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentLocation

`func (o *ConnectorAllOf) SetDeploymentLocation(v ConnectorAllOfDeploymentLocation)`

SetDeploymentLocation sets DeploymentLocation field to given value.

### HasDeploymentLocation

`func (o *ConnectorAllOf) HasDeploymentLocation() bool`

HasDeploymentLocation returns a boolean if a field has been set.

### GetConnectorTypeId

`func (o *ConnectorAllOf) GetConnectorTypeId() string`

GetConnectorTypeId returns the ConnectorTypeId field if non-nil, zero value otherwise.

### GetConnectorTypeIdOk

`func (o *ConnectorAllOf) GetConnectorTypeIdOk() (*string, bool)`

GetConnectorTypeIdOk returns a tuple with the ConnectorTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectorTypeId

`func (o *ConnectorAllOf) SetConnectorTypeId(v string)`

SetConnectorTypeId sets ConnectorTypeId field to given value.

### HasConnectorTypeId

`func (o *ConnectorAllOf) HasConnectorTypeId() bool`

HasConnectorTypeId returns a boolean if a field has been set.

### GetConnectorSpec

`func (o *ConnectorAllOf) GetConnectorSpec() map[string]interface{}`

GetConnectorSpec returns the ConnectorSpec field if non-nil, zero value otherwise.

### GetConnectorSpecOk

`func (o *ConnectorAllOf) GetConnectorSpecOk() (*map[string]interface{}, bool)`

GetConnectorSpecOk returns a tuple with the ConnectorSpec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectorSpec

`func (o *ConnectorAllOf) SetConnectorSpec(v map[string]interface{})`

SetConnectorSpec sets ConnectorSpec field to given value.

### HasConnectorSpec

`func (o *ConnectorAllOf) HasConnectorSpec() bool`

HasConnectorSpec returns a boolean if a field has been set.

### GetStatus

`func (o *ConnectorAllOf) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ConnectorAllOf) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ConnectorAllOf) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ConnectorAllOf) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


