# ChatMessageInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Content** | **string** |  | 
**Role** | **string** |  | 
**Name** | **string** |  | 
**Audio** | Pointer to [**MSGAudio**](MSGAudio.md) |  | [optional] 
**FunctionCall** | Pointer to [**MSGFunctionCall**](MSGFunctionCall.md) |  | [optional] 
**Refusal** | Pointer to **string** |  | [optional] 
**ToolCalls** | Pointer to [**[]ToolCallsInner**](ToolCallsInner.md) |  | [optional] 
**ToolCallId** | **string** |  | 

## Methods

### NewChatMessageInput

`func NewChatMessageInput(content string, role string, name string, toolCallId string, ) *ChatMessageInput`

NewChatMessageInput instantiates a new ChatMessageInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChatMessageInputWithDefaults

`func NewChatMessageInputWithDefaults() *ChatMessageInput`

NewChatMessageInputWithDefaults instantiates a new ChatMessageInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContent

`func (o *ChatMessageInput) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *ChatMessageInput) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *ChatMessageInput) SetContent(v string)`

SetContent sets Content field to given value.


### GetRole

`func (o *ChatMessageInput) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *ChatMessageInput) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *ChatMessageInput) SetRole(v string)`

SetRole sets Role field to given value.


### GetName

`func (o *ChatMessageInput) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ChatMessageInput) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ChatMessageInput) SetName(v string)`

SetName sets Name field to given value.


### GetAudio

`func (o *ChatMessageInput) GetAudio() MSGAudio`

GetAudio returns the Audio field if non-nil, zero value otherwise.

### GetAudioOk

`func (o *ChatMessageInput) GetAudioOk() (*MSGAudio, bool)`

GetAudioOk returns a tuple with the Audio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudio

`func (o *ChatMessageInput) SetAudio(v MSGAudio)`

SetAudio sets Audio field to given value.

### HasAudio

`func (o *ChatMessageInput) HasAudio() bool`

HasAudio returns a boolean if a field has been set.

### GetFunctionCall

`func (o *ChatMessageInput) GetFunctionCall() MSGFunctionCall`

GetFunctionCall returns the FunctionCall field if non-nil, zero value otherwise.

### GetFunctionCallOk

`func (o *ChatMessageInput) GetFunctionCallOk() (*MSGFunctionCall, bool)`

GetFunctionCallOk returns a tuple with the FunctionCall field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctionCall

`func (o *ChatMessageInput) SetFunctionCall(v MSGFunctionCall)`

SetFunctionCall sets FunctionCall field to given value.

### HasFunctionCall

`func (o *ChatMessageInput) HasFunctionCall() bool`

HasFunctionCall returns a boolean if a field has been set.

### GetRefusal

`func (o *ChatMessageInput) GetRefusal() string`

GetRefusal returns the Refusal field if non-nil, zero value otherwise.

### GetRefusalOk

`func (o *ChatMessageInput) GetRefusalOk() (*string, bool)`

GetRefusalOk returns a tuple with the Refusal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefusal

`func (o *ChatMessageInput) SetRefusal(v string)`

SetRefusal sets Refusal field to given value.

### HasRefusal

`func (o *ChatMessageInput) HasRefusal() bool`

HasRefusal returns a boolean if a field has been set.

### GetToolCalls

`func (o *ChatMessageInput) GetToolCalls() []ToolCallsInner`

GetToolCalls returns the ToolCalls field if non-nil, zero value otherwise.

### GetToolCallsOk

`func (o *ChatMessageInput) GetToolCallsOk() (*[]ToolCallsInner, bool)`

GetToolCallsOk returns a tuple with the ToolCalls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToolCalls

`func (o *ChatMessageInput) SetToolCalls(v []ToolCallsInner)`

SetToolCalls sets ToolCalls field to given value.

### HasToolCalls

`func (o *ChatMessageInput) HasToolCalls() bool`

HasToolCalls returns a boolean if a field has been set.

### GetToolCallId

`func (o *ChatMessageInput) GetToolCallId() string`

GetToolCallId returns the ToolCallId field if non-nil, zero value otherwise.

### GetToolCallIdOk

`func (o *ChatMessageInput) GetToolCallIdOk() (*string, bool)`

GetToolCallIdOk returns a tuple with the ToolCallId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToolCallId

`func (o *ChatMessageInput) SetToolCallId(v string)`

SetToolCallId sets ToolCallId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


