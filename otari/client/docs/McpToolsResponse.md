# McpToolsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServerId** | **string** |  | 
**ServerRevision** | **string** | An opaque revision of the stored server&#39;s URL, credential, enabled state and allowlist. It detects Otari-side and platform-side configuration changes only: a remote server that changes its own catalog or a tool&#39;s behavior behind an unchanged URL will not move it. | 
**Tools** | [**[]McpToolDefinition**](McpToolDefinition.md) |  | 
**Warnings** | [**[]McpToolWarning**](McpToolWarning.md) |  | 

## Methods

### NewMcpToolsResponse

`func NewMcpToolsResponse(serverId string, serverRevision string, tools []McpToolDefinition, warnings []McpToolWarning, ) *McpToolsResponse`

NewMcpToolsResponse instantiates a new McpToolsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMcpToolsResponseWithDefaults

`func NewMcpToolsResponseWithDefaults() *McpToolsResponse`

NewMcpToolsResponseWithDefaults instantiates a new McpToolsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServerId

`func (o *McpToolsResponse) GetServerId() string`

GetServerId returns the ServerId field if non-nil, zero value otherwise.

### GetServerIdOk

`func (o *McpToolsResponse) GetServerIdOk() (*string, bool)`

GetServerIdOk returns a tuple with the ServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerId

`func (o *McpToolsResponse) SetServerId(v string)`

SetServerId sets ServerId field to given value.


### GetServerRevision

`func (o *McpToolsResponse) GetServerRevision() string`

GetServerRevision returns the ServerRevision field if non-nil, zero value otherwise.

### GetServerRevisionOk

`func (o *McpToolsResponse) GetServerRevisionOk() (*string, bool)`

GetServerRevisionOk returns a tuple with the ServerRevision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerRevision

`func (o *McpToolsResponse) SetServerRevision(v string)`

SetServerRevision sets ServerRevision field to given value.


### GetTools

`func (o *McpToolsResponse) GetTools() []McpToolDefinition`

GetTools returns the Tools field if non-nil, zero value otherwise.

### GetToolsOk

`func (o *McpToolsResponse) GetToolsOk() (*[]McpToolDefinition, bool)`

GetToolsOk returns a tuple with the Tools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTools

`func (o *McpToolsResponse) SetTools(v []McpToolDefinition)`

SetTools sets Tools field to given value.


### GetWarnings

`func (o *McpToolsResponse) GetWarnings() []McpToolWarning`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *McpToolsResponse) GetWarningsOk() (*[]McpToolWarning, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *McpToolsResponse) SetWarnings(v []McpToolWarning)`

SetWarnings sets Warnings field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


