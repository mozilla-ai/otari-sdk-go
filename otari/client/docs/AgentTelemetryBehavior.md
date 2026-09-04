# AgentTelemetryBehavior

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiErrors** | Pointer to **int32** |  | [optional] [default to 0]
**ByTool** | Pointer to [**[]AgentTelemetryToolRow**](AgentTelemetryToolRow.md) |  | [optional] [default to {}]
**Sessions** | Pointer to **int32** |  | [optional] [default to 0]
**ToolAccepts** | Pointer to **int32** |  | [optional] [default to 0]
**ToolCalls** | Pointer to **int32** |  | [optional] [default to 0]
**ToolRejects** | Pointer to **int32** |  | [optional] [default to 0]
**Turns** | Pointer to **int32** |  | [optional] [default to 0]

## Methods

### NewAgentTelemetryBehavior

`func NewAgentTelemetryBehavior() *AgentTelemetryBehavior`

NewAgentTelemetryBehavior instantiates a new AgentTelemetryBehavior object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentTelemetryBehaviorWithDefaults

`func NewAgentTelemetryBehaviorWithDefaults() *AgentTelemetryBehavior`

NewAgentTelemetryBehaviorWithDefaults instantiates a new AgentTelemetryBehavior object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiErrors

`func (o *AgentTelemetryBehavior) GetApiErrors() int32`

GetApiErrors returns the ApiErrors field if non-nil, zero value otherwise.

### GetApiErrorsOk

`func (o *AgentTelemetryBehavior) GetApiErrorsOk() (*int32, bool)`

GetApiErrorsOk returns a tuple with the ApiErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiErrors

`func (o *AgentTelemetryBehavior) SetApiErrors(v int32)`

SetApiErrors sets ApiErrors field to given value.

### HasApiErrors

`func (o *AgentTelemetryBehavior) HasApiErrors() bool`

HasApiErrors returns a boolean if a field has been set.

### GetByTool

`func (o *AgentTelemetryBehavior) GetByTool() []AgentTelemetryToolRow`

GetByTool returns the ByTool field if non-nil, zero value otherwise.

### GetByToolOk

`func (o *AgentTelemetryBehavior) GetByToolOk() (*[]AgentTelemetryToolRow, bool)`

GetByToolOk returns a tuple with the ByTool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetByTool

`func (o *AgentTelemetryBehavior) SetByTool(v []AgentTelemetryToolRow)`

SetByTool sets ByTool field to given value.

### HasByTool

`func (o *AgentTelemetryBehavior) HasByTool() bool`

HasByTool returns a boolean if a field has been set.

### GetSessions

`func (o *AgentTelemetryBehavior) GetSessions() int32`

GetSessions returns the Sessions field if non-nil, zero value otherwise.

### GetSessionsOk

`func (o *AgentTelemetryBehavior) GetSessionsOk() (*int32, bool)`

GetSessionsOk returns a tuple with the Sessions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessions

`func (o *AgentTelemetryBehavior) SetSessions(v int32)`

SetSessions sets Sessions field to given value.

### HasSessions

`func (o *AgentTelemetryBehavior) HasSessions() bool`

HasSessions returns a boolean if a field has been set.

### GetToolAccepts

`func (o *AgentTelemetryBehavior) GetToolAccepts() int32`

GetToolAccepts returns the ToolAccepts field if non-nil, zero value otherwise.

### GetToolAcceptsOk

`func (o *AgentTelemetryBehavior) GetToolAcceptsOk() (*int32, bool)`

GetToolAcceptsOk returns a tuple with the ToolAccepts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToolAccepts

`func (o *AgentTelemetryBehavior) SetToolAccepts(v int32)`

SetToolAccepts sets ToolAccepts field to given value.

### HasToolAccepts

`func (o *AgentTelemetryBehavior) HasToolAccepts() bool`

HasToolAccepts returns a boolean if a field has been set.

### GetToolCalls

`func (o *AgentTelemetryBehavior) GetToolCalls() int32`

GetToolCalls returns the ToolCalls field if non-nil, zero value otherwise.

### GetToolCallsOk

`func (o *AgentTelemetryBehavior) GetToolCallsOk() (*int32, bool)`

GetToolCallsOk returns a tuple with the ToolCalls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToolCalls

`func (o *AgentTelemetryBehavior) SetToolCalls(v int32)`

SetToolCalls sets ToolCalls field to given value.

### HasToolCalls

`func (o *AgentTelemetryBehavior) HasToolCalls() bool`

HasToolCalls returns a boolean if a field has been set.

### GetToolRejects

`func (o *AgentTelemetryBehavior) GetToolRejects() int32`

GetToolRejects returns the ToolRejects field if non-nil, zero value otherwise.

### GetToolRejectsOk

`func (o *AgentTelemetryBehavior) GetToolRejectsOk() (*int32, bool)`

GetToolRejectsOk returns a tuple with the ToolRejects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToolRejects

`func (o *AgentTelemetryBehavior) SetToolRejects(v int32)`

SetToolRejects sets ToolRejects field to given value.

### HasToolRejects

`func (o *AgentTelemetryBehavior) HasToolRejects() bool`

HasToolRejects returns a boolean if a field has been set.

### GetTurns

`func (o *AgentTelemetryBehavior) GetTurns() int32`

GetTurns returns the Turns field if non-nil, zero value otherwise.

### GetTurnsOk

`func (o *AgentTelemetryBehavior) GetTurnsOk() (*int32, bool)`

GetTurnsOk returns a tuple with the Turns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTurns

`func (o *AgentTelemetryBehavior) SetTurns(v int32)`

SetTurns sets Turns field to given value.

### HasTurns

`func (o *AgentTelemetryBehavior) HasTurns() bool`

HasTurns returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


