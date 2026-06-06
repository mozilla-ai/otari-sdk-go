# CCChoiceLogprobs

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Content** | Pointer to [**[]CCChatCompletionTokenLogprob**](CCChatCompletionTokenLogprob.md) |  | [optional] 
**Refusal** | Pointer to [**[]CCChatCompletionTokenLogprob**](CCChatCompletionTokenLogprob.md) |  | [optional] 

## Methods

### NewCCChoiceLogprobs

`func NewCCChoiceLogprobs() *CCChoiceLogprobs`

NewCCChoiceLogprobs instantiates a new CCChoiceLogprobs object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCCChoiceLogprobsWithDefaults

`func NewCCChoiceLogprobsWithDefaults() *CCChoiceLogprobs`

NewCCChoiceLogprobsWithDefaults instantiates a new CCChoiceLogprobs object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContent

`func (o *CCChoiceLogprobs) GetContent() []CCChatCompletionTokenLogprob`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *CCChoiceLogprobs) GetContentOk() (*[]CCChatCompletionTokenLogprob, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *CCChoiceLogprobs) SetContent(v []CCChatCompletionTokenLogprob)`

SetContent sets Content field to given value.

### HasContent

`func (o *CCChoiceLogprobs) HasContent() bool`

HasContent returns a boolean if a field has been set.

### SetContentNil

`func (o *CCChoiceLogprobs) SetContentNil(b bool)`

 SetContentNil sets the value for Content to be an explicit nil

### UnsetContent
`func (o *CCChoiceLogprobs) UnsetContent()`

UnsetContent ensures that no value is present for Content, not even an explicit nil
### GetRefusal

`func (o *CCChoiceLogprobs) GetRefusal() []CCChatCompletionTokenLogprob`

GetRefusal returns the Refusal field if non-nil, zero value otherwise.

### GetRefusalOk

`func (o *CCChoiceLogprobs) GetRefusalOk() (*[]CCChatCompletionTokenLogprob, bool)`

GetRefusalOk returns a tuple with the Refusal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefusal

`func (o *CCChoiceLogprobs) SetRefusal(v []CCChatCompletionTokenLogprob)`

SetRefusal sets Refusal field to given value.

### HasRefusal

`func (o *CCChoiceLogprobs) HasRefusal() bool`

HasRefusal returns a boolean if a field has been set.

### SetRefusalNil

`func (o *CCChoiceLogprobs) SetRefusalNil(b bool)`

 SetRefusalNil sets the value for Refusal to be an explicit nil

### UnsetRefusal
`func (o *CCChoiceLogprobs) UnsetRefusal()`

UnsetRefusal ensures that no value is present for Refusal, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


