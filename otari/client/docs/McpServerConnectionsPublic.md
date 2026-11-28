# McpServerConnectionsPublic

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | **int32** | How many users are connected in all | 
**Data** | [**[]McpServerConnectionPublic**](McpServerConnectionPublic.md) |  | 

## Methods

### NewMcpServerConnectionsPublic

`func NewMcpServerConnectionsPublic(count int32, data []McpServerConnectionPublic, ) *McpServerConnectionsPublic`

NewMcpServerConnectionsPublic instantiates a new McpServerConnectionsPublic object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMcpServerConnectionsPublicWithDefaults

`func NewMcpServerConnectionsPublicWithDefaults() *McpServerConnectionsPublic`

NewMcpServerConnectionsPublicWithDefaults instantiates a new McpServerConnectionsPublic object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *McpServerConnectionsPublic) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *McpServerConnectionsPublic) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *McpServerConnectionsPublic) SetCount(v int32)`

SetCount sets Count field to given value.


### GetData

`func (o *McpServerConnectionsPublic) GetData() []McpServerConnectionPublic`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *McpServerConnectionsPublic) GetDataOk() (*[]McpServerConnectionPublic, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *McpServerConnectionsPublic) SetData(v []McpServerConnectionPublic)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


