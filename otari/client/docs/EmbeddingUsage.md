# EmbeddingUsage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CostUsd** | Pointer to **NullableString** | What the request cost, as a six-decimal USD string; absent when it was not priced | [optional] 
**PricingSource** | Pointer to **NullableString** | Which rate priced the request: &#x60;organization&#x60;, &#x60;deployment&#x60; or &#x60;defaults&#x60;; absent when it was not priced | [optional] 
**PromptTokens** | **int32** |  | 
**TotalTokens** | **int32** |  | 

## Methods

### NewEmbeddingUsage

`func NewEmbeddingUsage(promptTokens int32, totalTokens int32, ) *EmbeddingUsage`

NewEmbeddingUsage instantiates a new EmbeddingUsage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmbeddingUsageWithDefaults

`func NewEmbeddingUsageWithDefaults() *EmbeddingUsage`

NewEmbeddingUsageWithDefaults instantiates a new EmbeddingUsage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCostUsd

`func (o *EmbeddingUsage) GetCostUsd() string`

GetCostUsd returns the CostUsd field if non-nil, zero value otherwise.

### GetCostUsdOk

`func (o *EmbeddingUsage) GetCostUsdOk() (*string, bool)`

GetCostUsdOk returns a tuple with the CostUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostUsd

`func (o *EmbeddingUsage) SetCostUsd(v string)`

SetCostUsd sets CostUsd field to given value.

### HasCostUsd

`func (o *EmbeddingUsage) HasCostUsd() bool`

HasCostUsd returns a boolean if a field has been set.

### SetCostUsdNil

`func (o *EmbeddingUsage) SetCostUsdNil(b bool)`

 SetCostUsdNil sets the value for CostUsd to be an explicit nil

### UnsetCostUsd
`func (o *EmbeddingUsage) UnsetCostUsd()`

UnsetCostUsd ensures that no value is present for CostUsd, not even an explicit nil
### GetPricingSource

`func (o *EmbeddingUsage) GetPricingSource() string`

GetPricingSource returns the PricingSource field if non-nil, zero value otherwise.

### GetPricingSourceOk

`func (o *EmbeddingUsage) GetPricingSourceOk() (*string, bool)`

GetPricingSourceOk returns a tuple with the PricingSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPricingSource

`func (o *EmbeddingUsage) SetPricingSource(v string)`

SetPricingSource sets PricingSource field to given value.

### HasPricingSource

`func (o *EmbeddingUsage) HasPricingSource() bool`

HasPricingSource returns a boolean if a field has been set.

### SetPricingSourceNil

`func (o *EmbeddingUsage) SetPricingSourceNil(b bool)`

 SetPricingSourceNil sets the value for PricingSource to be an explicit nil

### UnsetPricingSource
`func (o *EmbeddingUsage) UnsetPricingSource()`

UnsetPricingSource ensures that no value is present for PricingSource, not even an explicit nil
### GetPromptTokens

`func (o *EmbeddingUsage) GetPromptTokens() int32`

GetPromptTokens returns the PromptTokens field if non-nil, zero value otherwise.

### GetPromptTokensOk

`func (o *EmbeddingUsage) GetPromptTokensOk() (*int32, bool)`

GetPromptTokensOk returns a tuple with the PromptTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromptTokens

`func (o *EmbeddingUsage) SetPromptTokens(v int32)`

SetPromptTokens sets PromptTokens field to given value.


### GetTotalTokens

`func (o *EmbeddingUsage) GetTotalTokens() int32`

GetTotalTokens returns the TotalTokens field if non-nil, zero value otherwise.

### GetTotalTokensOk

`func (o *EmbeddingUsage) GetTotalTokensOk() (*int32, bool)`

GetTotalTokensOk returns a tuple with the TotalTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTokens

`func (o *EmbeddingUsage) SetTotalTokens(v int32)`

SetTotalTokens sets TotalTokens field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


