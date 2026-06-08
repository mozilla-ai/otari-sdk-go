# CCChatCompletionMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Content** | Pointer to **NullableString** |  | [optional] 
**Refusal** | Pointer to **NullableString** |  | [optional] 
**Role** | **string** |  | 
**Annotations** | Pointer to **[]map[string]interface{}** |  | [optional] 
**Audio** | Pointer to [**NullableCCChatCompletionAudio**](CCChatCompletionAudio.md) |  | [optional] 
**FunctionCall** | Pointer to [**NullableCCFunctionCall**](CCFunctionCall.md) |  | [optional] 
**ToolCalls** | Pointer to [**[]CCChatCompletionMessageToolCallsInner**](CCChatCompletionMessageToolCallsInner.md) |  | [optional] 
**Reasoning** | Pointer to [**NullableCCKReasoning**](CCKReasoning.md) |  | [optional] 

## Methods

### NewCCChatCompletionMessage

`func NewCCChatCompletionMessage(role string, ) *CCChatCompletionMessage`

NewCCChatCompletionMessage instantiates a new CCChatCompletionMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCCChatCompletionMessageWithDefaults

`func NewCCChatCompletionMessageWithDefaults() *CCChatCompletionMessage`

NewCCChatCompletionMessageWithDefaults instantiates a new CCChatCompletionMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContent

`func (o *CCChatCompletionMessage) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *CCChatCompletionMessage) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *CCChatCompletionMessage) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *CCChatCompletionMessage) HasContent() bool`

HasContent returns a boolean if a field has been set.

### SetContentNil

`func (o *CCChatCompletionMessage) SetContentNil(b bool)`

 SetContentNil sets the value for Content to be an explicit nil

### UnsetContent
`func (o *CCChatCompletionMessage) UnsetContent()`

UnsetContent ensures that no value is present for Content, not even an explicit nil
### GetRefusal

`func (o *CCChatCompletionMessage) GetRefusal() string`

GetRefusal returns the Refusal field if non-nil, zero value otherwise.

### GetRefusalOk

`func (o *CCChatCompletionMessage) GetRefusalOk() (*string, bool)`

GetRefusalOk returns a tuple with the Refusal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefusal

`func (o *CCChatCompletionMessage) SetRefusal(v string)`

SetRefusal sets Refusal field to given value.

### HasRefusal

`func (o *CCChatCompletionMessage) HasRefusal() bool`

HasRefusal returns a boolean if a field has been set.

### SetRefusalNil

`func (o *CCChatCompletionMessage) SetRefusalNil(b bool)`

 SetRefusalNil sets the value for Refusal to be an explicit nil

### UnsetRefusal
`func (o *CCChatCompletionMessage) UnsetRefusal()`

UnsetRefusal ensures that no value is present for Refusal, not even an explicit nil
### GetRole

`func (o *CCChatCompletionMessage) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *CCChatCompletionMessage) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *CCChatCompletionMessage) SetRole(v string)`

SetRole sets Role field to given value.


### GetAnnotations

`func (o *CCChatCompletionMessage) GetAnnotations() []map[string]interface{}`

GetAnnotations returns the Annotations field if non-nil, zero value otherwise.

### GetAnnotationsOk

`func (o *CCChatCompletionMessage) GetAnnotationsOk() (*[]map[string]interface{}, bool)`

GetAnnotationsOk returns a tuple with the Annotations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotations

`func (o *CCChatCompletionMessage) SetAnnotations(v []map[string]interface{})`

SetAnnotations sets Annotations field to given value.

### HasAnnotations

`func (o *CCChatCompletionMessage) HasAnnotations() bool`

HasAnnotations returns a boolean if a field has been set.

### SetAnnotationsNil

`func (o *CCChatCompletionMessage) SetAnnotationsNil(b bool)`

 SetAnnotationsNil sets the value for Annotations to be an explicit nil

### UnsetAnnotations
`func (o *CCChatCompletionMessage) UnsetAnnotations()`

UnsetAnnotations ensures that no value is present for Annotations, not even an explicit nil
### GetAudio

`func (o *CCChatCompletionMessage) GetAudio() CCChatCompletionAudio`

