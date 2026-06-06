# MSGChatCompletionToolMessageParam

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Content** | [**Content1**](Content1.md) |  | 
**Role** | **string** |  | 
**ToolCallId** | **string** |  | 

## Methods

### NewMSGChatCompletionToolMessageParam

`func NewMSGChatCompletionToolMessageParam(content Content1, role string, toolCallId string, ) *MSGChatCompletionToolMessageParam`

NewMSGChatCompletionToolMessageParam instantiates a new MSGChatCompletionToolMessageParam object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMSGChatCompletionToolMessageParamWithDefaults

`func NewMSGChatCompletionToolMessageParamWithDefaults() *MSGChatCompletionToolMessageParam`

NewMSGChatCompletionToolMessageParamWithDefaults instantiates a new MSGChatCompletionToolMessageParam object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContent

`func (o *MSGChatCompletionToolMessageParam) GetContent() Content1`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *MSGChatCompletionToolMessageParam) GetContentOk() (*Content1, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *MSGChatCompletionToolMessageParam) SetContent(v Content1)`

SetContent sets Content field to given value.


### GetRole

`func (o *MSGChatCompletionToolMessageParam) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *MSGChatCompletionToolMessageParam) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *MSGChatCompletionToolMessageParam) SetRole(v string)`

SetRole sets Role field to given value.


### GetToolCallId

`func (o *MSGChatCompletionToolMessageParam) GetToolCallId() string`

GetToolCallId returns the ToolCallId field if non-nil, zero value otherwise.

### GetToolCallIdOk

`func (o *MSGChatCompletionToolMessageParam) GetToolCallIdOk() (*string, bool)`

GetToolCallIdOk returns a tuple with the ToolCallId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToolCallId

`func (o *MSGChatCompletionToolMessageParam) SetToolCallId(v string)`

SetToolCallId sets ToolCallId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


