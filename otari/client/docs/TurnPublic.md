# TurnPublic

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CacheReadTokens** | **int32** |  | 
**Continued** | **bool** | True for steps a session started with before any prompt was seen. | 
**Cost** | **float32** |  | 
**EndedAt** | **NullableTime** |  | 
**Errors** | **int32** |  | 
**Index** | **int32** |  | 
**InputTokens** | **int32** |  | 
**LlmCalls** | **int32** |  | 
**OutputTokens** | **int32** |  | 
**StartedAt** | **NullableTime** |  | 
**State** | **string** |  | 
**StepIds** | **[]string** |  | 
**ToolCalls** | **int32** |  | 

## Methods

### NewTurnPublic

`func NewTurnPublic(cacheReadTokens int32, continued bool, cost float32, endedAt NullableTime, errors int32, index int32, inputTokens int32, llmCalls int32, outputTokens int32, startedAt NullableTime, state string, stepIds []string, toolCalls int32, ) *TurnPublic`

NewTurnPublic instantiates a new TurnPublic object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTurnPublicWithDefaults

`func NewTurnPublicWithDefaults() *TurnPublic`

NewTurnPublicWithDefaults instantiates a new TurnPublic object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCacheReadTokens

`func (o *TurnPublic) GetCacheReadTokens() int32`

GetCacheReadTokens returns the CacheReadTokens field if non-nil, zero value otherwise.

### GetCacheReadTokensOk

`func (o *TurnPublic) GetCacheReadTokensOk() (*int32, bool)`

GetCacheReadTokensOk returns a tuple with the CacheReadTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheReadTokens

`func (o *TurnPublic) SetCacheReadTokens(v int32)`

SetCacheReadTokens sets CacheReadTokens field to given value.


### GetContinued

`func (o *TurnPublic) GetContinued() bool`

GetContinued returns the Continued field if non-nil, zero value otherwise.

### GetContinuedOk

`func (o *TurnPublic) GetContinuedOk() (*bool, bool)`

GetContinuedOk returns a tuple with the Continued field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContinued

`func (o *TurnPublic) SetContinued(v bool)`

SetContinued sets Continued field to given value.


### GetCost

`func (o *TurnPublic) GetCost() float32`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *TurnPublic) GetCostOk() (*float32, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *TurnPublic) SetCost(v float32)`

SetCost sets Cost field to given value.


### GetEndedAt

`func (o *TurnPublic) GetEndedAt() time.Time`

GetEndedAt returns the EndedAt field if non-nil, zero value otherwise.

### GetEndedAtOk

`func (o *TurnPublic) GetEndedAtOk() (*time.Time, bool)`

GetEndedAtOk returns a tuple with the EndedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndedAt

`func (o *TurnPublic) SetEndedAt(v time.Time)`

SetEndedAt sets EndedAt field to given value.


### SetEndedAtNil

`func (o *TurnPublic) SetEndedAtNil(b bool)`

 SetEndedAtNil sets the value for EndedAt to be an explicit nil

### UnsetEndedAt
`func (o *TurnPublic) UnsetEndedAt()`

UnsetEndedAt ensures that no value is present for EndedAt, not even an explicit nil
### GetErrors

`func (o *TurnPublic) GetErrors() int32`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *TurnPublic) GetErrorsOk() (*int32, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *TurnPublic) SetErrors(v int32)`

SetErrors sets Errors field to given value.


### GetIndex

`func (o *TurnPublic) GetIndex() int32`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *TurnPublic) GetIndexOk() (*int32, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *TurnPublic) SetIndex(v int32)`

SetIndex sets Index field to given value.


### GetInputTokens

`func (o *TurnPublic) GetInputTokens() int32`

GetInputTokens returns the InputTokens field if non-nil, zero value otherwise.

### GetInputTokensOk

`func (o *TurnPublic) GetInputTokensOk() (*int32, bool)`

GetInputTokensOk returns a tuple with the InputTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputTokens

`func (o *TurnPublic) SetInputTokens(v int32)`

SetInputTokens sets InputTokens field to given value.


### GetLlmCalls

`func (o *TurnPublic) GetLlmCalls() int32`

GetLlmCalls returns the LlmCalls field if non-nil, zero value otherwise.

### GetLlmCallsOk

`func (o *TurnPublic) GetLlmCallsOk() (*int32, bool)`

GetLlmCallsOk returns a tuple with the LlmCalls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLlmCalls

`func (o *TurnPublic) SetLlmCalls(v int32)`

SetLlmCalls sets LlmCalls field to given value.


### GetOutputTokens

`func (o *TurnPublic) GetOutputTokens() int32`

GetOutputTokens returns the OutputTokens field if non-nil, zero value otherwise.

### GetOutputTokensOk

`func (o *TurnPublic) GetOutputTokensOk() (*int32, bool)`

GetOutputTokensOk returns a tuple with the OutputTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputTokens

`func (o *TurnPublic) SetOutputTokens(v int32)`

SetOutputTokens sets OutputTokens field to given value.


### GetStartedAt

`func (o *TurnPublic) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *TurnPublic) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *TurnPublic) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.


### SetStartedAtNil

`func (o *TurnPublic) SetStartedAtNil(b bool)`

 SetStartedAtNil sets the value for StartedAt to be an explicit nil

### UnsetStartedAt
`func (o *TurnPublic) UnsetStartedAt()`

UnsetStartedAt ensures that no value is present for StartedAt, not even an explicit nil
### GetState

`func (o *TurnPublic) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *TurnPublic) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *TurnPublic) SetState(v string)`

SetState sets State field to given value.


### GetStepIds

`func (o *TurnPublic) GetStepIds() []string`

GetStepIds returns the StepIds field if non-nil, zero value otherwise.

### GetStepIdsOk

`func (o *TurnPublic) GetStepIdsOk() (*[]string, bool)`

GetStepIdsOk returns a tuple with the StepIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStepIds

`func (o *TurnPublic) SetStepIds(v []string)`

SetStepIds sets StepIds field to given value.


### GetToolCalls

`func (o *TurnPublic) GetToolCalls() int32`

GetToolCalls returns the ToolCalls field if non-nil, zero value otherwise.

### GetToolCallsOk

`func (o *TurnPublic) GetToolCallsOk() (*int32, bool)`

GetToolCallsOk returns a tuple with the ToolCalls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToolCalls

`func (o *TurnPublic) SetToolCalls(v int32)`

SetToolCalls sets ToolCalls field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