GetAudio returns the Audio field if non-nil, zero value otherwise.

### GetAudioOk

`func (o *CCChatCompletionMessage) GetAudioOk() (*CCChatCompletionAudio, bool)`

GetAudioOk returns a tuple with the Audio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudio

`func (o *CCChatCompletionMessage) SetAudio(v CCChatCompletionAudio)`

SetAudio sets Audio field to given value.

### HasAudio

`func (o *CCChatCompletionMessage) HasAudio() bool`

HasAudio returns a boolean if a field has been set.

### SetAudioNil

`func (o *CCChatCompletionMessage) SetAudioNil(b bool)`

 SetAudioNil sets the value for Audio to be an explicit nil

### UnsetAudio
`func (o *CCChatCompletionMessage) UnsetAudio()`

UnsetAudio ensures that no value is present for Audio, not even an explicit nil
### GetFunctionCall

`func (o *CCChatCompletionMessage) GetFunctionCall() CCFunctionCall`

GetFunctionCall returns the FunctionCall field if non-nil, zero value otherwise.

### GetFunctionCallOk

`func (o *CCChatCompletionMessage) GetFunctionCallOk() (*CCFunctionCall, bool)`

GetFunctionCallOk returns a tuple with the FunctionCall field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctionCall

`func (o *CCChatCompletionMessage) SetFunctionCall(v CCFunctionCall)`

SetFunctionCall sets FunctionCall field to given value.

### HasFunctionCall

`func (o *CCChatCompletionMessage) HasFunctionCall() bool`

HasFunctionCall returns a boolean if a field has been set.

### SetFunctionCallNil

`func (o *CCChatCompletionMessage) SetFunctionCallNil(b bool)`

 SetFunctionCallNil sets the value for FunctionCall to be an explicit nil

### UnsetFunctionCall
`func (o *CCChatCompletionMessage) UnsetFunctionCall()`

UnsetFunctionCall ensures that no value is present for FunctionCall, not even an explicit nil
### GetToolCalls

`func (o *CCChatCompletionMessage) GetToolCalls() []CCChatCompletionMessageToolCallsInner`

GetToolCalls returns the ToolCalls field if non-nil, zero value otherwise.

### GetToolCallsOk

`func (o *CCChatCompletionMessage) GetToolCallsOk() (*[]CCChatCompletionMessageToolCallsInner, bool)`

GetToolCallsOk returns a tuple with the ToolCalls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToolCalls

`func (o *CCChatCompletionMessage) SetToolCalls(v []CCChatCompletionMessageToolCallsInner)`

SetToolCalls sets ToolCalls field to given value.

### HasToolCalls

`func (o *CCChatCompletionMessage) HasToolCalls() bool`

HasToolCalls returns a boolean if a field has been set.

### SetToolCallsNil

`func (o *CCChatCompletionMessage) SetToolCallsNil(b bool)`

 SetToolCallsNil sets the value for ToolCalls to be an explicit nil

### UnsetToolCalls
`func (o *CCChatCompletionMessage) UnsetToolCalls()`

UnsetToolCalls ensures that no value is present for ToolCalls, not even an explicit nil
### GetReasoning

`func (o *CCChatCompletionMessage) GetReasoning() CCKReasoning`

GetReasoning returns the Reasoning field if non-nil, zero value otherwise.

### GetReasoningOk

`func (o *CCChatCompletionMessage) GetReasoningOk() (*CCKReasoning, bool)`

GetReasoningOk returns a tuple with the Reasoning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasoning

`func (o *CCChatCompletionMessage) SetReasoning(v CCKReasoning)`

SetReasoning sets Reasoning field to given value.

### HasReasoning

`func (o *CCChatCompletionMessage) HasReasoning() bool`

HasReasoning returns a boolean if a field has been set.

### SetReasoningNil

`func (o *CCChatCompletionMessage) SetReasoningNil(b bool)`

 SetReasoningNil sets the value for Reasoning to be an explicit nil

### UnsetReasoning
`func (o *CCChatCompletionMessage) UnsetReasoning()`

UnsetReasoning ensures that no value is present for Reasoning, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


