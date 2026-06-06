# CCKCompletionUsage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompletionTokens** | **int32** |  | 
**PromptTokens** | **int32** |  | 
**TotalTokens** | **int32** |  | 
**CompletionTokensDetails** | Pointer to [**NullableCCKCompletionTokensDetails**](CCKCompletionTokensDetails.md) |  | [optional] 
**PromptTokensDetails** | Pointer to [**NullableCCKPromptTokensDetails**](CCKPromptTokensDetails.md) |  | [optional] 

## Methods

### NewCCKCompletionUsage

`func NewCCKCompletionUsage(completionTokens int32, promptTokens int32, totalTokens int32, ) *CCKCompletionUsage`

NewCCKCompletionUsage instantiates a new CCKCompletionUsage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCCKCompletionUsageWithDefaults

`func NewCCKCompletionUsageWithDefaults() *CCKCompletionUsage`

NewCCKCompletionUsageWithDefaults instantiates a new CCKCompletionUsage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompletionTokens

`func (o *CCKCompletionUsage) GetCompletionTokens() int32`

GetCompletionTokens returns the CompletionTokens field if non-nil, zero value otherwise.

### GetCompletionTokensOk

`func (o *CCKCompletionUsage) GetCompletionTokensOk() (*int32, bool)`

GetCompletionTokensOk returns a tuple with the CompletionTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletionTokens

`func (o *CCKCompletionUsage) SetCompletionTokens(v int32)`

SetCompletionTokens sets CompletionTokens field to given value.


### GetPromptTokens

`func (o *CCKCompletionUsage) GetPromptTokens() int32`

GetPromptTokens returns the PromptTokens field if non-nil, zero value otherwise.

### GetPromptTokensOk

`func (o *CCKCompletionUsage) GetPromptTokensOk() (*int32, bool)`

GetPromptTokensOk returns a tuple with the PromptTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromptTokens

`func (o *CCKCompletionUsage) SetPromptTokens(v int32)`

SetPromptTokens sets PromptTokens field to given value.


### GetTotalTokens

`func (o *CCKCompletionUsage) GetTotalTokens() int32`

GetTotalTokens returns the TotalTokens field if non-nil, zero value otherwise.

### GetTotalTokensOk

`func (o *CCKCompletionUsage) GetTotalTokensOk() (*int32, bool)`

GetTotalTokensOk returns a tuple with the TotalTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTokens

`func (o *CCKCompletionUsage) SetTotalTokens(v int32)`

SetTotalTokens sets TotalTokens field to given value.


### GetCompletionTokensDetails

`func (o *CCKCompletionUsage) GetCompletionTokensDetails() CCKCompletionTokensDetails`

GetCompletionTokensDetails returns the CompletionTokensDetails field if non-nil, zero value otherwise.

### GetCompletionTokensDetailsOk

`func (o *CCKCompletionUsage) GetCompletionTokensDetailsOk() (*CCKCompletionTokensDetails, bool)`

GetCompletionTokensDetailsOk returns a tuple with the CompletionTokensDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletionTokensDetails

`func (o *CCKCompletionUsage) SetCompletionTokensDetails(v CCKCompletionTokensDetails)`

SetCompletionTokensDetails sets CompletionTokensDetails field to given value.

### HasCompletionTokensDetails

`func (o *CCKCompletionUsage) HasCompletionTokensDetails() bool`

HasCompletionTokensDetails returns a boolean if a field has been set.

### SetCompletionTokensDetailsNil

`func (o *CCKCompletionUsage) SetCompletionTokensDetailsNil(b bool)`

 SetCompletionTokensDetailsNil sets the value for CompletionTokensDetails to be an explicit nil

### UnsetCompletionTokensDetails
`func (o *CCKCompletionUsage) UnsetCompletionTokensDetails()`

UnsetCompletionTokensDetails ensures that no value is present for CompletionTokensDetails, not even an explicit nil
### GetPromptTokensDetails

`func (o *CCKCompletionUsage) GetPromptTokensDetails() CCKPromptTokensDetails`

GetPromptTokensDetails returns the PromptTokensDetails field if non-nil, zero value otherwise.

### GetPromptTokensDetailsOk

`func (o *CCKCompletionUsage) GetPromptTokensDetailsOk() (*CCKPromptTokensDetails, bool)`

GetPromptTokensDetailsOk returns a tuple with the PromptTokensDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromptTokensDetails

`func (o *CCKCompletionUsage) SetPromptTokensDetails(v CCKPromptTokensDetails)`

SetPromptTokensDetails sets PromptTokensDetails field to given value.

### HasPromptTokensDetails

`func (o *CCKCompletionUsage) HasPromptTokensDetails() bool`

HasPromptTokensDetails returns a boolean if a field has been set.

### SetPromptTokensDetailsNil

`func (o *CCKCompletionUsage) SetPromptTokensDetailsNil(b bool)`

 SetPromptTokensDetailsNil sets the value for PromptTokensDetails to be an explicit nil

### UnsetPromptTokensDetails
`func (o *CCKCompletionUsage) UnsetPromptTokensDetails()`

UnsetPromptTokensDetails ensures that no value is present for PromptTokensDetails, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


