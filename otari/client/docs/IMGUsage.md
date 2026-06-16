# IMGUsage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InputTokens** | **int32** |  | 
**InputTokensDetails** | [**IMGUsageInputTokensDetails**](IMGUsageInputTokensDetails.md) |  | 
**OutputTokens** | **int32** |  | 
**TotalTokens** | **int32** |  | 
**OutputTokensDetails** | Pointer to [**NullableIMGUsageOutputTokensDetails**](IMGUsageOutputTokensDetails.md) |  | [optional] 

## Methods

### NewIMGUsage

`func NewIMGUsage(inputTokens int32, inputTokensDetails IMGUsageInputTokensDetails, outputTokens int32, totalTokens int32, ) *IMGUsage`

NewIMGUsage instantiates a new IMGUsage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIMGUsageWithDefaults

`func NewIMGUsageWithDefaults() *IMGUsage`

NewIMGUsageWithDefaults instantiates a new IMGUsage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInputTokens

`func (o *IMGUsage) GetInputTokens() int32`

GetInputTokens returns the InputTokens field if non-nil, zero value otherwise.

### GetInputTokensOk

`func (o *IMGUsage) GetInputTokensOk() (*int32, bool)`

GetInputTokensOk returns a tuple with the InputTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputTokens

`func (o *IMGUsage) SetInputTokens(v int32)`

SetInputTokens sets InputTokens field to given value.


### GetInputTokensDetails

`func (o *IMGUsage) GetInputTokensDetails() IMGUsageInputTokensDetails`

GetInputTokensDetails returns the InputTokensDetails field if non-nil, zero value otherwise.

### GetInputTokensDetailsOk

`func (o *IMGUsage) GetInputTokensDetailsOk() (*IMGUsageInputTokensDetails, bool)`

GetInputTokensDetailsOk returns a tuple with the InputTokensDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputTokensDetails

`func (o *IMGUsage) SetInputTokensDetails(v IMGUsageInputTokensDetails)`

SetInputTokensDetails sets InputTokensDetails field to given value.


### GetOutputTokens

`func (o *IMGUsage) GetOutputTokens() int32`

GetOutputTokens returns the OutputTokens field if non-nil, zero value otherwise.

### GetOutputTokensOk

`func (o *IMGUsage) GetOutputTokensOk() (*int32, bool)`

GetOutputTokensOk returns a tuple with the OutputTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputTokens

`func (o *IMGUsage) SetOutputTokens(v int32)`

SetOutputTokens sets OutputTokens field to given value.


### GetTotalTokens

`func (o *IMGUsage) GetTotalTokens() int32`

GetTotalTokens returns the TotalTokens field if non-nil, zero value otherwise.

### GetTotalTokensOk

`func (o *IMGUsage) GetTotalTokensOk() (*int32, bool)`

GetTotalTokensOk returns a tuple with the TotalTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTokens

`func (o *IMGUsage) SetTotalTokens(v int32)`

SetTotalTokens sets TotalTokens field to given value.


### GetOutputTokensDetails

`func (o *IMGUsage) GetOutputTokensDetails() IMGUsageOutputTokensDetails`

GetOutputTokensDetails returns the OutputTokensDetails field if non-nil, zero value otherwise.

### GetOutputTokensDetailsOk

`func (o *IMGUsage) GetOutputTokensDetailsOk() (*IMGUsageOutputTokensDetails, bool)`

GetOutputTokensDetailsOk returns a tuple with the OutputTokensDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputTokensDetails

`func (o *IMGUsage) SetOutputTokensDetails(v IMGUsageOutputTokensDetails)`

SetOutputTokensDetails sets OutputTokensDetails field to given value.

### HasOutputTokensDetails

`func (o *IMGUsage) HasOutputTokensDetails() bool`

HasOutputTokensDetails returns a boolean if a field has been set.

### SetOutputTokensDetailsNil

`func (o *IMGUsage) SetOutputTokensDetailsNil(b bool)`

 SetOutputTokensDetailsNil sets the value for OutputTokensDetails to be an explicit nil

### UnsetOutputTokensDetails
`func (o *IMGUsage) UnsetOutputTokensDetails()`

UnsetOutputTokensDetails ensures that no value is present for OutputTokensDetails, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


