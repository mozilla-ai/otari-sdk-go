# McpExecuteRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Arguments** | Pointer to **map[string]interface{}** | The exact caller-authorized JSON-object arguments. | [optional] 
**ClientExecutionId** | **string** | A caller-generated UUID, for correlation only. It is not proof of approval and not an idempotency key: repeating a request with the same value may execute the tool again, so this request must never be retried automatically. | 
**McpServerId** | **string** | The stored MCP server to execute against. | 
**ServerRevision** | **string** | The stored-server revision returned by tool discovery. Required, and compared against the current one so a configuration change since the caller authorized this call is refused rather than executed. Not an approval credential. | 
**ToolName** | **string** | The remote MCP tool name the caller authorized for this one execution. | 

## Methods

### NewMcpExecuteRequest

`func NewMcpExecuteRequest(clientExecutionId string, mcpServerId string, serverRevision string, toolName string, ) *McpExecuteRequest`

NewMcpExecuteRequest instantiates a new McpExecuteRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMcpExecuteRequestWithDefaults

`func NewMcpExecuteRequestWithDefaults() *McpExecuteRequest`

NewMcpExecuteRequestWithDefaults instantiates a new McpExecuteRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArguments

`func (o *McpExecuteRequest) GetArguments() map[string]interface{}`

GetArguments returns the Arguments field if non-nil, zero value otherwise.

### GetArgumentsOk

`func (o *McpExecuteRequest) GetArgumentsOk() (*map[string]interface{}, bool)`

GetArgumentsOk returns a tuple with the Arguments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArguments

`func (o *McpExecuteRequest) SetArguments(v map[string]interface{})`

SetArguments sets Arguments field to given value.

### HasArguments

`func (o *McpExecuteRequest) HasArguments() bool`

HasArguments returns a boolean if a field has been set.

### GetClientExecutionId

`func (o *McpExecuteRequest) GetClientExecutionId() string`

GetClientExecutionId returns the ClientExecutionId field if non-nil, zero value otherwise.

### GetClientExecutionIdOk

`func (o *McpExecuteRequest) GetClientExecutionIdOk() (*string, bool)`

GetClientExecutionIdOk returns a tuple with the ClientExecutionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientExecutionId

`func (o *McpExecuteRequest) SetClientExecutionId(v string)`

SetClientExecutionId sets ClientExecutionId field to given value.


### GetMcpServerId

`func (o *McpExecuteRequest) GetMcpServerId() string`

GetMcpServerId returns the McpServerId field if non-nil, zero value otherwise.

### GetMcpServerIdOk

`func (o *McpExecuteRequest) GetMcpServerIdOk() (*string, bool)`

GetMcpServerIdOk returns a tuple with the McpServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMcpServerId

`func (o *McpExecuteRequest) SetMcpServerId(v string)`

SetMcpServerId sets McpServerId field to given value.


### GetServerRevision

`func (o *McpExecuteRequest) GetServerRevision() string`

GetServerRevision returns the ServerRevision field if non-nil, zero value otherwise.

### GetServerRevisionOk

`func (o *McpExecuteRequest) GetServerRevisionOk() (*string, bool)`

GetServerRevisionOk returns a tuple with the ServerRevision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerRevision

`func (o *McpExecuteRequest) SetServerRevision(v string)`

SetServerRevision sets ServerRevision field to given value.


### GetToolName

`func (o *McpExecuteRequest) GetToolName() string`

GetToolName returns the ToolName field if non-nil, zero value otherwise.

### GetToolNameOk

`func (o *McpExecuteRequest) GetToolNameOk() (*string, bool)`

GetToolNameOk returns a tuple with the ToolName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToolName

`func (o *McpExecuteRequest) SetToolName(v string)`

SetToolName sets ToolName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


