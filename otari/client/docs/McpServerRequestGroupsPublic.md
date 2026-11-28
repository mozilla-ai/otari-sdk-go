# McpServerRequestGroupsPublic

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | **int32** | How many requests are pending in all | 
**Data** | [**[]McpServerRequestGroupPublic**](McpServerRequestGroupPublic.md) |  | 
**ServerCount** | **int32** | How many catalog servers have a request pending, for paging &#x60;data&#x60; | 

## Methods

### NewMcpServerRequestGroupsPublic

`func NewMcpServerRequestGroupsPublic(count int32, data []McpServerRequestGroupPublic, serverCount int32, ) *McpServerRequestGroupsPublic`

NewMcpServerRequestGroupsPublic instantiates a new McpServerRequestGroupsPublic object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMcpServerRequestGroupsPublicWithDefaults

`func NewMcpServerRequestGroupsPublicWithDefaults() *McpServerRequestGroupsPublic`

NewMcpServerRequestGroupsPublicWithDefaults instantiates a new McpServerRequestGroupsPublic object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *McpServerRequestGroupsPublic) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *McpServerRequestGroupsPublic) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *McpServerRequestGroupsPublic) SetCount(v int32)`

SetCount sets Count field to given value.


### GetData

`func (o *McpServerRequestGroupsPublic) GetData() []McpServerRequestGroupPublic`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *McpServerRequestGroupsPublic) GetDataOk() (*[]McpServerRequestGroupPublic, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *McpServerRequestGroupsPublic) SetData(v []McpServerRequestGroupPublic)`

SetData sets Data field to given value.


### GetServerCount

`func (o *McpServerRequestGroupsPublic) GetServerCount() int32`

GetServerCount returns the ServerCount field if non-nil, zero value otherwise.

### GetServerCountOk

`func (o *McpServerRequestGroupsPublic) GetServerCountOk() (*int32, bool)`

GetServerCountOk returns a tuple with the ServerCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerCount

`func (o *McpServerRequestGroupsPublic) SetServerCount(v int32)`

SetServerCount sets ServerCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


