# Connector

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Kind** | Pointer to **string** |  | [optional] 
**Href** | Pointer to **string** |  | [optional] 
**Metadata** | Pointer to [**ConnectorAllOfMetadata**](Connector_allOf_metadata.md) |  | [optional] 
**DeploymentLocation** | Pointer to [**ConnectorAllOfDeploymentLocation**](Connector_allOf_deployment_location.md) |  | [optional] 
**ConnectorTypeId** | Pointer to **string** |  | [optional] 
**ConnectorSpec** | Pointer to **map[string]interface{}** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 

## Methods

### NewConnector

`func NewConnector() *Connector`

NewConnector instantiates a new Connector object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectorWithDefaults

`func NewConnectorWithDefaults() *Connector`

NewConnectorWithDefaults instantiates a new Connector object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Connector) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Connector) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Connector) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Connector) HasId() bool`

HasId returns a boolean if a field has been set.

### GetKind

`func (o *Connector) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *Connector) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *Connector) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *Connector) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetHref

`func (o *Connector) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *Connector) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *Connector) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *Connector) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetMetadata

`func (o *Connector) GetMetadata() ConnectorAllOfMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *Connector) GetMetadataOk() (*ConnectorAllOfMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *Connector) SetMetadata(v ConnectorAllOfMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *Connector) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetDeploymentLocation

`func (o *Connector) GetDeploymentLocation() ConnectorAllOfDeploymentLocation`

GetDeploymentLocation returns the DeploymentLocation field if non-nil, zero value otherwise.

### GetDeploymentLocationOk

`func (o *Connector) GetDeploymentLocationOk() (*ConnectorAllOfDeploymentLocation, bool)`

GetDeploymentLocationOk returns a tuple with the DeploymentLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentLocation

`func (o *Connector) SetDeploymentLocation(v ConnectorAllOfDeploymentLocation)`

SetDeploymentLocation sets DeploymentLocation field to given value.

### HasDeploymentLocation

`func (o *Connector) HasDeploymentLocation() bool`

HasDeploymentLocation returns a boolean if a field has been set.

### GetConnectorTypeId

`func (o *Connector) GetConnectorTypeId() string`

GetConnectorTypeId returns the ConnectorTypeId field if non-nil, zero value otherwise.

### GetConnectorTypeIdOk

`func (o *Connector) GetConnectorTypeIdOk() (*string, bool)`

GetConnectorTypeIdOk returns a tuple with the ConnectorTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectorTypeId

`func (o *Connector) SetConnectorTypeId(v string)`

SetConnectorTypeId sets ConnectorTypeId field to given value.

### HasConnectorTypeId

`func (o *Connector) HasConnectorTypeId() bool`

HasConnectorTypeId returns a boolean if a field has been set.

### GetConnectorSpec

`func (o *Connector) GetConnectorSpec() map[string]interface{}`

GetConnectorSpec returns the ConnectorSpec field if non-nil, zero value otherwise.

### GetConnectorSpecOk

`func (o *Connector) GetConnectorSpecOk() (*map[string]interface{}, bool)`

GetConnectorSpecOk returns a tuple with the ConnectorSpec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectorSpec

`func (o *Connector) SetConnectorSpec(v map[string]interface{})`

SetConnectorSpec sets ConnectorSpec field to given value.

### HasConnectorSpec

`func (o *Connector) HasConnectorSpec() bool`

HasConnectorSpec returns a boolean if a field has been set.

### GetStatus

`func (o *Connector) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Connector) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Connector) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Connector) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


