# PlaygroundToolsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CodeExecution** | [**PlaygroundToolStatus**](PlaygroundToolStatus.md) |  | 
**McpServers** | [**[]PlaygroundMcpServer**](PlaygroundMcpServer.md) |  | 
**WebSearch** | [**PlaygroundToolStatus**](PlaygroundToolStatus.md) |  | 

## Methods

### NewPlaygroundToolsResponse

`func NewPlaygroundToolsResponse(codeExecution PlaygroundToolStatus, mcpServers []PlaygroundMcpServer, webSearch PlaygroundToolStatus, ) *PlaygroundToolsResponse`

NewPlaygroundToolsResponse instantiates a new PlaygroundToolsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlaygroundToolsResponseWithDefaults

`func NewPlaygroundToolsResponseWithDefaults() *PlaygroundToolsResponse`

NewPlaygroundToolsResponseWithDefaults instantiates a new PlaygroundToolsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCodeExecution

`func (o *PlaygroundToolsResponse) GetCodeExecution() PlaygroundToolStatus`

GetCodeExecution returns the CodeExecution field if non-nil, zero value otherwise.

### GetCodeExecutionOk

`func (o *PlaygroundToolsResponse) GetCodeExecutionOk() (*PlaygroundToolStatus, bool)`

GetCodeExecutionOk returns a tuple with the CodeExecution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodeExecution

`func (o *PlaygroundToolsResponse) SetCodeExecution(v PlaygroundToolStatus)`

SetCodeExecution sets CodeExecution field to given value.


### GetMcpServers

`func (o *PlaygroundToolsResponse) GetMcpServers() []PlaygroundMcpServer`

GetMcpServers returns the McpServers field if non-nil, zero value otherwise.

### GetMcpServersOk

`func (o *PlaygroundToolsResponse) GetMcpServersOk() (*[]PlaygroundMcpServer, bool)`

GetMcpServersOk returns a tuple with the McpServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMcpServers

`func (o *PlaygroundToolsResponse) SetMcpServers(v []PlaygroundMcpServer)`

SetMcpServers sets McpServers field to given value.


### GetWebSearch

`func (o *PlaygroundToolsResponse) GetWebSearch() PlaygroundToolStatus`

GetWebSearch returns the WebSearch field if non-nil, zero value otherwise.

### GetWebSearchOk

`func (o *PlaygroundToolsResponse) GetWebSearchOk() (*PlaygroundToolStatus, bool)`

GetWebSearchOk returns a tuple with the WebSearch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebSearch

`func (o *PlaygroundToolsResponse) SetWebSearch(v PlaygroundToolStatus)`

SetWebSearch sets WebSearch field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


