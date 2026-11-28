# McpServersResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LeftOut** | Pointer to [**[]McpServerLeftOut**](McpServerLeftOut.md) | Servers the key&#39;s access list grants that are not in &#x60;servers&#x60;, each with the reason. Empty for a key that belongs to a user. | [optional] 
**Servers** | [**[]McpServerSummary**](McpServerSummary.md) |  | 

## Methods

### NewMcpServersResponse

`func NewMcpServersResponse(servers []McpServerSummary, ) *McpServersResponse`

NewMcpServersResponse instantiates a new McpServersResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMcpServersResponseWithDefaults

`func NewMcpServersResponseWithDefaults() *McpServersResponse`

NewMcpServersResponseWithDefaults instantiates a new McpServersResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLeftOut

`func (o *McpServersResponse) GetLeftOut() []McpServerLeftOut`

GetLeftOut returns the LeftOut field if non-nil, zero value otherwise.

### GetLeftOutOk

`func (o *McpServersResponse) GetLeftOutOk() (*[]McpServerLeftOut, bool)`

GetLeftOutOk returns a tuple with the LeftOut field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeftOut

`func (o *McpServersResponse) SetLeftOut(v []McpServerLeftOut)`

SetLeftOut sets LeftOut field to given value.

### HasLeftOut

`func (o *McpServersResponse) HasLeftOut() bool`

HasLeftOut returns a boolean if a field has been set.

### GetServers

`func (o *McpServersResponse) GetServers() []McpServerSummary`

GetServers returns the Servers field if non-nil, zero value otherwise.

### GetServersOk

`func (o *McpServersResponse) GetServersOk() (*[]McpServerSummary, bool)`

GetServersOk returns a tuple with the Servers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServers

`func (o *McpServersResponse) SetServers(v []McpServerSummary)`

SetServers sets Servers field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


