# CCKChoiceLogprobs

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Content** | Pointer to [**[]CCKChatCompletionTokenLogprob**](CCKChatCompletionTokenLogprob.md) |  | [optional] 
**Refusal** | Pointer to [**[]CCKChatCompletionTokenLogprob**](CCKChatCompletionTokenLogprob.md) |  | [optional] 

## Methods

### NewCCKChoiceLogprobs

`func NewCCKChoiceLogprobs() *CCKChoiceLogprobs`

NewCCKChoiceLogprobs instantiates a new CCKChoiceLogprobs object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCCKChoiceLogprobsWithDefaults

`func NewCCKChoiceLogprobsWithDefaults() *CCKChoiceLogprobs`

NewCCKChoiceLogprobsWithDefaults instantiates a new CCKChoiceLogprobs object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContent

`func (o *CCKChoiceLogprobs) GetContent() []CCKChatCompletionTokenLogprob`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *CCKChoiceLogprobs) GetContentOk() (*[]CCKChatCompletionTokenLogprob, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *CCKChoiceLogprobs) SetContent(v []CCKChatCompletionTokenLogprob)`

SetContent sets Content field to given value.

### HasContent

`func (o *CCKChoiceLogprobs) HasContent() bool`

HasContent returns a boolean if a field has been set.

### SetContentNil

`func (o *CCKChoiceLogprobs) SetContentNil(b bool)`

 SetContentNil sets the value for Content to be an explicit nil

### UnsetContent
`func (o *CCKChoiceLogprobs) UnsetContent()`

UnsetContent ensures that no value is present for Content, not even an explicit nil
### GetRefusal

`func (o *CCKChoiceLogprobs) GetRefusal() []CCKChatCompletionTokenLogprob`

GetRefusal returns the Refusal field if non-nil, zero value otherwise.

### GetRefusalOk

`func (o *CCKChoiceLogprobs) GetRefusalOk() (*[]CCKChatCompletionTokenLogprob, bool)`

GetRefusalOk returns a tuple with the Refusal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefusal

`func (o *CCKChoiceLogprobs) SetRefusal(v []CCKChatCompletionTokenLogprob)`

SetRefusal sets Refusal field to given value.

### HasRefusal

`func (o *CCKChoiceLogprobs) HasRefusal() bool`

HasRefusal returns a boolean if a field has been set.

### SetRefusalNil

`func (o *CCKChoiceLogprobs) SetRefusalNil(b bool)`

 SetRefusalNil sets the value for Refusal to be an explicit nil

### UnsetRefusal
`func (o *CCKChoiceLogprobs) UnsetRefusal()`

UnsetRefusal ensures that no value is present for Refusal, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


