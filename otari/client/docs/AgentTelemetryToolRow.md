# AgentTelemetryToolRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Calls** | **int32** |  | 
**Tool** | **NullableString** |  | 

## Methods

### NewAgentTelemetryToolRow

`func NewAgentTelemetryToolRow(calls int32, tool NullableString, ) *AgentTelemetryToolRow`

NewAgentTelemetryToolRow instantiates a new AgentTelemetryToolRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentTelemetryToolRowWithDefaults

`func NewAgentTelemetryToolRowWithDefaults() *AgentTelemetryToolRow`

NewAgentTelemetryToolRowWithDefaults instantiates a new AgentTelemetryToolRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCalls

`func (o *AgentTelemetryToolRow) GetCalls() int32`

GetCalls returns the Calls field if non-nil, zero value otherwise.

### GetCallsOk

`func (o *AgentTelemetryToolRow) GetCallsOk() (*int32, bool)`

GetCallsOk returns a tuple with the Calls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCalls

`func (o *AgentTelemetryToolRow) SetCalls(v int32)`

SetCalls sets Calls field to given value.


### GetTool

`func (o *AgentTelemetryToolRow) GetTool() string`

GetTool returns the Tool field if non-nil, zero value otherwise.

### GetToolOk

`func (o *AgentTelemetryToolRow) GetToolOk() (*string, bool)`

GetToolOk returns a tuple with the Tool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTool

`func (o *AgentTelemetryToolRow) SetTool(v string)`

SetTool sets Tool field to given value.


### SetToolNil

`func (o *AgentTelemetryToolRow) SetToolNil(b bool)`

 SetToolNil sets the value for Tool to be an explicit nil

### UnsetTool
`func (o *AgentTelemetryToolRow) UnsetTool()`

UnsetTool ensures that no value is present for Tool, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


