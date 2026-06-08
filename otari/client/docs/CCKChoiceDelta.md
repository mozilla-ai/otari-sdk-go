# CCKChoiceDelta

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Content** | Pointer to **NullableString** |  | [optional] 
**FunctionCall** | Pointer to [**NullableCCKChoiceDeltaFunctionCall**](CCKChoiceDeltaFunctionCall.md) |  | [optional] 
**Refusal** | Pointer to **NullableString** |  | [optional] 
**Role** | Pointer to **NullableString** |  | [optional] 
**ToolCalls** | Pointer to [**[]CCKChoiceDeltaToolCall**](CCKChoiceDeltaToolCall.md) |  | [optional] 
**Reasoning** | Pointer to [**NullableCCKReasoning**](CCKReasoning.md) |  | [optional] 

## Methods

### NewCCKChoiceDelta

`func NewCCKChoiceDelta() *CCKChoiceDelta`

NewCCKChoiceDelta instantiates a new CCKChoiceDelta object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCCKChoiceDeltaWithDefaults

`func NewCCKChoiceDeltaWithDefaults() *CCKChoiceDelta`

NewCCKChoiceDeltaWithDefaults instantiates a new CCKChoiceDelta object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContent

`func (o *CCKChoiceDelta) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *CCKChoiceDelta) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *CCKChoiceDelta) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *CCKChoiceDelta) HasContent() bool`

HasContent returns a boolean if a field has been set.

### SetContentNil

`func (o *CCKChoiceDelta) SetContentNil(b bool)`

 SetContentNil sets the value for Content to be an explicit nil

### UnsetContent
`func (o *CCKChoiceDelta) UnsetContent()`

UnsetContent ensures that no value is present for Content, not even an explicit nil
### GetFunctionCall

`func (o *CCKChoiceDelta) GetFunctionCall() CCKChoiceDeltaFunctionCall`

GetFunctionCall returns the FunctionCall field if non-nil, zero value otherwise.

### GetFunctionCallOk

`func (o *CCKChoiceDelta) GetFunctionCallOk() (*CCKChoiceDeltaFunctionCall, bool)`

GetFunctionCallOk returns a tuple with the FunctionCall field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctionCall

`func (o *CCKChoiceDelta) SetFunctionCall(v CCKChoiceDeltaFunctionCall)`

SetFunctionCall sets FunctionCall field to given value.

### HasFunctionCall

`func (o *CCKChoiceDelta) HasFunctionCall() bool`

HasFunctionCall returns a boolean if a field has been set.

### SetFunctionCallNil

`func (o *CCKChoiceDelta) SetFunctionCallNil(b bool)`

 SetFunctionCallNil sets the value for FunctionCall to be an explicit nil

### UnsetFunctionCall
`func (o *CCKChoiceDelta) UnsetFunctionCall()`

UnsetFunctionCall ensures that no value is present for FunctionCall, not even an explicit nil
### GetRefusal

`func (o *CCKChoiceDelta) GetRefusal() string`

GetRefusal returns the Refusal field if non-nil, zero value otherwise.

### GetRefusalOk

`func (o *CCKChoiceDelta) GetRefusalOk() (*string, bool)`

GetRefusalOk returns a tuple with the Refusal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefusal

`func (o *CCKChoiceDelta) SetRefusal(v string)`

SetRefusal sets Refusal field to given value.

### HasRefusal

`func (o *CCKChoiceDelta) HasRefusal() bool`

HasRefusal returns a boolean if a field has been set.

### SetRefusalNil

`func (o *CCKChoiceDelta) SetRefusalNil(b bool)`

 SetRefusalNil sets the value for Refusal to be an explicit nil

### UnsetRefusal
`func (o *CCKChoiceDelta) UnsetRefusal()`

UnsetRefusal ensures that no value is present for Refusal, not even an explicit nil
### GetRole

`func (o *CCKChoiceDelta) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *CCKChoiceDelta) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *CCKChoiceDelta) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *CCKChoiceDelta) HasRole() bool`

HasRole returns a boolean if a field has been set.

### SetRoleNil

`func (o *CCKChoiceDelta) SetRoleNil(b bool)`

 SetRoleNil sets the value for Role to be an explicit nil

### UnsetRole
`func (o *CCKChoiceDelta) UnsetRole()`

UnsetRole ensures that no value is present for Role, not even an explicit nil
### GetToolCalls

`func (o *CCKChoiceDelta) GetToolCalls() []CCKChoiceDeltaToolCall`

GetToolCalls returns the ToolCalls field if non-nil, zero value otherwise.

### GetToolCallsOk

`func (o *CCKChoiceDelta) GetToolCallsOk() (*[]CCKChoiceDeltaToolCall, bool)`

GetToolCallsOk returns a tuple with the ToolCalls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToolCalls

`func (o *CCKChoiceDelta) SetToolCalls(v []CCKChoiceDeltaToolCall)`

SetToolCalls sets ToolCalls field to given value.

### HasToolCalls

`func (o *CCKChoiceDelta) HasToolCalls() bool`

HasToolCalls returns a boolean if a field has been set.

### SetToolCallsNil

`func (o *CCKChoiceDelta) SetToolCallsNil(b bool)`

 SetToolCallsNil sets the value for ToolCalls to be an explicit nil

### UnsetToolCalls
`func (o *CCKChoiceDelta) UnsetToolCalls()`

UnsetToolCalls ensures that no value is present for ToolCalls, not even an explicit nil
### GetReasoning

`func (o *CCKChoiceDelta) GetReasoning() CCKReasoning`

GetReasoning returns the Reasoning field if non-nil, zero value otherwise.

### GetReasoningOk

`func (o *CCKChoiceDelta) GetReasoningOk() (*CCKReasoning, bool)`

GetReasoningOk returns a tuple with the Reasoning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasoning

`func (o *CCKChoiceDelta) SetReasoning(v CCKReasoning)`

SetReasoning sets Reasoning field to given value.

### HasReasoning

`func (o *CCKChoiceDelta) HasReasoning() bool`

HasReasoning returns a boolean if a field has been set.

### SetReasoningNil

`func (o *CCKChoiceDelta) SetReasoningNil(b bool)`

 SetReasoningNil sets the value for Reasoning to be an explicit nil

### UnsetReasoning
`func (o *CCKChoiceDelta) UnsetReasoning()`

UnsetReasoning ensures that no value is present for Reasoning, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


