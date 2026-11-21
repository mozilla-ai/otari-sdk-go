# McpConnectionsPublic

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | **int32** | How many of the workspace&#39;s servers the caller is connected to in all | 
**Data** | [**[]McpConnectionPublic**](McpConnectionPublic.md) |  | 

## Methods

### NewMcpConnectionsPublic

`func NewMcpConnectionsPublic(count int32, data []McpConnectionPublic, ) *McpConnectionsPublic`

NewMcpConnectionsPublic instantiates a new McpConnectionsPublic object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMcpConnectionsPublicWithDefaults

`func NewMcpConnectionsPublicWithDefaults() *McpConnectionsPublic`

NewMcpConnectionsPublicWithDefaults instantiates a new McpConnectionsPublic object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *McpConnectionsPublic) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *McpConnectionsPublic) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *McpConnectionsPublic) SetCount(v int32)`

SetCount sets Count field to given value.


### GetData

`func (o *McpConnectionsPublic) GetData() []McpConnectionPublic`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *McpConnectionsPublic) GetDataOk() (*[]McpConnectionPublic, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *McpConnectionsPublic) SetData(v []McpConnectionPublic)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


