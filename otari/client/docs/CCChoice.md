# CCChoice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FinishReason** | **string** |  | 
**Index** | **int32** |  | 
**Logprobs** | Pointer to [**NullableCCChoiceLogprobs**](CCChoiceLogprobs.md) |  | [optional] 
**Message** | [**CCChatCompletionMessage**](CCChatCompletionMessage.md) |  | 

## Methods

### NewCCChoice

`func NewCCChoice(finishReason string, index int32, message CCChatCompletionMessage, ) *CCChoice`

NewCCChoice instantiates a new CCChoice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCCChoiceWithDefaults

`func NewCCChoiceWithDefaults() *CCChoice`

NewCCChoiceWithDefaults instantiates a new CCChoice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFinishReason

`func (o *CCChoice) GetFinishReason() string`

GetFinishReason returns the FinishReason field if non-nil, zero value otherwise.

### GetFinishReasonOk

`func (o *CCChoice) GetFinishReasonOk() (*string, bool)`

GetFinishReasonOk returns a tuple with the FinishReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinishReason

`func (o *CCChoice) SetFinishReason(v string)`

SetFinishReason sets FinishReason field to given value.


### GetIndex

`func (o *CCChoice) GetIndex() int32`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *CCChoice) GetIndexOk() (*int32, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *CCChoice) SetIndex(v int32)`

SetIndex sets Index field to given value.


### GetLogprobs

`func (o *CCChoice) GetLogprobs() CCChoiceLogprobs`

GetLogprobs returns the Logprobs field if non-nil, zero value otherwise.

### GetLogprobsOk

`func (o *CCChoice) GetLogprobsOk() (*CCChoiceLogprobs, bool)`

GetLogprobsOk returns a tuple with the Logprobs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogprobs

`func (o *CCChoice) SetLogprobs(v CCChoiceLogprobs)`

SetLogprobs sets Logprobs field to given value.

### HasLogprobs

`func (o *CCChoice) HasLogprobs() bool`

HasLogprobs returns a boolean if a field has been set.

### SetLogprobsNil

`func (o *CCChoice) SetLogprobsNil(b bool)`

 SetLogprobsNil sets the value for Logprobs to be an explicit nil

### UnsetLogprobs
`func (o *CCChoice) UnsetLogprobs()`

UnsetLogprobs ensures that no value is present for Logprobs, not even an explicit nil
### GetMessage

`func (o *CCChoice) GetMessage() CCChatCompletionMessage`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *CCChoice) GetMessageOk() (*CCChatCompletionMessage, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *CCChoice) SetMessage(v CCChatCompletionMessage)`

SetMessage sets Message field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


