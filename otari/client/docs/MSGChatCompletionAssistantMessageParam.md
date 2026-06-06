# MSGChatCompletionAssistantMessageParam

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Role** | **string** |  | 
**Audio** | Pointer to [**NullableMSGAudio**](MSGAudio.md) |  | [optional] 
**Content** | Pointer to [**NullableContent**](Content.md) |  | [optional] 
**FunctionCall** | Pointer to [**NullableMSGFunctionCall**](MSGFunctionCall.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Refusal** | Pointer to **NullableString** |  | [optional] 
**ToolCalls** | Pointer to [**[]ToolCallsInner**](ToolCallsInner.md) |  | [optional] 

## Methods

### NewMSGChatCompletionAssistantMessageParam

`func NewMSGChatCompletionAssistantMessageParam(role string, ) *MSGChatCompletionAssistantMessageParam`

NewMSGChatCompletionAssistantMessageParam instantiates a new MSGChatCompletionAssistantMessageParam object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMSGChatCompletionAssistantMessageParamWithDefaults

`func NewMSGChatCompletionAssistantMessageParamWithDefaults() *MSGChatCompletionAssistantMessageParam`

NewMSGChatCompletionAssistantMessageParamWithDefaults instantiates a new MSGChatCompletionAssistantMessageParam object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRole

`func (o *MSGChatCompletionAssistantMessageParam) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *MSGChatCompletionAssistantMessageParam) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *MSGChatCompletionAssistantMessageParam) SetRole(v string)`

SetRole sets Role field to given value.


### GetAudio

`func (o *MSGChatCompletionAssistantMessageParam) GetAudio() MSGAudio`

GetAudio returns the Audio field if non-nil, zero value otherwise.

### GetAudioOk

`func (o *MSGChatCompletionAssistantMessageParam) GetAudioOk() (*MSGAudio, bool)`

GetAudioOk returns a tuple with the Audio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudio

`func (o *MSGChatCompletionAssistantMessageParam) SetAudio(v MSGAudio)`

SetAudio sets Audio field to given value.

### HasAudio

`func (o *MSGChatCompletionAssistantMessageParam) HasAudio() bool`

HasAudio returns a boolean if a field has been set.

### SetAudioNil

`func (o *MSGChatCompletionAssistantMessageParam) SetAudioNil(b bool)`

 SetAudioNil sets the value for Audio to be an explicit nil

### UnsetAudio
`func (o *MSGChatCompletionAssistantMessageParam) UnsetAudio()`

UnsetAudio ensures that no value is present for Audio, not even an explicit nil
### GetContent

`func (o *MSGChatCompletionAssistantMessageParam) GetContent() Content`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *MSGChatCompletionAssistantMessageParam) GetContentOk() (*Content, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *MSGChatCompletionAssistantMessageParam) SetContent(v Content)`

SetContent sets Content field to given value.

### HasContent

`func (o *MSGChatCompletionAssistantMessageParam) HasContent() bool`

HasContent returns a boolean if a field has been set.

### SetContentNil

`func (o *MSGChatCompletionAssistantMessageParam) SetContentNil(b bool)`

 SetContentNil sets the value for Content to be an explicit nil

### UnsetContent
`func (o *MSGChatCompletionAssistantMessageParam) UnsetContent()`

UnsetContent ensures that no value is present for Content, not even an explicit nil
### GetFunctionCall

`func (o *MSGChatCompletionAssistantMessageParam) GetFunctionCall() MSGFunctionCall`

GetFunctionCall returns the FunctionCall field if non-nil, zero value otherwise.

### GetFunctionCallOk

`func (o *MSGChatCompletionAssistantMessageParam) GetFunctionCallOk() (*MSGFunctionCall, bool)`

GetFunctionCallOk returns a tuple with the FunctionCall field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctionCall

`func (o *MSGChatCompletionAssistantMessageParam) SetFunctionCall(v MSGFunctionCall)`

SetFunctionCall sets FunctionCall field to given value.

### HasFunctionCall

`func (o *MSGChatCompletionAssistantMessageParam) HasFunctionCall() bool`

HasFunctionCall returns a boolean if a field has been set.

### SetFunctionCallNil

`func (o *MSGChatCompletionAssistantMessageParam) SetFunctionCallNil(b bool)`

 SetFunctionCallNil sets the value for FunctionCall to be an explicit nil

### UnsetFunctionCall
`func (o *MSGChatCompletionAssistantMessageParam) UnsetFunctionCall()`

UnsetFunctionCall ensures that no value is present for FunctionCall, not even an explicit nil
### GetName

`func (o *MSGChatCompletionAssistantMessageParam) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MSGChatCompletionAssistantMessageParam) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MSGChatCompletionAssistantMessageParam) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MSGChatCompletionAssistantMessageParam) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRefusal

`func (o *MSGChatCompletionAssistantMessageParam) GetRefusal() string`

GetRefusal returns the Refusal field if non-nil, zero value otherwise.

### GetRefusalOk

`func (o *MSGChatCompletionAssistantMessageParam) GetRefusalOk() (*string, bool)`

GetRefusalOk returns a tuple with the Refusal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefusal

`func (o *MSGChatCompletionAssistantMessageParam) SetRefusal(v string)`

SetRefusal sets Refusal field to given value.

### HasRefusal

`func (o *MSGChatCompletionAssistantMessageParam) HasRefusal() bool`

HasRefusal returns a boolean if a field has been set.

### SetRefusalNil

`func (o *MSGChatCompletionAssistantMessageParam) SetRefusalNil(b bool)`

 SetRefusalNil sets the value for Refusal to be an explicit nil

### UnsetRefusal
`func (o *MSGChatCompletionAssistantMessageParam) UnsetRefusal()`

UnsetRefusal ensures that no value is present for Refusal, not even an explicit nil
### GetToolCalls

`func (o *MSGChatCompletionAssistantMessageParam) GetToolCalls() []ToolCallsInner`

GetToolCalls returns the ToolCalls field if non-nil, zero value otherwise.

### GetToolCallsOk

`func (o *MSGChatCompletionAssistantMessageParam) GetToolCallsOk() (*[]ToolCallsInner, bool)`

GetToolCallsOk returns a tuple with the ToolCalls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToolCalls

`func (o *MSGChatCompletionAssistantMessageParam) SetToolCalls(v []ToolCallsInner)`

SetToolCalls sets ToolCalls field to given value.

### HasToolCalls

`func (o *MSGChatCompletionAssistantMessageParam) HasToolCalls() bool`

HasToolCalls returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


