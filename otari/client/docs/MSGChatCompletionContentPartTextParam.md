# MSGChatCompletionContentPartTextParam

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Text** | **string** |  | 
**Type** | **string** |  | 
**PromptCacheBreakpoint** | Pointer to [**MSGPromptCacheBreakpoint**](MSGPromptCacheBreakpoint.md) |  | [optional] 

## Methods

### NewMSGChatCompletionContentPartTextParam

`func NewMSGChatCompletionContentPartTextParam(text string, type_ string, ) *MSGChatCompletionContentPartTextParam`

NewMSGChatCompletionContentPartTextParam instantiates a new MSGChatCompletionContentPartTextParam object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMSGChatCompletionContentPartTextParamWithDefaults

`func NewMSGChatCompletionContentPartTextParamWithDefaults() *MSGChatCompletionContentPartTextParam`

NewMSGChatCompletionContentPartTextParamWithDefaults instantiates a new MSGChatCompletionContentPartTextParam object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetText

`func (o *MSGChatCompletionContentPartTextParam) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *MSGChatCompletionContentPartTextParam) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *MSGChatCompletionContentPartTextParam) SetText(v string)`

SetText sets Text field to given value.


### GetType

`func (o *MSGChatCompletionContentPartTextParam) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MSGChatCompletionContentPartTextParam) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MSGChatCompletionContentPartTextParam) SetType(v string)`

SetType sets Type field to given value.


### GetPromptCacheBreakpoint

`func (o *MSGChatCompletionContentPartTextParam) GetPromptCacheBreakpoint() MSGPromptCacheBreakpoint`

GetPromptCacheBreakpoint returns the PromptCacheBreakpoint field if non-nil, zero value otherwise.

### GetPromptCacheBreakpointOk

`func (o *MSGChatCompletionContentPartTextParam) GetPromptCacheBreakpointOk() (*MSGPromptCacheBreakpoint, bool)`

GetPromptCacheBreakpointOk returns a tuple with the PromptCacheBreakpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromptCacheBreakpoint

`func (o *MSGChatCompletionContentPartTextParam) SetPromptCacheBreakpoint(v MSGPromptCacheBreakpoint)`

SetPromptCacheBreakpoint sets PromptCacheBreakpoint field to given value.

### HasPromptCacheBreakpoint

`func (o *MSGChatCompletionContentPartTextParam) HasPromptCacheBreakpoint() bool`

HasPromptCacheBreakpoint returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


