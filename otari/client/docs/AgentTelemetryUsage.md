# AgentTelemetryUsage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cost** | Pointer to **float32** |  | [optional] [default to 0.0]
**Requests** | Pointer to **int32** |  | [optional] [default to 0]

## Methods

### NewAgentTelemetryUsage

`func NewAgentTelemetryUsage() *AgentTelemetryUsage`

NewAgentTelemetryUsage instantiates a new AgentTelemetryUsage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentTelemetryUsageWithDefaults

`func NewAgentTelemetryUsageWithDefaults() *AgentTelemetryUsage`

NewAgentTelemetryUsageWithDefaults instantiates a new AgentTelemetryUsage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCost

`func (o *AgentTelemetryUsage) GetCost() float32`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *AgentTelemetryUsage) GetCostOk() (*float32, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *AgentTelemetryUsage) SetCost(v float32)`

SetCost sets Cost field to given value.

### HasCost

`func (o *AgentTelemetryUsage) HasCost() bool`

HasCost returns a boolean if a field has been set.

### GetRequests

`func (o *AgentTelemetryUsage) GetRequests() int32`

GetRequests returns the Requests field if non-nil, zero value otherwise.

### GetRequestsOk

`func (o *AgentTelemetryUsage) GetRequestsOk() (*int32, bool)`

GetRequestsOk returns a tuple with the Requests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequests

`func (o *AgentTelemetryUsage) SetRequests(v int32)`

SetRequests sets Requests field to given value.

### HasRequests

`func (o *AgentTelemetryUsage) HasRequests() bool`

HasRequests returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


