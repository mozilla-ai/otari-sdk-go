# MSGChatCompletionSystemMessageParam

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Content** | [**Content1**](Content1.md) |  | 
**Role** | **string** |  | 
**Name** | Pointer to **string** |  | [optional] 

## Methods

### NewMSGChatCompletionSystemMessageParam

`func NewMSGChatCompletionSystemMessageParam(content Content1, role string, ) *MSGChatCompletionSystemMessageParam`

NewMSGChatCompletionSystemMessageParam instantiates a new MSGChatCompletionSystemMessageParam object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMSGChatCompletionSystemMessageParamWithDefaults

`func NewMSGChatCompletionSystemMessageParamWithDefaults() *MSGChatCompletionSystemMessageParam`

NewMSGChatCompletionSystemMessageParamWithDefaults instantiates a new MSGChatCompletionSystemMessageParam object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContent

`func (o *MSGChatCompletionSystemMessageParam) GetContent() Content1`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *MSGChatCompletionSystemMessageParam) GetContentOk() (*Content1, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *MSGChatCompletionSystemMessageParam) SetContent(v Content1)`

SetContent sets Content field to given value.


### GetRole

`func (o *MSGChatCompletionSystemMessageParam) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *MSGChatCompletionSystemMessageParam) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *MSGChatCompletionSystemMessageParam) SetRole(v string)`

SetRole sets Role field to given value.


### GetName

`func (o *MSGChatCompletionSystemMessageParam) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MSGChatCompletionSystemMessageParam) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MSGChatCompletionSystemMessageParam) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MSGChatCompletionSystemMessageParam) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


