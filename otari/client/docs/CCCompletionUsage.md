# CCCompletionUsage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompletionTokens** | **int32** |  | 
**PromptTokens** | **int32** |  | 
**TotalTokens** | **int32** |  | 
**CompletionTokensDetails** | Pointer to [**NullableCCCompletionTokensDetails**](CCCompletionTokensDetails.md) |  | [optional] 
**PromptTokensDetails** | Pointer to [**NullableCCPromptTokensDetails**](CCPromptTokensDetails.md) |  | [optional] 
**InferenceGeo** | Pointer to **NullableString** |  | [optional] 
**ServerToolUse** | Pointer to [**NullableCCServerToolUsage**](CCServerToolUsage.md) |  | [optional] 

## Methods

### NewCCCompletionUsage

`func NewCCCompletionUsage(completionTokens int32, promptTokens int32, totalTokens int32, ) *CCCompletionUsage`

NewCCCompletionUsage instantiates a new CCCompletionUsage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCCCompletionUsageWithDefaults

`func NewCCCompletionUsageWithDefaults() *CCCompletionUsage`

NewCCCompletionUsageWithDefaults instantiates a new CCCompletionUsage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompletionTokens

`func (o *CCCompletionUsage) GetCompletionTokens() int32`

GetCompletionTokens returns the CompletionTokens field if non-nil, zero value otherwise.

### GetCompletionTokensOk

`func (o *CCCompletionUsage) GetCompletionTokensOk() (*int32, bool)`

GetCompletionTokensOk returns a tuple with the CompletionTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletionTokens

`func (o *CCCompletionUsage) SetCompletionTokens(v int32)`

SetCompletionTokens sets CompletionTokens field to given value.


### GetPromptTokens

`func (o *CCCompletionUsage) GetPromptTokens() int32`

GetPromptTokens returns the PromptTokens field if non-nil, zero value otherwise.

### GetPromptTokensOk

`func (o *CCCompletionUsage) GetPromptTokensOk() (*int32, bool)`

GetPromptTokensOk returns a tuple with the PromptTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromptTokens

`func (o *CCCompletionUsage) SetPromptTokens(v int32)`

SetPromptTokens sets PromptTokens field to given value.


### GetTotalTokens

`func (o *CCCompletionUsage) GetTotalTokens() int32`

GetTotalTokens returns the TotalTokens field if non-nil, zero value otherwise.

### GetTotalTokensOk

`func (o *CCCompletionUsage) GetTotalTokensOk() (*int32, bool)`

GetTotalTokensOk returns a tuple with the TotalTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTokens

`func (o *CCCompletionUsage) SetTotalTokens(v int32)`

SetTotalTokens sets TotalTokens field to given value.


### GetCompletionTokensDetails

`func (o *CCCompletionUsage) GetCompletionTokensDetails() CCCompletionTokensDetails`

GetCompletionTokensDetails returns the CompletionTokensDetails field if non-nil, zero value otherwise.

### GetCompletionTokensDetailsOk

`func (o *CCCompletionUsage) GetCompletionTokensDetailsOk() (*CCCompletionTokensDetails, bool)`

GetCompletionTokensDetailsOk returns a tuple with the CompletionTokensDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletionTokensDetails

`func (o *CCCompletionUsage) SetCompletionTokensDetails(v CCCompletionTokensDetails)`

SetCompletionTokensDetails sets CompletionTokensDetails field to given value.

### HasCompletionTokensDetails

`func (o *CCCompletionUsage) HasCompletionTokensDetails() bool`

HasCompletionTokensDetails returns a boolean if a field has been set.

### SetCompletionTokensDetailsNil

`func (o *CCCompletionUsage) SetCompletionTokensDetailsNil(b bool)`

 SetCompletionTokensDetailsNil sets the value for CompletionTokensDetails to be an explicit nil

### UnsetCompletionTokensDetails
`func (o *CCCompletionUsage) UnsetCompletionTokensDetails()`

UnsetCompletionTokensDetails ensures that no value is present for CompletionTokensDetails, not even an explicit nil
### GetPromptTokensDetails

`func (o *CCCompletionUsage) GetPromptTokensDetails() CCPromptTokensDetails`

GetPromptTokensDetails returns the PromptTokensDetails field if non-nil, zero value otherwise.

### GetPromptTokensDetailsOk

`func (o *CCCompletionUsage) GetPromptTokensDetailsOk() (*CCPromptTokensDetails, bool)`

GetPromptTokensDetailsOk returns a tuple with the PromptTokensDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromptTokensDetails

`func (o *CCCompletionUsage) SetPromptTokensDetails(v CCPromptTokensDetails)`

SetPromptTokensDetails sets PromptTokensDetails field to given value.

### HasPromptTokensDetails

`func (o *CCCompletionUsage) HasPromptTokensDetails() bool`

HasPromptTokensDetails returns a boolean if a field has been set.

### SetPromptTokensDetailsNil

`func (o *CCCompletionUsage) SetPromptTokensDetailsNil(b bool)`

 SetPromptTokensDetailsNil sets the value for PromptTokensDetails to be an explicit nil

### UnsetPromptTokensDetails
`func (o *CCCompletionUsage) UnsetPromptTokensDetails()`

UnsetPromptTokensDetails ensures that no value is present for PromptTokensDetails, not even an explicit nil
### GetInferenceGeo

`func (o *CCCompletionUsage) GetInferenceGeo() string`

GetInferenceGeo returns the InferenceGeo field if non-nil, zero value otherwise.

### GetInferenceGeoOk

`func (o *CCCompletionUsage) GetInferenceGeoOk() (*string, bool)`

GetInferenceGeoOk returns a tuple with the InferenceGeo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInferenceGeo

`func (o *CCCompletionUsage) SetInferenceGeo(v string)`

SetInferenceGeo sets InferenceGeo field to given value.

### HasInferenceGeo

`func (o *CCCompletionUsage) HasInferenceGeo() bool`

HasInferenceGeo returns a boolean if a field has been set.

### SetInferenceGeoNil

`func (o *CCCompletionUsage) SetInferenceGeoNil(b bool)`

 SetInferenceGeoNil sets the value for InferenceGeo to be an explicit nil

### UnsetInferenceGeo
`func (o *CCCompletionUsage) UnsetInferenceGeo()`

UnsetInferenceGeo ensures that no value is present for InferenceGeo, not even an explicit nil
### GetServerToolUse

`func (o *CCCompletionUsage) GetServerToolUse() CCServerToolUsage`

GetServerToolUse returns the ServerToolUse field if non-nil, zero value otherwise.

### GetServerToolUseOk

`func (o *CCCompletionUsage) GetServerToolUseOk() (*CCServerToolUsage, bool)`

GetServerToolUseOk returns a tuple with the ServerToolUse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerToolUse

`func (o *CCCompletionUsage) SetServerToolUse(v CCServerToolUsage)`

SetServerToolUse sets ServerToolUse field to given value.

### HasServerToolUse

`func (o *CCCompletionUsage) HasServerToolUse() bool`

HasServerToolUse returns a boolean if a field has been set.

### SetServerToolUseNil

`func (o *CCCompletionUsage) SetServerToolUseNil(b bool)`

 SetServerToolUseNil sets the value for ServerToolUse to be an explicit nil

### UnsetServerToolUse
`func (o *CCCompletionUsage) UnsetServerToolUse()`

UnsetServerToolUse ensures that no value is present for ServerToolUse, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


