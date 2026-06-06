# CCKChunkChoice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Delta** | [**CCKChoiceDelta**](CCKChoiceDelta.md) |  | 
**FinishReason** | Pointer to **NullableString** |  | [optional] 
**Index** | **int32** |  | 
**Logprobs** | Pointer to [**NullableCCKChoiceLogprobs**](CCKChoiceLogprobs.md) |  | [optional] 

## Methods

### NewCCKChunkChoice

`func NewCCKChunkChoice(delta CCKChoiceDelta, index int32, ) *CCKChunkChoice`

NewCCKChunkChoice instantiates a new CCKChunkChoice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCCKChunkChoiceWithDefaults

`func NewCCKChunkChoiceWithDefaults() *CCKChunkChoice`

NewCCKChunkChoiceWithDefaults instantiates a new CCKChunkChoice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDelta

`func (o *CCKChunkChoice) GetDelta() CCKChoiceDelta`

GetDelta returns the Delta field if non-nil, zero value otherwise.

### GetDeltaOk

`func (o *CCKChunkChoice) GetDeltaOk() (*CCKChoiceDelta, bool)`

GetDeltaOk returns a tuple with the Delta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelta

`func (o *CCKChunkChoice) SetDelta(v CCKChoiceDelta)`

SetDelta sets Delta field to given value.


### GetFinishReason

`func (o *CCKChunkChoice) GetFinishReason() string`

GetFinishReason returns the FinishReason field if non-nil, zero value otherwise.

### GetFinishReasonOk

`func (o *CCKChunkChoice) GetFinishReasonOk() (*string, bool)`

GetFinishReasonOk returns a tuple with the FinishReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinishReason

`func (o *CCKChunkChoice) SetFinishReason(v string)`

SetFinishReason sets FinishReason field to given value.

### HasFinishReason

`func (o *CCKChunkChoice) HasFinishReason() bool`

HasFinishReason returns a boolean if a field has been set.

### SetFinishReasonNil

`func (o *CCKChunkChoice) SetFinishReasonNil(b bool)`

 SetFinishReasonNil sets the value for FinishReason to be an explicit nil

### UnsetFinishReason
`func (o *CCKChunkChoice) UnsetFinishReason()`

UnsetFinishReason ensures that no value is present for FinishReason, not even an explicit nil
### GetIndex

`func (o *CCKChunkChoice) GetIndex() int32`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *CCKChunkChoice) GetIndexOk() (*int32, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *CCKChunkChoice) SetIndex(v int32)`

SetIndex sets Index field to given value.


### GetLogprobs

`func (o *CCKChunkChoice) GetLogprobs() CCKChoiceLogprobs`

GetLogprobs returns the Logprobs field if non-nil, zero value otherwise.

### GetLogprobsOk

`func (o *CCKChunkChoice) GetLogprobsOk() (*CCKChoiceLogprobs, bool)`

GetLogprobsOk returns a tuple with the Logprobs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogprobs

`func (o *CCKChunkChoice) SetLogprobs(v CCKChoiceLogprobs)`

SetLogprobs sets Logprobs field to given value.

### HasLogprobs

`func (o *CCKChunkChoice) HasLogprobs() bool`

HasLogprobs returns a boolean if a field has been set.

### SetLogprobsNil

`func (o *CCKChunkChoice) SetLogprobsNil(b bool)`

 SetLogprobsNil sets the value for Logprobs to be an explicit nil

### UnsetLogprobs
`func (o *CCKChunkChoice) UnsetLogprobs()`

UnsetLogprobs ensures that no value is present for Logprobs, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


