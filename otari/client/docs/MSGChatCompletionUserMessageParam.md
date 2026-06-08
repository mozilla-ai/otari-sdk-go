# MSGChatCompletionUserMessageParam

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Content** | [**Content2**](Content2.md) |  | 
**Role** | **string** |  | 
**Name** | Pointer to **string** |  | [optional] 

## Methods

### NewMSGChatCompletionUserMessageParam

`func NewMSGChatCompletionUserMessageParam(content Content2, role string, ) *MSGChatCompletionUserMessageParam`

NewMSGChatCompletionUserMessageParam instantiates a new MSGChatCompletionUserMessageParam object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMSGChatCompletionUserMessageParamWithDefaults

`func NewMSGChatCompletionUserMessageParamWithDefaults() *MSGChatCompletionUserMessageParam`

NewMSGChatCompletionUserMessageParamWithDefaults instantiates a new MSGChatCompletionUserMessageParam object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContent

`func (o *MSGChatCompletionUserMessageParam) GetContent() Content2`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *MSGChatCompletionUserMessageParam) GetContentOk() (*Content2, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *MSGChatCompletionUserMessageParam) SetContent(v Content2)`

SetContent sets Content field to given value.


### GetRole

`func (o *MSGChatCompletionUserMessageParam) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *MSGChatCompletionUserMessageParam) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *MSGChatCompletionUserMessageParam) SetRole(v string)`

SetRole sets Role field to given value.


### GetName

`func (o *MSGChatCompletionUserMessageParam) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MSGChatCompletionUserMessageParam) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MSGChatCompletionUserMessageParam) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MSGChatCompletionUserMessageParam) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


